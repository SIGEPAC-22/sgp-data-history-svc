package bootstrap

import (
	"database/sql"
	"fmt"
	"github.com/dimiro1/health"
	kitlog "github.com/go-kit/log"
	_ "github.com/go-sql-driver/mysql"
	goconfig "github.com/iglin/go-config"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sgp-data-history-svc/internal/getHistorical/getHistoricalService"
	"sgp-data-history-svc/internal/getHistorical/platform/handler"
	"sgp-data-history-svc/internal/getHistorical/platform/storage/mysql"
	"sgp-data-history-svc/internal/getOneHistorical/getOneHistoricalService"
	handler2 "sgp-data-history-svc/internal/getOneHistorical/platform/handler"
	mysql2 "sgp-data-history-svc/internal/getOneHistorical/platform/storage/mysql"
	"syscall"
)

func Run() {
	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	port := config.GetString("server.port")

	var kitlogger kitlog.Logger
	kitlogger = kitlog.NewJSONLogger(os.Stderr)
	kitlogger = kitlog.With(kitlogger, "time", kitlog.DefaultTimestamp)

	mux := http.NewServeMux()
	errs := make(chan error, 2)

	////////////////////////////////////////////////////////////////////////
	////////////////////////CORS///////////////////////////////////////////
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodPut,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	handlerCORS := cors.Handler(mux)
	////////////////////////CORS///////////////////////////////////////////

	db, err := sql.Open("mysql", getStrConnection())
	if err != nil {
		log.Fatalf("unable to open database connection %s", err.Error())
	}

	//////////////////////GET HISTORICAL////////////////////////////////////////////////
	getHistoricalRepo := mysql.NewGetHistoricalRepo(db, kitlogger)
	getHistoricalService := getHistoricalService.NewService(getHistoricalRepo, kitlogger)
	getHistoricalEndpoint := handler.MakeGetHistoricalEndpoints(getHistoricalService)
	getHistoricalEndpoint = handler.GetHistoricalTransportMiddleware(kitlogger)(getHistoricalEndpoint)
	getHistoricalHandler := handler.NewHttpGetHistoricalHandler(config.GetString("paths.getHistorical"), getHistoricalEndpoint)
	//////////////////////GET HISTORICAL////////////////////////////////////////////////

	//////////////////////GET ONE HISTORICAL////////////////////////////////////////////////
	getOneHistoricalRepo := mysql2.NewGetOneHistoricalRepo(db, kitlogger)
	getOneHistoricalService := getOneHistoricalService.NewService(getOneHistoricalRepo, kitlogger)
	getOneHistoricalEndpoint := handler2.MakeGetOneHistoricalEndpoints(getOneHistoricalService)
	getHistoricalEndpoint = handler2.GetOneHistoricalTransportMiddleware(kitlogger)(getOneHistoricalEndpoint)
	getOneHistoricalHandler := handler2.NewHttpGetOneHistoricalHandler(config.GetString("paths.getOneHistorical"), getHistoricalEndpoint)
	//////////////////////GET ONE HISTORICAL////////////////////////////////////////////////

	mux.Handle(config.GetString("paths.getHistorical"), getHistoricalHandler)
	mux.Handle(config.GetString("paths.getOneHistorical"), getOneHistoricalHandler)
	mux.Handle("/health", health.NewHandler())

	go func() {
		kitlogger.Log("listening", "transport", "http", "address", port)
		errs <- http.ListenAndServe(":"+port, handlerCORS)
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		signal.Notify(c, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
		db.Close()
	}()
	kitlogger.Log("terminated", <-errs)
}

func getStrConnection() string {
	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	host := config.GetString("datasource.host")
	user := config.GetString("datasource.user")
	pass := config.GetString("datasource.pass")
	dbname := config.GetString("datasource.dbname")
	strconn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True", user, pass, host, dbname)
	return strconn
}
