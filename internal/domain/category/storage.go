package category

type storage interface {
	GetAllCategories(userId int) ([]*Category, error)
	CreateCategory(userId int, category *Category) (*Category, error)
	UpdateCategory(userId int, category *Category) (*Category, error)
	DeleteCategory(userId, categoryId int) (*Category, error)
}
