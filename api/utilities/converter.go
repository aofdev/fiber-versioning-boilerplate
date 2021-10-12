package utilities

import jsoniter "github.com/json-iterator/go"

func ConvertInterfaceToStruct(document interface{}, rawStruct interface{}) interface{} {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var temporaryBytes []byte

	temporaryBytes, errMarshal := json.Marshal(document)
	if errMarshal != nil {
		panic(errMarshal)
	}

	errUnmarshal := json.Unmarshal(temporaryBytes, &rawStruct)
	if errUnmarshal != nil {
		panic(errUnmarshal)
	}

	return rawStruct
}
