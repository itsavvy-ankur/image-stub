package swagger

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var IC ImageConfig

func NewRouter(imgYamlConfigPath string) *mux.Router {
	log.Printf("Image Config File Path : %s\n", imgYamlConfigPath)
	fileBytes, _ := ioutil.ReadFile(imgYamlConfigPath)
	err := yaml.Unmarshal(fileBytes, &IC)
	if err != nil {
		fmt.Println(err)
	}

	log.Printf("Image Config router: %+v\n", IC)
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v1/scanimages/",
		Index,
	},

	Route{
		"Getimage",
		strings.ToUpper("Get"),
		"/v1/scanimages/getimage",
		Getimage,
	},
}
