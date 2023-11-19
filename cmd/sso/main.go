// `cmd/sso/main.go`
package main

import "github.com/SHALfEY088/sso/internal/config"

func main() {
	cfg := config.MustLoad()

	// ...

	// TODO: инициализировать логгер

	// TODO: инициализировать приложение (app)

	// TODO: запустить gRPC-сервер приложения
}
