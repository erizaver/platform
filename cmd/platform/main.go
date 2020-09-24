package main

import (
	"context"
	app "github.com/EridanSilver/platform/cmd/platform/app"
)

func main() {
	ctx := context.Background()

	application := app.NewApp()
	if err := application.Run(ctx); err != nil {
		panic("panic")
	}
}
