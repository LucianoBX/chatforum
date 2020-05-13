package models

import (
	"time"
)

type Session struct {
	Id          int
	Uuid, Email string
	UserId      int
	CreatedAt   time.Time
}

// 检查session在数据库中有效性
func (s *Session) Check() (valid bool, err error) {
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = ?", s.Uuid).Scan(&s.Id, &s.Uuid, &s.Email, &s.UserId, &s.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if s.Id != 0 {
		valid = true
	}
	return
}

// 删除线程
func (s *Session) DeleteByUUID() (err error) {
	statement := "delete from sessions where uuid = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(s.Uuid)
	return
}

// 从Session中获取用户
func (s *Session) User() (u User, err error) {
	u = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = ?", s.UserId).Scan(&u.Id, &u.Uuid, &u.Email, &u.CreatedAt)
	return
}

// 从数据库中删除所有session
func SessionDeleteAll() (err error) {
	statement := "delete from sessions"
	_, err = Db.Exec(statement)
	return
}
