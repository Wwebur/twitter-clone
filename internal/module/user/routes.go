package user

import (
	"github.com/HotPotatoC/twitter-clone/internal/module/user/action"
	"github.com/HotPotatoC/twitter-clone/internal/module/user/service"
	"github.com/HotPotatoC/twitter-clone/internal/server/middleware"
	"github.com/HotPotatoC/twitter-clone/pkg/aws"
	"github.com/HotPotatoC/twitter-clone/pkg/cache"
	"github.com/HotPotatoC/twitter-clone/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func Routes(r fiber.Router, db database.Database, s3 *aws.S3Bucket, cache cache.Cache) {
	authMiddleware := middleware.NewAuthMiddleware()
	r.Post("/register", buildRegisterHandler(db))
	r.Get("/:username", authMiddleware.Execute(), buildGetUserHandler(db))
	r.Get("/:username/tweets", authMiddleware.Execute(), buildGetUserTweetsHandler(db))
	r.Patch("/profile", authMiddleware.Execute(), buildUpdateUserHandler(db))
	r.Patch("/profile_image", authMiddleware.Execute(), buildUpdateProfileImageHandler(db, s3))
}

func buildRegisterHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewRegisterService(db)
		action := action.NewRegisterAction(service)

		return action.Execute(c)
	}
}

func buildGetUserHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewGetUserService(db)
		action := action.NewGetUserAction(service)

		return action.Execute(c)
	}
}

func buildGetUserTweetsHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewGetUserTweetsService(db)
		action := action.NewGetUserTweetsAction(service)

		return action.Execute(c)
	}
}

func buildUpdateUserHandler(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewUpdateUserService(db)
		action := action.NewUpdateUserAction(service)

		return action.Execute(c)
	}
}

func buildUpdateProfileImageHandler(db database.Database, s3 *aws.S3Bucket) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := service.NewUpdateProfileImageService(db, s3)
		action := action.NewUpdateProfileImageAction(service)

		return action.Execute(c)
	}
}
