package imgcontroller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"os"
	"outstagram/server/constants"
	"outstagram/server/models"
	"outstagram/server/utils"
	"regexp"
	"strconv"
	"strings"
)

func (ic *Controller) GetImage(c *gin.Context) {
	imageID, err := utils.StringToUint(c.Param("imageID"))
	if err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid userID", nil)
		return
	}

	image, err := ic.imageService.FindByID(imageID)
	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving user", err.Error())
		return
	}

	size := c.Query("size")
	file, err := readFileBySize(image, size)
	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving user", err.Error())
		return
	}

	defer func() {
		file.Close()
	}()

	fileStat, err := file.Stat()
	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while reading file", err.Error())
		return
	}

	c.DataFromReader(200, fileStat.Size(), "image/png", file, map[string]string{"Content-Disposition": `inline`})
}

func (ic *Controller) GetUserAvatar(c *gin.Context) {
	userID, err := utils.StringToUint(c.Param("userID"))
	if err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid userID", nil)
		return
	}

	user, err := ic.userService.FindByID(userID)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ResponseWithError(c, http.StatusNotFound, "User not fond", err.Error())
			return
		}

		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving user", err.Error())
		return
	}

	image, err := ic.imageService.FindByID(user.AvatarImageID)
	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving user", err.Error())
		return
	}

	size := c.Query("size")
	file, err := readFileBySize(image, size)
	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving user", err.Error())
		return
	}

	defer func() {
		file.Close()
	}()

	fileStat, err := file.Stat()
	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while reading file", err.Error())
		return
	}

	c.DataFromReader(200, fileStat.Size(), "image/png", file, map[string]string{"Content-Disposition": `inline`})
}

func readFileBySize(image *models.Image, size string) (*os.File, error) {
	fileURL := "./images/"

	switch size {
	case "":
		fileURL += utils.GetImageSize(image, constants.STDImageSizes[len(constants.STDImageSizes)-1])
	case "mini", "tiny", "small", "medium", "big", "huge", "origin":
		fileURL += utils.GetImageSize(image, strings.Title(size))
	default:
		if ok, err := regexp.MatchString(`^([0-9]+)x([0-9]+)$`, size); err != nil {
			return nil, err
		} else if !ok {
			return nil, errors.New("invalid size format")
		} else {
			reg, err := regexp.Compile(`[0-9]+`)
			if err != nil {
				return nil, errors.New("invalid regex")
			}
			dim := reg.FindAllString(size, 2)
			width, _ := strconv.ParseInt(dim[0], 10, 24)
			height, _ := strconv.ParseInt(dim[1], 10, 24)
			fileURL += utils.GetImageSize(image, strings.Title(getBestFitSize(int(width), int(height))))
		}
	}

	file, err := os.Open(fileURL)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func getBestFitSize(width, height int) string {
	// Smaller than the smallest of STDImageWidths
	if width <= constants.STDImageWidths[0] {
		return constants.STDImageSizes[0] // The smallest
	}

	// Larger than the largest of STDImageWidths
	if width > constants.STDImageWidths[len(constants.STDImageWidths)-1] {
		return constants.STDImageSizes[len(constants.STDImageSizes)-1] // The origin
	}

	for i := 1; i < len(constants.STDImageWidths); i++ {
		if width > constants.STDImageWidths[i-1] && width <= constants.STDImageWidths[i] {
			return constants.STDImageSizes[i] // Mini to Origin exclusive
		}
	}

	// IMPORTANT: This case never happens, must returns due to Go compiler
	return ""
}
