package user

import (
	"sippm/model/domain"
	"sippm/model/web/dosen"
	"sippm/model/web/lppm"
	"sippm/model/web/uppm"
)

type Response struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func ToResponse(user domain.Users) Response {
	return Response{
		Id:       user.Id,
		Username: user.Username,
	}
}

type DosenResponse struct {
	User  Response       `json:"user"`
	Dosen dosen.Response `json:"dosen"`
}

func ToDosenResponse(user domain.Users) *DosenResponse {
	return &DosenResponse{
		User:  ToResponse(user),
		Dosen: dosen.ToResponse(user.Dosen),
	}
}

func ToDosenResponses(users []domain.Users) []DosenResponse {
	var results []DosenResponse

	for _, user := range users {
		results = append(results, *ToDosenResponse(user))
	}

	return results
}

type UppmResponse struct {
	User Response      `json:"user"`
	Uppm uppm.Response `json:"uppm"`
}

func ToUppmResponse(user domain.Users) *UppmResponse {
	return &UppmResponse{
		User: ToResponse(user),
		Uppm: uppm.ToResponse(user.Uppm),
	}
}

func ToUppmResponses(users []domain.Users) []UppmResponse {
	var results []UppmResponse

	for _, user := range users {
		results = append(results, *ToUppmResponse(user))
	}

	return results
}

type LppmResponse struct {
	User Response      `json:"user"`
	Lppm lppm.Response `json:"lppm"`
}

func ToLppmResponse(user domain.Users) *LppmResponse {
	return &LppmResponse{
		User: ToResponse(user),
		Lppm: lppm.ToResponse(user.Lppm),
	}
}

func ToLppmResponses(users []domain.Users) []LppmResponse {
	var results []LppmResponse

	for _, user := range users {
		results = append(results, *ToLppmResponse(user))
	}

	return results
}
