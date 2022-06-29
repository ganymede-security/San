package netlog_test

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ganymede-security/san/server/middleware/netlog"
)

//func TestNewProxyHandler(t *testing.T) {
//	log.New()

//	netlog.NewProxyHandler
//}

func TestHandler(t *testing.T) {
	var logger netlog.ProxyLogger

	netlog.NewProxyHandler(logger, )
}
