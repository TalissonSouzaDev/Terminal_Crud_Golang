package main

import (
	"CrudGolang/config"
	"CrudGolang/internal/Controller"
	"fmt"
	"log"
	"time"
)

func main() {
	for {
		cfg, err := config.LoadConfig()
		if err != nil {
			panic(err)
		}
		db, err := config.Connect(cfg.Dbdriver, cfg.Dbhost, cfg.Dbport, cfg.Dbname, cfg.Dbuser, cfg.Dbpassword)
		if err != nil {
			log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
		}
		fmt.Println("********************************** Painel De Crud **************************************")
		fmt.Println("1 Lista Usuarios")
		fmt.Println("2 Cadastra Usuario")
		fmt.Println("3 Atualizar Usuario")
		fmt.Println("4 Deleta Usuario")
		fmt.Println("5 Filtrar um Usuario")
		fmt.Println("6 Sair")
		fmt.Println("****************************************************************************************")

		var opcao string
		_, err = fmt.Scanln(&opcao)
		if err != nil {
			fmt.Println("Erro ao ler a entrada. Tente novamente.")
			continue
		}

		switch opcao {
		case "1":
			Controller.Index(db)
			break
		case "2":
			var (
				firstname string
				lastname  string
				email     string
				telefone  string
			)
			var dt_nascimento time.Time
			Controller.Store(db, firstname, lastname, email, telefone, dt_nascimento)
			break
		case "5":
			fmt.Println("Digite o Id do usuario")
			var id string
			fmt.Scanln(&id)
			Controller.Show(db, id)
			break
		}

	}

}
