package main

import (
	"fmt"
	"myproject/Cgpt/internal"
	"time"
)

func main() {

	// экземпляр проекта
	project := internal.NewProject()

	// задачи, которые добавляем в проект
	task1 := internal.Task{
		ID:          1,
		Title:       "Задача 1",
		Description: "Описание задачи 1",
		Status:      internal.StatusPending,
		DueDate:     time.Now().AddDate(0, 0, 7), // Дедлайн через 7 дней
	}

	task2 := internal.Task{
		ID:          2,
		Title:       "Задача 2",
		Description: "Описание задачи 2 ",
		Status:      internal.StatusInProgress,
		DueDate:     time.Now().AddDate(0, 0, 14),
	}

	project.AddTask(task1)
	project.AddTask(task2)

	// Выводим список задач
	tasks := project.ListTasks()
	fmt.Println("Список задач:")
	for _, task := range tasks {
		fmt.Printf("ID: %d, Заголовок: %s, Статус: %v\n", task.ID, task.Title, internal.TaskStatusStrings[task.Status])
	}

	// Изменяем статус задачи
	err := project.UpdateTaskStatus(1, internal.StatusCompleted)
	if err != nil {
		fmt.Println("Ошибка при изменении статуса задачи:", err)
	} else {
		fmt.Println("Статус задачи изменен.")
	}

	// Выводим обновленный список задач
	tasks = project.ListTasks()
	fmt.Println("Обновленный список задач:")
	for _, task := range tasks {
		fmt.Printf("ID: %d, Заголовок: %s, Статус %v\n", task.ID, task.Title, internal.TaskStatusStrings[task.Status])
	}
}
