package main
//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
import (
	"fmt"
	"strings"
)

//3

func main() {
	conferenceName := "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v re still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Введите имя: ")
		fmt.Scan(&firstName)

		fmt.Println("Введите фамилию: ")
		fmt.Scan(&lastName)

		fmt.Println("Введите email: ")
		fmt.Scan(&email)

		fmt.Println("Введите количество билетов: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings[0] = append(bookings, firstName+" "+lastName)

		fmt.Printf("Весь срез: %v\n", bookings)
		fmt.Printf("Первое значение: %v\n", bookings[0])
		fmt.Printf("Тип среза: %T\n", bookings)
		fmt.Printf("Длина среза: %v\n", len(bookings))

		fmt.Printf("Спасибо %v %v за бронирование %v билетов. Подтверждение прийдет на ваш email: %v \n ", lastName, firstName, userTickets, email)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)

			firstName = append(firstNames, names[0])
		}
		fmt.Printf("The first names of bookings are: %v\n ", firstNames)
	}

}
