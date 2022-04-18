package service

import (
	"github.com/wsjcko/shoporder/domain/model"
	"github.com/wsjcko/shoporder/domain/repository"
)

type IOrderService interface {
	AddOrder(*model.Order) (int64, error)
	DeleteOrder(int64) error
	UpdateOrder(*model.Order) error
	FindOrderByID(int64) (*model.Order, error)
	FindAllOrder() ([]model.Order, error)
	UpdateShipStatus(int64, int32) error
	UpdatePayStatus(int64, int32) error
}

//创建
func NewOrderService(orderRepository repository.IOrderRepository) IOrderService {
	return &OrderService{orderRepository}
}

type OrderService struct {
	OrderRepository repository.IOrderRepository
}

//插入
func (u *OrderService) AddOrder(order *model.Order) (int64, error) {
	return u.OrderRepository.CreateOrder(order)
}

//删除
func (u *OrderService) DeleteOrder(orderID int64) error {
	return u.OrderRepository.DeleteOrderByID(orderID)
}

//更新
func (u *OrderService) UpdateOrder(order *model.Order) error {
	return u.OrderRepository.UpdateOrder(order)
}

//查找
func (u *OrderService) FindOrderByID(orderID int64) (*model.Order, error) {
	return u.OrderRepository.FindOrderByID(orderID)
}

//查找
func (u *OrderService) FindAllOrder() ([]model.Order, error) {
	return u.OrderRepository.FindAll()
}

func (u *OrderService) UpdateShipStatus(orderID int64, shipStatus int32) error {
	return u.OrderRepository.UpdateShipStatus(orderID, shipStatus)
}

func (u *OrderService) UpdatePayStatus(orderID int64, payStatus int32) error {
	return u.OrderRepository.UpdatePayStatus(orderID, payStatus)
}
