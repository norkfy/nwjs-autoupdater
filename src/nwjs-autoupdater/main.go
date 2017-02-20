package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"nwjs-autoupdater/updater"

	"github.com/skratchdot/open-golang/open"
)

func main() {
	var bundle, instDir, appName string

	fmt.Println("WTF!!111")

	flag.StringVar(&bundle, "bundle", "", "Path to the update package")
	flag.StringVar(&instDir, "inst-dir", "", "Path to the application install dir")
	flag.StringVar(&appName, "app", "", "Application name")
	flag.Parse()

	cwd, _ := os.Getwd()
	logfile, err := os.Create(filepath.Join(cwd, "updater.log"))
	if err != nil {
		panic(err)
	}
	defer logfile.Close()

	logger := log.New(logfile, "", log.LstdFlags)

	var appExec string
	err, appExec = updater.Update(bundle, instDir, appName)
	if err != nil {
		logger.Fatal(err)
	}
	open.Start(appExec)
}
