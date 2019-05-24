package book_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/yuryanliang/go_playground/ginkgo/book"
)

//var _ = Describe("Book", func() {
//	var(
//		longBook Book
//		shortBook Book
//	)
//
//	BeforeEach(func() {
//		longBook=Book{
//			Title:"Les Miserables",
//			Author: "Victor Hugo",
//			Page:1488,
//
//		}
//		shortBook=Book{
//			Title:"Fox in Socks",
//			Author: "Dr. Seuss",
//			Page:24,
//
//		}
//		Expect(shortBook.Title).Should(Equal("Fox in Socks"))
//	})
//
//	Describe("Categorizing book length", func() {
//		//defer GinkgoRecover()
//
//		Context("With more than 300 pages", func() {
//			It("should be a novel", func() {
//				Expect(longBook.CategoryByLength()).Should(Equal("NOVEL"))
//			})
//		})
//		//Fail("goog")
//
//		Context("With fewer than 300 pages", func() {
//			It("should be a short story", func() {
//				Expect(shortBook.CategoryByLength()).Should(Equal("SHORT STORY"))
//
//			})
//		})
//	})
//})
var _ = Describe("Book", func() {
	var (
		book Book
		err  error
	)

	BeforeEach(func() {
		book, err = NewBookFromJSON(`{
            "title":"Les Miserables",
            "author":"Victor Hugo",
            "pages":1488
        }`)
	})

	Describe("loading from JSON", func() {
		Context("when the JSON parses succesfully", func() {
			It("should populate the fields correctly", func() {
				Expect(book.Title).To(Equal("Les Miserables"))
				Expect(book.Author).To(Equal("Victor Hugo"))
				Expect(book.Pages).To(Equal(1488))
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the JSON fails to parse", func() {
			BeforeEach(func() {
				book, err = NewBookFromJSON(`{
                    "title":"Les Miserables",
                    "author":"Victor Hugo",
                    "pages":1488oops
                }`)
			})

			It("should return the zero-value for the book", func() {
				Expect(book).To(BeZero())
			})

			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Extracting the author's last name", func() {
		It("should correctly identify and return the last name", func() {
			Expect(book.AuthorLastName()).To(Equal("Hugo"))
		})
	})
})
