package user

import (
	"net/http"
	"net/url"

	"github.com/Folderr/Mirror/user/endpoints"
)

func SetupEndpoints(domain *url.URL, mux *http.ServeMux) {

	endpoints.SetupLinks(&instanceOnline, domain, mux)

	CheckStatus(domain)
}
