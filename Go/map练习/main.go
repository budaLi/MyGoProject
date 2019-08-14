package main

import (
	"fmt"
)

func main() {
	users := make(map[string]map[string]string, 10)
	users["bab"] = make(map[string]string, 2)
	users["bab"]["pwd"] = "123"
	modifyUser(users, "libuda")
	fmt.Println(users)

}

func modifyUser(users map[string]map[string]string, name string) {
	if users[name] != nil {
		users[name]["pwd"] = "888"
	} else {
		users[name] = make(map[string]string, 2)
		users[name]["pwd"] = "888"
		users[name]["nickname"] = name
	}
}
