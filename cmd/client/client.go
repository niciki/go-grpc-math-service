package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	mathservice "github.com/niciki/go-grpc-math-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	args := os.Args[1:]
	if len(args) < 3 {
		grpclog.Fatalf("Please write correct number of aruments: <value1> <operator> <value2>")
	}
	conn, err := grpc.Dial("127.0.0.1:5300", opts...)
	if err != nil {
		grpclog.Fatalf("failed to listen port:  %v", err)
	}
	defer conn.Close()
	client := mathservice.NewMathServiceClient(conn)
	value1, err1 := strconv.Atoi(args[0])
	if err1 != nil {
		grpclog.Fatalf("Please write correct value1:  %v", err)
	}
	value2, err2 := strconv.Atoi(args[2])
	if err2 != nil {
		grpclog.Fatalf("Please write correct value1:  %v", err)
	}
	request := &mathservice.OperationRequest{
		Number1:  int64(value1),
		Operator: args[1],
		Number2:  int64(value2),
	}
	response, err := client.MakeOperation(context.TODO(), request)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	fmt.Println(response.GetResult())
}
