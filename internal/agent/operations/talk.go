package operations

import (
	"fmt"

	"github.com/Angel-del-dev/bee/internal/utils/types"
)

func Talk(op types.Operation) error {
	fmt.Println("Talking")
	fmt.Println(op)
	return nil
}
