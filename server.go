package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// var todoList string

// func getCmd(input string) string {
// 	inputArr := strings.Split(input, " ")
// 	return inputArr[0]
// }

// func getMessage(input string) string {
// 	inputArr := strings.Split(input, " ")
// 	var result string
// 	for i := 1; i < len(inputArr); i++ {
// 		result += inputArr[i]
// 	}
// 	return result
// }

// func updateTodoList(input string) {
// 	tmpList := todoList
// 	todoList = []string{}
// 	for _, val := range tmpList {
// 		if val == input {
// 			continue
// 		}
// 		todoList = append(todoList, val)
// 	}
// }

func doFirstTask(input string) string {
	//inputArr := string.Split(input, "1")
	var result string
	// for i := 1; i < len(input); i++ {
	// 	result += inputArr[i]
	// }
	var answer = []byte(input)
	answer[0] = ' '
	var b string
	for i := 1; i < len(answer); i++ {
		if answer[i] != ' ' {
			b = b + string(answer[i])
		} else {
			if n, err := strconv.Atoi(b); err == nil {
				s := strconv.Itoa(n * n)
				result += string(s) + " "
			} else {
				log.Println(b, " не является целым числом.")
				result += "Не является целым числом. "
			}
			b = b[:0]
		}
	}
	return result
}

func doSecondTask(input string) string {
	var answer = []byte(input)
	answer[0] = ' '
	result := string(answer)
	return result
}

func doTasks(input string) string {
	var result string
	if input[0] == '1' {
		result = doFirstTask(input)
	}
	if input[0] == '2' {
		result = doSecondTask(input)
	}

	return result
}

func doTheerdTaskF(input string) string {
	var result string
	// for i := 1; i < len(input); i++ {
	// 	result += inputArr[i]
	// }
	var answer = []byte(input)
	answer[0] = ' '
	var b string
	for i := 1; i < len(answer); i++ {
		b = b + string(answer[i])
	}
	if n, err := strconv.Atoi(b); err == nil {
		s := strconv.Itoa(n * n)
		result += string(s)
	} else {
		log.Println(b, " не является целым числом.")
		result += "Не является целым числом. "
	}
	return result
}

func doTheerdTaskS(input string) string {
	var result string = "Количество символов в строке равно: "
	// for i := 1; i < len(input); i++ {
	// 	result += inputArr[i]
	// }
	var answer = []byte(input)
	answer[0] = ' '
	var b int
	for i := 1; i < len(answer); i++ {
		b++
	}
	result += strconv.Itoa(b)
	return result
}

func main() {

	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		// Upgrade upgrades the HTTP server connection to the WebSocket protocol.
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade failed: ", err)
			return
		}
		defer conn.Close()

		var checkingValue string
		var answer string
		// Continuosly read and write message
		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read failed:", err)
				break
			}
			input := string(message)
			// cmd := getCmd(input)
			// msg := getMessage(input)
			checkInput := []byte(input)

			if checkInput[0] == '3' && checkingValue == "" {
				switch string(checkInput) {
				case "30":
					checkingValue = "0"
					answer = "Сервер ожидает ввода цифры"
				case "31":
					checkingValue = "1"
					answer = "Сервер ожидает ввода строки"
				case "33":
					checkingValue = "3"
					conn.Close()
				default:
					answer = "Неопознанная команда"
				}
			} else if checkInput[0] != '3' {
				answer = doTasks(input)
			} else if checkingValue != "" {
				switch checkingValue {
				case "0":
					answer = doTheerdTaskF(input)
					checkingValue = checkingValue[:0]

				case "1":
					answer = doTheerdTaskS(input)
					checkingValue = checkingValue[:0]
				default:
				}
			}

			// if cmd == "add" {
			// 	todoList = append(todoList, msg)
			// } else if cmd == "done" {
			// 	updateTodoList(msg)
			// }
			output := "Current Answers: \n"

			output += "\n - " + answer + "\n"

			output += "\n----------------------------------------"
			message = []byte(output)
			err = conn.WriteMessage(mt, message)
			if err != nil {
				log.Println("write failed:", err)
				break
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "src/websockets.html")
	})

	http.HandleFunc("/index.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "src/index.js")
	})

	http.ListenAndServe(":8080", nil)
}
