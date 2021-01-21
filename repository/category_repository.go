package repository

import "awesomeProject/entity"

type CateGoryRepository interface {
	FindById(id int) *entity.Category
}
