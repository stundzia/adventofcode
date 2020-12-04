package day4

import (
	"fmt"
	"github.com/stundzia/adventofcode/utils"
	"regexp"
	"strconv"
	"strings"
)

var RequiredFields = map[string]bool{
	"byr": true,
	"iyr": true,
	"eyr": true,
	"hgt": true,
	"hcl": true,
	"ecl": true,
	"pid": true,
	"cid": false,
}

var eyeColors = map[string]interface{}{
	"amb": struct {}{},
	"blu": struct {}{},
	"brn": struct {}{},
	"gry": struct {}{},
	"grn": struct {}{},
	"hzl": struct {}{},
	"oth": struct {}{},
}


func fieldValid(field, value string) bool {
	switch field {
	case "byr":
		v, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		if v < 1920 || v > 2002 {
			return false
		}
		return true
	case "iyr":
		v, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		if v < 2010 || v > 2020 {
			return false
		}
		return true
	case "eyr":
		v, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		if v < 2020 || v > 2030 {
			return false
		}
		return true
	case "hgt":
		re := regexp.MustCompile("^(\\d{2,3})(cm|in)$")
		res := re.FindAllStringSubmatch(value, -1)
		if len(res) == 0 {
			return false
		}
		if len(res[0]) != 3 {
			return false
		}
		h, err := strconv.Atoi(res[0][1])
		if err != nil {
			return false
		}
		units := res[0][2]
		if units == "cm" {
			if h >= 150 && h <= 193 {
				return true
			}
		}
		if units == "in" {
			if h >= 59 && h <= 76 {
				return true
			}
		}
		return false
	case "hcl":
		valid, _ := regexp.MatchString("^#[0-9a-f]{6}$", value)
		return valid
	case "ecl":
		_, valid := eyeColors[value]
		return valid
	case "pid":
		valid, _ := regexp.MatchString("^[0-9]{9}$", value)
		return valid
	case "cid":
		return true
	}
	return false
}

func parsePassportScore(s string) int {
	fieldsValues := strings.Split(s, " ")
	fieldsPresent := []string{}
	score := 0
	for _, fv := range fieldsValues {
		fieldsPresent = append(fieldsPresent, strings.Split(fv, ":")[0])
	}
	for _, passField := range fieldsPresent {
		if required, exists := RequiredFields[passField]; required == true && exists == true {
			score++
		}
	}
	return score
}

func parsePassportFieldsValid(p string) bool {
	fieldsValues := strings.Split(p, " ")
	for _, fv := range fieldsValues {
		asd := strings.Split(fv, ":")
		if len(asd) == 2 {
			if fieldValid(asd[0], asd[1]) {
				continue
			} else {
				return false
			}
		}
	}
	return true
}


func DoSilver() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 4, "\n")
	passports := []string{}
	passport := ""
	for _, line := range input {
		if line == "" {
			passports = append(passports, passport)
			passport = ""
			continue
		}
		passport = fmt.Sprintf("%s %s", passport, line)
	}
	validCount := 0
	for _, p := range passports {
		if parsePassportScore(p) == 7 {
			validCount++
		}
	}
	return strconv.Itoa(validCount)
}

func DoGold() string {
	input, _ := utils.ReadInputFileContentsAsStringSlice(2020, 4, "\n")
	passports := []string{}
	passport := ""
	for _, line := range input {
		if line == "" {
			passports = append(passports, passport)
			passport = ""
			continue
		}
		passport = fmt.Sprintf("%s %s", passport, line)
	}
	validCount := 0
	for _, p := range passports {
		if parsePassportScore(p) == 7 && parsePassportFieldsValid(p) {
			validCount++
		}
	}
	return strconv.Itoa(validCount)
}