package funciones

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
)

type Puntuaciones struct {
	Nombre   string
	Apellido string
	Puntos   float64
	Fecha    string
}

// type Valores struct {
// 	Ano      string
// 	Fecha    string
// 	Concepto string
// 	Credito  float64
// 	Debito   float64
// 	Saldo    float64
// 	Ncuenta  string
// 	Razon    string
// 	Cuit     string
// 	TipoCta  string
// }

func LeerArchivo(ruta string, w http.ResponseWriter) {

	f, err := os.Open(ruta)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// read 1024 bytes at a time
	buf := make([]byte, 1024)

	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			// there is no more data to read
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 {
			//f.Write(buf[:n])
			//fmt.Println("leyendo ", buf[:n])
			w.Write(buf[:n])
		}
	}

}

func WriteJson(payload interface{}) (int, string) {

	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to Marshal JSON file: %v ", payload)
		return 0, ""
	}

	return 1, string(dat)
}

func Wait(c int) {

	i := 0
	for range time.Tick(time.Second) {
		if i == c {
			break
		}
		i++
	}

}

func PadLeft(str, pad string, lenght int) string {
	for {
		str = pad + str
		if len(str) > lenght {
			return str[1 : lenght+1]
		}
	}

}

func PadRight(str, pad string, lenght int) string {

	for len(str) < lenght {

		str = str + pad

	}

	return str
}

// Archivo de logs

func Log(msg string, path string) {

	var file *os.File
	var err error

	// FECHA ACTUAL
	//currentDate := time.Now().Format(time.DateOnly)

	if _, findErr := os.Stat(path); os.IsNotExist(findErr) {

		// el archivo existe ?

		file, err = os.Create(path)

		if err != nil {
			fmt.Printf("%s", "\n error al crear el archivo")

			return
		}
	} else {

		file, _ = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	}

	defer file.Close()

	if msg != "" {
		//fmt.Printf("   [%s]", msg)
		_, err = file.WriteString(msg + "\n")
	}

	if err != nil {
		m := "error al escribir el archivo"
		fmt.Printf("%s err: %s", m, err)

	}

}

func OrdenaVector(vec []string) {

	var aux string

	for j := 0; j < len(vec); j++ {

		for i := 0; i < len(vec)-j-1; i++ {

			if vec[i] > vec[i+1] {

				aux = vec[i]
				vec[i] = vec[i+1]
				vec[i+1] = aux

			}
		}

	}

}

func NumberFinder(line string, num int) int {

	line = line[num:]

	//line = strings.ReplaceAll(line, " ", "a")
	f := strings.Split(line, "")

	// fmt.Printf("\n vector de linea: %s", f)
	g := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	var index int
	index = -1

	for j, v := range f {
		// fmt.Printf("\n j: %d", j)
		for _, k := range g {
			// fmt.Printf("\n v: %s k: %s", v, k)
			if v == k {

				index = j
				// fmt.Printf("\n index: %d", index)
				break
			}

		}

		if index != -1 {
			break
		}

	}

	return index

}

func DownloadFile(w http.ResponseWriter, filePath string, msg string) {
	// Open the file
	f, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Unable to open file for download", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	// Get the filename
	filename := filepath.Base(filePath)

	// w.Header().Set("Content-Type", "application/octet-stream")
	// Set headers to make the browser download the file
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", filename))
	//w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", getFileSize(f)))

	//w.Write([]byte(msg))

	// Reset file read pointer to beginning
	f.Seek(0, io.SeekStart)

	// Copy file to response
	io.Copy(w, f)
}

func getFileSize(f *os.File) int64 {
	info, err := f.Stat()
	if err != nil {
		return 0
	}
	return info.Size()
}

func AddStyle(i int) *xlsx.Style {

	var style *xlsx.Style

	switch i {

	case 1:

		// Create a style
		style = xlsx.NewStyle()
		// Font formatting
		font := xlsx.DefaultFont()
		//font.Bold = true
		font.Size = 8
		font.Color = "FF000000" // black text
		style.Font = *font
		style.ApplyFont = true
		// Fill (background color)
		fill := xlsx.Fill{
			PatternType: "solid",
			FgColor:     "FF00FF00", // white background
			BgColor:     "FF00FF00",
		}
		style.Fill = fill
		style.ApplyFill = true
		// Alignment
		style.Alignment = xlsx.Alignment{
			Horizontal: "left",
			Vertical:   "bottom",
		}
		style.ApplyAlignment = true

	case 2:

		// Create a style
		style = xlsx.NewStyle()
		// Font formatting
		font := xlsx.DefaultFont()
		//font.Bold = true
		font.Size = 9
		font.Color = "FFFFFF" // white text
		style.Font = *font
		style.ApplyFont = true
		// Fill (background color)
		fill := xlsx.Fill{
			PatternType: "solid",
			FgColor:     "FF808080", // Grey background
			BgColor:     "FF808080",
		}
		style.Fill = fill
		style.ApplyFill = true
		// Alignment
		style.Alignment = xlsx.Alignment{
			Horizontal: "left",
			Vertical:   "bottom",
		}
		style.ApplyAlignment = true

	default:

		// Create a style
		style = xlsx.NewStyle()
		// Font formatting
		font := xlsx.DefaultFont()
		//font.Bold = true
		font.Size = 9
		font.Color = "FF000000" // black text
		style.Font = *font
		style.ApplyFont = true
		// Fill (background color)
		fill := xlsx.Fill{
			PatternType: "solid",
			FgColor:     "FFFFA500", // orange background
			BgColor:     "FFFFA500",
		}
		style.Fill = fill
		style.ApplyFill = true
		// Alignment
		style.Alignment = xlsx.Alignment{
			Horizontal: "left",
			Vertical:   "bottom",
		}
		style.ApplyAlignment = true

	}

	return style

}

// func AddSheet(i *xlsx.Sheet) *xlsx.Row {

// }
