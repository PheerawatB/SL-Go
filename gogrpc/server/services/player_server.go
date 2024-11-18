package services

import (
	context "context"
	"fmt"
	"io"
)

type playerServer struct {
}

func NewPlayerService() PlayerServer {
	return playerServer{}
}

func (playerServer) mustEmbedUnimplementedPlayerServer() {}

func (playerServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	result := fmt.Sprintf("Hello %v", req.Name)
	res := HelloResponse{
		Results: result,
	}
	return &res, nil
}

func (playerServer) Avg(stream Player_AvgServer) error {

	sum := 0.0
	count := 0.0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		sum += req.Number
		count++
	}
	res := AvgResponse{
		Results: sum / count,
	}
	return stream.SendAndClose(&res)
}

func (playerServer) Sum(stream Player_SumServer) error {
	sum := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		sum += req.Number
		res := SumResponse{
			Results: sum,
		}
		err = stream.Send(&res)
		if err != nil {
			return err
		}
	}
	return nil
}
