package modules

import "project/internal/storage"

func Setup(s *storage.Storage) *ApiModule {
	repo := SetupRepoModule(s)
	uc := SetupUsecaseModule(repo)
	api := SetupApiModule(uc)
	return api
}

type RepoModule struct {
}

func SetupRepoModule(s *storage.Storage) *RepoModule {
	return &RepoModule{}
}

type UsecaseModule struct {
}

func SetupUsecaseModule(r *RepoModule) *UsecaseModule {
	return &UsecaseModule{}
}

type ApiModule struct {
}

func SetupApiModule(uc *UsecaseModule) *ApiModule {
	return &ApiModule{}
}

func (a *ApiModule) Serve() error {
	return nil
}
