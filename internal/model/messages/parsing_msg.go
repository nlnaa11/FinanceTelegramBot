package messages

import (
	"strconv"
	"time"

	"github.com/nlnaa11/FinanceTelegramBot/internal/model/data"
	"github.com/pkg/errors"
)

func parseAdd(request []string) (*data.Expense, error) {
	idx := 1
	end := len(request)

	expense, err := data.NewExpense()
	if err != nil {
		return nil, errors.Wrap(err, "adding new expenses: ")
	}

	for idx != end {
		key := request[idx]
		idx++
		if idx > len(request)-1 {
			return nil,
				errors.New("The parameter isn't set for the key: " + key)
		}

		switch key {
		case "-c":
			expense.Category = request[idx]
		case "-s":
			{
				sum, err := strconv.ParseFloat(request[idx], 64)
				if err != nil {
					return nil,
						errors.New("The sum must be a digit")
				}
				expense.Sum = sum
			}
		case "-d":
			{
				data, err := time.Parse("01-02-2006", request[idx])
				if err != nil {
					return nil,
						errors.New("Failed to parse a date (format: MM-DD-YYYY")
				}
				expense.Data = data
			}
		case "-cmt":
			expense.Comment = request[idx]
		default:
			return nil,
				errors.New("Couldn't find a key for adding command: ")
		}
		idx++
	}

	if expense.Sum <= 0.0 || expense.Category == "" {
		return nil,
			errors.New("Not enough data to create a new expense")
	}

	return expense, nil
}

// func parseShow(request []string) (string, error) {
// 	return "", nil
// }
