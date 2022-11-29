package logics

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRandomQuiz(context *gin.Context) {
	var (
		index int
		quiz  interface{}
	)
	if len(quizzes) > 0 {
		index = rand.Intn(len(quizzes)) + 1
		quiz = quizzes[index-1]
	}
	context.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":   index,
			"quiz": quiz,
		},
	})
}
