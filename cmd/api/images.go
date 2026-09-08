package main

import (
	"errors"
	"fmt"
	"image"
	_ "image/jpeg" // registers the JPEG decoder with image.DecodeConfig
	_ "image/png"  // registers the PNG decoder with image.DecodeConfig
	"io"
	"net/http"
	"os"

	"github.com/Kelvin-Justin-Gordon/test1-imagelab/internal/data"
	"github.com/Kelvin-Justin-Gordon/test1-imagelab/internal/validator"
)

// maxUploadBytes has a 10 MB limit, applied to the actual file size
const maxUploadBytes = 10 << 20 // 10 MB

// uploadFormField is the form field name the frontend must use
const uploadFormField = "image"

// createImageHandler validates the upload, determines its real type by attempting to decode it, stores the originial under a server-controlled filename, and creates the image record
func (app *application) createImageHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadBytes+1<<20)

	if err := r.ParseMultipartForm(maxUploadBytes + 1<<20); err != nil {
		var maxBytesErr *http.MaxBytesError
		if errors.As(err, &maxBytesErr) {
			app.requestEntityTooLargeResponse(w, r, maxUploadBytes)
			return
		}
		app.badRequestResponse(w, r, err)
		return
	}

	file, header, err := r.FormFile(uploadFormField)
	if err != nil {
		app.badRequestResponse(w, r, fmt.Errorf("must include a file in the %q field", uploadFormField))
		return
	}
	defer file.Close()

	if header.Size <= 0 {
		app.badRequestResponse(w, r, errors.New("uploaded file must not be empty"))
		return
	}
	if header.Size > maxUploadBytes {
		app.requestEntityTooLargeResponse(w, r, maxUploadBytes)
		return
	}

	//attempts to actually decode the image
	_, format, err := image.DecodeConfig(file)
	if err != nil || !validator.In(format, "jpeg", "png") {
		app.unsupportedMediaTypeResponse(w, r)
		return
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	ext := ".jpg"
	if format == "png" {
		ext = ".png"
	}

	storedFilename, sizeBytes, err := app.originals.SaveOriginal(file, ext)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	img := &data.Image{
		OriginalFilename: header.Filename,
		StoredFilename:   storedFilename,
		MediaType:        "image/" + format,
		SizeBytes:        sizeBytes,
	}
	if err := app.models.Images.Insert(img); err != nil {
		os.Remove(app.originals.Path(storedFilename))
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/images/%s", img.ID))
	if err := app.writeJSON(w, http.StatusCreated, envelope{"image": img}, headers); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
