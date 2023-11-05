package recipes

type NpmRecipeParser struct {
	itemParser ItemsParser
}

func NewNpmRecipeParser(itemParser ItemsParser) *NpmRecipeParser {
	return &NpmRecipeParser{
		itemParser,
	}
}
