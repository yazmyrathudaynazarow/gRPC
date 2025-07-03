package main

import (
	"context"
	"fmt"
	"github.com/tmbiz.info/yazmyrathudaynazarow/gRPC/api"
	"google.golang.org/grpc"
	"math/rand"
	"net"
	"time"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	api.RegisterWeatherServiceServer(srv, &myWeatherService{})
	fmt.Println("Starting server...")
	panic(srv.Serve(lis))
}

type myWeatherService struct {
	api.UnimplementedWeatherServiceServer
}

func (m *myWeatherService) ListCities(ctx context.Context,
	req *api.ListCitiesRequest) (*api.ListCitiesResponse, error) {

	return &api.ListCitiesResponse{
		Items: []*api.CityEntry{
			&api.CityEntry{CityCode: "tr_ank", CityName: "Ankara"},
			&api.CityEntry{CityCode: "tr_ist", CityName: "Istanbul"},
		},
	}, nil
}

func (m *myWeatherService) QueryWeather(req *api.WeatherRequest,
	resp api.WeatherService_QueryWeatherServer) error {
	for {
		err := resp.Send(&api.WeatherResponse{
			Temperature: rand.Float32()*10 + 10})
		if err != nil {
			break
		}
		time.Sleep(time.Second)
	}
	return nil
}
