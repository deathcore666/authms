package dbclient

import (
	"errors"
	"log"

	"github.com/deathcore666/authms/model"
	"github.com/gocql/gocql"
)

const address string = "127.0.0.1"
const keyspace string = "bships"

//CreateSession lol
func CreateSession(address, keyspace string) (*gocql.Session, error) {
	cluster := gocql.NewCluster(address)
	cluster.Keyspace = keyspace
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	return session, nil
}

//GetUserID lol
func GetUserID(username string) (int, error) {
	session, err := CreateSession(address, keyspace)
	if err != nil {
		return 0, err
	}
	defer session.Close()

	query := "SELECT id FROM users WHERE username = ? ALLOW FILTERING"
	iter := session.Query(query, username).Iter()

	var id int
	if !iter.Scan(&id) {
		return 0, errors.New("iteration error cassandra")
	}
	return id, nil
}

//InsertUser lol
func InsertUser(user model.UserAccount) (int, error) {
	log.Println("attempting to insert a user: ", user.UserName, user.Password)
	session, err := CreateSession(address, keyspace)
	if err != nil {
		return 0, err
	}
	defer session.Close()

	checkQuery := "SELECT userName FROM users WHERE userName = ? ALLOW FILTERING"
	err = session.Query(checkQuery, user.UserName).Exec()
	if err != nil {
		return 0, err
	}
	iter := session.Query(checkQuery, user.UserName).Iter()

	if iter.NumRows() == 0 {
		checkID := "SELECT * FROM users"
		iterID := session.Query(checkID).Iter()
		currentID := iterID.NumRows() + 10000
		query := "INSERT INTO users (id, userName, password) VALUES (?, ?, ?)"
		err = session.Query(query, currentID, user.UserName, user.Password).Exec()
		return currentID, err
	}
	return 0, errors.New("101-username-already-exists")
}

//QueryUser lol
func QueryUser(user model.UserAccount) error {
	session, err := CreateSession(address, keyspace)
	var pwd string
	if err != nil {
		return err
	}
	defer session.Close()

	iter := session.Query("SELECT password FROM users WHERE userName = ? ALLOW FILTERING", user.UserName).Iter()
	if iter.NumRows() == 0 {
		err := errors.New("001-wrong-username")
		return err
	}

	for iter.Scan(&pwd) {
		if pwd == user.Password {
			return nil
		}
	}
	err = errors.New("002-wrong-password")
	return err
}
