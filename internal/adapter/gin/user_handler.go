package gin

import (
	"net/http"

	dmuser "github.com/fbriansyah/bank-ina-test/internal/application/domain/user"
	"github.com/gin-gonic/gin"
)

func (s *GinAdapter) ListUsers(ctx *gin.Context) {
	listUsers, err := s.service.ListUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, listUsers)
}

func (s *GinAdapter) GetUserByID(ctx *gin.Context) {
	var req getUserByIDRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := s.service.GetUserByID(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type updateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func (s *GinAdapter) UpdateUser(ctx *gin.Context) {
	var reqUri getUserByIDRequest
	var req updateUserRequest

	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userUpdate := dmuser.User{
		Name:  req.Name,
		Email: req.Email,
	}

	user, err := s.service.UpdateUser(reqUri.ID, userUpdate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}