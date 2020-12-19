package sys

import (
	"runtime"

	"github.com/golang/protobuf/proto"
	"github.com/shirou/gopsutil/v3/cpu"
	//"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"

	pb "cpustats/proto/github.com/jroosing/cpustats/grpc/v1/sys"
)

func (s *Stats) GetHardwareInfoProto() []byte {
	bytes, err := proto.Marshal(s.GetHardwareInfo())
	if err != nil {
		s.log.Errorf("could not marshal hardware info to proto: %v", err)
	}
	return bytes;
}

func (s *Stats) GetHardwareInfo() *pb.HardwareInfo {
	return &pb.HardwareInfo{
		Host: s.GetHostInfo(),
		Cpu:  s.GetCpuInfo(),
	}
}

func (s *Stats) GetHostInfo() *pb.Host {
	hostInfo, err := host.Info()
	if err != nil {
		s.log.Errorf("unable to get host info: %s", err.Error())
		return &pb.Host{}
	}

	return &pb.Host{
		HostName:           hostInfo.Hostname,
		Platform:           hostInfo.Platform,
		PlatformFamily:     hostInfo.PlatformFamily,
		PlatformVersion:    hostInfo.PlatformVersion,
		KernelVersion:      hostInfo.KernelVersion,
		KernelArchitecture: hostInfo.KernelArch,
		OperatingSystem:    runtime.GOOS,
	}
}

func (s *Stats) GetCpuInfo() *pb.Cpu {
	cpuInfo, err := cpu.Info()
	if err != nil {
		s.log.Errorf("unable to retrieve cpu information: %v", err)
		return &pb.Cpu{}
	}

	c := &pb.Cpu{
		ModelName:     cpuInfo[0].ModelName,
		PhysicalCores: 0,
		LogicalCores:  0,
		CacheSize:     cpuInfo[0].CacheSize,
		Architecture:  runtime.GOARCH,
		Mhz:           cpuInfo[0].Mhz,
	}

	physicalCores, err := cpu.Counts(false)
	if err != nil {
		s.log.Errorf("unable to retrieve physical core count: %v", err)
	}
	c.PhysicalCores = int32(physicalCores)
	logicalCores, err := cpu.Counts(true)
	if err != nil {
		s.log.Errorf("unable to retrieve logical core count: %v", err)
	}
	c.LogicalCores = int32(logicalCores)

	return c
}
