package handler

import (
	"context"
	"github.com/wsjcko/shoporder/common"
	"github.com/wsjcko/shoporder/domain/model"
	"github.com/wsjcko/shoporder/domain/service"
	pb "github.com/wsjcko/shoporder/protobuf/pb"
)

type ShopOrder struct {
	OrderService service.IOrderService
}

//根据订单ID查询订单
func (o *ShopOrder) GetOrderByID(ctx context.Context, request *pb.OrderID, response *pb.OrderInfo) error {
	order, err := o.OrderService.FindOrderByID(request.OrderId)
	if err != nil {
		return err
	}
	if err := common.SwapTo(order, response); err != nil {
		return err
	}
	return nil
}

//查找所有订单
func (o *ShopOrder) GetAllOrder(ctx context.Context, request *pb.AllOrderRequest, response *pb.AllOrder) error {
	orderAll, err := o.OrderService.FindAllOrder()
	if err != nil {
		return err
	}

	for _, v := range orderAll {
		order := &pb.OrderInfo{}
		if err := common.SwapTo(v, order); err != nil {
			return err
		}
		response.OrderInfo = append(response.OrderInfo, order)
	}
	return nil
}

//创建订单
func (o *ShopOrder) CreateOrder(ctx context.Context, request *pb.OrderInfo, response *pb.OrderID) error {
	orderAdd := &model.Order{}
	if err := common.SwapTo(request, orderAdd); err != nil {
		return err
	}
	orderID, err := o.OrderService.AddOrder(orderAdd)
	if err != nil {
		return err
	}
	response.OrderId = orderID
	return nil
}

//删除订单
func (o *ShopOrder) DeleteOrderByID(ctx context.Context, request *pb.OrderID, response *pb.Response) error {
	if err := o.OrderService.DeleteOrder(request.OrderId); err != nil {
		return err
	}
	response.Msg = "删除成功"
	return nil
}

//更新订单支付状态
func (o *ShopOrder) UpdateOrderPayStatus(ctx context.Context, request *pb.PayStatus, response *pb.Response) error {
	if err := o.OrderService.UpdatePayStatus(request.OrderId, request.PayStatus); err != nil {
		return err
	}
	response.Msg = "支付状态更新成功"
	return nil
}

//更新发货状态
func (o *ShopOrder) UpdateOrderShipStatus(ctx context.Context, request *pb.ShipStatus, response *pb.Response) error {
	if err := o.OrderService.UpdateShipStatus(request.OrderId, request.ShipStatus); err != nil {
		return err
	}
	response.Msg = "发货状态更新成功"
	return nil
}

//更新订单状态
func (o *ShopOrder) UpdateOrder(ctx context.Context, request *pb.OrderInfo, response *pb.Response) error {
	order := &model.Order{}
	if err := common.SwapTo(request, order); err != nil {
		return err
	}
	if err := o.OrderService.UpdateOrder(order); err != nil {
		return err
	}
	response.Msg = "订单更新成功"
	return nil
}
