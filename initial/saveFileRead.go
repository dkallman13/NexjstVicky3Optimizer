package initial

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var DecodedSaveFile []string
var saveFileLoc string
var saveFileLocBuilder strings.Builder

func saveFileLocSetter() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	saveFileLocBuilder.WriteString(homedir)
	saveFileLocBuilder.WriteString("\\Documents\\Paradox Interactive\\Victoria 3\\save games\\ \n")
	saveFileLoc = saveFileLocBuilder.String()
}
func SaveFileRead() {
	saveFileLocSetter()
	fmt.Print(saveFileLoc)
}
