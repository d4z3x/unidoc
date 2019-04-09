/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package creator

import (
	"os"

	"github.com/unidoc/unidoc/pdf/model"
)

// Loads the template from path as a list of pages.
func loadPagesFromFile(f *os.File) ([]*model.PdfPage, error) {
	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		return nil, err
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return nil, err
	}

	// Load the pages.
	var pages []*model.PdfPage
	for i := 0; i < numPages; i++ {
		page, err := pdfReader.GetPage(i + 1)
		if err != nil {
			return nil, err
		}

		pages = append(pages, page)
	}

	return pages, nil
}
