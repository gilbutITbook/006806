package controllers

import (
	"goblog/app/models"
	"goblog/app/routes"

	"github.com/revel/revel"
)

type Comment struct {
	App
}

func (c Comment) CheckUser() revel.Result {
	if c.MethodName != "Destroy" {
		return nil
	}

	if c.CurrentUser == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(App.Login)
	}

	if c.CurrentUser.Role != "admin" {
		c.Response.Status = 401 // Unauthorized
		c.Flash.Error("You are not admin")
		return c.Redirect(App.Login)
	}
	return nil
}

func (c Comment) Create(postId int, body, commenter string) revel.Result {
	comment := models.Comment{PostId: postId, Body: body, Commenter: commenter}
	c.Txn.Create(&comment)
	c.Flash.Success("댓글 작성 완료")
	return c.Redirect(routes.Post.Show(postId))
}

func (c Comment) Destroy(postId, id int) revel.Result {
	c.Txn.Where("id = ?", id).Delete(&models.Comment{})
	c.Flash.Success("댓글 삭제 완료")
	return c.Redirect(routes.Post.Show(postId))
}
