package updater

import (
	"path/filepath"
	"os"
	"github.com/mholt/archiver"
)

func Update(bundle, instDir, appName string) (error, string) {
	appExecName := appName + ".exe"
  appExec := filepath.Join(instDir, appExecName)

	err := os.RemoveAll(filepath.Join(instDir, "package.nw"));

	if err != nil {
		return err, appExec
	}

  err = archiver.Zip.Open(bundle, instDir)
	if err != nil {
		return err, appExec
	}

  return nil, appExec
}
