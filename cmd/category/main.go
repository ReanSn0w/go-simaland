package main

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"

	"github.com/ReanSn0w/go-simaland/pkg/simaland"
	"github.com/go-pkgz/lgr"
	"gopkg.in/yaml.v3"
)

var (
	log        = lgr.New(lgr.Debug)
	categories = []simaland.Category{}
	items      = []Item{}
	result     = []*Category{}
)

type Item struct {
	ID    int
	Name  string
	Path  []int
	Level int
}

func (i *Item) ParentID() int {
	if len(i.Path) == 1 {
		return i.Path[0]
	}

	return i.Path[len(i.Path)-2]
}

type Category struct {
	Title      string      `yaml:"title"`
	Image      string      `yaml:"image,omitempty"`
	SimalandID []int       `yaml:"simaland_id"`
	Items      []*Category `yaml:"items,omitempty"`
}

func main() {
	loadCategoryList()

	// Оставляем только нужную информацию из списка
	for i := range categories {
		items = append(items, Item{
			ID:   int(categories[i].ID),
			Name: categories[i].Name,
			Path: func() []int {
				pathString := categories[i].Path

				pathStrings := strings.Split(pathString, ".")

				path := make([]int, 0)

				for i := range pathStrings {
					v, err := strconv.Atoi(pathStrings[i])
					if err != nil {
						continue
					}
					path = append(path, v)
				}

				return path
			}(),
			Level: int(categories[i].Level),
		})
	}

	// строим список категорий
	{
		stage := map[int]*Category{}

		for i := range items {
			item := items[i]

			stage[item.ID] = &Category{
				Title:      item.Name,
				SimalandID: []int{item.ID},
				Items:      make([]*Category, 0),
			}
		}

		for i := range items {
			item := stage[items[i].ID]
			parentID := items[i].ParentID()
			if parentID == item.SimalandID[0] {
				continue
			}

			parentItem, ok := stage[parentID]
			if !ok {
				continue
			}

			parentItem.Items = append(parentItem.Items, item)
		}

		for i := range items {
			if len(items[i].Path) == 1 {
				result = append(result, stage[items[i].ID])
			}
		}
	}

	resultData, err := yaml.Marshal(result)
	if err != nil {
		log.Logf("[ERROR] marshal file error: %v", err)
		os.Exit(2)
	}

	err = os.WriteFile("list.yaml", resultData, 0666)
	if err != nil {
		log.Logf("[ERROR] write file error: %v", err)
		os.Exit(2)
	}
}

func loadCategoryList() {
	if tryToLoadPlain() {
		return
	}

	cl, err := simaland.NewClient(
		simaland.WithLogger(log),
		simaland.WithApiKey("deb770fbb87c87cd415806c1f4260c3d8aeec54f0a2a351722403b9e546d11ea66aff0e082518e0cd2d7630bb5f5aa11995b0cea6b20ac9f5a010297409f458c"),
		simaland.WithJWT("papkovda@me.com", "benryj-dirjIr-2rafge"),
	)

	if err != nil {
		log.Logf("[ERROR] create client error: %v", err)
		os.Exit(2)
	}

	// Получаем список категорий
	{
		iterator := 1

		for {
			cats, err := cl.ListCategory(iterator, 100)
			if err != nil {
				log.Logf("[ERROR] category loading error: %v", err)
				os.Exit(2)
			}

			categories = append(categories, cats.Items...)

			if cats.Meta.PageCount <= int64(iterator) {
				break
			}

			iterator++
		}
	}

	savePlain()
}

func savePlain() {
	data, err := json.Marshal(categories)
	if err != nil {
		return
	}

	os.WriteFile("plain.json", data, 0664)
}

func tryToLoadPlain() bool {
	file, err := os.ReadFile("plain.json")
	if err != nil {
		return false
	}

	err = json.Unmarshal(file, &categories)
	return err == nil
}
