package db

import (
	"crm"
	"database/sql"
)

// GetAllContacts returns all stored contacts
func GetAllContacts() (contacts []crm.Contact, err error) {
	_, err = mapper.Select(&contacts, "select * from contact")

	return
}

// GetContact returns the contact using the specified id
func GetContact(id int64) (*crm.Contact, error) {
	contact := crm.Contact{}

	err := mapper.SelectOne(&contact, "select * from contact where id=$1", id)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &contact, err
}
