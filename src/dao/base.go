package dao

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq" // here
	"strconv"
	"sync"
	"user-service/src/common/utils"
	"user-service/src/models"
)

// InitDatabase registers the database
func InitDatabase(db *models.Database) error {
	if err := utils.TestTCPConn(fmt.Sprintf("%s:%d", db.Host, db.Port), 60, 2); err != nil {
		return err
	}

	if err := orm.RegisterDriver("postgres", orm.DRPostgres); err != nil {
		return err
	}

	info := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		db.Host, strconv.Itoa(db.Port), db.Username, db.Password, db.Database, db.SSLMode)

	if err := orm.RegisterDataBase("default", "postgres", info, db.MaxIdleConns, db.MaxOpenConns); err != nil {
		return err
	}

	database, _ := orm.GetDB()
	database.SetMaxOpenConns(db.MaxOpenConns)

	fmt.Println("Register database completed")
	return nil
}

var globalOrm orm.Ormer
var once sync.Once

// GetOrmer :set ormer singleton
func GetOrmer() orm.Ormer {
	once.Do(func() {
		// override the default value(1000) to return all records when setting no limit
		orm.DefaultRowsLimit = -1
		globalOrm = orm.NewOrm()
	})
	return globalOrm
}
