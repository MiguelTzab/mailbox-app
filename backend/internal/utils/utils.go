package utils

import (
	"log"
	"os"
	"runtime/pprof"
	"strings"
)

func StartCPUProfile(filename string) func() {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	return func() {
		pprof.StopCPUProfile()
		f.Close()
	}
}

func StartMemProfile(filename string) func() {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	return func() {
		pprof.WriteHeapProfile(f)
		f.Close()
	}
}

func SplitAndTrim(s, delimiter string) []string {
	if s == "" {
		return nil
	}
	parts := strings.Split(s, delimiter)
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}
