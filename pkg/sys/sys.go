package sys

import (
	"time"

	"github.com/golang/protobuf/proto"
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
			bytes, err := proto.Marshal(s.GetHardwareUsage())
			if err != nil {
				s.log.Errorf("could not marshal hardware load to proto %v")
			}
			runtime.Events.Emit("hwUsage", bytes)
			time.Sleep(1 * time.Second)
		}
	}()

	return nil
}
