package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/davimelovasc/go-simple-api/internal/infra/akafka"
	"github.com/davimelovasc/go-simple-api/internal/infra/repository"
	"github.com/davimelovasc/go-simple-api/internal/infra/web"
	"github.com/davimelovasc/go-simple-api/internal/usecase"
	"github.com/go-chi/chi/v5"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/products")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := repository.NewProductRepositoryMysql(db)
	createProductUseCase := usecase.NewCreateProductUseCase(repository)
	listProductUseCase := usecase.NewListProductsUseCase(repository)

	productHandler := web.NewProductHandler(createProductUseCase, listProductUseCase)

	r := chi.NewRouter()
	r.Post("/products", productHandler.CreateProductHandler)
	r.Get("/products", productHandler.ListProductHandler)

	go http.ListenAndServe(":8000", r)

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"products"}, "host.docker.internal:9092", msgChan)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDto{}
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			println("Error on json parse", err)
		}
		createProductUseCase.Execute(dto)
	}
}
