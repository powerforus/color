package lib

import "fmt"

func Red(words string) string {
	return fmt.Sprintf("\033[1;31;40m%s\033[0m", words)
}

func Green(words string) string {
	return fmt.Sprintf("\033[1;32;40m%s\033[0m", words)
}

func Yellow(words string) string {
	return fmt.Sprintf("\033[1;33;40m%s\033[0m", words)
}

func Blue(words string) string {
	return fmt.Sprintf("\033[1;34;40m%s\033[0m", words)
}
