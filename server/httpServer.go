package main

import (
	"SimpleHttpServer/conf"
	"SimpleHttpServer/dao"
	"SimpleHttpServer/data"
	util "SimpleHttpServer/util"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

var env string

func init() {
	//Get environment
	env = os.Getenv("GO_APP_ENV")
}
func main() {

	_, filePath, _, _ := runtime.Caller(0)
	p := path.Dir(filePath)
	p = path.Dir(p)

	addr := conf.Addr + ":" + conf.Port
	if env == util.EnvPro {
		httpRun(addr)
	} else {
		log.Println("please check your os variables, let GO_APP_ENV='PRO' and restart server")
	}

}

func httpRun(addr string) {
	router := mux.NewRouter()
	router.HandleFunc("/v{version}/{action}", calAPIHandler).Methods("GET", "POST")
	router.HandleFunc("/{namespace}/{resource}/{action}", listHandler).Methods("GET")
	router.HandleFunc("/user/self/{option}", authenHandler).Methods("POST")
	http.ListenAndServe(addr, router)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		vars := mux.Vars(r)
		namespace := vars["namespace"]
		resource := vars["resource"]
		action := vars["action"]
		log.Println(namespace, resource, action)

		if namespace == util.NamespaceTutorial && resource == util.ResourceStudent {
			if action == util.ActionList { //list action
				t := template.Must(template.ParseFiles(conf.TemplatesPath + "studentList.html"))
				// tuser := &data.TUser{FirstName: "yo", LastName: "yoyo1", Age: 12}
				// studentList := make([]*data.TUser, 0)
				// studentList = append(studentList, tuser)
				studentList, err := dao.ListSimpleUser()
				if err != nil {
					log.Fatalf("list student failed:%v", err)
				}
				w.Header().Set("content-type", "text/html")
				t.Execute(w, studentList)
				return
			}
		}
	}
}
func calAPIHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		versionParam, actionParam := filterVersionAndAction(r)
		//-------------for Test API--------------
		if actionParam == util.TestAPI {
			testReturn := &data.TestRes{Ret: util.ResponseSucess}
			returnData, err := json.Marshal(testReturn)
			if err != nil {
				log.Fatalf("json marshal failed: %v", err)
			}
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(returnData)
			return
		}
		//----------for Normal API---------------

		u, err := url.Parse(r.URL.String())
		if err != nil {
			log.Fatalf("url parse failed: %v", err)
			return
		}
		fmt.Println(u.RawQuery)
		m, _ := url.ParseQuery(u.RawQuery)
		if len(m) < 2 {
			log.Fatal("query less than 2")
			testData := &data.TestRes{
				Ret: util.ResponseFailed,
			}
			returnData, err := json.Marshal(testData)
			if err != nil {
				panic(err)
			}
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(returnData)
			return
		}
		//op1
		a, aerr := strconv.Atoi(m["a"][0])
		if aerr != nil {
			log.Fatalf("a format: %v", aerr)
		}
		//op2
		b, berr := strconv.Atoi(m["b"][0])
		if berr != nil {
			log.Fatalf("b format: %v", berr)
		}

		var result int
		if actionParam == util.ActionPlus {
			result = a + b
		} else { //action is not "plus"
			result = -9999999
		}
		testData := &data.TestRes{
			Ret:     util.ResponseSucess,
			Version: versionParam,
			Action:  actionParam,
			Result:  result,
		}
		returnData, err := json.Marshal(testData)
		if err != nil {
			panic(err)
		}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(returnData)
		return
	}
	if r.Method == http.MethodPost {
		versionParam, actionParam := filterVersionAndAction(r)
		a, aerr := strconv.Atoi(r.FormValue("a"))
		if aerr != nil {
			log.Fatalf("a format: %v", aerr)
		}
		b, berr := strconv.Atoi(r.FormValue("b"))
		if berr != nil {
			log.Fatalf("b format: %v", berr)
		}
		var result int
		if actionParam == util.ActionPlus {
			result = a + b
		} else { //action is illegal
			log.Println("action not correct")
			return
		}
		testData := &data.TestRes{
			Ret:     util.ResponseSucess,
			Version: versionParam,
			Action:  actionParam,
			Result:  result,
		}
		returnData, err := json.Marshal(testData)
		if err != nil {
			panic(err)
		}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(returnData)
	}
}

//Authenticate Handler
func authenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		token := r.FormValue("intAuthToken")
		fmt.Println(token)
		if len(token) < 1 { //no token passed
			testData := &data.AuthenRes{Ret: util.ResponseSucess}
			returnData, err := json.Marshal(testData)
			if err != nil {
				log.Fatalf("json marshal authen failed:%v", err)
				// panic(err)
			}
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(returnData)
			return

		}
		vars := mux.Vars(r)
		option := vars["option"]
		//Get Name
		if option == util.AuthenName {
			authenRes, err := dao.CheckSelf(token)
			if err != nil {
				// w.Header().Set("content-type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				// w.Write(returnData)
				return
			}
			returnData, err := json.Marshal(authenRes)
			if err != nil {
				log.Fatalf("json marshal authen failed:%v", err)
			}
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(returnData)
			return
		} else if option == util.AuthenAge { //Get Age

		}

	}
}

//This is for filtering version and action field
func filterVersionAndAction(r *http.Request) (int, string) {
	vars := mux.Vars(r)
	version := vars["version"]
	versionParam, verr := strconv.Atoi(version)
	checkVersion(verr)

	actionParam := strings.ToLower(vars["action"])

	return versionParam, actionParam
}
func checkVersion(err error) {
	if err != nil {
		log.Fatalf("Illegal version: %v", err)
	}
}
