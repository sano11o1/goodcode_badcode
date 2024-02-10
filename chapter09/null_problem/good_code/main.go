package main

import "errors"

func main() {

}

type EquipmentName string

const EQUIPMENT_NOTING EquipmentName = "Noting"
const EQUIPMENT_HEAD EquipmentName = "Iron Helmet"

type Equipment struct {
	Name    EquipmentName
	Defence int
}

func NewEquipment(name EquipmentName, defence int) (Equipment, error) {
	if name == "" {
		return Equipment{}, errors.New("name is empty")
	}
	return Equipment{
		Name:    name,
		Defence: defence,
	}, nil
}

func NewNotingEquipment() Equipment {
	return Equipment{Name: "Noting", Defence: 0}
}

type Member struct {
	Head    Equipment
	Body    Equipment
	Arm     Equipment
	Defence int
}

func (m Member) TotalDefence() int {
	d := m.Defence
	d += m.Head.Defence
	d += m.Body.Defence
	d += m.Arm.Defence
	return d
}

func (m *Member) takeOfAllEquipment() {
	m.Head = NewNotingEquipment()
	m.Body = NewNotingEquipment()
	m.Arm = NewNotingEquipment()
}
