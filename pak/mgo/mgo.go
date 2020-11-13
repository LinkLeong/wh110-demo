package mgo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"wh110api/model"
	"wh110api/pak/logging"
	"wh110api/pak/setting"
)

type mgo struct {
	database   string
	collection string //要连接的集合
}

func NewMgo(collection string) *mgo {
	return &mgo{
		database:   setting.DBSetting.DB_NAME,
		collection: collection,
	}
}

//func (m *mgo) Connect() *mongo.Collection {
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.1.176"))
//	if err != nil {
//		log.Println(err)
//	}
//	collection := client.Database("News").Collection(m.collection)
//	return collection
//}

//根据id查下一个数据
func (m *mgo) FindOneById(id string) *mongo.SingleResult {
	client := DB.Mongo
	//collection, _ := client.Database(m.database).Collection(m.collection).Clone()
	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logging.Error("根据id获取新闻详情错误", err)
	}
	filter := bson.D{{"_id", objid}}
	//var news =News{}
	//singleResult := collection.FindOne(context.TODO(), filter)
	singleResult := client.Database(m.database).Collection(m.collection).FindOne(context.TODO(), filter)
	return singleResult
}

//根据order查下一个数据
func (m *mgo) FindOneByOrder(order string) *mongo.SingleResult {
	client := DB.Mongo
	//collection, _ := client.Database(m.database).Collection(m.collection).Clone()

	filter := bson.D{{"Order", order}}
	//var news =News{}
	//singleResult := collection.FindOne(context.TODO(), filter)
	singleResult := client.Database(m.database).Collection(m.collection).FindOne(context.TODO(), filter)
	return singleResult
}

//查询列表
func (m *mgo) FindTopList(page, size int64, tp int) *[]model.News {
	client := DB.Mongo
	//filter := bson.M{"State": 1,"Top": true}
	filter := bson.D{{"State", 1}, {"Top", true}, {"Type", tp}}
	opt := options.Find()
	opt.SetLimit(size)
	opt.SetSkip(size * (page - 1))
	opt.Sort = bson.D{{"Creatime", 1}}
	singleResult, err := client.Database(m.database).Collection(m.collection).Find(context.TODO(), filter, opt)

	if err != nil {
		logging.Error("查询推荐列表错误", err)
	}
	var news []model.News
	errr := singleResult.All(context.TODO(), &news)
	if errr != nil {
		logging.Error("查询推荐列表解析错误", err)
	}

	defer singleResult.Close(context.TODO())

	return &news
}

//查询列表
func (m *mgo) FindList(page, size int64, ty int, isacs bool) (*[]model.News, int) {
	client := DB.Mongo
	filter := bson.D{{"State", 1}, {"Type", ty}}
	opt := options.Find()
	opt.SetLimit(size)
	opt.SetSkip(size * (page - 1))
	acs := 1
	if !isacs {
		acs = -1
	}
	opt.Sort = bson.D{{"Creatime", acs}}
	result, err := client.Database(m.database).Collection(m.collection).Find(context.TODO(), filter)
	singleResult, err := client.Database(m.database).Collection(m.collection).Find(context.TODO(), filter, opt)

	if err != nil {
		logging.Error("查询列表错误", err)
	}
	var news []model.News
	errr := singleResult.All(context.TODO(), &news)
	if errr != nil {
		logging.Error("查询列表解析错误", err)
	}

	defer singleResult.Close(context.TODO())
	defer result.Close(context.TODO())

	return &news, result.RemainingBatchLength()
}

//查询友链列表
func (m *mgo) FindFriendlyList(ty int) *[]model.Friend {
	client := DB.Mongo
	filter := bson.D{{"IsEnable", true}, {"Type", ty}}
	singleResult, err := client.Database(m.database).Collection(m.collection).Find(context.TODO(), filter)

	if err != nil {
		logging.Error("查询友链错误", err)
	}

	var friend []model.Friend
	erri := singleResult.All(context.TODO(), &friend)

	if erri != nil {
		logging.Error("查询友链解析错误", err)
	}

	defer singleResult.Close(context.TODO())

	return &friend
}

//
func (m *mgo) GetRecommendRandomNews(size int64, ty int) *[]model.News {
	client := DB.Mongo
	filter := mongo.Pipeline{
		bson.D{{"$match", bson.M{"Type": ty}}},
		bson.D{{"$sample", bson.M{"size": size}}},
	}
	opt := options.Find()
	opt.SetLimit(size)
	singleResult, err := client.Database(m.database).Collection(m.collection).Aggregate(context.TODO(), filter)

	if err != nil {
		logging.Error("查询随机新闻错误", err)
	}

	var news []model.News
	erri := singleResult.All(context.TODO(), &news)

	if erri != nil {
		logging.Error("查询随机新闻解析错误", err)
	}

	defer singleResult.Close(context.TODO())

	return &news
}
func (m *mgo) GetUpDownNewsById(tm time.Time, ty int) (up *model.News, down *model.News) {
	client := DB.Mongo
	//filter :=mongo.Pipeline{
	//	bson.D{{"$match", bson.M{"Type":ty}}},
	//	bson.D{{"$sample", bson.M{"size":size}}},
	//
	//}
	upfilter := bson.D{{"Type", ty}, {"Creatime", bson.D{{"$gt", tm}}}}
	downfilter := bson.D{{"Type", ty}, {"Creatime", bson.D{{"$lt", tm}}}}
	opt := options.Find()
	opt.SetLimit(1)

	//获取上一条
	u, ue := client.Database(m.database).Collection(m.collection).Find(context.TODO(), upfilter, opt)
	defer u.Close(context.TODO())

	if ue != nil {
		logging.Error("获取上一条", up)
	}

	var unews []model.News
	u.All(context.TODO(), &unews)
	if unews != nil {
		up = &unews[0]
	}

	optd := options.Find()
	optd.SetLimit(1)
	optd.Sort = bson.D{{"Creatime", -1}}

	//获取下一条
	v, err := client.Database(m.database).Collection(m.collection).Find(context.TODO(), downfilter, optd)
	defer v.Close(context.TODO())
	if err != nil {
		logging.Error("获取下一条错误", err)
	}

	var news []model.News
	v.All(context.TODO(), &news)
	if news != nil {
		down = &news[0]
	}
	return
}
