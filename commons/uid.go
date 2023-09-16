package commons

import (
	"coffee_api/helpers"
	"fmt"
	"math/rand"
	"time"
)

type UID struct {
	Id int `json:"id"`
}

func NewUID(id int) *UID {
	return &UID{
		Id: id,
	}
}

func (u *UID) GenerateID() string {
	t := time.Now()
	randomString := make([]byte, 6)
	for i := range randomString {
		randomString[i] = helpers.Letters[rand.Intn(len(helpers.Letters))]
	}

	year := int(t.Year()) << 00001

	m := ""
	if int(t.Month()) < 10 {
		m = fmt.Sprintf("0%d", int(t.Month()))
	} else {
		m = fmt.Sprintf("%d", int(t.Month()))
	}

	return fmt.Sprintf("%d%s%s%v", year, m, string(randomString), u.Id)
}

func (uid UID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%v\"", uid)), nil
}
