package sample

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/petrostrak/gRPC-with-Go/pb/pb"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTZ
	case 2:
		return pb.Keyboard_QWERTY
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomInt(min, max int) int {
	// min + 0 = min
	// min + max - min = max
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomCPUName(name string) string {
	if name == "Intel" {
		return randomStringFromSet(
			"Xeon E-2286M",
			"Core i9-9980HK",
			"Core i7-9980HK",
			"Core i5-9980HK",
			"Core i3-9980HK",
		)
	}

	return randomStringFromSet(
		"Ryzen 7 PRO 2700U",
		"Ryzen 5 PRO 2700U",
		"Ryzen 3 PRO 2700U",
	)
}

func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"RTX 2060",
			"RTX 2070",
			"GTX 1600-Ti",
			"GTX 1070",
		)
	}

	return randomStringFromSet(
		"RX 590",
		"RX 580",
		"RX 5700-XT",
		"RX Vega-56",
	)
}

func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}

	return a[rand.Intn(n)]
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)

}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9

	return &pb.Screen_Resolution{
		Height: uint32(height),
		Width:  uint32(width),
	}
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}

	return pb.Screen_OLED
}

func randomID() string {
	return uuid.New().String()
}

func randomLaptopBrand() string {
	return randomStringFromSet(
		"Apple",
		"Dell",
		"Lenovo",
		"Huawei",
	)
}

func randomLaptopName(name string) string {
	switch name {
	case "Apple":
		return randomStringFromSet("MacBook Air", "Macbook Pro")
	case "Dell":
		return randomStringFromSet("Latitude", "XPS", "Alienware")
	case "Lenovo":
		return randomStringFromSet("Thinkpad X1", "Thinkpad P1", "Thingpad P53")
	default:
		return "Matebook 13"
	}
}
