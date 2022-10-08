package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoLeandro := ContaCorrente{titular: "Leandro",
		numeroAgencia: 589, numeroConta: 65432, saldo: 125.5}

	contaDaAna := ContaCorrente{"Ana Paula", 589, 85246, 200.00}

	fmt.Println(contaDoLeandro)
	fmt.Println(contaDaAna)
}
