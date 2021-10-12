package routes

import (
	"github.com/aofdev/fiber-versioning-boilerplate/api/handlers"
	"github.com/aofdev/fiber-versioning-boilerplate/api/versions/v2/entities"
	media "github.com/aofdev/fiber-versioning-boilerplate/api/versions/v2/usecases"
	"github.com/gofiber/fiber/v2"
)

// MediaRouter is group routers for media
func MediaRouter(app fiber.Router, mediaUsecase media.MediaUsecase) {
	app.Get("/media/:id", getMediaByID(mediaUsecase))
	app.Post("/media", getMediaByIDS(mediaUsecase))
}

// getMediaByID
// @Summary object media
// @Description get media by id
// @Accept  json
// @Produce  json
// @Param id path int true "Media id"
// @Success 200 {object} handlers.HTTPSuccess{data=entities.Post}
// @Router /v2/media/:id [get]
func getMediaByID(mediaUsecase media.MediaUsecase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		paramID := c.Params("id")

		if paramID == "" {
			return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid ID")
		}

		result, err := mediaUsecase.GetMediaByID(paramID)

		if err != nil {
			_ = c.JSON(&handlers.HTTPError{Status: "error", Error: err.Error()})
		}

		return c.JSON(&handlers.HTTPSuccess{Status: "success", Data: result})
	}
}

// getMediaByIDS
// @Summary Lists item
// @Description get object media
// @Accept  json
// @Produce  json
// @Param ids body entities.RequestFindByListID true "Lists Media id"
// @Success 200 {object} handlers.HTTPSuccess{data=[]entities.Post}
// @Failure 400 {object} handlers.HTTPError
// @Failure 404 {object} handlers.HTTPError
// @Failure 500 {object} handlers.HTTPError
// @Router /v2/media [post]
func getMediaByIDS(mediaUsecase media.MediaUsecase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		mediaBodyRequest := new(entities.RequestFindByListID)
		if err := handlers.ParseBodyAndValidate(c, mediaBodyRequest); err != nil {
			return err
		}

		result, err := mediaUsecase.GetMediaByIDS(mediaBodyRequest.IDS)
		if err != nil {
			_ = c.JSON(&handlers.HTTPError{Status: "error", Error: err.Error()})
		}

		return c.JSON(&handlers.HTTPSuccess{Status: "success", Data: result})
	}
}
