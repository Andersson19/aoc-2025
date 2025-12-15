package day11

import (
	"strings"
)

type Device struct {
	name string
	outputs []*Device
}

type VisitedDevice struct  {
	name string
	seenFft bool
	seenDac bool
}

func createDevices(lines []string) map[string]*Device {
	devices := make(map[string]*Device)

	for _, line := range lines {
		strSplit := strings.Split(line, ": ")
		deviceName := strSplit[0]
		outputNames := strings.Split(strSplit[1], " ")

		// create and alloc memory for outputs
		devices[deviceName] = &Device{
			name: deviceName,
			outputs: make([]*Device, len(outputNames)),
		}
	}
	// manually add out
	devices["out"] = &Device{
		name: "out",
	}

	// fill with outputs
	for _, line := range lines {
		strSplit := strings.Split(line, ": ")
		deviceName := strSplit[0]
		outputList := strings.Split(strSplit[1], " ")
		
		for i, o := range outputList {
			devices[deviceName].outputs[i] = devices[o]
		}
	}
	return devices
}

func PartOne(lines []string, extras ...any) any {
	devices := createDevices(lines)

	return searchOutFromDevice(*devices["you"])
}

func searchOutFromDevice(device Device) int {
	paths := 0
	for _, output := range device.outputs {
		if output.name == "out" {
			return 1
		} else {
			paths += searchOutFromDevice(*output)
		}
	}

	return paths
}

func PartTwo(lines []string, extras ...any) any {
	devices := createDevices(lines)
	
	// 
	visited := make(map[VisitedDevice]int)

	return searchTree(*devices["svr"], false, false, visited)
}

func searchTree(device Device, hasPassedFFT, hasPassedDAC bool, visited map[VisitedDevice]int) int {
	if device.name == "out" {
		if hasPassedFFT && hasPassedDAC {
			return 1
		} else {
			return 0
		}
	}

	paths := 0
	for _, output := range device.outputs {
		seenFFT := hasPassedFFT || output.name == "fft"
		seenDAC := hasPassedDAC || output.name == "dac"
		key := VisitedDevice{
			name: output.name,
			seenFft: seenFFT,
			seenDac: seenDAC,
		}

		if v, ok := visited[key]; ok {
			paths += v
		} else {
			v = searchTree(*output, seenFFT, seenDAC, visited)
			visited[key] = v
			paths += v
		}
	}
	return paths
}