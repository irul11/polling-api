package handler

import (
	"fmt"
	"log"
	"poll-api/database"
	"poll-api/models"
)

func GetPolls() ([]models.Polls, error) {
	query := "SELECT * FROM polls ORDER BY id ASC;"
	rows, err := database.DB.Query(query)

	if err != nil {
		log.Printf("Error querying polls: %v", err)
		return nil, err
	}

	polls := []models.Polls{}
	for rows.Next() {
		var poll models.Polls
		err := rows.Scan(
			&poll.Id,
			&poll.Question,
			&poll.AnswerA,
			&poll.AnswerB,
			&poll.CreatedAt,
			&poll.CountA,
			&poll.CountB,
		)

		if err != nil {
			log.Printf("Error scanning polls: %v", err)
			return nil, err
		}

		polls = append(polls, poll)
	}
	rows.Close()

	return polls, nil
}

func GetPollsById(pollsId int) (models.Polls, error) {
	var poll models.Polls
	query := "SELECT * FROM polls WHERE id=$1"
	err := database.DB.QueryRow(query, pollsId).Scan(
		&poll.Id,
		&poll.Question,
		&poll.AnswerA,
		&poll.AnswerB,
		&poll.CreatedAt,
		&poll.CountA,
		&poll.CountB,
	)

	if err != nil {
		log.Printf("Error querying polls: %v", err)
		return models.Polls{}, err
	}

	return poll, nil
}

func CreatePolls(body models.Polls) error {
	query := `
		INSERT INTO polls (question, answer_a, answer_b)
		VALUES ($1, $2, $3)
	`

	_, err := database.DB.Exec(query, body.Question, body.AnswerA, body.AnswerB)
	if err != nil {
		log.Printf("Error creating polls: %v", err)
		return err
	}

	return nil
}

func UpdatePolls(pollsId int, body models.Polls) error {
	query := `
		UPDATE polls
		SET question=$1, answer_a=$2, answer_b=$3, count_a=0, count_b=0
		WHERE id=$4
	`

	_, err := database.DB.Exec(query, body.Question, body.AnswerA, body.AnswerB, pollsId)
	if err != nil {
		log.Printf("Error updating polls: %v", err)
		return err
	}

	return nil
}

func UpdatePollsVote(pollsId int, option string) error {
	query := fmt.Sprintf(`
		UPDATE polls 
		SET count_%s = count_%s + 1
		WHERE id=$1;
	`, option, option)

	_, err := database.DB.Exec(query, pollsId)

	if err != nil {
		log.Printf("Error updating polls's vote: %v", err)
		return err
	}
	return nil
}

func DeletePolls(pollsId int) error {
	query := `
		DELETE FROM polls WHERE id=$1;
	`

	_, err := database.DB.Exec(query, pollsId)
	if err != nil {
		log.Printf("Error deleting polls: %v", err)
		return err
	}
	return nil
}
