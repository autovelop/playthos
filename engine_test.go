package engine_test

import (
	"github.com/autovelop/playthos"
	"testing"
)

func TestNewEngine(t *testing.T) {
	eng := engine.New("TestNewEngine")
	eng.Once()
}

func TestNewEntity(t *testing.T) {
	eng := engine.New("TestNewEntity")

	eng.NewEntity()
	if len(eng.Entities()) != 1 {
		t.Fatal("new entity created but engine slice is empty", eng.Entities())
	}
}

func TestFindEntityByID(t *testing.T) {
	eng := engine.New("TestFindEntityByID")

	ent := eng.NewEntity()
	entFetch := eng.Entity(ent.ID())
	if ent != entFetch {
		t.Fatal("expected", ent, "got", entFetch)
	}
}

func TestDeleteEntity(t *testing.T) {
	eng := engine.New("TestDeleteEntity")

	ent := eng.NewEntity()
	eng.DeleteEntity(ent)
	if len(eng.Entities()) == 1 {
		t.Fatal("entity delete but engine slice not empty", eng.Entities())
	}
}

func TestNewUpdaterSystem(t *testing.T) {
	eng := engine.New("TestNewUpdaterSystem")
	engine.NewSystem(&MockUpdaterSystem{})
	if len(eng.Updaters()) != 1 {
		t.Fatal("new updater system created but engine slice is empty", eng.Updaters())
	}
}

func TestNewDrawerSystem(t *testing.T) {
	eng := engine.New("TestNewDrawerSystem")
	engine.NewSystem(&MockDrawerSystem{})
	if eng.Drawer() == nil {
		t.Fatal("new drawer system created but engine returns nil", eng.Drawer())
	}
}

// TODO(F): Write plain integrant tests

func TestNewListenerIntegrant(t *testing.T) {
	eng := engine.New("TestNewListenerIntegrant")
	engine.NewIntegrant(&MockListenerIntegrant{})
	if len(eng.Listeners()) != 1 {
		t.Fatal("new listener integrant created but engine slice is empty", eng.Listeners())
	}
}

func TestNewPlatformerIntegrant(t *testing.T) {
	eng := engine.New("TestNewPlatformerIntegrant")
	engine.NewIntegrant(&MockPlatformerIntegrant{})
	if eng.Platformer() == nil {
		t.Fatal("new platformer integrant created but engine returns nil", eng.Platformer())
	}
}
