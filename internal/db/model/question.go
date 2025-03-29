package model

import "gorm.io/gorm"

type Question struct {
	ID         int    `gorm:"primaryKey; autoIncrement" json:"id"`
	QuestionID int    `gorm:"type:integer; not null" json:"question_id"`
	Response   string `gorm:"type:varchar(255); not null" json:"response"`
}

func AddQuestion(db *gorm.DB, questionID int, response string) (*Question, error) {
	question := Question{
		QuestionID: questionID,
		Response:   response,
	}
	err := db.Create(&question).Error
	if err != nil {
		return nil, err
	}
	return &question, nil
}
