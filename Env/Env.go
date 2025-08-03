package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
)

//Переписать всё на cleanenv
//Очистить глобальный env
//Продублировать и выложить.

//Все переменные от os.Setenv сохраняются в рамках текущего процесса и его дочерних ветвей. Это значит, что при завершении программы они стираются.
//В таком случае, когда мы запускаем программу нам нужно считывать данные из .env потом записывать их в виде os.Setenv(value from file)
//Так же надо записывать при необходимости вроде Write("NameEnv=", os.Getenv(NameEnv))

func ConfRead() {
	path, _ := os.Getwd()
	os.Chdir(path + "/module/Env")
	fmt.Println(path)
	f, err := os.Open(".env")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		str := scanner.Text()
		var arr []string
		arr = strings.Split(str, "=")
		os.Setenv(arr[0], arr[1])
	}
}

func main() {
	type config struct {
		//В кавычках определяется название переменной, стандартное знаечение. Без неё, он не будет видеть его как необходимое ему значение
		Key  string `env:"KEY"`
		KEY2 string `env:"KEY2" env-default:"None"`
		//Пример ошибки неправильного названия
		KEY3 string `env:"Key3"`
	}

	var cfg config

	os.Setenv("KEY", "Value")
	//Неправильное название, которое не будет определено, но всё равно запишется.
	os.Setenv("KEY3", "NotNone")

	//Надо было правильно читать. Была функция ReadConfig. Если заглянуть в библиотеку cleanenv и посмотреть что она делает, то можно заметить что она принимает path и cfg
	//ReadEnv принимает только cfg. А почему? Потому что она считывает из env и скорее всего  помещает данные в структуру.
	ConfRead()
	cleanenv.ReadEnv(&cfg)
	fmt.Println(&cfg)

	ConfRead()
	cleanenv.ReadEnv(&cfg)
	os.Setenv("Key3", "OtherValue")
	cleanenv.ReadEnv(&cfg)

	fmt.Println(&cfg)

	text, err := os.ReadFile("/home/cpp/project/module/Env/env/.env")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(text))

	//Верный вариант работы с конфигами.
	//Из-за махинаций выше, у меня выводится 3 значения вместо 2-х.
	var cfg2 config
	cleanenv.ReadConfig("/home/cpp/project/module/Env/.env", &cfg2)
	fmt.Println(cfg2)
}
