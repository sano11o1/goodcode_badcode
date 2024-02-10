package model

import "errors"

type Member struct {
	ID int
}

type Party struct {
	members []Member // 別パッケージからの変更を防ぐために小文字にする　取得はGetMembersメソッドを使う
}

func NewParty(members []Member) Party {
	return Party{
		members: members,
	}
}

func (p Party) Add(member Member) (Party, error) {
	if p.isExist(member) {
		return Party{}, errors.New("Member already exists")
	}
	if p.isFull() {
		return Party{}, errors.New("members are full!")
	}
	// 副作用を抑えるために値型を返す。レシーバの構造体には影響がない。
	p.members = append(p.members, member)
	return p, nil
}

func (p Party) isFull() bool {
	return len(p.members) >= 10
}

func (p Party) isExist(member Member) bool {
	for _, m := range p.members {
		if m.ID == member.ID {
			return true
		}
	}
	return false
}

// 別パケージからMembersを取得するためのメソッド.
func (p Party) GetMembers() []Member {
	return p.members
}
