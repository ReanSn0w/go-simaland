package simaland_test

import (
	"testing"

	"github.com/ReanSn0w/go-simaland/pkg/simaland"
	"github.com/go-pkgz/lgr"
)

var (
	cl, _ = simaland.NewClient(
		simaland.WithLogger(lgr.New(lgr.Debug)),
		simaland.WithApiKey("4eddeaea9be05a00886c986d0b22adb6aec8cee798d7ecc437aa669ae9ac1d3a786a5d9e7d6a5b7c80e8df57b48d51624bbf0c0979387780a46a01eb01d032f1"),
		simaland.WithJWT("papkovda@me.com", "benryj-dirjIr-2rafge"),
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
	item, err := cl.GetItem(4381727)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(item)
}

func TestClient_FindSettlement(t *testing.T) {
	item, err := cl.FindSettlement("Москва")
	if err != nil {
		t.Error(err)
		return
	}

	if len(item.Items) < 1 {
		t.Error("fail: no items")
		return
	}

	st := item.Items[0]
	if !(st.Name == "Москва" || st.Region == "Московская область" || st.ID == 1686293227) {
		t.Error("unvalid data")
	}
}

func TestClient_ListDeliveryAddress(t *testing.T) {
	item, err := cl.FindSettlement("Москва")
	if err != nil {
		t.Error(err)
		return
	}

	if len(item.Items) == 0 {
		t.Error("no settlement")
	}

	list, err := cl.ListDeliveryAddress(item.Items[0].ID, 1, 100)
	if err != nil {
		t.Error(err)
		return
	}

	for _, obj := range list.Items {
		if obj.SettlementID != item.Items[0].ID {
			t.Error("point without settlement id")
		}
	}
}

func TestClient_CheckoutByProducts(t *testing.T) {
	list, err := cl.ListItems(1, 10, 19258)
	if err != nil {
		t.Error(err)
		return
	}

	cart := map[int64]int{
		list.Items[0].Sid: 3,
	}

	// for i := range list.Items {
	// 	cart[list.Items[i].Sid] = 10

	// 	if i >= 3 {
	// 		break
	// 	}
	// }

	data, err := cl.CheckoutByProducts("jet_mail@icloud.com", "79267027006", "Артур", cart)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(data)
}

func TestClient_GetOrder(t *testing.T) {
	data, err := cl.GetOrder(45808024)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(data)
}
