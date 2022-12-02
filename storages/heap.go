package storages

import (
	"github.com/nlnaa11/FinanceTelegramBot/internal/model/data"
)

type Heap struct {
	expenses map[string][]data.Expense
}

func New() *Heap {
	return &Heap{
		expenses: make(map[string][]data.Expense),
	}
}

func (h *Heap) AddExpense(e data.Expense) (string, error) {
	h.expenses[e.Category] = append(h.expenses[e.Category], e)

	result := e.ToString()

	return result, nil
}

func (h *Heap) ShowExpenses() (string, error) {
	var result string
	for cat, expenses := range h.expenses {
		for _, e := range expenses {
			result = cat + e.ToString()
		}
	}
	return result, nil
}
