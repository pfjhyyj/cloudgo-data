package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

func count() int {
	return len(UserInfoService.FindAll())
}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	u.UID = int(count())
	_, err := engine.InsertOne(u)

	checkErr(err)
	return nil
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	everyone := make([]UserInfo, 0)
	err := engine.Find(&everyone)
	checkErr(err)
	return everyone
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	var user UserInfo
	users := make([]UserInfo, 0)
	err := engine.Where("u_i_d = ?", id).Find(&users)
	if len(users) != 0 {
		user = users[0]
	} else {
		return &user
	}
	checkErr(err)
	return &user
}
