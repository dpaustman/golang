package main

import "fmt"

func main() {
	m := map[string] string{
		"course":  "golang",
		"teacher": "xiaowang",
		"time":    "friday",
		"name":    "history",
	}
	m2 := make(map[string]int)
	var m3 map[string] int
	fmt.Println(m, m2, m3)
	fmt.Println("traversing map")
	for k, v := range m {
		fmt.Println(k, v )
	}
	for _,v := range m {
		fmt.Println(v)
	}
	for k := range m {
		fmt.Println(k)
	}
	courseName, ok:= m["course"]
	fmt.Println(courseName, ok)

	coursName, ok:= m["cours"]
	fmt.Println(coursName, ok)
	if courseName, ok := m["course"]; ok {
		fmt.Println(courseName)
	} else {
		fmt.Println("key not exist,please check!")
	}

	fmt.Println("deleting keys from map")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
