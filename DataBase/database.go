package main

import (
	//Это всего лишь универсальный интерфейс с реазизациями. Необходим драйвер базы данных, который ниже.
	"database/sql"
	"fmt"

	//Драйвер базы данных.
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/lib/pq"
)

type product struct {
	id      int
	model   string
	company string
	price   int
}

func main() {
	//Если не будет работать, надо будет перепроверить пользователя, пароль, базу данных, права.
	type config struct {
		User     string `env:"USER"`
		Password string `env:"PASSWORD"`
		DBName   string `env:"DBNAME"`
	}
	var cfg config
	cleanenv.ReadConfig("/home/cpp/project/module/DataBase/.env", &cfg)

	//Для работы этих строчек нужен .env файл вида:
	//USER=user
	//PASSWORD=password
	//DBNAME=name
	var connStr = fmt.Sprintf("user= %s password= %s dbname= %s sslmode=disable", cfg.User, cfg.Password, cfg.DBName)
	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	add_db(db)
	search_db(db)
	del_db(db)
	fmt.Print("end")
}

func add_db(db *sql.DB) {

	result, err := db.Exec("insert into Products (model, company, price) values ('iPhone X', $1, $2)", "Apples", 72000)

	if err != nil {
		panic(err)
	}
	fmt.Print(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}

func del_db(db *sql.DB) {
	del_arr := []int{7, 8}
	for _, i := range del_arr {
		result, err := db.Exec("delete from Products where id = $1", i)
		if err != nil {
			panic(err)
		}
		fmt.Println(result.RowsAffected())
	}
}

func search_db(db *sql.DB) {
	products := []product{}

	rows, err := db.Query("SELECT * FROM public.products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		p := product{}
		err := rows.Scan(&p.id, &p.model, &p.company, &p.price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}

	for _, p := range products {
		fmt.Println(p.id, p.model, p.company, p.price)
	}
}
