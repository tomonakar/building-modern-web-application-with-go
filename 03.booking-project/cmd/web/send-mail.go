package main

import (
	"log"
	"time"

	"github.com/tomonar/booking/internal/models"
	mail "github.com/xhit/go-simple-mail/v2"
)

func listenForMail() {
	go func() {
		for {
			msg := <-app.MailChan
			sendMsg(msg)
		}
	}()
}

func sendMsg(m models.MailData) {
	sever := mail.NewSMTPClient()
	sever.Host = "localhost"
	sever.Port = 1025
	sever.KeepAlive = false
	sever.ConnectTimeout = 10 * time.Second
	sever.SendTimeout = 10 * time.Second

	client, err := sever.Connect()
	if err != nil {
		errorLog.Println(err)
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	email.SetBody(mail.TextHTML, m.Content)

	err = email.Send(client)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email sent!")

	}
}
