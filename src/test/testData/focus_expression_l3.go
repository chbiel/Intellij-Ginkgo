package test

var _ = Describe("Book", func() {
	Describe("Categorizing book length", func() {
		Context("With more than 300 pages", func() {
			<caret>It("should be a novel", func() {
				Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
			})
		})
	})
})
