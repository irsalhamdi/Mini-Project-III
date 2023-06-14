package main

import (
	"customer-relationship-management/entity"
	"customer-relationship-management/modules/actor"
	"customer-relationship-management/modules/customer"
	db2 "customer-relationship-management/utils/db"
	"fmt"
	"path"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	//main function
	_ = godotenv.Load()
	db := db2.GormMysql()

	migrationSource := &migrate.FileMigrationSource{
		Dir: path.Join("."),
	}
	sqlDb, _ := db.DB()
	n, err := migrate.Exec(sqlDb, "mysql", migrationSource, migrate.Up)

	fmt.Println(n, err)

	router := gin.New()

	router.Use(cors.Default())
	router.Use(helmet.Default())
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Minute * 1,
		Limit: 20,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: entity.ErrorHandler,
		KeyFunc:      entity.KeyFunc,
	})
	router.Use(mw)
	actorHandler := actor.NewRouter(db)
	actorHandler.Handle(router)

	customerHandler := customer.NewRouter(db)
	customerHandler.Handle(router)

	errRouter := router.Run(":8081")
	if errRouter != nil {
		fmt.Println("error running server", errRouter)
		return
	}
}
