package day13

import (
	"strconv"
	"strings"
)

type BusSchedule struct {
	Buses []*Bus
}

type Bus struct {
	Schedule *BusSchedule
	ID int
	DepartureTimes map[int]interface{}
}

func NewBusSchedule(buses string) *BusSchedule {
	busSchedule := &BusSchedule{}
	bs := strings.Split(buses, ",")
	for _, b := range bs {
		if b != "x" {
			busId, err := strconv.Atoi(b)
			if err == nil {
				bus := &Bus{
					Schedule:       busSchedule,
					ID:             busId,
					DepartureTimes: map[int]interface{}{},
				}
				busSchedule.Buses = append(busSchedule.Buses, bus)
			}
		}
	}
	return busSchedule
}

func (b *Bus) GenerateScheduleUntil(timestamp int)  {
	for i := 0; i < timestamp; i++ {
		if i % b.ID == 0 {
			b.DepartureTimes[i] = struct {}{}
		}
	}
}

func (b *Bus) LeavesAtTimestamp(timestamp int) bool {
	_, ok := b.DepartureTimes[timestamp]
	return ok
}
