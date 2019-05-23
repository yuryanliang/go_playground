package cart_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/yuryanliang/go_playground/ginkgo/cart"
)

var _ = Describe("Shopping cart", func() {
	itemA := Item{
		ID:    "itemA",
		Name:  "Item A",
		Price: 10.20,
		Qty:   0,
	}
	itemB := Item{
		ID:    "itemB",
		Name:  "Item B",
		Price: 7.66,
		Qty:   0,
	}

	//var cart = Cart{}

	Context("initially", func() {

		cart := Cart{}
		It("has 0 items", func() {
			Expect(cart.TotalUniqueItems()).Should(BeZero())
		})
		It("has 0 units", func() {
			Expect(cart.TotalUnit()).Should(BeZero())
		})
		Specify("the total amount is 0.00", func() {
			Expect(cart.TotalAmount()).Should(BeZero())

		})
	})
	Context("when a new item is added", func() {
		cart := Cart{}
		originalItemCount := cart.TotalUniqueItems()
		originalUnitCount := cart.TotalUnit()
		originalAmount := cart.TotalAmount()

		cart.AddItem(itemA)

		Context("the shopping cart", func() {
			It("has 1 more unique item than it had earlier", func() {
				Expect(cart.TotalUniqueItems()).Should(Equal(originalItemCount + 1))
			})
			It("has 1 more unit than it had earlier", func() {
				Expect(cart.TotalUnit()).Should(Equal(originalUnitCount + 1))
			})
			Specify("the total amount increases by item price", func() {
				Expect(cart.TotalAmount()).Should(Equal(originalAmount + itemA.Price))
			})
		})
	})

	Context("when an existing item is added", func() {
		cart := Cart{}
		cart.AddItem(itemA)
		originalItemCount := cart.TotalUniqueItems()
		originalUnitCount := cart.TotalUnit()
		originalAmount := cart.TotalAmount()

		cart.AddItem(itemA)

		Context("the shopping cart", func() {
			It("has the same number of unique items as earlier", func() {
				Expect(cart.TotalUniqueItems()).Should(Equal(originalItemCount))
			})
			It("has 1 more unit than it had earlier", func() {
				Expect(cart.TotalUnit()).Should(Equal(originalUnitCount + 1))
			})
			Specify("total amount increases by item price", func() {
				Expect(cart.TotalAmount()).Should(Equal(originalAmount + itemA.Price))
			})
		})
	})

	Context("that has 0 unit of item A", func() {
		cart := Cart{}

		cart.AddItem(itemB)
		cart.AddItem(itemB)

		originalItemCount := cart.TotalUniqueItems()
		originalUnitCount := cart.TotalUnit()
		originalAmount := cart.TotalAmount()

		Context("removing item A", func() {
			cart.RemoveItem("itemA", 1)

			It("should not change the number of items", func() {
				Expect(cart.TotalUniqueItems()).Should(Equal(originalItemCount))
			})
			It("should not change the number of units", func() {
				Expect(cart.TotalUnit()).Should(Equal(originalUnitCount))
			})
			It("should not change the amount", func() {
				Expect(cart.TotalAmount()).Should(Equal(originalAmount))
			})
		})
	})

	Context("that has 1 unit of item A", func() {
		cart := Cart{}

		cart.AddItem(itemB)
		cart.AddItem(itemA)
		originalItemCount := cart.TotalUniqueItems()
		originalUnitCount := cart.TotalUnit()
		originalAmount := cart.TotalAmount()

		Context("removing 1 unit item A", func() {
			cart.RemoveItem("itemA", 1)
			It("should reduce the number of items by 1", func() {
				Expect(cart.TotalUniqueItems()).Should(Equal(originalItemCount - 1))
			})
			It("should reduce the number of units by 1", func() {
				Expect(cart.TotalUnit()).Should(Equal(originalUnitCount - 1))
			})
			It("should reduce the amount by item price", func() {
				Expect(cart.TotalAmount()).Should(Equal(originalAmount - itemA.Price))
			})
		})
	})

	Context("that has 2 units of item A", func() {

		Context("removing 1 unit of item A", func() {
			cart := Cart{}
			cart.AddItem(itemA)
			cart.AddItem(itemA)
			originalItemCount := cart.TotalUniqueItems()
			originalUnitCount := cart.TotalUnit()
			originalAmount := cart.TotalAmount()
			cart.RemoveItem(itemA.ID, 1)
			It("should not reduce the number of items", func() {
				Expect(cart.TotalUniqueItems()).Should(Equal(originalItemCount))
			})
			It("should reduce the number of units by 1", func() {
				Expect(cart.TotalUnit()).Should(Equal(originalUnitCount - 1))
			})
			It("should reduce the amount by the item price", func() {
				Expect(cart.TotalAmount()).Should(Equal(originalAmount - itemA.Price))
			})
		})

		Context("removing 2 units of item A", func() {
			cart := Cart{}
			cart.AddItem(itemA)
			cart.AddItem(itemA)
			originalItemCount := cart.TotalUniqueItems()
			originalUnitCount := cart.TotalUnit()
			originalAmount := cart.TotalAmount()
			cart.RemoveItem("itemA", 2)

			It("should reduce the number of items by 1", func() {
				Expect(cart.TotalUniqueItems()).Should(Equal(originalItemCount - 1))

			})
			It("should reduce the number of units by 2", func() {
				Expect(cart.TotalUnit()).Should(Equal(originalUnitCount - 2))

			})
			It("should reduce the amount by twice the item price", func() {
				Expect(cart.TotalAmount()).Should(Equal(originalAmount - itemA.Price*2))

			})
		})
	})
})
