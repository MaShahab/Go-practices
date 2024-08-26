package main

import (
	"github.com/key-R-hash/monsterslayer/actions"
	"github.com/key-R-hash/monsterslayer/interaction"
)

var currentRound = 0
var gameRounds = []interaction.RoundData{}

func main() {
	startGame()

	winner := ""

	for winner == "" {
		winner = executeRound()
	}

	endGame(winner)
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0

	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)
	var palyerAttackDmg int
	var playerHealValue int
	var monsterAttackDmg int

	if userChoice == "ATTACK" {
		palyerAttackDmg = actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		playerHealValue = actions.HealPlayer()
	} else {
		palyerAttackDmg = actions.AttackMonster(true)
	}

	monsterAttackDmg = actions.AttackPlayer()

	monsterHealth, playerHealth := actions.GetHealthAmounts()

	roundData := interaction.RoundData{
		Action:           userChoice,
		PlayerAttackDmg:  palyerAttackDmg,
		PlayerHealValue:  playerHealValue,
		MonsterAttackDmg: monsterAttackDmg,
		PlayerHealth:     playerHealth,
		MonsterHealth:    monsterHealth,
	}

	gameRounds = append(gameRounds, roundData)

	interaction.PringRoundStatistics(&roundData)

	if playerHealth <= 0 {
		return "MONSTER"
	} else if monsterHealth <= 0 {
		return "PLAYER"
	}

	return ""
}

func endGame(winner string) {
	interaction.DeclareWinner(winner)
	interaction.WriteLogFile(&gameRounds)
}
