package server

import (
	"errors"
	"github.com/sirodeneko/QQBOT/coolq"
	"github.com/sirodeneko/QQBOT/util/config"
)

type Client struct {
	HttpClient *HttpClient
	WsClient   *WsClient
}

func NewClient(config *config.Config, bot *coolq.QQBoT) (*Client, error) {

	if config == nil {
		return nil, errors.New("配置为nil")
	}

	client := &Client{}
	var err error
	client.WsClient, err = newWebSocketClient(config.WsUrl, config.Token, bot)
	if err != nil {
		return nil, err
	}

	client.HttpClient = newHttpClient(config.HttpUrl, config.Token, bot)
	return client, nil
}

func (c *Client) WsListen() {
	c.WsClient.listen()
}
