package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	mathservice "github.com/niciki/go-grpc-math-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	listener, err := net.Listen("tcp", ":5300")
	if err != nil {
		grpclog.Fatalf("failed to listen port:  %v", err)
	}
	opts := []grpc.ServerOption{}
	grpcServer := *grpc.NewServer(opts...)
	mathservice.RegisterMathServiceServer(&grpcServer, &server{})
	grpcServer.Serve(listener)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGINT,
		syscall.SIGHUP,
	)
	go func() {
		<-quit
		grpcServer.GracefulStop()
	}()
}

type server struct {
}

func (s *server) MakeOperation(ctx context.Context, req *mathservice.OperationRequest) (
	*mathservice.OperationResponse, error) {
	grpclog.Infoln(req)
	var resp mathservice.OperationResponse
	switch req.GetOperator() {
	case "+":
		resp.Result = req.GetNumber1() + req.GetNumber2()
	case "-":
		resp.Result = req.GetNumber1() - req.GetNumber2()
	case "*":
		resp.Result = req.GetNumber1() * req.GetNumber2()
	case "/":
		resp.Result = req.GetNumber1() / req.GetNumber2()
	default:
		fmt.Print(req.Operator)
		return &mathservice.OperationResponse{}, errors.New("Invalid operator")
	}
	return &resp, nil
}
