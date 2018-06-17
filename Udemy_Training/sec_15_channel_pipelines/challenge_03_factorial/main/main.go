package main

import (
    "flag"
    "os"
    "log"
    "runtime/pprof"
    "runtime"
    "github.com/KaoruDev/learning_go/Udemy_Training/sec_15_channel_pipelines/challenge_03_factorial/limited_routines"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
   flag.Parse()

   if *cpuprofile != "" {
       f, err := os.Create(*cpuprofile)
       if err != nil {
           log.Fatal(err)
       }
       pprof.StartCPUProfile(f)
       defer pprof.StopCPUProfile()
   }

    workCount := 100000
    // with_array.Run(workCount)
    // without_array.Run(workCount)
    limited_routines.Run(workCount)

    if *memprofile != "" {
        f, err := os.Create(*memprofile)
        if err != nil {
            log.Fatal("could not create memory profile: ", err)
        }
        runtime.GC() // get up-to-date statistics
        if err := pprof.WriteHeapProfile(f); err != nil {
            log.Fatal("could not write memory profile: ", err)
        }
        f.Close()
    }
}

