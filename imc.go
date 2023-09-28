package main

import (
	"fmt"
	"math"
)

func main() {
	var peso, altura float64

	// Solicitar ao usuário que insira o peso e a altura
	fmt.Print("Digite o peso em quilogramas: ")
	fmt.Scan(&peso)

	fmt.Print("Digite a altura em metros: ")
	fmt.Scan(&altura)

	// Calcular o IMC
	imc := peso / (altura * altura)

	// Arredondar o IMC para duas casas decimais
	imc = math.Round(imc*100) / 100

	// Exibir o resultado
	fmt.Printf("Seu IMC é: %.2f\n", imc)

	// Interpretar o IMC
	interpretarIMC(imc)
}

func interpretarIMC(imc float64) {
	if imc < 16.00 {
		fmt.Println("Magreza grave")
	} else if imc >= 16.00 && imc < 16.99 {
		fmt.Println("Magreza moderada")
	} else if imc >= 17.00 && imc < 18.49 {
		fmt.Println("Magreza leve")
	} else if imc >= 18.50 && imc < 24.99 {
		fmt.Println("Normal")
	} else if imc >= 25.00 && imc < 29.99 {
		fmt.Println("Sobrepeso")
	} else if imc >= 30.00 && imc < 34.99 {
		fmt.Println("Obesidade Grau I")
	} else if imc >= 35.00 && imc < 39.99 {
		fmt.Println("Obesidade Grau II (severa)")
	} else {
		fmt.Println("Obesidade Grau III (mórbida)")
	}
}
