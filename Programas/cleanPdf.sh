#!/bin/bash


#arch="file.pdf"

#current_datetime=$(date)

     gs -dPDFA=1 -dBATCH -dNOPAUSE -dNOOUTERSAVE -sDEVICE=txtwrite -sOutputFile=./Programas/SALIDA.txt -dPDFACompatibilityPolicy=1 ./Programas/file.pdf > output.txt
	mv -T ./output.txt ./Programas/output.txt
     echo -n "lines: " >> ./Programas/output.txt
     wc -l < "./Programas/SALIDA.txt" >> ./Programas/output.txt


echo "fin script"


	
