package auth

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var credentialsService CredentialsService

var apiKeyConfigMock *ApiKeyModel

type ApiKeyRepoMock struct {
}

func (apiKeyRepo *ApiKeyRepoMock) GetApiKeyConfig() *ApiKeyModel {
	return apiKeyConfigMock
}

func TestCredentialsServiceSpec(t *testing.T) {
	Convey("CredentialsService Tests", t, func() {

		apiKeyConfigMock = &ApiKeyModel{PrivateKey: "foo", PublicKey: "bar"}

		credentialsService = &CredentialsServiceImpl{ApiKeyRepoInferace: &ApiKeyRepoMock{}}

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

					So(credentialsModel.Hash, ShouldEqual, "bar")
				})

				Convey("Should not return empty TimeStamp", func() {

					credentialsModel := credentialsService.GenerateCredentials()

					So(credentialsModel.TimeStamp, ShouldEqual, "bar")
				})

			})

		})

	})
}
