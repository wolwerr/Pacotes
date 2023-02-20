//aceita um conjunto de dados e o reduz a um tamenho fixo menor
//hashes são usadas em TUDO em programação, principalmente para buscar DADOS e DETECTAR ALTERAÇÕES
//em Go são divididas em 3 categorias: criptográficas, não criptográficas e de verificação
//lista das NÃO CRIPTOGRÁFICAS: adler32, crc32, crc64, fnv
//nesse código vamos criar um HASH e escrever nossos dados nele

package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	//criar um hash
	h := crc32.NewIEEE()
	//escrever dados no hash
	h.Write([]byte("código com pacotes hash"))
	//calcular o valor do hash
	v := h.Sum32()
	fmt.Println(v)
}
