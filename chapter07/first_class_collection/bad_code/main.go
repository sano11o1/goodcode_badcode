package main

import "errors"

func main() {

}

type Member struct {
	ID int
}

type FieldManager struct {
	Members []Member
}

func NewFiledManager(members []Member) FieldManager {
	return FieldManager{
		Members: members,
	}
}

func (f *FieldManager) Add(member Member) error {
	for _, m := range f.Members {
		if m.ID == member.ID {
			return errors.New("Member already exists")
		}
	}

	if len(f.Members) >= 10 {
		return errors.New("members are full!")
	}

	f.Members = append(f.Members, member)

	return nil
}

type SpecialEventManager struct {
	Members []Member
}

func NewSpecialEventManager(members []Member) SpecialEventManager {
	return SpecialEventManager{
		Members: members,
	}
}

func (f *SpecialEventManager) Add(member Member) error {
	for _, m := range f.Members {
		if m.ID == member.ID {
			return errors.New("Member already exists")
		}
	}
	if len(f.Members) >= 10 {
		return errors.New("members are full!")
	}
	f.Members = append(f.Members, member)

	return nil
}
