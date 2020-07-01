package global

import "testing"

func TestRandomNumber(t *testing.T) {
	for i := 0; i < 1000; i++ {
		number, err := RandomNumber(6)
		if err != nil {
			t.Error(err)
		}
		t.Log(number)
	}
}

func TestRandomString(t *testing.T) {
	for i := 0; i < 1000; i++ {
		str := RandomString(6)
		t.Log(str)
	}
}

func TestHashPassword(t *testing.T) {
	for i := 0; i < 1000; i++ {
		sourceStr := RandomString(6)
		t.Log("Source String:", sourceStr)
		hashStr, salt := HashPassword(sourceStr)
		t.Log("Hash String:", hashStr, ",Salt:", salt)

		b := ValidatePassword(sourceStr, salt, hashStr)
		t.Log("Validate:", b)
		t.Log("---------------------------------------")
	}

}
