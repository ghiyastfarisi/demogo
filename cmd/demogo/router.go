package demogo

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	config "github.com/ghiyastfarisi/demogo/configs"
	"github.com/gin-gonic/gin"
)

func router() {
	// set release mode if env is production
	if config.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	// for healtcheck purpose
	r.GET("/healthcheck", healthCheckHandler)

	// enable recovery middleware (recover from panic)
	r.Use(gin.Recovery())
	// enable logger middleware with custom format
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	// enable custom middleware
	// r.Use(defaultMiddleware())

	// Routes `api` group
	// apiGroup := r.Group("/api")
	// Routes `v1` group
	// apiV1Group := apiGroup.Group("/v1")
	// Router `post` group
	// v1PostRouter := apiV1Group.Group("/post")
	// List of post routes
	// v1PostRouter.GET("/:post-slug", post.GetHandler)
	// ...............
	// ...............
	// ...............

	log.Println("\n>>>>> Running Application on Port", config.Element.Application.Port)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Element.Application.Port),
		Handler: r,
	}

	// listening signal
	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, os.Interrupt)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	<-stopSignal
	log.Println("Receiving signal to stop, shutting down service in 2sec")

	time.Sleep(2 * time.Second)

	// timeout shutdown for 2 secs
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Service exited, bye bye...")
}

func healthCheckHandler(c *gin.Context) {
	c.String(http.StatusOK, config.Element.Application.Name+" sehat selalu gaes!")
}
