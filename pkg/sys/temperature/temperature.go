package temperature

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
