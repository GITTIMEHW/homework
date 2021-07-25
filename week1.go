package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type person struct {
	id int
	name string
	age int
}

func main() {

	rs, err := selectById(1)
	if err != nil{
		print(err.Error())
	}
	println(rs.name)
	println(rs.age)
	println(rs.id)

}

func selectById(id int)(person, error){
	p := person{}
	db, err := sql.Open("mysql",
		"user:password@tcp(**.aliyuncs.com:3306)/mytest")
	if err != nil {
		log.Fatal(err)
		return p, err
	}
	dbRs := db.QueryRow("select id,name,age from user where id=?", id).Scan(&p.id,&p.name,&p.age)
	if dbRs != nil{
		if dbRs == sql.ErrNoRows{
			return p, nil
		}else {
			log.Fatalln(dbRs)
			return p, dbRs
		}
	}
	return p, nil
}

