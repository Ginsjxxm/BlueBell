package mysql

import (
	"BlueBell/models"
)

func CreatePost(p *models.Post) error {
	sqlStr := `insert into post(
                 post_id,title,content,author_id,community_id)
                 values(?,?,?,?,?)`
	_, err := db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	if err != nil {
		return err
	}
	return nil
}

func GetPostByID(pid int64) (post *models.Post, err error) {
	sqlStr := `select post_id,title,content,author_id,community_id,create_time from post where post.post_id=?`
	post = new(models.Post)
	err = db.Get(post, sqlStr, pid)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func GetPostList(offset int64, limit int64) (posts []*models.Post, err error) {
	// 定义 SQL 查询语句，使用占位符 `?`
	sqlStr := `SELECT post_id, title, content, author_id, community_id, create_time 
               FROM post 
               LIMIT ?, ?`

	posts = make([]*models.Post, 0, limit) // 根据 limit 设定初始容量

	// 执行查询并绑定结果
	err = db.Select(&posts, sqlStr, offset, limit)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
