package api

import (
	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

// @router /articles [get]
func (this *ArticleController) Articles(tag string, author string, favorited string, limit int, offset int) {
}

// @router /articles/feed [get]
func (this *ArticleController) Feed(limit int, offset int) {
}

// @router /articles/:slug [get]
func (this *ArticleController) Article(slug string) {
}

// @router /articles [post]
func (this *ArticleController) CreateArticle() {
}

// @router /articles/:slug [put]
func (this *ArticleController) UpdateArticle(slug string) {
}

// @router /articles/:slug [delete]
func (this *ArticleController) DeleteArticle(slug string) {
}