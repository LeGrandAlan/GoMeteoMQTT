package airportspackage

import (
	"fmt"
)

type Airport struct {
	Id   string
	Name string
}

func (o Airport) String() string {

	return fmt.Sprintf("{ Id: %s, Name: %s }",
		o.Id, o.Name)

}

func MakeFromMap(m map[string]interface{}) Airport {
	res := Airport{
		Id:   m["Id"].(string),
		Name: m["Name"].(string),
	}
	return res
}
