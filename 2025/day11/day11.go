package day11

import (
	"fmt"
	"strings"
)

type Device struct {
	name string
	outputs []Device
}

type ServerRack struct {
	devices []Device
}


func PartOne(lines []string, extras ...any) any {
	var rack ServerRack

	for _, line := range lines {
		strSplit := strings.Split(line, ":")

		deviceName := strSplit[0]
		outputNames := strings.Split(strings.Trim(strSplit[1], " "), " ")

		outputList := make([]Device, len(outputNames))
		for i := range outputNames {
			outputList[i] = Device{
				name: outputNames[i],
				outputs: nil,
			}
		}
		
		rack.devices = append(rack.devices, Device{
			name: deviceName,
			outputs: outputList,
		})
	}

	paths := 0
	paths = searchOutFromDevice(rack, "you", paths)

	return paths
}

func searchOutFromDevice(rack ServerRack, deviceName string, paths int) int {
	fmt.Println("> Searching:", deviceName)
	
	currDevice := Device{}
	for _, device := range rack.devices {
		if deviceName == device.name {
			currDevice = device
		}
	}

	for _, output := range currDevice.outputs {
		if output.name == "out" {
			paths += 1
			fmt.Println(paths)
		} else {
			paths += searchOutFromDevice(rack, output.name, paths)
		}
	}

	return paths
}

func PartTwo(lines []string, extras ...any) any {
	_ = lines

	return 0
}
