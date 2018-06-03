package profiling

import (
	// "bytes"
	"fmt"
	"github.com/autovelop/playthos"
	"log"
	"os"
	"os/exec"
	"runtime/pprof"
)

var profCPU, profMem bool

// StartProfiling starts either or both CPU and memory profiling. Will output results to terminal or you can manually run pprof on the .pprof file where the test is located
func StartProfiling(c bool, m bool) {
	profCPU = c
	profMem = m
	if profCPU {
		cpuBuffer, err := os.Create("cpu.pprof")
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}

		if err := pprof.StartCPUProfile(cpuBuffer); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
	}
}

// ReportUPS outputs to terminal the average updates per second the engine is reporting
func ReportUPS(e *engine.Engine) {
	fmt.Printf("> Profiling: %v average updates per second\n", e.UPS())
}

// StopProfiling stops whichever profiling is running
func StopProfiling() {
	var (
		cmdOut []byte
		cmdErr error
	)
	if profCPU {
		pprof.StopCPUProfile()

		if cmdOut, cmdErr = exec.Command("go", "tool", "pprof", "--text", "top", "cpu.pprof").Output(); cmdErr != nil {
			fmt.Fprintln(os.Stderr, cmdErr)
			os.Exit(1)
		}
		fmt.Println(string(cmdOut))
	}

	if profMem {
		memBuffer, err := os.Create("mem.pprof")
		if err != nil {
			log.Fatal("could not create Mem profile: ", err)
		}
		pprof.WriteHeapProfile(memBuffer)
		memBuffer.Close()

		if cmdOut, cmdErr = exec.Command("go", "tool", "pprof", "--text", "mem.pprof").Output(); err != nil {
			fmt.Fprintln(os.Stderr, cmdErr)
			os.Exit(1)
		}
		fmt.Println(string(cmdOut))
	}
}
