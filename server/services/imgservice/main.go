package imgservice

import (
	"crypto/sha1"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"io"
	"mime/multipart"
	"os"
	"outstagram/server/models"
	"outstagram/server/repos/imgrepo"
	"time"
)

var ThumbnailSizes = []int{32, 150, 300, 400, 500}

type ImageService struct {
	imageRepo *imgrepo.ImageRepo
}

func New(imageRepo *imgrepo.ImageRepo) *ImageService {
	return &ImageService{imageRepo: imageRepo}
}

func (s *ImageService) Save(file *multipart.FileHeader, userID uint) (*models.Image, error) {
	names, err := s.createThumbnail(file, userID)
	if err != nil {
		return nil, err
	}

	img := models.Image{Tiny: names[0], Small: names[1], Medium: names[2], Big: names[3], Huge: names[4], Origin: names[5]}
	if err := s.imageRepo.Save(&img); err != nil {
		return nil, err
	}

	return &img, nil
}

func (s *ImageService) createThumbnail(fileHeader *multipart.FileHeader, userID uint) ([]string, error) {
	var names []string

	// Get filename for original image
	originalFilename := s.getRandomName(userID, len(ThumbnailSizes))

	// Save uploaded image to /images/<originalSizeFile>
	if err := s.saveFile(fileHeader, originalFilename); err != nil {
		return nil, err
	}

	// Open uploaded image for creating thumbnail
	originalFile, err := imaging.Open(fmt.Sprintf("images/%v.png", originalFilename))
	if err != nil {
		return nil, err
	}
	// Get image's width
	originalWidth := originalFile.Bounds().Max.X - 1

	for i, v := range ThumbnailSizes {
		var thumbnail *image.NRGBA
		// If thumbnail's width <= original image width
		if v <= originalWidth {
			thumbnail = imaging.Thumbnail(originalFile, v, v, imaging.Lanczos)
		} else {
			thumbnail = imaging.Thumbnail(originalFile, originalWidth, originalWidth, imaging.Lanczos)
		}

		// Get random filename for thumbnail
		randomName := s.getRandomName(userID, i)

		// Save thumbnail to images/<filename>.png
		if err = imaging.Save(thumbnail, fmt.Sprintf("images/%v.png", randomName)); err != nil {
			return nil, err
		}
		names = append(names, randomName+".png")
	}

	names = append(names, originalFilename+".png")
	return names, nil
}

func (s *ImageService) getRandomName(userID uint, i int) string {
	randomName := fmt.Sprintf("%x", sha1.Sum([]byte(fmt.Sprint(time.Now().Unix())+fmt.Sprint(userID)+fmt.Sprint(i))))
	return randomName
}

func (s *ImageService) saveFile(fileHeader *multipart.FileHeader, filename string) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	out, err := os.Create(fmt.Sprintf("images/%v.png", filename))
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	return err
}
