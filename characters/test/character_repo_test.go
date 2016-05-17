package characterstest

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/trwalker/marvel-go/auth"
	"github.com/trwalker/marvel-go/characters"
	"github.com/trwalker/marvel-go/rest/test"
	"net/http"
	"testing"
	"time"
)

func TestCharacterRepoSpec(t *testing.T) {

	Convey("CharacterRepo Tests", t, func() {

		credentialsModelMock := &auth.CredentialsModel{}
		restClientAdapterMock := &resttest.RestClientAdapterMock{}
		restClientAdapterMock.GetMock = func(url string, timeout time.Duration) (resp *http.Response, body string, err error) {
			resp = &http.Response{StatusCode: 200}
			body = SpiderManResponseMock
			err = nil

			return
		}

		characterRepo := characters.NewCharacterRepo(restClientAdapterMock)

		Convey("GetCharacter Function", func() {

			Convey("When invalid state", func() {

				Convey("When 404 response", func() {

					restClientAdapterMock.GetMock = func(url string, timeout time.Duration) (resp *http.Response, body string, err error) {
						resp = &http.Response{StatusCode: 404}
						body = "{}"
						err = nil

						return
					}

					Convey("Should set found to false", func() {
						_, found, _ := characterRepo.GetCharacter(1234, credentialsModelMock)

						So(found, ShouldBeFalse)
					})

					Convey("Should set err to nil", func() {
						_, _, err := characterRepo.GetCharacter(1234, credentialsModelMock)

						So(err, ShouldBeNil)
					})

					Convey("Should set character to nil", func() {
						character, _, _ := characterRepo.GetCharacter(1234, credentialsModelMock)

						So(character, ShouldBeNil)
					})

				})

				Convey("When non 200 response", func() {

					restClientAdapterMock.GetMock = func(url string, timeout time.Duration) (resp *http.Response, body string, err error) {
						resp = &http.Response{StatusCode: 500}
						body = ""
						err = errors.New("Server Error")

						return
					}

					Convey("Should set found to false", func() {
						_, found, _ := characterRepo.GetCharacter(1234, credentialsModelMock)

						So(found, ShouldBeFalse)
					})

					Convey("Should set err to server error", func() {
						_, _, err := characterRepo.GetCharacter(1234, credentialsModelMock)

						So(err.Error(), ShouldEqual, "Server Error")
					})

					Convey("Should set character to nil", func() {
						character, _, _ := characterRepo.GetCharacter(1234, credentialsModelMock)

						So(character, ShouldBeNil)
					})

				})

			})

			Convey("When valid state", func() {

				Convey("Should set found to true", func() {
					_, found, _ := characterRepo.GetCharacter(1234, credentialsModelMock)

					So(found, ShouldBeTrue)
				})

				Convey("Should set err to nil", func() {
					_, _, err := characterRepo.GetCharacter(1234, credentialsModelMock)

					So(err, ShouldBeNil)
				})

				Convey("Should set character name", func() {
					character, _, _ := characterRepo.GetCharacter(1234, credentialsModelMock)

					So(character.Name, ShouldEqual, "Spider-Man")
				})

				Convey("Should set character description", func() {
					character, _, _ := characterRepo.GetCharacter(1234, credentialsModelMock)

					So(character.Description, ShouldEqual, "Bitten by a radioactive spider, high school student Peter Parker gained the speed, strength and powers of a spider. Adopting the name Spider-Man, Peter hoped to start a career using his new abilities. Taught that with great power comes great responsibility, Spidey has vowed to use his powers to help people.")
				})

				Convey("Should set character id", func() {
					character, _, _ := characterRepo.GetCharacter(1234, credentialsModelMock)

					So(character.Id, ShouldEqual, 1009610)
				})

				Convey("Should set character image", func() {
					character, _, _ := characterRepo.GetCharacter(1234, credentialsModelMock)

					So(character.Image, ShouldEqual, "http://i.annihil.us/u/prod/marvel/i/mg/3/50/526548a343e4b.jpg")
				})

				Convey("Should set character comics", func() {
					character, _, _ := characterRepo.GetCharacter(1234, credentialsModelMock)

					So(len(character.Comics), ShouldEqual, 20)
				})

			})

		})

	})

}
