package util

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/cshum/vipsgen/vips"
)

var (
	vipsOnce    sync.Once
	vipsInitErr error
)

func vipsInit() error {
	vipsOnce.Do(func() {
		// Startup 会检查版本并初始化 libvips
		vips.Startup(nil)
	})
	return vipsInitErr
}

// GetImageSize 只读取图片头部获取尺寸，避免加载整图与 CGO 依赖
func GetImageSizeFromPath(path string) (width, height int, err error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, 0, err
	}
	defer func() {
		_ = f.Close()
	}()

	cfg, _, err := image.DecodeConfig(f)
	if err != nil {
		return 0, 0, err
	}

	return cfg.Width, cfg.Height, nil
}

// GetImageSizeFromFile 从文件获取图片尺寸
func GetImageSizeFromFile(file *multipart.FileHeader) (width, height int, err error) {
	reader, err := file.Open()
	if err != nil {
		return 0, 0, err
	}
	defer func() {
		_ = reader.Close()
	}()
	return GetImageSizeFromReader(reader)
}

// GetImageSizeFromReader 从 Reader 获取图片尺寸
func GetImageSizeFromReader(reader io.Reader) (width, height int, err error) {
	cfg, _, err := image.DecodeConfig(reader)
	if err != nil {
		return 0, 0, err
	}
	return cfg.Width, cfg.Height, nil
}

// ConvertImage 转换图片格式
func ConvertImage(path, outputFormat string) error {
	if err := vipsInit(); err != nil {
		return err
	}

	img, err := vips.NewImageFromFile(path, nil)
	if err != nil {
		return err
	}
	defer img.Close()

	format := strings.TrimPrefix(strings.ToLower(outputFormat), ".")
	if format == "" {
		return fmt.Errorf("output format required")
	}

	base := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	outPath := filepath.Join(filepath.Dir(path), base+"."+format)

	switch format {
	case "webp":
		opts := vips.DefaultWebpsaveOptions()
		opts.Q = 85
		return img.Webpsave(outPath, opts)
	case "avif":
		opts := vips.DefaultHeifsaveOptions()
		opts.Q = 80
		opts.Compression = vips.HeifCompressionAv1
		opts.Encoder = vips.HeifEncoderAom
		return img.Heifsave(outPath, opts)
	case "png":
		return img.Pngsave(outPath, nil)
	case "jpeg", "jpg":
		opts := vips.DefaultJpegsaveOptions()
		opts.Q = 85
		return img.Jpegsave(outPath, opts)
	default:
		return fmt.Errorf("unsupported format: %s", outputFormat)
	}
}
