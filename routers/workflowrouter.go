package routers

import (
	"backend-test/controllers"
	"backend-test/repositories"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var controller = &controllers.Controller{Repo: &repositories.MockRepository{}}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{
		"[GET]: /workflow - List all workflows",
		"GET",
		"/workflow",
		controller.ListWorkflows,
	},

	Route{
		"[POST]: /workflow - Insert a workflow on database and on queue and respond request with the inserted workflow",
		"POST",
		"/workflow",
		controller.AddWorkflow,
	},

	Route{
		"[PATCH]: /workflow/{UUID} - Update status from specific workflow",
		"PATCH",
		"/workflow",
		controller.UpdateWorkflow,
	},

	Route{
		"[GET]: /consume - Consume a workflow from queue and generete a CSV file with workflow.Data",
		"GET",
		"/consume",
		controller.ConsumeWorkflows,
	}}

// ,

// ,
// Route{
// 	"AddProduct",
// 	"POST",
// 	"/AddProduct",
// 	AuthenticationMiddleware(controller.AddProduct),
// },
// Route{
// 	"UpdateProduct",
// 	"PUT",
// 	"/UpdateProduct",
// 	AuthenticationMiddleware(controller.UpdateProduct),
// },
// // Get Product by {id}
// Route{
// 	"GetProduct",
// 	"GET",
// 	"/products/{id}",
// 	controller.GetProduct,
// },
// // Delete Product by {id}
// Route{
// 	"DeleteProduct",
// 	"DELETE",
// 	"/deleteProduct/{id}",
// 	AuthenticationMiddleware(controller.DeleteProduct),
// },
// // Search product with string
// Route{
// 	"SearchProduct",
// 	"GET",
// 	"/Search/{query}",
// 	controller.SearchProduct,
// }}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
