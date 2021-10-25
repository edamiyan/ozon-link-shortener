package in_memory

type Repository struct {
	briefToFull map[string]string
	fullToBrief map[string]string
}

func NewRepository(n int) *Repository {
	briefToFull := make(map[string]string, n)
	fullToBrief := make(map[string]string, n)
	return &Repository{
		briefToFull: briefToFull,
		fullToBrief: fullToBrief,
	}
}
