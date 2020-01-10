package crm

// Contact is a single contact, i.e. a person in our CRM
type Contact struct {
	ID      int    `db:"id" json:"id"`
	Name    string `db:"name" json:"name"`
	EMail   string `db:"email" json:"email"`
	Company string `db:"company" json:"company"`
}
