package charrepos

type CharacterMapRepo interface {
	GetCharacterMap() map[string]int
}
