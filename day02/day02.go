package day02

import (
	"log"
)

type Program struct {
	Code   []int
	Cursor int
	Name   string
}

func NewProgram(code []int) *Program {
	return &Program{Code: code, Cursor: 0}
}

//next operation which returns false if program should be halted
func (p *Program) next() bool {
	if p == nil || p.Code == nil {
		return false
	}
	opCode := p.Code[p.Cursor]
	switch opCode {
	case 99:
		//halt program
		return false
	case 1:
		//add the next 2 values and overwrite the value on the position given as the next value
		//ignore safety checks
		sum := p.Code[p.Code[p.Cursor+1]] + p.Code[p.Code[p.Cursor+2]]
		p.Code[p.Code[p.Cursor+3]] = sum
		p.Cursor += 4
	case 2:
		//multiple the next 2 values and overwrite the value on the position given as the next value
		//ignore safety checks
		product := p.Code[p.Code[p.Cursor+1]] * p.Code[p.Code[p.Cursor+2]]
		p.Code[p.Code[p.Cursor+3]] = product
		p.Cursor += 4
	default:
		log.Printf("Invalid opCode: %v\n", opCode)
		return false
	}
	return true
}

func (p *Program) RunCode(noun, verb int) []int {
	if p == nil || p.Code == nil {
		return nil
	}
	p.Code[1] = noun
	p.Code[2] = verb
	for p.next() {
	}

	return p.Code
}
