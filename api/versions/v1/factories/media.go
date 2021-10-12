package factories

import (
	"github.com/aofdev/fiber-versioning-boilerplate/api/versions/v1/entities"
)

// GetStructMedia is select type struct
func GetStructMedia(MediaType string) interface{} {
	var selectedType interface{}
	switch MediaType {
	case "post":
		selectedType = entities.Post{}
	case "comment":
		selectedType = entities.Comment{}
	case "reply-comment":
		selectedType = entities.ReplyComment{}
	}
	return selectedType
}
