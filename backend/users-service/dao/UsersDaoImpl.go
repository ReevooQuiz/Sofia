package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/zhanghanchong/users-service/entity"
	"os"
)

type UsersDaoImpl struct {
	db *sql.DB
}

var mysqlUrl string

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	mysqlUrl = "root:root@tcp(localhost)/sofia"
}

func (u *UsersDaoImpl) Init() (err error) {
	u.db, err = sql.Open("mysql", mysqlUrl)
	if err == nil {
		log.Info("Successfully connect to mysql.")
	}
	return err
}

func (u *UsersDaoImpl) Destruct() {
	_ = u.db.Close()
}

func (u *UsersDaoImpl) FindByEmail(email string) (user entity.Users, err error) {
	if u.db == nil {
		panic("UNINITIALIZED.")
	}
	var stmt *sql.Stmt
	stmt, err = u.db.Prepare("select * from users where email = ?")
	if err != nil {
		return user, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(email).Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Role)
	return user, err
}

func (u *UsersDaoImpl) FindById(id int64) (user entity.Users, err error) {
	if u.db == nil {
		panic("UNINITIALIZED.")
	}
	var stmt *sql.Stmt
	stmt, err = u.db.Prepare("select * from users where id = ?")
	if err != nil {
		return user, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Role)
	return user, err
}

func (u *UsersDaoImpl) FindByUsername(username string) (user entity.Users, err error) {
	if u.db == nil {
		panic("UNINITIALIZED.")
	}
	var stmt *sql.Stmt
	stmt, err = u.db.Prepare("select * from users where username = ?")
	if err != nil {
		return user, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(username).Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Role)
	return user, err
}

func (u *UsersDaoImpl) Insert(user entity.Users) (id int64, err error) {
	if u.db == nil {
		panic("UNINITIALIZED.")
	}
	var stmt *sql.Stmt
	stmt, err = u.db.Prepare("insert into users(username, password, email, role) values(?, ?, ?, ?)")
	if err != nil {
		return id, err
	}
	defer stmt.Close()
	var res sql.Result
	res, err = stmt.Exec(user.Username, user.Password, user.Email, user.Role)
	if err != nil {
		return id, err
	}
	id, err = res.LastInsertId()
	return id, err
}

func (u *UsersDaoImpl) Update(user entity.Users) (err error) {
	if u.db == nil {
		panic("UNINITIALIZED.")
	}
	var stmt *sql.Stmt
	stmt, err = u.db.Prepare("update users set username = ?, password = ?, email = ?, role = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Username, user.Password, user.Email, user.Role, user.Id)
	return err
}
