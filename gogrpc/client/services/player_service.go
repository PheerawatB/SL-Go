package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"
)

type PlayerService interface {
	Hello(name string) error
	Avg(numbers ...float64) error
	Sum(numbers ...int32) error
	CallHello(name string) (string, error) //Service func out proto
}

type playerService struct {
	playerClient PlayerClient
}

func NewPlayerService(playerClient PlayerClient) PlayerService {
	return playerService{playerClient}
}

func (base playerService) Hello(name string) error {
	req := HelloRequest{
		Name: name,
	}

	res, err := base.playerClient.Hello(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Service : Hello \n")
	fmt.Printf("Request : %v\n", req.Name)
	fmt.Printf("Response : %v\n", res.Results)
	return nil
}

func (base playerService) Avg(numbers ...float64) error {

	stream, err := base.playerClient.Avg(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("Service : Avg\n")
	for _, number := range numbers {
		req := AvgRequest{
			Number: number,
		}
		fmt.Printf("Request : %v\n", req.Number)
		stream.Send(&req)
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	fmt.Printf("Response : %v\n", res.Results)
	return nil
}

func (base playerService) Sum(numbers ...int32) error { //binary directional
	stream, err := base.playerClient.Sum(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("Service : Sum \n")
	go func() {
		for _, number := range numbers {
			req := SumRequest{
				Number: number,
			}
			fmt.Printf("Request : %v\n", req.Number)
			stream.Send(&req)
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()

	done := make(chan bool)
	errs := make(chan error)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				errs <- err
			}
			fmt.Printf("Response : %v\n", res.Results)
		}
		done <- true
	}()

	select {
	case <-done:
		return nil
	case err := <-errs:
		return err
	}
}

func (base playerService) CallHello(name string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := HelloRequest{Name: name}
	res, err := base.playerClient.Hello(ctx, &req)
	if err != nil {
		return "", fmt.Errorf("gRPC Hello call failed: %v", err)
	}

	return res.Results, nil
}
