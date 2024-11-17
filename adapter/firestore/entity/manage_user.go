package entity

import (
	"cloud.google.com/go/firestore"
	"github.com/ktm-m/playground-go-firebase/internal/model"
)

type FireStoreInsertUserReq struct {
	Firstname          string   `json:"firstname"`
	Lastname           string   `json:"lastname"`
	Age                int      `json:"age"`
	Email              string   `json:"email"`
	City               string   `json:"city"`
	FavoritePlace      []string `json:"favorite_place"`
	FavoritePlaceCount int      `json:"favorite_place_count"`
}

type FireStoreUpsertUserReq struct {
	Firstname          string   `json:"firstname"`
	Lastname           string   `json:"lastname"`
	Age                int      `json:"age"`
	Email              string   `json:"email"`
	City               string   `json:"city"`
	FavoritePlace      []string `json:"favorite_place"`
	FavoritePlaceCount int      `json:"favorite_place_count"`
}

type FireStoreUpdateCapitalCityFlagReq struct {
	IsCapitalCity bool `json:"is_capital_city"`
}

type FireStoreIncrementUpdateUserReq struct {
	Firstname          *string  `json:"firstname"`
	Lastname           *string  `json:"lastname"`
	Age                *int     `json:"age"`
	Email              *string  `json:"email"`
	City               *string  `json:"city"`
	FavoritePlace      []string `json:"favorite_place"`
	FavoritePlaceCount *int     `json:"favorite_place_count"`
}

type FireStoreAddOneFavoritePlaceReq struct {
	FavoritePlace string `json:"favorite_place"`
}

type FireStoreAddMultipleFavoritePlaceReq struct {
	FavoritePlace []string `json:"favorite_place"`
}

type FireStoreDeleteUserByIDReq struct {
	ID string `json:"id"`
}

type FireStoreDeleteCapitalCityFlagReq struct {
	ID string `json:"id"`
}

func ToFireStoreInsertUserReq(req *model.InsertUserReq) (string, *FireStoreInsertUserReq) {
	return req.ID, &FireStoreInsertUserReq{
		Firstname:          req.Firstname,
		Lastname:           req.Lastname,
		Age:                req.Age,
		Email:              req.Email,
		City:               req.City,
		FavoritePlace:      req.FavoritePlace,
		FavoritePlaceCount: req.FavoritePlaceCount,
	}
}

func ToFireStoreUpsertUserReq(req *model.UpsertUserReq) (string, *FireStoreUpsertUserReq) {
	return req.ID, &FireStoreUpsertUserReq{
		Firstname:          req.Firstname,
		Lastname:           req.Lastname,
		Age:                req.Age,
		Email:              req.Email,
		City:               req.City,
		FavoritePlace:      req.FavoritePlace,
		FavoritePlaceCount: req.FavoritePlaceCount,
	}
}

func ToFireStoreUpdateCapitalCityFlagReq(req *model.UpdateCapitalCityFlagReq) (string, *FireStoreUpdateCapitalCityFlagReq) {
	return req.ID, &FireStoreUpdateCapitalCityFlagReq{
		IsCapitalCity: req.IsCapitalCity,
	}
}

func ToFireStoreIncrementUpdateUserReq(req *model.IncrementUpdateUserReq) (string, *FireStoreIncrementUpdateUserReq) {
	return req.ID, &FireStoreIncrementUpdateUserReq{
		Firstname:          req.Firstname,
		Lastname:           req.Lastname,
		Age:                req.Age,
		Email:              req.Email,
		City:               req.City,
		FavoritePlace:      req.FavoritePlace,
		FavoritePlaceCount: req.FavoritePlaceCount,
	}
}

func ToFireStoreAddOneFavoritePlaceReq(req *model.AddOneFavoritePlaceReq) (string, *FireStoreAddOneFavoritePlaceReq) {
	return req.ID, &FireStoreAddOneFavoritePlaceReq{
		FavoritePlace: req.FavoritePlace,
	}
}

func ToFireStoreAddMultipleFavoritePlaceReq(req *model.AddMultipleFavoritePlaceReq) (string, *FireStoreAddMultipleFavoritePlaceReq) {
	return req.ID, &FireStoreAddMultipleFavoritePlaceReq{
		FavoritePlace: req.FavoritePlace,
	}
}

func ToFireStoreDeleteUserByIDReq(req *model.DeleteUserByIDReq) *FireStoreDeleteUserByIDReq {
	return &FireStoreDeleteUserByIDReq{
		ID: req.ID,
	}
}

func ToFireStoreDeleteCapitalCityFlagReq(req *model.DeleteCapitalCityFlagReq) *FireStoreDeleteCapitalCityFlagReq {
	return &FireStoreDeleteCapitalCityFlagReq{
		ID: req.ID,
	}
}

func (req *FireStoreIncrementUpdateUserReq) BuildIncrementalUpdateData() []firestore.Update {
	updateData := make([]firestore.Update, 0)

	if req.Firstname != nil {
		updateData = append(updateData, firestore.Update{
			Path:  "firstname",
			Value: req.Firstname,
		})
	}

	if req.Lastname != nil {
		updateData = append(updateData, firestore.Update{
			Path:  "lastname",
			Value: req.Lastname,
		})
	}

	if req.Age != nil {
		updateData = append(updateData, firestore.Update{
			Path:  "age",
			Value: req.Age,
		})
	}

	if req.Email != nil {
		updateData = append(updateData, firestore.Update{
			Path:  "email",
			Value: req.Email,
		})
	}

	if req.City != nil {
		updateData = append(updateData, firestore.Update{
			Path:  "city",
			Value: req.City,
		})
	}

	if len(req.FavoritePlace) > 0 {
		updateData = append(updateData, firestore.Update{
			Path:  "favorite_place",
			Value: req.FavoritePlace,
		})
	}

	if req.FavoritePlaceCount != nil {
		updateData = append(updateData, firestore.Update{
			Path:  "favorite_place_count",
			Value: req.FavoritePlaceCount,
		})
	}

	return updateData
}
