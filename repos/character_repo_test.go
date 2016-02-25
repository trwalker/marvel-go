package repos

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var characterRepo CharacterRepo

func TestCharacterRepoSpec(t *testing.T) {

	Convey("CharacterRepo Tests", t, func() {

		characterRepo = &CharacterRepoImpl{}

		Convey("GetCharacter Function", func() {

			Convey("When valid state", func() {

				Convey("Should not return nil CharacterDetailsModel", func() {

					characterDetails := characterRepo.GetCharacter(1234)

					So(characterDetails, ShouldNotEqual, nil)

				})

			})

		})

	})

}
