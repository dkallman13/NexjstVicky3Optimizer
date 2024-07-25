package initial

import (
	"log"
	"os"
	"strings"
	//"text/scanner"
)

var DecodedSaveFile []string
var SaveFiles []string
var SaveFileLoc string
var saveFileLocBuilder strings.Builder


func SaveFileLocSetter() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	saveFileLocBuilder.WriteString(homedir)
	saveFileLocBuilder.WriteString("\\Documents\\Paradox Interactive\\Victoria 3\\save games\\")
	SaveFileLoc = saveFileLocBuilder.String()
}

func SaveFileLister(){
	files, err := os.ReadDir(SaveFileLoc)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		SaveFiles = append(SaveFiles, file.Name())
	}
}
func TokenizeSave(saveFileName string) string{
	for _, file := range SaveFiles{
		switch file == saveFileName {
		case true:
			var saveFileBuilder strings.Builder
			//var scan scanner.Scanner
			var rawFileTextBuilder strings.Builder
			saveFileBuilder.WriteString(SaveFileLoc)
			saveFileBuilder.WriteString(saveFileName)
			rawFile, err := os.ReadFile(saveFileBuilder.String())
			if err != nil {
				log.Fatal(err)
			}
			for _, charbyte := range rawFile{
				if(charbyte == 10){
					break
				}
				rawFileTextBuilder.WriteByte(charbyte)
			}
			return rawFileTextBuilder.String()
			//scan.Init(strings.NewReader(rawFileTextBuilder.String()))
			//scan.Whitespace = 1<<'\n'
		}
	}
	return "fail"
}

func SaveFileRead() {

}
