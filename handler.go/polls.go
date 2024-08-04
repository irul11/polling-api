package handler

import (
	"fmt"
	"log"
	"poll-api/database"
	"poll-api/models"
)

func GetPolls() ([]models.Polls, error) {
	query := "SELECT * FROM polls;"
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
