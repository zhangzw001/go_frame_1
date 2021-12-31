package main

import (
	"bufio"
	"fmt"
	"go_frame/input"
	"os"
	"time"
)

func main() {
	os.Exit(mainWithExitCode())
}

func usage(programName string) {
	fmt.Printf("Usage: %s aunt_path order_path result_path elapsed_path\n", programName)
}

const ExitFailure = 1
const ExitSuccess = 0

func mainWithExitCode() (exitCode int) {
	var err error
	exitCode = ExitSuccess
	if len(os.Args) != 5 {
		usage(os.Args[0])
		exitCode = ExitFailure
		return
	}
	err = OpenOutputFile(os.Args[3])
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "fail when OpenOutputFile %s: %v\n", os.Args[3], err)
		exitCode = ExitFailure
		return
	}
	defer func() {
		err = CloseOutputFile()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fail when CloseOutputFile: %v\n", err)
			exitCode = ExitFailure
			return
		}
	}()

	// 执行主体
	var t1 = time.Now()
	Process(os.Args[1], os.Args[2])
	var t2 = time.Now()
	var elapsed = t2.Sub(t1)

	// 将执行 Process 的耗时输出到指定文件
	err = OutputTimeFile(os.Args[4], elapsed)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "fail when output time file: %v\n", err)
		exitCode = ExitFailure
		return
	}

	return
}

var File *os.File
var Writer *bufio.Writer

func OpenOutputFile(outputFilePath string) (err error) {
	File, err = os.OpenFile(outputFilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0664)
	if err != nil {
		return
	}
	Writer = bufio.NewWriter(File)
	return
}

// CloseOutputFile 只有在 OpenOutputFile 执行成功的情况下, 才能调用该函数
func CloseOutputFile() (err error) {
	err = Writer.Flush()
	if err != nil {
		err = fmt.Errorf("fail when Flush in CloseOutputFile: %w", err)
		return
	}
	// 虽然 Sync 很耗时, 但框架不会统计这里的耗时, 更准确的是框架不会统计 CloseOutputFile() 的耗时
	err = File.Sync()
	if err != nil {
		err = fmt.Errorf("fail when Sync in CloseOutputFile: %w", err)
		return
	}
	err = File.Close()
	if err != nil {
		err = fmt.Errorf("fail when Close in CloseOutputFile: %w", err)
		return
	}
	return
}

func AddSet(order, aunt int) {
	// 不用检查此处的 err. 如果这里发生 err, 框架在调用 CloseOutputFile 的时候会返回 err.
	_, _ = fmt.Fprintf(Writer, "%d,%d\n", order, aunt)
}

func AddSets(order2aunt map[int]int) {
	for order, aunt := range order2aunt {
		AddSet(order, aunt)
	}
}

// OutputTimeFile 将运行时长输出到指定文件
func OutputTimeFile(timeFilePath string, dur time.Duration) (err error) {
	timeFile, err := os.OpenFile(timeFilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0664)
	if err != nil {
		err = fmt.Errorf("fail when open time file %s: %w", timeFilePath, err)
		return
	}
	defer func() {
		errTmp := timeFile.Close()
		if errTmp != nil {
			err = fmt.Errorf("fail when close time file %s: %w", timeFilePath, err)
		}
	}()
	_, err = fmt.Fprintf(timeFile, "%d\n", dur.Milliseconds())
	if err != nil {
		err = fmt.Errorf("fail when write time file %s: %w", timeFilePath, err)
		return
	}
	return
}


func Process(auntPath, orderPath string) {
	// 计算答案
	// 对这个order进行分配aunt
	for len(input.OrderLists) > 0 {
		aid,oid := input.Worker(input.OrderLists,input.AuntLists)
		AddSet(oid, aid)
	}
	//...

	// 输出答案 (方式一) (多次调用会追加输出)
	//var orderID = 123
	//var auntID = 456
	//AddSet(orderID, auntID)

	// 输出答案 (方式二) (多次调用会追加输出)
	//var order2aunt = make(map[int]int)
	//AddSets(order2aunt)
}



