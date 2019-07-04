package imgservice

import (
	"crypto/sha1"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"io"
	"mime/multipart"
	"os"
	"outstagram/server/constants"
	"outstagram/server/models"
	"outstagram/server/repos/imgrepo"
	"time"
)

type ImageService struct {
	imageRepo *imgrepo.ImageRepo
}

func New(imageRepo *imgrepo.ImageRepo) *ImageService {
	return &ImageService{imageRepo: imageRepo}
}

func (s *ImageService) Save(file *multipart.FileHeader, userID uint, isThumbnail bool) (*models.Image, error) {
	names, err := s.processImage(file, userID, isThumbnail)
	if err != nil {
		return nil, err
	}

	img := models.Image{Mini: names[0], Tiny: names[1], Small: names[2], Medium: names[3], Big: names[4], Huge: names[5], Origin: names[6]}
	if err := s.imageRepo.Save(&img); err != nil {
		return nil, err
	}

	return &img, nil
}

func (s *ImageService) FindByID(id uint) (*models.Image, error) {
	return s.imageRepo.FindByID(id)
}

func (s *ImageService) processImage(fileHeader *multipart.FileHeader, userID uint, isThumbnail bool) ([]string, error) {
	var names []string

	// Get filename for original image
	originalFilename := s.getRandomName(userID, len(constants.STDImageWidths))

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

	for i, stdWidth := range constants.STDImageWidths {
		var image *image.NRGBA

		// If image's width <= original image width
		if stdWidth <= originalWidth {
			if isThumbnail {
				image = imaging.Thumbnail(originalFile, stdWidth, stdWidth, imaging.Lanczos)
			} else {
				image = imaging.Resize(originalFile, stdWidth, 0, imaging.Lanczos)
			}
		} else {
			if isThumbnail {
				image = imaging.Thumbnail(originalFile, originalWidth, originalWidth, imaging.Lanczos)
			} else {
				image = imaging.Resize(originalFile, originalWidth, 0, imaging.Lanczos)
			}
		}

		// Get random filename for image
		randomName := s.getRandomName(userID, i)

		// Save image to images/<filename>.png
		if err = imaging.Save(image, fmt.Sprintf("images/%v.png", randomName)); err != nil {
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
