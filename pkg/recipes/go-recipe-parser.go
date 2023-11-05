package recipes

type GoRecipeParser struct {
	itemParser ItemsParser
}

func NewGoRecipeParser(itemParser ItemsParser) *GoRecipeParser {
	return &GoRecipeParser{
		itemParser,
	}
}
