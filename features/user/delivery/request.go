package delivery

type UserRequest struct {
	Name		string	`json:"name" form:"name"`
	Username	string	`json:"username" form:"username"`
	Email		string	`json:"email" form:"email"`
	Password	string	`json:"password" form:"password"`
}