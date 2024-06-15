package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

// interface
type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

func main() {
	printSomething(1)
	printSomething(1.5)

	title, content := getNoteData()
	todoText := getUserInput("Todo text\n")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	err = saveData(todo)
	if err != nil {
		return
	}

	err = saveData(userNote)
	if err != nil {
		return
	}
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

// alternatively, instead of "interface{}" you can also write "any"
// that's an alias for this special "any value is allowed" type
func printSomething(value interface{}) {
	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer:", intVal)
	}

	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("Float:", floatVal)
	}

	stringVal, ok := value.(string)

	if ok {
		fmt.Println("string:", stringVal)
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println(value)
	// }
	// fmt.Println(value)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the data failed.")
		return err
	}
	data.Save()

	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

// generics
func add[T int | float64 | string](a, b T) T {
	return a + b
}
