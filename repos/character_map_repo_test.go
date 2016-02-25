package repos

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/trwalker/marvel-go/models"
	"testing"
)

var characterMapRepo CharacterMapRepo

func TestCharacterMapRepoSpec(t *testing.T) {

	Convey("CharacterMapRepo Tests", t, func() {

		characterMapRepo = &CharacterMapRepoImpl{characterMap: make(map[string]*models.CharacterModel)}

		Convey("GetCharacterMap Function", func() {

			Convey("When valid state", func() {

				characters := characterMapRepo.GetCharacterMap()

				Convey("Should contain 11 characters", func() {
					So(len(characters), ShouldEqual, 11)
				})

				Convey("Should contain \"spider-man\" character", func() {
					_, found := characters["spider-man"]
					So(found, ShouldBeTrue)
				})

			})

		})

	})

}
