package entity

import "github.com/ktm-m/playground-go-firebase/internal/model"

type FireStoreInquiryUserByIDReq struct {
	ID string `json:"id"`
}

type FireStoreInquiryUsersByCityReq struct {
	City string `json:"city"`
}

type FireStoreInquiryUsersResp struct {
	ID                 string   `json:"id"`
	Firstname          string   `json:"firstname"`
	Lastname           string   `json:"lastname"`
	Email              string   `json:"email"`
	City               string   `json:"city"`
	IsCapitalCity      bool     `json:"is_capital_city"`
	FavoritePlace      []string `json:"favorite_place"`
	FavoritePlaceCount int      `json:"favorite_place_count"`
}

type FireStoreInquiryUserByIDResp struct {
	ID                 string   `json:"id"`
	Firstname          string   `json:"firstname"`
	Lastname           string   `json:"lastname"`
	Email              string   `json:"email"`
	City               string   `json:"city"`
	IsCapitalCity      bool     `json:"is_capital_city"`
	FavoritePlace      []string `json:"favorite_place"`
	FavoritePlaceCount int      `json:"favorite_place_count"`
}

type FireStoreInquiryUsersByCityResp struct {
	ID                 string   `json:"id"`
	Firstname          string   `json:"firstname"`
	Lastname           string   `json:"lastname"`
	Email              string   `json:"email"`
	City               string   `json:"city"`
	IsCapitalCity      bool     `json:"is_capital_city"`
	FavoritePlace      []string `json:"favorite_place"`
	FavoritePlaceCount int      `json:"favorite_place_count"`
}

func (resp *FireStoreInquiryUserByIDResp) ToCoreFireStoreInquiryUserByIDResp() *model.InquiryUserByIDResp {
	return &model.InquiryUserByIDResp{
		ID:                 resp.ID,
		Firstname:          resp.Firstname,
		Lastname:           resp.Lastname,
		Email:              resp.Email,
		City:               resp.City,
		IsCapitalCity:      resp.IsCapitalCity,
		FavoritePlace:      resp.FavoritePlace,
		FavoritePlaceCount: resp.FavoritePlaceCount,
	}
}

func (resp *FireStoreInquiryUsersByCityResp) ToCoreFireStoreInquiryUsersByCityResp() model.InquiryUsersByCityResp {
	return model.InquiryUsersByCityResp{
		ID:                 resp.ID,
		Firstname:          resp.Firstname,
		Lastname:           resp.Lastname,
		Email:              resp.Email,
		City:               resp.City,
		IsCapitalCity:      resp.IsCapitalCity,
		FavoritePlace:      resp.FavoritePlace,
		FavoritePlaceCount: resp.FavoritePlaceCount,
	}
}

func (resp *FireStoreInquiryUsersResp) ToCoreFireStoreInquiryUsersResp() model.InquiryUsersResp {
	return model.InquiryUsersResp{
		ID:                 resp.ID,
		Firstname:          resp.Firstname,
		Lastname:           resp.Lastname,
		Email:              resp.Email,
		City:               resp.City,
		IsCapitalCity:      resp.IsCapitalCity,
		FavoritePlace:      resp.FavoritePlace,
		FavoritePlaceCount: resp.FavoritePlaceCount,
	}
}

func ToFireStoreInquiryUserByIDReq(coreReq *model.InquiryUserByIDReq) *FireStoreInquiryUserByIDReq {
	return &FireStoreInquiryUserByIDReq{
		ID: coreReq.ID,
	}
}

func ToFireStoreInquiryUsersByCityReq(coreReq *model.InquiryUsersByCityReq) *FireStoreInquiryUsersByCityReq {
	return &FireStoreInquiryUsersByCityReq{
		City: coreReq.City,
	}
}

func MappingCoreFireStoreInquiryUsersResp(respData []FireStoreInquiryUsersResp) []model.InquiryUsersResp {
	coreResp := make([]model.InquiryUsersResp, 0, len(respData))
	for _, data := range respData {
		coreResp = append(coreResp, data.ToCoreFireStoreInquiryUsersResp())
	}

	return coreResp
}

func MappingCoreFireStoreInquiryUsersByCityResp(respData []FireStoreInquiryUsersByCityResp) []model.InquiryUsersByCityResp {
	coreResp := make([]model.InquiryUsersByCityResp, 0, len(respData))
	for _, data := range respData {
		coreResp = append(coreResp, data.ToCoreFireStoreInquiryUsersByCityResp())
	}

	return coreResp
}
