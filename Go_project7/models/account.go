package models

import (
	u "contactsBook/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

func (account *Account) Validate() (map[string]interface{}, int) {
	emailRegex, _ := regexp.Compile("^[a-zA-Z0-9_+&*-]+(?:\\.[a-zA-Z0-9_+&*-]+)*@(?:[a-zA-Z0-9-]+\\.)+[a-zA-Z]{2,7}$")
	if !emailRegex.MatchString(account.Email) {
		return u.Message(http.StatusBadRequest, "Invalid email"), http.StatusBadRequest
	}
	if !strings.Contains(account.Email, "@") {
		return u.Message(http.StatusBadRequest, "Email address is not valid!"), http.StatusBadRequest
	}

	if len(account.Password) < 4 {
		return u.Message(http.StatusBadRequest, "Password must be longer then 4 symbols"), http.StatusBadRequest
	}

	acc := &Account{}

	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(acc).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(http.StatusBadRequest, "Connection error!"), http.StatusBadRequest
	}
	if acc.Email != "" {
		return u.Message(http.StatusBadRequest, "The email is already occupied by another user!"), http.StatusBadRequest
	}

	return u.Message(http.StatusOK, "Check is passed!"), http.StatusOK
}

func (account *Account) CreateAccount() map[string]interface{} {

	if resp, statusCode := account.Validate(); statusCode != http.StatusOK {
		return resp
	}

	pwd, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(pwd)

	GetDB().Create(account)

	if account.ID <= 0 {
		return u.Message(http.StatusBadRequest, "Failed to create account, connection error.")
	}

	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenStr, _ := token.SignedString([]byte(os.Getenv("token_pass")))
	account.Token = tokenStr

	account.Password = ""

	response := u.Message(http.StatusOK, "Account has been created!")
	response["account"] = account
	return response
}

func LoginAccount(email, password string) map[string]interface{} {

	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(http.StatusBadRequest, "Email address not found")
		}
		return u.Message(http.StatusBadRequest, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(http.StatusBadRequest, "Password does not match!")
	}
	account.Password = ""
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenStr, _ := token.SignedString([]byte(os.Getenv("token_pass")))
	account.Token = tokenStr

	resp := u.Message(http.StatusOK, "Logged In")
	resp["account"] = account
	return resp
}

func UpdateAccount(id uint, email, password string) map[string]interface{} {

	account := &Account{}
	err := GetDB().Table("accounts").Where("id = ?", id).First(account).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(http.StatusBadRequest, "Account id not found")
		}
		return u.Message(http.StatusInternalServerError, "Connection error. Please retry")
	}
	account.Email = email
	account.Password = password

	if resp, statusCode := account.Validate(); statusCode != http.StatusOK {
		return resp
	}

	pwd, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(pwd)

	GetDB().Save(account)

	account.Password = ""
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenStr, _ := token.SignedString([]byte(os.Getenv("token_pass")))
	account.Token = tokenStr

	resp := u.Message(http.StatusOK, "Account updated")
	resp["account"] = account
	return resp
}

func DeleteAccount(id uint) map[string]interface{} {
	account := &Account{}
	err := GetDB().Table("accounts").Where("id = ?", id).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(http.StatusBadRequest, "Account id not found")
		}
		return u.Message(http.StatusInternalServerError, "Connection error. Please retry")
	}

	GetDB().Delete(account).Where("id = ", id)

	resp := u.Message(http.StatusOK, "Account deleted")
	return resp
}

func GetUser(u uint) *Account {

	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" { //User not found!
		return nil
	}

	acc.Password = ""
	return acc
}
