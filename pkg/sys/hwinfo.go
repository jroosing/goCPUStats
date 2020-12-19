package sys

import (
	"runtime"

	"github.com/shirou/gopsutil/v3/cpu"
	//"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
)

type HardwareInfo struct {
	Host Host `json:"host"`
	Cpu  Cpu  `json:"cpu"`
}

type Host struct {
	Hostname           string `json:"hostname"`
	Platform           string `json:"platform"`
	PlatformFamily     string `json:"platformFamily"`
	PlatformVersion    string `json:"platformVersion"`
	KernelVersion      string `json:"kernelVersion"`
	KernelArchitecture string `json:"kernelArchitecture"`
	OperatingSystem    string `json:"operatingSystem"`
}

type Cpu struct {
	ModelName     string  `json:"modelName"`
	PhysicalCores int     `json:"physicalCores"`
	LogicalCores  int     `json:"logicalCores"`
	CacheSize     int32   `json:"cacheSize"`
	Architecture  string  `json:"architecture"`
	Mhz           float64 `json:"mhz"`
}

func (s *Stats) GetHardwareInfo() *HardwareInfo {
	return &HardwareInfo{
		Host: s.GetHostInfo(),
		Cpu:  s.GetCPUInfo(),
	}
}

func (s *Stats) GetHostInfo() Host {
	hostInfo, err := host.Info()
	if err != nil {
		s.log.Errorf("unable to get host info: %s", err.Error())
		return Host{}
	}
	return Host{
		Hostname:           hostInfo.Hostname,
		Platform:           hostInfo.Platform,
		PlatformFamily:     hostInfo.PlatformFamily,
		PlatformVersion:    hostInfo.PlatformVersion,
		KernelVersion:      hostInfo.KernelVersion,
		KernelArchitecture: hostInfo.KernelArch,
		OperatingSystem:    runtime.GOOS,
	}
}

// TODO: add support for a multi cpu setup
// not really high prio
func (s *Stats) GetCPUInfo() Cpu {
	cpuInfo, err := cpu.Info()
	if err != nil {
		s.log.Errorf("unable to retrieve cpu information: %v", err)
		return Cpu{}
	}

	c := Cpu{
		ModelName:    cpuInfo[0].ModelName,
		CacheSize:    cpuInfo[0].CacheSize,
		Architecture: runtime.GOARCH,
		Mhz:          cpuInfo[0].Mhz,
	}

	physicalCores, err := cpu.Counts(false)
	if err != nil {
		s.log.Errorf("unable to retrieve physical core count: %v", err)
	}
	c.PhysicalCores = physicalCores
	logicalCores, err := cpu.Counts(true)
	if err != nil {
		s.log.Errorf("unable to retrieve logical core count: %v", err)
	}
	c.LogicalCores = logicalCores

	return c
}
