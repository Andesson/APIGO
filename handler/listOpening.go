package handler

import (
	"net/http"

	"github.com/Andesson/APIGO/schemas"
	"github.com/gin-gonic/gin"
)

func ShowAllOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSucess(ctx, "list-opening", openings)
}
