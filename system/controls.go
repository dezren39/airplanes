package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"

	"github.com/m110/airplanes/archetypes"
	"github.com/m110/airplanes/component"
)

type Controls struct {
	query *query.Query
}

func NewControls() *Controls {
	return &Controls{
		query: query.NewQuery(
			filter.Contains(
				component.Position,
				component.Input,
				component.Velocity,
				component.Sprite,
				component.PlayerAirplane,
			),
		),
	}
}

func (i *Controls) Update(w donburi.World) {
	i.query.EachEntity(w, func(entry *donburi.Entry) {
		input := component.GetInput(entry)

		if input.Disabled {
			return
		}

		velocity := component.GetVelocity(entry)

		velocity.X = 0
		velocity.Y = -0.5

		if ebiten.IsKeyPressed(input.MoveUpKey) {
			velocity.Y = -input.MoveSpeed
		} else if ebiten.IsKeyPressed(input.MoveDownKey) {
			velocity.Y = input.MoveSpeed
		}

		if ebiten.IsKeyPressed(input.MoveRightKey) {
			velocity.X = input.MoveSpeed
		}
		if ebiten.IsKeyPressed(input.MoveLeftKey) {
			velocity.X = -input.MoveSpeed
		}

		// TODO Seems like a very complex way to get the weapon level and timer
		airplane := component.GetPlayerAirplane(entry)
		player := archetypes.MustFindPlayerByNumber(w, airplane.PlayerNumber)
		player.ShootTimer.Update()
		if ebiten.IsKeyPressed(input.ShootKey) && player.ShootTimer.IsReady() {
			position := component.GetPosition(entry).Position

			archetypes.NewBullet(w, player, position)
			player.ShootTimer.Reset()
		}
	})
}
