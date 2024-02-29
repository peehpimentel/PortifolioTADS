package main

import (
	"bufio"
	"fmt"
	"os"
)

type User struct {
	name      string
	age       string
	address   string
	cellphone string
	email     string
}

func main() {

	//createUser()
	userData()

}

/*
func createUser() {

	var newUser = [5]string{" ", " ", " ", " ", " "}
	var user = [0]newUser{}

}*/

func userData() {

	in := bufio.NewReader(os.Stdin)
	var user User

	for i := 0; i < 6; i++ {
		switch i {
		case 1:
			fmt.Print("Entre com seu nome: ")
			user.name, _ = in.ReadString('\n')
		case 2:
			fmt.Print("Entre com sua idade: ")
			user.age, _ = in.ReadString('\n')
		case 3:
			fmt.Print("Entre com seu endereço: ")
			user.address, _ = in.ReadString('\n')
		case 4:
			fmt.Print("Entre com seu número de telefone: ")
			user.cellphone, _ = in.ReadString('\n')
		case 5:
			fmt.Print("Entre com seu e-mail: ")
			user.email, _ = in.ReadString('\n')
		}
	}
	fmt.Print(user)
}
