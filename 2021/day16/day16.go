package day16

import (
	"fmt"
	"math"
	"strconv"

	"github.com/stundzia/adventofcode/utils"
)

const (
	packetTypeLiteralID = 4
)

type transmission struct {
	bits string
}

func (t *transmission) addBits(hexNum string) {
	num, err := strconv.ParseInt(hexNum, 16, 64)
	if err != nil {
		fmt.Println("FUUUCK!: ", err)
	}
	bits := strconv.FormatInt(num, 2)
	for len(bits) < 4 {
		bits = "0" + bits
	}
	t.bits += bits
}

type packet struct {
	version      int64
	typeID       int64
	val          int64
	lengthTypeID int
	lengthNum    int64
	subPackets   []*packet
}

func (p *packet) versionSum() int64 {
	var vSum int64 = p.version
	for _, sp := range p.subPackets {
		vSum += sp.versionSum()
	}
	return vSum
}

func (p *packet) getValue() int64 {
	var res int64
	switch p.typeID {
	case 0:
		for _, sp := range p.subPackets {
			res += sp.getValue()
		}
	case 1:
		var val int64 = 1
		for _, sp := range p.subPackets {
			val = val * sp.getValue()
		}
		res = val
	case 2:
		var minVal int64 = math.MaxInt
		for _, sp := range p.subPackets {
			val := sp.getValue()
			if val < minVal {
				minVal = val
			}
		}
		res = minVal
	case 3:
		var maxVal int64 = 0
		for _, sp := range p.subPackets {
			val := sp.getValue()
			if val > maxVal {
				maxVal = val
			}
		}
		res = maxVal
	case 4:
		return p.val
	case 5:
		if p.subPackets[0].getValue() > p.subPackets[1].getValue() {
			res = 1
		}
	case 6:
		if p.subPackets[0].getValue() < p.subPackets[1].getValue() {
			res = 1
		}
	case 7:
		if p.subPackets[0].getValue() == p.subPackets[1].getValue() {
			res = 1
		}
	}

	return res
}

func (t *transmission) parsePacket(offset int) (*packet, int) {
	p := &packet{subPackets: []*packet{}}
	endOffset := offset + 6
	versionS := t.bits[offset : offset+3]
	versionI, err := strconv.ParseInt(versionS, 2, 64)
	if err != nil {
		fmt.Println("version err: ", err)
	}
	p.version = versionI
	typeS := t.bits[offset+3 : offset+6]
	typeI, err := strconv.ParseInt(typeS, 2, 64)
	if err != nil {
		fmt.Println("type err: ", err)
	}
	p.typeID = typeI

	if p.typeID == packetTypeLiteralID {
		content := ""
		for i := offset + 6; ; i += 5 {
			part := t.bits[i : i+5]
			content += part[1:]
			endOffset += 5
			if part[0] == '0' {
				break
			}
		}
		p.val, err = strconv.ParseInt(content, 2, 64)
		if err != nil {
			fmt.Println("val err: ", err)
		}
	} else {
		p.lengthTypeID, _ = strconv.Atoi(string(t.bits[offset+6]))
		endOffset++
		if p.lengthTypeID == 0 {
			l := t.bits[offset+7 : offset+22]
			endOffset += 15
			p.lengthNum, _ = strconv.ParseInt(l, 2, 64)
			end := endOffset + int(p.lengthNum)
			for i := endOffset; i < end; {
				newPacket, newOff := t.parsePacket(i)
				p.subPackets = append(p.subPackets, newPacket)
				i = newOff
			}
			endOffset = end
		}
		if p.lengthTypeID == 1 {
			l := t.bits[offset+7 : offset+18]
			endOffset += 11
			p.lengthNum, _ = strconv.ParseInt(l, 2, 64)
			for i := 0; i < int(p.lengthNum); i++ {
				newPacket, newOff := t.parsePacket(endOffset)
				p.subPackets = append(p.subPackets, newPacket)
				endOffset = newOff
			}
		}
	}

	return p, endOffset
}

func DoSilver() string {
	num, _ := utils.ReadInputFileContentsAsString(2021, 16)
	t := transmission{bits: ""}
	for _, n := range num {
		t.addBits(string(n))
	}
	pp, _ := t.parsePacket(0)
	return fmt.Sprintf("Solution: %d", pp.versionSum())
}

func DoGold() string {
	num, _ := utils.ReadInputFileContentsAsString(2021, 16)
	t := transmission{bits: ""}
	for _, n := range num {
		t.addBits(string(n))
	}
	pp, _ := t.parsePacket(0)
	return fmt.Sprintf("Solution: %d", pp.getValue())
}
