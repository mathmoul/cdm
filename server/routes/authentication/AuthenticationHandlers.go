package authentication

import (
	"net/http"
	"io"
	"io/ioutil"
	"encoding/json"
	"errors"
	"log"
	"fmt"
	"cdm/server/muxrouter"
	"cdm/server/models"
)

type S_ResetPassword struct {
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
	Token                string `json:"token"`
}

// TODO protect back from false requests with cookie and check every datas

//PRIVATE FUNCTIONS
func init() {

}

func parseJsonBody(name string, r io.Reader) (ret string, err error) {
	var body []byte
	var t map[string]interface{}

	if body, err = ioutil.ReadAll(r); err != nil {
		return
	}
	if err = json.Unmarshal(body, &t); err != nil {
		return
	}
	if _, ok := t[name].(string); !ok {
		err = errors.New("error")
		return
	}
	ret = t[name].(string)
	return
}

//used in resetpassword request
//parse body to find passwords
func getPassword(reader io.Reader) (data S_ResetPassword, err error) {
	var body []byte
	var t = map[string]S_ResetPassword{
		"data": data,
	}
	body, err = ioutil.ReadAll(reader)
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &t); err != nil {
		return
	}
	data = t["data"]
	return
}

//Authentication == login function
//response for front login form
func Authentication(w http.ResponseWriter, r *http.Request) {
	u := models.NewUsers()
	j := muxrouter.JSONResponseWriter{ResponseWriter: w}
	c, err := u.ParseBody("credentials", r.Body)
	if err != nil {
		j.Error401(err)
		return
	}
	fu, err := u.GetOneByEmail(c.Email)
	if err != nil {
		j.Error401(err)
		return
	}
	if err := u.Login(fu); err != nil {
		j.Error401(err)
		return
	}
	log.Println(u)
	j.Success(muxrouter.JSON{"user": u.ToAuthJSON()})

}

//Signup function
//register a new user
//TODO need to send email to user to go to confirmation page

func Signup(w http.ResponseWriter, r *http.Request) {
	u := models.NewUsers()
	j := muxrouter.JSONResponseWriter{ResponseWriter: w}
	_, err := u.ParseBody("user", r.Body)
	if err != nil {
		j.Error401(err)
		return
	}
	if err := u.InsertNewUsers(); err != nil {
		j.Error401(err)
		return
	}
	log.Println(u.GenerateConfirmationURL())
	j.Success(muxrouter.JSON{"user": u.ToAuthJSON()})
}

//Confirmation function
//update user with confirmation / activation email
//send user with token see ToAuthJson()
func Confirmation(w http.ResponseWriter, r *http.Request) {
	u := models.NewUsers()
	j := muxrouter.JSONResponseWriter{ResponseWriter: w}
	token, err := parseJsonBody("token", r.Body)
	if err != nil {
		j.Error401(err)
		return
	}
	u.ConfirmationToken = token
	U, err := u.ConfirmConnection()
	if err != nil {
		fmt.Println(err)
		j.Error401(err)
		return
	}
	j.Success(muxrouter.JSON{"user": U.ToAuthJSON()})
}

//ResetPasswordRequest function
//if user needs to reset his password
//function will mail an new request with fresh token and expiration
func ResetPasswordRequest(w http.ResponseWriter, r *http.Request) {
	u := models.NewUsers()
	j := muxrouter.JSONResponseWriter{ResponseWriter: w}
	// TODO mailer
	// TODO save token in db with user to know if user actually send a reset request
	mail, err := parseJsonBody("email", r.Body)
	if err != nil {
		j.Error401(err)
		return
	}
	U, err := u.GetOneByEmail(mail)
	if err != nil {
		j.Error401(err)
		return
	}
	fmt.Println(U.GeneratePasswordLink())
	// return success
}

//ValiateToken function
//when user need to reset his password he got a mail, this function handle this mail request
//if send success the user can access to reset password form
func ValidateToken(w http.ResponseWriter, r *http.Request) {
	u := models.NewUsers()
	j := muxrouter.JSONResponseWriter{ResponseWriter: w}
	token, err := parseJsonBody("token", r.Body)
	if err != nil {
		j.Error401(err)
		return
	}
	fmt.Println(token)
	// TODO compare token with db token
	u.ConfirmationToken = token
	if err := u.ValidateToken(); err != nil {
		j.Error401(err)
		return
	}
	j.Success(nil)
	// return success
}

//ResetPasswordFunction
//New password for user
//After ask from user to reset his password
// TODO compare old password and new password ?
// TODO compare token from db token
// TODO mailing
func ResetPassword(w http.ResponseWriter, r *http.Request) {
	u := models.NewUsers()
	j := muxrouter.JSONResponseWriter{ResponseWriter: w}
	s, err := getPassword(r.Body)
	if err != nil {
		j.Error401(err)
		return
	}
	u.ConfirmationToken = s.Token
	if err := u.ValidateToken(); err != nil {
		j.Error401(err)
	}
	if err := u.FindWithId(s.Token); err != nil {
		j.Error401(err)
		return
	}
	u.PasswordHash = s.Password
	u.SetPassword()
	if _, err := u.Update(); err != nil {
		j.Error401(err)
		return
	}
	j.Success(nil)
}
