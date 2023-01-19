package handler

import "sosmedapps/features/user"

type Register struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"username"`
}

func RegisterResponse(data user.Core) Register {
	return Register{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		UserName: data.UserName,
	}
}

type Login struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
}

func LoginResponse(data user.Core) Login {
	return Login{
		ID:       data.ID,
		Name:     data.Name,
		UserName: data.UserName,
	}
}

type Search struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
}

func SearchResponse(data user.Core) Search {
	return Search{
		ID:       data.ID,
		Name:     data.Name,
		UserName: data.UserName,
	}
}
