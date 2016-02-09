package repos

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCharacterMapRepo(t *testing.T) {
	Convey("CharacterMapRepo Tests", t, func() {

		var charMapRepo CharacterMapRepo = CharacterMapRepoInstance

		Convey("GetCharacterMap Method", func() {

			Convey("When valid state", func() {

				characterMap := charMapRepo.GetCharacterMap()

				Convey("Should contain 11 Characters", func() {
					So(len(characterMap), ShouldEqual, 11)
				})

				Convey("Should contain \"spider-man\"", func() {
					_, found := characterMap["spider-man"]
					So(found, ShouldBeTrue)
				})

			})

		})

	})
}
