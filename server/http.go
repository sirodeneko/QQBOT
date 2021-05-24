package server

import (
	"github.com/asmcos/requests"
	"github.com/sirodeneko/QQBOT/coolq"
)

type HttpClient struct {
	bot *coolq.QQBoT

	httpUrl string
	token   string
}

func newHttpClient(urlStr string, token string, bot *coolq.QQBoT) *HttpClient {

	httpClient := &HttpClient{
		bot:     bot,
		httpUrl: urlStr,
		token:   token,
	}

	return httpClient
}

func (hc *HttpClient) newRequest() *requests.Request {
	req := requests.Requests()
	req.Header.Set("User-Agent", "QQBOT")
	req.Header.Set("Authorization", "Token "+hc.token)
	req.Header.Set("Content-Type", "application/json")
	return req
}
