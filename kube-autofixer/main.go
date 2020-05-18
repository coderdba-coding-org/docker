package main

import (
	//"crypto/tls"
	//"encoding/base64"
	//"encoding/json"
	//"flag"
	"fmt"
	//"io/ioutil"
	"log"
	//"net"
	//"net/http"
	//"os"
	"os/exec"
	//"path"
	"bytes"
	//"strings"
        "time"
)

func main() {

	//var (
		//buf    bytes.Buffer
		//logger = log.New(&buf, "logger: ", log.Lshortfile)
	//)

	// Maintenance functions
	//reinitPatroni()

	// Trial/test commands
	//listPods("mypostgresql", "default")
	//runTestShellCommand1()
	//runTestShellCommand2()

	//runTestKubectl1()

	for {
		//fmt.Println("Infinite Loop")
                //log.Fatal("Infinite Loop")

                //logger.Print("Infinite Loop")
                log.Println("Infinite Loop")
		time.Sleep(100 * time.Second)
	}

}

func patroniReinit() {

	stsName := "mypostgresql"

	cmd := fmt.Sprintf("kubectl exec -it $(kubectl get pods --no-headers -l %s-role=master | tail -n 1 | awk '{print $1}') -- bash -c \"cd /home/postgres && patronictl -c postgres.yml list\"", stsName)

	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Command is %s\n", cmd)
	log.Print("Command is %s\n", cmd)
	fmt.Printf("Output is %s\n", out)
	log.Print("Output is %s\n", out)

}

// TEST/TRIAL FUNCTIONS
func runTestKubectl1() {

	cmd := fmt.Sprintf("kubectl get nodes")

	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Command is %s\n", cmd)
	log.Print("Command is %s\n", cmd)
	fmt.Printf("Output is %s\n", out)
	log.Print("Output is %s\n", out)
}


func runTestShellCommand1() {
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}

func runTestShellCommand2() {
	shellcmd := "ls"
	cmd := exec.Command("bash", "-c", shellcmd)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Command output: %q\n", out.String())

}

func listPods(podNamePattern string, namespace string) (err error) {

	cmd := fmt.Sprintf("kubectl get pods -n %s -o wide | grep %s", namespace, podNamePattern)
	fmt.Println(cmd)

	if logger.LogLevel1() {
		log.Print(cmd)
	}
	if err := exec.Command("bash", "-c", cmd).Run(); err != nil {
		fmt.Println("Error running command")
		return err
	}
	return nil
}
