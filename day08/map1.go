package main

import "fmt"

type Personinfo struct {
	ID string
	Name string
	Address string
}
func main() {
	var personDB map[string] Personinfo
	personDB = make(map[string] Personinfo)
	personDB["1"] = Personinfo{"1", "Tom", "chengdu"}
	personDB["2"] = Personinfo{"2", "Jack", "nanchong"}
	personDB["3"] = Personinfo{"3", "Jack1", "nanbu"}
	personDB["4"] = Personinfo{"4", "Jack2", "panlong"}
	person, ok := personDB["1"]
	if ok {
		fmt.Println("Found person",person.Name,"with ID 1")
	} else {
		fmt.Println("Not found this person with ID　１")
	}
	delete(personDB, "4")
	fmt.Println(personDB["4"])
 }
