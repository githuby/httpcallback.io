package mongo

import (
	"github.com/pjvds/httpcallback.io/model"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type MgoUserRepository struct {
	session  *MgoSession
	database *mgo.Database
}

func NewUserRepository(session *MgoSession) *MgoUserRepository {
	return &MgoUserRepository{
		session:  session,
		database: session.database,
	}
}

func (r *MgoUserRepository) Add(user *model.User) error {
	return r.database.C("Users").Insert(user)
}

func (r *MgoUserRepository) Get(id model.ObjectId) (*model.User, error) {
	query := r.database.C("Users").Find(bson.M{"_id": id})
	var result model.User
	err := query.One(&result)

	return &result, err
}

func (r *MgoUserRepository) GetByAuth(username string, authToken model.AuthenticationToken) (*model.UserAuthInfo, error) {
	query := r.database.C("Users").Find(bson.M{"username": username, "authToken": authToken}).Select(bson.M{"id": 1, "username": 1})
	var result bson.M
	err := query.One(&result)

	return &model.UserAuthInfo{
		UserId:   result["id"].(model.ObjectId),
		Username: result["username"].(string),
	}, err
}

func (r *MgoUserRepository) List() ([]*model.User, error) {
	query := r.database.C("Users").Find(nil)
	var result []*model.User
	err := query.All(&result)
	if result == nil {
		result = make([]*model.User, 0)
	}

	return result, err
}
