package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"poll-api/models"
)

type PollHandler interface {
	GetPolls() ([]models.Polls, error)
	GetPollsById(pollsId int) (models.Polls, error)
	CreatePolls(body models.Polls) error
	UpdatePolls(pollsId int, body models.Polls) error
	UpdatePollsVote(pollsId int, option string) error
	DeletePolls(pollsId int) error
}

type PollHandlerImpl struct {
	DB *sql.DB
}

func NewPollHandler(db *sql.DB) PollHandler {
	return &PollHandlerImpl{
		DB: db,
	}
}

func (ph *PollHandlerImpl) GetPolls() ([]models.Polls, error) {
	query := "SELECT id, question, answer_a, answer_b, created_at, count_a, count_b FROM polls ORDER BY id ASC;"
	rows, err := ph.DB.Query(query)

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

func (ph *PollHandlerImpl) GetPollsById(pollsId int) (models.Polls, error) {
	var poll models.Polls
	query := "SELECT * FROM polls WHERE id=$1 LIMIT 1;"
	err := ph.DB.QueryRow(query, pollsId).Scan(
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

func (ph *PollHandlerImpl) CreatePolls(body models.Polls) error {
	query := `
		INSERT INTO polls (question, answer_a, answer_b)
		VALUES ($1, $2, $3)
	`

	_, err := ph.DB.Exec(query, body.Question, body.AnswerA, body.AnswerB)
	if err != nil {
		log.Printf("Error creating polls: %v", err)
		return err
	}

	return nil
}

func (ph *PollHandlerImpl) UpdatePolls(pollsId int, body models.Polls) error {
	query := `
		UPDATE polls
		SET question=$1, answer_a=$2, answer_b=$3, count_a=0, count_b=0
		WHERE id=$4
	`

	_, err := ph.DB.Exec(query, body.Question, body.AnswerA, body.AnswerB, pollsId)
	if err != nil {
		log.Printf("Error updating polls: %v", err)
		return err
	}

	return nil
}

func (ph *PollHandlerImpl) UpdatePollsVote(pollsId int, option string) error {
	query := fmt.Sprintf(`
		UPDATE polls 
		SET count_%s = count_%s + 1
		WHERE id=$1;
	`, option, option)

	_, err := ph.DB.Exec(query, pollsId)

	if err != nil {
		log.Printf("Error updating polls's vote: %v", err)
		return err
	}
	return nil
}

func (ph *PollHandlerImpl) DeletePolls(pollsId int) error {
	query := `
		DELETE FROM polls WHERE id=$1;
	`

	result, err := ph.DB.Exec(query, pollsId)
	if err != nil {
		log.Printf("Error deleting polls: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Database or database driver do not support RowsAffected().")
		return errors.New("database or database driver do not support RowsAffected()")
	}
	if rowsAffected == 0 {
		log.Printf("Error, deleting polls: %v", err)
		return errors.New("no data deleted")
	}
	return nil
}
