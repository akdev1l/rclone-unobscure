package main

import (
	"fmt"
	"github.com/rclone/rclone/fs/config"
	"github.com/rclone/rclone/fs/config/obscure"
	"log/slog"
)

func main() {
	for _, section := range config.FileSections() {
		remoteType := config.FileGet(section, "type")
		if remoteType == "crypt" {
			slog.Info("Found crypt remote", "remote", section)
			password := config.FileGet(section, "password")
			password2 := config.FileGet(section, "password2")
			fmt.Printf("[%s]\n", section)
			fmt.Println("password =", obscure.MustReveal(password))
			fmt.Println("password2 =", obscure.MustReveal(password2))
		}
	}
}
