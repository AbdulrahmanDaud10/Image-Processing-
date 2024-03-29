package tasks

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"time"

	"github.com/nfnt/resize"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
)

const TypeResizeImage = "image:resize"

var StandardWidths = []uint{16, 32, 128, 240, 320, 480, 540, 640, 800, 1024}

type ResizeImagePayLoad struct {
	ImageData []byte
	Width     uint
	Height    uint
	FileName  string
}

func NewImageResizeTask(imageData []byte, fileName string) ([]*asynq.Task, error) {
	img, _, err := image.Decode(bytes.NewReader(imageData))
	if err != nil {
		return nil, err
	}

	originalBounds := img.Bounds()
	originalWidth := uint(originalBounds.Dx())
	originalHeight := uint(originalBounds.Dy())

	var tasks []*asynq.Task
	for _, width := range StandardWidths {
		height := (width * originalHeight) / originalWidth
		payload := ResizeImagePayLoad{
			ImageData: imageData,
			Width:     width,
			Height:    height,
			FileName:  fileName,
		}

		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}

		task := asynq.NewTask(TypeResizeImage, payloadBytes)
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func HandleResizeImageTask(ctx context.Context, task *asynq.Task) error {
	var payload ResizeImagePayLoad
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to parse resize image task payload: %v", err)
	}

	img, _, err := image.Decode(bytes.NewReader(payload.ImageData))
	if err != nil {
		return fmt.Errorf("image decode failed: %v", err)
	}

	resizedImg := resize.Resize(payload.Width, payload.Height, img, resize.Lanczos3)
	outputUUID := uuid.New()
	outputFileName := fmt.Sprintf("images/%s/%s%s", time.Now().Format("2006-01-02"), outputUUID.String(), filepath.Ext(payload.FileName))

	outputDir := filepath.Dir(outputFileName)
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return err
		}
	}

	outFile, err := os.Create(outputFileName)
	if err != nil {
		return err
	}

	defer outFile.Close()
	if err := jpeg.Encode(outFile, resizedImg, nil); err != nil {
		return err
	}
	fmt.Printf("Output UUID for the processed image: %s\n", outputUUID.String())

	return nil
}
