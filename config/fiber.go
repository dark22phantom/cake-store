package config

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Fiber struct {
	Config fiber.Config
}

func (f Fiber) Initialize() {
	f.Config = fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			log.Fatal("[CONFIG] Failed to initialize fiber config | " + err.Error())
			return err
		},
	}
}
