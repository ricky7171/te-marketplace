package error_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricky7171/te-marketplace/internal/library_wrapper"
)

func HandleHttpError(ctx library_wrapper.MyGinContext, message string) {
	if r := recover(); r != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": message})
	}
}
