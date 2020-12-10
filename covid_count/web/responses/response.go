package responses

import (
	"net/http"

	"github.com/unrolled/render"
)

var out *render.Render

func init() {
	out = render.New()
}

func WriteText(w http.ResponseWriter, statusCode int, data string) {
	// nolint
	_ = out.Text(w, statusCode, data)
}

// WriteJSON writes a json response
func WriteJSON(w http.ResponseWriter, statusCode int, body interface{}) {
	// nolint
	_ = out.JSON(w, statusCode, body)
}

// WriteJSONError writes a JSON error
func WriteJSONError(w http.ResponseWriter, r *http.Request, err error) {
	resErr, ok := err.(*Error)
	if !ok {
		resErr = UnexpectedError(err.Error())
	}

	WriteJSON(w, resErr.Code, resErr)
}
