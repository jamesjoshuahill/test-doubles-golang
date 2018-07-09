package pegomocks_test

import (
	. "github.com/onsi/gomega"
	. "github.com/petergtz/pegomock"

	"errors"

	"github.com/jamesjoshuahill/test-doubles-golang/pegomocks"
)

var _ = Describe("BookService", func() {
	It("calls the repository with the name and category", func() {
		repo := NewMockrepository()
		service := pegomocks.NewBookService(repo)

		service.First("name")

		repo.VerifyWasCalledOnce().Query("name", "book")
	})

	Context("when no books are found", func() {
		It("returns an error", func() {
			repo := NewMockrepository()
			When(repo.Query(AnyString(), AnyString())).ThenReturn([]pegomocks.Record{}, nil)
			service := pegomocks.NewBookService(repo)

			_, err := service.First("name")

			Expect(err).To(MatchError(`no books found with name "name"`))
		})
	})

	Context("when one matching book is found", func() {
		It("returns the matching book", func() {
			repo := NewMockrepository()
			When(repo.Query(AnyString(), AnyString())).ThenReturn([]pegomocks.Record{
				{Name: "name", Category: "book"},
			}, nil)
			service := pegomocks.NewBookService(repo)

			record, err := service.First("name")

			Expect(err).NotTo(HaveOccurred())
			Expect(record.Name).To(Equal("name"))
		})
	})

	Context("when several books are found", func() {
		It("returns the first matching book", func() {
			repo := NewMockrepository()
			When(repo.Query(AnyString(), AnyString())).ThenReturn([]pegomocks.Record{
				{Name: "name", Category: "book"},
				{Name: "another name", Category: "book"},
				{Name: "the name", Category: "book"},
			}, nil)
			service := pegomocks.NewBookService(repo)

			record, err := service.First("name")

			Expect(err).NotTo(HaveOccurred())
			Expect(record.Name).To(Equal("name"))
		})
	})

	Context("when finding books fails", func() {
		It("returns an error", func() {
			repo := NewMockrepository()
			When(repo.Query(AnyString(), AnyString())).ThenReturn(nil, errors.New("find failed"))
			service := pegomocks.NewBookService(repo)

			_, err := service.First("name")

			Expect(err).To(MatchError(`finding books with name "name" failed: find failed`))
		})
	})
})
