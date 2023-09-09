package internal

import (
	"fmt"
	"time"
)

type TaskStatus int

const (
	StatusPending TaskStatus = iota
	StatusInProgress
	StatusCompleted
	StatusCanceled
)

var TaskStatusStrings = [...]string{
	StatusPending:    "Ожидает",
	StatusInProgress: "В процессе",
	StatusCompleted:  "Завершена",
	StatusCanceled:   "Отменена",
}

type Task struct {
	ID          int
	Title       string
	Description string
	Status      TaskStatus
	DueDate     time.Time
}

// Project представляет проект с списком задач.
// Структура содержит поле Tasks, которое является срезом задач,
// где каждая задача представлена структурой Task.
type Project struct {
	Tasks []Task
}

// NewProject создает и возвращает новый экземпляр проекта.
// Этот проект не содержит задач и может быть использован
// как отправная точка для добавления задач.
func NewProject() *Project {
	return &Project{}
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
func (p *Project) UpdateTaskStatus(taskID int, status TaskStatus) error {
	for i, task := range p.Tasks {
		if task.ID == taskID {
			p.Tasks[i].Status = status
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", taskID)
}

// ListTasks TODO: выводит список всех задач в проекте
func (p *Project) ListTasks() []Task {
	return p.Tasks
}
