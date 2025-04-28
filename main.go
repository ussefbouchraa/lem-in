package main

import (
	"bufio"
	"os"
	"strconv"

	// F "fmt"
	S "strings"
)


type s_Global struct{
	Ants      	  int;
	FirstRoomName string;
	EndRoomName   string;
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

func isWhiteSpaces(line string) bool{

if line == "" || line == "\t" || line == "\r" {
	return true
}
return false;
}

func main(){
args:= os.Args[1:]
if (len(args) != 1){
	os.Stderr.WriteString("Err : Usage Invalid !\n")
	return ;
}

if!(checkFile(args[0])){
	os.Stderr.WriteString("Err : Invalid Argument !\n")
	return ;
}

file,_ := os.Open(args[0]);
scanner := bufio.NewScanner(file)


if !scanner.Scan() {
    if scanner.Err() == nil {
        os.Stderr.WriteString("Err: Empty File!\n")
		return;
	}
}

for (scanner.Scan()){

	if 	isWhiteSpaces(scanner.Text()){
		os.Stderr.WriteString("Err : Invalid Format!! \n"); return
	}

	if scanner.Text() == "##start" || scanner.Text() == "##end"{
		continue;
	}

	if S.Count(scanner.Text(), "#") == 1 && scanner.Text()[:1] == "#"{
		//ignore comment
		continue;
	}

	if S.Contains(scanner.Text(),"-") && S.Count(scanner.Text(), "-") == 1 {
		tokens:= S.Split(scanner.Text(), "-")
		if len(tokens) != 2 { os.Stderr.WriteString("Err : Invalid Format! \n"); return}
		
		tkn1, err := strconv.Atoi(tokens[0])
		tkn2, err := strconv.Atoi(tokens[1])
		if err != nil || tkn1 < 0 || tkn2 < 0 { os.Stderr.WriteString("Err : Invalid Format! \n"); return}

		_= tkn1 ; _= tkn2	
		// 1- save the links to the stucts
		}

		// 2- If room ➔ save name and coordinates{

		// }


		// 3- else {
		// 	os.Stderr.WriteString("Err: " + scanner.Err().Error() + "\n")
		// }

	}





}