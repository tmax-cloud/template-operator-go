package controller

import (
	"github.com/jwkim1993/hypercloud-operator/pkg/controller/clustertemplate"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, clustertemplate.Add)
}
