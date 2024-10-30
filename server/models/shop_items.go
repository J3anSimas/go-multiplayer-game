package models

import "github.com/J3anSimas/game_multiplayer_go/types"

type ShopItem struct {
	Title       string
	Description string
	Cost        int
	Attribute   types.ShopItemAttributeModifier
	Modifier    int
}
