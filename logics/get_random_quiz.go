package logics

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRandomQuiz(context *gin.Context) {
	index := rand.Intn(len(quizzes))
	context.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":   index + 1,
			"quiz": quizzes[index],
		},
	})
}
