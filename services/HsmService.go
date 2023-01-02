package services

import (
	"errors"

	"github.com/metalcorpe/payshield-rest-api/interfaces"
	"github.com/metalcorpe/payshield-rest-api/models"
)

type HsmService struct {
	interfaces.IHsmRepository
}

func (service *HsmService) NewVerifypinResponse(r models.PinVer) error {
	ec := service.DA(r)
	if ec != "00" {
		return errors.New(ec)
	}
	return nil
}
func (service *HsmService) NewVersionResponse() (models.VersionResponse, error) {
	resp := models.VersionResponse{}
	ec, lmk, firmware := service.NC()
	if ec != "00" {
		return resp, errors.New(ec)
	}
	resp.LmkCheck = lmk
	resp.FirmwareNumber = firmware
	return resp, nil
}
func (service *HsmService) NewGenerateKeyPairResponse(r models.GeneratePair) (models.GeneratePairResp, error) {
	resp, ec := service.EI(r)
	if ec != "00" {
		return models.GeneratePairResp{}, errors.New(ec)
	}
	return resp, nil
}

func (service *HsmService) NewMigrateResponse(r models.Migrate) (models.MigrateRes, error) {
	resp, ec := service.BW(r)
	if ec != "00" {
		return models.MigrateRes{}, errors.New(ec)
	}
	return resp, nil
}

func (service *HsmService) NewMigratePrivateResponse(r models.TranslatePrivate) (models.TranslatePrivateResp, error) {
	resp, ec := service.EM(r)
	if ec != "00" {
		return models.TranslatePrivateResp{}, errors.New(ec)
	}
	return resp, nil
}

func (service *HsmService) NewGenerateKeyResponse(r models.GenerateKey) (models.GenerateKeyResp, error) {
	resp, ec := service.A0(r)
	if ec != "00" {
		return models.GenerateKeyResp{}, errors.New(ec)
	}
	return resp, nil
}

func (service *HsmService) NewExportKeyResponse(r models.ExportKey) (models.ExportKeyResp, error) {
	resp, ec := service.A8(r)
	if ec != "00" {
		return models.ExportKeyResp{}, errors.New(ec)
	}
	return resp, nil
}

func (service *HsmService) ImportKeyRSAResponce(r models.ImportKeyOrDataUnderRSAPubKey) (models.ImportKeyOrDataUnderRSAPubKeyResp, error) {
	resp, ec := service.GI(r)
	if ec != "00" {
		return models.ImportKeyOrDataUnderRSAPubKeyResp{}, errors.New(ec)
	}
	return resp, nil
}
