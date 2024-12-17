package tg

import (
	"fmt"

	i_logger "github.com/pefish/go-interface/i-logger"
	tg_sender "github.com/pefish/tg-sender"
)

type TgAgent struct {
	token    string
	tgSender *tg_sender.TgSender
	groupId  string
}

func New(logger i_logger.ILogger, token string, groupId string) *TgAgent {
	return &TgAgent{
		token:    token,
		tgSender: tg_sender.NewTgSender(logger, token),
		groupId:  groupId,
	}
}

func (i *TgAgent) Info(msg string) error {
	return i.tgSender.SendMsg(&tg_sender.MsgStruct{
		ChatId: i.groupId,
		Msg:    fmt.Sprintf("[INFO] %s", msg),
		Ats:    nil,
	}, 0)
}

func (i *TgAgent) Warn(msg string) error {
	return i.tgSender.SendMsg(&tg_sender.MsgStruct{
		ChatId: i.groupId,
		Msg:    fmt.Sprintf("[WARN] %s", msg),
		Ats:    nil,
	}, 0)
}

func (i *TgAgent) Error(msg string) error {
	return i.tgSender.SendMsg(&tg_sender.MsgStruct{
		ChatId: i.groupId,
		Msg:    fmt.Sprintf("[ERROR] %s", msg),
		Ats:    nil,
	}, 0)
}
