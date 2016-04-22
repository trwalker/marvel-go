package characters

type CharacterCacheRepo interface {
	Get(id int) (characterDetails *CharacterDetailsModel, found bool)
	Add(characterDetails *CharacterDetailsModel)
}