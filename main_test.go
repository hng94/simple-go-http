package main

import (
	"github.com/hoangnguyen94/simplesurance-coding/libs"
	"github.com/hoangnguyen94/simplesurance-coding/structs"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func Test_1000requests(t *testing.T) {

	testData := structs.Timestamps{}
	testData.Init("mockData.log")
	for i := 0; i < 1000; i++ {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		testData.Handler(w, req)
		res := w.Result()
		res.Body.Close()
		data, error := ioutil.ReadAll(res.Body)
		if error != nil {
			t.Errorf("expected error to be nil got %v", error)
		}
		strRes := string(data)
		actual, _ := strconv.Atoi(strRes)
		expected := len(testData.Data)
		if actual != expected {
			t.Errorf("expected %v got %v", expected, actual)
		}
	}
}

func Test_2requests_wait30s(t *testing.T) {
	testData := structs.Timestamps{}
	testData.Init("mockData.log")
	req1 := httptest.NewRequest(http.MethodGet, "/", nil)
	w1 := httptest.NewRecorder()
	testData.Handler(w1, req1)

	time.Sleep(30 * time.Second)
	req2 := httptest.NewRequest(http.MethodGet, "/", nil)
	w2 := httptest.NewRecorder()
	testData.Handler(w2, req2)
	res := w2.Result()
	res.Body.Close()
	data, error := ioutil.ReadAll(res.Body)
	if error != nil {
		t.Errorf("expected error to be nil got %v", error)
	}

	strRes := string(data)
	actual, _ := strconv.Atoi(strRes)
	expected := len(testData.Data)
	if actual != expected {
		t.Errorf("expected %v got %v", expected, actual)
	} else {
		t.Logf("expected %v got %v", expected, actual)
	}

	libs.RemoveFile(testData.Filename)
}

func Test_2requests_wait60s(t *testing.T) {
	testData := structs.Timestamps{}
	testData.Init("mockData.log")
	req1 := httptest.NewRequest(http.MethodGet, "/", nil)
	w1 := httptest.NewRecorder()
	testData.Handler(w1, req1)

	time.Sleep(30 * time.Second)
	req2 := httptest.NewRequest(http.MethodGet, "/", nil)
	w2 := httptest.NewRecorder()
	testData.Handler(w2, req2)
	res := w2.Result()
	res.Body.Close()
	data, error := ioutil.ReadAll(res.Body)
	if error != nil {
		t.Errorf("expected error to be nil got %v", error)
	}

	strRes := string(data)
	actual, _ := strconv.Atoi(strRes)
	expected := len(testData.Data)
	if actual != expected {
		t.Errorf("expected %v got %v", expected, actual)
	} else {
		t.Logf("expected %v got %v", expected, actual)
	}

	libs.RemoveFile(testData.Filename)
}
