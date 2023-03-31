package service

import (
	"e-commerce/internal/email"
	"fmt"
	"net/smtp"
)

type EmailService interface {
	SendMail(req email.Payload) error
	SendBatchEmail(req email.Payload) error
	// SendMailToSupport(req emailModels.SendEmailReq) error
}
type emailSrv struct {
	AuthEmail    string
	AuthPassword string
	AuthHost     string
	AuthPort     string
}

func (e emailSrv) GetAuth() smtp.Auth {
	return smtp.PlainAuth("", e.AuthEmail, e.AuthPassword, e.AuthHost)
}

func (e emailSrv) GetAddress() (str string) {
	str = fmt.Sprintf("%v:%v", e.AuthHost, e.AuthPort)
	return
}

func (e emailSrv) SendBatchEmail(req email.Payload) error {
	auth := e.GetAuth()
	addr := e.GetAddress()

	for _, email := range req.Recipients {
		header := fmt.Sprintf("From: %v\nTo: %v\n", e.AuthEmail, email)
		body := []byte(header + req.Subject + "\n" + req.Body)
		err := smtp.SendMail(addr, auth, e.AuthEmail, []string{email}, body)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e emailSrv) SendMail(req email.Payload) error {
	auth := e.GetAuth()
	addr := e.GetAddress()

	for _, email := range req.Recipients {
		header := fmt.Sprintf("From: %v\nTo: %v\n", e.AuthEmail, email)
		body := []byte(header + req.Subject + req.Body)

		err := smtp.SendMail(addr, auth, e.AuthEmail, []string{email}, body)
		if err != nil {
			return err
		}
	}

	return nil
}

// func (e emailSrv) SendMailToSupport(req emailModels.SendEmailReq) error {
// 	auth := smtp.PlainAuth("", e.FromEmail, e.Password, e.Host)
// 	addr := e.Host + ":" + e.Port
// 	header := fmt.Sprintf("From: %v\nTo: %v\n", req.EmailAddress, e.FromEmail)
// 	body := []byte(header + req.EmailSubject + req.EmailBody)
// 	err := smtp.SendMail(addr, auth, req.EmailAddress, []string{e.FromEmail}, body)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func NewEmailService(authEmail string, authPassword string, authHost string, authPort string) EmailService {
	return &emailSrv{AuthEmail: authEmail, AuthPassword: authPassword, AuthHost: authHost, AuthPort: authPort}
}
