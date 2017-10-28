package actor

import (
	"fmt"
)

var actorSystemInstance actorSystem = actorSystem{}

//func init() {
//	System().actorNames = make(map[string]actorRefInterface)
//	System().actors = make(map[actorRefInterface]actorInterface)
//}

func ActorSystem() *actorSystem {
	return &actorSystemInstance
}

type actorSystem struct {
	actorNames map[string]actorRefInterface
	actors     map[actorRefInterface]actorInterface
}

type actorSystemInterface interface {
	RegisterActor(actorInterface) error
	Actor(string) (actorRefInterface, error)
	actor(actorRefInterface) (actorInterface, error)
}

func (system *actorSystem) RegisterActor(name string, actor actorInterface) error {
	if system.actorNames == nil {
		system.actorNames = make(map[string]actorRefInterface)
		system.actors = make(map[actorRefInterface]actorInterface)
	}

	_, exists := system.actorNames[name]
	if exists {
		return fmt.Errorf("Actor %v already registered")
	}

	actorRef := &ActorRef{name}
	system.actorNames[name] = *actorRef
	system.actors[*actorRef] = actor

	actor.setMailbox(make(chan Message))
	go receive(actor)

	return nil
}

func (system *actorSystem) Actor(name string) (actorRefInterface, error) {
	ref, exists := system.actorNames[name]
	if !exists {
		return nil, fmt.Errorf("Actor %v not registered", name)
	}

	return ref, nil
}

func (system *actorSystem) actor(actorRef actorRefInterface) (actorInterface, error) {
	ref, exists := system.actors[actorRef]

	if !exists {
		return nil, fmt.Errorf("Actor %v not registered", actorRef)
	}

	return ref, nil
}
