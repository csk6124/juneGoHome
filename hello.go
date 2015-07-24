package main

import (
	"fmt"
 	"encoding/json"
	"github.com/csk6124/stringutil"
)

type Message struct {
    Name string
    Body string
    Time int64
}


func main() {
	fmt.Printf(stringutil.Reverse("!oG, olleH"))

	m := Message{"Hello", "June", 1242342343143423423}
	b, err := json.Marshal(m)
	

	if err == nil {
		json.Unmarshal(b, &m)
		fmt.Printf("test %s %s", m.Name, m.Body)
	}

}