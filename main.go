package main

import (
	"fmt"
	"container/list"
)

type Route struct {
	From string
	To   string
}

type State struct {
	City          string
	PreviousState *State
}

func main() {
	citys := []string{
		"0-Arad",
		"1-Timiosara",
		"2-Lugoj",
		"3-Mehadia",
		"4-Dobreta",
		"5-Craiova",
		"6-Rimnicu Vilcea",
		"7-Sibiu",
		"8-Fagaras",
		"9-Pitesti",
		"10-Bucharest",
	}
	routes := make([]Route, 0)
	routes = append(routes, Route{From: citys[0], To: citys[1]})
	routes = append(routes, Route{From: citys[0], To: citys[7]})
	routes = append(routes, Route{From: citys[1], To: citys[2]})
	routes = append(routes, Route{From: citys[2], To: citys[3]})
	routes = append(routes, Route{From: citys[3], To: citys[4]})
	routes = append(routes, Route{From: citys[4], To: citys[5]})
	routes = append(routes, Route{From: citys[5], To: citys[6]})
	routes = append(routes, Route{From: citys[5], To: citys[9]})
	routes = append(routes, Route{From: citys[6], To: citys[9]})
	routes = append(routes, Route{From: citys[6], To: citys[7]})
	routes = append(routes, Route{From: citys[7], To: citys[8]})
	routes = append(routes, Route{From: citys[8], To: citys[10]})
	routes = append(routes, Route{From: citys[9], To: citys[10]})

	fmt.Println(routes)
	visitedGo := make([]State, 0)
	visitedBack := make([]State, 0)
	boundaryGo := list.New()
	boundaryBack := list.New()
	run(boundaryGo, citys, visitedGo, routes, boundaryBack, visitedBack)
}

func run(boundary *list.List, citys []string, visited []State, routes []Route, boundaryBack *list.List, visitedBack []State) {

	stateGo := State{
		City:          citys[0],
		PreviousState: nil,
	}
	fmt.Println("ADD BOUNDARY", stateGo)
	boundary.PushBack(stateGo)
	stateBack := State{
		City:          citys[10],
		PreviousState: nil,
	}
	boundaryBack.PushBack(stateBack)
	fmt.Println("ADD BOUNDARY", stateBack)
	fmt.Println("RUN")
	resultState := handleBoundary(boundary, visited, routes, boundaryBack, visitedBack)
	if resultState == nil {
		fmt.Println("GOAL NOT FOUND")
		return
	}
	fmt.Println("GOAL FOUND")
	for {
		if resultState.PreviousState == nil {
			break
		}
		resultState = resultState.PreviousState
		fmt.Print(resultState, " ")
	}
}

func handleBoundary(boundary *list.List, visited []State, routes []Route, boundaryBack *list.List, visitedBack []State) *State {
	for e := boundary.Front(); e != nil; e = e.Next() {
		boundary.Remove(e)
		currentCity := e.Value.(State)
		fmt.Println("GET", currentCity)
		visited = append(visited, currentCity)
		matchCity := checkBoundaries(boundaryBack, &currentCity)
		if matchCity != nil {
			return matchCity
		}
		for _, route := range routes {
			if currentCity.City == route.From {
				if checkVisited(visited, route.To) {
					continue
				}
				state := State{City: route.To, PreviousState: &currentCity}
				fmt.Println("ADD BOUNDARY", state)
				boundary.PushBack(state)
			}
			if currentCity.City == route.To {
				if checkVisited(visited, route.From) {
					continue
				}
				state := State{City: route.From, PreviousState: &currentCity}
				fmt.Println("ADD BOUNDARY", state)
				boundary.PushBack(state)
			}
		}
		printLog(visited, boundary)
	}
	if boundary.Front() == nil {
		return nil
	}
	return handleBoundary(boundaryBack, visitedBack, routes, boundary, visited)
}

func printLog(visited []State, boundary *list.List) {
	fmt.Println("VISITED", visited, "<- CURRENT CITY")
	if boundary.Len() > 0 {
		fmt.Print("BOUNDARY ")
		for e := boundary.Front(); e != nil; e = e.Next() {
			fmt.Print(e.Value.(State).City, " ")
		}
		fmt.Println()
	}
}

func checkBoundaries(boundaryBack *list.List, currentCity *State) *State {
	for e2 := boundaryBack.Front(); e2 != nil; e2 = e2.Next() {
		currentCity2 := e2.Value.(State)
		if currentCity.City == currentCity2.City {
			fmt.Println("YEAAHHH!!!!!")
			fmt.Println(currentCity)
			return mirrorPrevious(currentCity, currentCity2.PreviousState)
		}
	}
	return nil
}

func mirrorPrevious(currentCity *State, currentCity2 *State) *State{
	fmt.Println(currentCity2.City + ".NewPrevious->" +currentCity.City)
	if currentCity2.PreviousState != nil {
		fmt.Println(currentCity.PreviousState.City)
	}
	if currentCity2.PreviousState == nil {
		currentCity2.PreviousState = currentCity
		return currentCity2
	}
	aux := *currentCity.PreviousState
	currentCity2.PreviousState = currentCity
	return mirrorPrevious(currentCity2, &aux)
}

func checkVisited(visiteds []State, city string) bool {
	for _, visited := range visiteds {
		if visited.City == city {
			return true
		}
	}
	return false
}
