package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowAllOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "LIST Opening",
	})
}
