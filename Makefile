.PHONY: prepare run-p1 run-p2 run-p4 run-api migrate

prepare:
	go mod tidy

run-p1:
	@echo "Running Problem 1 (Polycarp Sequence)"
	@echo "Enter number of test cases, followed by k for each case:"
	@go run p1/main.go

run-p2:
	@echo "Running Problem 2 (Palindrome Checker)"
	@echo "Enter a string to check:"
	@go run p2/main.go

run-p4:
	@echo "Running Problem 4 (Selection Sort)"
	@go run p4/main.go

run-api:
	@echo "Running Problem 3 API server on :8080"
	@go run p3/main.go

migrate:
	@echo "Running Database Migration for Problem 3"
	@go run p3/migration/main.go

migrate-fresh:
	@echo "Running Fresh Database Migration for Problem 3"
	@go run p3/migration/main.go -fresh

seed:
	@echo "Seeding Database (Users and Spendings)..."
	@go run p3/seeder/main.go
