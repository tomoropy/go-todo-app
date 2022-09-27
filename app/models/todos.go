package models

import (
	"log"
)

type Todo struct {
	ID        int
	Content   string
	UserID    int
}

func (u *User) CreateTodo(content string) (err error) {

	cmd := `insert into todos (
		content,
		user_id) values (?,?)`

	_, err = Db.Exec(cmd, content, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err

}

func GetTodo(id int) (todo Todo, ferr error) {
	cmd := `select id, content, user_id from todos
	where id = ?`
	todo = Todo{}

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
	)
	return todo, err
}

func GetTodos() (todos []Todo, err error) {
	cmd := `select id, content, user_id from todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
		)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `select id, content, user_id from 
	todos where user_id = ?`

	rows, _ := Db.Query(cmd, u.ID)
	if err != err {
		log.Fatalln(err)
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
		)

		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

func (t *Todo) UpdateTodo() error {
	cmd := `update todos set content = ?, user_id =? 
	where id = ?`

	_, err = Db.Exec(cmd, t.Content, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (t *Todo) DeleteTodo() error {
	cmd := `delete from todos where id = ?`
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
