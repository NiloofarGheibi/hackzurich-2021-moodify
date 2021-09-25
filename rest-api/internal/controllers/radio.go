package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/internal/exception"
	"rest-api/internal/models"
	"rest-api/internal/services"
)

type Radio struct {
	MusicService services.Music
}

func NewRadioController(service services.Music) *Radio {
	return &Radio{MusicService: service}
}

func (r *Radio) RegisterTrack(ctx *gin.Context) {
	track := &models.Track{}
	if err := ctx.ShouldBindJSON(track); err != nil {
		exception.AsHTTPError(ctx, exception.ErrInvalidInput.WithDetail(err.Error()))
		return
	}

	registeredTrack, err := r.MusicService.Store(ctx, track)
	if err != nil {
		exception.AsHTTPError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, registeredTrack)
}
