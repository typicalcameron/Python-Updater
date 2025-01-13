package uninstallPython

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"update-python-2/pkg/isPythonInstalled"
	logToFile "update-python-2/pkg/logtofile"
)

func UninstallPython() {
	psScript := `
$Products = Get-ChildItem -Path "HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\" | Where-Object { $_.GetValue('DisplayName') -like '*Python*' } | 
    ForEach-Object {
        [PSCustomObject]@{
            Name           = $_.PSChildName
            DisplayName    = $_.GetValue('DisplayName')
            UninstallString = $_.GetValue('UninstallString')
            DisplayVersion = $_.GetValue('DisplayVersion')
        }
}

$Remaining = @()
if ($Products) {
    foreach ($prod in $Products) {
	    if (($prod.DisplayName -like "*Development Libraries*")-or ($prod.Name -like "*utility scripts*") -or ($prod.DisplayName -like "*Add to Path*") -or ($prod.DisplayName -like "*Tcl/Tk Support*")) {
	        Start-Process -FilePath "msiexec.exe" -ArgumentList "/x $($prod.UninstallString.Split("/I")[2]) /quiet" -Wait -Verbose

	        $exists = Get-ChildItem -Path "HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\" | Where-Object { $_.GetValue('DisplayName') -eq $prod.DisplayName }
	        if ($exists) {
	            Write-Output "Failed to uninstall: [ $($prod.DisplayName) ]"
	        } else {
	            Write-Output "Successfully unintalled: [ $($prod.DisplayName) ]"
	        }
	    } else {
	        $Remaining += $prod
	    }
	}      
}

if ($Remaining) {
	foreach ($prod in $Remaining) {
	    Start-Process -FilePath "msiexec.exe" -ArgumentList "/x $($prod.UninstallString.Split("/I")[2]) /quiet" -Wait -Verbose

	    if ($exists) {
	        Write-Output "Failed to uninstall: [ $($prod.DisplayName) ]"
	    } else {
	        Write-Output "Successfully unintalled: [ $($prod.DisplayName) ]"
	    }
	}
}

$HKCUApp = Get-ChildItem -Path "HKCU:\SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\" | Where-Object { $_.GetValue('DisplayName') -like '*Python*' } | ForEach-Object {
    [PSCustomObject]@{
        Name           = $_.PSChildName
        DisplayName    = $_.GetValue('DisplayName')
        UninstallString = $_.GetValue('UninstallString')
        DisplayVersion = $_.GetValue('DisplayVersion')
    }
}

if ($HKCUApp.UninstallString -match '"(.*?)"') {
    $filePath = $Matches[1]
}

if ($HKCUApp) {
    Start-Process -FilePath $filePath -ArgumentList "/uninstall /quiet" -Wait -Verbose
    $exists = Get-ChildItem -Path "HKCU:\SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\" | Where-Object { $_.GetValue('DisplayName') -like '*Python*' }
    
    if ($exists) {
        Write-Output "Failed to uninstall: [ $($prod.DisplayName) ]"
    } else {
        Write-Output "Successfully unintalled: [ $($prod.DisplayName) ]"
    }
}
`
	uninstallCommand := exec.Command("powershell", "-Command", psScript)
	output, err := uninstallCommand.CombinedOutput()
	if err != nil {
		logToFile.Log.Println("Error", err)
		logToFile.Log.Println("Command Output:", string(output))
	} else {
		if string(output) != "" {
			logToFile.Log.Println(string(output))
		} else {
			logToFile.Log.Println("No Installations of Python Were Found.")
		}
	}
}

func UninstallWindowsAppPython() {
	paths, err := isPythonInstalled.FindPythonInstalls()
	if err != nil {
		logToFile.Log.Println("No Windows App Installations of Python Were Found.")
	}
	for _, path := range paths {
		if strings.Contains(path, "WindowsApps") {
			logToFile.Log.Printf("Removing Windows App Version of Python: [ %s ]", path)
			err := os.Remove(path)
			if err != nil {
				logToFile.Log.Printf("Failed to Remove File: %s", err)
			} else {
				logToFile.Log.Println("File Removed Successfully.")
			}
		}
	}
}

func UninstallBoth() {
	for {
		fmt.Print("Are You Sure You Want to Uninstall the Current Python Install(s)? Y | N: ")
		var uninstallInput string
		fmt.Scanln(&uninstallInput)
		uninstallInput = strings.ToUpper(uninstallInput)
		switch uninstallInput {
		case "Y":
			logToFile.Log.Println("Starting Uninstall Process..")
			UninstallPython()
			UninstallWindowsAppPython()
			return
		case "N":
			fmt.Println("Cancelling uninstall..")
			return
		default:
			fmt.Println("Invalid Choice, Please Try Again.")
		}
	}
}
