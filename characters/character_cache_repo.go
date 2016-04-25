package characters

type CharacterCacheRepo interface {
	Get(id int) (characterDetails *CharacterModel, found bool)
	Add(characterDetails *CharacterModel)
}
