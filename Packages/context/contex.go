package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// Контекст работает в горутинах, и встречена проблема с select, когда я не понимаю почему он структуру и канал пытается проверить на true.
// Если в context.WithTimeout(parent, timer) timer > run.main, то контекст не выполнится, так как основная ветка закончится раньше. Окончание горутин не отслеживается.
// Если нужно чтобы контекст выполнился, необходимо добавить ожидание в main или сделать WaitGroap из синхронизации.
// При использовании разных типов в switch case может не давать выполнить ctx.Done().
func ctxTimer(parent context.Context, i int) {
	ctx, err := context.WithTimeout(parent, 1500*time.Millisecond)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Duration(i) * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("Time is out.")
	default:
		fmt.Println("Ready.")
	}
	fmt.Println("Wait")
}

func main() {
	parent := context.Background()
	ctxTimer(parent, 1)
	ctxTimer(parent, 2)
	time.Sleep(3 * time.Second)
	fmt.Println("Main is ends.")
}
