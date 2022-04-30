package outputs

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

var (
	spaces *regexp.Regexp = regexp.MustCompile(`\s+`)
)

func Budget(purchaseDate, price, description string) string {
	cleanedDecs := spaces.ReplaceAllString(description, "_")
	cleanedDecs = strings.ToLower(cleanedDecs)
	parsedDate, err := time.Parse("02.01.2006", purchaseDate)
	if err != nil {
		panic(err)
	}
	strippedPrice := strings.TrimLeft(price, "-")

	return fmt.Sprintf("osto %s %s %s",
		cleanedDecs,
		parsedDate.Format("01-2006"),
		strippedPrice)
}
