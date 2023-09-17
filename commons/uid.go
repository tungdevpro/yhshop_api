package commons

import (
	"coffee_api/helpers"
	"fmt"
	"math/rand"
	"time"
)

type UID struct {
	LocalId string `json:"id"`
}

func NewUID(id string) *UID {
	return &UID{
		LocalId: id,
	}
}

func (u *UID) GenerateID() UID {
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

	res := fmt.Sprintf("%d%s%s%v", year, m, string(randomString), u.LocalId)

	uuid := UID{
		LocalId: res,
	}

	return uuid
}

func (uid *UID) String() string {
	return fmt.Sprintf("%v", uid.LocalId)
}

func (uid UID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", uid.String())), nil
}
