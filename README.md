# BusyCPUs

A tool for generating heavy CPU loads for testing and analysis purposes.

## Download

- **Windows x64:** [busycpus.exe](https://github.com/robert-mcdermott/busycpus/blob/master/bin/busycpus.exe?raw=true)
- **Linux   x64:** [busycpus](https://github.com/robert-mcdermott/busycpus/blob/master/bin/busycpus?raw=true)

## Usage

By default, BusyCPU will use a thread per logical CPU on the system and 4 concurent workers per thread (should saturate the CPU of any system) but you can tune the thread count (-t), number of concurent workers per thread (-c) and the stride of each worker (-s) if needed.
Once it's started it will run until stopped (ctrl-c).

```
./busycpus [-t <cpu threads> -c <concurrency> -s <stride>]

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

## Examples

Default mode, it will swamp every CPU in the system:

```
$ ./busycpus 
Starting 32 concurent workers, multiplexed over 8 CPU threads...
Working: *****
```
Specify how many threads and concurent workers to use:

```
$ ./busycpus -t 2 -c 16
Starting 16 concurent workers, multiplexed over 4 CPU threads...
Working: *****
```
