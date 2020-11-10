package model

type User struct {
	//gorm.Model
	ID       	uint   `json:"id"`
	UserId		string `json:"user_id"`
	Username	string `json:"name"`
	Password 	string `json:"password"`
	Title	 	string `json:"title"`
	Nickname 	string `json:"nickname"`
	Salt 	 	string `json:"nickname"`
}
