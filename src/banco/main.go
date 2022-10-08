package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func main() {
	contaDaCintia := ContaCorrente{}
	contaDaCintia.titular = "Silvia"
	contaDaCintia.saldo = 549.90

	fmt.Println(contaDaCintia.saldo)

	fmt.Println(contaDaCintia.Sacar(409.9))
	fmt.Println(contaDaCintia.saldo)

	// contaDoLeandro := ContaCorrente{titular: "Leandro",
	//
	//	numeroAgencia: 589, numeroConta: 65432, saldo: 125.5}
	//
	// contaDaAna := ContaCorrente{"Ana Paula", 589, 85246, 200.00}
	//
	// fmt.Println(contaDoLeandro)
	// fmt.Println(contaDaAna)
}
