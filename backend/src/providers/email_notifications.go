package providers

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"mime/quotedprintable"
	"net/smtp"
	"strings"

	"github.com/high-value-team/groupbox/backend/src/models"
)

type MessageTemplateData struct {
	Title        string
	PersonalLink string
}

const CreateBoxMessageTemplate = `
<!DOCTYPE HTML PULBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
<head>
<meta http-equiv="content-type" content="text/html"; charset="ISO-8859-1">
</head>
Willkommen zur Groupbox: {{ .Title }}<br>
<br>
Ihr persönlicher Link <a href="{{ .PersonalLink }}">{{ .PersonalLink }}</a><br>
<br>
Viel Spass!<br>
</body>
</html>
`

const AddItemMessageTemplate = `
<!DOCTYPE HTML PULBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
<head>
<meta http-equiv="content-type" content="text/html"; charset="ISO-8859-1">
</head>
Eine neue Nachricht wurde hinzugefügt.<br>
<br>
Groupbox: {{ .Title }}<br>
<br>
Ihr persönlicher Link <a href="{{ .PersonalLink }}">{{ .PersonalLink }}</a><br>
<br>
Viel Spass!<br>
</body>
</html>
`

type EmailNotifications struct {
	Domain            string
	NoReplyEmail      string
	SMTPServerAddress string
	Username          string
	Password          string
}

func (e *EmailNotifications) SendInvitations(title string, members []models.Member) {
	for i := range members {
		e.sendInvitation(title, &members[i])
	}
}

func (e *EmailNotifications) NotifyAudience(members []models.Member, title string) {
	for i := range members {
		e.notifyAudience(title, &members[i])
	}
}

func (e *EmailNotifications) notifyAudience(title string, member *models.Member) {
	message := e.buildAddItemMessage(title, member.Key)
	e.sendMail(member.Email, "Neue Nachricht", message)
}

func (e *EmailNotifications) sendInvitation(title string, member *models.Member) {
	messsage := e.buildCreateBoxMessage(title, member.Key)
	e.sendMail(member.Email, "Neue Groupbox", messsage)
}

func (e *EmailNotifications) buildCreateBoxMessage(title, key string) string {
	messageTemplateData := MessageTemplateData{
		Title:        title,
		PersonalLink: fmt.Sprintf("%s/%s", e.Domain, key),
	}
	return buildMessage(CreateBoxMessageTemplate, messageTemplateData)
}

func (e *EmailNotifications) buildAddItemMessage(title, key string) string {
	messageTemplateData := MessageTemplateData{
		Title:        title,
		PersonalLink: fmt.Sprintf("%s/%s", e.Domain, key),
	}
	return buildMessage(AddItemMessageTemplate, messageTemplateData)
}

func buildMessage(templateText string, messageTemplateData MessageTemplateData) string {
	messageTemplate, err := template.New("body").Parse(templateText)
	if err != nil {
		panic(models.SuprisingException{Err: err})
	}
	messageBuffer := bytes.Buffer{}
	err = messageTemplate.Execute(&messageBuffer, messageTemplateData)
	if err != nil {
		log.Print(err)
		panic(models.SuprisingException{Err: err})
	}
	return messageBuffer.String()
}

func (e *EmailNotifications) sendMail(receipient string, subject, message string) {
	// header
	header := make(map[string]string)
	header["From"] = e.NoReplyEmail
	header["To"] = receipient
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", "text/html")
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	// escape message
	var encodedMessageBuffer bytes.Buffer
	finalMessage := quotedprintable.NewWriter(&encodedMessageBuffer)
	_, err := finalMessage.Write([]byte(message))
	if err != nil {
		log.Print(err)
		panic(models.SuprisingException{Err: err})
	}
	err = finalMessage.Close()
	if err != nil {
		log.Print(err)
		panic(models.SuprisingException{Err: err})
	}

	// build email
	email := ""
	for key, value := range header {
		email += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	email += "\r\n" + encodedMessageBuffer.String()

	// send email
	SMTPServer := strings.Split(e.SMTPServerAddress, ":")[0]
	err = smtp.SendMail(
		e.SMTPServerAddress,
		smtp.PlainAuth("", e.Username, e.Password, SMTPServer),
		e.Username, []string{receipient}, []byte(email))
	if err != nil {
		log.Print(err)
		panic(models.SuprisingException{Err: err})
	}
}
