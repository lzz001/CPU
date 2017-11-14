
package main

import (  
   
    "fmt"  
    //"io/ioutil"  
    "net/http" 
    "strings" 
    //"net/url"
   "os"
   "io"
)  
  
func main() {
	//insert()
 // search()
  fmt.Println("lzz and qiqi")
  fmt.Println("I Love My Family!")
  fmt.Println("I Love coco!")
}




//插入数据
func insert(){
	 //生成client 参数为默认
	client := &http.Client{}
	
	//生成要访问的url
	url := "http://127.0.0.1:9200/v1/push1"
	//{\"timestamp\": 1499935041,\"metricName\": \"cpu.idle\",\"value\": {\"cnt\': 1,\"sum\": 99.0,\"max\": 99.0,\"min\": 99.0}}
	// {\"name\":\"wwww\",\"age\":\"38\"}
	//提交请求
	reqest, err := http.NewRequest("POST", url, strings.NewReader("{\"timestamp\": 1499935777,\"metricName\": \"cpu.idle2\",\"value\": {\"cnt\": 1,\"sum\": 99.0,\"max\": 99.0,\"min\": 99.0}}"))
	reqest.Header.Add("Content-type", "application/json");
	if err != nil {
		panic(err)
	}
	
	//处理返回结果
	response, _ := client.Do(reqest)
   
   //将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
	stdout := os.Stdout
	_, err = io.Copy(stdout, response.Body)
   
   //返回的状态码
	//status := response.StatusCode

	fmt.Println(response.Body)
}

func search(){
	//生成client 参数为默认
	client := &http.Client{}
	
	//生成要访问的url
	url := "http://localhost:9200/v1/push1/_search"
	//{\"timestamp\": 1499935041,\"metricName\": \"cpu.idle\",\"value\": {\"cnt\': 1,\"sum\": 99.0,\"max\": 99.0,\"min\": 99.0}}
	// {\"name\":\"wwww\",\"age\":\"38\"}
	//{\"query\":{\"match\":{\"metricName\":\"cpu.idle\"}}}
	//{\"query\":{\"bool\":{\"must\":{\"match\":{\"metricName\":\"cpu.idle2\"}},\"filter\":{\"range\":{\"timestamp\":{\"gt\":1499935041,\"lt\":1499935031}}}}}}
	// "@timestamp": {"gte": "2016-05-10T00:00:00.000Z", "lte": "2016-05-11T00:00:00.000Z"}
	//提交请求
	reqest, err := http.NewRequest("POST", url, strings.NewReader("{\"query\":{\"bool\":{\"must\":{\"match\":{\"metricName\":\"cpu.idle2\"}},\"filter\":{\"range\":{\"timestamp\":{\"gt\":1499935031, \"lt\":1499935041}}}}}}"))
	reqest.Header.Add("Content-type", "application/json");
	if err != nil {
		panic(err)
	}
	
	//处理返回结果
	response, _ := client.Do(reqest)
   
   //将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
	stdout := os.Stdout
	_, err = io.Copy(stdout, response.Body)
   
   //返回的状态码
	//status := response.StatusCode

	fmt.Println(response.Body)
}