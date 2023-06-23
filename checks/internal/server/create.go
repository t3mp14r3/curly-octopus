package server

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"github.com/signintech/gopdf"
	"go.uber.org/zap"

	"github.com/t3mp14r3/curly-octopus/checks/gen"
)

func (s *Service) Create(ctx context.Context, req *gen.CreateRequest) (*gen.CreateResponse, error) {
    pdf := gopdf.GoPdf{}
    pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595, H: 420}})

    pdf.AddPage()

    err := pdf.AddTTFFont("font", s.fontPath)

    if err != nil {
        s.logger.Error("failed to load the font", zap.Error(err))
        return nil, err
    }

    err = pdf.SetFont("font", "", 15)
    
    if err != nil {
        s.logger.Error("failed to set the font", zap.Error(err))
        return nil, err
    }

    template := pdf.ImportPage(s.templatePath, 1, "/MediaBox")

    pdf.UseImportedTemplate(template, 0, 0, 595, 420)
    
    pdf.SetXY(55, 90)
    pdf.Cell(nil, req.Barcode)

    pdf.SetFontSize(15)
    pdf.SetXY(55, 199)
    pdf.Cell(nil, req.Name)

    pdf.SetFontSize(15)
    pdf.SetXY(451, 317.5)
    pdf.Cell(nil, fmt.Sprintf("%d", req.Cost))

    bytes := pdf.GetBytesPdf()

    date := time.Now().Format("02.01.2006")

    filename := fmt.Sprintf("doc_%s_%s.pdf", req.Id, date)
    path := filepath.Join(s.storagePath, filename)

    pdf.WritePdf(path)
    err = pdf.Close()

    return &gen.CreateResponse{
        Filename: filename,
        Data: bytes,
    }, nil
}
