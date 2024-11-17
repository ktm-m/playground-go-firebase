package outbound

import (
	"context"
	"github.com/ktm-m/playground-go-firebase/internal/model"
)

type FireStorePort interface {
	InquiryUsers(ctx context.Context) ([]model.InquiryUsersResp, error)
	InquiryUserByID(ctx context.Context, req *model.InquiryUserByIDReq) (*model.InquiryUserByIDResp, error)
	InquiryUsersByCity(ctx context.Context, req *model.InquiryUsersByCityReq) ([]model.InquiryUsersByCityResp, error)
	InsertUser(ctx context.Context, req *model.InsertUserReq) error
	UpsertUser(ctx context.Context, req *model.UpsertUserReq) error
	UpdateCapitalCityFlag(ctx context.Context, req *model.UpdateCapitalCityFlagReq) error
	AddOneFavoritePlace(ctx context.Context, req *model.AddOneFavoritePlaceReq) error
	DeleteUserByID(ctx context.Context, req *model.DeleteUserByIDReq) error
	DeleteCapitalCityFlag(ctx context.Context, req *model.DeleteCapitalCityFlagReq) error
}
