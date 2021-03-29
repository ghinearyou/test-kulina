package users

import (
	"log"

	"test-kulina/databases"
)


type repository struct {
	db *databases.Mysql
}

func (r repository) CreateUser(user User) error {
	query := `INSERT into user (
		Id, 
		Name,
		Email,
		Phone,
		Address 
	) 
	VALUES 
	(
		:Id,
		:Name,
		:Email,
		:Phone,
		:Address
	)`

	_, err := r.db.GetActiveDB().NamedExec(query, user)
	if err != nil {
		log.Println("Create user prepare", err)
		return err
	}

	return nil
}