package modules_test

import (
	"testing"

	"test-golang/controllers/requests"
	"test-golang/modules"
	conveter "test-golang/utils/converter"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestClient(t *testing.T) {
	Convey("Client", t, func() {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		converterHandler := modules.NewConverterHandler()
		requestTest := requests.UrlRequest{
			RequestUrl: "test",
		}

		Convey("StatsUrl()", func() {
			Convey("When data not found", func() {
				conveter.ShortenUrlMemory = nil
				Convey("It returns error response", func() {
					response, err := converterHandler.StatsUrl(requestTest)
					So(err, ShouldNotBeNil)
					So(response, ShouldBeNil)
				})
			})

			Convey("When data is exist", func() {
				conveter.ShortenUrlMemory = []*conveter.ShortenUrlData{
					&conveter.ShortenUrlData{
						ShortUrl: "test",
					},
				}

				Convey("It returns error response", func() {
					response, err := converterHandler.StatsUrl(requestTest)
					So(err, ShouldBeNil)
					So(response, ShouldNotBeNil)
				})
			})
		})
	})
}
