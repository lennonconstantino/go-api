package cache

import (
	"encoding/json"
	"fmt"
	model "go-api/internal/core/domain"
)

type handleProducts func() ([]model.Product, error)

//type handleProduct func() (model.Product, error)

func Cache(key string, f handleProducts) ([]byte, error) {
	reply, err := Get(key)

	if err != nil {
		fmt.Println("going to db")
		objects, err := f()
		if err != nil {
			return nil, err
		}
		productBytes, _ := json.Marshal(objects)
		Set(key, productBytes)
		return productBytes, nil
	}

	fmt.Println("searching in redis")
	return reply, nil
}
