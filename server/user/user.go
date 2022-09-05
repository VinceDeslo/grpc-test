package user

import (
	"encoding/json"
	"log"

	"github.com/VinceDeslo/grpc-test/server/models"
	"github.com/VinceDeslo/grpc-test/server/user/pb"
	"github.com/VinceDeslo/grpc-test/server/utils"
)

type UserDb struct {
	users        []models.User
	usersPayload []*pb.User
	data         []byte
}

func NewUserDb() *UserDb {
	return &UserDb{}
}

func (db *UserDb) LoadUsersData(fp string) {
	bytes, err := utils.ReadJsonFile(fp)
	if err != nil {
		log.Fatalln("Failed to fetch data from User database.")
	}
	db.data = bytes
}

func (db *UserDb) GetUsersData() []byte {
	return db.data
}

func (db *UserDb) GetUsersPayload() []*pb.User {
	var payload []*pb.User
	json.Unmarshal(db.data, &db.users)

	for _, user := range db.users {
		payload = append(payload, &pb.User{
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Age:         user.Age,
			Permissions: user.Permissions,
		})
	}
	db.usersPayload = payload
	return db.usersPayload
}
