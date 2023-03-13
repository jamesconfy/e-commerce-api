package emailService

import (
	"e-commerce/internal/models/emailModels"
	"fmt"
	"net/smtp"
)

type EmailService interface {
	SendMail(req emailModels.SendEmailReq) error
	SendBatchEmail(req emailModels.SendBatchEmail) error
	SendMailToSupport(req emailModels.SendEmailReq) error
}
type emailSrv struct {
	FromEmail string
	Password  string
	Host      string
	Port      string
}

func (e emailSrv) SendBatchEmail(req emailModels.SendBatchEmail) error {
	auth := smtp.PlainAuth("", e.FromEmail, e.Password, e.Host)
	addr := e.Host + ":" + e.Port
	header := fmt.Sprintf("From: %v\nTo: %v\n", e.FromEmail, req.EmailAddresses)
	body := []byte(header + req.EmailSubject + "\n" + req.EmailBody)
	err := smtp.SendMail(addr, auth, e.FromEmail, req.EmailAddresses, body)
	if err != nil {
		return err
	}
	return nil
}

func (e emailSrv) SendMail(req emailModels.SendEmailReq) error {
	auth := smtp.PlainAuth("", e.FromEmail, e.Password, e.Host)

	addr := e.Host + ":" + e.Port

	header := fmt.Sprintf("From: %v\nTo: %v\n", e.FromEmail, req.EmailAddress)

	body := []byte(header + req.EmailSubject + req.EmailBody)

	err := smtp.SendMail(addr, auth, e.FromEmail, []string{req.EmailAddress}, body)
	if err != nil {
		return err
	}
	return nil
}

func (e emailSrv) SendMailToSupport(req emailModels.SendEmailReq) error {
	auth := smtp.PlainAuth("", e.FromEmail, e.Password, e.Host)
	addr := e.Host + ":" + e.Port
	header := fmt.Sprintf("From: %v\nTo: %v\n", req.EmailAddress, e.FromEmail)
	body := []byte(header + req.EmailSubject + req.EmailBody)
	err := smtp.SendMail(addr, auth, req.EmailAddress, []string{e.FromEmail}, body)

	if err != nil {
		return err
	}

	return nil
}

func NewEmailSrv(fromEmail string, password string, host string, port string) EmailService {
	return &emailSrv{FromEmail: fromEmail, Password: password, Host: host, Port: port}
}
