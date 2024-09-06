package controller

import (
	"errors"
	agent "github.com/cloud-barista/cm-honeybee/agent/common"
	"github.com/cloud-barista/cm-honeybee/server/dao"
	"github.com/cloud-barista/cm-honeybee/server/lib/config"
	"github.com/cloud-barista/cm-honeybee/server/pkg/api/rest/common"
	"github.com/cloud-barista/cm-honeybee/server/pkg/api/rest/model"
	"github.com/jollaman999/utils/logger"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"time"
)

func doImportInfra(connID string) (*model.SavedInfraInfo, error) {
	connectionInfo, err := dao.ConnectionInfoGet(connID)
	if err != nil {
		return nil, err
	}

	oldSavedInfraInfo, _ := dao.SavedInfraInfoGet(connectionInfo.ID)

	if oldSavedInfraInfo == nil {
		savedInfraInfo := new(model.SavedInfraInfo)
		savedInfraInfo.ConnectionID = connectionInfo.ID
		savedInfraInfo.InfraData = ""
		savedInfraInfo.Status = "importing"
		savedInfraInfo.SavedTime = time.Now()
		savedInfraInfo, err = dao.SavedInfraInfoRegister(savedInfraInfo)
		if err != nil {
			errMsg := "Error occurred while getting infra information." +
				" (ConnectionID = " + connectionInfo.ID + ")"
			logger.Println(logger.ERROR, false, errMsg)
			return nil, errors.New(errMsg)
		}
		oldSavedInfraInfo = savedInfraInfo
	}

	data, err := common.GetHTTPRequest("http://" + connectionInfo.IPAddress + ":" + config.CMHoneybeeConfig.CMHoneybee.Agent.Port + "/" + strings.ToLower(agent.ShortModuleName) + "/infra")
	if err != nil {
		oldSavedInfraInfo.Status = "failed"
		_ = dao.SavedInfraInfoUpdate(oldSavedInfraInfo)
		errMsg := "Error occurred while getting infra information." +
			" (ConnectionID = " + connectionInfo.ID + ")"
		logger.Println(logger.ERROR, false, errMsg)
		return nil, errors.New(errMsg)
	}

	oldSavedInfraInfo.InfraData = string(data)
	oldSavedInfraInfo.Status = "success"
	oldSavedInfraInfo.SavedTime = time.Now()
	err = dao.SavedInfraInfoUpdate(oldSavedInfraInfo)
	if err != nil {
		errMsg := "Error occurred while saving the infra information." +
			" (ConnectionID = " + connectionInfo.ID + ")"
		logger.Println(logger.ERROR, false, errMsg)
		return nil, errors.New(errMsg)
	}

	return oldSavedInfraInfo, nil
}

func doImportSoftware(connID string) (*model.SavedSoftwareInfo, error) {
	connectionInfo, err := dao.ConnectionInfoGet(connID)
	if err != nil {
		return nil, err
	}

	oldSavedSoftwareInfo, _ := dao.SavedSoftwareInfoGet(connectionInfo.ID)

	if oldSavedSoftwareInfo == nil {
		savedSoftwareInfo := new(model.SavedSoftwareInfo)
		savedSoftwareInfo.ConnectionID = connectionInfo.ID
		savedSoftwareInfo.SoftwareData = ""
		savedSoftwareInfo.Status = "importing"
		savedSoftwareInfo.SavedTime = time.Now()
		savedSoftwareInfo, err = dao.SavedSoftwareInfoRegister(savedSoftwareInfo)
		if err != nil {
			errMsg := "Error occurred while getting software information." +
				" (ConnectionID = " + connectionInfo.ID + ")"
			logger.Println(logger.ERROR, false, errMsg)
			return nil, errors.New(errMsg)
		}
		oldSavedSoftwareInfo = savedSoftwareInfo
	}

	data, err := common.GetHTTPRequest("http://" + connectionInfo.IPAddress + ":" + config.CMHoneybeeConfig.CMHoneybee.Agent.Port + "/" + strings.ToLower(agent.ShortModuleName) + "/software")
	if err != nil {
		oldSavedSoftwareInfo.Status = "failed"
		_ = dao.SavedSoftwareInfoUpdate(oldSavedSoftwareInfo)
		errMsg := "Error occurred while getting software information." +
			" (ConnectionID = " + connectionInfo.ID + ")"
		logger.Println(logger.ERROR, false, errMsg)
		return nil, errors.New(errMsg)
	}

	oldSavedSoftwareInfo.SoftwareData = string(data)
	oldSavedSoftwareInfo.Status = "success"
	oldSavedSoftwareInfo.SavedTime = time.Now()
	err = dao.SavedSoftwareInfoUpdate(oldSavedSoftwareInfo)
	if err != nil {
		errMsg := "Error occurred while saving the software information." +
			" (ConnectionID = " + connectionInfo.ID + ")"
		logger.Println(logger.ERROR, false, errMsg)
		return nil, errors.New(errMsg)
	}

	return oldSavedSoftwareInfo, nil
}

