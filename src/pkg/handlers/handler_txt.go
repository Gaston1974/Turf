package handlers

import (
	"bufio"
	"fmt"
	"hello/src/pkg/dao"
	fun "hello/src/pkg/funciones"
	"strconv"
	"strings"

	"os"
)

func handlerTXT(ruta string, stdout string) (int, string) {

	flag1 := 0
	flag2 := 0
	flag3 := 0
	flag4 := 0
	flag5 := 0
	flag6 := 0
	iii := 0
	aux1 := 0
	var nCarrera int64
	nCarrera = 1
	var nCarreras int64
	var totaLines int64

	f, err := os.Open(ruta)
	if err != nil {
		fmt.Printf("/nerror: %s", err.Error())
		fun.Log("\nerror de lectura archivo intermedio", "./Programas/"+"Logs.txt")
		return 0, "error de lectura archivo intermedio"
	}

	r, err := os.Open(stdout)
	if err != nil {
		fmt.Printf("/nerror: %s", err.Error())
		fun.Log("\nerror de lectura archivo stdout", "./Programas/"+"Logs.txt")
		return 0, "error de lectura archivo stdout"
	}

	defer f.Close()
	defer r.Close()

	// ************************ obtengo cantidad de carreras

	scannerP := bufio.NewScanner(r)

	for scannerP.Scan() {
		line := scannerP.Text()

		end := len(line)

		if strings.Contains(line, "Processing") {

			nCarreras, _ = strconv.ParseInt(line[27:end-1], 10, 64)

		}

		if strings.Contains(line, "lines: ") {

			totaLines, _ = strconv.ParseInt(strings.Trim(line[7:end], " "), 10, 64)
			fmt.Printf("\n\n lines: %d ", totaLines)

		}

		if err := scannerP.Err(); err != nil {
			fmt.Println("Error reading file:", err)
		}
	}

	// ************************

	scanner := bufio.NewScanner(f)

	li := 0
	var index int
	var l string

	var car dao.Carrera
	var cab dao.Caballo
	var jock dao.Jockey
	var cui dao.Cuidador
	var carDet dao.CarreraDet

	var carVec []dao.Carrera
	var carDetVec []dao.CarreraDet

	for scanner.Scan() {
		line := scanner.Text()
		li++

		if strings.Contains(line, "Ñ") {
			strings.ReplaceAll(line, "Ñ", "N")
		}

		//fmt.Printf("\n linea: %d", li)

		if strings.Contains(line, "REUNION") {

			ind := fun.NumberFinder(line, 0) // 121

			h := line[ind+7:]
			h = strings.Trim(h, " ")
			ind = fun.NumberFinder(h, 0)
			// fmt.Printf("\n index: %d", ind)
			// fmt.Printf("\n fecha: %s", h)

			dia := h[ind+1 : ind+2]
			mes := h[6+ind : ind+9]
			dd := fun.NumberFinder(h[ind+8:], 0)
			año := h[ind+8+dd : ind+8+dd+4]

			car.Fecha = dia + "-" + mes + "-" + año
			// fmt.Printf("\n dia: %s \nmes: %s \naño: %s", dia, mes, año)

		}

		// -----------------------------------------------------------

		if flag1 == 1 || flag1 == 2 || flag1 == 3 {

			flag1++

		} else if flag1 == 4 {

			car.Descripcion = line
			flag1 = 0

		}

		if strings.Contains(line, "El Hipódromo Argentino") || flag6 == 1 {

			flag6 = 1

			if strings.Contains(line, "Comisión") {

				flag6 = 0
				flag2 = 0

			} else {
				continue
			}

		}

		if strings.Contains(line, "Carrera ") {

			nCarrera++

			if nCarrera > 2 {
				car.Detalle = carDetVec
				carVec = append(carVec, car)
				flag5 = 1

			}

			ind := strings.Index(line, "-")

			car.Nombre = strings.Trim(line[ind-31:ind-1], " ")
			car.Distancia, _ = strconv.ParseInt(strings.Trim(line[ind+1:ind+7], " "), 10, 64)

			flag1++
			flag2 = 0

		}

		if flag3 == 1 {

			if strings.Contains(line, "DEBUTA") {
				index = strings.Index(line, "D")
			} else {
				index = fun.NumberFinder(line, 0)

			}

			l = line[index+18:]

			index = fun.NumberFinder(l, 0)

			if index > 80 {
				aux1 = 1

				continue
			}

			if aux1 == 1 {
				aux1 = 0

				index = fun.NumberFinder(line, 0)
				l = line[index+2:]

			} else {
				l = l[index+2:]
			}

			index = fun.NumberFinder(l, 0)

			if flag5 == 1 {
				carDetVec = nil
				flag5 = 0
			}

			// insert de caballos y recupero id
			cab.Nombre = strings.Trim(l[:index], " ")

			id_3, m := cab.LoadDB()
			if id_3 == 0 {
				fmt.Printf("\n mesg: %s", m)
			}

			carDet.Competidor = int(id_3)
			carDet.Handicap = strings.Trim(l[index:index+2], " ")

			l2 := strings.Trim(l[index+4:], " ")

			// insert de jockeys y recupero id
			jock.Nombre = strings.Trim(l2[:30], " ")
			jock.Apellido, jock.Nombre, _ = strings.Cut(jock.Nombre, " ")
			if fun.NumberFinder(jock.Apellido, 0) != -1 {
				jock.Apellido, jock.Nombre, _ = strings.Cut(jock.Nombre, " ")
			}

			id_1, m := jock.LoadDB()
			if id_1 == 0 {
				fmt.Printf("\n mesg: %s", m)
			}

			// insert de cuidadores y recupero id
			if line[iii-7:iii-6] == " " {
				cui.Nombre = strings.Trim(line[iii-6:], " ")
				cui.Apellido, cui.Nombre, _ = strings.Cut(cui.Nombre, " ")
			}

			cab.Padre = strings.Trim(line[30:iii-6], " ")

			id_2, m := cui.LoadDB()
			if id_2 == 0 {
				fmt.Printf("\n mesg: %s", m)
			}

			// armado detalle de la carrera
			carDet.Nombre = car.Nombre
			carDet.Jockey = int(id_1)
			carDet.Padre = strings.Trim(l2[30:80], " ")
			carDet.Cuidador = int(id_2)

			// carga vector detalle
			carDetVec = append(carDetVec, carDet)

			flag2 = 0
			flag3 = 0
			flag4 = 1
		}

		if flag2 == 1 {
			flag3 = 1
		}

		if strings.Contains(line, "Caballeriza") || flag4 == 1 {

			if flag4 == 0 {
				iii = strings.Index(line, "Entrenador")
			}

			flag2++
			flag4 = 0

			if nCarrera == nCarreras && li == int(totaLines) {

				// fmt.Printf("\n nCarrera: %d \nCarreras: %d \nli: %d \ntotaLines: %d", nCarrera, nCarreras, li, int(totaLines))
				car.Detalle = carDetVec
				carVec = append(carVec, car)
			}

		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			return 0, "Error reading file"
		}
		// fmt.Printf("\n li: %d", li)
	}

	var msg string
	var cars string

	// carga de carreras en BD
	// ki := len(carVec)
	// fmt.Printf("\n\n\n carVec Len: %d \n", ki)

	for _, v := range carVec {

		// fmt.Printf("\n nombre: %s \nfecha: %s \ndistancia: %d ", v.Nombre, v.Fecha, v.Distancia)

		res, msg := v.LoadDB()

		cars = strconv.FormatInt(nCarreras, 10)

		if res != 1 {
			// fmt.Printf("\n carreras: %d", nCarreras)

			return 0, msg + "\n carreras cargadas: " + cars
		}

	}

	return 1, msg + "\n carreras cargadas: " + cars

	// ********************************************************

}
