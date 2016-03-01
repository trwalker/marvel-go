package characters

type CharacterMapRepo interface {
	GetCharacterMap() map[string]*CharacterModel
}
