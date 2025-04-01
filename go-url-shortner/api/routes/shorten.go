package routes

import (
	"time"
    "github.com/shusin2798/Golang/go-url-shortner/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/asaskevich/govalidator"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Cannot parse JSON",
		})
	}

	//implement rate limiting
	//check if input is an actual URL
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid URL",
		})
	}
	//check for domain error
	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error":   true,
			"message": "Domain not allowed",
		})
	}
	//enforce http/ssl
	body.URL = helpers.EnforceHTTP(body.URL)

	//return c.Status(fiber.StatusOK).JSON(fiber.Map{
	//	"error":   false,
	//	"message": "URL shortened successfully",
	//	"data": response{
	//		URL:         body.URL,
	//		CustomShort: body.CustomShort,
	//		Expiry:      body.Expiry,
	//		XRateRemaining: 0,
	//		XRateLimitReset: 0,
	//	},
	//})
	
}
