package demo

import (
	"bufio"
	"fmt"
	"io"

	//"math"
	"os"
	"strings"
	"sync"
)

func Demo() {
	/*
		Sample file data
		{"remark": "call time: 2021 / 04 / 15 13:52:07 customer Tel: 13913xx39xx", "no": "600020510132021101310210547639", "title": "b-ae0e-0242ac100907", "call_in_date": "2021-04-15 13:52:12", "name": "Zhang San", "_date": "2021-06-15", "name": "Zhang San", "meet": "1"}
		1. We take out the call_ in_ Date ":" 2021-04-15 13:52:1 data is written to another file
	*/
	var (
		// S            time.Time // current time
		file         *os.File
		fileStat     os.FileInfo
		err          error
		lastLineSize int64
	)
	// s = time.Now()
	if file, err = os.Open("/Users/zhangsan/Downloads/log.txt"); err != nil {
		fmt.Println(err)
	}
	defer func() {
		err = file.Close() //close after checking err
	}()
	//queryStartTime, err := time.Parse("2006-01-02T15:04:05.0000Z", startTimeArg)
	//if err != nil {
	//	fmt.Println("Could not able to parse the start time", startTimeArg)
	//	return
	//}
	//
	//queryFinishTime, err := time.Parse("2006-01-02T15:04:05.0000Z", finishTimeArg)
	//if err != nil {
	//	fmt.Println("Could not able to parse the finish time", finishTimeArg)
	//	return
	//}
	/**
	* {name:"log.log", size:911100961, mode:0x1a4,
	modTime:time.Time{wall:0x656c25c, ext:63742660691,
	loc:(*time.Location)(0x1192c80)}, sys:syscall.Stat_t{Dev:16777220,
	Mode:0x81a4, Nlink:0x1, Ino:0x118cba7, Uid:0x1f5, Gid:0x14, Rdev:0,
	Pad_cgo_0:[4]uint8{0x0, 0x0, 0x0, 0x0}, Atimespec:syscall.Timespec{Sec:1607063899, Nsec:977970393},
	Mtimespec:syscall.Timespec{Sec:1607063891, Nsec:106349148}, Ctimespec:syscall.Timespec{Sec:1607063891,
	Nsec:258847043}, Birthtimespec:syscall.Timespec{Sec:1607063883, Nsec:425808150},
	Size:911100961, Blocks:1784104, Blksize:4096, Flags:0x0, Gen:0x0, Lspare:0, Qspare:[2]int64{0, 0}}
	*
	*/
	if fileStat, err = file.Stat(); err != nil {
		return
	}
	fileSize := fileStat.Size() //72849354767
	offset := fileSize - 1
	//Detect if all lines are empty only \ n
	for {
		var (
			b    []byte
			n    int
			char string
		)
		b = make([]byte, 1)
		//Read from the specified location
		if n, err = file.ReadAt(b, offset); err != nil {
			fmt.Println("Error reading file ", err)
			break
		}
		char = string(b[0])
		if char == "\n" {
			break
		}
		offset--
		//Gets the size of a row
		lastLineSize += int64(n)
	}
	var (
		lastLine  []byte
		logSlice  []string
		logSlice1 []string
	)
	//Initialize a row of space
	lastLine = make([]byte, lastLineSize)
	_, err = file.ReadAt(lastLine, offset)
	if err != nil {
		fmt.Println("Could not able to read last line with offset", offset, "and lastline size", lastLineSize)
		return
	}
	//Distinguish according to conditions
	logSlice = strings.Split(strings.Trim(string(lastLine), "\n"), "next_pay_date")
	logSlice1 = strings.Split(logSlice[1], "\"")
	if logSlice1[2] == "2021-06-15" {
		Process(file)
	}
	// fmt.Println("\nTime taken - ", time.Since(s))
	fmt.Println(err)
}
func Process(f *os.File) error {
	//Read the key of the data to reduce the GC pressure
	linesPool := sync.Pool{New: func() interface{} {
		lines := make([]byte, 250*1024)
		return lines
	}}
	//Read back data pool
	stringPool := sync.Pool{New: func() interface{} {
		lines := ""
		return lines
	}}
	//A file object itself implements io. Reader. Use buffio. Newreader to initialize a reader object. If it exists in the buffer, it will be emptied after reading once
	r := bufio.NewReader(f) //
	//Set the read buffer pool size to 16 by default
	r = bufio.NewReaderSize(r, 250*1024)
	var wg sync.WaitGroup
	for {
		buf := linesPool.Get().([]byte)
		//Read the contents of the reader object into a buf of type [] byte
		n, err := r.Read(buf)
		buf = buf[:n]
		if n == 0 {
			if err != nil {
				fmt.Println(err)
				break
			}
			if err == io.EOF {
				break
			}
			return err
		}
		//Make up the remaining unsatisfied surplus
		nextUntillNewline, err := r.ReadBytes('\n')
		//fmt.Println(string(nextUntillNewline))
		if err != io.EOF {
			buf = append(buf, nextUntillNewline...)
		}
		wg.Add(1)
		go func() {
			ProcessChunk(buf, &linesPool, &stringPool)
			wg.Done()
		}()
	}
	wg.Wait()
	return nil
}

func ProcessChunk(chunk []byte, linesPool *sync.Pool, stringPool *sync.Pool) {
	//Handle accordingly
}
