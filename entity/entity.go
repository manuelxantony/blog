package entity

type BlogPost struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedBy string `json:"createdby"`
}
