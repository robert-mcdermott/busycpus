// A simple utility to generate CPU load
package main

import (
	"flag"
	"fmt"
	"github.com/briandowns/spinner"
	"math"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {
	// define and set default command parameter flags
	var tFlag = flag.Int("t", runtime.NumCPU(), "Optional: set number of CPU threads to use; defaults to the number of logical CPUs in your system")
	var cFlag = flag.Int("c", runtime.NumCPU()*4, "Optional: set number of concurrent workers to use; defaults to 4 per logical CPU")
	var sFlag = flag.Int("s", 1000000, "Optional: set number of sines to compute per worker (stride); defaults to 1,000,000")
	var hFlag = flag.Bool("h", false, "print usage information")

	// usage function that's executed if a required flag is missing or user asks for help (-h)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nUsage: %s [-t <cpu threads> -c <concurrency>\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Println()
	}
	flag.Parse()

	// provide help (-h)
	if *hFlag == true {
		flag.Usage()
		os.Exit(0)
	}

	// set the number of CPU threads to use as provided by the -t flag
	runtime.GOMAXPROCS(*tFlag)
	wg := new(sync.WaitGroup)

	// about do do some work, start the timer
	start := time.Now()

	fmt.Printf("Starting %d concurent workers, multiplexed over %d CPU threads...\n", *cFlag, *tFlag)
	// Start the concurent workers
	for i := 0; i < *cFlag; i++ {
		wg.Add(1)
		go doWork(wg, *sFlag)
	}

	// start a spinner so people know it's working and not hung
	go func() {
		s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
		s.Prefix = "Working: "
		s.Start()
		s.Color("green")
	}()

	// wait for all goroutines to complete
	wg.Wait()

	// work is done, stop the timer
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

// Compute the sine of a bunch of random numbers to burn CPU
func doWork(wg *sync.WaitGroup, sFlag int) {
	defer wg.Done()

	for {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < sFlag; i++ {
			_ = math.Sin(float64(r.Intn(10000)))
		}
		runtime.Gosched()
	}
}
