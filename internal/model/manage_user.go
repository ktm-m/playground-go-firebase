package model

type InsertUserReq struct {
	ID                 string   `json:"id"`
	Firstname          string   `json:"firstname"`
	Lastname           string   `json:"lastname"`
	Email              string   `json:"email"`
	City               string   `json:"city"`
	FavoritePlace      []string `json:"favorite_place"`
	FavoritePlaceCount int      `json:"favorite_place_count"`
}

type UpsertUserReq struct {
	ID                 string   `json:"id"`
	Firstname          string   `json:"firstname"`
	Lastname           string   `json:"lastname"`
	Email              string   `json:"email"`
	City               string   `json:"city"`
	FavoritePlace      []string `json:"favorite_place"`
	FavoritePlaceCount int      `json:"favorite_place_count"`
}

type UpdateCapitalCityFlagReq struct {
	ID            string `json:"id"`
	IsCapitalCity bool   `json:"is_capital_city"`
}

type IncrementUpdateUserReq struct {
	ID                 string   `json:"id"`
	Firstname          *string  `json:"firstname"`
	Lastname           *string  `json:"lastname"`
	Email              *string  `json:"email"`
	City               *string  `json:"city"`
	FavoritePlace      []string `json:"favorite_place"`
	FavoritePlaceCount *int     `json:"favorite_place_count"`
}

type AddOneFavoritePlaceReq struct {
	ID            string `json:"id"`
	FavoritePlace string `json:"favorite_place"`
}

type AddMultipleFavoritePlaceReq struct {
	ID            string   `json:"id"`
	FavoritePlace []string `json:"favorite_place"`
}

type DeleteUserByIDReq struct {
	ID string `json:"id"`
}

type DeleteCapitalCityFlagReq struct {
	ID string `json:"id"`
}
