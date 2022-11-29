package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	quizzes = []Quiz{
		{
			Question: "Apa nama ibu kota Indonesia?",
			Answers: []Answer{
				{
					IsCorrect: true,
					Value:     "Jakarta",
				},
				{
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
		context.JSON(http.StatusOK, gin.H{
			"data": quizzes,
		})
	})

	if err := r.Run(); err != nil {
		log.Fatalln("something went wrong:", err)
	}
}
