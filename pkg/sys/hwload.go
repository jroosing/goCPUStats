package sys

import "C"
import (
	"math"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"

	"cpustats/pkg/sys/temp"
)

type HardwareLoad struct {
	Cpu     CpuLoad              `json:"cpuLoad"`
	Mem     MemoryLoad           `json:"memLoad"`
	SwapMem MemoryLoad           `json:"swapMemLoad"`
	Temps   HardwareTemperatures `json:"temps"`
}

type CpuLoad struct {
	UsedPercent               int   `json:"usedPercent"`
	UsedPerLogicalCorePercent []int `json:"usedPerLogicalCorePercent"`
}

type MemoryLoad struct {
	Total       uint64 `json:"total"`
	Used        uint64 `json:"used"`
	UsedPercent int    `json:"usedPercent"`
}

type HardwareTemperatures struct {
	Cpu    temp.TempStats
	Gpu    temp.TempStats
	Memory temp.TempStats
	Disk   temp.TempStats
}

func (s *Stats) GetHardwareUsage() *HardwareLoad {
	return &HardwareLoad{
		Cpu:     s.getCpuLoad(),
		Mem:     s.GetMemoryLoad(),
		SwapMem: s.GetSwapMemoryLoad(),
		Temps:   s.getTemperatures(),
	}
}

func (s *Stats) getTemperatures() HardwareTemperatures {
	hwTemp := HardwareTemperatures{}

	c, err := temp.GetCpuTemp()
	if err != nil {
		s.log.Errorf("unable to get cpu temp: %v", err)
	}
	hwTemp.Cpu = c

	g, err := temp.GetGpuTemp()
	if err != nil {
		s.log.Errorf("unable to get gpu temp: %v", err)
	}
	hwTemp.Gpu = g

	d, err := temp.GetDiskTemp()
	if err != nil {
		s.log.Errorf("unable to get disk temp: %v", err)
	}
	hwTemp.Disk = d

	m, err := temp.GetMemoryTemp()
	if err != nil {
		s.log.Errorf("unable to get memory temp: %v", err)
	}
	hwTemp.Memory = m

	return hwTemp
}

// GetCPUUsage .
func (s *Stats) getCpuLoad() CpuLoad {
	percent, err := cpu.Percent(1*time.Second, false)
	if err != nil {
		s.log.Errorf("unable to get cpu stats: %s", err.Error())
		return CpuLoad{}
	}

	cpuLoad := CpuLoad{
		UsedPercent: int(math.Round(percent[0])),
	}

	perCpuPercent, err := cpu.Percent(1*time.Second, true)
	if err != nil {
		s.log.Errorf("unable to get per cpu percentage stats: %s", err.Error())
	}

	var perLogicalCorePercent []int
	for _, v := range perCpuPercent {
		perLogicalCorePercent = append(perLogicalCorePercent, int(math.Round(v)))
	}
	cpuLoad.UsedPerLogicalCorePercent = perLogicalCorePercent

	return cpuLoad
}

func (s *Stats) GetSwapMemoryLoad() MemoryLoad {
	swapMem, err := mem.SwapMemory()
	if err != nil {
		s.log.Errorf("unable to retrieve swap memory: %s", err.Error())
		return MemoryLoad{}
	}
	return MemoryLoad{
		Total:       swapMem.Total,
		Used:        swapMem.Used,
		UsedPercent: int(math.Round(swapMem.UsedPercent)),
	}
}

// GetMemory returns the realtime memory usage
func (s *Stats) GetMemoryLoad() MemoryLoad {
	vMem, err := mem.VirtualMemory()
	if err != nil {
		s.log.Errorf("unable to retrieve swap memory: %s", err.Error())
		return MemoryLoad{}
	}
	return MemoryLoad{
		Total:       vMem.Total,
		Used:        vMem.Used,
		UsedPercent: int(math.Round(vMem.UsedPercent)),
	}
}
