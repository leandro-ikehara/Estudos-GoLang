package main

import (
	"fmt"

	"github.com/leandro-ikehara/Estudos-GoLang/banco/contas"
)

func main() {

	contaDoLeandro := contas.ContaCorrente{Titular: "Leandro", Saldo: 125.5}
	contaDaAna := contas.ContaCorrente{Titular: "Ana Paula", Saldo: 200.00}

	status := contaDaAna.Tranferir(-100, &contaDoLeandro)

	fmt.Println(status)
	fmt.Println(contaDoLeandro)
	fmt.Println(contaDaAna)
}
