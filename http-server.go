package main

import (
	"encoding/json"
	"fmt"
	pb "github.com/GoogleCloudPlatform/microservices-demo/src/productcatalogservice/genproto/hipstershop"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"net/http"
	"time"
)

func runHttpServer(port string) {
	http.Handle("/listproducts", otelhttp.NewHandler(http.HandlerFunc(ListProducts), "ListProducts"))

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatalf("failed to http serve: %v", err)
	}
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	log.Info("[ListProducts] received request")
	defer log.Info("[ListProducts] completed request")

	time.Sleep(extraLatency)

	var lpr ListProductsResponse
	var productsGrpc []*pb.Product = parseCatalog()
	var products []*Product = parseProduct(productsGrpc)

	lpr.Products = products

	json, err := json.Marshal(lpr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write(json)
}

func parseProduct(productsGrpc []*pb.Product) []*Product {
	var products []*Product

	for _, p := range productsGrpc {
		products = append(products, &Product{
			Id:          p.Id,
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			PriceUsd:    parseMoney(p.PriceUsd),
			Categories:  p.Categories,
		})
	}
	return products
}

func parseMoney(moneyGrpc *pb.Money) *Money {
	return &Money{
		CurrencyCode: moneyGrpc.CurrencyCode,
		Units:        moneyGrpc.Units,
		Nanos:        moneyGrpc.Nanos,
	}
}
