package util

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

const (
	mysqlConnStr = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"
)

func genMysqlConnStr() string{
	viper.SetConfigName("config")
	viper.AddConfigPath("conf")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil{
		panic(fmt.Errorf("config file error, please check: %s\n", err))
	}

	host := viper.Get("mysql.host")
	port := viper.Get("mysql.port")
	user := viper.Get("mysql.username")
	passWord := viper.Get("mysql.password")
	dbName := "user"
	fmt.Println(host, port)
	return fmt.Sprintf(mysqlConnStr, user, passWord, host, port, dbName)
}

var DB *gorm.DB
func init(){
	// 初始化创建一个数据库连接
	var err error
	connStr := genMysqlConnStr()
	fmt.Printf(connStr)
	DB, err = gorm.Open("mysql", connStr)
	if err != nil {
		panic("failed to connect db")
	}
	// sql语句打印在终端
	DB.LogMode(true)

	// 全局禁用表名复数
	DB.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`否则默认`users`,使用`TableName`设置的表名不受影响

	// 迁移数据库？？？
	DB.AutoMigrate()
}

