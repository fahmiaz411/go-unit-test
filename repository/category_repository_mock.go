package repository

import (
	"go-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}                                                     

func (repo *CategoryRepositoryMock) FindById (id string) *entity.Category {
	args := repo.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	} else {
		category := args.Get(0).(entity.Category)
		return &category
	}
}