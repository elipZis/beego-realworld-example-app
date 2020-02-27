package api

import (
	"github.com/astaxie/beego"
)

type CommentController struct {
	beego.Controller
}

// @router /articles/:slug/comments [get]
func (this *CommentController) Comments(slug string) {
}

// @router /articles/:slug/comments [post]
func (this *CommentController) AddComment(slug string) {
}

// @router /articles/:slug/comments/:id [delete]
func (this *CommentController) DeleteComment(slug string, id int) {
}

// @router /articles/:slug/favorite [post]
func (this *CommentController) Favorite(slug string) {
}

// @router /articles/:slug/favorite [delete]
func (this *CommentController) UnFavorite(slug string) {
}
