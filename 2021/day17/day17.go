package day17

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
	"strconv"
	"strings"
)

type ocean struct {
	grid map[string]bool
	targetBottom int
	targetRight int
	targetLeft int
	probe *probe
}

type probe struct {
	ocean *ocean
	positionX int
	positionY int
	velocityX int
	velocityY int
	highestY int
	hitTarget bool
}

func (p *probe) step() {
	p.positionX += p.velocityX
	p.positionY += p.velocityY
	if p.positionY > p.highestY {
		p.highestY = p.positionY
	}
	if p.inTarget() {
		p.hitTarget = true
	}
	if p.velocityX > 0 {
		p.velocityX--
	}
	p.velocityY--
}

func (p *probe) inTarget() bool {
	return p.ocean.grid[fmt.Sprintf("%d,%d", p.positionX, p.positionY)] == true
}

func newProbe(ocean *ocean) *probe {
	return &probe{
		ocean:     ocean,
		positionX: 0,
		positionY: 0,
		velocityX: 0,
		velocityY: 0,
		highestY:  0,
		hitTarget: false,
	}
}

func (p *probe) reset(velocityX, velocityY int) {
	p.positionX = 0
	p.positionY = 0
	p.velocityX = velocityX
	p.velocityY = velocityY
	p.highestY = 0
	p.hitTarget = false
}

func (p *probe) wontHit() bool {
	if p.positionY < p.ocean.targetBottom {
		return true
	}
	if p.positionX > p.ocean.targetRight {
		return true
	}
	if p.velocityX == 0 && p.positionX < p.ocean.targetLeft {
		return true
	}
	return false
}

func (p *probe) runTest(velocityX, velocityY int) (bool, int) {
	p.reset(velocityX, velocityY)
	for p.inTarget() != true && p.wontHit() != true {
		p.step()
	}
	if p.hitTarget {
		return true, p.highestY
	}
	return false, 0
}

func newOcean(minX, maxX, minY, maxY int) *ocean {
	o := &ocean{
		grid: map[string]bool{},
		probe: nil,
		targetBottom: minY,
		targetRight: maxX,
		targetLeft: minX,
	}
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			o.grid[fmt.Sprintf("%d,%d", x, y)] = true
		}
	}
	return o
}

func DoSilver() string {
	//minX, maxX, minY, maxY := 282, 314, -80, -45
	//target area: x=282..314, y=-80..-45
	targetInput, _ := utils.ReadInputFileContentsAsString(2021, 20)
	ti := strings.Split(targetInput, ": ")
	ti = strings.Split(ti[1], ", ")
	tiX := strings.Split(strings.Replace(ti[0], "x=", "", -1), "..")
	tiY := strings.Split(strings.Replace(ti[1], "y=", "", -1), "..")
	minX, _ := strconv.Atoi(tiX[0])
	maxX, _ := strconv.Atoi(tiX[1])
	minY, _ := strconv.Atoi(tiY[0])
	maxY, _ := strconv.Atoi(tiY[1])
	ocean := newOcean(minX, maxX, minY, maxY)
	ocean.probe = newProbe(ocean)
	highestY := 0
	for x := 15; x < 300; x++ {
		for y := 5; y < 3300; y++ {
			_, res := ocean.probe.runTest(x, y)
			if res > highestY {
				highestY = res
			}
		}
	}
	return fmt.Sprintf("Solution: %d", highestY)
}

func DoGold() string {
	targetInput, _ := utils.ReadInputFileContentsAsString(2021, 20)
	ti := strings.Split(targetInput, ": ")
	ti = strings.Split(ti[1], ", ")
	tiX := strings.Split(strings.Replace(ti[0], "x=", "", -1), "..")
	tiY := strings.Split(strings.Replace(ti[1], "y=", "", -1), "..")
	minX, _ := strconv.Atoi(tiX[0])
	maxX, _ := strconv.Atoi(tiX[1])
	minY, _ := strconv.Atoi(tiY[0])
	maxY, _ := strconv.Atoi(tiY[1])
	ocean := newOcean(minX, maxX, minY, maxY)
	ocean.probe = newProbe(ocean)
	okCount := 0
	for x := 20; x < 350; x++ {
		for y := -81; y < 15300; y++ {
			hit, _ := ocean.probe.runTest(x, y)
			if hit {
				okCount++
			}
		}
	}
	return fmt.Sprintf("Solution: %d", okCount)
}
