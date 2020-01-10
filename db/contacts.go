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

// GetAllNotes returns all stored notes from one contact
func GetAllNotes(contactID int64) (contacts []crm.Note, err error) {
	_, err = mapper.Select(&contacts, "select * from note where \"contactId\"=$1", contactID)

	return
}
