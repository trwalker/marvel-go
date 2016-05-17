package authmock

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/trwalker/marvel-go/auth"
	"testing"
)

func TestCredentialsServiceSpec(t *testing.T) {
	Convey("CredentialsService Tests", t, func() {

		apiKeyRepoMock := &ApiKeyRepoMock{}
		apiKeyRepoMock.GetApiKeyConfigMock = func() *auth.ApiKeyModel {
			return &auth.ApiKeyModel{PrivateKey: "foo", PublicKey: "bar"}
		}

		credentialsService := auth.NewCredentialsService(apiKeyRepoMock)

		Convey("GenerateCredentials Function", func() {

			Convey("When valid state", func() {

				Convey("Should not return nil CredentialsModel", func() {

					credentialsModel := credentialsService.GenerateCredentials()

					So(credentialsModel, ShouldNotEqual, nil)
				})

				Convey("Should not return empty PublicKey", func() {

					credentialsModel := credentialsService.GenerateCredentials()

					So(credentialsModel.PublicKey, ShouldEqual, "bar")
				})

				Convey("Should not return empty Hash", func() {

					credentialsModel := credentialsService.GenerateCredentials()

					So(len(credentialsModel.Hash), ShouldBeGreaterThan, 0)
				})

				Convey("Should not return empty TimeStamp", func() {
					credentialsModel := credentialsService.GenerateCredentials()

					So(len(credentialsModel.TimeStamp), ShouldBeGreaterThan, 0)
				})

			})

		})

	})
}
