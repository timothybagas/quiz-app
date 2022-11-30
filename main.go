package main

import (
	"log"

	"quiz-app/logics"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE")

		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(204)
			return
		}
		context.Next()
	})

	// returns a list of quizzes
	r.GET("/quizzes", logics.GetQuizzes)
	// returns a random quiz
	r.GET("/quiz", logics.GetRandomQuiz)
	// returns a result of a specific quiz answer
	r.GET("/answer/:quizID", logics.AnswerQuiz)
	// creates a new quiz
	r.POST("/quiz", logics.CreateQuiz)
	// removes a specific quiz
	r.DELETE("/quiz/:quizID", logics.DeleteQuiz)

	if err := r.Run(); err != nil {
		log.Fatalln("something went wrong:", err)
	}
}
