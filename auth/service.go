package auth

type serviceQueries struct {
}

type serviceMutations struct {
}

var (
	_ controllerQueries   = (*controller)(nil)
	_ controllerMutations = (*controller)(nil)
)

type service struct {
	repo    *repo
	session *session
}

func newService() *service {
	return &service{
		repo:    newRepo(),
		session: newSession(),
	}
}
