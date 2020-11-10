package serializer

type UResponse struct {
	UserId string 			`json:"user_id"`
	Username string 		`json:"username"`
	Phone string			`json:"phone"`
	Title string			`json:"title"`
}

//func UserInfoRes(data interface{}) UResponse {
//	var res UResponse
//	res.UserId = data.user_id
//	res.Username = data.user_name
//	res.Phone = data.phone
//	res.Title = data.title
//	return res
//}