package machine

import (
	"fmt"
	"github.com/dockersecuritytools/batten/batten"
)

type JSONResult struct {
	CheckResults	[]CheckResult
}

type CheckResult struct {
	checkid	int
	result	string
	name	string
}


func FormatResultsForMachine(idx int, results *batten.CheckResults) CheckResult {
	
	checkdefinition := results.CheckDefinition

	//fmt.Printf("[%d/%d]", idx+1, len(batten.Checks))
	//fmt.Printf("%s", checkdefinition.Name())
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
		checkid:	idx+1,
		result:		resultString,
		name:		checkdefinition.Name(),
	}
}

func ShowResults(results *JSONResult) {
	fmt.Printf("%s\n","{'checks:'")
	fmt.Printf("%s","{")
	for _, result := range results.CheckResults {
		fmt.Printf("[%d:%s:%s]", 
			result.checkid,
			result.result,
			result.name)
	}
	fmt.Printf("}\n%s","}")
}
