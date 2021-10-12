package repositories

import (
	"github.com/aofdev/fiber-versioning-boilerplate/api/adapters"
	utilities "github.com/aofdev/fiber-versioning-boilerplate/api/utilities"
	"github.com/aofdev/fiber-versioning-boilerplate/api/versions/v2/entities"
	"github.com/aofdev/fiber-versioning-boilerplate/api/versions/v2/factories"
	"go.mongodb.org/mongo-driver/bson"
)

//MediaRepository interface allows us to access the CRUD Operations in mongo here.
type MediaRepository interface {
	GetMediaByID(id string) (*interface{}, error)
	GetMediaByIDS(ids []string) (*[]interface{}, error)
}

type mediaRepository struct {
	database adapters.MongoAdapter
}

//NewMediaRepository is the single instance repo that is being created.
func NewMediaRepository(database adapters.MongoAdapter) MediaRepository {
	return &mediaRepository{database}
}

func (m *mediaRepository) GetMediaByID(id string) (*interface{}, error) {
	doc, err := m.database.FindOne(bson.M{"_id": id, "data_version": "v2.0"})
	if err != nil {
		return nil, err
	}

	var typeMediaFromStruct = entities.MediaType{}
	utilities.ConvertInterfaceToStruct(doc, &typeMediaFromStruct)

	var structByTypeMedia = factories.GetStructMedia(typeMediaFromStruct.Meta.Type)
	utilities.ConvertInterfaceToStruct(doc, &structByTypeMedia)

	return &structByTypeMedia, nil
}

func (m *mediaRepository) GetMediaByIDS(ids []string) (*[]interface{}, error) {
	var result []interface{}

	docs, err := m.database.FindMany(bson.M{"_id": bson.M{"$in": ids}, "data_version": "v2.0"})
	if err != nil {
		return nil, err
	}

	for _, doc := range docs {
		var typeMediaFromStruct = entities.MediaType{}
		utilities.ConvertInterfaceToStruct(doc, &typeMediaFromStruct)

		var structByTypeMedia = factories.GetStructMedia(typeMediaFromStruct.Meta.Type)
		utilities.ConvertInterfaceToStruct(doc, &structByTypeMedia)

		result = append(result, structByTypeMedia)
	}

	return &result, nil
}
