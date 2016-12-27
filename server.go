package main

import (
	"github.com/ianr0bkny/go-sonos/ssdp"
	"golang.org/x/net/context"

	pb "github.com/brotherlogic/sonosrpc/proto"
)

type deviceManager interface {
	Close() error
	Discover(intf string, port string, subscr bool) error
	Devices() ssdp.DeviceMap
}

// ListDevices calls out to list all the wink devices
func (s *Server) ListDevices(ctx context.Context, in *pb.Empty) (*pb.DeviceList, error) {
	if err := s.mgr.Discover("en1", "11209", false); err != nil {
		return nil, err
	}

	devices := &pb.DeviceList{}
	for _, dev := range s.mgr.Devices() {
		devices.Device = append(devices.Device, &pb.Device{Name: dev.Name()})
	}
	return devices, nil
}
