package yandexPdd

import (
	"log"
	"io/ioutil"
	"encoding/json"
	"net/http"
)

const (
	pddApiUrl  = "https://pddimp.yandex.ru/api2/admin"
	GET = "GET"
	POST = "POST"
)

type (
	session struct {
		token string
		domain string
	}

	MaillistResponse struct{
		Maillist string `json:"maillist"`
		Uid int `json:"uid"`
		Cnt int `json:"cnt"`
	}

	PddImpResponse struct {
		Domain string `json:"domain"`
		Uid int `json:"uid,omitempty"`
		Maillist string `json:"maillist,omitempty"`
		MaillistUid int `json:"maillist_uid,omitempty"`
		Maillists []MaillistResponse `json:"maillists,omitempty"`
		Subscriber string `json:"subscriber,omitempty"`
		SubscriberUid int `json:"subscriber_uid,omitempty"`
		Subscribers []string `json:"subscribers,omitempty"`
		CanSend interface{} `json:"can_send_on_behalf,omitempty"`
		Error string `json:"error"`
		Success string `json:"success"`
	}
)

func New(domain, token string) session {
	var s session
	s.token = token
	s.domain = domain
	return s
}

func (s *session) pddRequest(metod, url string, params map[string]string) PddImpResponse {
	client := &http.Client{}
	req, err := http.NewRequest(metod, pddApiUrl + url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("PddToken", s.token)
	req.Header.Add("Accept", "application/json")
	q := req.URL.Query()
	q.Add("domain", s.domain)
	for i, p := range params {
		q.Add(i, p)
	}
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var r PddImpResponse
	err = json.Unmarshal(data, &r)
	if err != nil {
		log.Fatal(err)
	}

	return r
}

