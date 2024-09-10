package Controller

import (
	"CrudGolang/internal/models"
	"database/sql"
	"fmt"
	"time"
)

// type PessoaController interface {
// 	Index()
// 	Show(id string)
// 	Store(pessoa models.Pessoa) error
// 	Update(pessoa models.Pessoa, id string) error
// 	Destroy(id string) error
// }

func Index(dbconexao *sql.DB) {
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
		fmt.Println("*******************************************************************************************")
	}
}

func Show(dbconexao *sql.DB, id string) {
	pessoa, err := models.FindByPessoa(dbconexao, id)
	if err != nil {
		panic(err)
	}
	// fmt.Println("Id: ", pessoa.ID)
	// fmt.Println("Primeiro Nome: ", pessoa.FirstName)
	// fmt.Println("Segundo Nome: ", pessoa.FirstName)
	// fmt.Println("Email: ", pessoa.Email)
	// fmt.Println("Telefone: ", pessoa.Telefone)
	fmt.Println("Data de Nascimento: ", pessoa.Dt_Nascimento)
	fmt.Println("*******************************************************************************************")
}

func Store(dbconexao *sql.DB, firstname, lastname, email, telefone, dt_nascimento string) {
	layout := "02/01/2006"
	data_nascimento, err := time.Parse(layout, dt_nascimento)
	if err != nil {
		panic(err)
	}
	ps := models.NewPessoa(firstname, lastname, email, telefone, data_nascimento)
	ps.Create(dbconexao)
	fmt.Println("Registro criado com sucesso!")
}

func Update(dbconexao *sql.DB, id, firstname, lastname, email, telefone, dt_nascimento string) {
	pessoa, err := models.FindByPessoa(dbconexao, id)

	if firstname == "" {
		firstname = pessoa.FirstName
	}
	if lastname == "" {
		lastname = pessoa.LastName
	}
	if email == "" {
		email = pessoa.Email
	}
	if telefone == "" {
		telefone = pessoa.Telefone
	}
	if dt_nascimento == "" {
		ps := models.NewPessoa(firstname, lastname, email, telefone, pessoa.Dt_Nascimento)
		ps.Update(dbconexao, id)
		fmt.Println("Registro criado com sucesso!")
		return
	}
	layout := "02/01/2006"
	data_nascimento, err := time.Parse(layout, dt_nascimento)
	if err != nil {
		panic(err)
	}
	ps := models.NewPessoa(firstname, lastname, email, telefone, data_nascimento)
	err = ps.Update(dbconexao, id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Registro Atualizado com sucesso!")
	return

}

func Delete(dbconexao *sql.DB, id string) {

	err := models.Delete(dbconexao, id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deletado com sucesso")
}
