package model

import (
	"fmt"
	"github.com/CaoJiayuan/go-workflow/workflow-engine/logger"
	"log"
	"time"

	config "github.com/CaoJiayuan/go-workflow/workflow-config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	dbLogger "gorm.io/gorm/logger"
)

var db *gorm.DB

// Model 其它数据结构的公共部分
type Model struct {
	ID int `gorm:"primary_key" json:"id,omitempty"`
}

// 配置
var conf = *config.Config

// Setup 初始化一个db连接
func Setup() {
	log.Println("数据库初始化！！")
	var err error
	db, err = getDb()
	if err != nil {
		log.Fatalf("数据库连接失败 err: %v", err)
	}

	db.AutoMigrate(&Procdef{})
	db.AutoMigrate(&Execution{})
	db.AutoMigrate(&Task{})
	db.AutoMigrate(&ProcInst{})
	db.AutoMigrate(&Identitylink{})
	db.AutoMigrate(&ExecutionHistory{})
	db.AutoMigrate(&IdentitylinkHistory{})
	db.AutoMigrate(&ProcInstHistory{})
	db.AutoMigrate(&TaskHistory{})
	db.AutoMigrate(&ProcdefHistory{})
	//---------------------历史纪录------------------------------
}

func getDb() (*gorm.DB, error) {
	var dialector gorm.Dialector

	dialector = mysql.Open(fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName))
	if conf.DbType == "postgres" {

		dialector = postgres.Open(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName))
	}

	l := dbLogger.New(
		logger.GetLogger("db"), // io writer
		dbLogger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      dbLogger.Info, // Log level
			Colorful:      false,         // 禁用彩色打印
		})

	return gorm.Open(dialector, &gorm.Config{
		Logger:      l,
		PrepareStmt: true,
	})
}

// GetDB getdb
func GetDB() *gorm.DB {
	return db
}

// GetTx GetTx
func GetTx() *gorm.DB {
	return db.Begin()
}
