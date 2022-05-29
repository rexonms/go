package property

// Returns the number or bathroom based on full and half bath
func GetBathRoomCount(baths_full uint8, baths_half uint8)  float32  {
	if(baths_half > 0) {
		return float32(baths_full) + 0.5
	} else {
		return float32(baths_full);
	}
}