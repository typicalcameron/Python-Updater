package main

import (
	"fmt"
	checkEOL "python-updater/pkg/check-eol"
	"python-updater/pkg/installPython"
	"python-updater/pkg/isPythonInstalled"
	"python-updater/pkg/uninstallPython"
)

func main() {
	tasks := []string{
		"Check Python Version",
		"Check Python File Location(s)",
		"Uninstall Python",
		"Install Python",
		"Check EOL",
	}

	for {
		fmt.Println("Please Select an Option:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task)
		}
		fmt.Println("0. Exit")

		var choice int

		fmt.Print("Enter Your Choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 0:
			fmt.Println("Exiting...")
			return
		case 1:
			isPythonInstalled.CheckVersion()
		case 2:
			isPythonInstalled.FindPythonInstalls()
		case 3:
			uninstallPython.UninstallBoth()
		case 4:
			installPython.InstallPython()
		case 5:
			checkEOL.CheckEOL()
		default:
			fmt.Println("Invalid Choice, Please Try Again.")
		}
	}
}
