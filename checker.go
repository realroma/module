package main

import (
	"context"
	"fmt"
	"time"
)

type Checker struct {
	items []interface{}
	//Поле items(*.items) состоит из слайса интерфейсов.
}

func (Checker *Checker) Run(ctx context.Context, change <-chan int, tm time.Duration) {
	//Создаём таймер на пять секунд
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	ticker := time.NewTicker(tm * time.Second)
	for _ = range ticker.C {
		select {
		case <-ctx.Done():
			fmt.Println("Process is stoped.")
			break
		case <-change:
			//Если произошли изменения будет ждать следующей проверки.
			fmt.Println("Wait.")
		default:
			Check(Checker)
		}
	}
}

func (Checker *Checker) Stop(cancel context.CancelFunc) {
	cancel()
}

func (Checker *Checker) Add(items *Checkable, change chan<- int) {
	//Отправляем в канал значение и ожидаем считывания.
	change <- 1
	//Метод для добавления в слайс интерфейсов с типом Checkable.
	Checker.items = append(Checker.items, *items)
}

func (Checker *Checker) String() string {
	//Метод вызываемый fmt.Print(*).
	var str string
	var heal string
	for _, value := range Checker.items {
		a, _ := value.(*HealthCheck)
		//Объявление переменной как структуры под интерфейсом.(Примерное описание интерфейса в конце).
		if a.health == true {
			heal = "true"
		} else {
			heal = "false"
		}
		str = str + "GetID: " + a.id + " Health: " + heal + ".\n"
		//Присвоение поля структуры под интерфейсом в переменную.
	}
	return fmt.Sprintf(str)
}

func Check(c *Checker) {
	for _, value := range c.items {
		a, _ := value.(*HealthCheck)
		if a.health == false {
			fmt.Println(a.id + " GetID:" + "false.")
		}
	}
}

type Measurable interface {
	GetMetrics() string
}

func GetMertics() string {
	return "nil"
}

type Checkable interface {
	Measurable
	//Встравание интерфейса и его методов.
	Ping() error
	//Метод для интерфейса. Все методы интерфейса джны быть и у объекта, для корректной работы интерфейса с объектом.
	GetID(bool) string
	Health(bool) bool
}

type HealthCheck struct {
	err    error
	id     string
	health bool
}

func (s *HealthCheck) Ping() error {
	//Метод для структуры.
	return nil
}

func (s *HealthCheck) GetID(in bool) string {
	s.id = "link"
	s.Health(in)
	return s.id
	//Изменяет поле структуры как возвращаемое значение.
}

func (s *HealthCheck) Health(in bool) bool {
	switch in {
	case true:
		s.health = true
	case false:
		s.health = false
	}
	return s.health
}

func (s *HealthCheck) GetMetrics() string {
	return "nil"
}

func main() {
	items := &Checker{}

	ctx, cancel := context.WithCancel(context.Background())
	change := make(chan int)
	//Присваивание ссылки(указателя) на структуру в переменную.
	var Checkable1 Checkable
	//Объявление переменной как тип интерфейс.
	var Checkable2 Checkable
	var Checkable3 Checkable
	var Checkable4 Checkable

	Checkable1 = &HealthCheck{}
	//Добавление структуры в переменную-интерфейс.
	Checkable2 = &HealthCheck{}
	Checkable3 = &HealthCheck{}
	Checkable1.GetID(false)
	//Метод изменяет поле структуры у переменной типа интерфейс.
	Checkable2.GetID(true)
	Checkable3.GetID(false)

	go items.Run(ctx, change, 5)

	items.Add(&Checkable1, change)
	items.Add(&Checkable2, change)
	items.Add(&Checkable3, change)

	time.Sleep(10 * time.Second)
	fmt.Println("Ch 2 False. Add Ch 4 false.")
	Checkable2.GetID(false)
	Checkable4 = &HealthCheck{}
	Checkable4.GetID(false)
	items.Add(&Checkable4, change)
	Checkable4.GetID(false)
	fmt.Println("Finish.")
	fmt.Print(items)
	items.Stop(cancel)
}

//Интерфейсы работают как промежуточный инструмент для разных структур с одними методами.
//Интерфейсы обладают полями которые нужно объявлять для доступа.
//В ином случае это воспринимается как метод интерфейса и выдаёт ошибку.
//В интерфейсы можно вложить значения.
