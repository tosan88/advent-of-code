package day07

import (
	"fmt"
	"github.com/tosan88/advent-of-code/day05"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
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

type Program struct {
	*day05.Program
	in    <-chan int
	out   chan<- int
	close chan int
}

func (p *Program) nextInstruction() bool {
	if p == nil || p.Code == nil {
		return false
	}
	opCode, modes := day05.GetOpCodeWithParamModes(p.Code[p.Cursor])
	switch opCode {
	case 99:
		if p.close != nil {
			close(p.close)
		}
		return false
	case 1:
		p.AddInstruction(modes)
	case 2:
		p.MultiplyInstruction(modes)
	case 3:
		p.ReadInputInstruction()
	case 4:
		p.OutputInstruction(modes)
	case 5:
		p.JumpIfTrueInstruction(modes)
	case 6:
		p.JumpIfFalseInstruction(modes)
	case 7:
		p.LessThanInstruction(modes)
	case 8:
		p.EqualsInstruction(modes)
	default:
		log.Printf("Invalid opCode: %v, Cursor: %v\n", opCode, p.Cursor)
		return false
	}
	return true
}

func (p *Program) ReadInputInstruction() {
	//read an input and store it at the position given as the next value
	in := p.ReadInput()
	p.Code[p.Code[p.Cursor+1]] = in
	p.Cursor += 2
}

func (p *Program) OutputInstruction(modes []int) {
	//prints a given value given as the first param
	var output int
	if modes[0] == 0 {
		output = p.Code[p.Code[p.Cursor+1]]
	} else {
		output = p.Code[p.Cursor+1]
	}
	if debug {
		p.PrintOutput(output)
	}
	p.out <- output
	p.Cursor += 2
}

func (p *Program) ReadInput() int {
	if debug {
		log.Printf("%v expecting input...\n", p.Name)
	}
	input := <-p.in
	if debug {
		log.Printf("%v got input: %v\n", p.Name, input)
	}
	return input
}

type Amplifier struct {
	Code    []int
	Timeout time.Duration
}

func (a *Amplifier) Amplify(phases []int, initialPhase int) (int, error) {
	var output int
	var wg sync.WaitGroup
	semaphore := make(chan int, len(phases))
	errCh := make(chan error, 33)
	finishedCh := make(chan struct{})
	for i, phase := range phases {
		code := make([]int, len(a.Code))
		copy(code, a.Code)
		inCh := make(chan int)
		outCh := make(chan int)
		pr := Program{in: inCh, out: outCh, Program: &day05.Program{Cursor: 0, Code: code}}
		wg.Add(2)
		go func(i int) {
			defer wg.Done()
			select {
			case inCh <- phase:
			case <-time.After(a.Timeout):
				errCh <- fmt.Errorf("sending on input channel timed out after %v seconds", a.Timeout)
				return
			}
			if i == 0 {
				select {
				case inCh <- initialPhase:
				case <-time.After(a.Timeout):
					errCh <- fmt.Errorf("sending on input channel timed out after %v seconds", a.Timeout)
					return
				}
			} else {
				var out int
				select {
				case out = <-semaphore:
				case <-time.After(a.Timeout):
					errCh <- fmt.Errorf("receiving on semaphore channel timed out after %v seconds", a.Timeout)
					return
				}
				select {
				case inCh <- out:
				case <-time.After(a.Timeout):
					errCh <- fmt.Errorf("sending on output channel timed out after %v seconds", a.Timeout)
					return
				}
			}
		}(i)
		go func(i int) {
			defer wg.Done()
			select {
			case output = <-outCh:
			case <-time.After(a.Timeout):
				errCh <- fmt.Errorf("receiving on output channel timed out after %v seconds", a.Timeout)
				return
			}

			if debug {
				if len(phases) == i+1 {
					fmt.Printf("Final output: %v\n", output)
				} else {
					fmt.Printf("Intermediate output: %v\n", output)
				}
			}
			select {
			case semaphore <- output:
			case <-time.After(a.Timeout):
				errCh <- fmt.Errorf("sending on semaphore channel timed out after %v seconds", a.Timeout)
				return
			}
		}(i)
		for pr.nextInstruction() {

		}
	}

	go func() {
		wg.Wait()
		finishedCh <- struct{}{}
	}()
	select {
	case <-finishedCh:
		return output, nil
	case err := <-errCh:
		return 0, err
	case <-time.After(a.Timeout):
		return 0, fmt.Errorf("amplify timed out after %v seconds", a.Timeout)

	}
}

func (a *Amplifier) FindMaxThrusterSignal(amplifyWithLoop bool) (int, error) {
	outputCh := make(chan int)
	errCh := make(chan error, 33)
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
							var amplified int
							var err error
							if amplifyWithLoop {
								amplified, err = a.AmplifyWithLoop(phases)
							} else {
								amplified, err = a.Amplify(phases, 0)
							}
							if err != nil {
								errCh <- err
								return
							}
							select {
							case outputCh <- amplified:
							case <-time.After(a.Timeout):
								errCh <- fmt.Errorf("sending on output channel timed out after %v seconds", a.Timeout)
								return
							}

						}()
						select {
						case out := <-outputCh:
							if out > maxOutput {
								maxOutput = out
							}
							if debug {
								fmt.Printf("Out: %v\n", out)
							}
						case <-time.After(a.Timeout):
							return 0, fmt.Errorf("FindMaxThrusterSignal timed out after %v seconds", a.Timeout)
						}
					}
				}
			}
		}
	}

	return maxOutput, nil
}

