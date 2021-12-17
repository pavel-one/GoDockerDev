package routes

import (
	"app/base"
)

type Route struct {
	DB     *base.DB
	Router *base.Router
}

func New(db *base.DB, router *base.Router) Route {
	return Route{
		DB:     db,
		Router: router,
	}
}
