package integ

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("start running integ test")
	closeFn := startWorkflowWorker()
	code := m.Run()
	closeFn()
	fmt.Println("finished running integ test with status code", code)
	os.Exit(code)
}

func apiV1WorkflowStateStart(c *gin.Context) {
	var req iwfidl.WorkflowStateStartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := workerService.HandleWorkflowStateStart(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
	return
}
func apiV1WorkflowStateDecide(c *gin.Context) {
	var req iwfidl.WorkflowStateDecideRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := workerService.HandleWorkflowStateDecide(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
	return
}

func startWorkflowWorker() (closeFunc func()) {
	router := gin.Default()
	router.POST(iwf.WorkflowStateStartApi, apiV1WorkflowStateStart)
	router.POST(iwf.WorkflowStateDecideApi, apiV1WorkflowStateDecide)

	wfServer := &http.Server{
		Addr:    ":" + iwf.DefaultWorkerPort,
		Handler: router,
	}
	go func() {
		if err := wfServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	return func() { wfServer.Close() }
}
