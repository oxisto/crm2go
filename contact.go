package crm

import "time"

type Note struct {
	ID        int64     `db:"id" json:"id"`
	ContactID int64     `db:"contactId" json:"contactId"`
	Text      string    `db:"text" json:"text"`
	Date      time.Time `db:"date" json:"date"`
}

// Contact is a single contact, i.e. a person in our CRM
type Contact struct {
	ID int64 `db:"id" json:"id"`
	// Name is a required attribute
	Name string `db:"name" json:"name"`
	// EMail might be optional
	EMail *string `db:"email" json:"email,omitempty"`
	// Company might be optional
	Company *string `db:"company" json:"company,omitempty"`
}
