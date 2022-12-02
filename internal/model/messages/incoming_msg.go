package messages

import (
	"strings"

	"github.com/nlnaa11/FinanceTelegramBot/internal/model/data"
	"github.com/nlnaa11/FinanceTelegramBot/storages"
	"github.com/pkg/errors"
)

var (
	ruleForAddRequest string = "To add new expenses:\n" +
		"/add -c <category>* -s <sum>* -d <data> -cmt <coment>"

	ruleForShowRequest string = "To show expenses for a period:\n" +
		"/show -p <week/month/year/all>* <specific_period> -c <category>"
)

type MessageSender interface {
	SendMessage(text string, userID int64) error
}

type StorageHandler interface {
	AddExpenses(e data.Expense) (string, error)
	ShowExpenses()
}

type Model struct {
	tgClient MessageSender
}

func New(tgClient MessageSender) *Model {
	return &Model{
		tgClient: tgClient,
	}
}

type Message struct {
	Text   string
	UserID int64
}

func (s *Model) IncomingMessage(msg Message) error {
	if len(msg.Text) == 0 {
		return s.tgClient.SendMessage("unknown command", msg.UserID)
	}
	if msg.Text == "/start" {
		return s.tgClient.SendMessage("hello", msg.UserID)
	}

	paramsOfRequest := strings.Split(msg.Text, " ")

	stor := storages.New()

	if paramsOfRequest[0] == "/add" {
		if len(paramsOfRequest) < 5 {
			return s.tgClient.SendMessage(ruleForAddRequest, msg.UserID)
		}
		expense, err := parseAdd(paramsOfRequest)
		if err != nil || expense == nil {
			return s.tgClient.SendMessage(err.Error(), msg.UserID)
			//return errors.Wrap(err, "parsing adding a expense: ")
		}
		result, err := stor.AddExpense(*expense)
		if err != nil {
			return s.tgClient.SendMessage(err.Error(), msg.UserID)
			//return errors.Wrap(err, "adding new expense: ")
		}
		return s.tgClient.SendMessage(result, msg.UserID)
	}

	if paramsOfRequest[0] == "/show" {
		if len(paramsOfRequest) < 3 {
			return s.tgClient.SendMessage(ruleForShowRequest, msg.UserID)
		}
		// foo, err := parseShow(paramsOfRequest)
		// if err != nil {
		// 	return errors.Wrap(err, "print expenses: ")
		// }
		result, err := stor.ShowExpenses()
		if err != nil {
			return errors.Wrap(err, "showing expenses: ")
		}
		return s.tgClient.SendMessage(result, msg.UserID)
	}

	return s.tgClient.SendMessage("unknown command", msg.UserID)
}
