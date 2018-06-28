package minimock_test

import (
	"errors"

	mm "github.com/gojuno/minimock"
	"github.com/jamesjoshuahill/test-doubles-golang/minimock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BookService", func() {
	It("calls the repository with the name and category", func() {
		mc := mm.NewController(GinkgoT())
		repo := minimock.NewrepositoryMock(mc)
		repo.QueryMock.Expect("name", "book").Return(nil, nil)
		service := minimock.NewBookService(repo)

		service.First("name")

		Expect(repo.QueryCounter).To(Equal(uint64(1)))
		mc.Finish()
	})

	Context("when no books are found", func() {
		It("returns an error", func() {
			mc := mm.NewController(GinkgoT())
			repo := minimock.NewrepositoryMock(mc)
			repo.QueryMock.Return([]minimock.Record{}, nil)
			service := minimock.NewBookService(repo)

			_, err := service.First("name")

			Expect(err).To(MatchError(`no books found with name "name"`))
			mc.Finish()
		})
	})

	Context("when one matching book is found", func() {
		It("returns the matching book", func() {
			mc := mm.NewController(GinkgoT())
			repo := minimock.NewrepositoryMock(mc)
			repo.QueryMock.Return([]minimock.Record{
				{Name: "name", Category: "book"},
			}, nil)
			service := minimock.NewBookService(repo)

			record, err := service.First("name")

			Expect(err).NotTo(HaveOccurred())
			Expect(record.Name).To(Equal("name"))
			mc.Finish()
		})
	})

	Context("when several books are found", func() {
		It("returns the first matching book", func() {
			mc := mm.NewController(GinkgoT())
			repo := minimock.NewrepositoryMock(mc)
			repo.QueryMock.Return([]minimock.Record{
				{Name: "name", Category: "book"},
				{Name: "another name", Category: "book"},
				{Name: "the name", Category: "book"},
			}, nil)
			service := minimock.NewBookService(repo)

			record, err := service.First("name")

			Expect(err).NotTo(HaveOccurred())
			Expect(record.Name).To(Equal("name"))
			mc.Finish()
		})
	})

	Context("when finding books fails", func() {
		It("returns an error", func() {
			mc := mm.NewController(GinkgoT())
			repo := minimock.NewrepositoryMock(mc)
			repo.QueryMock.Return(nil, errors.New("find failed"))
			service := minimock.NewBookService(repo)

			_, err := service.First("name")

			Expect(err).To(MatchError(`finding books with name "name" failed: find failed`))
			mc.Finish()
		})
	})

})
