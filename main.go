package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	id       int
	username string
	email    string
	age      int
}

var id int

var users map[int]User

func createUser() {
	clearConsole()
	fmt.Println("ingrese un valor para username")
	username := readLine()

	fmt.Println("ingrese un valor para email")
	email := readLine()

	fmt.Println("ingrese un valor para edad")
	age, err := strconv.Atoi(readLine())

	if err != nil {
		panic("no es posible convertir de un stringa un entero")
	}

	id++

	user := User{id, username, email, age}

	users[id] = user

	fmt.Println("\n")
}

func readUsers() {

	clearConsole()

	for id, user := range users {
		fmt.Println(id, "-", user.username)
	}

	fmt.Println("\n")
}

func updateUser() {
	fmt.Println("ingrese un valor para username")
	username := readLine()

	fmt.Println("ingrese un valor para email")
	email := readLine()

	fmt.Println("ingrese un valor para edad")
	age, _ := strconv.Atoi(readLine())

	fmt.Println("ingrese el Id del usuario a actualizar: ")
	id, _ := strconv.Atoi(readLine())

	_, ok := users[id]

	if ok {
		user := User{
			id,
			username,
			email,
			age,
		}
		users[id] = user
	}
}
func deleteUser() {
	clearConsole()
	fmt.Println("ingrese el id del usuario a eliminar: ")
	id, err := strconv.Atoi(readLine())

	if err != nil {
		panic("no es posible convertir de un stringa un entero")
	}

	_, ok := users[id]
	if ok {
		delete(users, id)
	}

	fmt.Println("usuario elimado success")
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func readLine() string {
	option, err := reader.ReadString('\n')

	if err != nil {
		panic("No es posible obtener el valor")
	}
	return strings.TrimSuffix(option, "\n")
}

func main() {

	var option string

	users = make(map[int]User)

	reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Opcion A), Crear")
		fmt.Println("Opcion B), Listar")
		fmt.Println("Opcion C), Actualizar")
		fmt.Println("Opcion D), Eliminar")

		fmt.Println("ingrese una opcion valida: ")
		option = readLine()

		switch option {
		case "a", "crear":
			createUser()
		case "b", "listar":
			readUsers()
		case "c", "actualizar":
			updateUser()
		case "d", "eliminar":
			deleteUser()
		default:
			fmt.Println("opcion no valida")
		}

		if option == "quit" || option == "q" {
			break
		}

	}
}
