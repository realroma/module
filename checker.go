package main

import ("fmt" )

type Checker struct {
        items []interface{}
	//Поле items(*.items) состоит из слайса интерфейсов.
}

func (Checker *Checker) Add(items Checkable) {
	//Метод для добавления в слайс интерфейсов с типом Checkable.
        Checker.items = append(Checker.items, items)
}

func (Checker *Checker) String() string {
	//Метод вызываемый fmt.Print(*).
        var str string
	var heal string
	for _, value  := range Checker.items {
		a, _:= value.(*HealthCheck)
		//Объявление переменной как структуры под интерфейсом.(Примерное описание интерфейса в конце).
		if a.health == true {
			heal = "true"
		} else {
			heal = "false"
		}
		str = str + "GetID: " + a.id + " Health: " + heal + "\n"
		//Присвоение поля структуры под интерфейсом в переменную.
	}
        return fmt.Sprintf(str)
}

func Check(c *Checker){
	for _, value := range c.items {
		a, _ := value.(*HealthCheck)
		if a.health == false {
			fmt.Println(a.id + " GetID:" + "false")
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
        Health() bool
}

type HealthCheck struct {
	err error
	id string
	health bool
}

func (s *HealthCheck) Ping() error {
	//Метод для структуры.
        return nil
}

func (s *HealthCheck) GetID(a bool) string {
	s.id = "link"
	if a == true {
		s.id = "url"
	}
	s.Health()
        return s.id
	//Изменяет поле структуры как возвращаемое значение.
}

func (s *HealthCheck) Health() bool {
	s.health = true
        return s.health
}

func (s *HealthCheck) GetMetrics() string {
	return "nil"
}

func main() {
        items := &Checker{}
	//Присваивание ссылки(указателя) на структуру в переменную.
        var Checkable1 Checkable
	//Объявление переменной как тип интерфейс.
        var Checkable2 Checkable
        var Checkable3 Checkable
	Checkable1 = &HealthCheck{}
	//Добавление структуры в переменную-интерфейс.
	Checkable2 = &HealthCheck{}
	Checkable3 = &HealthCheck{}
	Checkable1.GetID(false)
	//Метод изменяет поле структуры у переменной типа интерфейс.
	Checkable2.GetID(true)
	Checkable3.GetID(false)
        items.Add(Checkable1)
        items.Add(Checkable2)
        items.Add(Checkable3)
	fmt.Println(items)
	Check(items)
}
//Интерфейсы работают как промежуточный инструмент для разных структур с одними методами.
//Интерфейсы обладают полями которые нужно объявлять для доступа.
//В ином случае это воспринимается как метод интерфейса и выдаёт ошибку.
//В интерфейсы можно вложить значения.
