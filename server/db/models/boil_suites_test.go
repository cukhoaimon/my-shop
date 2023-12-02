// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dbmodels

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("BookCategories", testBookCategories)
	t.Run("Books", testBooks)
	t.Run("Discounts", testDiscounts)
	t.Run("Orders", testOrders)
	t.Run("SchemaMigrations", testSchemaMigrations)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("BookCategories", testBookCategoriesDelete)
	t.Run("Books", testBooksDelete)
	t.Run("Discounts", testDiscountsDelete)
	t.Run("Orders", testOrdersDelete)
	t.Run("SchemaMigrations", testSchemaMigrationsDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("BookCategories", testBookCategoriesQueryDeleteAll)
	t.Run("Books", testBooksQueryDeleteAll)
	t.Run("Discounts", testDiscountsQueryDeleteAll)
	t.Run("Orders", testOrdersQueryDeleteAll)
	t.Run("SchemaMigrations", testSchemaMigrationsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("BookCategories", testBookCategoriesSliceDeleteAll)
	t.Run("Books", testBooksSliceDeleteAll)
	t.Run("Discounts", testDiscountsSliceDeleteAll)
	t.Run("Orders", testOrdersSliceDeleteAll)
	t.Run("SchemaMigrations", testSchemaMigrationsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("BookCategories", testBookCategoriesExists)
	t.Run("Books", testBooksExists)
	t.Run("Discounts", testDiscountsExists)
	t.Run("Orders", testOrdersExists)
	t.Run("SchemaMigrations", testSchemaMigrationsExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("BookCategories", testBookCategoriesFind)
	t.Run("Books", testBooksFind)
	t.Run("Discounts", testDiscountsFind)
	t.Run("Orders", testOrdersFind)
	t.Run("SchemaMigrations", testSchemaMigrationsFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("BookCategories", testBookCategoriesBind)
	t.Run("Books", testBooksBind)
	t.Run("Discounts", testDiscountsBind)
	t.Run("Orders", testOrdersBind)
	t.Run("SchemaMigrations", testSchemaMigrationsBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("BookCategories", testBookCategoriesOne)
	t.Run("Books", testBooksOne)
	t.Run("Discounts", testDiscountsOne)
	t.Run("Orders", testOrdersOne)
	t.Run("SchemaMigrations", testSchemaMigrationsOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("BookCategories", testBookCategoriesAll)
	t.Run("Books", testBooksAll)
	t.Run("Discounts", testDiscountsAll)
	t.Run("Orders", testOrdersAll)
	t.Run("SchemaMigrations", testSchemaMigrationsAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("BookCategories", testBookCategoriesCount)
	t.Run("Books", testBooksCount)
	t.Run("Discounts", testDiscountsCount)
	t.Run("Orders", testOrdersCount)
	t.Run("SchemaMigrations", testSchemaMigrationsCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("BookCategories", testBookCategoriesHooks)
	t.Run("Books", testBooksHooks)
	t.Run("Discounts", testDiscountsHooks)
	t.Run("Orders", testOrdersHooks)
	t.Run("SchemaMigrations", testSchemaMigrationsHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("BookCategories", testBookCategoriesInsert)
	t.Run("BookCategories", testBookCategoriesInsertWhitelist)
	t.Run("Books", testBooksInsert)
	t.Run("Books", testBooksInsertWhitelist)
	t.Run("Discounts", testDiscountsInsert)
	t.Run("Discounts", testDiscountsInsertWhitelist)
	t.Run("Orders", testOrdersInsert)
	t.Run("Orders", testOrdersInsertWhitelist)
	t.Run("SchemaMigrations", testSchemaMigrationsInsert)
	t.Run("SchemaMigrations", testSchemaMigrationsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("BookToBookCategoryUsingCategory", testBookToOneBookCategoryUsingCategory)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("BookCategoryToCategoryBooks", testBookCategoryToManyCategoryBooks)
	t.Run("BookToOrders", testBookToManyOrders)
	t.Run("DiscountToOrders", testDiscountToManyOrders)
	t.Run("OrderToBooks", testOrderToManyBooks)
	t.Run("OrderToDiscounts", testOrderToManyDiscounts)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("BookToBookCategoryUsingCategoryBooks", testBookToOneSetOpBookCategoryUsingCategory)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("BookCategoryToCategoryBooks", testBookCategoryToManyAddOpCategoryBooks)
	t.Run("BookToOrders", testBookToManyAddOpOrders)
	t.Run("DiscountToOrders", testDiscountToManyAddOpOrders)
	t.Run("OrderToBooks", testOrderToManyAddOpBooks)
	t.Run("OrderToDiscounts", testOrderToManyAddOpDiscounts)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("BookToOrders", testBookToManySetOpOrders)
	t.Run("DiscountToOrders", testDiscountToManySetOpOrders)
	t.Run("OrderToBooks", testOrderToManySetOpBooks)
	t.Run("OrderToDiscounts", testOrderToManySetOpDiscounts)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("BookToOrders", testBookToManyRemoveOpOrders)
	t.Run("DiscountToOrders", testDiscountToManyRemoveOpOrders)
	t.Run("OrderToBooks", testOrderToManyRemoveOpBooks)
	t.Run("OrderToDiscounts", testOrderToManyRemoveOpDiscounts)
}

func TestReload(t *testing.T) {
	t.Run("BookCategories", testBookCategoriesReload)
	t.Run("Books", testBooksReload)
	t.Run("Discounts", testDiscountsReload)
	t.Run("Orders", testOrdersReload)
	t.Run("SchemaMigrations", testSchemaMigrationsReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("BookCategories", testBookCategoriesReloadAll)
	t.Run("Books", testBooksReloadAll)
	t.Run("Discounts", testDiscountsReloadAll)
	t.Run("Orders", testOrdersReloadAll)
	t.Run("SchemaMigrations", testSchemaMigrationsReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("BookCategories", testBookCategoriesSelect)
	t.Run("Books", testBooksSelect)
	t.Run("Discounts", testDiscountsSelect)
	t.Run("Orders", testOrdersSelect)
	t.Run("SchemaMigrations", testSchemaMigrationsSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("BookCategories", testBookCategoriesUpdate)
	t.Run("Books", testBooksUpdate)
	t.Run("Discounts", testDiscountsUpdate)
	t.Run("Orders", testOrdersUpdate)
	t.Run("SchemaMigrations", testSchemaMigrationsUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("BookCategories", testBookCategoriesSliceUpdateAll)
	t.Run("Books", testBooksSliceUpdateAll)
	t.Run("Discounts", testDiscountsSliceUpdateAll)
	t.Run("Orders", testOrdersSliceUpdateAll)
	t.Run("SchemaMigrations", testSchemaMigrationsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}