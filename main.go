package main
import ("fmt")
import("os")
import ( "os/exec")
import ("bufio")
import ("log")
import ("strings")
const shell ="/bin/bash"
var val int
var number = 0
var valu = 0
var path = "/packages/list.d/sour.list"
var curl1 = "/usr/bin/curl -s -o /packages/index.text "
var curlf string
var index = "/packages/index.text"
var packagen string
var lol = 0
var els string
var nub int
var els25 string
var els35 string
var output string
func varible() {
	val := 0
	fmt.Println(val)
	fmt.Println("Enter The package you want to install: ")
	fmt.Scanln(&packagen)
	fmt.Println(packagen)
}
func mkdir() {
    out, err := exec.Command("mkdir", "/packages").Output()
    if err != nil {
        fmt.Printf("%s", err)
    }
    output := string(out[:])
    fmt.Println(output)
}
func file1() {
	readFile, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string
	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}
	readFile.Close()
	for _, eachline := range fileTextLines {
		fmt.Println(eachline)
		number++
		curlf = curl1+eachline
		fmt.Println(curlf)
		fmt.Println(number)
		out, err := exec.Command(shell, "-c", curlf).Output()
		if err != nil {
			fmt.Printf("%s", err)
		}
		fmt.Println(out)
		curlf = curl1+eachline
		fmt.Println(curlf)
		readFile2, err := os.Open(index)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	fileScanner2 := bufio.NewScanner(readFile2)
	fileScanner2.Split(bufio.ScanLines)
	var fileTextLines2 []string
	for fileScanner2.Scan() {
		fileTextLines2 = append(fileTextLines2, fileScanner2.Text())
	}
	readFile2.Close()
	for _, eachline := range fileTextLines2 {
		//valu = 0
		fmt.Println(eachline)
		els := strings.Split(eachline, " ")
		els2, els3 := els[0], els[1]
		fmt.Println(els2,"els3:",els3)
		if packagen == els2 { //&& nub == number{
			fmt.Println("this package will be installed.")
			els25 = els2
			els35 = els3
			return
		}
}
func install() {
	path := "/packages/"+els25+".gz"
	fmt.Println(path)
	out, err := exec.Command("sudo", "curl", "-o", path, els35).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Command Successfully Executed")
    	output := string(out[:])
    	fmt.Println(output)
}
func install2() {
	path := "/packages/"+els25+".gz"
	fmt.Println(path)
	out, err := exec.Command("sudo", "gzip", "-d", path).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Command Successfully Executed")
    	output := string(out[:])
    	fmt.Println(output)
}
func install3() {
		path := "/packages/"+els25
	fmt.Println(path)
	out, err := exec.Command("sudo", "chmod", "a+x", path).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Command Successfully Executed")
    	output := string(out[:])
    	fmt.Println(output)
}
func install4() {
	path := "/packages/"+els25
	path9 := "/bin/"+els25
	fmt.Println(path)
	out, err := exec.Command("sudo", "mv", path, path9).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Command Successfully Executed")
    	output := string(out[:])
    	fmt.Println(output)	
}
func flil() {
	write := els25+"\n"
	file, err := os.OpenFile("/packages/installed.list", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	 _, err2 := file.WriteString(write)
	 if err2 != nil {
	 	fmt.Println("Could not write text to example.txt")
	 } else {
	 	fmt.Println("Operation successful! Text has been appended to example.txt")
	 }
}
func main() {
varible()
mkdir()
file1()
install()
install2()
install3()
install4()
flil()
}