func doImportKubernetes(connID string) (*model.SavedKubernetesInfo, error) {
	connectionInfo, err := dao.ConnectionInfoGet(connID)
	if err != nil {
		return nil, err
	}

	oldSavedKubernetesInfo, _ := dao.SavedKubernetesInfoGet(connectionInfo.ID)

	if oldSavedKubernetesInfo == nil {
		savedKubernetesInfo := new(model.SavedKubernetesInfo)
		savedKubernetesInfo.ConnectionID = connectionInfo.ID
		savedKubernetesInfo.KubernetesData = ""
		savedKubernetesInfo.Status = "importing"
		savedKubernetesInfo.SavedTime = time.Now()
		savedKubernetesInfo, err = dao.SavedKubernetesInfoRegister(savedKubernetesInfo)
		if err != nil {
			errMsg := "Error occurred while getting kubernetes information." +
				" (ConnectionID = " + connectionInfo.ID + ")"
			logger.Println(logger.ERROR, false, errMsg)
			return nil, errors.New(errMsg)
		}
		oldSavedKubernetesInfo = savedKubernetesInfo
	}

	logger.Println(logger.INFO, false, "http://"+connectionInfo.IPAddress+":"+config.CMHoneybeeConfig.CMHoneybee.Agent.Port+"/"+strings.ToLower(agent.ShortModuleName)+"/kubernetes")
	data, err := common.GetHTTPRequest("http://" + connectionInfo.IPAddress + ":" + config.CMHoneybeeConfig.CMHoneybee.Agent.Port + "/" + strings.ToLower(agent.ShortModuleName) + "/kubernetes")
	if err != nil {
		oldSavedKubernetesInfo.Status = "failed"
		_ = dao.SavedKubernetesInfoUpdate(oldSavedKubernetesInfo)
		errMsg := "Error occurred while getting software information." +
			" (ConnectionID = " + connectionInfo.ID + ")"
		logger.Println(logger.ERROR, false, errMsg)
		return nil, errors.New(errMsg)
	}

	oldSavedKubernetesInfo.KubernetesData = string(data)
	oldSavedKubernetesInfo.Status = "success"
	oldSavedKubernetesInfo.SavedTime = time.Now()
	err = dao.SavedKubernetesInfoUpdate(oldSavedKubernetesInfo)
	if err != nil {
		errMsg := "Error occurred while saving the kubernetes information." +
			" (ConnectionID = " + connectionInfo.ID + ")"
		logger.Println(logger.ERROR, false, errMsg)
		return nil, errors.New(errMsg)
	}

	return oldSavedKubernetesInfo, nil
}

