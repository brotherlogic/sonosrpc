package main

import (
	"errors"
	"testing"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/sonosrpc/proto"
	"github.com/ianr0bkny/go-sonos/ssdp"
)

type testDeviceManager struct {
	failOnDiscover bool
}

func (mgr testDeviceManager) Close() error {
	return nil
}

func (mgr testDeviceManager) Discover(intf string, port string, subscr bool) error {
	if mgr.failOnDiscover {
		return errors.New("Failed as promised")
	}
	return nil
}

type testDevice struct{}

func (t testDevice) Product() string {
	return ""
}
func (t testDevice) ProductVersion() string {
	return ""
}
func (t testDevice) Name() string {
	return "winner"
}
func (t testDevice) Location() ssdp.Location {
	return ""
}
func (t testDevice) UUID() ssdp.UUID {
	return ""
}
func (t testDevice) Service(key ssdp.ServiceKey) (service ssdp.Service, has bool) {
	return nil, false
}
func (t testDevice) Services() []ssdp.ServiceKey {
	return []ssdp.ServiceKey{}
}

func (mgr testDeviceManager) Devices() ssdp.DeviceMap {
	mapper := make(ssdp.DeviceMap)
	mapper["madeup"] = testDevice{}
	return mapper
}

// Gets a test server that'll pull from local files rather than reading out
func getTestServer(failOnDiscover bool) Server {
	s := Server{}
	s.mgr = testDeviceManager{failOnDiscover: failOnDiscover}
	return s
}

func TestListDevicesFailOnDiscover(t *testing.T) {
	s := getTestServer(true)
	_, err := s.ListDevices(context.Background(), &pb.Empty{})

	if err == nil {
		t.Error("Discover did not fail, it should have done")
	}
}

func TestListDevices(t *testing.T) {
	s := getTestServer(false)
	list, err := s.ListDevices(context.Background(), &pb.Empty{})

	if err != nil {
		t.Errorf("Failure to list devices: %v", err)
	}

	//Test results should have one device called winner
	if len(list.Device) != 1 || list.Device[0].Name != "winner" {
		t.Errorf("Error in listing devices: %v", list)
	}
}
