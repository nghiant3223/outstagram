package imgservice

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/segmentio/ksuid"
	"image"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"outstagram/server/constants"
	"outstagram/server/models"
	"outstagram/server/repos/imgrepo"
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

func (s *ImageService) SaveURL(url string, userID uint, isThumbnail bool) (*models.Image, error) {
	names, err := s.processImageURL(url, userID, isThumbnail)
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
	// Get filename for original image
	originalFilename := s.getRandomName()

	// Save uploaded image to /images/<originalSizeFile>
	if err := s.saveFileByHeader(fileHeader, originalFilename); err != nil {
		return nil, err
	}

	// Open uploaded image for creating thumbnail
	originalFile, err := imaging.Open(fmt.Sprintf("images/%v.png", originalFilename))
	if err != nil {
		return nil, err
	}

	return s.createDifferentSizes(originalFile, isThumbnail, userID, originalFilename)
}

func (s *ImageService) processImageURL(url string, userID uint, isThumbnail bool) ([]string, error) {
	// Get filename for original image
	originalFilename := s.getRandomName()

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	file, err := os.Create(fmt.Sprintf("images/%v.png", originalFilename))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return nil, err
	}

	// Open uploaded image for creating thumbnail
	originalFile, err := imaging.Open(fmt.Sprintf("images/%v.png", originalFilename))
	if err != nil {
		return nil, err
	}

	return s.createDifferentSizes(originalFile, isThumbnail, userID, originalFilename)
}

func (s *ImageService) createDifferentSizes(img image.Image, isThumbnail bool, userID uint, originalFilename string) ([]string, error) {
	var names []string
	originalWidth := img.Bounds().Max.X - 1

	for _, stdWidth := range constants.STDImageWidths {
		var resizeImage *image.NRGBA

		// If resizeImage's width <= original resizeImage width
		if stdWidth <= originalWidth {
			if isThumbnail {
				resizeImage = imaging.Thumbnail(img, stdWidth, stdWidth, imaging.Lanczos)
			} else {
				resizeImage = imaging.Resize(img, stdWidth, 0, imaging.Lanczos)
			}
		} else {
			if isThumbnail {
				resizeImage = imaging.Thumbnail(img, originalWidth, originalWidth, imaging.Lanczos)
			} else {
				resizeImage = imaging.Resize(img, originalWidth, 0, imaging.Lanczos)
			}
		}

		// Get random filename for resizeImage
		randomName := s.getRandomName()

		// Save resizeImage to images/<filename>.png
		if err := imaging.Save(resizeImage, fmt.Sprintf("images/%v.png", randomName)); err != nil {
			return nil, err
		}
		names = append(names, randomName+".png")
	}

	names = append(names, originalFilename+".png")
	return names, nil
}

func (s *ImageService) getRandomName() string {
	randomName := ksuid.New().String()
	return randomName
}

func (s *ImageService) saveFileByHeader(fileHeader *multipart.FileHeader, filename string) error {
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
