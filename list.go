package yandexPdd

import (
	"errors"
)

func (s *session) ListGet() ([]MaillistResponce,error) {
	j := s.pddRequest(GET, "/email/ml/list", map[string]string{})
	if j.Success != "ok" {
		return []MaillistResponce{}, errors.New(j.Error)
	}
	return j.Maillists, nil
}

func (s *session) ListAdd(mailList string) error {
	params := map [string]string{
		"maillist": mailList,
	}
	j := s.pddRequest(POST, "/email/ml/add", params)
	if j.Success != "ok" {
		return errors.New(j.Error)
	}
	return nil
}

func (s *session) ListDel(mailList string) error {
	params := map [string]string{
		"maillist": mailList,
	}
	j := s.pddRequest(POST, "/email/ml/del", params)
	if j.Success != "ok" {
		return errors.New(j.Error)
	}
	return nil
}

func (s *session) ListUnsubscribeEmail(mailList, email string) error {
	params := map [string]string{
		"maillist": mailList,
		"subscriber": email,
	}
	j := s.pddRequest(POST, "/email/ml/unsubscribe", params)
	if j.Success != "ok" {
		return errors.New(j.Error)
	}
	return nil
}

func (s *session) ListGetCanSend(mailList, email string) (bool, error) {
	var canSend bool
	params := map [string]string{
		"maillist": mailList,
		"subscriber": email,
	}
	j := s.pddRequest(POST, "/email/ml/get_can_send_on_behalf", params)
	if j.Success != "ok" {
		return false, errors.New(j.Error)
	}

	if j.CanSend == "yes" {
		canSend = true
	} else {
		canSend = false
	}

	return canSend, nil
}

func (s *session) ListSetCanSend(mailList, email string, canSend bool) error {
	cs := "no"
	if canSend {
		cs = "yes"
	}
	params := map [string]string{
		"maillist": mailList,
		"subscriber": email,
		"can_send_on_behalf": cs,
	}
	j := s.pddRequest(POST, "/email/ml/set_can_send_on_behalf", params)
	if j.Success != "ok" {
		return errors.New(j.Error)
	}
	return nil
}

func (s *session) ListSubscribeEmail(mailList, email string, canSend bool) error {
	cs := "no"
	if canSend {
		cs = "yes"
	}
	params := map [string]string{
		"maillist": mailList,
		"subscriber": email,
		"can_send_on_behalf": cs,
	}
	j := s.pddRequest(POST, "/email/ml/subscribe", params)
	if j.Success != "ok" {
		return errors.New(j.Error)
	}
	return nil
}

func (s *session) ListSubscribers(mailList string) ([]string, error){
	params := map [string]string{
		"maillist": mailList,
	}
	j := s.pddRequest(GET, "/email/ml/subscribers", params)

	if j.Success != "ok" {
		return []string{}, errors.New(j.Error)
	}
	return j.Subscribers, nil
}

func (s *session) ListSubscribersByUid(mailListUid int) ([]string, error){
	params := map [string]string{
		"maillist_uid": string(mailListUid),
	}
	j := s.pddRequest(GET, "/email/ml/subscribers", params)

	if j.Success != "ok" {
		return []string{}, errors.New(j.Error)
	}
	return j.Subscribers, nil
}
