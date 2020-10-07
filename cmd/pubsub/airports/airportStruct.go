package airportspackage

import (
	"fmt"
)

type Airport struct {
	Id   string
	Name string
	City string
}

func (o Airport) String() string {

	return fmt.Sprintf("{ Id: %s, Name: %s, City: %s }",
		o.Id, o.Name)

}

func MakeFromMap(m map[string]interface{}) Airport {
	res := Airport{
		Id:   m["Id"].(string),
		Name: m["Name"].(string),
		City: m["City"].(string),
	}
	return res
}
