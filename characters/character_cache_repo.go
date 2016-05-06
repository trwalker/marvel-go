package characters

type CharacterCacheRepo interface {
	Get(id int) (character *CharacterModel, found bool)
	Add(character *CharacterModel)
}
