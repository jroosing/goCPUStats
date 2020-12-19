// +build darwin

package temp

import (
	"fmt"
	"math"
)

const (
	cpuTempSensorKey = "TC0P"
	gpuTempSensorKey = "TG0P"
	memoryTempSensorKey = "TM0P"
	diskTempSensorKey = "TH0P"
)

func GetCpuTemp() (TempStats, error) {
	if temps == nil {
		if err := loadTemperatures(); err != nil {
			return TempStats{}, fmt.Errorf("could not load temperatures: %w", err)
		}
	}

	for _, t := range temps {
		if t.SensorKey != cpuTempSensorKey {
			continue
		}
		return TempStats{
			Temperature: int16(math.Round(t.Temperature)),
			High: int16(math.Round(t.High)),
			Critical: int16(math.Round(t.Critical)),
		}, nil
	}
	return TempStats{}, fmt.Errorf("could not find cpu temps")
}

func GetGpuTemp() (TempStats, error) {
	if temps == nil {
		if err := loadTemperatures(); err != nil {
			return TempStats{}, fmt.Errorf("could not load temperatures: %w", err)
		}
	}

	for _, t := range temps {
		if t.SensorKey != gpuTempSensorKey {
			continue
		}
		return TempStats{
			Temperature: int16(math.Round(t.Temperature)),
			High: int16(math.Round(t.High)),
			Critical: int16(math.Round(t.Critical)),
		}, nil
	}
	return TempStats{}, fmt.Errorf("could not find cpu temps")
}

func GetDiskTemp() (TempStats, error) {
	if temps == nil {
		if err := loadTemperatures(); err != nil {
			return TempStats{}, fmt.Errorf("could not load temperatures: %w", err)
		}
	}

	for _, t := range temps {
		if t.SensorKey != diskTempSensorKey {
			continue
		}
		return TempStats{
			Temperature: int16(math.Round(t.Temperature)),
			High: int16(math.Round(t.High)),
			Critical: int16(math.Round(t.Critical)),
		}, nil
	}
	return TempStats{}, fmt.Errorf("could not find cpu temps")
}

func GetMemoryTemp() (TempStats, error) {
	if temps == nil {
		if err := loadTemperatures(); err != nil {
			return TempStats{}, fmt.Errorf("could not load temperatures: %w", err)
		}
	}

	for _, t := range temps {
		if t.SensorKey != memoryTempSensorKey {
			continue
		}
		return TempStats{
			Temperature: int16(math.Round(t.Temperature)),
			High: int16(math.Round(t.High)),
			Critical: int16(math.Round(t.Critical)),
		}, nil
	}
	return TempStats{}, fmt.Errorf("could not find cpu temps")
}
