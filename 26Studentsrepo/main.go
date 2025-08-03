package main

import "fmt"

var students []Student

type Student struct {
	Rollno int
	Name   string
	Marks  float64
}

func main() {
	var choice int

	for {
		fmt.Println("\n=== Student Marks Manager ===")
		fmt.Println("1. Add Student")
		fmt.Println("2. List Students")
		fmt.Println("3. Update Marks")
		fmt.Println("4. Delete Student")

		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			AddStudent()
		case 2:
			listStudents()
		case 3:
			updateMark()
		case 4:
			DeleteStudent()

		case 5:
			fmt.Println("👋 Exiting...")
			return
		default:
			fmt.Println("❌ Invalid choice")
		}
	}
}

func AddStudent() {
	var rollno int
	var name string
	var marks float64

	fmt.Println("Enter the marks")
	fmt.Scan(&rollno)
	fmt.Println("Enter the name")
	fmt.Scan(&name)
	fmt.Println("Enter the marks")
	fmt.Scan(&marks)

	students = append(students, Student{Rollno: rollno, Name: name, Marks: marks})
	fmt.Println("Student added successfull")

}
func listStudents() {
	if len(students) == 0 {
		fmt.Println("invlaid")
		return
	} else {
		fmt.Println("student list")
		for _, i := range students {
			fmt.Println(i.Rollno, i.Marks, i.Name)
		}

	}
}
func updateMark() {
	var rollno int
	var name string
	var marks float64

	fmt.Println("enter the roolno")
	fmt.Scan(&rollno)
	fmt.Println("enter the name")
	fmt.Scan(&name)
	fmt.Println("Enter the marks to be added")
	fmt.Scan(&marks)
	for i := range students {
		if students[i].Rollno == rollno && students[i].Name == name {
			students[i].Marks += marks
			fmt.Println("updated marks for the student is ", students[i].Marks)

		} else {
			fmt.Println("invalide students")
		}
	}
}
func DeleteStudent() {
	var rollno int
	fmt.Println("enter the roolno")
	fmt.Scan(&rollno)
	index := -1
	for i, s := range students {
		if s.Rollno == rollno {
			index = i
			break
		}

	}
	if index == -1 {
		fmt.Println("student not found")
		return
	} else {
		fmt.Println(append(students[:index], students[index+1:]...))
		fmt.Println("Deleted student ", students[index].Name)

	}

}
