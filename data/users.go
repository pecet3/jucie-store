package data

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
)

const UsersTable = `
create table if not exists users (
	id integer primary key autoincrement,
	uuid text not null,
	name text not null,
	email text not null,
	password text not null,
	salt text not null,
	image_url text default '',
	is_active bool default false,
	created_at timestamp default current_timestamp
);
`

type User struct {
	Id        int    `json:"-"`
	Uuid      string `json:"uuid"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	Salt      string `json:"-"`
	ImageURL  string `json:"image_url"`
	CreatedAt string `json:"created_at"`
	IsActive  bool   `json:"-"`
}

func (u User) GetByEmail(db *sql.DB, email string) (*User, error) {
	query := "select id, name, email, password, salt, image_url, is_active from users where email = ?"
	err := db.QueryRow(query, email).Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.Salt, &u.ImageURL, &u.IsActive)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &u, nil
}

func (u User) GetById(db *sql.DB, id int) (*User, error) {
	query := "select id, name, email, password, salt, is_active from users where id = ?"
	err := db.QueryRow(query, id).Scan(&u.Id, &u.Name, &u.Password, &u.Password, &u.Salt, &u.IsActive)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &u, nil
}

func (u User) Add(db *sql.DB, name, email, password, salt string) (int, error) {
	uuid := uuid.NewString()
	query := "insert into users (uuid, name, email, password, salt) values (?,?,?,?,?)"

	result, err := db.Exec(query, uuid, name, email, password, salt)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	userId, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return int(userId), nil
}
