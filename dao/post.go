package dao

import (
	"database/sql"
	"log"
	"models"
)

func CountGetPostByCategory(cId int) (count int) {
	rows := DB.QueryRow("select count(1) from blog_post where category_id = ?", cId)
	_ = rows.Scan(&count)
	return
}

func CountGetAllPost() (count int) {
	rows := DB.QueryRow("select count(1) from blog_post")
	_ = rows.Scan(&count)
	return
}

func CountPostBySlug(slug string) (count int) {
	rows := DB.QueryRow("select count(1) from blog_post where slug = ?", slug)
	_ = rows.Scan(&count)
	return
}

// 从数据库返回结果填充posts数组
func rowsToPosts(r *sql.Rows) ([]models.Post, error) {
	var posts []models.Post

	for r.Next() {
		var post models.Post

		err := r.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)

		if err != nil {
			log.Println("posts填充错误")
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil

}

func GetPostPage(page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize

	rows, err := DB.Query("select* from blog_post limit ?,?", page, pageSize)

	if err != nil {
		return nil, err
	}

	posts, _ := rowsToPosts(rows)

	return posts, nil

}

// 获取所有posts
func GetAllPost() ([]models.Post, error) {
	rows, err := DB.Query("select* from blog_post")

	if err != nil {
		return nil, err
	}

	posts, _ := rowsToPosts(rows)

	return posts, nil
}

func GetPostPageByCategory(cId, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select* from blog_post where category_id  = ? limit ?,?", cId, page, pageSize)

	if err != nil {
		return nil, err
	}

	posts, _ := rowsToPosts(rows)

	return posts, nil

}

func GetPostPageBySlug(slug string, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize

	rows, err := DB.Query("select* from blog_post where slug  = ? limit ?,?", slug, page, pageSize)

	if err != nil {
		return nil, err
	}

	posts, _ := rowsToPosts(rows)

	// for rows.Next() {
	// 	var post models.Post

	// 	err := rows.Scan(
	// 		&post.Pid,
	// 		&post.Title,
	// 		&post.Content,
	// 		&post.Markdown,
	// 		&post.CategoryId,
	// 		&post.UserId,
	// 		&post.ViewCount,
	// 		&post.Type,
	// 		&post.Slug,
	// 		&post.CreateAt,
	// 		&post.UpdateAt,
	// 	)

	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	posts = append(posts, post)

	// }

	return posts, nil

}

func GetPostById(pId int) (*models.Post, error) {
	p := &models.Post{}

	err := DB.QueryOne(p, "select * from blog_post where pid = ?", pId)

	return p, err
	// row := DB.QueryRow("select * from blog_post where pid = ?", pId)

	// var post models.Post
	// if row.Err() != nil {
	// 	return post, row.Err()
	// }

	// err := row.Scan(
	// 	&post.Pid,
	// 	&post.Title,
	// 	&post.Content,
	// 	&post.Markdown,
	// 	&post.CategoryId,
	// 	&post.UserId,
	// 	&post.ViewCount,
	// 	&post.Type,
	// 	&post.Slug,
	// 	&post.CreateAt,
	// 	&post.UpdateAt,
	// )

	// if err != nil {
	// 	return post, row.Err()
	// }
}

func GetPostSearchRes(condition string) ([]models.Post, error) {
	rows, err := DB.Query("select* from blog_post where title like ?", "%"+condition+"%")

	if err != nil {
		return nil, err
	}

	posts, _ := rowsToPosts(rows)

	return posts, nil
}

func SavePost(post *models.Post) {
	res, err := DB.Exec("insert into blog_post"+
		"(title, content, markdown, category_id, user_id, view_count, type, slug, create_at, update_at)"+
		"values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.UserId,
		post.ViewCount,
		post.Type,
		post.Slug,
		post.CreateAt,
		post.UpdateAt)

	if err != nil {
		log.Panicln(err)
	}

	pid, _ := res.LastInsertId()

	post.Pid = int(pid)
}

func UpdatePost(post *models.Post) {
	_, err := DB.Exec("update blog_post set title = ?, content = ?, markdown = ?, category_id = ?, type = ?, slug = ?, update_at = ? where pid = ?",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.Type,
		post.Slug,
		post.UpdateAt,
		post.Pid)

	if err != nil {
		log.Panicln(err)
	}

}
