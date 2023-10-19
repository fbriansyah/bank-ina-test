package application

import (
	"context"
	"errors"
	"time"

	db "github.com/fbriansyah/bank-ina-test/internal/adapter/database"
	dmsession "github.com/fbriansyah/bank-ina-test/internal/application/domain/session"
	dmtoken "github.com/fbriansyah/bank-ina-test/internal/application/domain/token"
	dmuser "github.com/fbriansyah/bank-ina-test/internal/application/domain/user"
	"github.com/fbriansyah/bank-ina-test/util"
)

var (
	ErrorInvalidEmail    = errors.New("invalid email address")
	ErrorInvalidPassword = errors.New("invalid password")
)

func (s *Service) Login(email, password string) (dmsession.Session, error) {
	user, err := s.db.GetUserByEmail(context.Background(), email)
	if err != nil {
		return dmsession.Session{}, ErrorInvalidEmail
	}

	if err = util.CheckPassword(password, user.Password); err != nil {
		return dmsession.Session{}, ErrorInvalidPassword
	}

	access_token, payload, err := s.tokenMaker.CreateToken("access", user.ID, time.Hour*24)
	if err != nil {
		return dmsession.Session{}, err
	}

	return dmsession.Session{
		ID:                   payload.ID.String(),
		UserID:               user.ID,
		AccessToken:          access_token,
		AccessTokenExpiresAt: payload.ExpiredAt,
	}, nil
}

func (s *Service) CheckToken(token string) (*dmtoken.Payload, error) {
	return s.tokenMaker.VerifyToken(token)
}

func (s *Service) Register(email, password, name string) (dmuser.User, error) {
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return dmuser.User{}, err
	}

	user, err := s.db.CreateUser(context.Background(), db.CreateUserParams{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	})
	if err != nil {
		return dmuser.User{}, err
	}

	return dmuser.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
