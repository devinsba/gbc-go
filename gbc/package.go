package gbc

import "github.com/ventu-io/slf"

func logger() slf.StructuredLogger {
	return slf.WithContext("gbc")
}
