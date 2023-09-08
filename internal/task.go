package internal

import "time"

type TaskStatus int

const (
	StatusPending TaskStatus = iota
	StatusInProgress
	StatusCompleted
	StatusCanceled
)

type Task struct {
	ID          int
	Title       string
	Description string
	Status      TaskStatus
	DueDate     time.Time
}

type Project struct {
	Tasks []Task
}

// AddTask TODO: добавляет новую задачу в проект
func (p *Project) AddTask(task Task) {
	p.Tasks = append(p.Tasks, task)

}

// RemoveTask TODO: удаляет задачу из проекта по заданному идентификатору
func (p *Project) RemoveTask(taskID int) {
	for i, task := range p.Tasks {
		if task.ID == taskID {
			p.Tasks = append(p.Tasks[:i], p.Tasks[i+1:]...)
			return
		}
	}

}

// UpdateTaskStatus TODO: изменяет статус задачи по заданному идентификатору
func (p *Project) UpdateTaskStatus(taskID int, status string) {

}

// ListTasks TODO: выводит список всех задач в проекте
func (p *Project) ListTasks() {

}
