package handrolled_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"errors"

	"github.com/jamesjoshuahill/test-doubles-golang/handrolled"
)

var _ = Describe("BookService", func() {
	It("calls the repository with the name and category", func() {
		repo := new(FakeRepository)
		service := handrolled.NewBookService(repo)

		service.First("name")

		Expect(repo.QueryMock.CallCount).To(Equal(1))
		Expect(repo.QueryMock.Received.Name).To(Equal("name"))
		Expect(repo.QueryMock.Received.Category).To(Equal("book"))
	})

	Context("when no books are found", func() {
		It("returns an error", func() {
			repo := new(FakeRepository)
			service := handrolled.NewBookService(repo)

			_, err := service.First("name")

			Expect(err).To(MatchError(`no books found with name "name"`))
		})
	})

	Context("when one matching book is found", func() {
		It("returns the matching book", func() {
			repo := new(FakeRepository)
			repo.QueryMock.Returns.Records = []handrolled.Record{
				{Name: "name", Category: "book"},
			}
			service := handrolled.NewBookService(repo)

			record, err := service.First("name")

			Expect(err).NotTo(HaveOccurred())
			Expect(record.Name).To(Equal("name"))
		})
	})

	Context("when several books are found", func() {
		It("returns the first matching book", func() {
			repo := new(FakeRepository)
			repo.QueryMock.Returns.Records = []handrolled.Record{
				{Name: "name", Category: "book"},
				{Name: "another name", Category: "book"},
				{Name: "the name", Category: "book"},
			}
			service := handrolled.NewBookService(repo)

			record, err := service.First("name")

			Expect(err).NotTo(HaveOccurred())
			Expect(record.Name).To(Equal("name"))
		})
	})

	Context("when finding books fails", func() {
		It("returns an error", func() {
			repo := new(FakeRepository)
			repo.QueryMock.Returns.Error = errors.New("find failed")
			service := handrolled.NewBookService(repo)

			_, err := service.First("name")

			Expect(err).To(MatchError(`finding books with name "name" failed: find failed`))
		})
	})
})
