package advent_of_code

import (
	"log"
)

type program struct {
	code   []int
	cursor int
	name string
}

func NewProgram(code []int) *program {
	return &program{code: code, cursor: 0}
}

//next operation which returns false if program should be halted
func (p *program) next() bool {
	if p == nil || p.code == nil {
		return false
	}
	opCode := p.code[p.cursor]
	switch opCode {
	case 99:
		//halt program
		return false
	case 1:
		//add the next 2 values and overwrite the value on the position given as the next value
		//ignore safety checks
		sum := p.code[p.code[p.cursor+1]] + p.code[p.code[p.cursor+2]]
		p.code[p.code[p.cursor+3]] = sum
		p.cursor += 4
	case 2:
		//multiple the next 2 values and overwrite the value on the position given as the next value
		//ignore safety checks
		product := p.code[p.code[p.cursor+1]] * p.code[p.code[p.cursor+2]]
		p.code[p.code[p.cursor+3]] = product
		p.cursor += 4
	default:
		log.Printf("Invalid opCode: %v\n", opCode)
		return false
	}
	return true
}

func (p *program) RunCode(noun, verb int) []int {
	if p == nil || p.code == nil {
		return nil
	}
	p.code[1] = noun
	p.code[2] = verb
	for p.next() {
	}

	return p.code
}
