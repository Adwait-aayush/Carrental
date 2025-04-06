package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)
var connection *pgxpool.Pool
func Connect() error {
	_=godotenv.Load()
	user:=os.Getenv("db_user")
	password:=os.Getenv("db_password")
	host:=os.Getenv("db_host")
	db:=os.Getenv("db_name")
	var connstring string = fmt.Sprintf("postgres://%s:%s@%s:5432/%s",user,password,host,db)
	conn,err:=pgxpool.New(context.Background(),connstring)
	if err!=nil{
		return err
	}
	connection=conn
	return nil
}
func Disconnect(){
	connection.Close()
}
func Connection() *pgxpool.Pool {
	return connection
}
func Ping(con *pgxpool.Pool) error {
	err := con.Ping(context.Background())

	if err != nil {
		return err
	}
	return nil
}

