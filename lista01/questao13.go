package main

import "fmt"

func main () {
	var nota float32
	fmt.Print ("Digite a nota:")
	fmt.Scan (&nota)
	 if nota >= 9 { 
		fmt.Print ("A")
	 } else if nota >= 7.5 {
		fmt.Print ("B")
	  } else if nota >= 6 {
		fmt.Print ("C")
	   } else if nota < 6 {
		fmt.Print ("D")
	}