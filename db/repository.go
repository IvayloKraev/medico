package db

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"medico/config"
)

type Repository interface {
	Model(value interface{}) *gorm.DB
	Select(query interface{}, args ...interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	Exec(sql string, values ...interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	Raw(sql string, values ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Updates(value interface{}) *gorm.DB
	Delete(value interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	Preload(column string, conditions ...interface{}) *gorm.DB
	Scopes(funcs ...func(*gorm.DB) *gorm.DB) *gorm.DB
	ScanRows(rows *sql.Rows, result interface{}) error
	Transaction(fc func(tx Repository) error) (err error)
	Close() error
	DropTableIfExists(value interface{}) error
	AutoMigrate(value interface{}) error
}

type repository struct {
	name string
	db   *gorm.DB
}

func CreateNewRepository(name string, databaseConfig *config.DatabaseConfig) Repository {
	db, err := createConnection(databaseConfig)

	if err != nil {
		panic(err)
	}

	return &repository{name: name, db: db}
}

func createConnection(databaseConfig *config.DatabaseConfig) (*gorm.DB, error) {
	fmt.Println("Connecting to database...")
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		databaseConfig.Username, databaseConfig.Password,
		databaseConfig.Host, databaseConfig.DBName)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func (r *repository) Model(value interface{}) *gorm.DB {
	return r.db.Model(value)
}

func (r *repository) Select(query interface{}, args ...interface{}) *gorm.DB {
	return r.db.Select(query, args...)
}

func (r *repository) Find(out interface{}, where ...interface{}) *gorm.DB {
	return r.db.Find(out, where...)
}

func (r *repository) Exec(sql string, values ...interface{}) *gorm.DB {
	return r.db.Exec(sql, values...)
}

func (r *repository) First(out interface{}, where ...interface{}) *gorm.DB {
	return r.db.First(out, where...)
}

func (r *repository) Raw(sql string, values ...interface{}) *gorm.DB {
	return r.db.Raw(sql, values...)
}

func (r *repository) Create(value interface{}) *gorm.DB {
	return r.db.Create(value)
}

func (r *repository) Save(value interface{}) *gorm.DB {
	return r.db.Save(value)
}

func (r *repository) Updates(value interface{}) *gorm.DB {
	return r.db.Updates(value)
}

func (r *repository) Delete(value interface{}) *gorm.DB {
	return r.db.Delete(value)
}

func (r *repository) Where(query interface{}, args ...interface{}) *gorm.DB {
	return r.db.Where(query, args...)
}

func (r *repository) Preload(column string, conditions ...interface{}) *gorm.DB {
	return r.db.Preload(column, conditions...)
}

func (r *repository) Scopes(funcs ...func(*gorm.DB) *gorm.DB) *gorm.DB {
	return r.db.Scopes(funcs...)
}

func (r *repository) ScanRows(rows *sql.Rows, result interface{}) error {
	return r.db.ScanRows(rows, result)
}

func (r *repository) Close() error {
	sqlDB, _ := r.db.DB()
	return sqlDB.Close()
}

func (r *repository) DropTableIfExists(value interface{}) error {
	return r.db.Migrator().DropTable(value)
}

func (r *repository) AutoMigrate(value interface{}) error {
	return r.db.AutoMigrate(value)
}

func (r *repository) Transaction(fc func(tx Repository) error) (err error) {
	panicked := true
	tx := r.db.Begin()
	defer func() {
		if panicked || err != nil {
			tx.Rollback()
		}
	}()

	txrep := &repository{}
	txrep.db = tx
	err = fc(txrep)

	if err == nil {
		err = tx.Commit().Error
	}

	panicked = false
	return
}
