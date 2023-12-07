package service

import (
	"Blog/dao"
	"config"
	"html/template"
	"log"
	"models"
)

func SavePost(post *models.Post) {
	dao.SavePost(post)
}

func UpdatePost(post *models.Post) {
	dao.UpdatePost(post)
}

func GetPostDetail(pId int) (*models.PostRes, error) {
	post, err := dao.GetPostById(pId)

	if err != nil {
		return nil, err
	}

	categoryName := dao.GetCategoryById(post.CategoryId)
	userName := dao.GetUserNameById(post.UserId)

	postMore := models.PostMore{
		Pid:          post.Pid,
		Title:        post.Title,
		Slug:         post.Slug,
		Content:      template.HTML(post.Content),
		CategoryId:   post.CategoryId,
		CategoryName: categoryName,
		UserId:       post.UserId,
		UserName:     userName,
		ViewCount:    post.ViewCount,
		Type:         post.Type,
		CreateAt:     models.DateDay(post.CreateAt),
		UpdateAt:     models.DateDay(post.UpdateAt),
	}

	var postRes = &models.PostRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Article:      postMore,
	}

	return postRes, nil
}

func Writing() (wr models.WritingRes) {
	wr.Title = config.Cfg.Viewer.Title
	wr.CdnURL = config.Cfg.System.CdnURL

	category, err := dao.GetCategory()

	if err != nil {
		log.Println(err)
		return
	}

	wr.Categorys = category
	return
}

func SearchPost(condition string) []models.SearchResp {
	posts, _ := dao.GetPostSearchRes(condition)
	var searchResps []models.SearchResp
	for _, post := range posts {
		searchResps = append(searchResps, models.SearchResp{
			Pid:   post.Pid,
			Title: post.Title,
		})
	}

	return searchResps
}
