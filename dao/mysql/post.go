package mysql

import (
	"RedBubble/models"
)

// 新增帖子
func AddPost(post *models.Post) (err error) {
	result := mdb.Select("PostId", "UserId", "Username", "CommunityID", "Status", "Title", "Content").Create(post)
	return result.Error
}
