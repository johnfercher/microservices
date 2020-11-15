package infra

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetMysqlConnection() (*gorm.DB, error) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "AdminUser:AdminPassword@tcp(localhost:3306)/UserDb?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})

	/*db, err := gorm.Open(mysql.NewConnector(&mysql.Config{
		DriverName: "my_mysql_driver",
		DSN: "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local", // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
	}), &gorm.Config{})*/

	//return sql.Open("mysql", "AdminUser:AdminPassword@tcp(mysql:3306)/UserDb")
	//return sql.Open("mysql", "AdminUser:AdminPassword@tcp(localhost:3306)/UserDb")
}
