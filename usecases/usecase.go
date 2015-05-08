package usecases

import (
	"errors"
	"fmt"
	"log"
)

type IUseCaseManager interface {
	Execute(arg map[string]interface{}) error
}

type UseCaseManager struct {
	InjectedUseCaseManager IUseCaseManager
}

func (manager *UseCaseManager) Execute(arg map[string]interface{}) error {
	if manager.InjectedUseCaseManager == nil {
		return errors.New("Injected InjectedUseCaseManager cannot be null")
	}

	if err := manager.InjectedUseCaseManager.Execute(arg); err != nil {
		log.Println(fmt.Sprintf("Error executing UseCaseManager %s", err.Error()))
	}

	return nil
}
