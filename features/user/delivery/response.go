package delivery

import "github.com/iqbalpradipta/BalPay/features/user"

type userRespon struct {
	ID		uint	`json:"id"`
	Name	string	`json:"name"`
	Email	string	`json:"email,omitempty"`
	Role	string 	`json:"role"`
}

func fromCore(data user.Core) userRespon {
	return userRespon{
		ID:		data.ID,
		Name:	data.Name,
		Email:	data.Email,
		Role:	data.Role,
	}
}

func fromCoreList(data []user.Core) []userRespon {
	var dataResponse []userRespon
	for _, v := range data {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}

func fromCoreToResponse(dataCore user.Core) userRespon {
	dataResponse := userRespon {
		Name: dataCore.Name,
		Email: dataCore.Email,
		Role: dataCore.Role,
	}
	return dataResponse
}