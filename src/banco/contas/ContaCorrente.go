package contas

import "github.com/leandro-ikehara/Estudos-GoLang/banco/clientes"

type ContaCorrente struct {
	Titular clientes.Titular
	Agencia int
	Conta   int
	saldo   float64
}

func (c *ContaCorrente) Sacar(saque float64) string {
	podeSacar := saque > 0 && saque <= c.saldo
	if podeSacar {
		c.saldo -= saque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(deposito float64) (string, float64) {
	if deposito > 0 {
		c.saldo += deposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do deposito menor que zero", c.saldo
	}
}

func (c *ContaCorrente) Tranferir(transferencia float64, contaDestino *ContaCorrente) bool {
	if transferencia < c.saldo && transferencia > 0 {
		c.saldo -= transferencia
		contaDestino.Depositar(transferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
