package service

import (
	"github.com/unknwon/com"
	"reflect"
	"time"
	"wh110api/model"
	"wh110api/pak/cacheHelper"
	"wh110api/pak/logging"
	"wh110api/pak/mgo"
)

//获取新闻详情
func GetById(id string) model.News {
	var news = model.News{}
	key := "newsdetail_" + id
	n, ok := cacheHelper.GetCache().Get(key)
	if ok && n != nil {
		return n.(model.News)
	}

	err := mgo.NewMgo("Article").FindOneById(id).Decode(&news)
	if err != nil {
		logging.Error("获取新闻详情错误", err)
	}
	if err == nil && reflect.DeepEqual(news, model.News{}) {
		cacheHelper.GetCache().Set(key, news, time.Minute*2)
	}
	//	news.ID = news._id.String()
	return news
}

//获取新闻详情
func GetByOrder(order string) model.News {
	var news = model.News{}
	key := "newsdetailorder_" + order
	n, ok := cacheHelper.GetCache().Get(key)
	if ok && n != nil {
		return n.(model.News)
	}

	err := mgo.NewMgo("Article").FindOneByOrder(order).Decode(&news)
	if err != nil {
		logging.Error("获取新闻详情错误", err)
	}
	if err == nil && reflect.DeepEqual(news, model.News{}) {
		cacheHelper.GetCache().Set(key, news, time.Minute*2)
	}
	//	news.ID = news._id.String()
	return news
}

//获取推荐新闻列表
func GetTopList(page, size int64, ty int) *[]model.News {

	key := "newstopdetail_" + com.ToStr(page) + "_" + com.ToStr(size) + "_" + com.ToStr(ty)

	t, ok := cacheHelper.GetCache().Get(key)

	if ok && t != nil {
		temp := t.(*[]model.News)
		return temp
	}

	news := mgo.NewMgo("Article").FindTopList(page, size, ty)

	if news != nil {
		cacheHelper.GetCache().Set(key, news, time.Minute)
	}

	return news
}

//获取新闻列表
func GetList(page, size int64, ty int, isacs bool) (*[]model.News, int) {

	keyn := "newslist_" + com.ToStr(page) + "_" + com.ToStr(size) + "_" + com.ToStr(ty) + "_" + com.ToStr(isacs)
	keyc := "newslist_count_" + com.ToStr(page) + "_" + com.ToStr(size) + "_" + com.ToStr(ty) + "_" + com.ToStr(isacs)

	n, okn := cacheHelper.GetCache().Get(keyn)
	c, okc := cacheHelper.GetCache().Get(keyc)

	if okn && n != nil && okc && c != nil {
		temp := n.(*[]model.News)
		return temp, c.(int)
	}

	news, count := mgo.NewMgo("Article").FindList(page, size, ty, isacs)

	if news != nil && count > 0 {
		cacheHelper.GetCache().Set(keyn, news, time.Minute)
		cacheHelper.GetCache().Set(keyc, count, time.Minute)
	}
	return news, count
}

//获取随机推荐文章
func GetRecommendRandomNewsById(size int64, ty int, id string) *[]model.News {

	news := mgo.NewMgo("Article").GetRecommendRandomNews(size+1, ty)
	var sz = int(size)
	var ns = []model.News{}
	for key, value := range *news {
		if id != value.Id && key < sz {
			ns = append(ns, value)
		}
	}
	return &ns
}

func GetUpDownNewsById(ty int, id string) (up *model.News, down *model.News) {
	keyu := "newupbyid_" + id + "_" + com.ToStr(ty)
	keyd := "newdownbyid_" + id + "_" + com.ToStr(ty)

	u, oku := cacheHelper.GetCache().Get(keyu)

	d, okd := cacheHelper.GetCache().Get(keyd)

	if oku && okd && u != nil && d != nil {
		tempu := u.(*model.News)
		tempd := d.(*model.News)
		return tempu, tempd
	}

	var cnews = model.News{}
	mgo.NewMgo("Article").FindOneById(id).Decode(&cnews)
	up, down = mgo.NewMgo("Article").GetUpDownNewsById(cnews.Creatime, ty)

	if up != nil && down != nil {
		cacheHelper.GetCache().Set(keyu, up, time.Minute*30)
		cacheHelper.GetCache().Set(keyd, down, time.Minute*30)
	}

	return
}
func GetUpDownNewsByOrder(ty int, order string) (up *model.News, down *model.News) {

	keyu := "newupbyid_" + order + "_" + com.ToStr(ty)
	keyd := "newdownbyid_" + order + "_" + com.ToStr(ty)

	u, oku := cacheHelper.GetCache().Get(keyu)

	d, okd := cacheHelper.GetCache().Get(keyd)

	if oku && okd && u != nil && d != nil {
		tempu := u.(*model.News)
		tempd := d.(*model.News)
		return tempu, tempd
	}

	var cnews = model.News{}
	mgo.NewMgo("Article").FindOneByOrder(order).Decode(&cnews)
	up, down = mgo.NewMgo("Article").GetUpDownNewsById(cnews.Creatime, ty)

	if up != nil && down != nil {
		cacheHelper.GetCache().Set(keyu, up, time.Minute*30)
		cacheHelper.GetCache().Set(keyd, down, time.Minute*30)
	}

	return
}
