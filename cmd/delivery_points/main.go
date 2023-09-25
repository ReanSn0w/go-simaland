package main

import (
	"encoding/json"
	"os"

	"github.com/ReanSn0w/go-simaland/pkg/simaland"
	"github.com/go-pkgz/lgr"
)

func main() {
	points, err := loadPointsList()
	if err != nil {
		lgr.Default().Logf("[ERROR] %v", err)
		return
	}

	data, err := json.Marshal(points)
	if err != nil {
		lgr.Default().Logf("[ERROR] %v", err)
		return
	}

	_ = os.WriteFile("points.json", data, 0666)
}

func loadPointsList() ([]simaland.DeliveryPoint, error) {
	items := []simaland.DeliveryPoint{}

	cl, err := simaland.NewClient(
		simaland.WithLogger(lgr.Default()),
		simaland.WithApiKey("deb770fbb87c87cd415806c1f4260c3d8aeec54f0a2a351722403b9e546d11ea66aff0e082518e0cd2d7630bb5f5aa11995b0cea6b20ac9f5a010297409f458c"),
		simaland.WithJWT("papkovda@me.com", "benryj-dirjIr-2rafge"),
	)

	if err != nil {
		return items, err
	}

	page := 1

	for {
		res, err := cl.ListDeliveryAddress(0, page, 100)
		if err != nil {
			return items, err
		}

		items = append(items, res.Items...)

		if len(res.Items) != 100 {
			break
		}

		page++
	}

	return items, nil
}
