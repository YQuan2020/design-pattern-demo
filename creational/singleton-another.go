package creational

import (
	"fmt"
	"sync"
)

var once sync.Once

type anotherSingle struct{}

var anotherSingleInstance *anotherSingle

func getInstanceAnother() *anotherSingle {
	if anotherSingleInstance == nil {
		once.Do(func() {
			fmt.Println("creating single instance now")
			anotherSingleInstance = &anotherSingle{}
		})
	} else {
		fmt.Println("single instance already created")
	}
	return anotherSingleInstance
}

func DoAnotherSingle() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}
	fmt.Scanln()
}
