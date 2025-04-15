package data

import (
	"github.com/WH-5/push-service/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewPushRepo)

// Data .
type Data struct {
	WSD *WSData
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	ws := NewWSData()
	return &Data{WSD: ws}, cleanup, nil
}
