package tg_alert

import (
	"fmt"
	"time"

	i_logger "github.com/pefish/go-interface/i-logger"
	tg_sender "github.com/pefish/tg-sender"
)

type TgAgent struct {
	token    string
	tgSender *tg_sender.TgSender
	groupId  string
	interval time.Duration
}

func New(
	logger i_logger.ILogger,
	token string,
	groupId string,
	interval time.Duration,
) *TgAgent {
	return &TgAgent{
		token:    token,
		tgSender: tg_sender.NewTgSender(logger, token),
		groupId:  groupId,
		interval: interval,
	}
}

func (i *TgAgent) Infof(format string, a ...any) error {
	msg := fmt.Sprintf(format, a...)
	return i.tgSender.SendMsg(&tg_sender.MsgStruct{
		ChatId: i.groupId,
		Msg:    fmt.Sprintf("[INFO] %s", msg),
		Ats:    nil,
	}, i.interval)
}

func (i *TgAgent) Warnf(format string, a ...any) error {
	msg := fmt.Sprintf(format, a...)
	return i.tgSender.SendMsg(&tg_sender.MsgStruct{
		ChatId: i.groupId,
		Msg:    fmt.Sprintf("[WARN] %s", msg),
		Ats:    nil,
	}, i.interval)
}

func (i *TgAgent) Errorf(format string, a ...any) error {
	msg := fmt.Sprintf(format, a...)
	return i.tgSender.SendMsg(&tg_sender.MsgStruct{
		ChatId: i.groupId,
		Msg:    fmt.Sprintf("[ERROR] %s", msg),
		Ats:    nil,
	}, i.interval)
}
