package mailer

import (
	"cdm/server/models"
	"log"
)

type Mailer struct {
}

type IMailer interface {
	SendConfirmationEmail(user models.Users) error
	SendResetPasswordEmail(user models.Users) error
}

func init() {
	//initialisation connection w/ stmp protocol
}

func (m *Mailer) SendConfirmationEmail(u models.Users) error {
	log.Println("send confirmation Email to " + u.Email)
	return nil
}

func (m *Mailer) SendResetPasswordEmail(u models.Users) error {
	log.Println("Send Reset PasswordEmail to " + u.Email)
	return nil
}
