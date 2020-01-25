package advent_of_code

import (
	"log"
)

//next operation which returns false if program should be halted
func (p *program) nextInstruction(inputCh chan int) bool {
	if p == nil || p.code == nil {
		return false
	}
	opCode, modes := getOpCodeWithParamModes(p.code[p.cursor])
	switch opCode {
	case 99:
		//halt program
		return false
	case 1:
		p.addInstruction(modes)
	case 2:
		p.multiplyInstruction(modes)
	case 3:
		p.readInputInstruction(inputCh)
	case 4:
		p.printInstruction(modes)
	case 5:
		p.jumpIfTrueInstruction(modes)
	case 6:
		p.jumpIfFalseInstruction(modes)
	case 7:
		p.lessThanInstruction(modes)
	case 8:
		p.equalsInstruction(modes)
	default:
		log.Printf("Invalid opCode: %v, cursor: %v\n", opCode, p.cursor)
		return false
	}
	return true
}

func (p *program) equalsInstruction(modes []int) {
	//equals
	if modes[2] == 1 {
		log.Printf("Invalid mode for instruction write : %v, cursor: %v\n", p.code[p.cursor], p.cursor)
	}
	var firstPart, secondPart int
	if modes[0] == 0 {
		firstPart = p.code[p.code[p.cursor+1]]
	} else {
		firstPart = p.code[p.cursor+1]
	}

	if modes[1] == 0 {
		secondPart = p.code[p.code[p.cursor+2]]
	} else {
		secondPart = p.code[p.cursor+2]
	}

	if firstPart == secondPart {
		p.code[p.code[p.cursor+3]] = 1
	} else {
		p.code[p.code[p.cursor+3]] = 0
	}
	p.cursor += 4
}

func (p *program) lessThanInstruction(modes []int) {
	//less-than
	if modes[2] == 1 {
		log.Printf("Invalid mode for instruction write : %v, cursor: %v\n", p.code[p.cursor], p.cursor)
	}
	var firstPart, secondPart int
	if modes[0] == 0 {
		firstPart = p.code[p.code[p.cursor+1]]
	} else {
		firstPart = p.code[p.cursor+1]
	}

	if modes[1] == 0 {
		secondPart = p.code[p.code[p.cursor+2]]
	} else {
		secondPart = p.code[p.cursor+2]
	}

	if firstPart < secondPart {
		p.code[p.code[p.cursor+3]] = 1
	} else {
		p.code[p.code[p.cursor+3]] = 0
	}
	p.cursor += 4
}

func (p *program) jumpIfFalseInstruction(modes []int) {
	//jump-if-false
	var param int
	if modes[0] == 0 {
		param = p.code[p.code[p.cursor+1]]
	} else {
		param = p.code[p.cursor+1]
	}
	if param == 0 {
		if modes[1] == 0 {
			p.cursor = p.code[p.code[p.cursor+2]]
		} else {
			p.cursor = p.code[p.cursor+2]
		}
	} else {
		//just increment cursor
		p.cursor += 3
	}
}

func (p *program) jumpIfTrueInstruction(modes []int) {
	// jump-if-true
	var param int
	if modes[0] == 0 {
		param = p.code[p.code[p.cursor+1]]
	} else {
		param = p.code[p.cursor+1]
	}
	if param != 0 {
		if modes[1] == 0 {
			p.cursor = p.code[p.code[p.cursor+2]]
		} else {
			p.cursor = p.code[p.cursor+2]
		}
	} else {
		//just increment cursor
		p.cursor += 3
	}
}

func (p *program) printInstruction(modes []int) {
	//prints a given value given as the first param
	var output int
	if modes[0] == 0 {
		output = p.code[p.code[p.cursor+1]]
	} else {
		output = p.code[p.cursor+1]
	}
	printOutput(output, p.cursor)
	p.cursor += 2
}

func (p *program) readInputInstruction(inputCh chan int) {
	//read an input and store it at the position given as the next value
	in := readInputForSave(inputCh)
	p.code[p.code[p.cursor+1]] = in
	p.cursor += 2
}

func (p *program) multiplyInstruction(modes []int) {
	//multiple the next 2 values and overwrite the value on the position given as the next value
	//ignore safety checks
	if modes[2] == 1 {
		log.Printf("Invalid mode for instruction write : %v, cursor: %v\n", p.code[p.cursor], p.cursor)
	}
	var firstPart, secondPart int
	if modes[0] == 0 {
		firstPart = p.code[p.code[p.cursor+1]]
	} else {
		firstPart = p.code[p.cursor+1]
	}

	if modes[1] == 0 {
		secondPart = p.code[p.code[p.cursor+2]]
	} else {
		secondPart = p.code[p.cursor+2]
	}
	product := firstPart * secondPart
	p.code[p.code[p.cursor+3]] = product
	p.cursor += 4
}

func (p *program) addInstruction(modes []int) {
	//add the next 2 values and overwrite the value on the position given as the next value
	//ignore safety checks
	if modes[2] == 1 {
		log.Printf("Invalid mode for instruction write : %v, cursor: %v\n", p.code[p.cursor], p.cursor)
	}
	var firstPart, secondPart int
	if modes[0] == 0 {
		firstPart = p.code[p.code[p.cursor+1]]
	} else {
		firstPart = p.code[p.cursor+1]
	}

	if modes[1] == 0 {
		secondPart = p.code[p.code[p.cursor+2]]
	} else {
		secondPart = p.code[p.cursor+2]
	}
	sum := firstPart + secondPart
	p.code[p.code[p.cursor+3]] = sum
	p.cursor += 4
}

func getOpCodeWithParamModes(input int) (int, []int) {
	opCode := input % 100
	digits := getDigits(input / 100)
	if len(digits) != 3 {
		digits = append(digits, 0, 0, 0)
	}
	return opCode, digits[:3]
}

func printOutput(output, cursor int) {
	log.Printf("output: %v, cursor: %v\n", output, cursor)
}

func readInputForSave(inputCh chan int) int {
	log.Println("Expecting input...")
	return <-inputCh
}

func (p *program) RunCodeOfDay05(inputCh chan int) []int {
	if p == nil || p.code == nil {
		return nil
	}
	//p.code[1] = noun
	//p.code[2] = verb
	for p.nextInstruction(inputCh) {
	}

	return p.code
}
