package property

// Returns the number or bathroom based on full and half bath
func GetBathRoomCount(baths_full int8, baths_half int8)  float32  {
	if(baths_half > 0) {
		return float32(baths_full) + 0.5
	} else {
		return float32(baths_full);
	}
}