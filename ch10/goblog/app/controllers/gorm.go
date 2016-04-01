package controllers

import (
	"goblog/app/models"

	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm" // ➊ gorm 패키지 임포트
	"github.com/revel/revel"
)

const (
	DefaultName,
	DefaultRole,
	DefaultUsername,
	DefaultPassword = "Admin", "admin", "admin", "admin"
)

var (
	db gorm.DB
)

// ➋ GormController 정의
type GormController struct {
	*revel.Controller
	Txn *gorm.DB
}

// ➌ 데이터베이스 초기화
func InitDB() {
	var (
		driver, spec string
		found        bool
	)

	// Read configuration.
	if driver, found = revel.Config.String("db.driver"); !found {
		revel.ERROR.Fatal("No db.driver found.")
	}
	if spec, found = revel.Config.String("db.spec"); !found {
		revel.ERROR.Fatal("No db.spec found.")
	}

	// Open a connection.
	var err error
	db, err = gorm.Open(driver, spec)
	if err != nil {
		revel.ERROR.Fatal(err)
	}

	// Enable Logger
	db.LogMode(true)
	migrate()
}

// ➍ 테이블 생성
func migrate() {
	db.AutoMigrate(&models.Post{}, &models.Comment{}, &models.User{})
	bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(DefaultPassword), bcrypt.DefaultCost)
	db.Where(models.User{Name: DefaultName, Role: DefaultRole, Username: DefaultUsername}).
		Attrs(models.User{Password: string(bcryptPassword)}).
		FirstOrCreate(&models.User{})
}

// ➎ 트랜잭션 설정
// Begin a transaction
func (c *GormController) Begin() revel.Result {
	c.Txn = db.Begin()
	return nil
}

// Rollback if it's still going (must have panicked).
func (c *GormController) Rollback() revel.Result {
	if c.Txn != nil {
		c.Txn.Rollback()
		c.Txn = nil
	}
	return nil
}

// Commit the transaction.
func (c *GormController) Commit() revel.Result {
	if c.Txn != nil {
		c.Txn.Commit()
		c.Txn = nil
	}
	return nil
}
