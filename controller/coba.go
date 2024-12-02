package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/aiteung/musik"
	inimodel "github.com/ghaidafasya24/be-tubes/model"
	cek "github.com/ghaidafasya24/be-tubes/module"
	"github.com/ghaidafasya24/bp-tubes/config"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

// GetMenu godoc
// @Summary Get All Data Menu.
// @Description Mengambil semua data menu.
// @Tags Menu
// @Accept json
// @Produce json
// @Success 200 {object} Menu
// @Router /restoran [get]
func GetMenu(c *fiber.Ctx) error {
	ps := cek.GetAllMenu(config.Ulbimongoconn, "restoran")
	fmt.Println("Data yang akan dikirim: ", ps) // Tambahkan log ini
	return c.JSON(ps)
}

// GetMenuID godoc
// @Summary Get By ID Data Menu.
// @Description Ambil per ID data menu.
// @Tags Menu
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Menu
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /restoran/{id} [get]
func GetMenuID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := cek.GetMenuFromID(objID, config.Ulbimongoconn, "restoran")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}

// InsertDataMenu godoc
// @Summary Insert data menu.
// @Description Input data menu.
// @Tags Menu
// @Accept json
// @Produce json
// @Param request body ReqMenu true "Payload Body [RAW]"
// @Success 200 {object} Menu
// @Failure 400
// @Failure 500
// @Router /insert [post]
func InsertDataMenu(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var menurestoran inimodel.Menu
	if err := c.BodyParser(&menurestoran); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := cek.InsertMenu(db, "restoran", menurestoran)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// UpdateData godoc
// @Summary Update data menu.
// @Description Ubah data menu.
// @Tags Menu
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body ReqMenu true "Payload Body [RAW]"
// @Success 200 {object} Menu
// @Failure 400
// @Failure 500
// @Router /update/{id} [put]
func UpdateData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a Presensi object
	var menurestoran inimodel.Menu
	if err := c.BodyParser(&menurestoran); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdateMenu function with the parsed ID and the Menu object
	err = cek.UpdateMenu(db, "restoran",
		objectID,
		menurestoran.Nama,
		menurestoran.Harga,
		menurestoran.Deskripsi,
		menurestoran.Gambar,
		menurestoran.Kategori,
		menurestoran.BahanBaku)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

// DeleteMenuByID godoc
// @Summary Delete data menu.
// @Description Hapus data menu.
// @Tags Menu
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delete/{id} [delete]
func DeleteMenuByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}

	err = cek.DeleteMenuByID(objID, config.Ulbimongoconn, "restoran")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}
