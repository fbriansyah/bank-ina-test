package application

import (
	"context"

	db "github.com/fbriansyah/bank-ina-test/internal/adapter/database"
	dmuser "github.com/fbriansyah/bank-ina-test/internal/application/domain/user"
)

func (s *Service) ListUsers() ([]dmuser.User, error) {
	// for improvement, need to add limit and offset for paging
	users, err := s.db.GetAllUser(context.Background())
	if err != nil {
		return []dmuser.User{}, err
	}

	var listUsers []dmuser.User
	for _, user := range users {
		listUsers = append(listUsers, dmuser.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return listUsers, nil
}

func (s *Service) GetUserByID(id int32) (dmuser.User, error) {
	user, err := s.db.GetUserByID(context.Background(), id)
	if err != nil {
		return dmuser.User{}, err
	}

	return dmuser.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *Service) UpdateUser(id int32, user dmuser.User) (dmuser.User, error) {
	usr, err := s.db.GetUserByID(context.Background(), id)
	if err != nil {
		return dmuser.User{}, err
	}

	_, err = s.db.UpdateUser(context.Background(), db.UpdateUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Password: usr.Password,
		ID:       id,
	})

	user.ID = id

	return user, nil
}
func (s *Service) DeleteUser(id int32) error {
	return s.db.DeleteUser(context.Background(), id)
}
