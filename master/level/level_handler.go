package level

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/vivaldy22/eatnfit-auth-service/tools/respJson"
	"github.com/vivaldy22/eatnfit-auth-service/tools/vError"
	"net/http"
)

func (l *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := l.service.GetAll(context.Background(), new(empty.Empty))
	if err != nil {
		vError.WriteError("Get All Level failed!", err, &w)
	} else {
		respJson.WriteJSON(data, w)
	}
}