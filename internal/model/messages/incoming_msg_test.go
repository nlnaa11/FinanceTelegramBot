package messages

import (
	"testing"

	"github.com/golang/mock/gomock"
	mocks "github.com/nlnaa11/FinanceTelegramBot/internal/mocks/messages"
	"github.com/stretchr/testify/assert"
)

func Test_OnStartCommand_ShouldAnswerWithIntroMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockMessageSender(ctrl)
	model := New(sender)

	sender.EXPECT().SendMessage("hello", int64(123))

	err := model.IncomingMessage(Message{
		Text:   "/start",
		UserID: 123,
	})

	assert.NoError(t, err)
}

func Test_OnUnknownCommand_ShouldAnswerWithHelpMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockMessageSender(ctrl)
	model := New(sender)

	sender.EXPECT().SendMessage("unknown command", int64(123))

	err := model.IncomingMessage(Message{
		Text:   "/any",
		UserID: 123,
	})

	assert.NoError(t, err)
}
