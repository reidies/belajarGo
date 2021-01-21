package service

import (
	"awesomeProject/entity"
	"awesomeProject/repository"
	"errors"
)

type CategoryService struct{
	Repository repository.CateGoryRepository
}

func (service CategoryService) Get(id int) (*entity.Category, error){
	category := service.Repository.FindById(id)
	if category== nil{
		return  category, errors.New("Category Not Found")
	} else {
		return category,nil
	}
}
