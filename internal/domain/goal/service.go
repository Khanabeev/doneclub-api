package goal

type goalService struct {
	storage Storage
}

func NewGoalService(storage Storage) *goalService {
	return &goalService{
		storage: storage,
	}
}
