package logics

import (
	"net/http"
	"strings"

	"quiz-app/types"

	"github.com/gin-gonic/gin"
)

func GetQuizzes(context *gin.Context) {
	var (
		search = context.Query("search")
		list   = make([]struct {
			ID   int        `json:"id"`
			Quiz types.Quiz `json:"quiz"`
		}, 0)
	)
	for id, quiz := range quizzes {
		if strings.Contains(strings.ToLower(quiz.Question), search) {
			list = append(list, struct {
				ID   int        `json:"id"`
				Quiz types.Quiz `json:"quiz"`
			}{ID: id + 1, Quiz: quiz})
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"data": list,
	})
}
