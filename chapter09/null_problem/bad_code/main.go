package main

func main() {
}

type Equipment struct {
	Name    string
	Defence int
}

type Member struct {
	Head    *Equipment
	Body    *Equipment
	Arm     *Equipment
	Defence int
}

func (m Member) TotalDefence() int {
	d := m.Defence
	if m.Head != nil {
		d += m.Head.Defence
	}
	if m.Body != nil {
		d += m.Body.Defence
	}
	if m.Arm != nil {
		d += m.Arm.Defence
	}
	return d
}

func (m *Member) takeOfAllEquipment() {
	m.Head = nil
	m.Body = nil
	m.Arm = nil
}