// ImportInfra godoc
//
// @Summary		Import Infra
// @Description	Import the infra information.
// @Tags		[Import] Import source info
// @Accept		json
// @Produce		json
// @Param		sgId path string true "ID of the source group."
// @Param		connId path string true "ID of the connection info."
// @Success		200	{object}	model.SavedInfraInfo	"Successfully saved the infra information"
// @Failure		400	{object}	common.ErrorResponse	"Sent bad request."
// @Failure		500	{object}	common.ErrorResponse	"Failed to save the infra information"
// @Router		/source_group/{sgId}/connection_info/{connId}/import/infra [post]
func ImportInfra(c echo.Context) error {
	sgID := c.Param("sgId")
	if sgID == "" {
		return common.ReturnErrorMsg(c, "Please provide the sgId.")
	}

	connID := c.Param("connId")
	if connID == "" {
		return common.ReturnErrorMsg(c, "Please provide the connId.")
	}

	_, err := dao.SourceGroupGet(sgID)
	if err != nil {
		return common.ReturnErrorMsg(c, err.Error())
	}

	savedInfraInfo, err := doImportInfra(connID)
	if err != nil {
		return common.ReturnErrorMsg(c, err.Error())
	}

	return c.JSONPretty(http.StatusOK, savedInfraInfo, " ")
}

// ImportInfraSourceGroup godoc
//
// @Summary		Import Infra Source Group
// @Description	Import infra information for all connections in the source group.
// @Tags		[Import] Import source info
// @Accept		json
// @Produce		json
// @Param		sgId path string true "ID of the source group."
// @Success		200	{object}	[]model.SavedInfraInfo	"Successfully saved the infra information"
// @Failure		400	{object}	common.ErrorResponse	"Sent bad request."
// @Failure		500	{object}	common.ErrorResponse	"Failed to save the infra information"
// @Router		/source_group/{sgId}/import/infra [post]
func ImportInfraSourceGroup(c echo.Context) error {
	sgID := c.Param("sgId")
	if sgID == "" {
		return common.ReturnErrorMsg(c, "Please provide the sgId.")
	}

	_, err := dao.SourceGroupGet(sgID)
	if err != nil {
		return common.ReturnErrorMsg(c, err.Error())
	}

	list, err := dao.ConnectionInfoGetList(&model.ConnectionInfo{SourceGroupID: sgID}, 0, 0)
	if err != nil {
		return common.ReturnErrorMsg(c, err.Error())
	}

	var savedInfraInfoList []model.SavedInfraInfo

	for _, conn := range *list {
		savedInfraInfo, _ := doImportInfra(conn.ID)
		savedInfraInfoList = append(savedInfraInfoList, *savedInfraInfo)
	}

	return c.JSONPretty(http.StatusOK, savedInfraInfoList, " ")
}

// ImportSoftware godoc
//
// @Summary		Import software
// @Description	Import the software information.
// @Tags		[Import] Import source info
// @Accept		json
// @Produce		json
// @Param		sgId path string true "ID of the source group."
// @Param		connId path string true "ID of the connection info."
// @Success		200	{object}	model.SavedSoftwareInfo	"Successfully saved the software information"
// @Failure		400	{object}	common.ErrorResponse	"Sent bad request."
// @Failure		500	{object}	common.ErrorResponse	"Failed to save the software information"
// @Router		/source_group/{sgId}/connection_info/{connId}/import/software [post]
func ImportSoftware(c echo.Context) error {
	sgID := c.Param("sgId")
	if sgID == "" {
		return common.ReturnErrorMsg(c, "Please provide the sgId.")
	}

	connID := c.Param("connId")
	if connID == "" {
		return common.ReturnErrorMsg(c, "Please provide the connId.")
	}

	_, err := dao.SourceGroupGet(sgID)
	if err != nil {
		return common.ReturnErrorMsg(c, err.Error())
	}

	savedSoftwareInfo, err := doImportSoftware(connID)
	if err != nil {
		return common.ReturnErrorMsg(c, err.Error())
	}

	return c.JSONPretty(http.StatusOK, savedSoftwareInfo, " ")
}

