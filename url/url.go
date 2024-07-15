package url

import (
	"github.com/ghaidafasya24/bp-tubes/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger" // swagger handler
)

func Web(page *fiber.App) {
	// page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	// page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth

	page.Get("/", controller.Sink)
	page.Post("/", controller.Sink)
	page.Put("/", controller.Sink)
	page.Patch("/", controller.Sink)
	page.Delete("/", controller.Sink)
	page.Options("/", controller.Sink)

	page.Get("/checkip", controller.Homepage) //ujicoba panggil package musik
	page.Get("/restoran", controller.GetMenu)
	page.Get("/restoran/:id", controller.GetMenuID) //menampilkan data menu berdasarkan id
	page.Post("/insert", controller.InsertDataMenu)
	page.Put("/update/:id", controller.UpdateData)
	page.Delete("/delete/:id", controller.DeleteMenuByID)

	page.Get("/docs/*", swagger.HandlerDefault)
}
