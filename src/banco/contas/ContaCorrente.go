package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(saque float64) string {
	podeSacar := saque > 0 && saque <= c.Saldo
	if podeSacar {
		c.Saldo -= saque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(deposito float64) (string, float64) {
	if deposito > 0 {
		c.Saldo += deposito
		return "Deposito realizado com sucesso", c.Saldo
	} else {
		return "Valor do deposito menor que zero", c.Saldo
	}
}

func (c *ContaCorrente) Tranferir(transferencia float64, contaDestino *ContaCorrente) bool {
	if transferencia < c.Saldo && transferencia > 0 {
		c.Saldo -= transferencia
		contaDestino.Depositar(transferencia)
		return true
	} else {
		return false
	}
}
