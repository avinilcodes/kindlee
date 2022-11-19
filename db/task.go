package db

import (
	"context"
	"strings"
	"time"
)

const (
	findTaskByDescription   = "SELECT id,description,task_status_code,started_at,ended_at FROM TASKS WHERE description=$1"
	findTaskByTaskId        = "SELECT id,description,task_status_code,started_at,ended_at FROM TASKS WHERE id=$1"
	insertTask              = `INSERT INTO tasks (id,description,task_status_code,started_at,ended_at) VALUES ($1,$2,$3,$4,$5)`
	findAllTasks            = "select * from tasks"
	findAllUserNameTaskName = "select u.email,t.description from users_tasks ut Inner join users u on u.id = ut.user_id Inner join tasks t on t.id = ut.task_id"
	updateTaskStatus        = `update tasks set task_status_code=$1,ended_at=$2 where id =$3`
	findUserFromTaskId      = `select * from users where id = (select user_id from users_tasks where task_id = $1)`
	findTasksByUserId       = `select * from tasks where id in (select task_id from users_tasks where user_id= $1)`
)

type Task struct {
	ID             string    `json:"id" db:"id"`
	Description    string    `json:"description" db:"description"`
	TaskStatusCode string    `json:"task_status_code" db:"task_status_code"`
	StartedAt      time.Time `json:"started_at" db:"started_at"`
	EndedAt        time.Time `json:"ended_at" db:"ended_at"`
}

func (s *store) CreateTask(ctx context.Context, task Task) (err error) {
	if task.Description == "" {
		return ErrCannotAddEmptyTask
	}
	task.Description = strings.ToLower(task.Description)
	res, err := s.db.Exec(findTaskByDescription, task.Description)
	if err != nil {
		return
	}
	cnt, _ := res.RowsAffected()
	if cnt == 0 {
		_, err = s.db.Query(insertTask, task.ID, task.Description, task.TaskStatusCode, task.StartedAt, task.EndedAt)
		if err != nil {
			return err
		}
		return
	}
	return ErrTaskAlreadyExist
}

func (s *store) UpdateTaskStatus(ctx context.Context, id string, status string, userEmail string) (err error) {
	var task Task
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &task, findTaskByTaskId, id)
	})
	if err != nil {
		return err
	}
	var user User
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &user, findUserFromTaskId, task.ID)
	})
	if err != nil {
		return
	}
	var currentUser User
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &currentUser, findUserByEmailQuery, userEmail)
	})
	if err != nil {
		return
	}
	if status != "in_progress" && status != "mr_approved" && status != "code_review" {
		return ErrTaskCannotBeUpdated
	}
	if userEmail != user.Email && currentUser.RoleType != "admin" && currentUser.RoleType != "super_admin" {
		return ErrTaskAssignedToAnotherUser
	}
	if status == "mr_approved" && task.TaskStatusCode == "in_progress" {
		return ErrTaskStatusError
	}
	if status == "mr_approved" && task.TaskStatusCode == "mr_approved" {
		return ErrTaskIsCompleted
	}
	if status == "mr_approved" && task.TaskStatusCode != "code_review" {
		return ErrTaskStatusError
	}
	if status == "mr_approved" && currentUser.RoleType != "admin" {
		return ErrOnlyAdminAccess
	}
	if status == "code_review" && task.TaskStatusCode != "in_progress" {
		return ErrTaskStatusError
	}
	flag := status != "in_progress" && status != "mr_approved" && status != "code_review"
	if !flag {
		if status == "mr_approved" {
			task.EndedAt = time.Now()
		}
		_, err = s.db.Query(updateTaskStatus, status, task.EndedAt, task.ID)
		if err != nil {
			return err
		}
		return
	}
	return ErrTaskCannotBeUpdated
}

func (s *store) ListTasks(ctx context.Context, email string) (tasks []Task, err error) {
	var currentUser User
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &currentUser, findUserByEmailQuery, email)
	})
	if err != nil {
		return
	}
	if currentUser.RoleType == "admin" || currentUser.RoleType == "super_admin" {
		err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
			return s.db.SelectContext(ctx, &tasks, findAllTasks)
		})
	} else {
		err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
			return s.db.SelectContext(ctx, &tasks, findTasksByUserId, currentUser.ID)
		})
	}
	return
}

func (s *store) ListUserTask(ctx context.Context) (usertask []NameUserTask, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.SelectContext(ctx, &usertask, findAllUserNameTaskName)
	})
	return
}
