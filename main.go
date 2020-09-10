package main

import (
	"github.com/stundzia/adventofcode/2019/day1"
	"github.com/stundzia/adventofcode/utils"
)

func main() {
	utils.RunWithTimeMetricsAndPrintOutput(day1.DoSilver)
	utils.RunWithTimeMetricsAndPrintOutput(day1.DoGold)
}