package controllers

import (
	"contactsBook/models"
	u "contactsBook/utils"
	"encoding/json"
	"net/http"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(http.StatusBadRequest, "Error!"))
		return
	}

	contact.UserId = user
	resp := contact.CreateContact()
	u.Respond(w, resp)
}

var GetContacts = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetContacts(id)
	resp := u.Message(http.StatusOK, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var PutContact = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(http.StatusBadRequest, "Invalid request body"))
		return
	}

	resp := models.UpdateContactById(user, contact.ID, contact.Phone)
	u.Respond(w, resp)
}

var DeleteContact = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(http.StatusBadRequest, "Invalid request body"))
		return
	}

	resp, stat := models.DeleteContactById(user, contact.ID)
	if stat != http.StatusOK {
		u.Respond(w, resp)
		return
	}
	u.Respond(w, resp)
}
