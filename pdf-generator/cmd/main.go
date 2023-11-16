package main

import (
	"strings"

	"github.com/FavorDespaches/pdf-generator/pkg/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

/*
	func printAndReturnContents(dirPath string) string {
		var builder strings.Builder

		// Read the directory contents
		files, err := ioutil.ReadDir(dirPath)
		if err != nil {
			log.Fatalf("Error reading directory %s: %v", dirPath, err)
		}

		// Loop through and append each file or directory name to the builder
		for _, file := range files {
			fullPath := filepath.Join(dirPath, file.Name())
			if file.IsDir() {
				fmt.Println("Directory:", fullPath)
				builder.WriteString("Directory: " + fullPath + "\n")
				// Optionally, you can recursively call this function for sub-directories
			} else {
				fmt.Println("File:", fullPath)
				builder.WriteString("File: " + fullPath + "\n")
			}
		}

		return builder.String()
	}
*/
func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch strings.ToUpper(req.HTTPMethod) {
	case "GET":
		return handlers.HandleUnsupportedMethod()
	case "POST":
		//handlers.SolicitarEtiqueta(req) //handlers.SolicitarEtiqueta(req)

		/*
			dirPath := "/opt/bin/images" // Replace with the path of your directory
			contents := printAndReturnContents(dirPath)
			successBody := handlers.SuccessResponse{
				IdPlp:       0,
				Etiquetas:   []string{""},
				LabelBase64: contents,
			}
		*/

		return handlers.SolicitarEtiqueta(req)
	case "PUT":
		return handlers.HandleUnsupportedMethod()
	case "DELETE":
		return handlers.HandleUnsupportedMethod()
	default:
		return handlers.HandleUnsupportedMethod()
	}
}
