package machine

import (
	"fmt"
	"encoding/json"

	"github.com/dockersecuritytools/batten/batten"
)

type JSONResult struct {
	CheckResults	[]CheckResult
}

type CheckResult struct {
	Checkid	int		`json:"id"`
	Result	string	`json:"result"`
	Name	string	`json:"desc"`
}


func FormatResultsForMachine(idx int, results *batten.CheckResults) CheckResult {
	
	checkdefinition := results.CheckDefinition

    var resultString string	

	if results.Error != nil {
		resultString = "ERROR"
	} else {
		if results.Success {
			resultString = "PASSED"
		} else {
			resultString = "FAILED"
		}
	}

	return CheckResult{
		Checkid:	idx+1,
		Result:		resultString,
		Name:		checkdefinition.Name(),
	}
}

func ShowResults(results *JSONResult) {
	fmt.Printf("%s\n","{'checks:'")
	fmt.Printf("%s","{")
	for _, result := range results.CheckResults {
		jsonString,_ := json.Marshal(result)
		fmt.Println(string(jsonString))
	}
	fmt.Printf("}\n%s","}")
}
