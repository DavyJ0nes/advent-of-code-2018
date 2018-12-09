# --- Day 1: Chronal Calibration ---

"We've detected some temporal anomalies," one of Santa's Elves at the Temporal Anomaly Research and Detection Instrument Station tells you. She sounded pretty worried when she called you down here. "At 500-year intervals into the past, someone has been changing Santa's history!"

"The good news is that the changes won't propagate to our time stream for another 25 days, and we have a device" - she attaches something to your wrist - "that will let you fix the changes with no such propagation delay. It's configured to send you 500 years further into the past every few days; that was the best we could do on such short notice."

"The bad news is that we are detecting roughly fifty anomalies throughout time; the device will indicate fixed anomalies with stars. The other bad news is that we only have one device and you're the best person for the job! Good lu--" She taps a button on the device and you suddenly feel like you're falling. To save Christmas, you need to get all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

After feeling like you've been falling for a few minutes, you look at the device's tiny screen. "Error: Device must be calibrated before first use. Frequency drift detected. Cannot maintain destination lock." Below the message, the device shows a sequence of changes in frequency (your puzzle input). A value like +6 means the current frequency increases by 6; a value like -3 means the current frequency decreases by 3.

For example, if the device displays frequency changes of +1, -2, +3, +1, then starting from a frequency of zero, the following changes would occur:

```
Current frequency  0, change of +1; resulting frequency  1.
Current frequency  1, change of -2; resulting frequency -1.
Current frequency -1, change of +3; resulting frequency  2.
Current frequency  2, change of +1; resulting frequency  3.
```

In this example, the resulting frequency is 3.

Here are other example situations:

```
+1, +1, +1 results in  3
+1, +1, -2 results in  0
-1, -2, -3 results in -6
```

Starting with a frequency of zero, what is the resulting frequency after all of the changes in frequency have been applied?

## Something interesting

So I was looking at other ways of implementing the solution and was inspired by Francesc Campoys solution that used
a seeker from a reader to read the file line by line. I implemented this in same way but split it into another function
so that it could be tested along with my initial implementation.

I was then interested in seeing how the performance compared between the two implementations and I was surprised to see
that the Reader implementation was able to run 33 times less iterations compared with the original implementation.
```
⇒ GOCACHE=off go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/davyj0nes/advent-of-code/day1
BenchmarkCalcTotal-4                    100000000               25.3 ns/op
BenchmarkCalcTotalWithReader-4           3000000               526 ns/op
```

This led me to start profiling the two implementations in more depth and from the CPU perspective `CalcTotalWithReader` 
is around 4 times faster. Below is the output from pprof

```
ROUTINE ======================== github.com/davyj0nes/advent-of-code/day1.CalcTotal in /Users/davy.jones/go/src/github.com/davyj0nes/advent-of-code/day1/main.go
     270ms      980ms (flat, cum) 21.88% of Total
... OUTPUT CUT

ROUTINE ======================== github.com/davyj0nes/advent-of-code/day1.CalcTotalWithReader in /Users/davy.jones/go/src/github.com/davyj0nes/advent-of-code/day1/main.go
      20ms      210ms (flat, cum)  4.69% of Total
```

So let's look at the benchmarks from memory perspective.

