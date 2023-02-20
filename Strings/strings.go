// pacotes são caixinhas de funcões
// função contains
// contains: procura uma string dentro de outra string
//ex: Terra - rr (true)

package main	

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("computador", "dor")) // true
	fmt.Println(strings.Contains("test", "pes")) // false
}