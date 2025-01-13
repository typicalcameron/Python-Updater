package installPython

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	logToFile "update-python-2/pkg/logtofile"
)

func CleanExeInput(exeFileInput string) string {
	exeFile := strings.TrimSpace(exeFileInput)

	if len(exeFile) > 1 && ((exeFile[0] == '"' && exeFile[len(exeFile)-1] == '"') || (exeFile[0] == '\'' && exeFile[len(exeFile)-1] == '\'')) {
		exeFile := exeFile[1 : len(exeFile)-1]
		return exeFile
	} else {
		fmt.Println(exeFile)
		return exeFile
	}

}

func InstallPython() {
	for {
		fmt.Print("Please Enter Python Installer Location:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		exeFileInput := scanner.Text()
		cleanExeInput := CleanExeInput(exeFileInput)

		fileExist, err := os.Stat(cleanExeInput)
		if err != nil {
			fmt.Println(err)
		} else {
			logToFile.Log.Printf("The file exists: [ %s ]", fileExist.Name())
			var installInput string
			fmt.Print("Are You Sure You Want to Install the EXE?: Y | N: ")
			fmt.Scanln(&installInput)
			installInput = strings.ToUpper(installInput)
			switch installInput {
			case "Y":

				logToFile.Log.Printf("Starting Install Process..")
				psScript := fmt.Sprintf(`
			 			Start-Process -FilePath "%s" -ArgumentList "/quiet InstallAllUsers=1 PrependPath=1 Include_doc=1 Include_pip=1 Include_tcltk=1 Include_test=1 AssociateFiles=1 Include_launcher=1 InstallLauncherAllUsers=1 Include_lib=1 CompileAll=1" -Wait
			 			`, cleanExeInput)
				installCommand := exec.Command("powershell", "-Command", psScript)

				output, err := installCommand.CombinedOutput()
				if err != nil {
					fmt.Println("Error", err)
					fmt.Println("Command Output\n:", string(output))
					return
				} else {
					fmt.Println("Python was successfully installed!")
					return
				}
			case "N":
				logToFile.Log.Println("Cancelling Install..")
				return
			default:
				fmt.Println("Invalid Input. Please enter Y or N.")
			}
		}
	}
	// 		fmt.Print("Are You Sure You Want to Install the Executable?: Y | N: ")
	// 		fmt.Scanln(&installInput)
	// 		installInput = strings.ToUpper(installInput)
	// 		switch installInput {
	// 		case "Y":
	// 			logToFile.Log.Printf("Install File: %s", exeRemoveQuotes)
	// 			psScript := fmt.Sprintf(`
	// 			Start-Process -FilePath "%s" -ArgumentList "/quiet InstallAllUsers=1 PrependPath=1 Include_doc=1 Include_pip=1 Include_tcltk=1 Include_test=1 Include_launcher=1 InstallLauncherAllUsers=1 Include_lib=1 CompileAll=1"
	// 			`, exeRemoveQuotes)
	// 			installCommand := exec.Command("powershell", "-Command", psScript)
	//
	// 			output, err := installCommand.CombinedOutput()
	// 			if err != nil {
	// 				fmt.Println("Error", err)
	// 				fmt.Println("Command Output\n:", string(output))
	// 				return
	// 			} else {
	// 				fmt.Println("Python was successfully installed!")
	// 				return
	// 			}
	// 		case "N":
	// 			logToFile.Log.Println("Cancelling Install..")
	// 			return
	// 		default:
	// 			fmt.Println("Invalid Input. Please enter Y or N.")
	// 		}
	// 	}
	// }
	// fileExist, err := os.Stat(exeFileInput)
	// trimExeInput, err := strconv.Unquote(exeFileInput)
	// if err != nil {
	// 	fmt.Printf("Error Trimming File Input: [ %s ]", err)
	// 	return
	// } else {
	// 	if _, err := os.Stat(trimExeInput); err == nil {
	// 		fmt.Printf("File exists\n")
	// 	} else {
	// 		fmt.Printf("File does not exist\n")
	// 	}
	// }
	//
	// // fileExist, err := os.Stat(exeRemoveQuotes)
	// if os.IsNotExist(err) {
	// 	fmt.Println("File Does Not Exist. Try again.")
	// } else {
	// 	logToFile.Log.Println(fileExist.Name())
	// 	for {
	// 		var installInput string
	// 		fmt.Print("Are You Sure You Want to Install the Executable?: Y | N: ")
	// 		fmt.Scanln(&installInput)
	// 		installInput = strings.ToUpper(installInput)
	// 		switch installInput {
	// 		case "Y":
	// 			logToFile.Log.Printf("Install File: %s", exeRemoveQuotes)
	// 			psScript := fmt.Sprintf(`
	// 			Start-Process -FilePath "%s" -ArgumentList "/quiet InstallAllUsers=1 PrependPath=1 Include_doc=1 Include_pip=1 Include_tcltk=1 Include_test=1 Include_launcher=1 InstallLauncherAllUsers=1 Include_lib=1 CompileAll=1"
	// 			`, exeRemoveQuotes)
	// 			installCommand := exec.Command("powershell", "-Command", psScript)
	//
	// 			output, err := installCommand.CombinedOutput()
	// 			if err != nil {
	// 				fmt.Println("Error", err)
	// 				fmt.Println("Command Output\n:", string(output))
	// 				return
	// 			} else {
	// 				fmt.Println("Python was successfully installed!")
	// 				return
	// 			}
	// 		case "N":
	// 			logToFile.Log.Println("Cancelling Install..")
	// 			return
	// 		default:
	// 			fmt.Println("Invalid Input. Please enter Y or N.")
	// 		}
	// 	}
	// }
}
