package shared
import (
	"github.com/gin-gonic/gin"
	"github.com/twodigitss/reserv-go/configs" //env vars
)

// Returns a json object as response to display in handler functions depending on the env variables.
// (pretty printing for dev, json for prod).
func JSON(c *gin.Context, code int, obj any) {
    if configs.ENV == "development" {
        c.IndentedJSON(code, obj)
    } else {
        c.JSON(code, obj)
    }
}
