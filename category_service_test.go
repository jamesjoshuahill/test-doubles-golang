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

var _ = Describe("CategoryService#First", func() {
	Context("with a hand-rolled fake logger", func() {
		It("calls the repository with the category", func() {
			repo := new(handrolled.FakeRepository)
			service := doubles.NewCategoryService(repo)

			service.First("name")

			Expect(repo.QueryMock.CallCount).To(Equal(1))
			Expect(repo.QueryMock.Received.Name).To(Equal("name"))
			Expect(repo.QueryMock.Received.Kind).To(Equal("category"))
		})

		Context("when no records are found", func() {
			It("returns an error", func() {
				repo := new(handrolled.FakeRepository)
				service := doubles.NewCategoryService(repo)

				_, err := service.First("name")

				Expect(err).To(MatchError(`no categories found with name "name"`))
			})
		})

		Context("when one matching record is found", func() {
			It("returns the matching record", func() {
				repo := new(handrolled.FakeRepository)
				repo.QueryMock.Returns.Records = []doubles.Record{
					{Category: "name"},
				}
				service := doubles.NewCategoryService(repo)

				record, err := service.First("name")

				Expect(err).NotTo(HaveOccurred())
				Expect(record.Category).To(Equal("name"))
			})
		})

		Context("when several records are found", func() {
			It("returns the first matching record", func() {
				repo := new(handrolled.FakeRepository)
				repo.QueryMock.Returns.Records = []doubles.Record{
					{Category: "name"},
					{Category: "name"},
					{Category: "name"},
				}
				service := doubles.NewCategoryService(repo)

				record, err := service.First("name")

				Expect(err).NotTo(HaveOccurred())
				Expect(record.Category).To(Equal("name"))
			})
		})

		Context("when finding records fails", func() {
			It("returns an error", func() {
				repo := new(handrolled.FakeRepository)
				repo.QueryMock.Returns.Error = errors.New("find failed")
				service := doubles.NewCategoryService(repo)

				_, err := service.First("name")

				Expect(err).To(MatchError(`finding categories with name "name" failed: find failed`))
			})
		})
	})

	Context("with a charlatan fake logger", func() {
		It("calls the repository with the category", func() {
			repo := new(charlatan.Fakerepository)
			repo.SetQueryStub(nil, nil)
			service := doubles.NewCategoryService(repo)

			service.First("name")

			Expect(repo.QueryCalledOnceWith("name", "category")).To(BeTrue())
		})

		Context("when no records are found", func() {
			It("returns an error", func() {
				repo := new(charlatan.Fakerepository)
				repo.SetQueryStub([]doubles.Record{}, nil)
				service := doubles.NewCategoryService(repo)

				_, err := service.First("name")

				Expect(err).To(MatchError(`no categories found with name "name"`))
			})
		})

		Context("when one matching record is found", func() {
			It("returns the matching record", func() {
				repo := new(charlatan.Fakerepository)
				repo.SetQueryStub([]doubles.Record{
					{Category: "name"},
				}, nil)
				service := doubles.NewCategoryService(repo)

				record, err := service.First("name")

				Expect(err).NotTo(HaveOccurred())
				Expect(record.Category).To(Equal("name"))
			})
		})

		Context("when several records are found", func() {
			It("returns the first matching record", func() {
				repo := new(charlatan.Fakerepository)
				repo.SetQueryStub([]doubles.Record{
					{Category: "name"},
					{Category: "name"},
					{Category: "name"},
				}, nil)
				service := doubles.NewCategoryService(repo)

				record, err := service.First("name")

				Expect(err).NotTo(HaveOccurred())
				Expect(record.Category).To(Equal("name"))
			})
		})

		Context("when finding records fails", func() {
			It("returns an error", func() {
				repo := new(charlatan.Fakerepository)
				repo.SetQueryStub(nil, errors.New("find failed"))
				service := doubles.NewCategoryService(repo)

				_, err := service.First("name")

				Expect(err).To(MatchError(`finding categories with name "name" failed: find failed`))
			})
		})
	})

	Context("with a counterfeiter fake logger", func() {
		It("calls the repository with the category", func() {
			repo := new(counterfeiter.FakeRepository)
			service := doubles.NewCategoryService(repo)

			service.First("name")

			Expect(repo.QueryCallCount()).To(Equal(1))
			name, kind := repo.QueryArgsForCall(0)
			Expect(name).To(Equal("name"))
			Expect(kind).To(Equal("category"))
		})

		Context("when no records are found", func() {
			It("returns an error", func() {
				repo := new(counterfeiter.FakeRepository)
				repo.QueryReturns([]doubles.Record{}, nil)
				service := doubles.NewCategoryService(repo)

				_, err := service.First("name")

				Expect(err).To(MatchError(`no categories found with name "name"`))
			})
		})

		Context("when one matching record is found", func() {
			It("returns the matching record", func() {
				repo := new(counterfeiter.FakeRepository)
				repo.QueryReturns([]doubles.Record{
					{Category: "name"},
				}, nil)
				service := doubles.NewCategoryService(repo)

				record, err := service.First("name")

				Expect(err).NotTo(HaveOccurred())
				Expect(record.Category).To(Equal("name"))
			})
		})

		Context("when several records are found", func() {
			It("returns the first matching record", func() {
				repo := new(counterfeiter.FakeRepository)
				repo.QueryReturns([]doubles.Record{
					{Category: "name"},
					{Category: "name"},
					{Category: "name"},
				}, nil)
				service := doubles.NewCategoryService(repo)

				record, err := service.First("name")

				Expect(err).NotTo(HaveOccurred())
				Expect(record.Category).To(Equal("name"))
			})
		})

		Context("when finding records fails", func() {
			It("returns an error", func() {
				repo := new(counterfeiter.FakeRepository)
				repo.QueryReturns(nil, errors.New("find failed"))
				service := doubles.NewCategoryService(repo)

				_, err := service.First("name")

				Expect(err).To(MatchError(`finding categories with name "name" failed: find failed`))
			})
		})
	})

	Context("with a minimock mock logger", func() {
		It("calls the repository with the category", func() {
			mc := minimock.NewController(GinkgoT())
			repo := minimocks.NewrepositoryMock(mc)
			repo.QueryMock.Expect("name", "category").Return(nil, nil)
			service := doubles.NewCategoryService(repo)

			service.First("name")

			Expect(repo.QueryCounter).To(Equal(uint64(1)))
			mc.Finish()
		})

		Context("when no records are found", func() {
			It("returns an error", func() {
				mc := minimock.NewController(GinkgoT())
				repo := minimocks.NewrepositoryMock(mc)
				repo.QueryMock.Return([]doubles.Record{}, nil)
				service := doubles.NewCategoryService(repo)

				_, err := service.First("name")

				Expect(err).To(MatchError(`no categories found with name "name"`))
				mc.Finish()
			})
		})

		Context("when one matching record is found", func() {
			It("returns the matching record", func() {
				mc := minimock.NewController(GinkgoT())
				repo := minimocks.NewrepositoryMock(mc)
				repo.QueryMock.Return([]doubles.Record{
					{Category: "name"},
				}, nil)
				service := doubles.NewCategoryService(repo)

				record, err := service.First("name")

				Expect(err).NotTo(HaveOccurred())
				Expect(record.Category).To(Equal("name"))
				mc.Finish()
			})
		})

		Context("when several records are found", func() {
			It("returns the first matching record", func() {
				mc := minimock.NewController(GinkgoT())
				repo := minimocks.NewrepositoryMock(mc)
				repo.QueryMock.Return([]doubles.Record{
					{Category: "name"},
					{Category: "name"},
					{Category: "name"},
				}, nil)
				service := doubles.NewCategoryService(repo)

				record, err := service.First("name")

				Expect(err).NotTo(HaveOccurred())
				Expect(record.Category).To(Equal("name"))
				mc.Finish()
			})
		})

		Context("when finding records fails", func() {
			It("returns an error", func() {
				mc := minimock.NewController(GinkgoT())
				repo := minimocks.NewrepositoryMock(mc)
				repo.QueryMock.Return(nil, errors.New("find failed"))
				service := doubles.NewCategoryService(repo)

				_, err := service.First("name")

				Expect(err).To(MatchError(`finding categories with name "name" failed: find failed`))
				mc.Finish()
			})
		})
	})

	Context("with a pegomock mock logger", func() {
		It("calls the repository with the category", func() {
			repo := pegomock.NewMockrepository()
			service := doubles.NewCategoryService(repo)

			service.First("name")

			repo.VerifyWasCalledOnce().Query("name", "category")
		})

		Context("when no records are found", func() {
			It("returns an error", func() {
				repo := pegomock.NewMockrepository()
				When(repo.Query(AnyString(), AnyString())).ThenReturn([]doubles.Record{}, nil)
				service := doubles.NewCategoryService(repo)

				_, err := service.First("name")

				Expect(err).To(MatchError(`no categories found with name "name"`))
			})
		})

		Context("when one matching record is found", func() {
			It("returns the matching record", func() {
				repo := pegomock.NewMockrepository()
				When(repo.Query(AnyString(), AnyString())).ThenReturn([]doubles.Record{
					{Category: "name"},
				}, nil)
				service := doubles.NewCategoryService(repo)

				record, err := service.First("name")

				Expect(err).NotTo(HaveOccurred())
				Expect(record.Category).To(Equal("name"))
			})
		})

		Context("when several records are found", func() {
			It("returns the first matching record", func() {
				repo := pegomock.NewMockrepository()
				When(repo.Query(AnyString(), AnyString())).ThenReturn([]doubles.Record{
					{Category: "name"},
					{Category: "name"},
					{Category: "name"},
				}, nil)
				service := doubles.NewCategoryService(repo)

				record, err := service.First("name")

				Expect(err).NotTo(HaveOccurred())
				Expect(record.Category).To(Equal("name"))
			})
		})

		Context("when finding records fails", func() {
			It("returns an error", func() {
				repo := pegomock.NewMockrepository()
				When(repo.Query(AnyString(), AnyString())).ThenReturn(nil, errors.New("find failed"))
				service := doubles.NewCategoryService(repo)

				_, err := service.First("name")

				Expect(err).To(MatchError(`finding categories with name "name" failed: find failed`))
			})
		})
	})
})
