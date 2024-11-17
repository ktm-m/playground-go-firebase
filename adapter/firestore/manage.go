package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"fmt"
	"github.com/ktm-m/playground-go-firebase/adapter/firestore/entity"
	"github.com/ktm-m/playground-go-firebase/internal/model"
	"time"
)

func (a *adapter) InsertUser(ctx context.Context, req *model.InsertUserReq) error {
	appCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	ID, data := entity.ToFireStoreInsertUserReq(req)

	// Insert a document, if not specify ID, Firestore will generate random ID (no need .Doc(ID))
	_, err := a.client.Collection("users").Doc(ID).Set(appCtx, data)
	if err != nil {
		return err
	}

	// Another way to insert a document
	ref := a.client.Collection("users").Doc(a.buildBackupID(ID))
	_, err = ref.Set(appCtx, data)
	if err != nil {
		return err
	}

	return nil
}

func (a *adapter) UpsertUser(ctx context.Context, req *model.UpsertUserReq) error {
	appCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	ID, data := entity.ToFireStoreUpsertUserReq(req)

	// Insert or Update a document
	_, err := a.client.Collection("users").Doc(ID).Set(appCtx, data, firestore.MergeAll)
	if err != nil {
		return err
	}

	return nil
}

func (a *adapter) UpdateCapitalCityFlag(ctx context.Context, req *model.UpdateCapitalCityFlagReq) error {
	appCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	ID, data := entity.ToFireStoreUpdateCapitalCityFlagReq(req)

	// Increment update a document by append new field
	_, err := a.client.Collection("users").Doc(ID).Set(appCtx, data, firestore.MergeAll)
	if err != nil {
		return err
	}

	return nil
}

func (a *adapter) IncrementUpdateUser(ctx context.Context, req *model.IncrementUpdateUserReq) error {
	appCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	ID, data := entity.ToFireStoreIncrementUpdateUserReq(req)

	// Increment update a document
	_, err := a.client.Collection("users").Doc(ID).Update(appCtx, data.BuildIncrementalUpdateData())
	if err != nil {
		return err
	}

	return nil
}

func (a *adapter) AddOneFavoritePlace(ctx context.Context, req *model.AddOneFavoritePlaceReq) error {
	appCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	ID, data := entity.ToFireStoreAddOneFavoritePlaceReq(req)

	// Add an element to array and increment count by 1
	_, err := a.client.Collection("users").Doc(ID).Update(appCtx, []firestore.Update{
		{
			Path:  "favorite_place",
			Value: firestore.ArrayUnion(data.FavoritePlace), // Add an element to array
		},
		{
			Path:  "favorite_place_count",
			Value: firestore.Increment(1), // Increment count by 1
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *adapter) AddMultipleFavoritePlace(ctx context.Context, req *model.AddMultipleFavoritePlaceReq) error {
	appCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	ID, data := entity.ToFireStoreAddMultipleFavoritePlaceReq(req)
	ref := a.client.Collection("users").Doc(ID)

	// Add an element to array and increment count by the number of elements using transaction
	err := a.client.RunTransaction(appCtx, func(ctx context.Context, tx *firestore.Transaction) error {
		doc, err := tx.Get(ref)
		if err != nil {
			return err
		}

		err = tx.Update(ref, []firestore.Update{
			{
				Path:  "favorite_place",
				Value: firestore.ArrayUnion(data.FavoritePlace),
			},
		})
		if err != nil {
			return err
		}

		count, err := doc.DataAt("favorite_place_count")
		if err != nil {
			return err
		}
		updateCount := count.(int) + len(data.FavoritePlace)
		if updateCount > 3 {
			return errors.New("favorite place count exceeds the limit")
		}

		err = tx.Update(ref, []firestore.Update{
			{
				Path:  "favorite_place_count",
				Value: updateCount,
			},
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *adapter) DeleteUserByID(ctx context.Context, req *model.DeleteUserByIDReq) error {
	appCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	request := entity.ToFireStoreDeleteUserByIDReq(req)

	// Delete a document by ID
	_, err := a.client.Collection("users").Doc(request.ID).Delete(appCtx)
	if err != nil {
		return err
	}

	return nil
}

func (a *adapter) DeleteCapitalCityFlag(ctx context.Context, req *model.DeleteCapitalCityFlagReq) error {
	appCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	request := entity.ToFireStoreDeleteCapitalCityFlagReq(req)

	// Delete a field in a document
	_, err := a.client.Collection("users").Doc(request.ID).Update(appCtx, []firestore.Update{
		{
			Path:  "is_capital_city",
			Value: firestore.Delete,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *adapter) buildBackupID(ID string) string {
	return fmt.Sprintf("%s_%s_%s", ID, time.Now().Format("20060102"), "backup")
}
