
package ghasedakapi

import(
	"fmt"
)


func arrayToString(i interface{}) string {
	return strings.Trim(strings.Replace(fmt.Sprint(i), " ", ",", -1), "[]")
}
