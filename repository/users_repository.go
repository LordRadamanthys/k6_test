package repository

import (
	"context"

	"github.com/test/model"
	"github.com/test/utils/config"
)

func SaveAnime(anime model.Anime) error {
	_, err := config.MongoCollection.InsertOne(context.TODO(), anime)

	if err != nil {
		return err
	}
	return nil
}
