package Controller

import (
	"CrudGolang/internal/db"
	"CrudGolang/internal/models"
	"fmt"
)

type PessoaController interface {
	Index()
	Show(id string)
	Store(pessoa models.Pessoa) error
	Update(pessoa models.Pessoa, id string) error
	Destroy(id string) error
}

func Index() {
	dbconexao := db.Conn("root@tcp(localhost:3306)/pessoa?parseTime=true")
	pessoas, err := models.IndexListPessoa(dbconexao)
	if err != nil {
		panic(err)
	}

	for _, value := range pessoas {
		fmt.Println("Id: ", value.ID)
		fmt.Println("Primeiro Nome: ", value.FirstName)
		fmt.Println("Segundo Nome: ", value.FirstName)
		fmt.Println("Email: ", value.Email)
		fmt.Println("Telefone: ", value.Telefone)
		fmt.Println("Data de Nascimento: ", value.Dt_Nascimento)
	}
}

func Show(id string) {
	dbconexao := db.Conn("root@tcp(localhost:3306)/pessoa?parseTime=true")
	pessoa, err := models.FindByPessoa(dbconexao, id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Id: ", pessoa.ID)
	fmt.Println("Primeiro Nome: ", pessoa.FirstName)
	fmt.Println("Segundo Nome: ", pessoa.FirstName)
	fmt.Println("Email: ", pessoa.Email)
	fmt.Println("Telefone: ", pessoa.Telefone)
	fmt.Println("Data de Nascimento: ", pessoa.Dt_Nascimento)
}
