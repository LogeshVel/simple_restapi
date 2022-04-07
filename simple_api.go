package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Employee struct {
	Name       string `json:"name"`
	EmpID      int    `json:"emp_id"`
	EmpDetails EmployeeDetails
}

type EmployeeDetails struct {
	Email   string `json:"email"`
	Address string `json:"addr"`
	PhoneNo int    `json:"phone_no"`
}

var my_employees []Employee

func findMax(a []int) (int, error) {
	if len(a) < 1 {
		return 0, errors.New("Given the slice with len 0")
	}
	max := a[0]
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max, nil
}

func getEmployee(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if request.Method != "GET" {
		json.NewEncoder(w).Encode(map[string]string{"status": "Failed", "Desc": "Incorrect method"})
		return
	}
	query_params := request.URL.Query()
	query_emp_id := query_params["emp_id"]
	if len(query_emp_id) < 1 {
		json.NewEncoder(w).Encode(map[string]string{"status": "Failed", "Desc": "Expected query parameters for emp id"})
		return
	}
	for _, employee := range my_employees {

		if query_emp_id[0] == strconv.Itoa(employee.EmpID) {
			json.NewEncoder(w).Encode(employee)
			return
		}
	}
	json.NewEncoder(w).Encode(Employee{})
}

func getAllEmployee(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if request.Method != "GET" {
		json.NewEncoder(w).Encode(map[string]string{"status": "Failed", "Desc": "Incorrect method"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(my_employees)
}

func addEmployee(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if request.Method != "POST" {
		json.NewEncoder(w).Encode(map[string]string{"status": "Failed", "Desc": "Incorrect method"})
		return
	}
	var emp Employee
	json.NewDecoder(request.Body).Decode(&emp)
	emp_id_list := []int{}
	for _, emp_value := range my_employees {
		emp_id_list = append(emp_id_list, emp_value.EmpID)
	}
	max_id, err := findMax(emp_id_list)
	if err != nil {
		fmt.Println(err)
		max_id = 0
	}
	emp.EmpID = max_id + 1
	my_employees = append(my_employees, emp)
	json.NewEncoder(w).Encode(map[string]string{"status": "Success", "Desc": "Employee added"})

}

// url : baseurl/delemp/{emp_id}
// Ex: localhost:8080/delemp/2
func deleteEmployee(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if request.Method != "DELETE" {
		json.NewEncoder(w).Encode(map[string]string{"status": "Failed", "Desc": "Incorrect method"})
		return
	}
	id := strings.TrimPrefix(request.URL.Path, "/delemp/")
	for index, emp_items := range my_employees {
		if strconv.Itoa(emp_items.EmpID) == id {
			my_employees = append(my_employees[:index], my_employees[index+1:]...)
			json.NewEncoder(w).Encode(map[string]string{"status": "Success", "Desc": "Employee deleted"})
			return
		}
	}
	json.NewEncoder(w).Encode(map[string]string{"status": "Failed", "Desc": "Given Employee ID not found"})

}

// url : baseurl/update_emp/{emp_id}
// Ex: localhost:8080/update_emp/2
func updateEmployee(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if request.Method != "PUT" {
		json.NewEncoder(w).Encode(map[string]string{"status": "Failed", "Desc": "Incorrect method"})
		return
	}
	id := strings.TrimPrefix(request.URL.Path, "/update_emp/")
	var emp Employee
	json.NewDecoder(request.Body).Decode(&emp)
	for index, emp_items := range my_employees {
		if strconv.Itoa(emp_items.EmpID) == id {
			my_employees = append(my_employees[:index], my_employees[index+1:]...)
			emp.EmpID = emp_items.EmpID
			my_employees = append(my_employees, emp)
			json.NewEncoder(w).Encode(map[string]string{"status": "Success", "Desc": "Employee updated"})
			return
		}
	}
	json.NewEncoder(w).Encode(map[string]string{"status": "Failed", "Desc": "Given Employee ID not found"})

}

func rootfunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Success", "Desc": "RootPage Endpoint Hitted"})

}

func main() {
	// initial mock data
	my_employees = append(my_employees, Employee{Name: "Logesh", EmpID: 1, EmpDetails: EmployeeDetails{Email: "logesh@my.org", Address: "some addr", PhoneNo: 94389489}})
	http.HandleFunc("/", rootfunc)
	http.HandleFunc("/emp", getEmployee)
	http.HandleFunc("/allEmp", getAllEmployee)
	http.HandleFunc("/addemp", addEmployee)
	http.HandleFunc("/delemp/", deleteEmployee)
	http.HandleFunc("/update_emp/", updateEmployee)

	http.ListenAndServe(":8090", nil)

}
