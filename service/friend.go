package service

import (
	"github.com/unknwon/com"
	"time"
	"wh110api/model"
	"wh110api/pak/cacheHelper"
	"wh110api/pak/mgo"
)

//获取友链列表
func GetFriendList(ty int) *[]model.Friend {
	key := "friendlist_" + com.ToStr(ty)
	frs, ok := cacheHelper.GetCache().Get(key)
	if ok && frs != nil {
		temp := frs.(*[]model.Friend)
		return temp
	}

	fr := mgo.NewMgo("Friendly").FindFriendlyList(ty)

	if fr != nil {
		cacheHelper.GetCache().Set(key, fr, 2*time.Minute)
	}

	return fr
}
