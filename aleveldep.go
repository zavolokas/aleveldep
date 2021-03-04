package aleveldep

import (
	"fmt"

	"github.com/zavolokas/bleveldep"
)

func GetId() string {
	return fmt.Sprintf("%s and %s", getId(), bleveldep.GetId())
}

func getId() string {
	return "this is a github aleveldep"
}
