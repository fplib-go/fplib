package fplib

import (
	"fmt"
)

func InDebug() bool {
	return is_Debug
}

func Debug(msgs ...interface{}) {
	if is_Debug {
		fmt.Print("-----Debug Log-----:")
		fmt.Println(Datetime_ms())

		for _, value := range msgs {
			fmt.Print(value)
			fmt.Print("  ")
		}
		fmt.Println("\n-------------------")
	}

}
func ShowType(msgs ...interface{}) {

	fmt.Print("-----LogType-----:")
	fmt.Println(Datetime_ms())
	fmt.Print(Typeof(msgs...))
	fmt.Println("\n-------------------")
}

func Error(msgs ...interface{}) {
	if len(msgs) > 0 {
		fmt.Print("-----Error Log-----:")
		fmt.Println(Datetime_ms())

		for _, value := range msgs {
			fmt.Print(value)
			fmt.Print("  ")
		}
		fmt.Println("\n-------------------")
		panic(msgs[0])
	} else {
		panic("Error")
	}

}
