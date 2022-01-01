// Day 1: Sonar Sweep
package main

import (
	"bufio"
	"fmt"

	// "io/ioutil"
	"log"
	"os"
	"strconv"
)

func CheckFile(name string) bool {
	path, err := os.Getwd() // To get working dirctory
	if err != nil {
		fmt.Println("Error in getting working directory ", err)
		return false
	} else {
		fmt.Printf("File: %v \n file_type: %T \n", path, path)

		fileinfo, err := os.Stat(path + "/" + name) //check if file is present
		if err == nil {
			// fmt.Println(fileinfo)
			println(fileinfo.Name()) // gets filename
			fmt.Println("Sonar Data found")
			return true
		} else {
			fmt.Println("Error | File Not Found ", err)
			return false
		}
	}
}

func main() {
	fmt.Println("Calculate increase in sea depth from sonar data")
	fmt.Println("1. Check if sonar data exists")
	name := "sonardata.txt"
	if CheckFile(name) {
		fmt.Println("Read File Data")

		fileStream, err := os.Open(name) // open a file
		if err != nil {
			log.Fatal(err) // log error and exit
		}
		defer fileStream.Close()                 // use defer to close file after end of tasks
		scanData := bufio.NewScanner(fileStream) // reads from fileStream
		workData := make([]int, 0)               // makes an empty slice of int

		for scanData.Scan() { // scans line by line till EOF
			if strToInt, err := strconv.Atoi(scanData.Text()); err == nil { // convert string to Integer
				workData = append(workData, strToInt) // append in workData as int
			}
		}
		defer waterLevel(workData) // check water level change one after another reading

		// addThree := make(map[int]int)
		addThree := make([]int, 0)
		for i, j, k := 0, 1, 2; k < len(workData); i, j, k = i+1, j+1, k+1 { // iterate on i and j and k
			addThree = append(addThree, (workData[i] + workData[j] + workData[k])) // add and append
		}
		// fmt.Println(addThree, len(addThree))
		waterLevel(addThree) // check water level change one after another reading

	}

}

func waterLevel(workData []int) {
	status := make(map[int]string) // create a empty dictonary
	// println(status)
	// fmt.Printf("%v\n", status)
	// fmt.Printf("Len %d\n", len(workData))
	for i, j := 0, 1; j < len(workData); i, j = i+1, j+1 { // iterate on i and j
		// println(i, j)
		if workData[j] > workData[i] {
			// fmt.Println(workData[j], workData[i], "Increased")
			status[i] = "Increased" // append in status
		}
	}
	// fmt.Printf("%v", status)
	fmt.Printf("Total range : %d\nIncreament in range : %d\n", len(workData), len(status))
}