package system

import (
	"math/rand"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"

	"github.com/m110/airplanes/archetypes"
	"github.com/m110/airplanes/component"
)

type Health struct {
	query *query.Query
}

func NewHealth() *Health {
	return &Health{
		query: query.NewQuery(filter.Contains(component.Health)),
	}
}

func (h *Health) Update(w donburi.World) {
	h.query.EachEntity(w, func(entry *donburi.Entry) {
		health := component.GetHealth(entry)

		if health.JustDamaged {
			health.DamageIndicatorTimer.Update()
			if health.DamageIndicatorTimer.IsReady() {
				health.JustDamaged = false
			}
		} else {
			if health.Health <= 0 {
				r := rand.Intn(10)
				if r < 7 {
					archetypes.NewRandomCollectible(w, component.GetTransform(entry).LocalPosition)
				}

				// TODO: It seems like a good candidate to be triggered by an event.
				component.MustFindGame(w).AddScore(1)
				Destroy(w, entry)
			}
		}
	})
}
