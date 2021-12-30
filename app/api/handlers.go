package api

import (
	"github.com/gofiber/fiber/v2"
	"pg_slowest/app/domain"
	"pg_slowest/models"
	"strconv"
)

type handler struct {
	app    *fiber.App
	domain domain.Stater
}

func NewHandler(app *fiber.App, d domain.Stater) *handler {
	return &handler{domain: d, app: app}
}

func Register(h *handler) *fiber.App {
	h.app.Group("/api/v1").
		Get("/stats", h.Stats)

	return h.app
}

func (h *handler) Stats(ctx *fiber.Ctx) error{
	l := ctx.Query("limit", "10") // limit
	offs := ctx.Query("offset", "0")
	filter := ctx.Query("filter", "")

	var (
		limit, offset int
		err           error
	)

	limit, err = strconv.Atoi(l)
	if err != nil {
		limit = 10
	}

	offset, err = strconv.Atoi(offs)
	if err != nil {
		offset = 0
	}

	resp, err := h.domain.Stat(ctx.Context(), &models.APIRequest{
		Filter: filter,
		Pagination: &models.PaginationRequest{
			Limit:  limit,
			Offset: offset,
		},
	})

	if err != nil {
		return ctx.SendString("error:" + err.Error())

	}

	return ctx.SendString(resp.Query)
}
