package repos

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var apiKeyRepo ApiKeyRepo

func TestApiKeyRepoSpec(t *testing.T) {
	Convey("ApiKeyRepo Tests", t, func() {

		apiKeyRepo = &ApiKeyRepoImpl{apiKeyConfigModel: nil}

		Convey("GetApiKeyConfig Function", func() {

			Convey("When valid state", func() {

				Convey("Should not return nil ApiKeyConfigModel", func() {

					config := apiKeyRepo.GetApiKeyConfig()

					So(config, ShouldNotEqual, nil)

				})

				Convey("Should not return empty public key", func() {

					config := apiKeyRepo.GetApiKeyConfig()

					So(len(config.PublicKey), ShouldBeGreaterThan, 0)

				})

				Convey("Should not return empty private key", func() {

					config := apiKeyRepo.GetApiKeyConfig()

					So(len(config.PrivateKey), ShouldBeGreaterThan, 0)

				})

			})

		})

	})
}
