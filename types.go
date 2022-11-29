package main

type Quiz struct {
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
}

type Answer struct {
	IsCorrect bool   `json:"isCorrect"`
	Value     string `json:"value"`
}
