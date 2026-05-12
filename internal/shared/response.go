package shared
import (
	"reflect"
	"github.com/gin-gonic/gin"
	"github.com/twodigitss/reserv-go/configs" //env vars
)

type response[T any] struct {
	Error string `json:"error"`
	Data  []T   	 `json:"data"`
}

// Returns a json object as response to display in handler functions depending on the env variables.
// development = pretty printed json. 
// production = unformatted json.
func JSON(g *gin.Context, _code int, _body any, _err error) {
	var finalData []any

	if _body == nil {
		finalData = []any{}
	} else {
		val := reflect.ValueOf(_body)
		if val.Kind() == reflect.Slice {
			finalData = make([]any, val.Len())
			for i := range val.Len() {
				finalData[i] = val.Index(i).Interface()
			}
		} else {
			finalData = []any{_body}
		}
	}

	var errMsg string = ""
	if _err != nil {
		errMsg = _err.Error()
	}

	switch configs.ENV {
	case "development":
		g.IndentedJSON(_code, response[any]{
			Data:  finalData,
			Error: errMsg,
		})

	default:
		g.JSON(_code, response[any]{
			Data:  finalData,
			Error: errMsg,
		})

	}
}

