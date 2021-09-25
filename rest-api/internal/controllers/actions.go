package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/internal/models"
)

type Action struct {
}

func NewActionController() *Action {
	return &Action{}
}

func (r *Action) GetActionRecommendations(ctx *gin.Context) {
	actions := []models.Action{
		{"Setup a 1-1 with your team members", "sad"},
		{"What about a team event? To spice things up", "calm"},
		{"Keep up the good work and the vibe!", "energetic"},
		{"Celebrate small things, it's important to show your team member's value", "happy"}}

	ctx.JSON(http.StatusOK, actions)
}
