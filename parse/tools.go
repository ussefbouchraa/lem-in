package parse

import (
	"os"
	"strconv"
	"strings"
	"unicode"
)

func IsValidLine(line string) bool {
	if len(line) < 1 {
		return false
	}

	for _, r := range line {
		if !unicode.IsPrint(r) && !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

func IsComment(line string) bool {
	return strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "##")
}

func IsOnlyInt(line string) (int, bool) {
	for _, r := range line {
		if !unicode.IsDigit(r) {
			return -1, false
		}
	}

	nbr, err := strconv.Atoi(line)
	if err != nil || nbr <= 0 {
		return -1, false
	}
	return nbr, true
}

func IsEmptyFile(file *os.File) bool {
	stat, err := file.Stat()
	if err != nil {
		return true
	}
	return stat.Size() == 0
}

func IsRoom(line string) bool {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return false
	}

	if strings.HasPrefix(parts[0], "L") || strings.HasPrefix(parts[0], "#") {
		return false
	}
	
	x, err1 := strconv.Atoi(parts[1])
	y, err2 := strconv.Atoi(parts[2])

	if err1 != nil || err2 != nil || x < 0 || y < 0 {
    return false
}

	return true
}

func IsLink(line string) bool {
	if strings.Count(line, "-") != 1 {
		return false
	}

	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return false
	}

	if parts[0] == "" || parts[1] == "" {
		return false
	}

	if strings.HasPrefix(parts[0], "L") || strings.HasPrefix(parts[1], "L") {
		return false
	}

	return true
}