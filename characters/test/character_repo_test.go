package characterstest

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/trwalker/marvel-go/auth"
	"github.com/trwalker/marvel-go/characters"
	"github.com/trwalker/marvel-go/rest/test"
	"testing"
)

func TestCharacterRepoSpec(t *testing.T) {

	Convey("CharacterRepo Tests", t, func() {

		credentialsModelMock := &auth.CredentialsModel{}
		restClientAdapterMock := &resttest.RestClientAdapterMock{}

		characterRepo := characters.NewCharacterRepo(restClientAdapterMock)

		Convey("GetCharacter Function", func() {

			Convey("When valid state", func() {

				Convey("Should not return nil CharacterDetailsModel", func() {
					character, _, _ := characterRepo.GetCharacter(1234, credentialsModelMock)

					So(character, ShouldNotEqual, nil)

				})

			})

		})

	})

}
