package hello

import (
	"fmt"

	"go.uber.org/zap"
)

func Hello() {
	logger := zap.NewExample() // or NewProduction, or NewDevelopment
	defer logger.Sync()

	fmt.Println("Hello SRE!")
	logger.Info("printed fancy greeting")
}
