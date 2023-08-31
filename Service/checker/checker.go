package checker

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


//Интерфейсы работают как промежуточный инструмент для разных структур с одними методами.
//Интерфейсы обладают полями которые нужно объявлять для доступа.
//В ином случае это воспринимается как метод интерфейса и выдаёт ошибку.
//В интерфейсы можно вложить значения.
