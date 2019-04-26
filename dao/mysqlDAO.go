package dao

import (
	"SimpleHttpServer/conf"
	"SimpleHttpServer/data"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {

	addr := conf.MySQLUser + ":" + conf.MYSQLPWD + "@" +
		"tcp(" + conf.MySQLAddr + ":" + conf.MySQLPort + ")/" +
		conf.MYSQLDB + "?charset=utf8"
	var err error
	db, err = sql.Open("mysql", addr)
	if err != nil {
		log.Fatalf("mysql initialize problem: %v", err)
	}
	db.SetMaxIdleConns(1000)
	db.SetMaxOpenConns(2000)
	err = db.Ping()
	log.Println(err)
}

func ListSimpleUser() ([]*data.TUser, error) {

	rows, err := db.Query("select firstname,lastname,age from simpleuser")
	if err != nil {
		log.Fatalf("get simpleUser failed:%v", err)
		return nil, err
	}
	defer rows.Close()

	studentList := make([]*data.TUser, 0)
	for rows.Next() {
		tuser := &data.TUser{}
		rows.Scan(&tuser.FirstName, &tuser.LastName, &tuser.Age)
		fmt.Println(tuser)
		studentList = append(studentList, tuser)
	}
	return studentList, err
}

func CheckSelf(token string) (*data.AuthenRes, error) {
	var (
		id     string
		name   string
		device string
		ip     string
	)
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("check transaction failed:%v", err)
		return nil, err
	}

	err = tx.QueryRow("select u.id,u.name,i.device,i.ip from int_auth_token_cache as i inner join user as u on u.id=i.id where i.int_auth_token = ?", token).Scan(&id, &name, &device, &ip)
	if err != nil {
		log.Fatalf("err%v", err)
	}
	tx.Commit()

	authRes := &data.AuthenRes{Name: name, UserId: id, Greeting: "Welcome " + device + " user from " + ip}
	return authRes, err

}
