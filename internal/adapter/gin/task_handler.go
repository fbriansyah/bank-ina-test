package gin

import (
	"fmt"
	"net/http"

	dmtask "github.com/fbriansyah/bank-ina-test/internal/application/domain/task"
	dmtoken "github.com/fbriansyah/bank-ina-test/internal/application/domain/token"
	"github.com/gin-gonic/gin"
)

type createTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (s *GinAdapter) CreateTask(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*dmtoken.Payload)

	var req createTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	task, err := s.service.CreateTask(authPayload.UserID, req.Title, req.Description)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (s *GinAdapter) ListTaks(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*dmtoken.Payload)

	tasks, err := s.service.ListTasks(authPayload.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if tasks == nil {
		ctx.JSON(http.StatusOK, []dmtask.Task{})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (s *GinAdapter) GetTask(ctx *gin.Context) {
	var reqUri getUserByIDRequest
	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	task, err := s.service.GetTaskByID(reqUri.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, task)
}

type updateTaskRequest struct {
	UserID      int32  `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func (s *GinAdapter) UpdateTask(ctx *gin.Context) {
	var reqUri getUserByIDRequest
	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var req updateTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	task, err := s.service.UpdateTask(dmtask.Task{
		ID:          reqUri.ID,
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (s *GinAdapter) DeleteTask(ctx *gin.Context) {
	var reqUri getUserByIDRequest
	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("success delete task with id %d", reqUri.ID),
	})
}
