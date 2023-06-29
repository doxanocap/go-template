package main

import (
	"app/cmd"
	"app/pkg/banner"
	"runtime"
	"time"
)

func main() {
	banner.Default(map[string]interface{}{
		"now":      time.Now().Format(time.ANSIC),
		"numCPU":   runtime.NumCPU(),
		"GOOS":     runtime.GOOS,
		"GOARCH":   runtime.GOARCH,
		"Compiler": runtime.Compiler,
	})
	cmd.Run()
}
