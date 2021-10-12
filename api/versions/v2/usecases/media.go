package usecases

import (
	repositories "github.com/aofdev/fiber-versioning-boilerplate/api/versions/v2/repositories"
)

//MediaUsecase is an interface from which our api module can access our repository of all our models
type MediaUsecase interface {
	GetMediaByID(id string) (*interface{}, error)
	GetMediaByIDS(ids []string) (*[]interface{}, error)
}

type mediaUsecase struct {
	mediaRepo repositories.MediaRepository
}

//NewMediaUsecase is used to create a single instance of the usecase
func NewMediaUsecase(repo repositories.MediaRepository) MediaUsecase {
	return &mediaUsecase{
		mediaRepo: repo,
	}
}

func (m *mediaUsecase) GetMediaByID(id string) (*interface{}, error) {
	return m.mediaRepo.GetMediaByID(id)
}

func (m *mediaUsecase) GetMediaByIDS(ids []string) (*[]interface{}, error) {
	return m.mediaRepo.GetMediaByIDS(ids)
}
