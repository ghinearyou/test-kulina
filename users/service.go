package users

import (
	"test-kulina/databases"
)
type Service struct {
	repository    *repository
}

func New() *Service {
	return &Service{
		repository: &repository{
			db: databases.New(),
		},
	}
}

func (s *Service) CreateUser(params User) (error){
	err := s.repository.CreateUser(params)
	if err != nil {
		return err
	}

	return nil
}