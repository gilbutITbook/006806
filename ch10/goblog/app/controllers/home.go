package controllers

import "github.com/revel/revel"

type Home struct {
	*revel.Controller
}

func (c Home) Index() revel.Result {
	return c.Render()
}
