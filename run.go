package main

import "fmt"

func testTask(args ...interface{}) {
	// fmt.Print(args)
	if len(args) != 0 {
		for i, v := range args {
			fmt.Printf("Func with i: %d", i)
			fmt.Printf(", v: %x\n", v)
		}

		return
	}
	fmt.Print("Func with no args exec...\n")
}

// func testTask2(a string, b string) {
// 	fmt.Printf("Func with a: %s, b: %s", a, b)
// }

func main() {
	// queue := NewThreadSafeQueue(0)

	// go producer(queue)
	// go consumer(queue)

	// time.Sleep(time.Second * 5000)

	task := NewTask()
	task.Exec(testTask, "hello", 666, []int{7, 8, 9})

	// t2Args := map[string]string{
	// 	"a": "Hello",
	// 	"b": "Norman",
	// }
	// task2 := NewTask()
	// task2.Exec(testTask2, t2Args)
}
