package common

import (
	"go.uber.org/fx"
)

var CommonModule = fx.Options(
	fx.Provide(NewSlugGenerator),
	fx.Provide(NewUUIDGenerator),
)
