package helper

//variables and functions defined outside any function, can be accessed in all other files within the same package
// biến và hàm nằm bên ngoài function chính có thể được truy cập ở bất kỳ file nào nếu chúng cùng chung package
import "fmt"

//export function bằng cách viết hoa chữ cái đầu, tương tự với variables
func GetUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
