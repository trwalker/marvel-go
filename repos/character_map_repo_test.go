package repos

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var characterMapRepo CharacterMapRepo

func TestCharacterMapRepoSpec(t *testing.T) {

	Convey("CharacterMapRepo Tests", t, func() {

		CharacterMapRepoInstance = constructor()
		characterMapRepo = CharacterMapRepoInstance

		Convey("GetCharacterMap Function", func() {

			Convey("When valid state", func() {

				characterMap := characterMapRepo.GetCharacterMap()

				Convey("Should contain 11 characters", func() {
					So(len(characterMap), ShouldEqual, 11)
				})

				Convey("Should contain \"spider-man\" character", func() {
					_, found := characterMap["spider-man"]
					So(found, ShouldBeTrue)
				})

			})

		})

	})

}
