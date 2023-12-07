package service

import (
	"Blog/dao"
	"config"
	"models"
)

func FindPostPigeonhole() models.PigeonholeRes {
	// 查询所有的文章 进行月份整理
	// 查询所有的分类
	categorys, _ := dao.GetCategory()

	posts, _ := dao.GetAllPost()
	pigeonholeMap := make(map[string][]models.Post, 0)
	for _, post := range posts {
		creatAt := post.CreateAt
		month := creatAt.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], post)
	}

	return models.PigeonholeRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Categorys:    categorys,
		Lines:        pigeonholeMap,
	}
}
