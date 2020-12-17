package temp

import "github.com/shirou/gopsutil/v3/host"

var temps []host.TemperatureStat

func loadTemperatures() error {
	tempStat, err := host.SensorsTemperatures()
	if err != nil {
		return err
	}
	temps = tempStat
	return nil
}

type TempStats struct {
	Temperature int16 `json:"temperature"`
	High        int16 `json:"sensorHigh"`
	Critical    int16 `json:"sensorCritical"`
}
