package main

import (
    "log"

    "github.com/rwnkl/go-package-inspector/internal/app"
)

func main() {
    if err := app.Run(); err != nil {
        log.Fatal(err)
    }
}