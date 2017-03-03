package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"time"

	"nwjs-autoupdater/updater"

	"github.com/skratchdot/open-golang/open"
)

func main() {
	var bundle, instDir, appName string
	var timeout int

	time.Sleep(2000 * time.Millisecond)

	flag.StringVar(&bundle, "bundle", "", "Path to the update package")
	flag.StringVar(&instDir, "inst-dir", "", "Path to the application install dir")
	flag.StringVar(&appName, "app", "", "Application name")
	flag.IntVar(&timeout, "timeout", 0, "Delay before copying files")
	flag.Parse()

	time.Sleep(time.Duration(timeout) * time.Millisecond)

	cwd, _ := os.Getwd()
	logfile, err := os.Create(filepath.Join(cwd, "updater.log"))
	if err != nil {
		panic(err)
	}
	defer logfile.Close()

	logger := log.New(logfile, "", log.LstdFlags)

	logger.Println(instDir)

	var appExec string
	err, appExec = updater.Update(bundle, instDir, appName)
	if err != nil {
		logger.Fatal(err)
	}
	open.Start(appExec)
}
