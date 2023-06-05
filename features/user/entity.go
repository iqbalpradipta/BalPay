package user

import (
	"time"

	"github.com/google/uuid"
)

type Core struct {
	ID			uuid.UUID
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
	ID			uuid.UUID
	Balance		uint
	CreatedAt	time.Time
}

type IusecaseInterface interface {
	GetUser() (data []Core, err error)
	GetUserById(id int) (data Core, err error)
	CreateData(data Core) (row int, err error)
	UpdateData(id int, updateData Core) (int , error)
	DeleteData(id int, deleteData Core) (int, error) 
}

type IdataInterface	interface {
	SelectAllData() (data []Core, err error)
	SelectDataById(id int) (data Core, err error)
	InsertData(data Core) (row int, err error)
	PostData(id int, updatedData Core) (int ,error)
	DelData(id int, deletedData Core) (int, error)
}

type AuthRequest struct {
	Email		string `json:"email" form:"email"`
	Password	string `json:"password" form:"password"`
}