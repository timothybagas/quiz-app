package logics

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AnswerQuiz(context *gin.Context) {
	var (
		quizID     = context.Param("quizID")
		index, err = strconv.ParseInt(quizID, 10, 32)
	)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	index--
	if index < 0 || int(index) >= len(quizzes) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid quiz id",
		})
		return
	}

	var (
		quiz     = quizzes[index]
		answer   = context.Query("key")
		regex, _ = regexp.Compile("^[A-Z]$")
	)
	if !regex.MatchString(answer) || int(answer[0]-'A') >= len(quiz.Answers) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "please provide a valid answer key",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"result": quiz.Answers[int(answer[0]-'A')].IsCorrect,
		},
	})
}
