package main

func squareIsWhite(coordinates string) bool {
	if (coordinates[0]-'a')%2 == 0 {
		return (coordinates[1]-'1')%2 == 1
	} else {
		return (coordinates[1]-'1')%2 == 0
	}
}