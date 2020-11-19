package mgo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"wh110api/pak/setting"
)

type Database struct {
	Mongo *mongo.Client
}

//链接数据库
func SetConnect() *mongo.Client {

	url := setting.DBSetting.DB_IP
	fmt.Println("数据库链接" + setting.DBSetting.DB_IP)
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		fmt.Println("数据库连接失败", err)
	}
	return client
}

var DB *Database

func Init() {
	DB = &Database{
		Mongo: SetConnect(),
	}
}
