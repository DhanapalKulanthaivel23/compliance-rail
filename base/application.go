package base

import (
	"github.com/DhanapalKulanthaivel23/compliance-rail/controller"
	"net/http"
)

func New() http.Handler { return controller.New() }
