package service

import (
	"errors"
	"mime/multipart"
	"os"

	"superTools-background/global"
	"superTools-background/pkg/upload"
)

/**
* @Author: super
* @Date: 2020-09-23 19:07
* @Description: 文件上传service代码
**/

type FileInfo struct {
	Name     string `json:"name"`
	TempPath string `json:"temp_path"`
	Url      string `json:"url"`
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported")
	}
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit")
	}

	uploadSavePath := upload.GetSavePath()
	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions")
	}

	dst := uploadSavePath + "/" + fileName
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessUrl := uploadSavePath + "/" + fileName
	url := global.AppSetting.UploadServerUrl + "/" + accessUrl
	return &FileInfo{Name: fileName, TempPath: accessUrl, Url: url}, nil
}
