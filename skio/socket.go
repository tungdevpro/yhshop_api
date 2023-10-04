package skio

import (
	"coffee_api/commons"
	"net"
	"net/http"
	"net/url"
)

type Conn interface {
	ID() string
	Close() error
	URL() url.URL
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	RemoteHeader() http.Header

	Context() interface{}
	SetContext(v interface{})
	Namespace() string
	Emit(msg string, v ...interface{})

	Join(room string)
	Leave(room string)
	LeaveAll()
}

type AppSocket interface {
	Conn
	commons.Requester
}

type appSocket struct {
	Conn
	commons.Requester
}

func NewAppSocket(conn Conn, requester commons.Requester) *appSocket {
	return &appSocket{
		Conn:      conn,
		Requester: requester,
	}
}
