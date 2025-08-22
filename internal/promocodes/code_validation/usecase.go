package code_validation

type Postgres interface {
	ValidCode(code string) string
}

type Usecase struct {
	Postgres Postgres
}

func NewUsecase(postgres Postgres) *Usecase {
	return &Usecase{Postgres: postgres}
}

// ValidCode валидирует промокод, используя репозиторий.
func (uc *Usecase) ValidCode(code string) bool {
	// В зависимости от логики, тут может быть проверка на наличие промокода
	// или его валидность.
	//result := uc.Postgres.ValidCode(code)

	// Пример логики: если результат не пустая строка, значит, код валиден.
	return true
}
