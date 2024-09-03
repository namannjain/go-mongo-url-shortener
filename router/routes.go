package router

import (
	"net/http"
	"urlShortenerMongo/constant"
	"urlShortenerMongo/controller"
)

var urlShortener = Routes{
	Route{"Url Shortening Service", http.MethodPost, constant.UrlShortenerPath, controller.ShortTheUrl},
	Route{"Redirect to Url", http.MethodGet, constant.RedirectUrlPath, controller.RedirectUrl},
}
