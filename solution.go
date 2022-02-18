package solution

import (
	"github.com/kyokomi/emoji"
)

// GetMessage returns string Hello with world map emoji
func GetMessage() string {
	return emoji.Sprint("Hello :world_map:!")
}
