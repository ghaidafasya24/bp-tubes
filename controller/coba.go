package controller

import (
	"errors"
	"fmt"
	"github.com/aiteung/musik"
	inimodel "github.com/ghaidafasya24/be-tubes/model"
	cek "github.com/ghaidafasya24/be-tubes/module"
	"github.com/ghaidafasya24/bp-tubes/config"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetMenu(c *fiber.Ctx) error {
	ps := cek.GetAllMenu(config.Ulbimongoconn, "restoran")
	return c.JSON(ps)
}

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

func InsertDataMenu(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var menurestoran inimodel.Menu
	if err := c.BodyParser(&menurestoran); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := cek.InsertMenu(db, "restoran",
		menurestoran.Nama,
		menurestoran.Harga,
		menurestoran.Deskripsi,
		menurestoran.Kategori,
		menurestoran.BahanBaku)
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
