package shopeego

import (
	"io/fs"
	"io/ioutil"
)

func (s *ShopeeClient) LogisticsDownloadShippingDocument(saveFilePath string) func(req *LogisticsDownloadShippingDocumentRequest) (err error) {
	return func(req *LogisticsDownloadShippingDocumentRequest) (err error) {
		b, err := s.postDownloadFile("LogisticsDownloadShippingDocument", req)
		if err != nil {
			return nil
		}

		err = ioutil.WriteFile(saveFilePath, b, fs.ModePerm|fs.ModeAppend)

		return nil
	}
}
