package logics

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteQuiz(context *gin.Context) {
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

	quizzes = append(quizzes[:index], quizzes[index+1:]...)
	context.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"message": "successfully delete a quiz",
		},
	})
}
