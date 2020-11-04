package service

import (
	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/domain/contract"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

type cloudVisionService struct {
	svc *Service
}

//newCloudVisionService return a new instance of the service
func newCloudVisionService(svc *Service) contract.CloudVisionService {
	return &cloudVisionService{
		svc: svc,
	}
}

func (s *cloudVisionService) SendImageToCloudVision(filePath string) (text string, err resterrors.RestErr) {

	text = "ALBERT EINSTEIN\nHOSPITAL ISRAELITA\nReceituário\nPaciente: Giovanni Coutinho Bernini\nEndereço: RUA INHAMBU 635, APTO 111 -SAO PAULO, SP 04520012\nUso interno\n1\nDIPIRONA SODICA 1000 MG COMPRIMIDOS\n1 comprimido(s), via oral, 4x/dia, se dor ou febre, Até de 6/6 horas\n2\nPARACETAMOL 750 MG CP\n1 comprimido(s), via oral, 4x/dia, se dor ou febre, Até de 6/6 horas\n3\nALLEGRA 180 MG COMPRIMIDOS REV\n1 comprimido(s), via oral, 1x/dia, Durante 5 dias\nUso nasal\n4\nMARESIS AEROSOL SPRAY NASAL 0,9%\n1 jato, via nasal, 6/6h\nDr. Henrique Vicente Haussauer Jr.\nMaeco\nCRM-S 168409\nSão Paulo-SP ,19/7/2020\nHenrique Vicente Haussauer Junior - 168409CRMSP - 168409CRMSP\nIbirapuera\nAv. Republica do Libano, 501, Ibirapuera, São Paulo-04501-000 Telefone 1121511233\n"
	return text, nil
}
