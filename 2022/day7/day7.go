package day7

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/stundzia/adventofcode/utils"
)

type filesystem struct {
	currentDir *dir
	root       *dir
}

type dir struct {
	name    string
	parent  *dir
	subdirs []*dir
	files   []*file
	size    int
}

type file struct {
	name string
	size int
}

func (fs *filesystem) ls(output []string) {
	for _, l := range output {
		if strings.HasPrefix(l, "dir") {
			fs.createDirInCurrentDirIfNotExists(l[4:])
			continue
		} else {
			parts := strings.Split(l, " ")
			size, _ := strconv.Atoi(parts[0])
			name := parts[1]
			fs.createFileInCurrentDirIfNotExists(name, size)
		}
	}
}

func (fs *filesystem) cd(op string) {
	switch op {
	case "..":
		fs.currentDir = fs.currentDir.parent
	case "/":
		fs.currentDir = fs.root
	default:
		for _, d := range fs.currentDir.subdirs {
			if d.name == op {
				fs.currentDir = d
				return
			}
		}
	}
}

func (fs *filesystem) createDirInCurrentDirIfNotExists(name string) {
	for _, d := range fs.currentDir.subdirs {
		if d.name == name {
			return
		}
	}
	newDir := &dir{
		name:    name,
		parent:  fs.currentDir,
		subdirs: []*dir{},
		files:   []*file{},
		size:    -1,
	}
	fs.currentDir.subdirs = append(fs.currentDir.subdirs, newDir)
}

func (fs *filesystem) createFileInCurrentDirIfNotExists(name string, size int) {
	for _, f := range fs.currentDir.files {
		if f.name == name {
			return
		}
	}
	newFile := &file{
		name: name,
		size: size,
	}
	fs.currentDir.files = append(fs.currentDir.files, newFile)
}

func (d *dir) calcSize(res *atomic.Uint64, limit int) int {
	if d.size != -1 {
		return d.size
	}
	size := 0
	for _, f := range d.files {
		size += f.size
	}
	for _, sd := range d.subdirs {
		size += sd.calcSize(res, limit)
	}
	d.size = size
	if size < limit {
		res.Add(uint64(size))
	}

	return size
}

func (d *dir) calcSizeWithSizeMap(m map[int]struct{}) int {
	if d.size != -1 {
		m[d.size] = struct{}{}
		return d.size
	}
	size := 0
	for _, f := range d.files {
		size += f.size
	}
	for _, sd := range d.subdirs {
		size += sd.calcSizeWithSizeMap(m)
	}
	d.size = size
	m[size] = struct{}{}

	return size
}

func DoSilver() string {
	lines, _ := utils.ReadInputFileContentsAsStringSlice(2022, 7, "\n")
	fs := &filesystem{
		root: &dir{
			name:    "/",
			parent:  nil,
			subdirs: []*dir{},
			files:   []*file{},
			size:    -1,
		},
	}
	fs.currentDir = fs.root

	output := []string{}
	for i, l := range lines {
		if i == 0 {
			if l != "$ cd /" {
				panic("ummmm")
			}
			continue
		}

		switch true {
		case !strings.HasPrefix(l, "$"):
			output = append(output, l)
			continue
		case strings.HasPrefix(l, "$ cd"):
			if len(output) > 0 {
				fs.ls(output)
				output = []string{}
			}
			fs.cd(l[5:])
			continue
		}

	}
	if len(output) > 0 {
		fs.ls(output)
		output = []string{}
	}
	res := &atomic.Uint64{}

	fs.root.calcSize(res, 100000)

	return fmt.Sprintf("Solution: %d", res.Load())
}

func DoGold() string {
	lines, _ := utils.ReadInputFileContentsAsStringSlice(2022, 7, "\n")
	fs := &filesystem{
		root: &dir{
			name:    "/",
			parent:  nil,
			subdirs: []*dir{},
			files:   []*file{},
			size:    -1,
		},
	}
	fs.currentDir = fs.root

	output := []string{}
	for i, l := range lines {
		if i == 0 {
			if l != "$ cd /" {
				panic("ummmm")
			}
			continue
		}

		switch true {
		case !strings.HasPrefix(l, "$"):
			output = append(output, l)
			continue
		case strings.HasPrefix(l, "$ cd"):
			if len(output) > 0 {
				fs.ls(output)
				output = []string{}
			}
			fs.cd(l[5:])
			continue
		}
	}
	if len(output) > 0 {
		fs.ls(output)
		output = []string{}
	}

	available := 70000000
	needed := 30000000
	sizeMap := map[int]struct{}{}
	rootSize := fs.root.calcSizeWithSizeMap(sizeMap)
	need := needed - (available - rootSize)
	res := math.MaxInt
	for size, _ := range sizeMap {
		if size >= need && size < res {
			res = size
		}
	}

	return fmt.Sprintf("Solution: %d", res)
}
