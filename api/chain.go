package api

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func (a *Api) chainQuery(c *gin.Context) {
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		RespErr(c, fmt.Errorf(fmt.Sprintf("read body err:(%v)", err)))
		return
	}
	res, err := a.Manager.ChainQuery(string(b))
	if err != nil {
		RespErr(c, fmt.Errorf("chainQuery err:(%v)", err))
		return
	}
	Resp(c, res)
}
func (a *Api) chainSetupChannel(c *gin.Context) {
	if err := a.Manager.ChainSetupChannel(); err != nil {
		RespErr(c, fmt.Errorf("ChainSetupChannel err:(%v)", err))
		return
	}
	Resp(c, "ok")
}
func (a *Api) installAndInstantiateCC(c *gin.Context) {
	if err := a.Manager.InstallAndInstantiateCC(); err != nil {
		RespErr(c, fmt.Errorf("installAndInstantiateCC err:(%v)", err))
		return
	}
	Resp(c, "ok")
}
