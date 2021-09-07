package school

type Address struct {
	Street string
	City string
	State string
	zipCode int
}

type Student struct {
	Studentid int 
	Name string
	Address	
}

type Teacher struct {
	EmployeeID int
	Name string
	Salary float64
	Addresses
}