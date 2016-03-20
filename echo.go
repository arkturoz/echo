package main

import(
	"fmt"
	"os"
	"strings"
)

func main(){
	fmt.Println(strings.Join(os.Args[0:], " "))

	for index, value := range os.Args[1:] {
		fmt.Println(index, value)
	}
}
