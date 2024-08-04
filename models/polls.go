package models

type Polls struct {
	Id        int    `json:"id"`
	Question  string `json:"question"`
	AnswerA   string `json:"answer_a"`
	AnswerB   string `json:"answer_b"`
	CountA    int    `json:"count_a"`
	CountB    int    `json:"count_b"`
	CreatedAt string `json:"created_at"`
}
