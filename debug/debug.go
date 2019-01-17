package debug

import (
	"encoding/json"
	"fmt"
)

func PrintArray(result interface{}) {
	p, _ := json.MarshalIndent(result, "", " ")
	fmt.Println(string(p))
}
