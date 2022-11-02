package handler

import (
	models "Server/model"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	ctx context.Context
}

func NewHandler(context context.Context) *Handler {
	return &Handler{
		context,
	}
}

func (h *Handler) BioHandler(c *gin.Context) {
	myProfile := models.NewBioProfile()
	myProfile.CreateMyBio()
	c.IndentedJSON(http.StatusOK, myProfile)
	return
}
