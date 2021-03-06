package handler

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImage_rectangles(t *testing.T) {
	i := NewImage("testData/haarcascade_frontalface_alt.xml")
	src, err := ioutil.ReadFile("testData/dest.jpg")
	respWriter := httptest.NewRecorder()
	if assert.NoError(t, err) {
		err = i.rectangles(src, respWriter)
		if assert.NoError(t, err) {
			assert.NotNil(t, respWriter.Body.Bytes())
		}
	}
}

func TestImage_draw(t *testing.T) {
	i := NewImage("testData/haarcascade_frontalface_alt.xml")
	src, err := ioutil.ReadFile("testData/dest.jpg")
	respWriter := httptest.NewRecorder()
	if assert.NoError(t, err) {
		err = i.draw(src, respWriter)
		if assert.NoError(t, err) {
			assert.NotNil(t, respWriter.Body.Bytes())
		}
	}
}
