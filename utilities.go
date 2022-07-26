package utilities

import (
	_ "fmt"
	"time"
)

type DataObject struct {
	name, botType                    string
	workDone, attack, defense, speed int
	createdAt, updatedAt             time.Time
}

type DestroyerType struct {
	name, botType                  string
	health, attack, defense, speed int
	DestroyerInterface
}

type DestroyerCoreProperties struct {
	description, name string
	eta               int
}
type DestroyerChoreMethodType struct {
	DestroyerCoreProperties
}

type DestroyerAttackMethodType struct {
	DestroyerCoreProperties
	attack int
}

type DestroyerDefenseMethodType struct {
	DestroyerCoreProperties
	defense int
}

type DestroyerStatMethodType struct {
	description, health string
	eta                 int
}

type DestroyerInterface interface {
	bakeSomeCookies(name string, botType string) DestroyerChoreMethodType
	doTheDishes(name string, botType string) DestroyerChoreMethodType
	doTheLaundry(name string, botType string) DestroyerChoreMethodType
	makeASammich(name string, botType string) DestroyerChoreMethodType
	sweetTheHouse(name string, botType string) DestroyerChoreMethodType
	giveTheDogABath(name string, botType string) DestroyerChoreMethodType
	mowTheLaw(name string, botType string) DestroyerChoreMethodType
	rakeTheLeaves(name string, botType string) DestroyerChoreMethodType
	takeOutTheRecycling(name string, botType string) DestroyerChoreMethodType
	washTheCar(name string, botType string) DestroyerChoreMethodType
	attackValue(name string) DestroyerAttackMethodType
	defenseValue(name string) DestroyerDefenseMethodType
	healthValue() int
	speedValue(name string) struct {
		speed, eta int
	}
	stats(name string) DestroyerStatMethodType
}


type WorkTaskProp struct {
	workDone int
	currentTask string
	nextTask int
	choreList string
	taskIsComplete bool
}

type CounterProp struct {
	choreClick int
	submitClick int
	progressInterval int
}

type ExecutionerStateProps struct {
	workTasks WorkTaskProp
	counters CounterProp
	ExecutionerStateInterface
}

type ExecutionerStateInterface interface {
	//TODO: These interface methods need better return types. Since these will be used outside of the context of react initially - more exporation needed.
	setWorkTasks(workTask WorkTaskProp) any
	setCounters(counters CounterProp) any
}

func Executioner (nextArray []string, bot DestroyerType, scoreUpdate any, count int, state ExecutionerStateProps) {
	executionCount := count

	//BUGFIX: wrong approach - executioner needs full refactor as it is a react based function. State should be sent via API call, and handled here, with fully transformed data supplied to FE
	// state.setWorkTasks()
}

func main() {
}
