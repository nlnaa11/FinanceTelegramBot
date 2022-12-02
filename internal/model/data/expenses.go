package data

import (
	"fmt"
	"time"
)

type Expense struct {
	Category string
	Sum      float64
	Data     time.Time
	Comment  string
}

func NewExpense() (*Expense, error) {
	return &Expense{
		Sum:     0.0,
		Data:    time.Now(),
		Comment: "",
	}, nil
}

func (e *Expense) ToString() string {
	result := fmt.Sprintf("cat:\t%s, sum:\t%v, data:\t%s",
		e.Category, e.Sum, e.Data.Format("02-01-2006"))
	if e.Comment != "" {
		result += ", comment:\t" + e.Comment
	}
	return result
}
