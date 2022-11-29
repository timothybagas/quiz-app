package logics

import (
	. "quiz-app/types"
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
