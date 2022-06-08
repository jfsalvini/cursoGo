package main

import (
	"fmt"
)

func main() {
	contaDoJorge := contas.contaCorrente{Titular: "Jorge", Saldo: 300}
	contaDoFelipe := contas.contaCorrente{Titular: "Felipe", Saldo: 100}

	status := contaDoJorge.Transferir(200, &contaDoFelipe)

	fmt.Println(status)
	fmt.Println(contaDoJorge)
	fmt.Println(contaDoFelipe)
}
