package main

import "github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/bootstrap"

func main() {
	if err := bootstrap.Initialize(); err != nil {
		panic(err)
	}
}
