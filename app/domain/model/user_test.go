package model_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/architect-basis/app/domain/model"
)

var _ = Describe("User", func() {
	var (
		user model.User
	)

	BeforeEach(func() {
		user = model.User{
			Name:  "John Doe",
			Email: "john.doe@example.com",
		}
	})

	It("should create a user", func() {
		Expect(user.Name).To(Equal("John Doe"))
		Expect(user.Email).To(Equal("john.doe@example.com"))
	})
})
