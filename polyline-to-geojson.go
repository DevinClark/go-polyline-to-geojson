package main

import "os"
import "log"
import "io"
import "io/ioutil"
import "github.com/paulmach/go.geojson" // https://godoc.org/github.com/paulmach/go.geojson

func decodeToken(bytes []byte) (pos int, value float64) {
  var token int64 = 0
  var shift uint = 0
  var result float64
  var factor float64 = 1e5

  for i, v := range bytes {
    current := int64(v) - 63
    token |= (current & 0x1f) << uint(shift)
    shift += 5

    if current & 0x20 == 0 {
      pos = i + 1

      if token & 1 != 0 {
        result = float64(^(token >> 1))
      } else {
        result = float64(token >> 1)
      }

      value = result / factor
      return
    }
  }

  pos = 0
  return
}

func decodePolyline(bytes []byte) ([]byte) {

  fc := geojson.NewFeatureCollection()
  coords := make([][]float64, 0)
  var pos int = 0
  var lat, lng float64

  for pos < len(bytes) {
    current, current_lat := decodeToken(bytes[pos:len(bytes)])
    pos += current
    lat += current_lat

    current, current_lng := decodeToken(bytes[pos:len(bytes)])
    pos += current
    lng += current_lng

    coord := []float64{lng, lat}
    coords = append(coords, coord)
  }

  line := geojson.NewLineStringFeature(coords)

  fc.AddFeature(line)

  json, _ := fc.MarshalJSON()

  return json
}

func main() {
  args := os.Args
  bytes := []byte("")
  err := error(nil)

  if len(args) > 1 && args[1] != "" {
    bytes = []byte(args[1])
  } else {
    bytes, err = ioutil.ReadAll(os.Stdin)

    if err != nil {
      if err == io.EOF {
        return
      }

      log.Println(err)
    }
  }

  geojson_output := decodePolyline(bytes)

  os.Stdout.Write(geojson_output)

}
