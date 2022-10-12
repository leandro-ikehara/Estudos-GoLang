# ORIENTAÇÃO A OBJETOS EM GO

## Go OO - Aula 1 - "Minha primeira struct"
> Criamos nossa primeira struct chamada ContaCorrente;
> Em seguida criamos um tipo com base na struct Conta corrente sem atribuir valor para os campos e descobrimos o Zero-initialization ou Inicialização zero;
> Para finalizar, vimos duas formas utilizar a struct criada: informando ou não os nomes dos campos.

## Go OO - Aula 2 - "Referência, ponteiro e método
> Criamos um nova conta corrente utilizando a palavra new;
> Em seguida, comparamos os tipos criados comparando suas referências e entendemos o que são ponteiros na prática;
> Para finalizar, desenvolvemos o método sacar que verifica se o valor do saque é maior do que zero e se a conta possui saldo.

## Go OO - Aula 3 - "Retornos, pacotes e visibilidade"
> Criamos um método chamado depositar, que quando invocado nos retorna uma mensagem (string) e o valor do novo saldo (float64);
> Em seguida, criamos o método transferir, que tira um valor de uma conta e transfere para uma conta destino reutilizando o método depositar;
> Para finalizar, criamos um pacote chamado contas e criamos um arquivo chamado contaCorrente.go para manter todo código referente a este tipo de conta.

## Go OO - Aula 4 - "Composição e encapsulamento"
> Criamos um novo pacote chamado clientes e um arquivo chamado clientes.go, onde desenvolvemos uma nova struct chamada Titular com nome, CPF e profissão;
> Em seguida, alteramos o campo titular da struct contaCorrente para ser do tipo da struct Titular;
> Para finalizar, alteramos a visibilidade do campo saldo e criamos um método chamado obterSaldo.

# Go OO - Aula 5 - "Interface e novo tipo de conta"
> - Criamos um novo tipo de conta: a ContaPoupança;
> - Para finalizar, criamos um novo tipo interface onde podemos utilizar tanto a conta corrente como poupança para pagar um boleto através da função Sacar.

![](/go_alura_logo.png)