// ImportSoftwareSourceGroup godoc
//
// @Summary		Import Software Source Group
// @Description	Import software information for all connections in the source group.
// @Tags		[Import] Import source info
// @Accept		json
// @Produce		json
// @Param		sgId path string true "ID of the source group."
// @Success		200	{object}	[]model.SavedSoftwareInfo	"Successfully saved the software information"
// @Failure		400	{object}	common.ErrorResponse	"Sent bad request."
// @Failure		500	{object}	common.ErrorResponse	"Failed to save the software information"
// @Router		/source_group/{sgId}/import/software [post]
func ImportSoftwareSourceGroup(c echo.Context) error {
	sgID := c.Param("sgId")
	if sgID == "" {
		return common.ReturnErrorMsg(c, "Please provide the sgId.")
	}

	_, err := dao.SourceGroupGet(sgID)
	if err != nil {
		return common.ReturnErrorMsg(c, err.Error())
	}

	list, err := dao.ConnectionInfoGetList(&model.ConnectionInfo{SourceGroupID: sgID}, 0, 0)
	if err != nil {
		return common.ReturnErrorMsg(c, err.Error())
	}

	var savedSoftwareInfoList []model.SavedSoftwareInfo

	for _, conn := range *list {
		savedSoftwareInfo, _ := doImportSoftware(conn.ID)
		savedSoftwareInfoList = append(savedSoftwareInfoList, *savedSoftwareInfo)
	}

	return c.JSONPretty(http.StatusOK, savedSoftwareInfoList, " ")
}

// ImportKubernetes godoc
//
// @Summary		Import kubernetes
// @Description	Import the kubernetes information.
// @Tags		[Import] Import source info
// @Accept		json
// @Produce		json
// @Param		sgId path string true "ID of the source group."
// @Param		connId path string true "ID of the connection info."
// @Success		200	{object}	model.SavedSoftwareInfo	"Successfully saved the kubernetes information"
// @Failure		400	{object}	common.ErrorResponse	"Sent bad request."
// @Failure		500	{object}	common.ErrorResponse	"Failed to save the kubernetes information"
// @Router		/source_group/{sgId}/connection_info/{connId}/import/kubernetes [post]
func ImportKubernetes(c echo.Context) error {
	sgID := c.Param("sgId")
	if sgID == "" {
		return common.ReturnErrorMsg(c, "Please provide the sgId.")
	}

	connID := c.Param("connId")
	if connID == "" {
		return common.ReturnErrorMsg(c, "Please provide the connId.")
	}

	_, err := dao.SourceGroupGet(sgID)
	if err != nil {
		return common.ReturnErrorMsg(c, err.Error())
	}

	savedKubernetesInfo, err := doImportKubernetes(connID)
	if err != nil {
		return common.ReturnErrorMsg(c, err.Error())
	}

	return c.JSONPretty(http.StatusOK, savedKubernetesInfo, " ")
}

// ImportKubernetesSourceGroup godoc
//
// @Summary		Import Kubernetes Source Group
// @Description	Import kubernetes information for all connections in the source group.
// @Tags		[Import] Import source info
// @Accept		json
// @Produce		json
// @Param		sgId path string true "ID of the source group."
// @Success		200	{object}	[]model.SavedKubernetesInfo	"Successfully saved the kubernetes information"
// @Failure		400	{object}	common.ErrorResponse	"Sent bad request."
// @Failure		500	{object}	common.ErrorResponse	"Failed to save the kubernetes information"
// @Router		/source_group/{sgId}/import/kubernetes [post]
func ImportKubernetesSourceGroup(c echo.Context) error {
	sgID := c.Param("sgId")
	if sgID == "" {
		return common.ReturnErrorMsg(c, "Please provide the sgId.")
	}

	_, err := dao.SourceGroupGet(sgID)
	if err != nil {
		return common.ReturnErrorMsg(c, err.Error())
	}

	list, err := dao.ConnectionInfoGetList(&model.ConnectionInfo{SourceGroupID: sgID}, 0, 0)
	if err != nil {
		return common.ReturnErrorMsg(c, err.Error())
	}

	var savedKubernetesInfoList []model.SavedKubernetesInfo

	for _, conn := range *list {
		savedKubernetesInfo, _ := doImportKubernetes(conn.ID)
		savedKubernetesInfoList = append(savedKubernetesInfoList, *savedKubernetesInfo)
	}

	return c.JSONPretty(http.StatusOK, savedKubernetesInfoList, " ")
}
