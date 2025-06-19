package utils

import (
	"context"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadToCloudinary(filePath string) (string, error) {
	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		return "", err
	}

	ctx := context.Background()

	resp, err := cld.Upload.Upload(ctx, filePath, uploader.UploadParams{
		Folder: "Balpay_image",
	})
	if err != nil {
		return "", err
	}

	return resp.SecureURL, nil
}
