package model

type BaseModelInterface interface {
	GetPointer4DB(string) interface{}
	GetValue4DB(string) interface{}
	GetId() string
	SetId(string)

	GetTableName() string
	GetDBNames() []string
	GetMapNames() []string
	GetValue4Map(string) interface{}
	GetDBFromMapName(string) string
}

func GetPointers4DB(obj BaseModelInterface, names []string) []interface{} {
	pointers := make([]interface{}, 0, 5)
	for _, name := range names {
		pointers = append(pointers, obj.GetPointer4DB(name))
	}
	return pointers
}

func GetValues4DB(obj BaseModelInterface, names []string) []interface{} {
	values := make([]interface{}, 0, 5)
	for _, name := range names {
		values = append(values, obj.GetValue4DB(name))
	}
	return values
}

func GetDBFromMapName(obj BaseModelInterface, name string) string {
	for i, n := range obj.GetMapNames() {
		if name == n {
			return obj.GetDBNames()[i]
		}
	}
	return ""
}