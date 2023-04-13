package monobank_action

import (
	"MonobankTracker/pkg/currency/adapter"
	"context"
)

type Repository interface {
	Create(monoAction adapter.Action) (*adapter.Action, error)
	All(ctx context.Context) ([]adapter.Action, error)
	GetByName(ctx context.Context, name string) (*adapter.Action, error)
	Update(ctx context.Context, id int, updated adapter.Action) (*adapter.Action, error)
	Delete(ctx context.Context, id int) error
	NewAction(action *MonobankAction) *adapter.Action
}
