package book

type Book struct {
	Title  string
	Author string
	Page   int
}

func (book *Book) CategoryByLength() (category string) {
	if book.Page > 300 {
		return "NOVEL"
	}
	return "SHORT STORY"
}
