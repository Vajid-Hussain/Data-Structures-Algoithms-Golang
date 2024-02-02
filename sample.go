package main

import (
	"fmt"
	"strconv"
)

func main() {
	value := strconv.FormatInt(3, 2)
	intVal,_:=strconv.Atoi(value)
	fmt.Println(value,intVal)

}
