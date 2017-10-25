
package ghasedakapi

import(
	"fmt"
	"strings"
)


func arrayToString(i interface{}) string {
	return strings.Trim(strings.Replace(fmt.Sprint(i), " ", ",", -1), "[]")
}
