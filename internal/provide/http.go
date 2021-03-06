package provide

import (
	"github.com/vseinstrumentiru/lego/v2/transport/http/httpclient"
	"github.com/vseinstrumentiru/lego/v2/transport/http/httpserver"
)

func Http() []interface{} {
	return []interface{}{
		httpserver.Provide,
		httpclient.Provide,
		httpclient.ConstructorProvider,
	}
}
