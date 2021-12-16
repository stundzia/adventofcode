package day13

import (
	"sort"
	"strconv"
	"strings"
)

type BusSchedule struct {
	Buses map[int]*Bus
}

type Bus struct {
	Schedule       *BusSchedule
	ID             int
	DepartureTimes map[int]interface{}
}

func NewBusSchedule(buses string) *BusSchedule {
	busSchedule := &BusSchedule{
		Buses: map[int]*Bus{},
	}
	bs := strings.Split(buses, ",")
	busSequenceId := 0
	for _, b := range bs {
		if b == "x" {
			busSequenceId++
			continue
		}
		busId, err := strconv.Atoi(b)
		if err == nil {
			bus := &Bus{
				Schedule:       busSchedule,
				ID:             busId,
				DepartureTimes: map[int]interface{}{},
			}
			busSchedule.Buses[busSequenceId] = bus
			busSequenceId++
		}
	}
	return busSchedule
}

func (bs *BusSchedule) IsSequentialDeparture(timestamp int) bool {
	for i, bus := range bs.Buses {
		if !bus.LeavesAtTimestampOpt(timestamp + i) {
			return false
		}
	}
	return true
}

func (bs *BusSchedule) GenerateSchedules(from int, to int) {
	for _, bus := range bs.Buses {
		bus.GenerateSchedule(from, to)
	}
}

func (b *Bus) GenerateSchedule(from int, to int) {
	for i := from; i < to; i++ {
		if i%b.ID == 0 {
			b.DepartureTimes[i] = struct{}{}
		}
	}
}

func (b *Bus) LeavesAtTimestamp(timestamp int) bool {
	_, ok := b.DepartureTimes[timestamp]
	return ok
}

func (b *Bus) LeavesAtTimestampOpt(timestamp int) bool {
	return timestamp%b.ID == 0
}

func (b *Bus) LeavesAtTimestampOptBig(timestamp uint64) bool {
	return timestamp%uint64(b.ID) == 0
}

func (bs *BusSchedule) IsSequentialDepartureBig(timestamp uint64) bool {
	for i, bus := range bs.Buses {
		if !bus.LeavesAtTimestampOptBig(timestamp + uint64(i)) {
			return false
		}
	}
	return true
}

func (bs *BusSchedule) SuggestSequential(start uint64, maxBus int, step uint64) uint64 {
	var t uint64
mainLoop:
	for t = start + step; ; t += step {
	busLoop:
		for i, bus := range bs.Buses {
			if i <= maxBus {
				if !bus.LeavesAtTimestampOptBig(t + uint64(i)) {
					continue mainLoop
				}
			} else {
				continue busLoop
			}
		}
		return t
	}
}

func (bs *BusSchedule) getFirstSequentialAndDiffToNext(start uint64, maxBus int, step uint64) (uint64, uint64) {
	t0 := bs.SuggestSequential(start, maxBus, step)
	t1 := bs.SuggestSequential(t0, maxBus, step)
	return t0, t1 - t0
}

func (bs *BusSchedule) GetFirstSequentialTimestamp() uint64 {
	busSequences := make([]int, 0, len(bs.Buses))
	for k := range bs.Buses {
		busSequences = append(busSequences, k)
	}
	sort.Ints(busSequences)
	startFrom, step := bs.getFirstSequentialAndDiffToNext(0, busSequences[0], 1)
	for i := 1; i < len(busSequences); i++ {
		startFrom, step = bs.getFirstSequentialAndDiffToNext(startFrom, busSequences[i], step)
	}
	return startFrom
}
