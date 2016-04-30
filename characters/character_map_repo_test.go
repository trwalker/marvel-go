package characters

import (
	. "github.com/smartystreets/goconvey/convey"
	"sync"
	"testing"
)

func TestCharacterMapRepoSpec(t *testing.T) {

	Convey("CharacterMapRepo Tests", t, func() {

		characterMapRepo := &CharacterMapRepoImpl{
			characterMap: make(map[string]int),
			lock:         &sync.Mutex{},
		}

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
