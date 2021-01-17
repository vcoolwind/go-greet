package greet

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("%s, 你好！ --from v1.0.2", name)
}
