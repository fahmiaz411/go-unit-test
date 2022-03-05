package service

import (
	"go-unit-test/entity"
	"go-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{
	Mock: mock.Mock{},
}
var categoryService = CategoryService{
	Repository: categoryRepository,
}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)
	c, err := categoryService.Get("1")

	assert.Nil(t, c, "Category Not Nil")
	assert.NotNil(t, err, "Error not expected")
}

func TestCategoryService_GetFound(t *testing.T) {
	cat := entity.Category{
		Id: "2",
		Name: "Fahmi",
	}
	categoryRepository.Mock.On("FindById", "2").Return(cat)
	res, err := categoryService.Get(cat.Id)
	assert.Nil(t, err, "Error expected")
	assert.NotNil(t, res, "Category Nil")
	assert.Equal(t, cat.Id, res.Id)
	assert.Equal(t, cat.Name, res.Name)
}