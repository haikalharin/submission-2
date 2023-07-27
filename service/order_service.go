package service

import (
	"submission-2/model/domain"
	"submission-2/model/web/request"
)

type OrderService interface {
	Create(request request.OrderCreate) error
	Update(request request.OrderUpdate, id int) error
	Delete(id int) error
	FindById(id int) (domain.Order, error)
	FindAll() ([]domain.Order, error)
}
