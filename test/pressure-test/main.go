package main

import (
	"_template_/proto/helloworld"
	"fmt"
	"github.com/bojand/ghz/printer"
	"github.com/bojand/ghz/runner"
	"os"
)



func main() {
	report, err := runner.Run(
		"helloworld.HelloWorld.GetResult",
		"localhost:8000",
		runner.WithProtoFile("helloworld.proto", []string{`./proto/helloworld`}),
		runner.WithData(helloworld.GetResultRequest{
			Text: `0.0.0.0`,
		}),
		runner.WithInsecure(true),
		runner.WithTotalRequests(1000),
	)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	printer := printer.ReportPrinter{
		Out:    os.Stdout,
		Report: report,
	}

	printer.Print("summary")
}
