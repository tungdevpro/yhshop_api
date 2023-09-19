package commons

const (
	CurrentUser = "current_user"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
