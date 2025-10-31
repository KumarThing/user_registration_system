package main
import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"encoding/json"
	

)
type User struct {
	Name string
	Password string

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var user [] User
	admin_password := "admin123"

	fmt.Println("\nWelcome to Secure Login System!")
	fmt.Println("---------------------------------")
	fmt.Println(`
1. register
2. login
3. show user
4. save
5. load
6. exit
	`)
	fmt.Println("---------------------------------")
	fmt.Print("Enter command: ")

	for {
		

		input,_ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "register" {
			fmt.Print("Enter user name: ")

			name,_ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			exists := false

			for i := range user {
				if strings.EqualFold(name, user[i].Name) {
					exists = true

				}
			}
			if exists {
				fmt.Println("Name already exists.")
				
			}

			fmt.Print("Enter Password: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)

			user = append(user, User{name,password})
			fmt.Printf("User %s registered succesfully!", name)


		} else if input == "login" {
			fmt.Print("Enter user name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Enter password: ")
			password,_ := reader.ReadString('\n')
			password = strings.TrimSpace(password)

			found := false

			for i := range user {
				if strings.EqualFold(name,user[i].Name) && strings.EqualFold(password,user[i].Password) {
					found = true
				}
			}
			if found {
				fmt.Printf("✅ Login successful! Welcome, %s!",name)
			}
			if !found {
				fmt.Printf("❌ Invalid password.")
			}
		} else if input == "show user" {
			if len(user) == 0 {
				fmt.Println("Empty data.")
			} else {
				fmt.Print("Enter admin password: ")

				password,_ := reader.ReadString('\n')
				password = strings.TrimSpace(password)

				if password == admin_password {
					fmt.Println("Registered User:")
					for i, u := range user {
						fmt.Printf("%d. %s\n",i+1, u.Name)
					} 
				} else {
					fmt.Println("Wrong admin password.")
				}

				
			}
		} else if input == "exit" {
			fmt.Println("\nbyee byee\n")
			break
		} else if input == "save" {
			if len(user) == 0 {
				fmt.Println("Empty data. Nothing to save.")
			} else {
				file,err := os.Create("user.json")
				if err != nil {
					fmt.Println("Error saving file.",err)
				}
				defer file.Close()

				data, err := json.MarshalIndent(user,"", " ")
				if err != nil {
					fmt.Println("Error converting .txt to json.",err)
				}
				file.Write(data)

				fmt.Println("Saved succefully.")
			}
		} else if input == "load" {
			data,err := os.ReadFile("user.json")
			if err != nil {
				fmt.Println("Error loading file",err)
			}
			err = json.Unmarshal(data, &user)
			if err != nil {
				fmt.Println("Error",err)
			}
			fmt.Println("Load succefully.")
		}

		fmt.Println(`
1. register
2. login
3. show user
4. save
5. load
6. exit
`)
	}

}