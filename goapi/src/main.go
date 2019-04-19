/*
 * Secret Server
 *
 * This is an API of a secret service. You can save your secret by using the API. You can restrict the access of a secret after the certen number of views or after a certen period of time.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	sw "github.com/olivernadj/secret-server-task/goapi/src/go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	sh := http.StripPrefix("/v1/ui/", http.FileServer(http.Dir("./swaggerui/")))
	router.PathPrefix("/v1/ui/").Handler(sh)

	router.Path("/metrics").Handler(promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8081", router))
}
