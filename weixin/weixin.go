package weixin

import (
	"fmt"
	"time"

	go_http "github.com/pefish/go-http"
	i_logger "github.com/pefish/go-interface/i-logger"
	"github.com/pkg/errors"
)

type WeiXinAgent struct {
	url    string
	logger i_logger.ILogger
}

func New(logger i_logger.ILogger, url string) *WeiXinAgent {
	return &WeiXinAgent{
		logger: logger,
		url:    url,
	}
}

func (i *WeiXinAgent) send(msg string) error {
	var httpResult struct {
		ErrCode uint64 `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}
	_, _, err := go_http.NewHttpRequester(
		go_http.WithLogger(i.logger),
		go_http.WithTimeout(5*time.Second),
	).PostForStruct(
		&go_http.RequestParams{
			Url: i.url,
			Params: map[string]interface{}{
				"msgtype": "text",
				"text": map[string]interface{}{
					"content":        msg,
					"mentioned_list": []string{"@all"},
				},
			},
		},
		&httpResult,
	)
	if err != nil {
		return err
	}
	if httpResult.ErrCode != 0 {
		return errors.Errorf(httpResult.ErrMsg)
	}

	return nil
}

func (i *WeiXinAgent) Info(msg string) error {
	return i.send(fmt.Sprintf("[INFO] %s", msg))
}

func (i *WeiXinAgent) Warn(msg string) error {
	return i.send(fmt.Sprintf("[WARN] %s", msg))
}

func (i *WeiXinAgent) Error(msg string) error {
	return i.send(fmt.Sprintf("[ERROR] %s", msg))
}