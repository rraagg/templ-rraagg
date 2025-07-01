package main

import (
	"context"
	"os"
)

func main() {
	component := hello("Rob")
	component.Render(context.Background(), os.Stdout)
}
