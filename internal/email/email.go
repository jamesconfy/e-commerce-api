package email

type Payload struct {
	From       string
	FromName   string
	Recipients map[string]string
	Subject    string
	Body       string
}
