package pet

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Persistence interface {
	FindById(id int64) *Pet
	FindByStatus(Status) []*Pet
	Update(pet *Pet) error
	Insert(pet *Pet) (*Pet, error)
	Delete(id int64) error
	GetStatusCounts() Inventory
}

func InitMySQLPersistence(host, port, user, pass, dbName string) (*MySQLPersistence, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbName)
	dbConn, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}
	return &MySQLPersistence{connString:connString, db:dbConn}, nil
}

type MySQLPersistence struct {
	connString string
	db *sql.DB
}

func (p *MySQLPersistence) FindById(id int64) *Pet {}
func (p *MySQLPersistence) FindByStatus(Status) []*Pet {}
func (p *MySQLPersistence) Update(pet *Pet) error {}
func (p *MySQLPersistence) Insert(pet *Pet) (*Pet, error) {}
func (p *MySQLPersistence) Delete(id int64) error {}
func (p *MySQLPersistence) GetStatusCounts() Inventory {}


