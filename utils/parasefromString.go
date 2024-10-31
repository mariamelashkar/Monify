package utils
import
(
	"fmt"
	"strconv"
	"time"
)


func ParseFromString(value interface{}, targetType string) (interface{}, error) {
	var strVal string
	switch v := value.(type) {
	case string:
		strVal = v
	default:
		strVal = fmt.Sprintf("%v", v)
	}

	switch targetType {
	case "bool":
		return strconv.ParseBool(strVal)
	case "int":
		return strconv.Atoi(strVal)
	case "float64":
		return strconv.ParseFloat(strVal, 64)
	case "int64":
		return strconv.ParseInt(strVal, 10, 64)
	case "uint64":
		return strconv.ParseUint(strVal, 10, 64)
	case "time":
		return time.Parse(time.RFC3339, strVal)
	default:
		return strVal, fmt.Errorf("unsupported target type: %s", targetType)
	}
}
