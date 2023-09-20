package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type OsType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (j *OsType) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img OsType
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img
	return nil
}

func (j *OsType) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}
