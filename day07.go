package advent_of_code

import (
	"fmt"
	"log"
	"sync"
)

type programD07 struct {
	*program
	in  <-chan int
	out chan<- int
}

func (p *programD07) nextInstruction() bool {
	if p == nil || p.code == nil {
		return false
	}
	opCode, modes := getOpCodeWithParamModes(p.code[p.cursor])
	switch opCode {
	case 99:
		//close(p.out)
		return false
	case 1:
		p.addInstruction(modes)
	case 2:
		p.multiplyInstruction(modes)
	case 3:
		p.readInputInstruction()
	case 4:
		p.outputInstruction(modes)
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

func (p *programD07) readInputInstruction() {
	//read an input and store it at the position given as the next value
	in := readInputForSave(p.in)
	p.code[p.code[p.cursor+1]] = in
	p.cursor += 2
}

func (p *programD07) outputInstruction(modes []int) {
	//prints a given value given as the first param
	var output int
	if modes[0] == 0 {
		output = p.code[p.code[p.cursor+1]]
	} else {
		output = p.code[p.cursor+1]
	}
	printOutput(output, p.cursor)
	p.out <- output
	p.cursor += 2
}

type amplifier struct {
	code []int
}

func (a *amplifier) amplify(phases []int, initialPhase int) int {
	var output int
	var wg sync.WaitGroup
	for i, phase := range phases {
		code := make([]int, len(a.code))
		copy(code, a.code)
		inCh := make(chan int)
		outCh := make(chan int)
		pr := programD07{in: inCh, out: outCh, program: &program{cursor: 0, code: code}}
		wg.Add(2)
		go func(i int) {
			defer wg.Done()
			inCh <- phase
			if i == 0 {
				inCh <- initialPhase
			} else {
				inCh <- output
			}
		}(i)
		go func(i int) {
			defer wg.Done()
			output = <-outCh
			//if !ok {
			//	fmt.Println("Program halted")
			//	return
			//}
			if len(phases) == i+1 {
				fmt.Printf("Final output: %v\n", output)
			} else {
				fmt.Printf("Intermediate output: %v\n", output)

			}

		}(i)
		for pr.nextInstruction() {

		}
	}
	wg.Wait()
	return output
}

func (a *amplifier) findMaxThrusterSignal() int {
	outputCh := make(chan int)
	var maxOutput int

	//TODO should be an easier way to do this
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == i {
				continue
			}
			for k := 0; k < 5; k++ {
				if k == i || k == j {
					continue
				}
				for l := 0; l < 5; l++ {
					if l == i || l == j || l == k {
						continue
					}
					for m := 0; m < 5; m++ {
						if m == i || m == j || m == k || m == l {
							continue
						}
						phases := []int{i, j, k, l, m}
						fmt.Printf("Phases: %v\n", phases)
						go func() {
							outputCh <- a.amplify(phases, 0)
						}()
						out := <-outputCh
						if out > maxOutput {
							maxOutput = out
						}
					}
				}
			}
		}
	}

	return maxOutput
}

func (a *amplifier) findMaxThrusterSignalPart2() int {
	outputCh := make(chan int)
	var maxOutput int

	//TODO should be an easier way to do this
	for i := 5; i < 10; i++ {
		for j := 5; j < 10; j++ {
			if j == i {
				continue
			}
			for k := 5; k < 10; k++ {
				if k == i || k == j {
					continue
				}
				for l := 5; l < 10; l++ {
					if l == i || l == j || l == k {
						continue
					}
					for m := 5; m < 10; m++ {
						if m == i || m == j || m == k || m == l {
							continue
						}
						phases := []int{i, j, k, l, m}
						fmt.Printf("Phases: %v\n", phases)

						go func() {
							outputCh <- a.amplify(phases, 0)
						}()

						out := <-outputCh
						if out > maxOutput {
							maxOutput = out
						}
						fmt.Printf("Out: %v\n", out)

					}
				}
			}
		}
	}

	return maxOutput
}
