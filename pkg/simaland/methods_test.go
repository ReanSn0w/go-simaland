package simaland_test

import (
	"testing"

	"github.com/ReanSn0w/go-simaland/pkg/simaland"
	"github.com/go-pkgz/lgr"
)

var (
	cl, _ = simaland.NewClient(
		simaland.WithLogger(lgr.New(lgr.Debug)),
	)
)

func TestClient_ListAuthors(t *testing.T) {
	_, err := cl.ListAuthors(1, 50)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestClient_Author(t *testing.T) {
	author, err := cl.GetAuthor(116)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(author.Name)
}

func TestClient_ListDeliveryAddress(t *testing.T) {
	_, err := cl.ListDeliveryAddress()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestClient_ListCategory(t *testing.T) {
	list, err := cl.ListCategory(1, 10)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(list)
}

func TestClient_ListItems(t *testing.T) {
	catList, err := cl.ListCategory(1, 10)
	if err != nil {
		t.Error(err)
		return
	}

	list, err := cl.ListItems(1, 10, int(catList.Items[len(catList.Items)-1].ID))
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(list)
}

func TestClient_GetItem(t *testing.T) {
	item, err := cl.GetItem(6600428)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(item)
}
