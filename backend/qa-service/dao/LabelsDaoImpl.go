package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/zhanghanchong/qa-service/entity"
	"os"
)

type LabelsDaoImpl struct {
	db *sql.DB
}

var mysqlUrl string

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	_ = godotenv.Load(os.Getenv("WORK_DIR") + "credentials.env")
	mysqlUrl = os.Getenv("MYSQL_URL")
}

func (l *LabelsDaoImpl) Init() (err error) {
	l.db, err = sql.Open("mysql", mysqlUrl)
	return err
}

func (l *LabelsDaoImpl) Destruct() {
	_ = l.db.Close()
}

func (l *LabelsDaoImpl) FindByTitle(title string) (label entity.Labels, err error) {
	var stmt *sql.Stmt
	stmt, err = l.db.Prepare("select * from labels where title = ?")
	if err != nil {
		return label, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(title).Scan(&label.Lid, &label.Title)
	return label, err
}

func (l *LabelsDaoImpl) Insert(label entity.Labels) (lid int64, err error) {
	var stmt *sql.Stmt
	stmt, err = l.db.Prepare("insert into labels(title) values(?)")
	if err != nil {
		return lid, err
	}
	defer stmt.Close()
	var res sql.Result
	res, err = stmt.Exec(label.Title)
	if err != nil {
		return lid, err
	}
	lid, err = res.LastInsertId()
	return lid, err
}
