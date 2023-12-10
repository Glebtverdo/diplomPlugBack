package restTransport

import (
	"fmt"
	"net/http"
)

// func encodeJson(body io.ReadCloser, obj *models.AnyObj) error {
// 	dec := json.NewDecoder(body)
// 	dec.DisallowUnknownFields()
// 	return dec.Decode(&obj)
// }

func errorHandler(e error, w http.ResponseWriter) {
	fmt.Println(e)
	w.WriteHeader(400)
	w.Write([]byte(e.Error()))
}
