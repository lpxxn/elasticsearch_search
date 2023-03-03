package geo_demo

import (
	"fmt"
	"testing"

	"github.com/golang/geo/s2"
)

func TestGeoDistinct(t *testing.T) {
	// 伦敦
	london := s2.LatLngFromDegrees(51.507222, -0.1275)
	// 巴黎
	paris := s2.LatLngFromDegrees(48.8567, 2.3508)

	// 计算伦敦和巴黎之间的距离（以米为单位）
	distance := s2.PointFromLatLng(london).Distance(s2.PointFromLatLng(paris)).Radians() * 6371000

	fmt.Printf("The distance between London and Paris is %.2f km\n", distance/1000)
}
