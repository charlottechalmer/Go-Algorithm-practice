package main

import (
	"fmt"
	"strconv"
	"time"
)

var frameID = 0
var frameName = ""
var assemblyArrangement [3]string

func main() {
	assemblyArrangement[0] = "frame"
	assemblyArrangement[1] = "body"
	assemblyArrangement[2] = "interior"

	numberOfCars := 3

	frameToCreate := len(assemblyArrangement)
	frameInfoChan := make(chan string)

	for i := 1; i < numberOfCars+1; i++ {
		for stageNumber := 0; stageNumber < frameToCreate; stageNumber++ {
			go assemblyStage(frameInfoChan, assemblyArrangement[stageNumber], stageNumber, frameToCreate, i)
			time.Sleep(time.Millisecond * 1000)
		}
	}
	// for i := 0; i < framesToCreate; i++ {
	// 	go assembleFrame(frameInfoChan)
	// 	go addBody(frameInfoChan)
	// 	go addInterior(frameInfoChan)
	// 	time.Sleep(time.Millisecond * 3000)
	// }
}

func assemblyStage(frameInfoChan chan string, stage string, stageNumber int, frameToCreate int, carID int) {
	nextStage := "paint"
	if stageNumber < frameToCreate {
		frameName = "FrameID" + strconv.Itoa(carID)
		if stageNumber != frameToCreate-1 {
			nextStage = assemblyArrangement[stageNumber+1]
		}
	}
	fmt.Println("add", stage, "to", frameName, "and proceed to", nextStage)
	frameInfoChan <- frameName
	time.Sleep(time.Millisecond * 1000)
}

func assembleFrame(frameInfoChan chan string) {
	frameID++
	frameName = "Frame ID" + strconv.Itoa(frameID)
	fmt.Println("Frame assembly complete", frameName, "Proceed to body")
	frameInfoChan <- frameName
	time.Sleep(time.Millisecond * 5)
}

func addBody(frameInfoChan chan string) {
	body := frameInfoChan
	fmt.Println("Add body to", body, "and proceed to interior.")
	frameInfoChan <- frameName
	time.Sleep(time.Millisecond * 5)
}

func addInterior(frameInfoChan chan string) {
	interior := frameInfoChan
	fmt.Println("Add interior to", interior, "and proceed to paint.")
	frameInfoChan <- frameName
	time.Sleep(time.Millisecond * 5)
}
