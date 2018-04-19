package mailer

import (
	"cdm/server/models"
	"log"
)

type Mailer struct {
}

type IMailer interface {
	SendConfirmationEmail(user models.User) error
	SendResetPasswordEmail(user models.User) error
}

func init() {
	//initialisation connection w/ stmp protocol
}

func (m *Mailer) SendConfirmationEmail(u models.User) error {
	log.Println("send confirmation Email to " + u.Email)
	return nil
}

func (m *Mailer) SendResetPasswordEmail(u models.User) error {
	log.Println("Send Reset PasswordEmail to " + u.Email)
	return nil
}
