package main

import "fmt"

type BeltSize int
type Shipping int

const (
	Small = iota
	Medium
	Large
)

const (
	Ground = iota
	Air
)

func (b BeltSize) String() string {
	return []string{"Small", "Medium", "Large"}[b]
}

func (s Shipping) String() string {
	return []string{"Ground", "Air"}[s]
}

type Conveyor interface {
	Convey() BeltSize
}

type Shipper interface {
	Ship() Shipping
}

type WarehouseAutomator interface {
	Conveyor
	Shipper
}

// 这里的类型别名使用大写不会报unused错
type SpamMail struct {
	amount int
}

func (s SpamMail) String() string {
	return "Spam Mail"
}

func (s *SpamMail) Convey() BeltSize {
	return Small
}

func (s *SpamMail) Ship() Shipping {
	return Air
}

func automate(item WarehouseAutomator) {
	fmt.Printf("Convery %v on %v conveyor", item, item.Convey())
	fmt.Printf("Ship %v via %v shipper", item, item.Ship())
}

type ToxicWaste struct {
	weight int
}

func (t *ToxicWaste) Ship() Shipping {
	fmt.Println("Weight:", t.weight)
	return Ground
}

// 感觉embedding很像类的interface的继承extends
func main() {
	mail := SpamMail{40000}
	automate(&mail)

	// waste := ToxicWaste{300}
	// automate(&waste)
	// 这里由于没有implement所有的接口，所以会报错
}
