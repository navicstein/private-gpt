package api

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// FindDocument is the controller for finding a document by id
func (a *API) FindDocument(ctx *fiber.Ctx) error {
	var (
		documentID = ctx.Params("id")
		err        error
	)

	if documentID == "" {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	document, err := a.documentRepository.FindDocumentByID(documentID)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(err)
	}

	return ctx.JSON(document)
}
