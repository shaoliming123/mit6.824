package mr

import (
	"log"
	"path/filepath"
)
import "net"
import "os"
import "net/rpc"
import "net/http"



type Master struct {
	// Your definitions here.
	inputFiles []string

	state State
}

// Your code here -- RPC handlers for the worker to call.

//
// an example RPC handler.
//
// the RPC argument and reply types are defined in rpc.go.
//
func (m *Master) Example(args *ExampleArgs, reply *ExampleReply) error {
	reply.Y = args.X + 1
	return nil
}


//
// start a thread that listens for RPCs from worker.go
//
func (m *Master) server() {
	rpc.Register(m)
	rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	sockname := masterSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

//
// main/mrmaster.go calls Done() periodically to find out
// if the entire job has finished.
//
func (m *Master) Done() bool {
	ret :=()


	return ret
}

func (m *Master) SetUp(files []string, nReduce int) error{
	// 获得输入文件
	inputFiles,err := getAllFile(files)
	if err !=nil{
		return err
	}
	m.inputFiles = inputFiles
	return nil
}

//
// create a Master.
// main/mrmaster.go calls this function.
// nReduce is the number of reduce tasks to use.
//

func MakeMaster(files []string, nReduce int) *Master {
	m := Master{}

	// Your code here.
	if err := m.SetUp(files,nReduce); err!=nil{
		panic(err)
	}

	m.server()
	return &m
}

func getAllFile(files []string)([]string, error){
	var allFiles []string
	for _,file := range files{
		glob, err := filepath.Glob(file)
		if err !=nil{
			return nil,err
		}
		if glob != nil{
			allFiles = append(allFiles, glob...)
		}
	}
	return allFiles,nil
}