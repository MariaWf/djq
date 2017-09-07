package model

type BaseModelInterface interface {
	GetPointer4DB(name string) interface{}
	GetValue4DB(name string) interface{}
	GetId() string
	SetId(id string)
}

func GetPointers4DB(names []string, obj BaseModelInterface) []interface{} {
	pointers := make([]interface{}, 0, 5)
	for _, name := range names {
		pointers = append(pointers, obj.GetPointer4DB(name))
	}
	return pointers
}

func GetValues4DB(names []string, obj BaseModelInterface) []interface{} {
	values := make([]interface{}, 0, 5)
	for _, name := range names {
		values = append(values, obj.GetValue4DB(name))
	}
	return values
}