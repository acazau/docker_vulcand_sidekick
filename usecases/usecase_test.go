package usecases

import (
	"errors"
	"github.com/acazau/docker_vulcand_sidekick/domain"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type _Mock_Good_UseCaseManager struct{}
type _Mock_Bad_UseCaseManager struct{}

func (manager *_Mock_Good_UseCaseManager) Execute(config *domain.Config) error {
	return nil
}

func (manager *_Mock_Bad_UseCaseManager) Execute(config *domain.Config) error {
	return errors.New("Bad UseCaseManager")
}

func TestExecute(t *testing.T) {
	Convey("Validate Execute", t, func() {

		Convey("Validate when UseCaseManager is null, returns errors", func() {
			useCaseManager := new(UseCaseManager)
			err := useCaseManager.Execute(new(domain.Config))
			So(err, ShouldNotBeNil)
		})

		Convey("Validate when UseCaseManager is valid, returns no errors", func() {
			useCaseManager := &UseCaseManager{InjectedUseCaseManager: &_Mock_Good_UseCaseManager{}}
			err := useCaseManager.Execute(new(domain.Config))
			So(err, ShouldBeNil)
		})

		Convey("Validate when UseCaseManager is not valid, returns errors", func() {
			useCaseManager := &UseCaseManager{InjectedUseCaseManager: &_Mock_Bad_UseCaseManager{}}
			err := useCaseManager.Execute(new(domain.Config))
			So(err, ShouldNotBeNil)
		})

	})
}
