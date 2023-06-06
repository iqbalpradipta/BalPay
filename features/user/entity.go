package user

import (
	"time"
)

type Core struct {
	ID			uint
	Name		string
	Username	string
	Email		string
	Password	string
	Role		string
	CreatedAt	time.Time
	UpdatedAt	time.Time
	Account		[]AccountCore	
}

type AccountCore struct  {
	ID			uint
	Balance		uint
	CreatedAt	time.Time
}

type IusecaseInterface interface {
	GetUser() (data []Core, err error)
	GetUserById(id int) (data Core, err error)
	CreateData(data Core) (row int, err error)
	UpdateData(id int, updateData Core) (int , error)
	DeleteData(id int, deleteData Core) (int, error) 
	LoginData(email, password string) (string, string, string)
}

type IdataInterface	interface {
	SelectAllData() (data []Core, err error)
	SelectDataById(id int) (data Core, err error)
	InsertData(data Core) (row int, err error)
	PostData(id int, updatedData Core) (int ,error)
	DelData(id int, deletedData Core) (int, error)
	AuthData(email string) (Core, error)
}

type AuthRequest struct {
	Email		string `json:"email" form:"email"`
	Password	string `json:"password" form:"password"`
}