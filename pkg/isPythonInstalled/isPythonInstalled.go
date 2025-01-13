package isPythonInstalled

import (
	"os/exec"
	"runtime"
	"strings"
	logToFile "update-python-2/pkg/logtofile"
)

func CheckVersion() {
	logToFile.Log.Println("Checking Python Version..")
	checkVersionCmd := exec.Command("python", "--version")
	checkVersionOutput, err := checkVersionCmd.CombinedOutput()

	if err != nil {
		logToFile.Log.Printf("Unable to Retrieve Python Version Or Python is Not Installed: %v\n", err)
		return
	}

	trim := strings.TrimSpace(string(checkVersionOutput))
	logToFile.Log.Println("Python Version: " + trim)
}

func FindPythonInstalls() ([]string, error) {
	logToFile.Log.Println("Checking Python Locations..")
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("where", "python")
	} else {
		cmd = exec.Command("which", "-a", "python3", "python")
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		logToFile.Log.Printf("No Installations of Python were Found: [ %s ]", err)
		return nil, err
	}

	paths := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(paths) > 0 {
		for _, path := range paths {
			logToFile.Log.Println(path)
		}

		return paths, nil
	}
	return nil, err
}
