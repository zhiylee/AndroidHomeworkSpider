package model

type Article struct {
	Id int `gorm:"primary_key" json:"id"`

	Title string `json:"title"`
	Content string `json:"content"`
	Cover string `json:"cover"`
	CreateTime int `json:"createTime"`
	CategoryId int `json:categoryId`
	CommentCount int `json:"commentCount"`
	Author string `json:"author"`
}

func AddArticle(article Article) Article {
	db.Create(&article)

	return article
}

func getArticle()  {

}

func ArticleIsExist(article Article) bool {
	articleOut := Article{}
	db.Find(&articleOut,article)
	if articleOut.Id<1 {
		return false
	}

	return true
}