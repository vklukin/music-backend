package internal

import (
	"net/http"
	"os"

	"music-backend/services/logging"
)

type getFolderDataError struct {
	status int
	message string
}

type getFolderDataResult struct {

}

func GetFoldersData(path string) any {
	files, err := os.ReadDir(path)
	if err != nil {
		logging.LogWithError("Open folder", err)
		return &getFolderDataError{
			status: http.StatusInternalServerError,
			message: "Cannot open folder with music",
		}	
	}

	
}