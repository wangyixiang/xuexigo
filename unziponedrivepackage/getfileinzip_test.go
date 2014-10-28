package unziponedrivepackage

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"testing"
)

func Test_getIndexMap(t *testing.T) {
	correctMap := map[string]string{
		"File1.pdf":  "多元样条函数及其应用.pdf",
		"File2.pdf":  "大哉言数.pdf",
		"File3.pdf":  "奥数教程 初一年级  第一版.pdf",
		"File4.pdf":  "奥数教程 初三年级  第一版.pdf",
		"File5.pdf":  "奥数教程 初二年级  第一版.pdf",
		"File6.pdf":  "奥数教程 高一年级  （第3版）.pdf",
		"File7.pdf":  "奥数教程 高三年级  （第3版）.pdf",
		"File8.pdf":  "奥数教程 高二年级  （第3版）.pdf",
		"File9.pdf":  "奥林匹克数竞赛解迷（高中部分）（康纪权）.pdf",
		"File10.pdf": "好玩的数学.pdf",
		"File11.pdf": "孤子理论（逆问题方法）.pdf",
		"File12.pdf": "孤子理论和微扰方法.pdf",
		"File13.pdf": "实变函数与泛函分析概要（第二版）上册.pdf",
		"File14.pdf": "实变函数与泛函分析概要（第二版）下册.pdf",
		"File15.pdf": "实变函数论与泛函分析上册第二版（夏道行+吴卓人+严绍宗+舒五昌）.pdf",
		"File16.pdf": "实变函数论与泛函分析下册第二版（夏道行+吴卓人+严绍宗+舒五昌）.pdf",
		"File17.pdf": "实变函数论的典型问题与方法(张喜堂).pdf",
		"File18.pdf": "实变函数论讲义(王昆阳）.pdf",
		"File19.pdf": "对应_王子侠单墫.pdf",
		"File20.exe": "小学生1年级数学奥数.exe",
		"File21.exe": "小学生2年级数学奥数.exe",
		"File22.exe": "小学生3年级数学奥数.exe",
		"File23.exe": "小学生4年级数学奥数.exe",
		"File24.exe": "小学生5年级数学奥数.exe",
		"File25.exe": "小学生6年级数学奥数.exe",
		"File26.pdf": "小波分析与信号处理.pdf",
		"File27.pdf": "小波分析导论.pdf",
		"File28.pdf": "工程控制论.pdf",
		"File29.pdf": "工程控制论上、下（钱学森，宋健）.pdf"}
	testDataFile := "testdata" + string(os.PathSeparator) + "Encoding Errors.txt"

	f, _ := os.Open(testDataFile)
	if im, ok := getIndexMap(f); ok != nil {
		t.Log("running getIndexMap failed!!!")
		t.Fail()
	} else {
		if len(correctMap) != len(im) {
			t.Log("The index Map is not same length as the verification map")
			t.Fail()
		} else {
			for k, v := range im {
				if correctMap[k] != v {
					t.Log("The value in one pair is not same as the one in the verifcation map")
					t.Fail()
				}
			}
		}
	}
}

func TestUnzipPackageTo(t *testing.T) {
	cmds, err := checkGoExeAndGetCMDList()
	if err != nil {
		t.Log("go running environment is not ready, please check!!!")
		t.Fail()
	}
	if err = setupTempZipFile(cmds); err != nil {
		t.Log("set up temprary files failed.")
		t.Fail()
	}
	testingpackage := "testdata" + string(os.PathSeparator) + "packagefortesting.zip"
	unpackageDir := "unpackagefortesting"
	if exists(unpackageDir) {
		if err = os.RemoveAll(unpackageDir); err != nil {
			t.Log("cleaning environment failed.")
			t.Fail()
		}
	}
	if err = UnzipPackageTo(testingpackage, "unpackagefortesting", true); err != nil {
		t.Log("UnzipPackageTo runing failed.")
		t.Fail()
	}
	files, err := ioutil.ReadDir(unpackageDir)
	if err != nil {
		t.Log("reading unpackageDir failed.")
		t.Fail()
	}
	if len(cmds) != len(files) {
		t.Log("the files in the package are not same as input data.")
		t.Fail()
	}
	for _, cmd := range cmds {
		stdout := &bytes.Buffer{}
		cmdcmd := exec.Command("go", "help", cmd)
		cmdcmd.Stdout = stdout
		cmdfile, err := os.Open(path.Join(unpackageDir, "go_"+cmd))
		defer cmdfile.Close()
		if err != nil {
			t.Log("opening the file in unpackageDir failed.")
			t.Fail()
		}
		cmdfilebuf, err := ioutil.ReadAll(cmdfile)
		if err = cmdcmd.Run(); err != nil {
			t.Log("It's so strange, the 'go help " + cmd + "' running failed.")
			t.Fail()
		}
		if same := bytes.Compare(stdout.Bytes(), cmdfilebuf); same != 0 {
			t.Log("failed at 'go help " + cmd + "' .")
			t.Fail()
		}

	}
}

