package logics

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetQuizzes(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"data": quizzes,
	})
}
