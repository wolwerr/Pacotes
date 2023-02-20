// a cripto são os exemplos criptograficos
//vamos usar uma muito comum: SHA-1
package main

import (
	"fmt"
	"crypto/sha1"
)

func main() {
	//criar um hash
	h := sha1.New()
	//escrever dados no hash
	h.Write([]byte("código com pacotes hash"))
	//calcular o valor do hash
	bs := h.Sum([]byte{})
	fmt.Println(bs)
}