package advent_of_code

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

var debug = func() bool {
	env := os.Getenv("DEBUG")
	if env == "" {
		log.Println("Using default debug flag (false)")
		return false
	}
	debug, err := strconv.ParseBool(env)
	if err != nil {
		log.Println("Using default debug flag (false)")
		return false
	}
	log.Println("Using debug=" + env)
	return debug
}()

type programD07 struct {
	*program
	in    <-chan int
	out   chan<- int
	close chan int
}

func (p *programD07) nextInstruction() bool {
	if p == nil || p.code == nil {
		return false
	}
	opCode, modes := getOpCodeWithParamModes(p.code[p.cursor])
	switch opCode {
	case 99:
		if p.close != nil {
			close(p.close)
		}
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
	in := p.readInput()
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
	if debug {
		p.printOutput(output)
	}
	p.out <- output
	p.cursor += 2
}

func (p *programD07) readInput() int {
	if debug {
		log.Printf("%v expecting input...\n", p.name)
	}
	input := <-p.in
	if debug {
		log.Printf("%v got input: %v\n", p.name, input)
	}
	return input
}

type amplifier struct {
	code []int
}

func (a *amplifier) amplify(phases []int, initialPhase int) int {
	var output int
	var wg sync.WaitGroup
	semaphore := make(chan int, len(phases))
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
				out := <-semaphore
				inCh <- out
			}
		}(i)
		go func(i int) {
			defer wg.Done()
			output = <-outCh

			if len(phases) == i+1 {
				fmt.Printf("Final output: %v\n", output)
			} else {
				fmt.Printf("Intermediate output: %v\n", output)
			}
			semaphore <- output
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
						if debug {
							fmt.Printf("Phases: %v\n", phases)
						}
						go func() {
							outputCh <- a.amplifyWithLoop(phases)
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

func (a *amplifier) amplifyWithLoop(phases []int) int {
	output := make([]int, len(phases))
	var wg sync.WaitGroup
	var mux sync.Mutex
	var inChannels []chan int
	var outChannels []chan int
	var closeChannels []chan int
	for range phases {
		inChannels = append(inChannels, make(chan int, 2)) //make buffered, so it wont' block when the output is received
		outChannels = append(outChannels, make(chan int))
		closeChannels = append(closeChannels, make(chan int))
	}

	for i, phase := range phases {
		code := make([]int, len(a.code))
		copy(code, a.code)
		outCh := outChannels[i]
		inCh := inChannels[i]
		closeCh := closeChannels[i]
		pr := programD07{in: inCh, out: outCh, program: &program{cursor: 0, code: code, name: fmt.Sprintf("P%v", i)}, close: closeCh}
		wg.Add(1)
		go func(p *programD07, i int) {
			//defer wg.Done()
			if debug {
				fmt.Printf("P%v : %p ; in : %p, out : %p\n", i, p, p.in, p.out)
			}
			for p.nextInstruction() {

			}
		}(&pr, i)
		go func(i, phase int, inCh chan int) {
			defer wg.Done()

			inCh <- phase
			if i == 0 {
				inCh <- 0
			}
		}(i, phase, inCh)
	}
	wg.Wait()
	for i, phase := range phases {
		mux.Lock()
		code := make([]int, len(a.code))
		copy(code, a.code)
		outCh := outChannels[i]
		inCh := inChannels[i]
		closeCh := closeChannels[i]
		var nextInch chan int
		if len(phases) == i+1 {
			nextInch = inChannels[0]
		} else {
			nextInch = inChannels[i+1]
		}
		mux.Unlock()
		wg.Add(1)
		go func(i, phase int, inCh, nextInch, outCh, closeCh chan int) {
			defer wg.Done()

		forLoop:
			for {
				select {
				case out := <-outCh:
					mux.Lock()
					output[i] = out
					fmt.Printf("Intermediate output for %v: %v\n", i, output[i])
					nextInch <- out
					mux.Unlock()
				case <-closeCh:
					fmt.Printf("Closed %v\n", i)
					mux.Lock()
					fmt.Printf("Final output for %v: %v\n", i, output[i])
					mux.Unlock()
					break forLoop
				}

			}

		}(i, phase, inCh, nextInch, outCh, closeCh)
	}

	wg.Wait()
	return output[len(output)-1]
}
