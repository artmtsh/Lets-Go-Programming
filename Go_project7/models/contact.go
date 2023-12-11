package models

import (
	u "contactsBook/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"regexp"
)

type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserId uint   `json:"user_id"`
}

func (contact *Contact) ValidateContact() (map[string]interface{}, int) {

	if contact.Name == "" {
		return u.Message(http.StatusBadRequest, "Name cannot be empty!"), http.StatusBadRequest
	}
	phoneRegex, _ := regexp.Compile("^\\+(?:[0-9]){6,14}[0-9]$")
	if !phoneRegex.MatchString(contact.Phone) {
		return u.Message(http.StatusBadRequest, "Invalid phone number"), http.StatusBadRequest
	}
	if contact.Phone == "" {
		return u.Message(http.StatusBadRequest, "Phone number cannot be empty!"), http.StatusBadRequest
	}

	if contact.UserId <= 0 {
		return u.Message(http.StatusBadRequest, "User not found!"), http.StatusBadRequest
	}

	return u.Message(http.StatusOK, "success"), http.StatusOK
}

func (contact *Contact) CreateContact() map[string]interface{} {

	if response, statusCode := contact.ValidateContact(); statusCode != http.StatusOK {
		return response
	}

	GetDB().Create(contact)

	resp := u.Message(http.StatusOK, "success")
	resp["contact"] = contact
	return resp
}

func GetContact(id uint) *Contact {

	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

func GetContacts(user uint) []*Contact {

	contactsSlice := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&contactsSlice).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contactsSlice
}

func UpdateContactById(user uint, userID uint, newPhoneNumber string) map[string]interface{} {
	var contact Contact
	if err := GetDB().Where("ID = ?", userID).First(&contact).Error; err != nil {
		return u.Message(http.StatusNotFound, "Contact with given ID doesn`t exist")
	}

	contact.Phone = newPhoneNumber
	GetDB().Save(&contact)
	return u.Message(http.StatusOK, "Contact updated successfully")
}

func DeleteContactById(user uint, contactId uint) (map[string]interface{}, int) {
	var contact Contact
	if err := GetDB().Where("ID = ?", contactId).First(&contact).Error; gorm.IsRecordNotFoundError(err) {
		return u.Message(http.StatusNotFound, "Contact with given ID doesn`t exists"), http.StatusNotFound
	} else if err != nil {
		return u.Message(http.StatusInternalServerError, "Error while finding contact"), http.StatusInternalServerError
	}

	if err := GetDB().Delete(&contact).Error; err != nil {
		return u.Message(http.StatusInternalServerError, "Error deleting contact"), http.StatusInternalServerError
	}
	return u.Message(http.StatusOK, "Contact successfully deleted"), http.StatusOK
}
