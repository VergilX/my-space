package main

import (
	"errors"
	"io"
	"net/http"

	"github.com/VergilX/my-space/internal/response"
)

func (app *application) uploadFile(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value(userIDKey).(int64)
	if !ok {
		app.serverError(w, r, errors.New("type assertion error: user-id"))
		return
	}

	userId := int(id)

	app.lock.Lock()
	transfer, exists := app.transfers[userId]

	if !exists {
		// uploader is first
		app.logger.Info("Uploader: Listening...")

		// create new transfer
		transfer = &Transfer{
			reader:  r.Body,
			ready:   make(chan bool),
			done:    make(chan bool),
			errChan: make(chan error, 1),
		}

		app.transfers[userId] = transfer
		app.lock.Unlock()

		app.logger.Info("Uploader: Waiting for downloader")
		select {
		case <-transfer.ready:
			// writer is ready
			app.logger.Info("Uploader: Downloader ready signal received")

		case err := <-transfer.errChan:
			app.serverError(w, r, err)
			return
		}
	} else {
		// downloader was first
		app.logger.Info("Uploader: Connecting waiting downloader")
		transfer.reader = r.Body
		app.lock.Unlock()

		// tell the downloader that the uploader is ready
		transfer.ready <- true

	}

	// wait for transfer to finish
	select {
	case <-transfer.done:
		app.logger.Info("Uploader: Transfer complete")
		response.JSON(w, http.StatusOK, envelope{"message": "transfer complete"})

	case err := <-transfer.errChan:
		app.serverError(w, r, err)
	}

}

func (app *application) downloadFile(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value(userIDKey).(int64)
	if !ok {
		app.serverError(w, r, errors.New("type assertion error: user-id"))
		return
	}

	userId := int(id)

	app.lock.Lock()
	transfer, exists := app.transfers[userId]
	if !exists {
		// downloader is first

		transfer = &Transfer{
			writer:  w,
			ready:   make(chan bool),
			done:    make(chan bool),
			errChan: make(chan error, 1),
		}

		app.transfers[userId] = transfer
		app.lock.Unlock()

		app.logger.Info("Downloader: Waiting for uploader...")
		select {
		case <-transfer.ready:
			// writer is ready
			app.logger.Info("Downloader: Uploader ready signal received")

		case err := <-transfer.errChan:
			app.serverError(w, r, err)
			return

		case <-r.Context().Done():
			app.serverError(w, r, errors.New("downloader disconnected while waiting"))
			return
		}
	} else {
		// uploader was first
		app.logger.Info("Downloader: Connecting to waiting uploader")
		transfer.writer = w
		app.lock.Unlock()

		transfer.ready <- true
	}

	// force header
	// w.Header().Set("Content-Type", "application/octet-stream")

	app.logger.Info("Downloader: Starting stream...")
	_, err := io.Copy(transfer.writer, transfer.reader)

	// Clean up the map *before* handling errors, as the transfer is over
	app.lock.Lock()
	delete(app.transfers, userId)
	app.lock.Unlock()

	if err != nil {
		// Send the error to the uploader
		app.serverError(w, r, err)
		transfer.errChan <- err
		return
	}

	// Success!
	app.logger.Info("Downloader: Stream finished successfully")
	transfer.reader.Close()
	close(transfer.done) // Signal success to the uploader
}
