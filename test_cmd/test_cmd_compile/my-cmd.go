package main

import (
	"github.com/micro/go-micro/agent/command"
)

func Ping() command.Command {
	usage := "yang"
	description := "hello wilker!!"
	return command.NewCommand("ping", usage, description, func(args ...string) ([]byte, error) {
		return []byte("Returns xuan 666"), nil
	})
}

func init() {
	command.Commands["^yang$"] = Ping()
}
