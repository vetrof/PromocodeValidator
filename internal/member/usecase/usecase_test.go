// file: usecase/usecase_test.go
package usecase

import "testing"

// самый простой мок: всегда возвращает "exists", nil
type mockRepoOK struct{}

func (m *mockRepoOK) Exist(id int) (string, error) {
	// игнорируем id — для первого теста нам это не важно
	return "exists", nil
}

func TestMemberExists_ReturnsValueFromRepo(t *testing.T) {
	// Arrange: собираем UseCase с нашим мок-репозиторием
	uc := New(&mockRepoOK{})

	// Act: вызываем тестируемый метод
	got := uc.MemberExists(123)

	// Assert: проверяем результат
	want := "exists"
	if got != want {
		t.Fatalf("ожидали %q, получили %q", want, got)
	}
}
