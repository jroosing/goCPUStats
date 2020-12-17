package sys

import (
	"time"

	"github.com/wailsapp/wails"
)

// Stats .
type Stats struct {
	log *wails.CustomLogger
}

// WailsInit .
func (s *Stats) WailsInit(runtime *wails.Runtime) error {
	s.log = runtime.Log.New("Stats")

	go func() {
		for {
			runtime.Events.Emit("hwUsage", s.GetHardwareUsage())
			time.Sleep(3 * time.Second)
		}
	}()

	return nil
}
