package charlatan_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"errors"

	"github.com/jamesjoshuahill/test-doubles-golang/charlatan"
)

var _ = Describe("BookService", func() {
	It("calls the repository with the name and name and category", func() {
		repo := new(charlatan.Fakerepository)
		repo.SetQueryStub(nil, nil)
		service := charlatan.NewBookService(repo)

		service.First("name")

		Expect(repo.QueryCalledOnceWith("name", "book")).To(BeTrue())
	})

	Context("when no books are found", func() {
		It("returns an error", func() {
			repo := new(charlatan.Fakerepository)
			repo.SetQueryStub([]charlatan.Record{}, nil)
			service := charlatan.NewBookService(repo)

			_, err := service.First("name")

			Expect(err).To(MatchError(`no books found with name "name"`))
		})
	})

	Context("when one matching book is found", func() {
		It("returns the matching book", func() {
			repo := new(charlatan.Fakerepository)
			repo.SetQueryStub([]charlatan.Record{
				{Name: "name", Category: "book"},
			}, nil)
			service := charlatan.NewBookService(repo)

			record, err := service.First("name")

			Expect(err).NotTo(HaveOccurred())
			Expect(record.Name).To(Equal("name"))
		})
	})

	Context("when several books are found", func() {
		It("returns the first matching book", func() {
			repo := new(charlatan.Fakerepository)
			repo.SetQueryStub([]charlatan.Record{
				{Name: "name", Category: "book"},
				{Name: "another name", Category: "book"},
				{Name: "the name", Category: "book"},
			}, nil)
			service := charlatan.NewBookService(repo)

			record, err := service.First("name")

			Expect(err).NotTo(HaveOccurred())
			Expect(record.Name).To(Equal("name"))
		})
	})

	Context("when finding books fails", func() {
		It("returns an error", func() {
			repo := new(charlatan.Fakerepository)
			repo.SetQueryStub(nil, errors.New("find failed"))
			service := charlatan.NewBookService(repo)

			_, err := service.First("name")

			Expect(err).To(MatchError(`finding books with name "name" failed: find failed`))
		})
	})
})
