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
	name	string
}


func FormatResultsForMachine(idx int, results *batten.CheckResults) CheckResult {
	
	checkdefinition := results.CheckDefinition

	//fmt.Printf("[%d/%d]", idx+1, len(batten.Checks))
	//fmt.Printf("%s", checkdefinition.Name())
	
	return CheckResult{
		checkid:	idx+1,
		name:		checkdefinition.Name(),
	}
}

func ShowResults(results *JSONResult) {
	for _, result := range results.CheckResults {
		fmt.Printf("[%d:%s]", result.checkid, result.name)
	}
	//fmt.Printf("%s",results)
}
