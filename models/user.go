package models

import "time"

type User struct {
	Id        int64  `db:"id" json:"id" facebook:"-"`
	Uid       string `db:"uid" json:"uid" facebook:"id"`
	CreatedAt int64  `db:"created_at" json:"created_at" facebook:"-"`
	Email     string `db:"email" json:"email"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
}

func FindOrCreateUser(user *User) {
	existing := &User{}
	err := DbMap.SelectOne(existing, "SELECT * FROM users WHERE uid = ?", user.Uid)

	if err != nil {
		// create new user
		user.CreatedAt = time.Now().UnixNano()
		DbMap.Insert(user)
	} else {
		// copy existing attributes
		user.Id = existing.Id
		user.CreatedAt = existing.CreatedAt
		DbMap.Update(user)
	}
}
