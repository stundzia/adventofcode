Advent of Code Solutions
===

Just a repo holding solutions for various years and days of advent of code puzzles (https://adventofcode.com/) 
written in GoLang.

This repo contains packages named after years for the according AoC year and those packages have sub-packages for
the days. So package 2019 (github.com/stundzia/adventofcode/2019) contains packages day1, day2, day3 etc. for the
according puzzle day, e.g. github.com/stundzia/adventofcode/2019/day1 contains the solutions for the first day.
The function that solves the first part is always called DoSilver (i.e. do the solution for the silver star part (part 1)) and
DoGold (i.e. do the solution for the gold star part (part 2)).

Running a specific day involves editing `main.go`. You will find it has the following lines:
```go
package main

import (
	"github.com/stundzia/adventofcode/2020/day1"
	"github.com/stundzia/adventofcode/utils"
)


func main() {
	utils.RunWithTimeMetricsAndPrintOutput(day1.DoSilver)
	utils.RunWithTimeMetricsAndPrintOutput(day1.DoGold)
}
```

So to run e.g. 2020s day 14 solutions you would need to change it to:
```go
package main

import (
	"github.com/stundzia/adventofcode/2020/day14"
	"github.com/stundzia/adventofcode/utils"
)


func main() {
	utils.RunWithTimeMetricsAndPrintOutput(day14.DoSilver)
	utils.RunWithTimeMetricsAndPrintOutput(day14.DoGold)
}
```

Then just execute the following command in repo root:
```
go run main.go
```

And you should see output similar to this:
```
Solution is:  Solution: 793524
Solution took:  294.77µs
Solution is:  Solution: 61515678
Solution took:  3.714729ms
```