package postgres

import (
	"os"
	"project2/model/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable TimeZone=Asia/Jakarta"
	// dsn := "host=localhost user=postgres password=dirham dbname=todos port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{})

	return db
}
