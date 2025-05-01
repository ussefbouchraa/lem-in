package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"

	F "fmt"
	U "lemin/utils"
	S "strings"
)


type s_Global struct{
	Ants      	  int;
	FirstRoom     s_Rooms;
	LastRoom      s_Rooms;
	rooms[]   	  s_Rooms;
	links[]   	  s_Links;
} 


type s_Rooms struct{
	r_name string;
	r_x	   int;
	r_y    int;

}

type s_Links struct {
	room1 string;
	room2 string;
}


func checkFile(arg string) bool{
	if len(arg) == 0 {
		return false ;
	}
	
	if S.Count(arg,".txt") != 1 || arg[len(arg) - 4 :] != ".txt" {
		return false ;
	}

	if _, err := os.Stat(arg); err != nil { return false }
	
	return true ;
}



func ParseFile(file *os.File) ( *s_Global, error){

	var data s_Global;
	var start, end int;
	
	scanner := bufio.NewScanner(file)
	
	if U.IsEmptyFile(file) {
		return nil, errors.New("Empty File!") 
	}

	// is number of ants
	if (scanner.Scan()){
		if nbr, stat := U.IsOnlyInt(scanner.Text()); stat {
			data.Ants = nbr;
		}else{ return nil, errors.New("invalid ants number") }
	}

	for (scanner.Scan()) {
		F.Println(len(scanner.Text()) , scanner.Text())
	
		//isComment #comment
		if  U.IsComment(scanner.Text()) { continue; }


		// standard checkLine  !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		if !U.IsValidLine(scanner.Text()){
			return nil, errors.New("Invalid Data Format ")
		}

		//isCommand ##start or ##end or ##err
		if S.HasPrefix(scanner.Text(), "##"){
			if scanner.Text() == "##start" {
				start++
				if scanner.Scan() && start == 1 && end == 0  && U.IsRoom(scanner.Text()) {
					parts := strings.Fields(scanner.Text())
			
					r_x, _:= strconv.Atoi(parts[1])
					r_y, _:= strconv.Atoi(parts[2])

					data.FirstRoom = s_Rooms{parts[0],r_x, r_y}	
					data.rooms = append(data.rooms, data.FirstRoom)
					
				}else{ return nil, errors.New("Invalid Data Format!") }

			}else if scanner.Text() == "##end" {
				end++
				if scanner.Scan() && start == 1 && end == 1 && U.IsRoom(scanner.Text()) {
					parts := strings.Fields(scanner.Text())
			
					r_x, _:= strconv.Atoi(parts[1])
					r_y, _:= strconv.Atoi(parts[2])

					data.LastRoom = s_Rooms{parts[0], r_x, r_y}	
					data.rooms = append(data.rooms, data.LastRoom)
				
				}else{ return nil, errors.New("Invalid Data Format!!") }
			}
			continue;
		}

		//isRoom
		if U.IsRoom(scanner.Text()) {
			parts := strings.Fields(scanner.Text())
			r_x, _:= strconv.Atoi(parts[1])
			r_y, _:= strconv.Atoi(parts[2])
			data.rooms = append(data.rooms, s_Rooms{parts[0], r_x, r_y})
			continue ;
		}

		//islink between rooms
		if U.IsLink(scanner.Text()) {
			parts := strings.Split(scanner.Text() , "-")
			data.links = append(data.links, s_Links{parts[0], parts[1]})
			continue ;
		}
		// if !(comment || command || room || links)
		return nil, errors.New("Invalid Data Format!!! ---> "+ scanner.Text()) 		
	}

	return &data, nil
}

func main(){

	args:= os.Args[1:]
	if (len(args) != 1){
		os.Stderr.WriteString("ERROR : Usage Invalid !\n")
		return ;
	}

	if!(checkFile(args[0])){
		os.Stderr.WriteString("ERROR : Invalid Argument [" + args[0] + "] !\n")
		return ;
	}

	file,_ := os.Open(args[0]);
	data, err := ParseFile(file)

	if err != nil{ os.Stderr.WriteString("ERROR :"+ err.Error() + "\n") ;return }

	_= data

	defer file.Close()

}