package model

//jwt认证用户的信息
type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	//var auth Auth
	//db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	//if auth.ID > 0 {
	//	return true
	//
	//}
	//return false

	if username == "lynn" && password == "d58wsdfd5w2f8s1z25q" {
		return true
	}
	return false

}
