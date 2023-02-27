package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"hris-gofiber/models"
	"hris-gofiber/responses"

	"github.com/gofiber/fiber/v2"
)

func GetProfilPribadi(c *fiber.Ctx) error {
	// Mendapatkan id sdm dari parameter URL
	idSdm := c.Params("id_sdm")

	// Membuat HTTP request ke endpoint API SISTER
	req, err := http.NewRequest("GET", fmt.Sprintf("https://sister.ulbi.ac.id/ws-sandbox.php/1.0/data_pribadi/profil/%s", idSdm), nil)
	if err != nil {
		return c.JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed to create HTTP request",
			Data:    nil,
		})
	}

	// Menambahkan header Authorization dengan token Bearer pada HTTP request
	req.Header.Add("Authorization", "Bearer <YOUR-TOKEN-HERE>")

	// Membuat HTTP client untuk mengirim request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "Failed to send HTTP request",
			Data:    nil,
		})
	}

	defer resp.Body.Close()

	// Membaca response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(responses.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed to read response body",
			Data:    nil,
		})
	}

	// Parsing response body ke struct ProfilPribadi
	var profilPribadi models.ProfilPribadi
	if err := json.Unmarshal(respBody, &profilPribadi); err != nil {
		return c.JSON(responses.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed to unmarshal response body",
			Data:    nil,
		})
	}

	// Mengembalikan data profil pribadi dalam bentuk JSON
	return c.JSON(responses.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    profilPribadi,
	})
}
