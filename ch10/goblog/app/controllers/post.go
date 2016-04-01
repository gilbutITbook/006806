package controllers

import (
	"goblog/app/models"
	"goblog/app/routes"

	"github.com/revel/revel"
)

// ➊ App을 임베디드 필드로 지정
type Post struct {
	App
}

func (c Post) CheckUser() revel.Result {
	// Index, Show는 권한 체크를 하지 않음.
	switch c.MethodName {
	case "Index", "Show":
		return nil
	}

	// CurrentUser 정보가 없으면 로그인 페이지로 이동
	if c.CurrentUser == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(App.Login)
	}

	// CurrentUser가 admin이 아니면 로그인 페이지로 이동
	if c.CurrentUser.Role != "admin" {
		c.Response.Status = 401 // Unauthorized
		c.Flash.Error("You are not admin")
		return c.Redirect(App.Login)
	}
	return nil
}

// ➋ Order메쏘드와 Find 메쏘드를 통해 전체 post 조회
func (c Post) Index() revel.Result {
	var posts []models.Post
	c.Txn.Order("created_at desc").Find(&posts)
	return c.Render(posts)
}

// ➌ First 메쏘드를 통해 id에 해당하는 포스트 조회
func (c Post) Show(id int) revel.Result {
	var post models.Post
	c.Txn.First(&post, id)
	c.Txn.Where(&models.Comment{PostId: id}).Find(&post.Comments)
	return c.Render(post)
}

// ➍ Save 메쏘드를 통해 포스트 수정
func (c Post) Update(id int, title, body string) revel.Result {
	var post models.Post
	c.Txn.First(&post, id)
	post.Title = title
	post.Body = body

	c.Txn.Save(&post)

	c.Flash.Success("포스트 수정 완료")
	return c.Redirect(routes.Post.Show(id))
}

// ➎ Create 메쏘드를 통해 포스트 생성
func (c Post) Create(title, body string) revel.Result {
	post := models.Post{Title: title, Body: body}
	c.Txn.Create(&post)
	c.Flash.Success("포스트 작성 완료")
	return c.Redirect(routes.Post.Index())
}

// ➏ Delete 메쏘를 통해 포스트 삭제
func (c Post) Destroy(id int) revel.Result {
	c.Txn.Where("post_id = ?", id).Delete(&models.Comment{})
	c.Txn.Where("id = ?", id).Delete(&models.Post{})
	c.Flash.Success("포스트 삭제 완료")
	return c.Redirect(routes.Post.Index())
}

func (c Post) New() revel.Result {
	post := models.Post{}
	return c.Render(post)
}

func (c Post) Edit(id int) revel.Result {
	var post models.Post
	c.Txn.First(&post, id)
	return c.Render(post)
}
