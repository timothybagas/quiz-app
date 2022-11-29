package main

import (
	"log"

	"quiz-app/logics"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// returns a random quiz
	r.GET("/quiz", logics.GetRandomQuiz)
	// returns a result of a specific quiz answer
	r.GET("/answer/:quizID", logics.AnswerQuiz)
	// creates a new quiz
	r.POST("/quiz", logics.CreateQuiz)

	if err := r.Run(); err != nil {
		log.Fatalln("something went wrong:", err)
	}
}
