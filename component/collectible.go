package component

import "github.com/yohamta/donburi"

type CollectibleType int

const (
	CollectibleTypeHealth CollectibleType = iota
	CollectibleTypeWeaponUpgrade
	CollectibleTypeShield
)

type CollectibleData struct {
	Type CollectibleType
}

var Collectible = donburi.NewComponentType[CollectibleData]()

func GetCollectible(entry *donburi.Entry) *CollectibleData {
	return donburi.Get[CollectibleData](entry, Collectible)
}
