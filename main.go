package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"pawn/src/gadai"
	"sync"
	"syscall"

	gadais "pawn/src/gadai/repository"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func main() {

	ctx := context.Background()

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = &serializedLogger{Logger: logger}
	logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)

	err := godotenv.Load()
	if err != nil {
		logger.Log("Error loading .env file")
	}

	//db, err := gorm.Open("sqlite3", "test.db")
	db, err := gorm.Open("mysql", "root:mysql@/gadai2?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&gadais.Taksir{})

	env := os.Getenv("ENV")
	if env == "DEV" {
		db.DropTable(&gadais.Taksir{})
		db.AutoMigrate(&gadais.Taksir{})
		var ss gadais.SeedService
		ss = gadais.Config(db)
		ss.SeedAll()
	}

	var (
		addr     = os.Getenv("PORT")
		httpAddr = flag.String("http.addr", ":"+addr, "HTTP listen address")
	)

	flag.Parse()

	fieldKeys := []string{"method"}

	var taksirs gadais.Repository
	taksirs = gadais.Init(db)

	var gs gadai.GadaiService
	gs = gadai.NewService(taksirs)
	gs = gadai.NewLoggingService(log.NewContext(logger).With("component", "gadai"), gs)
	gs = gadai.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "gadai_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "gadai_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		gs,
	)

	mux := http.NewServeMux()
	mux.Handle("/gadai/v2/", gadai.MakeHandler(ctx, gs))

	mux.Handle("/", accessControl(mux))
	mux.Handle("/metrics", stdprometheus.Handler())

	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddr, mux)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)

}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

type serializedLogger struct {
	mtx sync.Mutex
	log.Logger
}

func (l *serializedLogger) Log(keyvals ...interface{}) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	return l.Logger.Log(keyvals...)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		//h.ServeHTTP(w, r)
	})
}
