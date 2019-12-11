package main

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"net"
	"protoc_demo/message"
	"time"
)


type OrderService struct {

}

func (os * OrderService) GetOrderInfo(ctx context.Context, request *message.OrderRequest) (*message.OrderInfo, error){
	OrderMap := map[string]message.OrderInfo{
		"100":message.OrderInfo{OrderId:"1",OrderName:"test1",OrderStatus:"ok"},
		"200":message.OrderInfo{OrderId:"2",OrderName:"test2",OrderStatus:"ok"},
		"300":message.OrderInfo{OrderId:"3",OrderName:"test3",OrderStatus:"ok"},
	}
	current := time.Now().Unix()
	if request.TimeStamp == current {
		return nil, errors.New("system error")
	}else{
		k, ok:= OrderMap[request.OrderId]
		if !ok{
			return nil, errors.New("not exist")
		}else{
			return &k,nil
		}
	}
}

func main(){
	server := grpc.NewServer()
	message.RegisterOrderServiceServer(server,new(OrderService))
	lis ,err := net.Listen("tcp",":8080")
	if err!=nil{
		panic(err.Error())
	}
	_ = server.Serve(lis)
}