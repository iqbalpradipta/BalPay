package delivery

import "github.com/iqbalpradipta/BalPay/features/user"

type userRespon struct {
	Name	string	`json:"name"`
	Email	string	`json:"email,omitempty"`
	Role	string 	`json:"role"`
	Token	string	`json:"token"`
}

func fromCore(data user.Core) userRespon {
	return userRespon{
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

func fromCoreToResponse(name, role, token string) userRespon {
	dataResponse := userRespon {
		Name: name,
		Role: role,
		Token: token,
	}
	return dataResponse
}