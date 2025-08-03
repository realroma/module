package main

import (
	"fmt"

	bt "github.com/joeycumines/go-behaviortree"
)

func main() {
	fmt.Println(bt.New(
		bt.Any(bt.All),
		bt.New(func(children []bt.Node) (bt.Status, error) {
			fmt.Println(1)
			return bt.Success, nil
		}),
		bt.New(func(children []bt.Node) (bt.Status, error) {
			fmt.Println(2)
			return bt.Success, nil
		}),
		bt.New(func(children []bt.Node) (bt.Status, error) {
			fmt.Println(3)
			return bt.Success, nil
		}),
		bt.New(func(children []bt.Node) (bt.Status, error) {
			fmt.Println(4)
			return bt.Success, nil
		}),
		bt.New(func(children []bt.Node) (bt.Status, error) {
			fmt.Println(5)
			return bt.Failure, nil
		}),
		bt.New(func(children []bt.Node) (bt.Status, error) {
			fmt.Println(6)
			return bt.Success, nil
		}),
	).Tick())
}
