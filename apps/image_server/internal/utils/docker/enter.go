package docker

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

const manifestFile = "manifest.json"

func ParseImageMetadata(filePath string) (string, string, string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", "", "", err
	}
	defer file.Close()

	var reader io.Reader = file

	// Handle gzipped files
	if strings.HasSuffix(filePath, ".gz") {
		gzReader, err := gzip.NewReader(file)
		if err != nil {
			return "", "", "", err
		}
		defer gzReader.Close()
		reader = gzReader
	}

	tarReader := tar.NewReader(reader)

	var imageID, imageName, imageTag string

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", "", "", err
		}

		switch header.Name {
		case manifestFile:
			// Parse manifest.json to get image ID
			manifestData, err := io.ReadAll(tarReader)
			if err != nil {
				return "", "", "", err
			}
			// This is a simplified parsing - you might need a proper JSON parser
			data, err := extractImage(string(manifestData))
			if err != nil {
				return "", "", "", err
			}
			return data.ImageID, data.ImageName, data.ImageTag, nil
		}
	}

	if imageID == "" || imageName == "" || imageTag == "" {
		return "", "", "", fmt.Errorf("无法从镜像文件中提取完整的元数据")
	}

	return imageID, imageName, imageTag, nil
}

type manifestType struct {
	Config   string   `json:"Config"`
	RepoTags []string `json:"RepoTags"`
}

type manifestData struct {
	ImageID   string
	ImageName string
	ImageTag  string
}

// extractImage 解析manifest文件
func extractImage(manifest string) (data manifestData, err error) {
	// Simplified extraction - real implementation should parse JSON properly
	var t []manifestType
	err = json.Unmarshal([]byte(manifest), &t)
	if err != nil {
		err = fmt.Errorf("解析manifest文件失败 %s", err)
		return
	}
	if len(t) == 0 {
		err = fmt.Errorf("解析manifest文件内容失败 %s", manifest)
		return
	}
	if len(t[0].RepoTags) == 0 {
		err = fmt.Errorf("传入的镜像没有tag，无法解析 %s", manifest)
		return
	}
	repoTags := t[0].RepoTags[0]
	_list := strings.Split(repoTags, ":")
	data.ImageName = _list[0]
	data.ImageTag = _list[1]
	data.ImageID = strings.Split(t[0].Config, "/")[2][:12]
	return
}
