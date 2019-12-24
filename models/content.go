package models

//easyjson
type (
	ContentModel struct {
		ContentId   uint64 `json:"contentId" gorm:"primary_key;column:content_id"`
		ContentType int    `json:"contentType" gorm:"column:content_type"`
		UserId      uint64 `json:"userId" gorm:"column:user_id"`
		Title       string `json:"title" gorm:"column:title"`
		Content     string `json:"content" gorm:"column:content"`
		Cover       string `json:"cover" gorm:"column:cover"`
		AtUserIds   string `json:"atUserIds" gorm:"column:at_user_ids"`
		ShowStatus  int    `json:"showStatus" gorm:"column:show_status"`
		SourceApp   string `json:"sourceApp" gorm:"column:source_app"`
		Del         int    `json:"del" gorm:"column:del"`
		AddTime     int64  `json:"addTime" gorm:"column:add_time"`
		EditTime    int64  `json:"editTime" gorm:"column:edit_time"`
	}
)
