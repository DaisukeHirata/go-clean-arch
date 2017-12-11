package main

import (
	"database/sql"
	"fmt"
	"net/url"

	cfg "github.com/bxcodec/go-clean-arch/config/env"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

var (
	a      App
	config cfg.Config
)

// App is application struct hold router and db connection
type App struct {
	Router *echo.Echo
	DB     *sql.DB
}

func init() {
	config = cfg.NewViperConfig()
	if config.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	a = App{}

	a.initializeDB(
		config.GetString(`database.user`),
		config.GetString(`database.name`),
		config.GetString(`database.pass`),
		config.GetString(`database.host`),
		config.GetString(`database.port`),
	)

	a.initializeRouter()

	a.setupHandler()

	a.run(config.GetString("server.address"))
}

func (a *App) initializeDB(user, dbname, password, host, port string) {

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		user,
		password,
		host,
		port,
		dbname)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Tokyo")
	dsn := fmt.Sprintf("%s?%s", connectionString, val.Encode())

	var err error
	a.DB, err = sql.Open(`mysql`, dsn)
	if err != nil && config.GetBool("debug") {
		fmt.Println(err)
	}
}

func (a *App) run(serveraddr string) {
	a.Router.Start(serveraddr)
}
