package interfaces

import (
	"context"
	"ewallet-ums/external"
)

//go:generate mockgen -source=IExternalWallet.go -destination=../services/external_mock_test.go -package=services
type IExternal interface {
	CreateWallet(ctx context.Context, userID uint64) (*external.Wallet, error)
	SendNotification(ctx context.Context, recipient string, templateName string, placeHolder map[string]string) error
}
