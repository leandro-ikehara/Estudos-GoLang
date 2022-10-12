package contas

import "github.com/leandro-ikehara/Estudos-GoLang/banco/clientes"

type ContaPoupanca struct {
	Titular                  clientes.Titular
	Agencia, Conta, Operacao int
	saldo                    float64
}

func (c *ContaPoupanca) Sacar(saque float64) string {
	podeSacar := saque > 0 && saque <= c.saldo
	if podeSacar {
		c.saldo -= saque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaPoupanca) Depositar(deposito float64) (string, float64) {
	if deposito > 0 {
		c.saldo += deposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do deposito menor que zero", c.saldo
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
