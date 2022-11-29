package logics

import (
	"net/http"
	"regexp"

	"quiz-app/types"

	"github.com/gin-gonic/gin"
)

func CreateQuiz(context *gin.Context) {
	var (
		quiz = struct {
			Question string `json:"question"`
			Answers  []struct {
				IsCorrect bool   `json:"isCorrect"`
				Key       string `json:"key"`
				Value     string `json:"value"`
			} `json:"answers"`
		}{}
		err = context.ShouldBindJSON(&quiz)
	)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if quiz.Question == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "please fill in the question field",
		})
		return
	}
	if len(quiz.Answers) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "at least 1 answer must be provided to create a new quiz",
		})
		return
	}

	var (
		regex, _      = regexp.Compile("^[A-Z]$")
		correctAnswer = 0
		answers       = make([]types.Answer, 0)
		usedKey       = make([]bool, 26)
	)
	for _, answer := range quiz.Answers {
		if !regex.MatchString(answer.Key) {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "please provide a valid answer key",
			})
			return
		}
		if index := int(answer.Key[0] - 'A'); usedKey[index] {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "answer key must be unique",
			})
			return
		} else {
			usedKey[index] = true
		}
		if answer.Value == "" {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "answer cannot be empty",
			})
			return
		}
		if answer.IsCorrect {
			correctAnswer++
		}
		answers = append(answers, types.Answer{
			Key:       answer.Key,
			Value:     answer.Value,
			IsCorrect: answer.IsCorrect,
		})
	}
	if correctAnswer != 1 {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "only 1 answer should be a correct answer",
		})
		return
	}

	quizzes = append(quizzes, types.Quiz{
		Question: quiz.Question,
		Answers:  answers,
	})
	context.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"message": "successfully create a new quiz",
		},
	})
}
