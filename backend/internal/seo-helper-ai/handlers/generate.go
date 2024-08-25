package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Fibig/seo-helper-ai/internal/seo-helper-ai/templates/views"
	"github.com/Fibig/seo-helper-ai/internal/seo-helper-ai/utils"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
)

type APIError struct {
	StatusCode int
	Message    error
}

type OllamaResponse struct {
	Model       string `json:"model"`
	Created_at  string `json:"created_at"`
	Response    string `json:"response"`
	Done        bool   `json:"done"`
	Done_reason string `json:"done_reason"`
	Context     []int  `json:"context"`
}

func (err *APIError) Error() string {
	return fmt.Sprintf("status %d: err %v", err.StatusCode, err.Message)
}

func Generate() gin.HandlerFunc {
	return utils.RenderTemplComponent(views.Generate())
}

func APIGenerate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// prepare Form Data
		err := ctx.Request.ParseForm()
		if err != nil {
			ctx.String(http.StatusBadRequest, "%s", err)
			return
		}

		// reading Form Data Values
		prompt := ctx.Request.FormValue("prompt")

		// generate AI Response
		response, err := seoHelperAIGenerate(prompt)
		if err != nil {
			apiError := err.(*APIError)
			ctx.String(apiError.StatusCode, "%s", apiError.Message)
			return
		}

		ctx.String(http.StatusOK, "%s", response.Response)
	}
}

func seoHelperAIGenerate(prompt string) (*OllamaResponse, error) {
	// Example Prompt:
	// Generate website content for my startup called b4c connect, which is an online portal for the customers of the company b4c consulting. At b4c connect users can upload their files and enter their personal information to get a loan more quickly.

	// build a JSON Body
	body := map[string]any{
		"model":  "seo-helper-ai:latest",
		"prompt": prompt,
		"stream": false,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, &APIError{http.StatusInternalServerError, err}
	}

	// convert JSON to bytes
	bytesReader := bytes.NewReader(jsonBody)

	// requesting from Ollama API
	response, err := http.Post("http://localhost:11434/api/generate", mimetype.Detect(jsonBody).String(), bytesReader)
	if err != nil {
		if response != nil {
			return nil, &APIError{http.StatusBadRequest, err}
		}

		return nil, &APIError{http.StatusServiceUnavailable, err}
	}
	defer response.Body.Close()

	// reading Response Body
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, &APIError{http.StatusInternalServerError, err}
	}

	// convert AI Response to Ollama Response
	ollamaResponse := OllamaResponse{}
	err = json.Unmarshal(responseBody, &ollamaResponse)
	if err != nil {
		return nil, &APIError{http.StatusInternalServerError, err}
	}

	return &ollamaResponse, nil
}
