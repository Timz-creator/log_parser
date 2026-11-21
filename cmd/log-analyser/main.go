package main 
import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	
)

type LogEntry struct {
	timestamp string
	GPU_id string
	XID_code int
	message string
}

type XIDClassification struct {
	Description string
	Severity string
	Rationale string
}

//classifications -
	//Critical - needs immediate attention
	//Warning - should investigate issue in a timely manner
	//Informational - merely notifying operations 
	//Unknown - not a know XID code

var XIDPolicies = map[int]XIDClassification{
	79: {
		Description: "GPU driver attempts to connect to GPU unsuccessfully",
		Severity: "Critical",
		Rationale: "GPU connectivity, probably by failing GPU hardward",
	},
}







func ReadFile() () {
	fmt.Println("reading the XID logs, hold on bro")
	fileName := "example.txt"
	data, err := os.Open(fileName)
	
	if err != nil {
		panic(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	
	for scanner.Scan() {
		
		
		
		result, error := parseXID(scanner.Text())
		if error == nil {
			fmt.Println(result)
		}
		if error != nil {
			fmt.Println(error)
		}
		
	}
	

	}

	func parseXID(line string) (LogEntry, error) {
		rest, message, flag := strings.Cut(line, ",")
		if flag != true {
			return LogEntry{}, fmt.Errorf("invalid log")
		}
		rest2, xidCODE, flag2 := strings.CutLast(rest, ":")
		_ = rest2

		if flag2 != true {
			return LogEntry{}, fmt.Errorf("invalid log")
		}

		rest3, gpuID, flag3 := strings.CutLast(rest2, "Xid")
		message = strings.TrimSpace(message)
		gpuID= strings.TrimSpace(gpuID)
		gpuID = strings.Trim(gpuID, ")(")
		xidCODE = strings.TrimSpace(xidCODE)
		xidCODE2, err := strconv.Atoi(xidCODE)
		if err != nil {
			return LogEntry{}, fmt.Errorf("invalid log")
		}
		_ = rest3
		_ = err
		if flag3 != true {
			return LogEntry{}, fmt.Errorf("invalid log")
		}
		

		logentry := LogEntry{
			GPU_id: gpuID,
			XID_code: xidCODE2,
			message: message,
		}
		if err != nil {
			return LogEntry{}, fmt.Errorf("invalid log")
		}
		classifyXID(logentry.XID_code)

		return logentry, nil


	}

	func classifyXID(code int) XIDClassification {
		fmt.Println(code)
		value, ok := XIDPolicies[code]
		if ok != true {
			fmt.Println("XID code doesn't exist")
		}
		fmt.Println(value.Severity)
		return value

	}

	



func main() {
	
	ReadFile()
}

