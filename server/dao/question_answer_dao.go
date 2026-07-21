package dao

import (
	"wgame-server/server/db"
	"wgame-server/server/model"
)

type QuestionAnswerDAO = BaseDAO[model.QuestionAnswer]

func NewQuestionAnswerDAO() *QuestionAnswerDAO {
	return NewBaseDAO[model.QuestionAnswer](
		db.GORM(),
		db.Cache(),
		"question_answer",
		func(t *model.QuestionAnswer) int64 {
			return int64(t.ID)
		},
	)
}
