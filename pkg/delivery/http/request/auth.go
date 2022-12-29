package request

import (
	"encoding/json"
	"github.com/enfil/metamask-auth/internal/domain/user"
	"net/http"
)

func BindReqBody(r *http.Request, obj any) error {
	return json.NewDecoder(r.Body).Decode(obj)
}

func GetUserFromReqContext(r *http.Request) user.Entity {
	ctx := r.Context()
	key := ctx.Value("user").(user.Entity)
	return key
}
