package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// type IPessoa interface {
// 	IndexListPessoa(db *sql.DB) ([]Pessoa, error)
// 	FindByPessoa(db *sql.DB, id string) (*Pessoa, error)
// 	Create(db *sql.DB, pessoa *Pessoa) error
// 	Update(db *sql.DB, pessoa *Pessoa) error
// 	Delete(db *sql.DB, pessoa *Pessoa) error
// }

type Pessoa struct {
	ID            string
	FirstName     string
	LastName      string
	Email         string
	Telefone      string
	Dt_Nascimento time.Time
}

func NewPessoa(firname, lastname, email, telefone string, dt_nascimento time.Time) *Pessoa {
	return &Pessoa{
		ID:            uuid.NewString(),
		FirstName:     firname,
		LastName:      lastname,
		Email:         email,
		Telefone:      telefone,
		Dt_Nascimento: dt_nascimento,
	}
}

func IndexListPessoa(db *sql.DB) ([]Pessoa, error) {
	rows, err := db.Query("SELECT * FROM pessoas")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var pessoas []Pessoa
	for rows.Next() {
		var pessoa Pessoa
		err = rows.Scan(&pessoa.ID, &pessoa.FirstName, &pessoa.LastName, &pessoa.Email, &pessoa.Telefone, &pessoa.Dt_Nascimento)
		if err != nil {
			panic(err)
		}
		pessoas = append(pessoas, pessoa)
	}
	return pessoas, nil
}

func FindByPessoa(db *sql.DB, id string) (*Pessoa, error) {
	stmt, err := db.Prepare("SELECT id,firstname,lastname,email,telefone,dt_nascimento FROM pessoas where id = ?")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	var p Pessoa
	err = stmt.QueryRow(id).Scan(&p.ID, &p.FirstName, &p.LastName, &p.Email, &p.Telefone, &p.Dt_Nascimento)
	if err != nil {
		panic(err)
	}
	return &p, nil
}

func Create(db *sql.DB, pessoa *Pessoa) error {
	stmt, err := db.Prepare("INSERT INTO Pessoas (firstname,lastname,email,telefone,dt_nascimento) values (?,?,?,?,?)")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = stmt.Exec(pessoa.ID, pessoa.FirstName, pessoa.Email, pessoa.Telefone, pessoa.Dt_Nascimento)
	if err != nil {
		panic(err)
	}
	return nil

}

func Update(db *sql.DB, pessoa *Pessoa) error {
	stmt, err := db.Prepare("UPDATE SET firstname = ?, lastname = ?, email = ?, telefone = ?, dt_nascimento = ? where id = ? ")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = stmt.Exec(pessoa.FirstName, pessoa.LastName, pessoa.Email, pessoa.Telefone, pessoa.Dt_Nascimento, pessoa.ID)
	if err != nil {
		panic(err)
	}
	return nil
}

func Delete(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM pessoa where id = ?")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	return nil
}
