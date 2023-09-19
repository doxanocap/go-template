package banner

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

const goTemplateBanner = `
 ________  ________                _________  _______   _____ ______   ________  ___       ________  _________  _______      
|\   ____\|\   __  \              |\___   ___\\  ___ \ |\   _ \  _   \|\   __  \|\  \     |\   __  \|\___   ___\\  ___ \     
\ \  \___|\ \  \|\  \  ___________\|___ \  \_\ \   __/|\ \  \\\__\ \  \ \  \|\  \ \  \    \ \  \|\  \|___ \  \_\ \   __/|    
 \ \  \  __\ \  \\\  \|\____________\  \ \  \ \ \  \_|/_\ \  \\|__| \  \ \   ____\ \  \    \ \   __  \   \ \  \ \ \  \_|/__  
  \ \  \|\  \ \  \\\  \|____________|   \ \  \ \ \  \_|\ \ \  \    \ \  \ \  \___|\ \  \____\ \  \ \  \   \ \  \ \ \  \_|\ \ 
   \ \_______\ \_______\                 \ \__\ \ \_______\ \__\    \ \__\ \__\    \ \_______\ \__\ \__\   \ \__\ \ \_______\
    \|_______|\|_______|                  \|__|  \|_______|\|__|     \|__|\|__|     \|_______|\|__|\|__|    \|__|  \|_______|

`

func Default() {
	data := map[string]interface{}{
		"now":      time.Now().Format(time.ANSIC),
		"numCPU":   runtime.NumCPU(),
		"GOOS":     runtime.GOOS,
		"GOARCH":   runtime.GOARCH,
		"Compiler": runtime.Compiler,
	}
	err := Show(os.Stdout, strings.NewReader(goTemplateBanner), data)
	if err != nil {
		log.Println(err)
	}
}

func Show(out io.Writer, in io.Reader, data map[string]interface{}) error {
	if in == nil {
		return fmt.Errorf("the input is nil")
	}

	banner, err := io.ReadAll(in)
	if err != nil {
		return fmt.Errorf("error trying to read the banner, err: %v", err)
	}

	t, err := template.New("banner").Parse(string(banner))
	if err != nil {
		return fmt.Errorf("error trying to parse the banner file, err: %v", err)
	}

	err = t.Execute(out, data)
	if err != nil {
		return fmt.Errorf("error trying to execute template: %v", err)
	}

	return nil
}
