package services

import (
	"context"
	"time"

	"github.com/GustavoSaraivap/gRPCWithGo/pb/vehicle"
)

// type VehicleServiceServer interface {
// 	AddVehicle(context.Context, *Vehicle) (*Vehicle, error)
// 	mustEmbedUnimplementedVehicleServiceServer()
// }

type VehicleService struct {
	vehicle.UnimplementedVehicleServiceServer
}

func NewVehicleService() *VehicleService {
	return &VehicleService{}
}

func (*VehicleService) AddVehicle(ctx context.Context, req *vehicle.Vehicle) (*vehicle.Vehicle, error) {
	return &vehicle.Vehicle{
		Id:    "123",
		Brand: req.GetBrand(),
		Vin:   req.GetVin(),
	}, nil
}

func (*VehicleService) AddVehicleVerbose(req *vehicle.Vehicle, stream vehicle.VehicleService_AddVehicleVerboseServer) error {
	// stream.Send(&pb.UserResultStream{
	// 	Status: "Init",
	// 	User:   &pb.User{},
	// })

	stream.Send(&vehicle.VehicleResultStream{
		Status:  "Init",
		Vehicle: &vehicle.Vehicle{},
	})

	time.Sleep(time.Second * 3)

	// stream.Send(&pb.UserResultStream{
	// 	Status: "Inserting",
	// 	User:   &pb.User{},
	// })

	stream.Send(&vehicle.VehicleResultStream{
		Status:  "Inserting",
		Vehicle: &vehicle.Vehicle{},
	})

	time.Sleep(time.Second * 3)

	// stream.Send(&pb.UserResultStream{
	// 	Status: "User has been inserted",
	// 	User: &pb.User{
	// 		Id:    "123",
	// 		Name:  req.GetName(),
	// 		Email: req.GetEmail(),
	// 	},
	// })

	stream.Send(&vehicle.VehicleResultStream{
		Status: "Vehicle has been inserted",
		Vehicle: &vehicle.Vehicle{
			Id:    "123",
			Brand: req.GetBrand(),
			Vin:   req.GetVin(),
		},
	})

	time.Sleep(time.Second * 3)

	// stream.Send(&pb.UserResultStream{
	// 	Status: "Completed",
	// 	User: &pb.User{
	// 		Id:    "123",
	// 		Name:  req.GetName(),
	// 		Email: req.GetEmail(),
	// 	},
	// })

	stream.Send(&vehicle.VehicleResultStream{
		Status: "Completed",
		Vehicle: &vehicle.Vehicle{
			Id:    "123",
			Brand: req.GetBrand(),
			Vin:   req.GetVin(),
		},
	})

	time.Sleep(time.Second * 3)

	return nil
}
