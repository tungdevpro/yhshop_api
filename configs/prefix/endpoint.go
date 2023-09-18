package prefix

const (
	V1    = "/v1"
	Empty = ""

	// Authentication
	Auth     = "/auth"
	Login    = "/login"
	Register = "/register"

	// User
	User    = "/user"
	Profile = "/profile"

	// Upload
	Upload = "/upload"

	// Shop
	Shop     = "/shop"
	ListShop = "/shops"
	GetShop  = "/:id"
	DelShop  = "/:id"

	// Shipper
	Shipper     = "/shipper"
	ListShipper = "/shipers"
)
