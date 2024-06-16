package common

import (
	"newsletter/app/common/service"

	"go.uber.org/fx"
)

var CommonModule = fx.Options(
	fx.Provide(service.NewSlugGenerator),
	fx.Provide(service.NewUUIDGenerator),
)
