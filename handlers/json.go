package handlers

import(
	"net/http"
	"os"
	//"encoding/json"
	"io/ioutil"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/zeropage/mukgoorm/setting"
)

func Json(c *gin.Context){
	dir := setting.GetDirectory().Path

	file , err1 := os.Stat(dir)
	if err1 != nil {
		log.Error(err1)
		c.HTML(http.StatusNotFound, "error/404.tmpl", gin.H{})
	}

	ret, err2 := MakeJsTree(file, dir)
	fmt.Print("\n\n")
	if err2 != nil {
		log.Error(err1)
		c.HTML(http.StatusNotFound, "error/404.tmpl", gin.H{})
	}
	jstreeFormat := map[string]JsTreeContainer{
		"data": ret,
	}
	fmt.Print("Json Result : ")
	fmt.Print(jstreeFormat)
	fmt.Print("\n\n")

	c.JSON(http.StatusOK, gin.H{
		"core": jstreeFormat,
	})
}

func MakeJsTree(rootFile os.FileInfo, rootDir string)(JsTreeContainer, error){
	var subtree JsTreeContainer
	if rootFile.IsDir() {	// i think this wasn't work correctly
		fmt.Print("\n")
		fmt.Print("this is folder : ")
		fmt.Print(rootDir)
		files, err := ioutil.ReadDir(rootDir)
		if err != nil {
			panic(err)
		} else {
			var subByte []JsTreeContainer
			for i := 0; i < len(files); i++ {
				subdata, _ := MakeJsTree(files[i], rootDir + "/" + files[i].Name())
				subByte = append(subByte, subdata)
			}
			subtree = JsTreeContainer{rootFile.Name(), subByte}
		}
	} else {
		fmt.Print("\n")
		fmt.Print("this is file : ")
		fmt.Print(rootDir)
		subtree = JsTreeContainer{rootFile.Name(), []JsTreeContainer{}}
	}

	fmt.Print("\n")
	fmt.Print("now subtree : ")
	fmt.Print(subtree)

	return subtree, nil
}

type JsTreeContainer struct{
	text string
	Children []JsTreeContainer
}
