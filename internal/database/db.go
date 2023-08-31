package database

import (
	"avito/internal/config"
	"avito/internal/models"
	"avito/internal/tools"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func CreateUser(user models.User) (uint, error) {
	result := config.DB.Create(&user)
	if result.RowsAffected != 0 {
		return user.ID, nil
	}
	return 0, errors.New("user not created")
}

func GetUsers() ([]models.User, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return []models.User{}, err
	}
	return users, nil
}

func CreateSegment(segment models.Segment) (uint, error) {
	result := config.DB.Create(&segment)
	if result.RowsAffected != 0 {
		return segment.ID, nil
	}
	return 0, errors.New("segment not created")
}

func AddSegmentToUserWithExpiredTime(tx *gorm.DB, segmentName string, userId string, expired string) error {
	date := strings.Split(expired, "-")
	year, _ := strconv.Atoi(date[0])
	month, _ := strconv.Atoi(date[1])
	day, _ := strconv.Atoi(date[2])
	user := models.User{}
	segment := models.Segment{}
	if err := tx.First(&user, userId).Error; err != nil {
		return err
	}
	if err := tx.Where("Name = ?", segmentName).Find(&segment).Error; err != nil {
		return err
	}
	userSegmentAssociation := models.UserSegments{
		UserID:    int(user.ID),
		SegmentID: int(segment.ID),
		ExpiresAt: time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local),
	}
	if err := tx.Create(&userSegmentAssociation).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&user).Association("Segments").Append(&segment); err != nil {
		tx.Rollback()
		return err
	}
	_ = AddSegmentHistory(userId, segmentName, tools.AddSegment)
	return tx.Error
}

func AddSegmentToUser(tx *gorm.DB, segmentName string, userId string) error {
	user := models.User{}
	segment := models.Segment{}
	if err := tx.First(&user, userId).Error; err != nil {
		return err
	}
	if err := tx.Where("Name = ?", segmentName).Find(&segment).Error; err != nil {
		return err
	}
	err := tx.Model(&user).Association("Segments").Append(&segment)
	if err == nil {
		_ = AddSegmentHistory(userId, segmentName, tools.AddSegment)
	}
	return tx.Error
}
func DeleteSegmentFromUser(tx *gorm.DB, segmentName string, userId string) error {
	user := models.User{}
	segment := models.Segment{}
	if err := tx.First(&user, userId).Error; err != nil {
		return err
	}
	if err := tx.Where("Name = ?", segmentName).First(&segment).Error; err != nil {
		return err
	}
	segments, _ := GetUserWithSegments(userId)
	if !tools.Contains(segments, segment) {
		return errors.New("user already doesn't have this segment")
	}
	err := tx.Model(&user).Association("Segments").Delete(&segment)
	if err == nil {
		_ = AddSegmentHistory(userId, segmentName, tools.DeleteSegment)
	}
	return tx.Error
}

func GetUserWithSegments(userId string) ([]*models.Segment, error) {
	var user models.User
	if err := config.DB.Preload("Segments").First(&user, userId).Error; err != nil {
		return user.Segments, err
	}
	return user.Segments, nil
}

func DeleteSegment(segment models.Segment) error {
	if err := config.DB.Where("Name = ?", segment.Name).Delete(&segment).Error; err != nil {
		return err
	}
	return nil
}
func AddSegmentHistory(userID, segmentName, operation string) error {
	history := models.SegmentHistory{
		UserID:      userID,
		SegmentName: segmentName,
		Operation:   operation,
		Timestamp:   time.Now(),
	}
	return config.DB.Create(&history).Error
}
func GenerateCSVReport(userID string, year int, month time.Month) (string, error) {
	var history []models.SegmentHistory
	startDate := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, 0)

	if err := config.DB.Where("user_id = ? AND timestamp >= ? AND timestamp < ?", userID, startDate, endDate).Find(&history).Error; err != nil {
		return "", err
	}

	var lines []string
	for _, entry := range history {
		line := fmt.Sprintf("%s;%s;%s;%s", entry.UserID, entry.SegmentName, entry.Operation, entry.Timestamp.String())
		lines = append(lines, line)
	}

	reportContent := strings.Join(lines, "\n")
	return reportContent, nil
}

func DeleteExpiredSegments() {
	now := time.Now()
	var expiredSegments []models.UserSegments
	config.DB.Where("expires_at < ?", now).Find(&expiredSegments)
	for _, segment := range expiredSegments {
		tx := config.DB.Begin()
		var cur models.Segment
		tx.Where("id = ?", segment.SegmentID).First(&cur)
		err := DeleteSegmentFromUser(tx, cur.Name, strconv.Itoa(segment.UserID))
		tx.Delete(&segment)
		if err != nil {
			fmt.Println(err)
			return
		}
		errTx := tx.Commit().Error
		if errTx != nil {
			fmt.Println(errTx)
			return
		}
	}
}

func RandomApply(percentage int, segmentName string) error {
	var count int64
	if err := config.DB.Model(&models.User{}).Count(&count).Error; err != nil {
		return err
	}
	m := make(map[int]bool)
	tx := config.DB.Begin()
	for i := 0; i < int(math.Round(float64(int(count)*percentage/100))); {

		randNum := rand.Intn(int(count)) + 1
		if m[randNum] == false {
			err := AddSegmentToUser(tx, segmentName, strconv.Itoa(randNum))
			if err != nil {
				tx.Rollback()
				return err
			}
			i++
			m[randNum] = true
		}
	}
	errTx := tx.Commit().Error
	if errTx != nil {
		return errTx
	}
	return nil
}
