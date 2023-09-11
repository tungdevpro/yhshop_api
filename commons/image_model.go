package commons

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	Id        int    `json:"id" gorm:"column:id;"`
	Url       string `json:"url"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	CloudName string `json:"cloud_name"`
}

func (Image) TableName() string {
	return "images"
}

func (j *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img
	return nil
}

func (j *Image) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}
