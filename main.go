package main

import (
	"errors"
	"fmt"
)

//ТЫ УЕБИЩЕ ЕСЛИ КОД ЧЕКАЕШЬ ТУТ НЕТ РАТКИ СМОТРИ ИМПОРТЫ))))

func main() {
	s, d, c := cartholder()
	da := hot()
	if da == true {
		balik(s, d, c)
	} else if da == false {
		fmt.Println("Иди нахуй скатина неблагодарная сам пиши рас такой умный")
	}
}
func cartholder() (string, string, string) {
	asdasd := [3]string{}
	fmt.Println("Введите ваши банковские карты 💳")
	fmt.Println("Максимальное количество 3 шт")
	fmt.Scan(&asdasd[0])
	fmt.Scan(&asdasd[1])
	fmt.Scan(&asdasd[2])
	fmt.Println("Ваш Картхолдер 💳")
	s := asdasd[0]
	d := asdasd[1]
	c := asdasd[2]
	return s, d, c
}
func hot() bool {
	for {
		fmt.Println("Хотите Заполнить Баланс карт? y/n")
		var od string
		fmt.Scan(&od)
		if od == "y" {
			return true
		} else if od == "n" {
			return false
		} else {
		}
	}
}
func balik(s, d, c string) {
	balance := make([]float64, 3)
	fmt.Print("Сколько денег у вас на  1-й карте ", s, ": ")
	fmt.Scan(&balance[0])
	fmt.Print("Сколько денег у вас на  2-й карте ", d, ": ")
	fmt.Scan(&balance[1])
	fmt.Print("Сколько денег у вас на  3-й карте ", c, ": ")
	fmt.Scan(&balance[2])
	if balance[0] <= 0 {
		err := errors.New("НОРМ пиши или у тебя внатуре смерть в нищите на первой карте")
		fmt.Println(err)
	} else if balance[1] <= 0 {
		err := errors.New("НОРМ пиши или у тебя внатуре смерть в нищите на второй карте")
		fmt.Println(err)
	} else if balance[2] <= 0 {
		err := errors.New("НОРМ пиши или у тебя внатуре смерть в нищите на третьей карте")
		fmt.Println(err)
	}
	fmt.Println("Ваш Баланс:")
	fmt.Println("Ваш Баланс 1-й 💳", s, ":", balance[0])
	fmt.Println("Ваш Баланс 2-й 💳", d, ":", balance[1])
	fmt.Println("Ваш Баланс 3-й 💳", c, ":", balance[2])
	for {
		fmt.Println("Хотите поддержать автора? ДА/НЕТ")
		var dant string
		fmt.Scan(&dant)
		switch dant {
		case "ДА":
			fmt.Println("Пишите Сюда Отправлю Реквезиты")
			fmt.Println("https://t.me/meri_kynemlavat")
			return
		case "НЕТ":
			fmt.Println("Ты че сыр собаки ахуел совсем???")
		default:
			fmt.Println("Сука необразхованая пиши нормально")
		}

	}

}
