package models

import (
	"time"
)


type Post struct {
	Id               int
	Uuid, Body       string
	UserId, ThreadId int
	CreatedAt        time.Time
}

func (p *Post) CreatedAtDate() string {
	return p.CreatedAt.Format("Jan 2, 2020 at 3:43pm")
}

// 获得发表用户
func (p *Post) User() (u User) {
	u = User{}
	Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = ?", p.UserId).
		Scan(&u.Id, &u.Uuid, &u.Name, &u.Email, &u.CreatedAt)
	return
}
