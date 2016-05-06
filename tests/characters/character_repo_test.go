package characterstests

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/trwalker/marvel-go/characters"
	"testing"
)

var characterRepo characters.CharacterRepo

func TestCharacterRepoSpec(t *testing.T) {

	Convey("CharacterRepo Tests", t, func() {

		characterRepo = &characters.CharacterRepoImpl{}

		Convey("GetCharacter Function", func() {

			Convey("When valid state", func() {

				Convey("Should not return nil CharacterDetailsModel", func() {

					//characterDetails := characterRepo.GetCharacter(1234)

					//So(characterDetails, ShouldNotEqual, nil)

				})

			})

		})

	})

}
