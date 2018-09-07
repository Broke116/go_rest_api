package service

import (
	"fmt"
	"go_rest_api/app/config/db/mongo"
	"go_rest_api/app/config/db/mongo/dao"
	"go_rest_api/app/model"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MemberService is used store collection information.
type MemberService struct {
	collection *mgo.Collection
	//hash       security.Hash
}

// MemberServiceConstructor is an instance of the MemberService
func MemberServiceConstructor(session *mongo.Session, database string, collectionName string /*hash security.Hash*/) *MemberService {
	collection := session.GetCollection(database, collectionName)
	collection.EnsureIndex(dao.MemberModelIndex())
	return &MemberService{collection /*hash*/}
}

// GetMembers is a method of MemberService
func (ms *MemberService) GetMembers() (*model.Members, error) {
	members := model.Members{}
	err := ms.collection.Find(bson.M{}).All(&members)
	if err != nil {
		fmt.Println("Error was occured when fetching the list of members ", err)
	}

	return &members, err
}

// GetMemberByID is a method of MemberService
func (ms *MemberService) GetMemberByID(id string) (*model.Member, error) {
	member := dao.MemberModel{}
	err := ms.collection.FindId(bson.ObjectIdHex(id)).One(&member)
	return member.ToMember(), err
}

// InsertMember is a method of MemberService
func (ms *MemberService) InsertMember(m *model.Member) error {
	member := dao.NewMemberModel(m)
	fmt.Println("Member service insert member ", member)

	//member.Password = ms.hash.Generate(m.Password)

	return ms.collection.Insert(&member)
}

// UpdateMember is a method of MemberService
func (ms *MemberService) UpdateMember(m *model.Member, _id string) error {
	member := dao.NewMemberModel(m)
	//member.Password = ms.hash.Generate(m.Password)
	return ms.collection.UpdateId(bson.ObjectIdHex(_id), member)
}

// DeleteMember is a method of MemberService
func (ms *MemberService) DeleteMember(memberID string) error {
	return ms.collection.RemoveId(bson.ObjectIdHex(memberID))
}
