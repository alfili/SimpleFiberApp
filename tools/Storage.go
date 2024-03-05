package tools

import "github.com/gofiber/fiber/v2/middleware/session"

var Store Storage

type Storage struct {
	Sessions *session.Store
}
