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

/*
在这个示例中，我们使用了 s2.LatLngFromDegrees() 函数将经纬度转换为 s2.LatLng 类型，并将它们传递给 s2.PointFromLatLng() 函数来创建一个代表地球表面上一个点的 s2.Point 类型。
然后，我们使用 Distance() 函数来计算两个点之间的距离，并将结果转换为以米为单位的距离。
请注意，我们在距离计算中使用了地球的半径（6371 公里），这是一个常见的近似值。如果需要更高的精度，可以考虑使用更准确的地球半径值。
*/
