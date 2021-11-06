package main

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/wqrqwerqrw/cust_gam_backend/store"
	"github.com/wqrqwerqrw/cust_gam_backend/utils"
)

func TestCreateUser(t *testing.T) {
	store.Init()

	for i := 0; i < 10; i++ {
		_, err := store.CreateUser(utils.RandStr(10), utils.RandStr(10), utils.RandStr(10), 1)
		if err != nil {
			t.Fail()
		}
	}
	runtime.Gosched()
}

func TestQueryUser(t *testing.T) {
	store.Init()

	s, err := store.QueryUser(7)
	fmt.Print(s)
	if err != nil {
		t.Fail()
	}
	runtime.Gosched()
}

func TestUpdateUser(t *testing.T) {

	store.Init()

	for i := 0; i < 10; i++ {
		s, err := store.UpdateUser(7, utils.RandStr(10), utils.RandStr(10), utils.RandStr(10), 1)
		fmt.Print(s)
		if err != nil {
			t.Fail()
		}
	}
	runtime.Gosched()
}
