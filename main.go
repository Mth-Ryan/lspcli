package main

import (
	"fmt"

	"github.com/Mth-Ryan/lspcli/core"
)

const VERSION = "0.1.0"
const RECIPES = "runtime/recipes.yml"

func versionDialog() {
	fmt.Printf("lspcli %s\n", VERSION)
}

func printRecipes() {
	recipes, err := core.LoadAndParseRecipes(RECIPES)
	if err != nil {
		fmt.Println(err)
	}

	for idx, recipe := range recipes {
		fmt.Printf("%d - %v\n", idx, recipe)
	}
}

func main() {
	printRecipes()
}
