package types

type BookRequest struct {
	ID          uint   `json:"bookID"`
	BookName    string `json:"bookName"`
	Author      string `json:"author"`
	Publication string `json:"publication,omitempty"`
}
