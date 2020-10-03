package main

import (
	"andres/calculator/calculatorpb"
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (*server) Add(ctx context.Context, req *calculatorpb.AddRequest) (*calculatorpb.AddResponse, error) {
	num1 := req.GetNumbers().GetNumber_1()
	num2 := req.GetNumbers().GetNumber_2()
	result := num1 + num2
	fmt.Println(strconv.Itoa(int(num1)) + "+" + strconv.Itoa(int(num2)) + " = " + strconv.Itoa(int(result)))
	res := &calculatorpb.AddResponse{
		Result: result,
	}

	return res, nil

}

func main() {
	fmt.Println("Calc server starts...")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}
}
