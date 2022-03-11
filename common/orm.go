package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"myGin/model"
)

type Orm struct {
	*gorm.DB
}

var DbEngine *Orm

func InitDB() (*Orm, error) {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}
	db.SingularTable(true) //如果使用gorm来帮忙创建表时，这里填写false的话gorm会给表添加s后缀，填写true则不会
	//db.LogMode(true)               //打印sql语句
	//开启连接池
	//db.DB().SetMaxIdleConns(100)   //最大空闲连接
	//db.DB().SetMaxOpenConns(100)   //最大连接数
	//db.DB().SetConnMaxLifetime(30) //最大生存时间(s)
	db.AutoMigrate(model.User{})
	orm := new(Orm)
	orm.DB = db
	DbEngine = orm
	return orm, err
}
func GetDB() *Orm {
	return DbEngine
}
