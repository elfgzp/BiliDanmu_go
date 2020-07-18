package bilive

import "net/http"


var (
	// Client 客户端
	Client *http.Client
)

func init() {
	Client = &http.Client{}
}
