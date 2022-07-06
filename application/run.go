package application

import "sut-order-go/config"

func (app *Application) Run(cfg *config.Config) error {
	return grpcRun(cfg)(app)
}
