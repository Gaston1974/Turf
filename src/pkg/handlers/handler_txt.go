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
	aux1 := 0
	var nCarrera int64
	nCarrera = 1
	var nCarreras int64

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

		if strings.Contains(line, "Processing") {

			end := len(line)
			nCarreras, _ = strconv.ParseInt(line[27:end-1], 10, 64)

		}

		if err := scannerP.Err(); err != nil {
			fmt.Println("Error reading file:", err)
		}
	}

	// ************************

	scanner := bufio.NewScanner(f)

	var li int
	var index int
	var l string

	var car dao.Carrera
	var cab dao.Caballo
	var carDet dao.CarreraDet

	var carVec []dao.Carrera
	var carDetVec []dao.CarreraDet

	for scanner.Scan() {
		line := scanner.Text()
		li++

		fmt.Printf("\n linea: %d", li)

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

			l = line[index+14:]

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

			cab.Nombre = strings.Trim(l[:index], " ")
			carDet.Nombre = strings.Trim(l[:index], " ")
			carDet.Handicap = strings.Trim(l[index:index+2], " ")

			carDetVec = append(carDetVec, carDet)

			l2 := strings.Trim(l[index+2:], " ")

			carDet.Jockey = strings.Trim(l2[:30], " ")
			carDet.Padre = strings.Trim(l2[30:80], " ")
			carDet.Cuidador = strings.Trim(l2[80:], " ")

			//fmt.Printf("\n jockey: %s \n padre: %s \n cuidador: %s", carDet.Jockey, carDet.Padre, carDet.Cuidador)

			flag2 = 0
			flag3 = 0
			flag4 = 1
		}

		if flag2 == 1 {
			flag3 = 1
		}

		if strings.Contains(line, "Caballeriza") || flag4 == 1 {

			flag2++
			flag4 = 0

			if nCarrera == nCarreras {

				car.Detalle = carDetVec
				carVec = append(carVec, car)
			}

		}

		// car.Detalle = carDetVec

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
		}

	}

	// for _, v := range carVec {

	// 	fmt.Printf("\n nombre: %s \nfecha: %s \ndistancia: %d ", v.Nombre, v.Fecha, v.Distancia)

	// 	for _, r := range v.Detalle {

	// 		fmt.Printf("\n\n")
	// 		fmt.Printf("\n nombre: %s \nhandicap: %s ", r.Nombre, r.Handicap)

	// 	}
	// }

	return 1, "parsed file created successfully."

	// ********************************************************

}
