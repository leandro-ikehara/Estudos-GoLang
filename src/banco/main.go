package main

import (
	"fmt"

	"github.com/leandro-ikehara/Estudos-GoLang/banco/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoLeandro := contas.ContaPoupanca{}
	contaDoLeandro.Depositar(150)
	PagarBoleto(&contaDoLeandro, 80)

	fmt.Println(contaDoLeandro.ObterSaldo())

	contaDoAnaPaula := contas.ContaCorrente{}
	contaDoAnaPaula.Depositar(700)
	PagarBoleto(&contaDoAnaPaula, 250)

	fmt.Println(contaDoAnaPaula.ObterSaldo())
}
