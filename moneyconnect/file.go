package moneyconnect

import (
	"errors"
	"github.com/cyjaysong/lakala-go/common"
	"github.com/cyjaysong/lakala-go/moneyconnect/models"
	"os"
)

// ImageUploadFromBytes [imageType]可选值jpg、png、bmp、gif;
func (t *Client) ImageUploadFromBytes(imageType, imageName string, imageBytes []byte) (body *models.BaseResponse[models.ImageUploadResult], err error) {
	val := t.getBaseBody("image.upload")
	val["image_type"], val["image_name"] = imageType, imageName
	var req = t.buildUploadRequest(val, imageType, imageName)
	req.SetFileBytes("image_content", imageName, imageBytes)
	body = new(models.BaseResponse[models.ImageUploadResult])
	err = t.reqPost(req, body)
	return
}

// ImageUpload [imageType]可选值jpg、png、bmp、gif;
func (t *Client) ImageUpload(imageType, imageName string, imageFile *os.File) (body *models.BaseResponse[models.ImageUploadResult], err error) {
	val := t.getBaseBody("image.upload")
	val["image_type"], val["image_name"] = imageType, imageName
	var req = t.buildUploadRequest(val, imageType, imageName)
	req.SetFileReader("image_content", imageName, imageFile)
	body = new(models.BaseResponse[models.ImageUploadResult])
	err = t.reqPost(req, body)
	return
}

// FileUploadFromBytes [imageType]可选值pdf、doc、docx、zip
func (t *Client) FileUploadFromBytes(fileType, fileName string, fileBytes []byte) (body *models.BaseResponse[models.FileUploadResult], err error) {
	val := t.getBaseBody("file.upload")
	val["file_type"], val["file_name"] = fileType, fileName
	var req = t.buildUploadRequest(val, fileType, fileName)
	req.SetFileBytes("file_content", fileName, fileBytes)
	body = new(models.BaseResponse[models.FileUploadResult])
	err = t.reqPost(req, body)
	return
}

// FileUpload [imageType]可选值pdf、doc、docx、zip
func (t *Client) FileUpload(fileType, fileName string, imageFile *os.File) (body *models.BaseResponse[models.FileUploadResult], err error) {
	val := t.getBaseBody("file.upload")
	val["file_type"], val["file_name"] = fileType, fileName
	var req = t.buildUploadRequest(val, fileType, fileName)
	req.SetFileReader("file_content", fileName, imageFile)
	body = new(models.BaseResponse[models.FileUploadResult])
	err = t.reqPost(req, body)
	return
}

// DownloadFileBytes [fileNo]文件编号
func (t *Client) DownloadFileBytes(fileNo string) (bytes []byte, err error) {
	val := t.getBaseBody("file.download")
	val["file_no"] = fileNo
	var req = t.buildRequest(val, fileNo)
	response, err := t.reqPostDL(req)
	if err != nil {
		return nil, err
	}
	return response.Bytes(), nil
}

// DownloadFile [fileNo]文件编号
func (t *Client) DownloadFile(fileNo string, outputFilePath string) (err error) {
	val := t.getBaseBody("file.download")
	val["file_no"] = fileNo
	var req = t.buildRequest(val, fileNo).SetOutputFile(outputFilePath)
	response, err := t.reqPostDL(req)
	if err != nil {
		return err
	} else if response.IsErrorState() {
		return errors.New(response.Status)
	}
	return
}

// ImageRelation 图片关联业务接口
func (t *Client) ImageRelation(params models.ImageRelationParams) (body *models.BaseResponse[models.ImageRelationResult], err error) {
	val := t.getBaseBody("other.imagerelation")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr)
	body = new(models.BaseResponse[models.ImageRelationResult])
	err = t.reqPost(req, body)
	return
}

// DownloadOrderCheckFile 商户对账文件
func (t *Client) DownloadOrderCheckFile(params models.DownloadOrderCheckFileParams, outputFilePath string) (err error) {
	val := t.getBaseBody("order.check.download")
	jsonRaw, jsonStr, _ := common.JsonRow(params)
	val["params"] = jsonRaw
	var req = t.buildRequest(val, jsonStr).SetOutputFile(outputFilePath)
	response, err := t.reqPostDL(req)
	if err != nil {
		return err
	} else if response.IsErrorState() {
		return errors.New(response.Status)
	}
	return
}
