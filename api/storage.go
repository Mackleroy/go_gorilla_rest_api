package api

import "sync"

type ProjectStorage struct {
	sync.Mutex
	categories map[string][]Good
}

type Good struct {
	Name   string
	Tags   []string
	Price  uint32
	Length uint16
	Width  uint16
}

func (ps *ProjectStorage) Seed() {
	ps.Lock()
	ps.categories = make(map[string][]Good)
	ps.categories["sport"] = []Good{
		{Name: "soccer ball", Length: 30, Width: 30, Price: 100, Tags: []string{"football", "sport"}},
		{Name: "basketball ball", Price: 200, Length: 40, Width: 40, Tags: []string{"basketball", "sport"}},
		{Name: "tennis ball", Price: 50, Length: 10, Width: 10, Tags: []string{"tennis", "sport"}},
	}
	ps.categories["food"] = []Good{
		{Name: "white bread", Price: 100, Length: 10, Width: 10, Tags: []string{"bread", "food"}},
		{Name: "milk", Price: 100, Length: 10, Width: 10, Tags: []string{"milk-like", "food"}},
		{Name: "cheese", Price: 100, Length: 10, Width: 10, Tags: []string{"milk-like", "food"}},
	}
	ps.categories["tech"] = []Good{
		{Name: "laptop 2000", Price: 100000, Length: 10, Width: 10, Tags: []string{"portable", "tech"}},
		{Name: "smartphone pro max", Price: 100, Length: 10, Width: 10, Tags: []string{"smartphone", "tech"}},
		{Name: "tablet tab", Price: 100, Length: 10, Width: 10, Tags: []string{"portable", "tech"}},
	}
	ps.Unlock()
}

func (ps *ProjectStorage) AddCategory(category string) error {
	ps.Lock()
	if _, ok := ps.categories[category]; ok {
		return &ErrObjectAlreadyExists{}
	}
	ps.categories[category] = []Good{}
	ps.Unlock()
	return nil
}

func (ps *ProjectStorage) AddGood(data CreateGoodRequest) error {
	ps.Lock()
	if _, ok := ps.categories[data.CategoryName]; !ok {
		return &ErrObjectNotFound{}
	}

	for _, good := range ProjectStorageInstance.categories[data.CategoryName] {
		if good.Name == data.GoodName {
			return &ErrObjectAlreadyExists{}
		}
	}
	ps.categories[data.CategoryName] = append(
		ProjectStorageInstance.categories[data.CategoryName],
		Good{
			Name:   data.GoodName,
			Tags:   data.Tags,
			Price:  data.Price,
			Length: data.Length,
			Width:  data.Width,
		})
	ps.Unlock()
	return nil
}

func (ps *ProjectStorage) HasCategory(category string) bool {
	_, ok := ps.categories[category] // kind of Exists query
	return ok
}

func GetProjectStorage() *ProjectStorage {
	instance := &ProjectStorage{}
	instance.Seed()
	return instance
}

var ProjectStorageInstance = GetProjectStorage()
