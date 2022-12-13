package daytwelve

import (
	"fmt"
	"os"
)

func main() {
	raw, err := os.ReadFile("./test.input")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(raw))
}
