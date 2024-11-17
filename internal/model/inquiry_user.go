package model

type InquiryUserByIDReq struct {
	ID string `json:"id"`
}

type InquiryUsersByCityReq struct {
	City string `json:"city"`
}

type InquiryUsersResp struct {
	ID                 string   `json:"id"`
	Firstname          string   `json:"firstname"`
	Lastname           string   `json:"lastname"`
	Age                int      `json:"age"`
	Email              string   `json:"email"`
	City               string   `json:"city"`
	IsCapitalCity      bool     `json:"is_capital_city"`
	FavoritePlace      []string `json:"favorite_place"`
	FavoritePlaceCount int      `json:"favorite_place_count"`
}

type InquiryUserByIDResp struct {
	ID                 string   `json:"id"`
	Firstname          string   `json:"firstname"`
	Lastname           string   `json:"lastname"`
	Age                int      `json:"age"`
	Email              string   `json:"email"`
	City               string   `json:"city"`
	IsCapitalCity      bool     `json:"is_capital_city"`
	FavoritePlace      []string `json:"favorite_place"`
	FavoritePlaceCount int      `json:"favorite_place_count"`
}

type InquiryUsersByCityResp struct {
	ID                 string   `json:"id"`
	Firstname          string   `json:"firstname"`
	Lastname           string   `json:"lastname"`
	Age                int      `json:"age"`
	Email              string   `json:"email"`
	City               string   `json:"city"`
	IsCapitalCity      bool     `json:"is_capital_city"`
	FavoritePlace      []string `json:"favorite_place"`
	FavoritePlaceCount int      `json:"favorite_place_count"`
}

type SummaryUserBatchResp struct {
	UserCount      int     `json:"user_count"`
	SumUserAge     int     `json:"sum_user_age"`
	AverageUserAge float64 `json:"average_user_age"`
}
