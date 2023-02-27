package controllers

import (
	"bytes"
	"encoding/json"
	"hris-gofiber/models"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Authentication godoc
// @Summary Authentication to Get Token
// @Description Get Token with the input payload
// @Tags Token
// @Accept  json
// @Produce  json
// @Param token body models.Credentials true "Get the Token"
// @Success 200 {object} models.TokenResponse
// @Failure 400 {object} models.TokenResponse
// @Router /authorize [post]
func GetTokenHandler(c *fiber.Ctx) error {
	creds := models.Credentials{}
	err := json.Unmarshal(c.Body(), &creds)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid request body")
	}

	url := "https://sister.ulbi.ac.id/ws-sandbox.php/1.0/authorize"
	method := "POST"

	payload, err := json.Marshal(creds)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Internal server error")
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Internal server error")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Cookie", "PHPSESSID=fkaecjg8tso7ivesoghseckhv6")

	res, err := client.Do(req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Internal server error")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Internal server error")
	}

	if res.StatusCode != http.StatusOK {
		return c.Status(res.StatusCode).SendString(string(body))
	}

	var tokenResp models.TokenResponse
	err = json.Unmarshal(body, &tokenResp)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Internal server error")
	}

	token := tokenResp.Token
	return c.JSON(fiber.Map{"token": token})
}
