package utils

import "fmt"

func Info(msg string) {
	fmt.Println("📦", msg)
}

func Success(msg string) {
	fmt.Println("✅", msg)
}

func Error(msg string) {
	fmt.Println("⛔", msg)
}

func Step(msg string) {
	fmt.Println("🔧", msg)
}
