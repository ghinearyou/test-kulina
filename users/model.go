package users

type User struct {
	ID			int		`db:"Id" json:"id"`
	Name		string	`db:"Name" json:"name"`
	Email		string	`db:"Email" json:"email"`
	Phone		string	`db:"Phone" json:"phone"`
	Address		string	`db:"Address" json:"address"`
}

type Users []User