package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

func (s *GinAdapter) Login(ctx *gin.Context) {
	var req loginUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	sess, err := s.service.Login(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, sess)
}

type registerUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

func (s *GinAdapter) Register(ctx *gin.Context) {
	var req registerUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := s.service.Register(req.Email, req.Password, req.Name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
