package sys

import "C"
import (
	"math"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"

	"cpustats/pkg/sys/temperature"

	pb "cpustats/proto/github.com/jroosing/cpustats/grpc/v1/sys"
)

func (s *Stats) GetHardwareUsage() *pb.HardwareLoad {
	return &pb.HardwareLoad{
		Cpu:     s.getCpuLoad(),
		Mem:     s.GetMemoryLoad(),
		SwapMem: s.GetSwapMemoryLoad(),
		Temps:   s.getTemperatures(),
	}
}

func (s *Stats) getTemperatures() *pb.HardwareTemperatures {
	hwTemp := &pb.HardwareTemperatures{}

	c, err := temperature.GetCpuTemp()
	if err != nil {
		s.log.Errorf("unable to get cpu temp: %v", err)
	}
	hwTemp.Cpu = c

	g, err := temperature.GetGpuTemp()
	if err != nil {
		s.log.Errorf("unable to get gpu temperature: %v", err)
	}
	hwTemp.Gpu = g

	d, err := temperature.GetDiskTemp()
	if err != nil {
		s.log.Errorf("unable to get disk temperature: %v", err)
	}
	hwTemp.Disk = d

	m, err := temperature.GetMemoryTemp()
	if err != nil {
		s.log.Errorf("unable to get memory temperature: %v", err)
	}
	hwTemp.Memory = m

	return hwTemp
}

// GetCPUUsage .
func (s *Stats) getCpuLoad() *pb.CpuLoad {
	percent, err := cpu.Percent(1*time.Second, false)
	if err != nil {
		s.log.Errorf("unable to get cpu stats: %s", err.Error())
		return &pb.CpuLoad{}
	}

	cpuLoad := &pb.CpuLoad{
		UsedPercent: int32(math.Round(percent[0])),
	}

	perCpuPercent, err := cpu.Percent(1*time.Second, true)
	if err != nil {
		s.log.Errorf("unable to get per cpu percentage stats: %s", err.Error())
	}

	var perLogicalCorePercent []int32
	for _, v := range perCpuPercent {
		perLogicalCorePercent = append(perLogicalCorePercent, int32(math.Round(v)))
	}
	cpuLoad.UsedPerLogicalCorePercent = perLogicalCorePercent

	return cpuLoad
}

func (s *Stats) GetSwapMemoryLoad() *pb.MemoryLoad {
	swapMem, err := mem.SwapMemory()
	if err != nil {
		s.log.Errorf("unable to retrieve swap memory: %s", err.Error())
		return &pb.MemoryLoad{}
	}
	return &pb.MemoryLoad{
		Total:       swapMem.Total,
		Used:        swapMem.Used,
		UsedPercent: int32(math.Round(swapMem.UsedPercent)),
	}
}

// GetMemory returns the realtime memory usage
func (s *Stats) GetMemoryLoad() *pb.MemoryLoad {
	vMem, err := mem.VirtualMemory()
	if err != nil {
		s.log.Errorf("unable to retrieve swap memory: %s", err.Error())
		return &pb.MemoryLoad{}
	}
	return &pb.MemoryLoad{
		Total:       vMem.Total,
		Used:        vMem.Used,
		UsedPercent: int32(math.Round(vMem.UsedPercent)),
	}
}
