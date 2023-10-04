package tests

import (
	"fmt"
	"testing"
)

type ClassePessoaComMethod struct {
	Nome  string
	idade int16
}

func (c *ClassePessoaComMethod) exibeNome() (nome string) {
	return c.Nome
}

func (c *ClassePessoaComMethod) exibePessoa() {
	fmt.Printf("Nome: %s , Idade: %v \n", c.Nome, c.idade)
}

func TestOO(t *testing.T) {
	pessoa1 := ClassePessoaComMethod{
		Nome:  "Thiago",
		idade: 36,
	}

	fmt.Println(pessoa1.exibeNome())

	pessoa1.exibePessoa()

	if pessoa1.exibeNome() != "Thiago" {
		t.Error("Metodo exibeNome não funcionou!")
	}
}
