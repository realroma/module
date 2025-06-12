package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// Если в context.WithTimeout(parent, timer) timer > run.main, то контекст не выполнится, так как основная ветка закончится раньше. Окончание горутин не отслеживается.
// Если нужно чтобы контекст выполнился, необходимо добавить ожидание в main или сделать WaitGroap из синхронизации.
// При использовании разных типов в switch case может не давать выполнить ctx.Done().
func ctxTimer() {
	parent := context.Background()
	ctx, err := context.WithTimeout(parent, 1500*time.Millisecond)
	if err != nil {
		log.Fatal(err)
	}
	go went(ctx, 1)
	go went(ctx, 2)
	fmt.Println("Wait")
}
func went(ctx context.Context, i int) {
	time.Sleep(time.Duration(i) * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("Time is out.")
	default:
		fmt.Println("Ready.")
	}
}

func main() {
	ctxTimer()
	time.Sleep(3 * time.Second)
	fmt.Println("Main is ends.")
}
