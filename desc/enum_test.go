package desc

import (
	"fmt"
)

func ExampleEnum() {
	NormalId := Enum("NormalId",
		Item{"Top", 1},
		Item{"Bottom", 4},
		Item{"Back", 2},
		Item{"Front", 5},
		Item{"Right", 0},
		Item{"Left", 3},
	)
	fmt.Println(NormalId)
	fmt.Println(NormalId.Item("Front"))
	fmt.Println(NormalId.Item("Front").Value())
	// Output:
	// Enum.NormalId
	// Enum.NormalId.Front
	// 5
}
