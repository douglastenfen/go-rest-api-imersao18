package repository

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/douglastenfen/go-rest-api/internal/models"
)

var (
	data     *Data
	once     sync.Once
	jsonFile = "data.json"
)

type Data struct {
	Events []models.Event `json:"events"`
	Spots  []models.Spot  `json:"spots"`
}

func LoadData() (*Data, error) {
	once.Do(func() {
		file, err := os.ReadFile(jsonFile)
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal(file, &data); err != nil {
			panic(err)
		}
	})

	return data, nil
}
