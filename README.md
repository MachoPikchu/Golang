# Golang
Golang assignments

in main.go file connection to db mysql with name order_service and port 3306
Also in main.go file made simple crud operations and structures was made using automigrate and Classes.go file with structures in it 

Migrations made auto 


Queries are written in special file QueryMethods
getStudentsByDepartment: It finds the names of students in a given department from the database.

getCoursesByInstructor: This one finds the names of courses taught by a specific instructor.

getEnrollmentsByStudent: This function gets details of courses a student is enrolled in, like the course name and their grade.


I made some classes referencing each other Like smany-to-many relationship between Students and Courses through the Enrollments table, one-to-many relationship between Instructors and Courses.

Transaction i write to side Golang.txt file 
We start a new transaction, which is like a container for multiple database operations
Deferred Rollback on Panic:
If something goes wrong (like an error or unexpected issue), we want to make sure we don't mess up the database.
Inside the transaction, we do what we need to do: add a new student, create a new course, and enroll the student in the course
Commit Transaction:
If everything is ok we " by commit the transaction. 

Hooks made using Time in classes file BeforeCreate, BeforeUpdate

Soft deleting
Used Gorm's Unscoped method to perform a soft delete operation.
Call the Delete method with the Unscoped option, passing the student's ID.
The Delete method updates the deleted_at column of the student record with the current timestamp.


Function to Get Students Enrolled in a Course:
This function helps us find which students are taking a particular course.
We make a list to hold the names of these students.
We only keep the students who are actually enrolled in that course. We will get the list of students.

Function to Get Total Students in a Department:
This function shows us how many students are in a specific department.
We start by counting how many students belong to that department.
We get the total count of students in the department.
