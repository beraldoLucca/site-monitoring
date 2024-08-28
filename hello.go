package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const timeSleep = 5

func main() {

	//playingWithSlices()
	sites := readFile()
	showIntro()

	for {
		showMenu()

		comando := readCommand()

		switch comando {
		case 1:
			timeRepeat := timeRepeated()
			startMonitoring(timeRepeat, sites)
		case 2:
			showLogs()
		case 0:
			fmt.Println("Exiting the program.")
			os.Exit(0)
		default:
			fmt.Println("I dont know this command")
			os.Exit(-1)
		}
	}

}

func showIntro() {
	name := "Beraldini"
	var version float32 = 1.3

	fmt.Println("Your first name is", name)
	fmt.Println("a versão que está usando é:", version)
	fmt.Println("")
}

func showMenu() {
	fmt.Println("1 - Start of monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")
	fmt.Println("")
}

func readCommand() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)
	return comando
}

func startMonitoring(timeRepeat int, sites []string) {
	fmt.Println("Monitoring...")
	//site := "https://httpbin.org/status/200"

	for i := 0; i < timeRepeat; i++ {
		for _, site := range sites {
			resp, err := http.Get(site)

			if err != nil {
				fmt.Println("Ocorreu um erro:", err)
			}

			if resp.StatusCode == 200 {
				registerLog(site, true)
				fmt.Println("Site", site, "foi carregado com sucesso")
			} else {
				registerLog(site, false)
				fmt.Println("Site", site, "com problemas")
			}
		}
		fmt.Println("")
		time.Sleep(timeSleep * time.Second)
	}

}

func playingWithSlices() {
	nomes := []string{"Lucca", "Camila", "Mãe", "Pai"}
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println(len(nomes))
	fmt.Println(cap(nomes))
	nomes = append(nomes, "Little legs")
	fmt.Println(cap(nomes))
}

func timeRepeated() int {
	var time int
	fmt.Println("")
	fmt.Println("Type how many times you want to repeat the monitoring...")
	fmt.Println("")
	fmt.Scan(&time)
	return time
}

func readFile() []string {

	sites := []string{}

	file, err := os.Open("sites.txt")
	//file, err := os.ReadFile("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}
	file.Close()
	return sites
}

func registerLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func showLogs() {

	// file, err := os.Open("log.txt")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fileLogs := bufio.NewReader(file)

	// for {
	// 	line, err := fileLogs.ReadString('\n')
	// 	line = strings.TrimSpace(line)

	// 	fmt.Println(line)

	// 	if err == io.EOF {
	// 		break
	// 	}

	// }

	file, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}
