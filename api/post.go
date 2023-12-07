package api

import (
	"Blog/common"
	"Blog/dao"
	"Blog/service"
	"Blog/utils"
	"errors"
	"models"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func paramToInt(para interface{}) int {
	var res int
	switch para.(type) {
	case string:
		str := para.(string)
		res, _ = strconv.Atoi(str)
	case float64:
		fl := para.(float64)
		res = int(fl)
	}

	return res

}

func (*APIHandler) SaveandPost(w http.ResponseWriter, r *http.Request) {
	// 判断用户登录 验证token是否合法
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		common.Error(w, errors.New("用户过期"))
		return
	}
	uId := claim.Uid

	// POST save
	method := r.Method
	switch method {
	case http.MethodPost:
		params := common.GetRequestJsonParam(r)
		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pType := int(postType)
		post := &models.Post{
			Pid:        -1,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: categoryId,
			UserId:     uId,
			ViewCount:  0,
			Type:       pType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}

		service.SavePost(post)
		common.Success(w, post)
	case http.MethodPut:
		// update
		params := common.GetRequestJsonParam(r)

		categoryId := paramToInt(params["categoryId"])
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		pType := paramToInt(params["type"])
		pId := paramToInt(params["pid"])
		post := &models.Post{
			Pid:        pId,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: categoryId,
			UserId:     uId,
			ViewCount:  0,
			Type:       pType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}

		service.UpdatePost(post)
		common.Success(w, post)
	}

}

func (*APIHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	// fmt.Println(path)
	pIdStr := strings.TrimPrefix(path, "/api/v1/post/")
	pId, err := strconv.Atoi(pIdStr)

	if err != nil {
		common.Error(w, errors.New("路径不匹配"))
		return
	}

	post, err := dao.GetPostById(pId)

	if err != nil {
		common.Error(w, err)
		return
	}

	common.Success(w, post)

}

func (*APIHandler) SearchPost(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	condition := r.Form.Get("val")
	searchRes := service.SearchPost(condition)

	common.Success(w, searchRes)

}
