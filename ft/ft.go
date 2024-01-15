package ft

import (
	"fmt"
	"strings"
)

var fileExtensions = map[string]string{
	".mp3":  "Audio",
	".wav":  "Audio",
	".aac":  "Audio",
	".flac": "Audio",
	".m4a":  "Audio",
	".mp4":  "Video",
	".avi":  "Video",
	".mkv":  "Video",
	".mov":  "Video",
	".wmv":  "Video",
	".jpg":  "Images",
	".jpeg": "Images",
	".png":  "Images",
	".gif":  "Images",
	".bmp":  "Images",
	".svg":  "Images",
	".tiff": "Images",
	".tif":  "Images",
	".pdf":  "Documents",
	".docx": "Documents",
	".xlsx": "Documents",
	".pptx": "Documents",
	".txt":  "Documents",
	".mobi": "Documents",
	".epub": "Documents",
	".zip":  "Archives",
	".rar":  "Archives",
	".7z":   "Archives",
	".tar":  "Archives",
	".gz":   "Archives",
	".ttf":  "Archives",
	".otf":  "Archives",
	".woff": "Archives",
	".exe":  "Executables",
	".app":  "Executables",
	".dmg":  "Executables",
	".bin":  "Executables",
	".iso":  "Executables",
	".msi":  "Executables",
	".deb":  "Executables",
	".csv":  "Data",
	".json": "Data",
	".xml":  "Data",
}

func GetFileCategory(extension string) (string, error) {
	if !strings.HasPrefix(extension, ".") {
		extension = fmt.Sprintf(".%s", extension)
	}

	category, found := fileExtensions[extension]
	if !found {
		return "", fmt.Errorf("%s not found ", extension)
	}

	return category, nil
}
