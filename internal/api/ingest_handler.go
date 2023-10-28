package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"navicstein/private-gpt/internal/api/service"
	"net/http"
	"os"
)

type IngestURLParams struct {
	URL          string         `json:"url"`
	CollectionID string         `json:"collectionId"`
	Meta         map[string]any `json:"meta"`
}

//func (a *API) QueryData(ctx *fiber.Ctx) error {
//	var (
//		params service.QueryIngestParams
//	)
//	if err := ctx.BodyParser(&params); err != nil {
//		return ctx.Status(http.StatusBadRequest).JSON(err)
//	}
//
//	response, err := a.service.QueryStore(user, params)
//	if err != nil {
//		return ctx.Status(http.StatusBadRequest).JSON(err)
//	}
//	return ctx.JSON(response)
//}

// IngestData uploads the file and creates a document from it
func (a *API) IngestData(ctx *fiber.Ctx) error {
	var (
		params service.IngestDataParam
	)

	files, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	file := files.File["file"][0]

	// parse the body
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}

	// temp directory to save the file
	uploadDir := fmt.Sprintf("%s/vidpress-uploads", os.TempDir())
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		_ = os.MkdirAll(uploadDir, 0755)
	}

	savePath := fmt.Sprintf("%s/%s", uploadDir, file.Filename)

	// save the file
	err = ctx.SaveFile(file, savePath)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	_, err = a.service.IngestFiles(ctx.Context(), service.IngestDataParam{
		CollectionID:  params.CollectionID,
		Meta:          params.Meta,
		SavedFilePath: savePath,
	})

	if err != nil {
		//return ctx.Status(http.StatusInternalServerError).JSON(errs.ErrorHTTP{
		//	Error: err,
		//	Code:  errs.UnknownError,
		//})
		log.Err(err).Msg("unable to ingest file")
		return err
	}

	log.Info().Msg("ingested successfully")
	return ctx.JSON("successfully ingested, please return to UI")
}
