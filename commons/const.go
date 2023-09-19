package commons

const (
	CurrentUser = "current_user"
)

type Requester interface {
	GetUserId() string
	GetEmail() string
	GetRole() string
}
