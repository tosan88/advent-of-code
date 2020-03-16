package day05

import (
	"github.com/tosan88/advent-of-code/day02"
	"github.com/tosan88/advent-of-code/day04"
	"log"
)
type Program day02.Program

//next operation which returns false if program should be halted
func (p *Program) nextInstruction(inputCh chan int) bool {
	if p == nil || p.Code == nil {
		return false
	}
	opCode, modes := GetOpCodeWithParamModes(p.Code[p.Cursor])
	switch opCode {
	case 99:
		//halt program
		return false
	case 1:
		p.AddInstruction(modes)
	case 2:
		p.MultiplyInstruction(modes)
	case 3:
		p.ReadInputInstruction(inputCh)
	case 4:
		p.PrintInstruction(modes)
	case 5:
		p.JumpIfTrueInstruction(modes)
	case 6:
		p.JumpIfFalseInstruction(modes)
	case 7:
		p.LessThanInstruction(modes)
	case 8:
		p.EqualsInstruction(modes)
	default:
		log.Printf("Invalid opCode: %v, cursor: %v\n", opCode, p.Cursor)
		return false
	}
	return true
}

func (p *Program) EqualsInstruction(modes []int) {
	//equals
	if modes[2] == 1 {
		log.Printf("Invalid mode for instruction write : %v, cursor: %v\n", p.Code[p.Cursor], p.Cursor)
	}
	var firstPart, secondPart int
	if modes[0] == 0 {
		firstPart = p.Code[p.Code[p.Cursor+1]]
	} else {
		firstPart = p.Code[p.Cursor+1]
	}

	if modes[1] == 0 {
		secondPart = p.Code[p.Code[p.Cursor+2]]
	} else {
		secondPart = p.Code[p.Cursor+2]
	}

	if firstPart == secondPart {
		p.Code[p.Code[p.Cursor+3]] = 1
	} else {
		p.Code[p.Code[p.Cursor+3]] = 0
	}
	p.Cursor += 4
}

func (p *Program) LessThanInstruction(modes []int) {
	//less-than
	if modes[2] == 1 {
		log.Printf("Invalid mode for instruction write : %v, cursor: %v\n", p.Code[p.Cursor], p.Cursor)
	}
	var firstPart, secondPart int
	if modes[0] == 0 {
		firstPart = p.Code[p.Code[p.Cursor+1]]
	} else {
		firstPart = p.Code[p.Cursor+1]
	}

	if modes[1] == 0 {
		secondPart = p.Code[p.Code[p.Cursor+2]]
	} else {
		secondPart = p.Code[p.Cursor+2]
	}

	if firstPart < secondPart {
		p.Code[p.Code[p.Cursor+3]] = 1
	} else {
		p.Code[p.Code[p.Cursor+3]] = 0
	}
	p.Cursor += 4
}

func (p *Program) JumpIfFalseInstruction(modes []int) {
	//jump-if-false
	var param int
	if modes[0] == 0 {
		param = p.Code[p.Code[p.Cursor+1]]
	} else {
		param = p.Code[p.Cursor+1]
	}
	if param == 0 {
		if modes[1] == 0 {
			p.Cursor = p.Code[p.Code[p.Cursor+2]]
		} else {
			p.Cursor = p.Code[p.Cursor+2]
		}
	} else {
		//just increment cursor
		p.Cursor += 3
	}
}

func (p *Program) JumpIfTrueInstruction(modes []int) {
	// jump-if-true
	var param int
	if modes[0] == 0 {
		param = p.Code[p.Code[p.Cursor+1]]
	} else {
		param = p.Code[p.Cursor+1]
	}
	if param != 0 {
		if modes[1] == 0 {
			p.Cursor = p.Code[p.Code[p.Cursor+2]]
		} else {
			p.Cursor = p.Code[p.Cursor+2]
		}
	} else {
		//just increment cursor
		p.Cursor += 3
	}
}

func (p *Program) PrintInstruction(modes []int) {
	//prints a given value given as the first param
	var output int
	if modes[0] == 0 {
		output = p.Code[p.Code[p.Cursor+1]]
	} else {
		output = p.Code[p.Cursor+1]
	}
	p.PrintOutput(output)
	p.Cursor += 2
}

func (p *Program) ReadInputInstruction(inputCh chan int) {
	//read an input and store it at the position given as the next value
	in := ReadInputForSave(inputCh)
	p.Code[p.Code[p.Cursor+1]] = in
	p.Cursor += 2
}

func (p *Program) MultiplyInstruction(modes []int) {
	//multiple the next 2 values and overwrite the value on the position given as the next value
	//ignore safety checks
	if modes[2] == 1 {
		log.Printf("Invalid mode for instruction write : %v, cursor: %v\n", p.Code[p.Cursor], p.Cursor)
	}
	var firstPart, secondPart int
	if modes[0] == 0 {
		firstPart = p.Code[p.Code[p.Cursor+1]]
	} else {
		firstPart = p.Code[p.Cursor+1]
	}

	if modes[1] == 0 {
		secondPart = p.Code[p.Code[p.Cursor+2]]
	} else {
		secondPart = p.Code[p.Cursor+2]
	}
	product := firstPart * secondPart
	p.Code[p.Code[p.Cursor+3]] = product
	p.Cursor += 4
}

func (p *Program) AddInstruction(modes []int) {
	//add the next 2 values and overwrite the value on the position given as the next value
	//ignore safety checks
	if modes[2] == 1 {
		log.Printf("Invalid mode for instruction write : %v, cursor: %v\n", p.Code[p.Cursor], p.Cursor)
	}
	var firstPart, secondPart int
	if modes[0] == 0 {
		firstPart = p.Code[p.Code[p.Cursor+1]]
	} else {
		firstPart = p.Code[p.Cursor+1]
	}

	if modes[1] == 0 {
		secondPart = p.Code[p.Code[p.Cursor+2]]
	} else {
		secondPart = p.Code[p.Cursor+2]
	}
	sum := firstPart + secondPart
	p.Code[p.Code[p.Cursor+3]] = sum
	p.Cursor += 4
}

func GetOpCodeWithParamModes(input int) (int, []int) {
	opCode := input % 100
	digits := day04.GetDigits(input / 100)
	if len(digits) != 3 {
		digits = append(digits, 0, 0, 0)
	}
	return opCode, digits[:3]
}

func (p Program) PrintOutput(output int) {
	log.Printf("%v has output: %v, cursor: %v\n", p.Name, output, p.Cursor)
}

func ReadInputForSave(inputCh <-chan int) int {
	log.Println("Expecting input...")
	return <-inputCh
}

func (p *Program) RunCodeOfDay05(inputCh chan int) []int {
	if p == nil || p.Code == nil {
		return nil
	}
	//p.Code[1] = noun
	//p.Code[2] = verb
	for p.nextInstruction(inputCh) {
	}

	return p.Code
}
