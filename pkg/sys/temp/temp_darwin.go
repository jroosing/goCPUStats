// +build darwin

package temp

import (
	"fmt"
	"math"

	pb "cpustats/proto/github.com/jroosing/cpustats/grpc/v1/sys"
)

const (
	cpuTempSensorKey = "TC0P"
	gpuTempSensorKey = "TG0P"
	memoryTempSensorKey = "TM0P"
	diskTempSensorKey = "TH0P"
)

func GetCpuTemp() (*pb.TempStats, error) {
	if temps == nil {
		if err := loadTemperatures(); err != nil {
			return &pb.TempStats{}, fmt.Errorf("could not load temperatures: %w", err)
		}
	}

	for _, t := range temps {
		if t.SensorKey != cpuTempSensorKey {
			continue
		}
		return &pb.TempStats{
			Temperature: int32(math.Round(t.Temperature)),
			High: int32(math.Round(t.High)),
			Critical: int32(math.Round(t.Critical)),
		}, nil
	}
	return &pb.TempStats{}, fmt.Errorf("could not find cpu temps")
}

func GetGpuTemp() (*pb.TempStats, error) {
	if temps == nil {
		if err := loadTemperatures(); err != nil {
			return &pb.TempStats{}, fmt.Errorf("could not load temperatures: %w", err)
		}
	}

	for _, t := range temps {
		if t.SensorKey != gpuTempSensorKey {
			continue
		}
		return &pb.TempStats{
			Temperature: int32(math.Round(t.Temperature)),
			High: int32(math.Round(t.High)),
			Critical: int32(math.Round(t.Critical)),
		}, nil
	}
	return &pb.TempStats{}, fmt.Errorf("could not find cpu temps")
}

func GetDiskTemp() (*pb.TempStats, error) {
	if temps == nil {
		if err := loadTemperatures(); err != nil {
			return &pb.TempStats{}, fmt.Errorf("could not load temperatures: %w", err)
		}
	}

	for _, t := range temps {
		if t.SensorKey != diskTempSensorKey {
			continue
		}
		return &pb.TempStats{
			Temperature: int32(math.Round(t.Temperature)),
			High: int32(math.Round(t.High)),
			Critical: int32(math.Round(t.Critical)),
		}, nil
	}
	return &pb.TempStats{}, fmt.Errorf("could not find cpu temps")
}

func GetMemoryTemp() (*pb.TempStats, error) {
	if temps == nil {
		if err := loadTemperatures(); err != nil {
			return &pb.TempStats{}, fmt.Errorf("could not load temperatures: %w", err)
		}
	}

	for _, t := range temps {
		if t.SensorKey != memoryTempSensorKey {
			continue
		}
		return &pb.TempStats{
			Temperature: int32(math.Round(t.Temperature)),
			High: int32(math.Round(t.High)),
			Critical: int32(math.Round(t.Critical)),
		}, nil
	}
	return &pb.TempStats{}, fmt.Errorf("could not find cpu temps")
}
