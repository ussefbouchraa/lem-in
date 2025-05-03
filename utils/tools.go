package lemin

import (
	"strconv"
	"unicode"
	"strings"
	"os"
)



func IsValidLine(line string) bool{
	if len(line) <= 2  { return false }

	for _, r := range line {
		if !unicode.IsPrint(r) &&  !unicode.IsSpace(r) {
			return false 
		}
	}
	return true
}

func IsComment(line string) bool{

	if line[:1] == "#" && line[1:2] != "#" {
		return true
	}
	return false
}


func IsOnlyInt(line string) (int, bool){
	for _,it := range(line){
		if !unicode.IsNumber(it){
			return -1, false
		}
	}
	nbr, err := strconv.Atoi(line)
	if err != nil || nbr <= 0 { return -1, false }

	return  nbr, true
}

func IsEmptyFile( file *os.File) bool{
	stat, _ := file.Stat()
	
	if  stat.Size() == 0{
		return true
	}

	return false
}




func IsRoom(line string) bool {
	if len(line) < 5 { return false }

	parts := strings.Fields(line)
	if len(parts) != 3 { return false }

	name, x, y := parts[0], parts[1], parts[2]

	if strings.HasPrefix(name, "L") || strings.HasPrefix(name, "#") { return false }

	for _, r := range name {
		if !unicode.IsPrint(r) || unicode.IsSpace(r) {
			return false
		}
	}

	if _, err := strconv.Atoi(x); err != nil { return false }
	if _, err := strconv.Atoi(y); err != nil { return false }

	return true
}

func IsLink(line string) bool {
	if strings.Count(line, "-") != 1 { return false }

	parts := strings.Split(line, "-")
	if len(parts) != 2 { return false }

	if parts[0] == "" || parts[1] == "" || parts[0][0] == 'L' || parts[1][0] == 'L' {
		return false
		}

	return true
}

