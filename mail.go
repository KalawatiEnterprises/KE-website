// copied from https://github.com/tangingw/go_smtp/blob/master/send_mail.go

package main

import (
	"bytes"
	"fmt"
	"mime/quotedprintable"
	"net/smtp"
	"strings"
)

/**
	Modified from https://gist.github.com/jpillora/cb46d183eca0710d909a
	Thank you very much.
**/

const (
	/**
		Gmail SMTP Server
	**/
	SMTPServer = "smtp.gmail.com"
)

type Sender struct {
	User     string
	Password string
}
var sender Sender = Sender{"kalawatienterprises0@gmail.com", ""}

func (sender Sender) SendMail(Dest []string, Subject, bodyMessage string) {

	msg := "From: " + sender.User + "\n" +
		"To: " + strings.Join(Dest, ",") + "\n" +
		"Subject: " + Subject + "\n" + bodyMessage

	err := smtp.SendMail(SMTPServer+":587",
		smtp.PlainAuth("", sender.User, sender.Password, SMTPServer),
		sender.User, Dest, []byte(msg))

	if err != nil {

		fmt.Printf("smtp error: %s", err)
		return
	}

	fmt.Println("Mail sent successfully!")
}

func (sender Sender) WriteEmail(dest []string, contentType, subject, bodyMessage string) string {

	header := make(map[string]string)
	header["From"] = sender.User

	receipient := ""

	for _, user := range dest {
		receipient = receipient + user
	}

	header["To"] = receipient
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", contentType)
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	message := ""

	for key, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	var encodedMessage bytes.Buffer

	finalMessage := quotedprintable.NewWriter(&encodedMessage)
	finalMessage.Write([]byte(bodyMessage))
	finalMessage.Close()

	message += "\r\n" + encodedMessage.String()

	return message
}

func (sender *Sender) WriteHTMLEmail(dest []string, subject, bodyMessage string) string {

	return sender.WriteEmail(dest, "text/html", subject, bodyMessage)
}

func (sender *Sender) WritePlainEmail(dest []string, subject, bodyMessage string) string {

	return sender.WriteEmail(dest, "text/plain", subject, bodyMessage)
}

// modified code
func confirmMail (body string, r string) {
        //The receiver needs to be in slice as the receive supports multiple receiver
        receiver := []string{r}

        subject := "Thanks for contacting Kalawati Enterprises!"

        greeting := "Thanks for contacting Kalawati Enterprises! You will get a call or an emailregarding this following message you sent:<br>"
        message := `
        <!DOCTYPE HTML PULBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
        <html>
        <head>
        <meta http-equiv="content-type" content="text/html"; charset=ISO-8859-1">
        </head>
        <body>` + greeting + body + `<br>
        <div class="moz-signature">
        <b><br><br>Regards<br>Kalawati Enterprises(Not Alex)</b>
        </div>
        </body>
        </html>
        `
        bodyMessage := sender.WriteHTMLEmail(receiver, subject, message)
        sender.SendMail(receiver, subject, bodyMessage)
}

func notifySubmission (name string, email string, phone string, message string) {
        //The receiver needs to be in slice as the receive supports multiple receiver
        receiver := []string{"zerotsu@ctemplar.com"/*sender.username*/}

        subject := "New form submission received"

        msg := `
        <!DOCTYPE HTML PULBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
        <html>
        <head>
        <meta http-equiv="content-type" content="text/html"; charset=ISO-8859-1">
        </head>
        <body>
	Greetings, <b>` + name + `</b> filled the contact form posted on your website.
	here is the following message: <br><div>` + message + `</div><br><br>
	Their contact details are: <br>
	Phone: ` + phone + `<br> Email: ` + email + `
        </body>
        </html>
        `
        bodyMessage := sender.WriteHTMLEmail(receiver, subject, msg)
        sender.SendMail(receiver, subject, bodyMessage)
}
