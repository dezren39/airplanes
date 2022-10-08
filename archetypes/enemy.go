package archetypes

import (
	"time"

	"github.com/m110/airplanes/engine"
	"github.com/samber/lo"
	"github.com/yohamta/donburi"

	"github.com/m110/airplanes/assets"
	"github.com/m110/airplanes/component"
)

func NewEnemy(
	w donburi.World,
	position component.PositionData,
	rotation float64,
	speed float64,
	path []assets.Position,
) *donburi.Entry {
	enemy := w.Entry(
		w.Create(
			component.Position,
			component.Rotation,
			component.Velocity,
			component.Sprite,
			component.AI,
			component.Despawnable,
			component.Collider,
			component.Health,
		),
	)

	donburi.SetValue(enemy, component.Position, position)
	component.GetRotation(enemy).Angle = rotation

	image := assets.AirplaneGraySmall
	donburi.SetValue(enemy, component.Sprite, component.SpriteData{
		Image: image,
		Layer: component.SpriteLayerUnits,
		Pivot: component.SpritePivotCenter,
	})

	width, height := image.Size()

	donburi.SetValue(enemy, component.Collider, component.ColliderData{
		Width:  float64(width),
		Height: float64(height),
		Layer:  component.CollisionLayerEnemies,
	})

	if len(path) > 0 {
		componentPath := lo.Map(path, func(p assets.Position, _ int) component.PathPosition {
			return component.PathPosition(p)
		})

		donburi.SetValue(enemy, component.AI, component.AIData{
			Type:  component.AITypeFollowPath,
			Speed: speed,
			Path:  componentPath,
		})
	} else {
		donburi.SetValue(enemy, component.AI, component.AIData{
			Type:  component.AITypeConstantVelocity,
			Speed: speed,
		})
	}

	donburi.SetValue(enemy, component.Health, component.HealthData{
		Health:               3,
		DamageIndicatorTimer: engine.NewTimer(time.Millisecond * 100),
	})

	return enemy
}
