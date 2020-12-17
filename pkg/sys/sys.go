package sys

import (
	"math"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/wailsapp/wails"
)

// Stats .
type Stats struct {
	log *wails.CustomLogger
}

type HardwareInfo struct {
	HostInfo        *host.InfoStat `json:"hostInfo"`
	OperatingSystem string         `json:"operatingSystem"`

	CPUInfo         []cpu.InfoStat `json:"cpuInfo"`
	CPUThreadCount  int            `json:"cpuThreadCount"`
	CPUArchitecture string         `json:"architecture"`
}

type HardwareUsage struct {
	CPUAvg          int                    `json:"cpuAvg"`
	CPUThreadUsages []int                  `json:"cpuThreadUsages"`
	Swap            *mem.SwapMemoryStat    `json:"swap"`
	Mem             *mem.VirtualMemoryStat `json:"mem"`
	Temp            []host.TemperatureStat `json:"temp"`
}

// WailsInit .
func (s *Stats) WailsInit(runtime *wails.Runtime) error {
	s.log = runtime.Log.New("Stats")

	go func() {
		for {
			runtime.Events.Emit("hwUsage", s.GetHardwareUsage())
			time.Sleep(20 * time.Second)
		}
	}()

	return nil
}

func (s *Stats) GetHardwareUsage() *HardwareUsage {
	return &HardwareUsage{
		CPUAvg:          s.GetCPUUsage(),
		CPUThreadUsages: s.GetCPUThreadUsages(),
		Swap:            s.GetSwapMemory(),
		Mem:             s.GetMemory(),
		Temp:            s.GetTempStats(),
	}
}

func (s *Stats) GetHardwareInfo() *HardwareInfo {
	return &HardwareInfo{
		HostInfo:        s.GetHostInfo(),
		CPUInfo:         s.GetCPUInfo(),
		OperatingSystem: runtime.GOOS,
		CPUArchitecture: runtime.GOARCH,
		CPUThreadCount:  s.GetCPUCount(),
	}
}

func (s *Stats) GetHostInfo() *host.InfoStat {
	hostInfo, err := host.Info()
	if err != nil {
		s.log.Errorf("unable to get host info: %s", err.Error())
		return &host.InfoStat{}
	}
	return hostInfo
}

// GetCPUUsage .
func (s *Stats) GetCPUUsage() int {
	percent, err := cpu.Percent(1*time.Second, false)
	if err != nil {
		s.log.Errorf("unable to get cpu stats: %s", err.Error())
		return -1
	}

	return int(math.Round(percent[0]))
}

func (s *Stats) GetCPUThreadUsages() (perCPUUsage []int) {
	percentMap, err := cpu.Percent(1*time.Second, true)
	if err != nil {
		s.log.Errorf("unable to get per cpu percentage stats: %s", err.Error())
		return []int{}
	}
	for _, v := range percentMap {
		perCPUUsage = append(perCPUUsage, int(math.Round(v)))
	}
	return
}

func (s *Stats) GetCPUCount() int {
	count, err := cpu.Counts(true)
	if err != nil {
		s.log.Errorf("unable to retrieve cpu count: %s", err.Error())
		return -1
	}
	return count
}

func (s *Stats) GetSwapMemory() *mem.SwapMemoryStat {
	sms, err := mem.SwapMemory()
	if err != nil {
		s.log.Errorf("unable to retrieve swap memory: %s", err.Error())
		return &mem.SwapMemoryStat{}
	}
	return sms
}

// GetMemory returns the realtime memory usage
func (s *Stats) GetMemory() *mem.VirtualMemoryStat {
	ms, err := mem.VirtualMemory()
	if err != nil {
		s.log.Errorf("Unable to retrieve memory!")
		return &mem.VirtualMemoryStat{}
	}
	return ms
}

// GetCPUInfo returns a slice of type InfoStat
func (s *Stats) GetCPUInfo() []cpu.InfoStat {
	cpuInfo, err := cpu.Info()
	if err != nil {
		s.log.Errorf("unable to retrieve cpu information!")
		return []cpu.InfoStat{}
	}
	return cpuInfo
}

func (s *Stats) GetTempStats() []host.TemperatureStat {
	tempStat, err := host.SensorsTemperatures()
	if err != nil {
		s.log.Errorf("unable to retrieve temps")
		return []host.TemperatureStat{}
	}
	return tempStat
}
