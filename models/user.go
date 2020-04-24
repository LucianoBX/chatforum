package models

import "time"

// 用户类型
type User struct {
	Id                          int
	Uuid, Name, Email, Password string
	CreateAt                    time.Time
}

// Create a new session for an existing u
func (u *User) CreateSession() (s Session, err error) {
	statement := "insert into ss (uuid, email, u_id, created_at) values (?, ?, ?, ?)"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := createUUID()
	stmtin.Exec(uuid, u.Email, u.Id, time.Now())

	stmtout, err := Db.Prepare("select id, uuid, email, u_id, created_at from ss where uuid = ?")
	if err != nil {
		return
	}
	defer stmtout.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmtout.QueryRow(uuid).Scan(&s.Id, &s.Uuid, &s.Email, &s.UserId, &s.CreatedAt)
	return
}

// Get the s for an existing u
func (u *User) Session() (s Session, err error) {
	s = Session{}
	err = Db.QueryRow("SELECT id, uuid, email, u_id, created_at FROM ss WHERE u_id = ?", u.Id).
		Scan(&s.Id, &s.Uuid, &s.Email, &s.UserId, &s.CreatedAt)
	return
}

// Create a new user, save user info into the database
func (u *User) Create() (err error) {
	// Postgres does not automatically return the last insert id, because it would be wrong to assume
	// you're always using a sequence.You need to use the RETURNING keyword in your insert to get this
	// information from postgres.
	statement := "insert into users (uuid, name, email, password, created_at) values (?, ?, ?, ?, ?)"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := createUUID()
	stmtin.Exec(uuid, u.Name, u.Email, Encrypt(u.Password), time.Now())

	stmtout, err := Db.Prepare("select id, uuid, created_at from users where uuid = ?")
	if err != nil {
		return
	}
	defer stmtout.Close()
	// use QueryRow to return a row and scan the returned id into the User struct
	err = stmtout.QueryRow(uuid).Scan(&u.Id, &u.Uuid, &u.CreatedAt)
	return
}

// Delete user from database
func (u *User) Delete() (err error) {
	statement := "delete from users where id = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Id)
	return
}

// Update user information in the database
func (u *User) Update() (err error) {
	statement := "update users set name = ?, email = ? where id = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Name, u.Email, u.Id)
	return
}

// Delete all users from database
func UserDeleteAll() (err error) {
	statement := "delete from users"
	_, err = Db.Exec(statement)
	return
}

// Get all users in the database and returns it
func Users() (us []User, err error) {
	rows, err := Db.Query("SELECT id, uuid, name, email, password, created_at FROM users")
	if err != nil {
		return
	}
	for rows.Next() {
		u := User{}
		if err = rows.Scan(&u.Id, &u.Uuid, &u.Name, &u.Email, &u.Password, &u.CreatedAt); err != nil {
			return
		}
		us = append(us, u)
	}
	rows.Close()
	return
}

// Get a single user given the email
func UserByEmail(email string) (u User, err error) {
	u = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = ?", email).
		Scan(&u.Id, &u.Uuid, &u.Name, &u.Email, &u.Password, &u.CreatedAt)
	return
}

// Get a single user given the UUID
func UserByUUID(uuid string) (u User, err error) {
	u = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE uuid = ?", uuid).
		Scan(&u.Id, &u.Uuid, &u.Name, &u.Email, &u.Password, &u.CreatedAt)
	return
}
