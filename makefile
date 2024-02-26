run :
	go run ./cmd/api

wire: ## Generate wire_gen.go
	cd pkg/di && wire

swag: 
	swag init -g cmd/api/main.go -o ./cmd/api/docs
test:
	go test ./...

mock: ##make mock files using mockgen
	mockgen -source pkg/repository/interfaces/user.go -destination pkg/repository/mock/user_mock.go -package mock
	mockgen -source pkg/usecase/interfaces/user.go -destination pkg/usecase/mock/user_mock.go -package mock
	mockgen -source pkg/helper/interface/helper.go -destination pkg/helper/mock/helper_mock.go -package mock
	mockgen -source pkg/repository/interfaces/wallet.go -destination pkg/repository/mock/wallet_mock.go -package mock
	mockgen -source pkg/repository/interfaces/order.go -destination pkg/repository/mock/order_mock.go -package mock
	mockgen -source pkg/repository/interfaces/otp.go -destination pkg/repository/mock/otp_mock.go -package mock
	mockgen -source pkg/repository/interfaces/Inventory.go -destination pkg/repository/mock/inventory_mock.go -package mock
	mockgen 

