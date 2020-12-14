package day14

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Computer struct {
	Memory map[int]int64
	Mask1 int64
	Mask0 int64
	MemoryMask string
	Version int
}

var commandRegex = regexp.MustCompile("([a-z]{3,4})(\\[\\d+\\])?\\s=\\s(.*)")


func NewComputer(version int) *Computer {
	c := &Computer{
		Memory: map[int]int64{},
		Mask1:   0,
		Mask0:   0,
		MemoryMask: "",
		Version: version,
	}
	return c
}

func (c *Computer) handleMaskCommand(value string) {
	c.MemoryMask = value
	mask1Str := strings.Replace(value, "X", "0", -1)
	mask1, _ := strconv.ParseInt(mask1Str, 2, 64)
	c.Mask1 = mask1
	mask0Str := strings.Replace(value, "1", "X", -1)
	mask0Str = strings.Replace(mask0Str, "0", "1", -1)
	mask0Str = strings.Replace(mask0Str, "X", "0", -1)
	mask0, _ := strconv.ParseInt(mask0Str, 2, 64)
	c.Mask1 = mask1
	c.Mask0 = mask0
}

func (c *Computer) putToMemory(address int, value int64) {
	val := value | c.Mask1
	val = val & ^c.Mask0
	c.Memory[address] = val
}

func (c *Computer) GetMemoryMasks() (stable int64, floating []int64, floatingBase int64) {
	mStableStr := strings.Replace(c.MemoryMask, "X", "0", -1)
	mFloatingStr := strings.Replace(c.MemoryMask, "1", "0", -1)
	mFloatingBaseStr := strings.Replace(c.MemoryMask, "X", "1", -1)
	floating = c.memoryMaskStringToMasks(mFloatingStr)
	stable, _ = strconv.ParseInt(mStableStr, 2, 64)
	floatingBase, _ = strconv.ParseInt(mFloatingBaseStr, 2, 64)
	return stable, floating, floatingBase
}

func (c *Computer) putToMemoryV2(address int, value int64) {
	address64 := int64(address)
	stable, floating, floatingBase := c.GetMemoryMasks()
	address64 = address64 | stable
	address64 = address64 & ^floatingBase
	for _, m := range floating {
		add := address64 | m
		c.Memory[int(add)] = value
	}
}

func power(n int64, pow int) int64 {
	res := n
	for i := 0; i < pow; i++ {
		res *= n
	}
	return res
}

func (c *Computer) memoryMaskStringToMasks(mStr string) (masks []int64) {
	maskMap := map[int64]struct{}{}
	floatingBitCount := strings.Count(mStr, "X")
	if floatingBitCount == 0 {
		return masks
	}
	bits := []string{"1", "0"}
	varCount := int(power(2, floatingBitCount))
	rand.Seed(time.Now().Unix())
	for i := 0;i < varCount * 500; i++ {
		m := c.MemoryMask
		for i := 0; i < floatingBitCount; i++ {
			m = strings.Replace(m, "X", bits[rand.Intn(len(bits))], 1)
		}
		mInt, _ := strconv.ParseInt(m, 2, 64)
		maskMap[mInt] = struct{}{}
		if len(maskMap) != varCount {
			continue
		} else {
			break
		}
	}
	for mask, _ := range maskMap {
		masks = append(masks, mask)
	}
	return masks
}

func (c *Computer) handleMemoryCommand(address, value string) {
	// Lazy, I know
	address = strings.Replace(address, "[", "", -1)
	address = strings.Replace(address, "]", "", -1)
	addressInt, _ := strconv.Atoi(address)
	valueInt, _ := strconv.Atoi(value)
	if c.Version != 2 {
		c.putToMemory(addressInt, int64(valueInt))
	} else {
		c.putToMemoryV2(addressInt, int64(valueInt))
	}
}

func (c *Computer) parseInputCommand(command string) {
	parts := commandRegex.FindAllStringSubmatch(command, -1)
	switch parts[0][1] {
	case "mask":
		c.handleMaskCommand(parts[0][3])
		break
	case "mem":
		c.handleMemoryCommand(parts[0][2], parts[0][3])
		break
	default:
		fmt.Println("Unknown command: ", parts[0][1])
		break
	}
}

func (c *Computer) MemorySum() int {
	var sum int64
	for _, val := range c.Memory {
		sum += val
	}
	return int(sum)
}