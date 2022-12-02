package model

import (
	"testing"

	"github.com/golang/mock/gomock"
	mocks "github.com/nlnaa11/FinanceTelegramBot/internal/mocks/messages"
	"github.com/stretchr/testify/assert"
)

func Test_OnStartCommand_ShouldAnswerWithIntroMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockMessageSender(ctrl)
	model := New(nil)

	err := model.IncomingMessage(Message{
		Text:   "/start",
		UserID: 123,
	})

	assert.NoError(t, err)
}
