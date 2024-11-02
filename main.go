package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "projectgen",
		Short: "ProjectGen is a CLI tool to create project structures",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Please provide the project name")
				os.Exit(1)
			}
			projectName := args[0]
			createProjectStructure(projectName)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createProjectStructure(baseDir string) {
	folders := []string{
		"cmd/your-app",
		"internal/domain/payment",
		"internal/domain/customer",
		"internal/domain/order",
		"internal/application/checkout",
		"internal/application/payments",
		"internal/application/customers",
		"internal/application/uow",
		"pkg/shared",
		"pkg/middleware",
		"configs",
		"scripts",
		"test/integration",
		"test/unit",
	}

	files := map[string]string{
		"cmd/your-app/main.go":                               "package main\n\n// Main entry point",
		"internal/domain/payment/payment.go":                 "package payment\n\n// Payment entity definition",
		"internal/domain/payment/state.go":                   "package payment\n\n// Defines payment states",
		"internal/domain/customer/customer.go":               "package customer\n\n// Customer entity definition",
		"internal/domain/order/order.go":                     "package order\n\n// Order entity definition",
		"internal/application/checkout/checkout_service.go":  "package checkout\n\n// Checkout service logic",
		"internal/application/checkout/checkout_handler.go":  "package checkout\n\n// HTTP handler for checkout",
		"internal/application/checkout/saga_manager.go":      "package checkout\n\n// Saga manager specific to checkout",
		"internal/application/payments/payment_service.go":   "package payments\n\n// Payment service logic",
		"internal/application/payments/payment_handler.go":   "package payments\n\n// HTTP handler for payments",
		"internal/application/customers/customer_service.go": "package customers\n\n// Customer service logic",
		"internal/application/customers/customer_handler.go": "package customers\n\n// HTTP handler for customers",
		"internal/application/uow/transaction_manager.go":    "package uow\n\n// Transaction manager for unit of work",
		"pkg/shared/db_helpers.go":                           "package shared\n\n// Common database helpers",
		"pkg/shared/validators.go":                           "package shared\n\n// Validation functions",
		"pkg/shared/logger.go":                               "package shared\n\n// Logger utility",
		"pkg/middleware/auth_middleware.go":                  "package middleware\n\n// Authentication middleware",
		"configs/config.yaml":                                "# Main configuration file",
		"configs/db_config.yaml":                             "# Database configuration",
		"scripts/db_migration.sh":                            "#!/bin/bash\n\n# Database migration script",
		"test/integration/checkout_service_test.go":          "package integration\n\n// Integration test for checkout service",
		"test/unit/payment_repository_test.go":               "package unit\n\n// Unit test for payment repository",
		"go.mod":                                             "module your-saas-app\n\n// Go module definition",
	}

	// Create directories
	for _, folder := range folders {
		path := filepath.Join(baseDir, folder)
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", path, err)
		}
	}

	// Create files with package declarations
	for file, content := range files {
		path := filepath.Join(baseDir, file)
		f, err := os.Create(path)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", path, err)
			continue
		}
		if _, err := f.WriteString(content + "\n"); err != nil {
			fmt.Printf("Error writing to file %s: %v\n", path, err)
		}
		f.Close()
	}

	fmt.Printf("Project structure for '%s' created successfully!\n", baseDir)
}
