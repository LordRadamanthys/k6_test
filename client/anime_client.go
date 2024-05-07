package client

import (
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/test/model"
)

func GetAnime(id int) (*model.Data, error) {
	client := resty.New()
	anime := model.Data{}
	res, err := client.R().SetResult(&anime).Get(fmt.Sprintf("https://api.jikan.moe/v4/anime/%d", id))

	fmt.Println(res.RawResponse)
	if err != nil {
		panic(err)
	}

	if res.StatusCode() > 299 {
		return nil, errors.New("anime not found")
	}

	return &anime, nil

}
