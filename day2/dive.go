// Day 2: Dive!
package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	// "io/ioutil"
	"log"
	"os"
)

func CheckMovementFile(name string) bool {
	path, err := os.Getwd() // To get working dirctory
	if err != nil {
		fmt.Println("Error in getting working directory ", err)
		return false
	} else {
		// fmt.Printf("File: %v \n file_type: %T \n", path, path)

		fileinfo, err := os.Stat(path + "/" + name) //check if file is present
		if err == nil {
			// fmt.Println(fileinfo)
			println(fileinfo.Name()) // gets filename
			fmt.Println("Movement Data found")
			return true
		} else {
			fmt.Println("Error | File Not Found ", err)
			return false
		}
	}
}

func main() {
	fmt.Println("Submarine Movement")
	fmt.Println("1. Check if position data exists")
	name := "positiondata.txt"
	if CheckMovementFile(name) {
		fmt.Println("Read File Data")

		fileStream, err := os.Open(name) // open a file
		if err != nil {
			log.Fatal(err) // log error and exit
		}
		defer fileStream.Close()                 // use defer to close file after end of tasks
		scanData := bufio.NewScanner(fileStream) // reads from fileStream
		// workData := make([]string, 0)            // makes an empty slice of int
		// type Postion struct {
		// 	movementName  string
		// 	movementValue int
		// }
		// positionData := []Postion{
		// 	Postion{movementName: "forward", movementValue: 0},
		// 	Postion{movementName: "up", movementValue: 0},
		// 	Postion{movementName: "aim", movementValue: 0},
		// }
		// postitionChange := map[string]Postion{}
		// fmt.Println(positionData[0].movementName)
		// fmt.Println(positionData[0].movementValue)
		// positionData[0].movementValue = positionData[0].movementValue + 8
		// fmt.Println(postitionChange[positionData[0].movementName])

		// moveThree := make(map[string]int)
		moveThree := map[string]int{
			"forward": 0,
			"down":    0,
			"aim":     0,
		}
		// moveThree["forward"] = 0
		// moveThree["down"] = 0
		// moveThree["aim"] = 0

		fmt.Println(moveThree)
		for scanData.Scan() { // scans line by line till EOF
			splitData := strings.Split(scanData.Text(), " ") //split each line of string at space

			if strToInt, err := strconv.Atoi(splitData[1]); err == nil {
				if splitData[0] == "forward" {

					moveThree["forward"] = moveThree["forward"] + strToInt
					moveThree["down"] = moveThree["down"] + (moveThree["aim"] * strToInt)
				} else if splitData[0] == "up" {
					// moveThree["down"] = moveThree["down"] - strToInt
					moveThree["aim"] = moveThree["aim"] - strToInt
				} else if splitData[0] == "down" {
					// moveThree["down"] = moveThree["down"] + strToInt
					moveThree["aim"] = moveThree["aim"] + strToInt
				}
			}
		}
		fmt.Println(moveThree)
		fmt.Println(moveThree["forward"] * moveThree["down"])
	}
}