func checkZipExe() error {
	var waitStatus syscall.WaitStatus
	cmd := exec.Command("zip", "--version")
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			log.Println("zip command launching failed with exitcode=", waitStatus.ExitStatus())
		}
		return err
	}
	return nil
}

func checkGoExeAndGetCMDList() (cmds []string, err error) {
	var waitStatus syscall.WaitStatus
	stdout := &bytes.Buffer{}
	cmd := exec.Command("go", "help")
	cmd.Stdout = stdout
	if err = cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			log.Println("go command launching failed with exitcode=", waitStatus.ExitStatus())
		}
		return nil, err
	}
	stdCmdsRE := regexp.MustCompile(`(?s)The commands are:(.*)Use "go help \[command\]"`)
	addCmdsRE := regexp.MustCompile(`(?s)Additional help topics:(.*)Use "go help \[topic\]"`)
	stdCmdsS := stdCmdsRE.FindSubmatch(stdout.Bytes())
	addCmdsS := addCmdsRE.FindSubmatch(stdout.Bytes())

	if len(stdCmdsS) != 2 || len(addCmdsS) != 2 {
		return nil, errors.New("Can't analyze go help output.")
	}
	stdCmds, addCmds := strings.TrimSpace(string(stdCmdsS[1])), strings.TrimSpace(string(addCmdsS[1]))
	Cmds := stdCmds + "\n" + addCmds
	CmdsInputReader := bufio.NewReader(strings.NewReader(Cmds))
	for {
		line, err := CmdsInputReader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" && err == io.EOF {
			break
		}
		linefields := strings.Split(line, " ")
		if len(linefields) >= 2 {
			cmds = append(cmds, strings.TrimSpace(linefields[0]))
		}
		if err == io.EOF {
			break
		}
	}
	return cmds, nil
}

func setupTempZipFile(cmds []string) (err error) {
	testingdir := "packagefortesting"
	sStarter := "Original File Name  ->  New File Name"
	sMarker := "->"
	sEncodingErrorsFile := "Encoding Errors.txt"
	err = nil
	buf := &bytes.Buffer{}
	if exists(testingdir) {
		if err = os.RemoveAll(testingdir); err != nil {
			return err
		}
	}
	if err = os.Mkdir(testingdir, 0770); err != nil {
		return err
	}
	buf.WriteString(`wangyixiang made it for testing, \\(^.^)//`)
	buf.WriteString("\n")
	buf.WriteString(sStarter)
	buf.WriteString("\n")

	for i, cmd := range cmds {
		stdout := &bytes.Buffer{}
		cmdcmd := exec.Command("go", "help", cmd)
		cmdcmd.Stdout = stdout
		cmdcmd.Run()
		filename1 := testingdir + string(os.PathSeparator) + "go_" + cmd
		filename2 := testingdir + string(os.PathSeparator) + "File" + strconv.FormatInt(int64(i), 10)
		file, err := os.Create(filename1)
		if err != nil {
			log.Println(err)
			return err
		}
		if _, err = file.Write(stdout.Bytes()); err != nil {
			file.Close()
			log.Println(err)
			return err
		}
		file.Close()
		err = os.Rename(filename1, filename2)
		if err != nil {
			log.Println(err)
			return err
		}
		buf.WriteString("go_" + cmd + " " + sMarker + " " + "File" + strconv.FormatInt(int64(i), 10))
		buf.WriteString("\n")
	}

	eefile, err := os.Create(testingdir + string(os.PathSeparator) + sEncodingErrorsFile)
	if err != nil {
		return err
	}
	defer eefile.Close()
	_, err = eefile.Write(buf.Bytes())
	if err != nil {
		return err
	}

	zipcmd := exec.Command("zip", "-j", "-r", "testdata"+string(os.PathSeparator)+testingdir, testingdir)
	err = zipcmd.Run()
	return err
}

func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
