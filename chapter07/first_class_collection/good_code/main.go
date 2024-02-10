package main

import (
	"fmt"

	"github.com/sano11o1/goodcode_badcode/chapter07/first_class_collection/good_code/model"
)

func main() {
	members := []model.Member{
		{ID: 1},
		{ID: 2},
		{ID: 3},
	}
	party := model.NewParty(members)
	party, _ = party.Add(model.Member{ID: 4})
	fmt.Printf("%+v\n", party.GetMembers())
}
