package main

import (
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	quizzes = []Quiz{
		{
			Question: "Apa nama ibu kota Indonesia?",
			Answers: []Answer{
				{
					Key:       "A",
					IsCorrect: true,
					Value:     "Jakarta",
				},
				{
					Key:       "B",
					IsCorrect: false,
					Value:     "Semarang",
				},
			},
		},
	}
)

func main() {
	r := gin.Default()

	r.GET("/quiz", func(context *gin.Context) {
		index := rand.Intn(len(quizzes))
		context.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"id":   index + 1,
				"quiz": quizzes[index],
			},
		})
	})

	r.GET("/answer/:quizID", func(context *gin.Context) {
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
	})

	if err := r.Run(); err != nil {
		log.Fatalln("something went wrong:", err)
	}
}
