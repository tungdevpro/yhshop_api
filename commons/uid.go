package commons

type UID struct {
	LocalId string `json:"id"`
}

func NewUID(id string) *UID {
	return &UID{
		LocalId: id,
	}
}
