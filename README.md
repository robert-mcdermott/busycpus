# BusyCPUs

A tool for generating heavy CPU loads for testing and analysis purposes.

## Download

- **Windows x64:** [busycpus.exe](https://github.com/robert-mcdermott/busycpus/blob/master/bin/busycpus.exe)
- **Linux   x64:** [busycpus](https://github.com/robert-mcdermott/busycpus/blob/master/bin/busycpus)

## Usage

By default, BusyCPU will use a thread per logical CPU on the system and 2 concurent workers per thread (should saturate the CPU of any system) but you can tune the thread count (-t), number of concurent workers per thread (-c) and the stride of each worker (-s) if needed.
Once it's started it will run until stopped (ctrl-c).

```
./busycpus [-t <cpu threads> -c <concurrency> -s <stride>]

Example: ./busycpus -t 4 -t 8  (four threads with 8 workers per thread)

  -c int
  	Optional: set number of concurrent number of workers per thread; defaults to 2 per 
        logical CPU 
  
  -t int
    	Optional: set number of CPU threads to use; defaults to the number of logical 
        CPUs in your system (default 4)

  -s int
    	Optional: Specify the stride (numer or random sines to compute per interation) of each
        worker mount of time to wait in milliseconds; defaults to 100 

  -h	print usage information
```
