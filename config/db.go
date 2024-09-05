package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(dbdriver, dbhost, dbport, dbname, dbuser, dbpassword string) (*sql.DB, error) {

	// Verifique se a senha está vazia e construa a string de conexão de acordo.
	var dsn string
	if dbpassword != "" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbuser, dbpassword, dbhost, dbport, dbname)
	} else {
		dsn = fmt.Sprintf("%s@tcp(%s:%s)/%s?parseTime=true", dbuser, dbhost, dbport, dbname)
	}
	// Tente abrir a conexão com o banco de dados.
	conn, err := sql.Open(dbdriver, dsn)
	if err != nil {
		// Retorne o erro em vez de usar panic.
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %w", err)
	}

	// Verifique se a conexão é válida.
	if err = conn.Ping(); err != nil {
		// Feche a conexão antes de retornar o erro.
		conn.Close()
		return nil, fmt.Errorf("erro ao pingar o banco de dados: %w", err)
	}
	return conn, nil
}
