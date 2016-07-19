package main

import (
	"html/template"

	"./api"
	"./persistence"

	"github.com/iris-contrib/middleware/logger"
	"github.com/iris-contrib/template/pug"
	"github.com/kataras/iris"
)

// var log = logging.MustGetLogger("adventurous_mind")

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
// var format = logging.MustStringFormatter(
// `%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
// )

// Password is just an example type implementing the Redactor interface. Any
// time this is logged, the Redacted() function will be called.
// type Password string

// func (p Password) Redacted() interface{} {
// return logging.Redact(string(p))
// }

func main() {
	// backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	// backend2 := logging.NewLogBackend(os.Stderr, "", 0)
	// backend2Formatter := logging.NewBackendFormatter(backend2, format)
	// backend1Leveled := logging.AddModuleLevel(backend1)
	// backend1Leveled.SetLevel(logging.ERROR, "")
	// logging.SetBackend(backend1Leveled, backend2Formatter)

	cfg := pug.DefaultConfig()
	cfg.Funcs["bold"] = func(content string) (template.HTML, error) {
		return template.HTML("<b>" + content + "</b>"), nil
	}

	iris.UseTemplate(pug.New(cfg)).Directory("./templates", ".pug")

	iris.Favicon("./assets/images/favicon.ico")
	iris.Static("/assets", "./assets", 1)
	iris.Use(logger.New(iris.Logger))
	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.Render("errors/404.jade", iris.Map{"Title": iris.StatusText(iris.StatusNotFound)})
		iris.Logger.Warningf("%s %s not found", ctx.Method(), ctx.RequestURI())
	})

	iris.OnError(iris.StatusInternalServerError, func(ctx *iris.Context) {
		ctx.Render("errors/500.jade", nil)
		iris.Logger.Dangerf("%s %s internal server error", ctx.Method(), ctx.RequestURI())
	})

	// registerRoutes()
	registerAPI()
	iris.Listen("127.0.0.1:8000")
}

// func registerRoutes() {
// iris.Handle("GET", "/", routes.Index())
// iris.Get("/about", routes.About)
// iris.Get("/profile/:username", routes.Profile)("user-profile")
// iris.Get("/all", routes.UserList)
// }

func registerAPI() {
	iris.API("/users", api.UserAPI{})
	iris.API("/session", api.SessionAPI{Db: persistence.New()})
}
