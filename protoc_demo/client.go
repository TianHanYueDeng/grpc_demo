package main
import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"protoc_demo/message"
	"time"
)

func main(){
	con, err := grpc.Dial("127.0.0.1:8080",grpc.WithInsecure())
	if err!=nil{
		panic(err.Error())
	}
	defer con.Close()
	client := message.NewOrderServiceClient(con)
	request := message.OrderRequest{OrderId:"100",TimeStamp:time.Now().Unix()}
	orderInfo,err := client.GetOrderInfo(context.Background(),&request)
	if err!=nil{
		panic(err.Error())
	}
	fmt.Println(orderInfo.GetOrderId())
	fmt.Println(orderInfo.GetOrderName())
	fmt.Println(orderInfo.GetOrderStatus())
}


