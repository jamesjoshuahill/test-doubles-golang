package counterfeiter_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"errors"

	"github.com/jamesjoshuahill/test-doubles-golang/counterfeiter"
	"github.com/jamesjoshuahill/test-doubles-golang/counterfeiter/counterfeiterfakes"
)

var _ = Describe("BookService", func() {
	It("calls the repository with the name and category", func() {
		repo := new(counterfeiterfakes.FakeRepository)
		service := counterfeiter.NewBookService(repo)

		service.First("name")

		Expect(repo.QueryCallCount()).To(Equal(1))
		name, kind := repo.QueryArgsForCall(0)
		Expect(name).To(Equal("name"))
		Expect(kind).To(Equal("book"))
	})

	Context("when no books are found", func() {
		It("returns an error", func() {
			repo := new(counterfeiterfakes.FakeRepository)
			repo.QueryReturns([]counterfeiter.Record{}, nil)
			service := counterfeiter.NewBookService(repo)

			_, err := service.First("name")

			Expect(err).To(MatchError(`no books found with name "name"`))
		})
	})

	Context("when one matching book is found", func() {
		It("returns the matching book", func() {
			repo := new(counterfeiterfakes.FakeRepository)
			repo.QueryReturns([]counterfeiter.Record{
				{Name: "name", Category: "book"},
			}, nil)
			service := counterfeiter.NewBookService(repo)

			record, err := service.First("name")

			Expect(err).NotTo(HaveOccurred())
			Expect(record.Name).To(Equal("name"))
		})
	})

	Context("when several books are found", func() {
		It("returns the first matching book", func() {
			repo := new(counterfeiterfakes.FakeRepository)
			repo.QueryReturns([]counterfeiter.Record{
				{Name: "name", Category: "book"},
				{Name: "another name", Category: "book"},
				{Name: "the name", Category: "book"},
			}, nil)
			service := counterfeiter.NewBookService(repo)

			record, err := service.First("name")

			Expect(err).NotTo(HaveOccurred())
			Expect(record.Name).To(Equal("name"))
		})
	})

	Context("when finding books fails", func() {
		It("returns an error", func() {
			repo := new(counterfeiterfakes.FakeRepository)
			repo.QueryReturns(nil, errors.New("find failed"))
			service := counterfeiter.NewBookService(repo)

			_, err := service.First("name")

			Expect(err).To(MatchError(`finding books with name "name" failed: find failed`))
		})
	})
})
