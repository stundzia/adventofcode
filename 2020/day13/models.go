package day13

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
	"strconv"
	"strings"
)

type BusSchedule struct {
	Buses map[int]*Bus
}

type Bus struct {
	Schedule *BusSchedule
	ID int
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

func (b *Bus) GenerateSchedule(from int, to int)  {
	for i := from; i < to; i++ {
		if i % b.ID == 0 {
			b.DepartureTimes[i] = struct {}{}
		}
	}
}

func (b *Bus) LeavesAtTimestamp(timestamp int) bool {
	_, ok := b.DepartureTimes[timestamp]
	return ok
}

func (b *Bus) LeavesAtTimestampOpt(timestamp int) bool {
	return timestamp % b.ID == 0
}

func (b *Bus) LeavesAtTimestampOptBig(timestamp uint64) bool {
	return timestamp % uint64(b.ID) == 0
}

func (bs *BusSchedule) IsSequentialDepartureBig(timestamp uint64) bool {
	for i, bus := range bs.Buses {
		if !bus.LeavesAtTimestampOptBig(timestamp + uint64(i)) {
			return false
		}
	}
	return true
}

func (bs *BusSchedule) FindFirstSequentialDeparture() {
	forLcm := []int{}
	forLcm1 := []int{}
	forLcm2 := []int{}
	forLcm3 := []int{}
	for i, bus := range bs.Buses {
		forLcm = append(forLcm, bus.ID)
		forLcm1 = append(forLcm, bus.ID + i)
		forLcm3 = append(forLcm, bus.ID - i)
		if i > 0 {
			forLcm2 = append(forLcm2, i)
		}
	}
	// 1068788
	lcm := utils.LCM(forLcm[0], forLcm[1], forLcm[2:]...)
	gcd := utils.GCD(lcm, forLcm[0])

	lcmFirstLast := utils.LCM(29, 41)
	fmt.Println("lcm first last: ", lcmFirstLast)
	lcm2 := utils.LCM(forLcm2[0], forLcm2[1], forLcm2[2:]...)
	fmt.Println(gcd)
	fmt.Println("lcm: ", utils.LCM(forLcm[0], forLcm[1], forLcm[2:]...))
	fmt.Println("lcm1: ", utils.LCM(forLcm1[0], forLcm1[1], forLcm1[2:]...))
	fmt.Println("lcm2: ", utils.LCM(forLcm2[0], forLcm2[1], forLcm2[2:]...))
	fmt.Println("lcm3: ", utils.LCM(forLcm3[0], forLcm3[1], forLcm3[2:]...))
	fmt.Println(lcm2)
	fmt.Println(lcm * 409)
	fmt.Println(lcm / 409)
	fmt.Println("Actual answer: 1068781")

}

func (bs *BusSchedule) SequentialArrivals() {
	maxSeq := 0
	for seq, _ := range bs.Buses {
		if seq > maxSeq {
			maxSeq = seq
		}
	}
	for i := 0; i <= maxSeq; i++ {
		line := ""
		for s, bus := range bs.Buses {
			if s == i {
				line += fmt.Sprintf("|Main %d : %d |", s, bus.ID)
			}
			if s != i && i > 0 && i % bus.ID == 0 {
				line += fmt.Sprintf("|Sec %d : %d |",  s, bus.ID)
			}
		}
		fmt.Println(line)
		line = ""
	}
}

func (bs *BusSchedule) SuggestSequential(start uint64, maxBus int, step uint64) uint64 {
	var t uint64
	mainLoop:
	for t = start + step;; t+=step {
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

func (bs *BusSchedule) GetDiff(start uint64, maxBus int, step uint64) (uint64, uint64) {
	t0 := bs.SuggestSequential(start, maxBus, step)
	fmt.Println("t0: ", t0)
	t1 := bs.SuggestSequential(t0, maxBus, step)
	return t0, t1 - t0
}