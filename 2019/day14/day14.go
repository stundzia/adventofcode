package day14

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/stundzia/adventofcode/utils"
)

type Chemical struct {
	name          string
	inputs        map[*Chemical]uint64
	producedCount uint64
}

type Chemicals struct {
	chemicals map[string]*Chemical
	storage   map[*Chemical]*atomic.Uint64
	oreUsed   int
}

func (cs *Chemicals) setupStorage() {
	for _, chem := range cs.chemicals {
		cs.storage[chem] = &atomic.Uint64{}
	}
}

func (cs *Chemicals) produce(chem *Chemical, doNotProduceOre bool) {
	if chem.name == "ORE" {
		if doNotProduceOre {
			cs.printStorage()
			fmt.Println("Fuel in storage: ", cs.storage[cs.getOrCreateChemical("FUEL")].Load())
			os.Exit(0)
		}
		cs.oreUsed++
		cs.storage[chem].Add(1)
		return
	}
	for i, c := range chem.inputs {
		for {
			if cs.storage[i].Load() >= c {
				cs.storage[i].Add(-c)
				break
			}
			cs.produce(i, doNotProduceOre)
		}

	}
	cs.storage[chem].Add(chem.producedCount)
}

func (cs *Chemicals) getOrCreateChemical(name string) *Chemical {
	c, found := cs.chemicals[name]
	if !found {
		c = &Chemical{
			name:          name,
			inputs:        map[*Chemical]uint64{},
			producedCount: 1,
		}
		cs.chemicals[name] = c
	}
	return c
}

func (cs *Chemicals) updateChemicalInputs(chem *Chemical, inputsMap map[string]uint64, producedCount uint64) {
	for input, count := range inputsMap {
		chem.inputs[cs.getOrCreateChemical(input)] = count
	}
	chem.producedCount = producedCount
}

func (cs *Chemicals) printChems() {
	for _, chem := range cs.chemicals {
		fmt.Println("\nName: ", chem.name, " Produced: ", chem.producedCount)
		fmt.Println("\nInputs: ")
		for n, c := range chem.inputs {
			fmt.Println(c, " ", n.name)
		}
	}
}

func (cs *Chemicals) printStorage() {
	for chem, count := range cs.storage {
		fmt.Println(chem.name, " : ", count)
	}
}

func (cs *Chemicals) parseReaction(reaction string) {
	parts := strings.Split(reaction, " => ")
	inputChems := strings.Split(parts[0], ", ")
	resChemStr := parts[1]
	resChemName, resCount := parseChem(resChemStr)
	resChem := cs.getOrCreateChemical(resChemName)
	inputMap := map[string]uint64{}
	for _, chem := range inputChems {
		n, c := parseChem(chem)
		inputMap[n] = c
	}
	cs.updateChemicalInputs(resChem, inputMap, resCount)
}

func (cs *Chemicals) produceTillOreLasts(chem *Chemical, minOre uint64, storageChannel chan map[string]uint64) {
	ore := cs.getOrCreateChemical("ORE")
	for {
		cs.produce(chem, true)
		oreCount := cs.storage[ore].Load()
		if oreCount%5 == 0 {
			fmt.Println("Remaining: ", oreCount/100000000, "%")
		}
		if oreCount < minOre {
			res := map[string]uint64{}
			for chem, v := range cs.storage {
				res[chem.name] = v.Load()
			}
			storageChannel <- res
			return
		}
	}
}

func (cs *Chemicals) joinStorage(otherStorage map[string]uint64) {
	for chem, _ := range cs.storage {
		cs.storage[chem].Add(otherStorage[chem.name])
	}
}

func parseChem(chem string) (string, uint64) {
	parts := strings.Split(chem, " ")
	count, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	return parts[1], uint64(count)
}

func createChems() *Chemicals {
	reactions, _ := utils.ReadInputFileContentsAsStringSlice(2019, 14, "\n")
	cs := &Chemicals{chemicals: map[string]*Chemical{}, storage: map[*Chemical]*atomic.Uint64{}}
	for _, reaction := range reactions {
		cs.parseReaction(reaction)
	}
	cs.setupStorage()
	return cs
}

func DoSilver() string {
	reactions, _ := utils.ReadInputFileContentsAsStringSlice(2019, 14, "\n")
	cs := Chemicals{chemicals: map[string]*Chemical{}, storage: map[*Chemical]*atomic.Uint64{}}
	for _, reaction := range reactions {
		cs.parseReaction(reaction)
	}
	cs.setupStorage()
	cs.produce(cs.getOrCreateChemical("FUEL"), false)

	return strconv.Itoa(cs.oreUsed)
}

func DoGold() string {
	reactions, _ := utils.ReadInputFileContentsAsStringSlice(2019, 14, "\n")
	cs := &Chemicals{chemicals: map[string]*Chemical{}, storage: map[*Chemical]*atomic.Uint64{}}
	for _, reaction := range reactions {
		cs.parseReaction(reaction)
	}
	cs.setupStorage()

	master := createChems()

	storageChannel := make(chan map[string]uint64)
	haveOre := 1000000000000
	for i := 0; i < 100; i++ {
		cs := createChems()
		cs.storage[cs.getOrCreateChemical("ORE")].Store(uint64(haveOre / 100))
		go cs.produceTillOreLasts(cs.getOrCreateChemical("FUEL"), 997302, storageChannel)
	}

	done := 0
	for {
		time.Sleep(2 * time.Second)
		if done < 100 {
			select {
			case sto := <-storageChannel:
				master.joinStorage(sto)
				done++
			default:
			}
		} else {
			break
		}
	}

	for {
		master.printStorage()
		master.produce(master.getOrCreateChemical("FUEL"), true)
	}

	return strconv.Itoa(1)
}
