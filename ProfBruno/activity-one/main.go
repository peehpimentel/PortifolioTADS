package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     string `json:"age"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

var People []Person

func main() {
	for {
		showMenu()
		reader := bufio.NewReader(os.Stdin)
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(1)
		}
		cmdString = strings.TrimSuffix(cmdString, "\n")
		cmdString = strings.TrimSpace(cmdString)

		switch cmdString {
		case "1":
			fmt.Println("Enter Person Details")
			addUser()
		case "2":
			fmt.Println("Get People")
			searchPerson()
		case "3":
			fmt.Println("Delete Person")
			deletePerson()
		case "4":
			fmt.Println("Update Person")
			updatePerson() //Implement a func for updating a person
		case "5":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid Option")
		}
	}
}

func showMenu() {
	fmt.Println("1. Add Person")
	fmt.Println("2. Get People")
	fmt.Println("3. Delete Person")
	fmt.Println("4. Update Person")
	fmt.Println("5. Exit")
}

func addUser() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("----------")
	fmt.Print("Enter Name: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Enter Address: ")
	address, _ := reader.ReadString('\n')

	fmt.Print("Enter Age: ")
	age, _ := reader.ReadString('\n')

	fmt.Print("Enter Email: ")
	email, _ := reader.ReadString('\n')

	fmt.Print("Enter Phone: ")
	phone, _ := reader.ReadString('\n')
	fmt.Println("----------")
	person := Person{
		Name:    strings.TrimSpace(name),
		Address: strings.TrimSpace(address),
		Age:     strings.TrimSpace(age),
		Email:   strings.TrimSpace(email),
		Phone:   strings.TrimSpace(phone),
	}

	People = loadPeople()
	People = append(People, person)
	saveUsers()
}

func loadPeople() []Person {
	file, err := os.ReadFile("people.json")
	if err != nil {
		fmt.Println("Error reading file")
	}
	_ = json.Unmarshal(file, &People)
	return People
}

func saveUsers() {
	file, _ := json.MarshalIndent(People, "", " ")
	_ = os.WriteFile("people.json", file, 0644)
}

func deletePerson() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("----------")
	fmt.Print("Enter Name of Person to Delete: ")
	nameToDelete, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(os.Stderr, err)
		return
	}
	nameToDelete = strings.TrimSpace(nameToDelete)
	People = loadPeople()
	filteredPeople := make([]Person, 0)
	for _, person := range People {
		if person.Name != nameToDelete {
			filteredPeople = append(filteredPeople, person)
		}
	}
	if len(filteredPeople) == len(People) {
		fmt.Println("\nPerson not found.")
		fmt.Println("----------")
	} else {
		People = filteredPeople
		saveUsers()
		fmt.Println("\nPerson deleted successfully!")
	}
}

func searchPerson() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("----------")
	fmt.Print("Enter name of Person to Search: ")
	nameToSearch, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(os.Stderr, err)
		return
	}
	nameToSearch = strings.TrimSpace(nameToSearch)
	People = loadPeople()
	found := false
	for _, person := range People {
		if strings.Contains(strings.ToLower(person.Name), strings.ToLower(nameToSearch)) {
			found = true
			fmt.Println("Name:", person.Name)
			fmt.Println("Address:", person.Address)
			fmt.Println("Age:", person.Age)
			fmt.Println("Email:", person.Email)
			fmt.Println("Phone:", person.Phone)
			fmt.Println("----------")
		}
	}
	if !found {
		fmt.Println("Person not found.")
		fmt.Println("----------")
	}
}

func updatePerson() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("----------")
	fmt.Print("Enter name of Person to Edit: ")
	nameToEdit, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(os.Stderr, err)
		return
	}
	nameToEdit = strings.TrimSpace(nameToEdit)
	People = loadPeople()
	found := false
	for i, person := range People {
		if strings.Contains(strings.ToLower(person.Name), strings.ToLower(nameToEdit)) {
			found = true
			fmt.Println("1 - Name")
			fmt.Println("2 - Address")
			fmt.Println("3 - Age")
			fmt.Println("4 - Email")
			fmt.Println("5 - Phone")
			fmt.Println("----------")
			fmt.Println("Choose a number to edit: ")
			numberToEdit, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(os.Stderr, err)
				os.Exit(1)
			}
			numberToEdit = strings.TrimSuffix(numberToEdit, "\n")
			numberToEdit = strings.TrimSpace(numberToEdit)
			switch numberToEdit {
			case "1":
				fmt.Print("Enter new Name: ")
				newName, _ := reader.ReadString('\n')
				newName = strings.TrimSpace(newName)
				People[i].Name = newName
			case "2":
				fmt.Print("Enter new Address: ")
				newAddress, _ := reader.ReadString('\n')
				newAddress = strings.TrimSpace(newAddress)
				if newAddress != "" {
					People[i].Address = newAddress
				}
			case "3":
				fmt.Print("Enter new Age: ")
				newAge, _ := reader.ReadString('\n')
				newAge = strings.TrimSpace(newAge)
				if newAge != "" {
					People[i].Age = newAge
				}
			case "4":
				fmt.Print("Enter new Email: ")
				newEmail, _ := reader.ReadString('\n')
				newEmail = strings.TrimSpace(newEmail)
				if newEmail != "" {
					People[i].Email = newEmail
				}
			case "5":
				fmt.Print("Enter new Phone: ")
				newPhone, _ := reader.ReadString('\n')
				newPhone = strings.TrimSpace(newPhone)
				if newPhone != "" {
					People[i].Phone = newPhone
				}
				break
			}
			saveUsers()
			fmt.Println("Person updated successfully!")
			fmt.Println("----------")
			break
		}
	}
	if !found {
		fmt.Println("Person not found.")
		fmt.Println("----------")
	}
}

