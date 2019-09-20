package main

import (
	"encoding/json"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/fabiorubim/eng-zap-challenge-golang/models"
	"github.com/valyala/fasthttp"
	"log"
)

var properties *models.Properties

func Zap(ctx *fasthttp.RequestCtx) {
	json, err := json.Marshal(properties.GetZap())
	if err != nil {
		fmt.Println(err)
	}
	ctx.Response.SetBody(json)
}

func VivaReal(ctx *fasthttp.RequestCtx) {
	json, err := json.Marshal(properties.GetVivaReal())
	if err != nil {
		fmt.Println(err)
	}
	ctx.Response.SetBody(json)

}

func main() {
	router := fasthttprouter.New()
	router.GET("/zap", Zap);
	router.GET("/vivareal", VivaReal);

	properties = models.LoadProperties()

	log.Fatal(fasthttp.ListenAndServe(":8000", router.Handler))
}
