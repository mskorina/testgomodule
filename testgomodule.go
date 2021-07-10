package testgomodule

import "fmt"

func Welcome(name string) string {
	return fmt.Sprintf("Hi, my friend %s", name)
}