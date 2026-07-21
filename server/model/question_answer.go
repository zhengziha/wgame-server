package model

// QuestionAnswer 问答表（文曲星答题等）。
// type 为 SQL 保留字，column tag 直接写列名，GORM 会自动加反引号。
type QuestionAnswer struct {
	// ID 主键
	ID int32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`

	// No 题号
	No int32 `gorm:"column:no" json:"no"`

	// Question 题目
	Question string `gorm:"size:255;column:question" json:"question"`

	// Answer 答案
	Answer string `gorm:"type:text;column:answer" json:"answer"`

	// Type 题型（列名 type 为 SQL 保留字）
	Type string `gorm:"size:255;column:type" json:"type"`

	// Options 选项
	Options string `gorm:"size:255;column:options" json:"options"`
}

// TableName 显式指定表名
func (QuestionAnswer) TableName() string {
	return "question_answer"
}
