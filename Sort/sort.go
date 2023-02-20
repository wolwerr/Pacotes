//Ordena dados arbitrários
//usaremos a fubção sort.Interface
//sort.Interface requer 3 métodos(Len, Less, Swap)

package main

import (
	"fmt"
	"sort"
)

type Dados struct {
	Nome string
	Idade int
}

type ParaNome []Dados
func (ps ParaNome) Len() int {
	return len(ps)
}
func (ps ParaNome) Less(i, j int) bool { //item na posicao i é menor que o item na posição j
	return ps[i].Nome < ps[j].Nome //ordena por nome
}
func (ps ParaNome) Swap(i, j int) { //troca os itens na posição i e j
	ps[i], ps[j] = ps[j], ps[i]
}
func main() {
	criancas := []Dados{
		{"Julia", 5},
		{"João", 10},
		{"Pedro", 10},
	}
	sort.Sort(ParaNome(criancas))
	fmt.Println(criancas)
}