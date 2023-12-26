package constant

const LongitudeMin = 120.35
const LongitudeMax = 120.37
const LatitudeMin = 30.29
const LatitudeMax = 30.31

func InBound(longitude float64, latitude float64) bool {
	longitudeInBound := longitude >= LongitudeMin && longitude <= LongitudeMax
	latitudeInBound := latitude >= LatitudeMin && latitude <= LatitudeMax
	return longitudeInBound && latitudeInBound
}
