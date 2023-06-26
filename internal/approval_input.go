package internal

import (
	"bufio"
	"fmt"
	"gate/handlers"
	"os"
	"strings"
)

func ReadApprovalInput() {
	reader := bufio.NewReader(os.Stdin)

	for req := range handlers.Requests {
		fmt.Printf("Received request: %s %s\n", req.Method, req.Path)
		fmt.Print("Do you want to grant permission for this request? (Y/N): ")

		decision, _ := reader.ReadString('\n')
		decision = strings.TrimSpace(strings.ToLower(decision))

		permission := decision == "y" || decision == "yes"

		handlers.PermissionLock.Lock()
		handlers.PermissionMap[req] = permission
		handlers.PermissionLock.Unlock()

		fmt.Printf("Permission for request %s %s set to %v\n", req.Method, req.Path, permission)
	}
}
