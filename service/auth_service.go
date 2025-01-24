package service

import (
	"api/database"
	"api/models"
	"api/utils"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Fungsi Login
func Login(email, password string) (string, string, error) {
	// Cari user di database
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", "", errors.New("user not found")
	} else if err != nil {
		return "", "", err
	}

	// Bandingkan password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", "", errors.New("invalid email or password")
	}

	// Buat access token dan refresh token
	accessToken, err := utils.CreateToken(user.ID, user.Email)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := utils.CreateRefreshToken(user.ID, user.Email)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func OtpVerify(otp string) map[string]interface{} {
    // Validasi input OTP
    if otp == "" {
        return map[string]interface{}{
            "error": "OTP tidak boleh kosong",
        }
    }

    // Cari user temporary
    var userTemp models.UserTemp
    result := database.DB.Where("otp = ?", otp).First(&userTemp)
    if result.Error != nil {
        return map[string]interface{}{
            "error": "OTP tidak valid",
        }
    }

    // Mulai transaksi database
    tx := database.DB.Begin()
    if tx.Error != nil {
        return map[string]interface{}{
            "error": "Gagal memulai transaksi database",
        }
    }

    // Buat user baru
    user := models.User{
        Name:     userTemp.Name,
        Email:    userTemp.Email,
        Password: userTemp.Password,
    }

    // Simpan user ke database
    if err := tx.Create(&user).Error; err != nil {
        tx.Rollback()
        return map[string]interface{}{
            "error": "Gagal membuat user",
        }
    }

    // Buat info user
    userInfo := models.UserInfo{
        UserID: user.ID,
    }

    // Simpan info user ke database
    if err := tx.Create(&userInfo).Error; err != nil {
        tx.Rollback()
        return map[string]interface{}{
            "error": "Gagal membuat info user",
        }
    }

    // Hapus user temporary
    if err := tx.Delete(&userTemp).Error; err != nil {
        tx.Rollback()
        return map[string]interface{}{
            "error": "Gagal menghapus user temporary",
        }
    }

    // Commit transaksi
    if err := tx.Commit().Error; err != nil {
        return map[string]interface{}{
            "error": "Gagal commit transaksi",
        }
    }

    return map[string]interface{}{
        "message": "Verifikasi berhasil, silakan login!",
    }
}
