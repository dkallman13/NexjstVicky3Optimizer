package initial

import (
	"log"
	"os"
	"strings"

	//"github.com/dkallman13/Vicky3Optimizer/types"
	"github.com/bzick/tokenizer"
)

const (
	TokenCurlyOpen  = 1
	TokenCurlyClose = 2
	TokenEquals     = 3
)

// var DecodedSaveFile SaveFile
var SaveFiles []string
var SaveFileLoc string
var saveFileLocBuilder strings.Builder
var parser *tokenizer.Tokenizer

func SaveFileLocSetter() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	saveFileLocBuilder.WriteString(homedir)
	saveFileLocBuilder.WriteString("\\Documents\\Paradox Interactive\\Victoria 3\\save games\\")
	SaveFileLoc = saveFileLocBuilder.String()
}

func SaveFileLister() {
	files, err := os.ReadDir(SaveFileLoc)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		SaveFiles = append(SaveFiles, file.Name())
	}
}

func TokenizeSave(saveFileName string) string {
	for _, file := range SaveFiles {
		switch file == saveFileName {
		case true:
			var saveFileBuilder strings.Builder
			var rawFileTextBuilder strings.Builder
			saveFileBuilder.WriteString(SaveFileLoc)
			saveFileBuilder.WriteString(saveFileName)
			rawFile, err := os.ReadFile(saveFileBuilder.String())
			
			
			parser := tokenizer.New()
			parser.AllowKeywordSymbols(tokenizer.Underscore, tokenizer.Numbers)
			parser.DefineTokens(TokenCurlyOpen, []string{"{"})
			parser.DefineTokens(TokenCurlyClose, []string{"}"})
			parser.DefineTokens(TokenEquals, []string{"="})
			stream := parser.ParseBytes(rawFile)
			

			if err != nil {
				log.Fatal(err)
			}
			x := 0
			for stream.IsValid() {
				if stream.CurrentToken().Is(tokenizer.TokenKeyword) {
					rawFileTextBuilder.WriteString(stream.CurrentToken().ValueString())
					x++
				}
				if(x>5){
					break
				}
				stream.GoNext()
			}
			stream.Close()
			return rawFileTextBuilder.String()
		}
	}
	return "fail"
}