func (a *Amplifier) AmplifyWithLoop(phases []int) (int, error) {
	output := make([]int, len(phases))
	var wg sync.WaitGroup
	var mux sync.Mutex
	var inChannels []chan int
	var outChannels []chan int
	var closeChannels []chan int
	errCh := make(chan error, 33)
	finishedCh := make(chan struct{})
	for range phases {
		inChannels = append(inChannels, make(chan int, 2)) //make buffered, so it wont' block when the output is received
		outChannels = append(outChannels, make(chan int))
		closeChannels = append(closeChannels, make(chan int))
	}

	for i, phase := range phases {
		code := make([]int, len(a.Code))
		copy(code, a.Code)
		outCh := outChannels[i]
		inCh := inChannels[i]
		closeCh := closeChannels[i]
		pr := Program{in: inCh, out: outCh, Program: &day05.Program{Cursor: 0, Code: code, Name: fmt.Sprintf("P%v", i)}, close: closeCh}
		wg.Add(1)
		go func(p *Program, i int) {
			//defer wg.Done()
			if debug {
				fmt.Printf("P%v : %p ; in : %p, out : %p\n", i, p, p.in, p.out)
			}
			for p.nextInstruction() {

			}
		}(&pr, i)
		go func(i, phase int, inCh chan int) {
			defer wg.Done()
			select {
			case inCh <- phase:
			case <-time.After(a.Timeout):
				errCh <- fmt.Errorf("sending on input channel timed out after %v seconds", a.Timeout)
				return
			}
			if i == 0 {
				select {
				case inCh <- 0:
				case <-time.After(a.Timeout):
					errCh <- fmt.Errorf("sending on input channel timed out after %v seconds", a.Timeout)
					return
				}
			}
		}(i, phase, inCh)
	}
	wg.Wait()
	for i, phase := range phases {
		mux.Lock()
		code := make([]int, len(a.Code))
		copy(code, a.Code)
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
					if debug {
						fmt.Printf("Intermediate output for %v: %v\n", i, output[i])
					}
					nextInch <- out
					mux.Unlock()
				case <-closeCh:
					mux.Lock()
					if debug {
						fmt.Printf("Closed %v\n", i)
						fmt.Printf("Final output for %v: %v\n", i, output[i])
					}
					mux.Unlock()
					break forLoop
				case <-time.After(a.Timeout):
					errCh <- fmt.Errorf("receiving on output or close channel timed out after %v seconds", a.Timeout)
					return
				}

			}

		}(i, phase, inCh, nextInch, outCh, closeCh)
	}

	go func() {
		wg.Wait()
		finishedCh <- struct{}{}
	}()
	select {
	case <-finishedCh:
		return output[len(output)-1], nil
	case err := <-errCh:
		return 0, err
	case <-time.After(a.Timeout):
		return 0, fmt.Errorf("AmplifyWithLoop timed out after %v seconds", a.Timeout)

	}
}
