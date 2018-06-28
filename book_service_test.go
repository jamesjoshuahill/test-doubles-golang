package doubles_test

import (
	"errors"

	"github.com/gojuno/minimock"
	"github.com/jamesjoshuahill/test-doubles-golang"
	"github.com/jamesjoshuahill/test-doubles-golang/charlatan"
	"github.com/jamesjoshuahill/test-doubles-golang/counterfeiter"
	"github.com/jamesjoshuahill/test-doubles-golang/handrolled"
	"github.com/jamesjoshuahill/test-doubles-golang/minimocks"
	"github.com/jamesjoshuahill/test-doubles-golang/pegomock"
	. "github.com/onsi/gomega"
	. "github.com/petergtz/pegomock"
)

var _ = Describe("BookService#First", func() {
	Context("with a hand-rolled fake logger", func() {
		It("calls the repository with the name and category", func() {
			repo := new(handrolled.FakeRepository)
			service := doubles.NewBookService(repo)

			service.First("name")

			Expect(repo.QueryMock.CallCount).To(Equal(1))
			Expect(repo.QueryMock.Received.Name).To(Equal("name"))
			Expect(repo.QueryMock.Received.Kind).To(Equal("book"))
		})

		Context("when no books are found", func() {
			It("returns an error", func() {
				repo := new(handrolled.FakeRepository)
				service := doubles.NewBookService(repo)

				_, err := service.First("name")

				Expect(err).To(MatchError(`no books found with name "name"`))
			})
		})

		Context("when one matching book is found", func() {
			It("returns the matching book", func() {
				repo := new(handrolled.FakeRepository)
				repo.QueryMock.Returns.Records = []doubles.Record{
					{Name: "name", Category: "book"},
				}
				service := doubles.NewBookService(repo)

				record, err := service.First("name")

				Expect(err).NotTo(HaveOccurred())
				Expect(record.Name).To(Equal("name"))
			})
		})

		Context("when several books are found", func() {
			It("returns the first matching book", func() {
				repo := new(handrolled.FakeRepository)
				repo.QueryMock.Returns.Records = []doubles.Record{
					{Name: "name", Category: "book"},
					{Name: "another name", Category: "book"},
					{Name: "the name", Category: "book"},
				}
				service := doubles.NewBookService(repo)

				record, err := service.First("name")

				Expect(err).NotTo(HaveOccurred())
				Expect(record.Name).To(Equal("name"))
			})
		})

		Context("when finding books fails", func() {
			It("returns an error", func() {
				repo := new(handrolled.FakeRepository)
				repo.QueryMock.Returns.Error = errors.New("find failed")
				service := doubles.NewBookService(repo)

				_, err := service.First("name")

				Expect(err).To(MatchError(`finding books with name "name" failed: find failed`))
			})
		})
	})

	Context("with a charlatan fake logger", func() {
		It("calls the repository with the name and name and category", func() {
			repo := new(charlatan.Fakerepository)
			repo.SetQueryStub(nil, nil)
			service := doubles.NewBookService(repo)

			service.First("name")

			Expect(repo.QueryCalledOnceWith("name", "book")).To(BeTrue())
		})

		Context("when no books are found", func() {
			It("returns an error", func() {
				repo := new(charlatan.Fakerepository)
				repo.SetQueryStub([]doubles.Record{}, nil)
				service := doubles.NewBookService(repo)

				_, err := service.First("name")

				Expect(err).To(MatchError(`no books found with name "name"`))
			})
		})

		Context("when one matching book is found", func() {
			It("returns the matching book", func() {
				repo := new(charlatan.Fakerepository)
				repo.SetQueryStub([]doubles.Record{
					{Name: "name", Category: "book"},
				}, nil)
				service := doubles.NewBookService(repo)

				record, err := service.First("name")

				Expect(err).NotTo(HaveOccurred())
				Expect(record.Name).To(Equal("name"))
			})
		})

		Context("when several books are found", func() {
			It("returns the first matching book", func() {
				repo := new(charlatan.Fakerepository)
				repo.SetQueryStub([]doubles.Record{
					{Name: "name", Category: "book"},
					{Name: "another name", Category: "book"},
					{Name: "the name", Category: "book"},
				}, nil)
				service := doubles.NewBookService(repo)

				record, err := service.First("name")

				Expect(err).NotTo(HaveOccurred())
				Expect(record.Name).To(Equal("name"))
			})
		})

		Context("when finding books fails", func() {
			It("returns an error", func() {
				repo := new(charlatan.Fakerepository)
				repo.SetQueryStub(nil, errors.New("find failed"))
				service := doubles.NewBookService(repo)

				_, err := service.First("name")

				Expect(err).To(MatchError(`finding books with name "name" failed: find failed`))
			})
		})
	})

	Context("with a counterfeiter fake logger", func() {
		It("calls the repository with the name and category", func() {
			repo := new(counterfeiter.FakeRepository)
			service := doubles.NewBookService(repo)

			service.First("name")

			Expect(repo.QueryCallCount()).To(Equal(1))
			name, kind := repo.QueryArgsForCall(0)
			Expect(name).To(Equal("name"))
			Expect(kind).To(Equal("book"))
		})

		Context("when no books are found", func() {
			It("returns an error", func() {
				repo := new(counterfeiter.FakeRepository)
				repo.QueryReturns([]doubles.Record{}, nil)
				service := doubles.NewBookService(repo)

				_, err := service.First("name")

				Expect(err).To(MatchError(`no books found with name "name"`))
			})
		})

		Context("when one matching book is found", func() {
			It("returns the matching book", func() {
				repo := new(counterfeiter.FakeRepository)
				repo.QueryReturns([]doubles.Record{
					{Name: "name", Category: "book"},
				}, nil)
				service := doubles.NewBookService(repo)

				record, err := service.First("name")

				Expect(err).NotTo(HaveOccurred())
				Expect(record.Name).To(Equal("name"))
			})
		})

		Context("when several books are found", func() {
			It("returns the first matching book", func() {
				repo := new(counterfeiter.FakeRepository)
				repo.QueryReturns([]doubles.Record{
					{Name: "name", Category: "book"},
					{Name: "another name", Category: "book"},
					{Name: "the name", Category: "book"},
				}, nil)
				service := doubles.NewBookService(repo)

				record, err := service.First("name")

				Expect(err).NotTo(HaveOccurred())
				Expect(record.Name).To(Equal("name"))
			})
		})

		Context("when finding books fails", func() {
			It("returns an error", func() {
				repo := new(counterfeiter.FakeRepository)
				repo.QueryReturns(nil, errors.New("find failed"))
				service := doubles.NewBookService(repo)

				_, err := service.First("name")

				Expect(err).To(MatchError(`finding books with name "name" failed: find failed`))
			})
		})
	})

	Context("with a minimock mock logger", func() {
		It("calls the repository with the name and category", func() {
			mc := minimock.NewController(GinkgoT())
			repo := minimocks.NewrepositoryMock(mc)
			repo.QueryMock.Expect("name", "book").Return(nil, nil)
			service := doubles.NewBookService(repo)

			service.First("name")

			Expect(repo.QueryCounter).To(Equal(uint64(1)))
			mc.Finish()
		})

		Context("when no books are found", func() {
			It("returns an error", func() {
				mc := minimock.NewController(GinkgoT())
				repo := minimocks.NewrepositoryMock(mc)
				repo.QueryMock.Return([]doubles.Record{}, nil)
				service := doubles.NewBookService(repo)

				_, err := service.First("name")

				Expect(err).To(MatchError(`no books found with name "name"`))
				mc.Finish()
			})
		})

		Context("when one matching book is found", func() {
			It("returns the matching book", func() {
				mc := minimock.NewController(GinkgoT())
				repo := minimocks.NewrepositoryMock(mc)
				repo.QueryMock.Return([]doubles.Record{
					{Name: "name", Category: "book"},
				}, nil)
				service := doubles.NewBookService(repo)

				record, err := service.First("name")

				Expect(err).NotTo(HaveOccurred())
				Expect(record.Name).To(Equal("name"))
				mc.Finish()
			})
		})

		Context("when several books are found", func() {
			It("returns the first matching book", func() {
				mc := minimock.NewController(GinkgoT())
				repo := minimocks.NewrepositoryMock(mc)
				repo.QueryMock.Return([]doubles.Record{
					{Name: "name", Category: "book"},
					{Name: "another name", Category: "book"},
					{Name: "the name", Category: "book"},
				}, nil)
				service := doubles.NewBookService(repo)

				record, err := service.First("name")

				Expect(err).NotTo(HaveOccurred())
				Expect(record.Name).To(Equal("name"))
				mc.Finish()
			})
		})

		Context("when finding books fails", func() {
			It("returns an error", func() {
				mc := minimock.NewController(GinkgoT())
				repo := minimocks.NewrepositoryMock(mc)
				repo.QueryMock.Return(nil, errors.New("find failed"))
				service := doubles.NewBookService(repo)

				_, err := service.First("name")

				Expect(err).To(MatchError(`finding books with name "name" failed: find failed`))
				mc.Finish()
			})
		})
	})

	Context("with a pegomock mock logger", func() {
		It("calls the repository with the name and category", func() {
			repo := pegomock.NewMockrepository()
			service := doubles.NewBookService(repo)

			service.First("name")

			repo.VerifyWasCalledOnce().Query("name", "book")
		})

		Context("when no books are found", func() {
			It("returns an error", func() {
				repo := pegomock.NewMockrepository()
				When(repo.Query(AnyString(), AnyString())).ThenReturn([]doubles.Record{}, nil)
				service := doubles.NewBookService(repo)

				_, err := service.First("name")

				Expect(err).To(MatchError(`no books found with name "name"`))
			})
		})

		Context("when one matching book is found", func() {
			It("returns the matching book", func() {
				repo := pegomock.NewMockrepository()
				When(repo.Query(AnyString(), AnyString())).ThenReturn([]doubles.Record{
					{Name: "name", Category: "book"},
				}, nil)
				service := doubles.NewBookService(repo)

				record, err := service.First("name")

				Expect(err).NotTo(HaveOccurred())
				Expect(record.Name).To(Equal("name"))
			})
		})

		Context("when several books are found", func() {
			It("returns the first matching book", func() {
				repo := pegomock.NewMockrepository()
				When(repo.Query(AnyString(), AnyString())).ThenReturn([]doubles.Record{
					{Name: "name", Category: "book"},
					{Name: "another name", Category: "book"},
					{Name: "the name", Category: "book"},
				}, nil)
				service := doubles.NewBookService(repo)

				record, err := service.First("name")

				Expect(err).NotTo(HaveOccurred())
				Expect(record.Name).To(Equal("name"))
			})
		})

		Context("when finding books fails", func() {
			It("returns an error", func() {
				repo := pegomock.NewMockrepository()
				When(repo.Query(AnyString(), AnyString())).ThenReturn(nil, errors.New("find failed"))
				service := doubles.NewBookService(repo)

				_, err := service.First("name")

				Expect(err).To(MatchError(`finding books with name "name" failed: find failed`))
			})
		})
	})
})
