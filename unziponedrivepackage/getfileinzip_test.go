package unziponedrivepackage

import (
	"os"
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
