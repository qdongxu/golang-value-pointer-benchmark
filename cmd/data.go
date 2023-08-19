package cmd

import (
	"sync"
	"time"
)

type Data struct {
	ErrorCode   float64 `json:"errorCode"`
	Confidence  float64 `json:"confidence"`
	Stability   float64 `json:"stability"`
	Text        string  `json:"text"`
	PositionX   int     `json:"position_x"`
	PositionY   int     `json:"position_y"`
	ErrorCode2  float64 `json:"errorCode2"`
	Confidence2 float64 `json:"confidence2"`
	Stability2  float64 `json:"stability2"`
	Text2       string  `json:"text2"`
	PositionX2  int     `json:"position_x2"`
	PositionY2  int     `json:"position_y2"`
}

func NewDataValue() Data {
	return Data{
		ErrorCode:   0,
		Confidence:  0,
		Stability:   0,
		Text:        "",
		PositionX:   time.Now().Nanosecond(),
		PositionY:   2,
		ErrorCode2:  0,
		Confidence2: 0,
		Stability2:  0,
		Text2:       "",
		PositionX2:  3,
		PositionY2:  4,
	}
}

func NewDataPointer() *Data {
	return &Data{
		ErrorCode:   0,
		Confidence:  0,
		Stability:   0,
		Text:        "",
		PositionX:   time.Now().Nanosecond(),
		PositionY:   2,
		ErrorCode2:  0,
		Confidence2: 0,
		Stability2:  0,
		Text2:       "",
		PositionX2:  3,
		PositionY2:  4,
	}
}

func (d Data) calcWithValue(times int) int {
	if times == 0 {
		d.PositionX += d.PositionY
		return d.PositionX
	}

	return d.PositionX + d.PositionY + d.calcWithValue(times-1)
}

func (d *Data) calcWithPointer(times int) int {
	if times == 0 {
		d.PositionX += d.PositionY
		return d.PositionX
	}

	return d.PositionX + d.PositionY + d.calcWithPointer(times-1)
}

func calcWithValue() []Data {
	rel := make([]Data, 0, 100)
	for j := 0; j < 100; j++ {
		rel = append(rel, Data{
			ErrorCode:   0,
			Confidence:  0,
			Stability:   0,
			Text:        "",
			PositionX:   time.Now().Nanosecond(),
			PositionY:   2,
			ErrorCode2:  0,
			Confidence2: 0,
			Stability2:  0,
			Text2:       "",
			PositionX2:  3,
			PositionY2:  4,
		})
	}
	return rel
}

func sumWithValue(ds []Data) int {
	s := 0
	for _, d := range ds {
		s += d.PositionY
	}
	return s
}
func calcWithPointer() []*Data {
	rel := make([]*Data, 0, 100)
	for j := 0; j < 100; j++ {
		rel = append(rel, &Data{
			ErrorCode:   0,
			Confidence:  0,
			Stability:   0,
			Text:        "",
			PositionX:   time.Now().Nanosecond(),
			PositionY:   2,
			ErrorCode2:  0,
			Confidence2: 0,
			Stability2:  0,
			Text2:       "",
			PositionX2:  3,
			PositionY2:  4,
		})
	}
	return rel
}

func sumWithPointer(ds []*Data) int {
	s := 0
	for _, d := range ds {
		s += d.PositionY
	}
	return s
}

func ChanWithValue() int {
	wg := sync.WaitGroup{}

	dataChan := make(chan []Data, 100)
	sum := 0

	wg.Add(2)
	go func() {
		for i := 0; i < 1000; i++ {
			dataChan <- calcWithValue()
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			d := <-dataChan
			for _, v := range d {
				sum += v.PositionX2
			}
		}
		wg.Done()
	}()

	wg.Wait()
	return sum
}

func ChanWithPointer() int {
	wg := sync.WaitGroup{}

	dataChan := make(chan []*Data, 100)
	sum := 0

	wg.Add(2)
	go func() {
		for i := 0; i < 1000; i++ {
			dataChan <- calcWithPointer()
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			d := <-dataChan
			for _, v := range d {
				sum += v.PositionX2
			}
		}
		wg.Done()
	}()

	wg.Wait()
	return sum
}
