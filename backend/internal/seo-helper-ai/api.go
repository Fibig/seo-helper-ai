package seohelperai

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func chatWithAI() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := map[string]string{
			"model":  "seo-helper-ai:latest",
			"prompt": "Generate website content for my startup called b4c connect, which is an online portal for the customers of the company b4c consulting. At b4c connect users can upload their files and enter their personal information to get a loan more quickly.",
			"format": "json",
			"stream": "false",
		}

		jsonBody, err := json.Marshal(body)
		if err != nil {
			ctx.String(200, "%s", err)
			return
		}

		bytesReader := bytes.NewReader(jsonBody)
		resp, err := http.Post("http://localhost:11434/chat", "application/json", bytesReader)
		if err != nil {
			ctx.String(200, "%s", err)
			return
		}

		var target interface{}
		json.NewDecoder(resp.Body).Decode(target)
		defer resp.Body.Close()

		ctx.JSON(http.StatusOK, gin.H{
			"response": target,
		})
	}
}
