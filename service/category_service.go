package service

import (
	"errors"
	"go-unit-test/entity"
	"go-unit-test/repository"
)

type CategoryService struct{
	Repository repository.CategoryRepository
}

func (s CategoryService) Get (id string) (*entity.Category, error) {
	c := s.Repository.FindById(id)

	if c == nil {
		return c, errors.New("Not found")
	} else {
		return c, nil
	}
}