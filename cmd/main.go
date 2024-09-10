package main

import (
	"CrudGolang/config"
	"CrudGolang/internal/Controller"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	id            string
	firstname     string
	lastname      string
	email         string
	telefone      string
	dt_nascimento string
	opcao         string
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
			fmt.Println("Digite o primeiro nome")
			fmt.Scanln(&firstname)
			fmt.Println("Digite o segundo nome")
			fmt.Scanln(&lastname)
			fmt.Println("Digite o email")
			fmt.Scanln(&email)
			fmt.Println("Digite o telefone")
			fmt.Scanln(&telefone)
			fmt.Println("Digite a data de nascimento 00/00/0000")
			fmt.Scanln(&dt_nascimento)
			Controller.Store(db, firstname, lastname, email, telefone, dt_nascimento)
			break

		case "3":
			fmt.Println("Digite o id que você que atualizar")
			fmt.Scanln(&id)
			var campos string
			fmt.Println("Campos Disponíveis: nome, sobrenome, email, telefone, data_nascimento (ex: nome email)")
			fmt.Println("Digite os campos a serem atualizados:")
			fmt.Scanln(&campos)

			arraycampos := strings.Split(campos, ",")

			for i := 0; i < len(arraycampos); i++ {
				switch arraycampos[i] {
				case "nome":
					fmt.Println("Digite o primeiro nome:")
					fmt.Scanln(&firstname)
				case "sobrenome":
					fmt.Println("Digite o segundo nome:")
					fmt.Scanln(&lastname)
				case "email":
					fmt.Println("Digite o email:")
					fmt.Scanln(&email)
				case "telefone":
					fmt.Println("Digite o telefone:")
					fmt.Scanln(&telefone)
				case "data_nascimento":
					fmt.Println("Digite a data de nascimento (DD/MM/AAAA):")
					fmt.Scanln(&dt_nascimento)
				default:
					fmt.Printf("Campo inválido: %s\n", arraycampos[i])
				}
			}
			Controller.Update(db, id, firstname, lastname, email, telefone, dt_nascimento)
			break
		case "4":
			fmt.Println("Digite o Id do usuario")
			fmt.Scanln(&id)
			Controller.Delete(db, id)
			break
		case "5":
			fmt.Println("Digite o Id do usuario")
			fmt.Scanln(&id)
			Controller.Show(db, id)
			break

		default:
			fmt.Println("Opção inválida. Encerrando o programa.")
			os.Exit(1)
		}

	}

}
