// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tGormController struct {}
var GormController tGormController


func (_ tGormController) Begin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GormController.Begin", args).Url
}

func (_ tGormController) Rollback(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GormController.Rollback", args).Url
}

func (_ tGormController) Commit(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GormController.Commit", args).Url
}


type tHome struct {}
var Home tHome


func (_ tHome) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Home.Index", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tApp struct {}
var App tApp


func (_ tApp) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Login", args).Url
}

func (_ tApp) CreateSession(
		username string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("App.CreateSession", args).Url
}

func (_ tApp) DestroySession(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.DestroySession", args).Url
}


type tComment struct {}
var Comment tComment


func (_ tComment) CheckUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Comment.CheckUser", args).Url
}

func (_ tComment) Create(
		postId int,
		body string,
		commenter string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "postId", postId)
	revel.Unbind(args, "body", body)
	revel.Unbind(args, "commenter", commenter)
	return revel.MainRouter.Reverse("Comment.Create", args).Url
}

func (_ tComment) Destroy(
		postId int,
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "postId", postId)
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Comment.Destroy", args).Url
}


type tPost struct {}
var Post tPost


func (_ tPost) CheckUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Post.CheckUser", args).Url
}

func (_ tPost) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Post.Index", args).Url
}

func (_ tPost) Show(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Post.Show", args).Url
}

func (_ tPost) Update(
		id int,
		title string,
		body string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "title", title)
	revel.Unbind(args, "body", body)
	return revel.MainRouter.Reverse("Post.Update", args).Url
}

func (_ tPost) Create(
		title string,
		body string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "title", title)
	revel.Unbind(args, "body", body)
	return revel.MainRouter.Reverse("Post.Create", args).Url
}

func (_ tPost) Destroy(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Post.Destroy", args).Url
}

func (_ tPost) New(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Post.New", args).Url
}

func (_ tPost) Edit(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Post.Edit", args).Url
}


