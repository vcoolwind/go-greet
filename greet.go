package go_greet

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("%s, 你好！ --from v1.1.2", name)
}
