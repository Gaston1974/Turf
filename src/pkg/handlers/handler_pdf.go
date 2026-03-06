package handlers

import (
	//"encoding/json"

	"fmt"
	funciones "hello/src/pkg/funciones"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func HandlerPDF(w http.ResponseWriter, r *http.Request) {

	var err error

	// Path to the shell script
	scriptPath := "./Programas/cleanPdf.sh"

	absPath, err := filepath.Abs(scriptPath)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return
	}

	a := "./Programas/file.pdf"

	ab, err := filepath.Abs(a)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return
	}

	// Create a new command to run the shell script
	cmd := exec.Command("bash", absPath)

	// Set the command's output to be the standard output (terminal)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	flag := true

	for { // check for file
		_, findErr := os.Stat(ab)

		if !flag {

			break
		}

		if !os.IsNotExist(findErr) {

			flag = false
			// Run the command and check for errors

			err = cmd.Run()
			if err != nil {
				fmt.Println("Error executing script:", err.Error())
				return
			}

			fmt.Println("\nArchivo encontrado en /Programas")

		} else {

			funciones.ResponseWithJSON(w, 400, "Achivo no encontrado")
			return

		}

	}

	// *********************************************************************

	//_, find, _ := strings.Cut(path.Name(), ".")

	// ** lectura directorio Bancos
	dir, err := os.Open("./Programas")
	if err != nil {
		fmt.Println("\nerror lectura de directorio Bancos")
		return
	}

	defer dir.Close()

	// Read directory entries

	files, err := dir.Readdirnames(0) // 0 means read all files
	if err != nil {
		fmt.Println("\nerror lectura de directorio Bancos")
		return
	}

	turf := ""
	stdout := ""

	for _, v := range files {

		switch v {

		case "SALIDA.txt":
			turf = v

		case "output.txt":
			stdout = v

		}

	}

	// ********************************************

	// ** ---------------------

	// Parseo de archivos:

	res, msg := handlerTXT("./Programas/"+turf, "./Programas/"+stdout)

	fmt.Printf("\nresultado: %s", msg)

	if res == 1 {
		funciones.ResponseWithJSON(w, 200, msg)
	} else {
		funciones.ResponseWithJSON(w, 400, msg)
	}
	//funciones.Log(msg, ".log.txt")

	//os.Remove("./Bancos/" + bank)
	//os.Remove("./Convertidos/" + name)
	//os.Remove("./Bancos/" + "file.pdf")
	//os.Remove(xlsPath)

}

// *********************************************************************
