package main

import (
	"fmt"

	"github.com/leandro-ikehara/Estudos-GoLang/banco/contas"
)

func main() {
	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(100)

	fmt.Println(contaExemplo.ObterSaldo())
}
