package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type GateType int

const (
	OR  GateType = iota
	AND GateType = iota
	XOR GateType = iota
)

type Gate struct {
	gateType GateType
	ins      [2]Wire
	out      Wire
}

func (this *Gate) Update(updatedWire Wire) Wire {
	for i := range this.ins {
		if this.ins[i].name == updatedWire.name {
			this.ins[i].value = updatedWire.value
		}
	}
	if this.ins[0].value >= 0 && this.ins[1].value >= 0 {
		switch this.gateType {
		case OR:
			this.out.value = this.ins[0].value | this.ins[1].value
		case AND:
			this.out.value = this.ins[0].value & this.ins[1].value
		case XOR:
			this.out.value = this.ins[0].value ^ this.ins[1].value
		}
	}

	return this.out
}

type Wire struct {
	name  string
	value int
}

func day24() {
	initialValues, gates := parseInput()

	knownWires := make([]Wire, len(initialValues))
	copy(knownWires, initialValues)
	out := map[string]int{}
	for len(knownWires) > 0 {
		wire := knownWires[len(knownWires)-1]
		knownWires = knownWires[:len(knownWires)-1]

		wireWasInput := false
		for i := range gates {
			for _, inWire := range gates[i].ins {
				if wire.name == inWire.name {
					wireWasInput = true
					newWire := gates[i].Update(wire)
					if newWire.value >= 0 {
						knownWires = append(knownWires, newWire)
					}
				}
			}
		}

		if !wireWasInput {
			out[wire.name] = wire.value
		}
	}

	wireNames := []string{}
	for name := range out {
		wireNames = append(wireNames, name)
	}
	sort.Slice(wireNames, func(i, j int) bool {
		return wireNames[i] < wireNames[j]
	})

	outValue := 0
	for i, name := range wireNames {
		outValue += out[name] << i
	}

	println("Part1:", outValue)

	// Found through "findIncorrectGateOutputs"
	swapGates(gates, "vmv", "z07")
	swapGates(gates, "kfm", "z20")
	swapGates(gates, "hnv", "z28")
	swapGates(gates, "hth", "tqr")

	findIncorrectGateOutputs(gates)

	swappedGates := []string{"vmv", "z07", "kfm", "z20", "hnv", "z28", "hth", "tqr"}
	sort.Slice(swappedGates, func(i, j int) bool {
		return swappedGates[i] < swappedGates[j]
	})

	println("Part2:", strings.Join(swappedGates, ","))
}

func parseInput() ([]Wire, []Gate) {
	f, err := os.Open("day24.input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	initialValues := []Wire{}
	for scanner.Scan() && len(scanner.Bytes()) > 0 {
		splitLine := strings.Split(scanner.Text(), ": ")
		value, _ := strconv.Atoi(splitLine[1])
		initialValues = append(initialValues, Wire{splitLine[0], value})
	}

	gates := []Gate{}
	for scanner.Scan() {
		inOutSplit := strings.Split(scanner.Text(), " -> ")
		out := Wire{inOutSplit[1], -1}
		insAndGateType := strings.Split(inOutSplit[0], " ")
		var gateType GateType
		switch insAndGateType[1] {
		case "OR":
			gateType = OR
		case "AND":
			gateType = AND
		case "XOR":
			gateType = XOR
		default:
			log.Fatal("Unexpected gate string", insAndGateType[1])
		}
		ins := [2]Wire{{insAndGateType[0], -1}, {insAndGateType[2], -1}}
		gates = append(gates, Gate{gateType, ins, out})
	}

	return initialValues, gates
}

func findIncorrectGateOutputs(gates []Gate) {
	unlabeledGates := make([]Gate, len(gates))
	copy(unlabeledGates, gates)
	for len(unlabeledGates) > 0 {
		// fmt.Printf("unlabeledGates: %v\n", unlabeledGates)
		renames := map[string]string{}
		stillUnlabeledGates := []Gate{}
		for i := range unlabeledGates {
			if isLabeledWire(unlabeledGates[i].ins[0]) && isLabeledWire(unlabeledGates[i].ins[1]) {
				oldName, newName := unlabeledGates[i].relabel()
				println("Relabeled gate from", oldName, "to", newName)
				renames[oldName] = newName
			} else {
				stillUnlabeledGates = append(stillUnlabeledGates, unlabeledGates[i])
			}
		}

		for i := range stillUnlabeledGates {
			stillUnlabeledGates[i].updateInputs(renames)
		}

		unlabeledGates = stillUnlabeledGates
	}
}

func (this *Gate) relabel() (string, string) {
	oldLabel := this.out.name
	splitIns := map[byte]int{}
	for _, in := range this.ins {
		index, _ := strconv.Atoi(in.name[1:])
		category := in.name[0]
		splitIns[category] = index
	}

	if len(splitIns) < 2 {
		log.Fatal("Invalid gate", this.out.name)
	}

	xIndex, hasX := splitIns['x']
	yIndex, hasY := splitIns['y']
	aIndex, hasA := splitIns['a']
	bIndex, hasB := splitIns['b']
	cIndex, hasC := splitIns['c']
	dIndex, hasD := splitIns['d']

	if hasX && hasY && xIndex == yIndex && this.gateType == XOR {
		if xIndex == 0 {
			if this.out.name != "z00" {
				log.Fatal("Invalid gate ", this)
			}
		} else {
			this.out.name = fmt.Sprintf("a%02d", splitIns['x'])
		}
	} else if hasX && hasY && xIndex == yIndex && this.gateType == AND {
		if xIndex == 0 {
			this.out.name = fmt.Sprintf("c%02d", xIndex)
		} else {
			this.out.name = fmt.Sprintf("b%02d", xIndex)
		}
	} else if hasA && hasC && aIndex == cIndex+1 && this.gateType == XOR {
		if this.out.name != fmt.Sprintf("z%02d", aIndex) {
			log.Fatal("Invalid gate ", this)
		}
	} else if hasB && hasD && bIndex == dIndex && this.gateType == OR {
		this.out.name = fmt.Sprintf("c%02d", bIndex)
	} else if hasA && hasC && aIndex == cIndex+1 && this.gateType == AND {
		this.out.name = fmt.Sprintf("d%02d", aIndex)
	} else {
		log.Fatal("Invalid gate ", this)
	}

	return oldLabel, this.out.name
}

func (this *Gate) updateInputs(renames map[string]string) {
	for i := range this.ins {
		newName, isPresent := renames[this.ins[i].name]
		if isPresent {
			this.ins[i].name = newName
		}
	}
}

func isLabeledWire(wire Wire) bool {
	_, isValidNumber := strconv.Atoi(wire.name[1:])
	return isValidNumber == nil
}

func swapGates(gates []Gate, label1 string, label2 string) {
	for i := range gates {
		if gates[i].out.name == label1 {
			gates[i].out.name = label2
		} else if gates[i].out.name == label2 {
			gates[i].out.name = label1
		}
	}
}
