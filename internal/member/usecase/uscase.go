package usecase

type Repository interface {
	Exist(id int) (string, error)
}

type UseCase struct {
	repo Repository
}

func New(repo Repository) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) MemberExists(id int) string {

	result, err := uc.repo.Exist(id)
	if err != nil {
		return err.Error()
	}

	return result
}
