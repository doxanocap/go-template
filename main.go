package main

import (
	"app/cmd"
	"app/pkg/banner"
	"runtime"
	"time"
)

//	@title			go-template
//	@version		1.0
//	@description	Swagger документация для ТЗ.
//	@termsOfService	http://swagger.io/terms/

//	@host		localhost:8080
//	@BasePath	/

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
