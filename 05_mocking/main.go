package main

import "log"

type student struct {
	id   string
	name string
	age  int
}

type mySQLDriver struct{}

func (m mySQLDriver) ExecuteSQL(s string, vals ...interface{}) error {
	return nil
}

// StudentStorage defines generic persistance layer for student objects
type StudentStorage interface {
	InsertStudent(s student) error
}

// MockStudentStorage is a fake backend for storing student object
type MockStudentStorage struct{}

func NewMockStudentStorage() MockStudentStorage {
	return MockStudentStorage{}
}

func (ms MockStudentStorage) InsertStudent(s student) error {
	return nil
}

// MySQLStudentStorage is a MySQl backend for storing student object
type MySQLStudentStorage struct {
	driver mySQLDriver
}

func NewMySQLStudentStorage(driver mySQLDriver) MySQLStudentStorage {
	return MySQLStudentStorage{driver: driver}
}

func (ms MySQLStudentStorage) InsertStudent(s student) error {
	sql := "INSERT INTO students (id, name age) VALUES ($1, $2, $3)"
	return ms.driver.ExecuteSQL(sql, s.id, s.name, s.age)
}

// StudentService handles business logic for students
type StudentService struct {
	db StudentStorage
}

func NewStudentService(db StudentStorage) StudentService {
	return StudentService{db: db}
}

func (s *StudentService) Create() (id string, err error) {
	newStudent := student{id: "1002", name: "bob", age: 34}
	if err := s.db.InsertStudent(newStudent); err != nil {
		return "", err
	}
	return newStudent.id, nil
}

func main() {
	realDB := NewMySQLStudentStorage()
	service := NewStudentService(realDB)
	_, err := service.Create()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Success")
}
