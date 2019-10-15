package pet

import (
	"errors"
	"fmt"
)

type Service struct {
	db Persistence
}

func NewService(db Persistence) *Service{
	return &Service{db:db}
}

func (s *Service) AddPet(pet *Pet) (*Pet, error){
	p, err := s.db.Insert(pet)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *Service) UpdatePet(pet *Pet) error{
	err := s.db.Update(pet)
	return err
}

func (s *Service) UpdatePetStatus(id int64, status Status) error {
	p := s.db.FindById(id)
	if p == nil {
		return errors.New(fmt.Sprintf("pet id %d not found", id))
	}
	p.Status = status
	return s.UpdatePet(p)

}

func (s *Service) FindByStatus(status Status) []*Pet {
	return s.db.FindByStatus(status)
}

func (s *Service) FindById(id int64) *Pet {
	return s.db.FindById(id)
}

func (s *Service) DeletePet(id int64) error {
	return s.db.Delete(id)
}

func (s *Service) UploadImage(){
	// TODO: write this
}