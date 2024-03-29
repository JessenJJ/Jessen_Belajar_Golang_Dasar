package basic

type Post struct {
	Title   string
	Content string
	Owner   User
}

// conastructor => fungsi yang membentuk struct
func NewPost(title string, content string, owner User) *Post {
	return &Post{title, content, owner}
}

func (p *Post) EditContent(newContent string) {
	p.Content = newContent
}
