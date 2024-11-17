package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"github.com/ktm-m/playground-go-firebase/adapter/firestore/entity"
	"github.com/ktm-m/playground-go-firebase/internal/model"
	"google.golang.org/api/iterator"
	"time"
)

func (a *adapter) InquiryUsers(ctx context.Context) ([]model.InquiryUsersResp, error) {
	appCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	respData := make([]entity.FireStoreInquiryUsersResp, 0)

	// Get all documents in a collection
	iter := a.client.Collection("users").Documents(appCtx)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			return nil, err
		}

		respData = append(respData, a.mappingDocToFireStoreInquiryUsersResp(doc))
	}

	return entity.MappingCoreFireStoreInquiryUsersResp(respData), nil
}

func (a *adapter) InquiryUserByID(ctx context.Context, req *model.InquiryUserByIDReq) (*model.InquiryUserByIDResp, error) {
	appCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	request := entity.ToFireStoreInquiryUserByIDReq(req)
	respData := new(entity.FireStoreInquiryUserByIDResp)

	// Get a document by ID
	resp, err := a.client.Collection("users").Doc(request.ID).Get(appCtx)
	if err != nil {
		return nil, err
	}

	err = resp.DataTo(respData)
	if err != nil {
		return nil, err
	}
	respData.ID = resp.Ref.ID

	return respData.ToCoreFireStoreInquiryUserByIDResp(), nil
}

func (a *adapter) InquiryUsersByCity(ctx context.Context, req *model.InquiryUsersByCityReq) ([]model.InquiryUsersByCityResp, error) {
	appCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	request := entity.ToFireStoreInquiryUsersByCityReq(req)
	respData := make([]entity.FireStoreInquiryUsersByCityResp, 0)

	// Get all documents in a collection by where clause
	iter := a.client.Collection("users").Where("city", "==", request.City).Documents(appCtx)
	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			return nil, err
		}

		respData = append(respData, a.mappingDocToFireStoreInquiryUsersByCityResp(doc))
	}

	return entity.MappingCoreFireStoreInquiryUsersByCityResp(respData), nil
}

func (a *adapter) mappingDocToFireStoreInquiryUsersResp(doc *firestore.DocumentSnapshot) entity.FireStoreInquiryUsersResp {
	data := doc.Data()
	return entity.FireStoreInquiryUsersResp{
		ID:                 doc.Ref.ID,
		Firstname:          data["firstname"].(string),
		Lastname:           data["lastname"].(string),
		Email:              data["email"].(string),
		City:               data["city"].(string),
		FavoritePlace:      data["favorite_place"].([]string),
		FavoritePlaceCount: data["favorite_place_count"].(int),
	}
}

func (a *adapter) mappingDocToFireStoreInquiryUsersByCityResp(doc *firestore.DocumentSnapshot) entity.FireStoreInquiryUsersByCityResp {
	data := doc.Data()
	return entity.FireStoreInquiryUsersByCityResp{
		ID:                 doc.Ref.ID,
		Firstname:          data["firstname"].(string),
		Lastname:           data["lastname"].(string),
		Email:              data["email"].(string),
		City:               data["city"].(string),
		FavoritePlace:      data["favorite_place"].([]string),
		FavoritePlaceCount: data["favorite_place_count"].(int),
	}
}
