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

func main() {
}
