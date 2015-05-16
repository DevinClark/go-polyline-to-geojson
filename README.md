# Google Encoded Polyline to GeoJSON
Decodes [Google's encoded polyline format](https://developers.google.com/maps/documentation/utilities/polylinealgorithm?csw=1) and outputs a GeoJSON LineString. CLI tool. Accepts string polyline as an argument or it can be piped in. Uses [go.geojson](https://github.com/paulmach/go.geojson) to convert the lat/lng list to GeoJSON.

### Usage
```bash
./polyline-to-geojson '}vzuEvnxqQspYthE'
```

Output for that example will be [this GeoJSON](http://geojson.io/#id=gist:DevinClark/74e9f18602a17247f48e).
