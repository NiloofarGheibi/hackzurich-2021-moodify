package controllers

import (
	"net/http"
	"rest-api/internal/exception"
	"rest-api/internal/models"
	"rest-api/internal/services"
	"rest-api/internal/util"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserService services.User
}

func NewUserController(service services.User) *User {
	return &User{UserService: service}
}

func (u *User) GetAll(ctx *gin.Context) {
	offset := util.GetOffset(ctx.DefaultQuery("offset", "0"))
	limit := util.GetLimit(ctx.DefaultQuery("limit", "0"))

	user, err := u.UserService.Get(ctx.Request.Context(), offset, limit)
	if err != nil {
		exception.AsHTTPError(ctx, err)
		return
	}
	if user == nil {
		exception.AsHTTPError(ctx, exception.ErrNotFound.WithDetail("user"))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (u *User) GetByID(ctx *gin.Context) {
	userID := ctx.Param("userid")
	if userID == "" {
		exception.AsHTTPError(ctx, exception.ErrMissingInputParameters.WithDetail("userid"))
		return
	}

	user, err := u.UserService.GetByID(ctx.Request.Context(), userID)
	if err != nil {
		exception.AsHTTPError(ctx, err)
		return
	}
	if user == nil {
		exception.AsHTTPError(ctx, exception.ErrNotFound.WithDetail("user"))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (u *User) RegisterUser(ctx *gin.Context) {
	user := &models.User{}
	if err := ctx.ShouldBindJSON(user); err != nil {
		exception.AsHTTPError(ctx, exception.ErrInvalidInput.WithDetail(err.Error()))
		return
	}

	existingUser, err := u.UserService.GetByEmail(ctx.Request.Context(), user.Email)
	if err != nil {
		exception.AsHTTPError(ctx, err)
		return
	}

	if existingUser != nil {
		exception.AsHTTPError(ctx, exception.ErrConflict.WithDetail("user with this email exists already"))
		return
	}

	registeredUser, err := u.UserService.Store(ctx, user)
	if err != nil {
		exception.AsHTTPError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, registeredUser)
}

func (u *User) UpdateUser(ctx *gin.Context) {
	userID := ctx.Param("userid")
	if userID == "" {
		exception.AsHTTPError(ctx, exception.ErrMissingInputParameters.WithDetail("userid"))
		return
	}

	userUpdates := &models.User{}
	if err := ctx.ShouldBindJSON(userUpdates); err != nil {
		exception.AsHTTPError(ctx, exception.ErrInvalidInput.WithDetail(err.Error()))
		return
	}

	user, err := u.UserService.GetByID(ctx, userID)
	if err != nil {
		exception.AsHTTPError(ctx, err)
		return
	}

	if user == nil {
		exception.AsHTTPError(ctx, exception.ErrNotFound.WithDetail("user with id "+userID))
		return
	}

	// update user
	updatedUser, err := u.UserService.Update(ctx.Request.Context(), user, userUpdates)
	if err != nil {
		exception.AsHTTPError(ctx, exception.ErrOperationFailed.WithDetail("updating the user"))
		return
	}

	ctx.JSON(http.StatusOK, updatedUser)
}

func (u *User) DeleteUser(ctx *gin.Context) {
	userID := ctx.Param("userid")
	if userID == "" {
		exception.AsHTTPError(ctx, exception.ErrMissingInputParameters.WithDetail("userid"))
		return
	}

	userUpdates := &models.User{}
	if err := ctx.ShouldBindJSON(userUpdates); err != nil {
		exception.AsHTTPError(ctx, exception.ErrInvalidInput.WithDetail(err.Error()))
		return
	}

	user, err := u.UserService.GetByID(ctx, userID)
	if err != nil {
		exception.AsHTTPError(ctx, err)
		return
	}

	if user == nil {
		exception.AsHTTPError(ctx, exception.ErrNotFound.WithDetail("user with id "+userID))
		return
	}

	err = u.UserService.Delete(ctx.Request.Context(), user)
	if err != nil {
		exception.AsHTTPError(ctx, exception.ErrOperationFailed.WithDetail("deleting the user"))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
