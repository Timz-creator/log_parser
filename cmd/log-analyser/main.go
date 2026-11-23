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

type AnalysisReport struct {
	total_events int
	counts_by_XID_code map[int]int
	Severity map[string]int
	GPU_id map[string]int
}

var entries []LogEntry

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
	13: {
		Description: "Out of bounds error or illegal instruction or illegal register",
		Severity: "Warning",
		Rationale: "This code is a general user application faults",

	},
	32: {
		Description: "GPU memory fault",
		Severity: "Critical",
		Rationale: "Event logged when a fault is reported by MMU",
	},
	43: {
		Description: "GPU stopped processing",
		Severity: "Critical",
		Rationale: "User application hits a software induced fault and mus terminate",
	},

}
var occurs_code = make(map[int]int)

var occurs_severity = make(map[string]int)

var occurs_gpu = make(map[string]int)








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
		value := classifyXID(logentry.XID_code)
		entries = append(entries, logentry)
		occurs_code[xidCODE2] += 1
		occurs_gpu[logentry.GPU_id] += 1
		analysis := AnalysisReport{
			total_events:  ,
			counts_by_XID_code: occurs_code,
			GPU_id: occurs_gpu,
			Severity: value.Severity,
		}
		return logentry, nil


	}

	func classifyXID(code int) XIDClassification {
		fmt.Println(code)
		value, ok := XIDPolicies[code]
		if ok != true {
			fmt.Println("XID code doesn't exist")
		}
		fmt.Println(value.Severity)
		occurs_severity[value.Severity] += 1
		return value

	}

	func analysisReport()() {

	}

	



func main() {
	
	ReadFile()
}

