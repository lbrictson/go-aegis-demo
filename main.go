package main

import (
	"bytes"
	"context"
	"html/template"
	"log"
	"net/url"

	aegis "github.com/tmaiaroto/aegis/framework"
)

func main() {
	// Handle an APIGatewayProxyRequest event with a URL reqeust path Router
	router := aegis.NewRouter(fallThrough)
	router.Handle("GET", "/", root, helloMiddleware)
	router.Handle("GET", "/say/:message", rootWithParams, helloMiddleware)

	// Use an Aegis interface to inject optional dependencies into handlers
	// and start listening for events.
	app := aegis.New(aegis.Handlers{
		Router: router,
	})
	app.Start()
}

// fallThrough handles any path that couldn't be matched to another handler
func fallThrough(ctx context.Context, d *aegis.HandlerDependencies, req *aegis.APIGatewayProxyRequest, res *aegis.APIGatewayProxyResponse, params url.Values) error {
	res.StatusCode = 404
	return nil
}

// root is handling GET "/" in this case
func root(ctx context.Context, d *aegis.HandlerDependencies, req *aegis.APIGatewayProxyRequest, res *aegis.APIGatewayProxyResponse, params url.Values) error {
	// lc, _ := lambdacontext.FromContext(ctx)
	// res.JSON(200, map[string]interface{}{"event": req, "context": lc})

	type data struct {
		Message string
	}
	mes := data{Message: "working"}
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println(err)
		return nil
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, mes); err != nil {
		log.Println(err)
		return nil
	}

	result := tpl.String()

	res.HTML(200, result)
	return nil
}

// helloMiddleware is a simple example of middleware
func helloMiddleware(ctx context.Context, d *aegis.HandlerDependencies, req *aegis.APIGatewayProxyRequest, res *aegis.APIGatewayProxyResponse, params url.Values) bool {
	log.Println("Hello CloudWatch!")
	return true
}

// root is handling GET "/" in this case
func rootWithParams(ctx context.Context, d *aegis.HandlerDependencies, req *aegis.APIGatewayProxyRequest, res *aegis.APIGatewayProxyResponse, params url.Values) error {
	// lc, _ := lambdacontext.FromContext(ctx)
	// res.JSON(200, map[string]interface{}{"event": req, "context": lc})

	type data struct {
		Message string
	}
	para := params.Get("message")
	mes := data{Message: para}
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println(err)
		return nil
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, mes); err != nil {
		log.Println(err)
		return nil
	}

	result := tpl.String()

	res.HTML(200, result)
	return nil
}
