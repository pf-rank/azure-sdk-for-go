// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azopenai_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

func ExampleClient_GetChatCompletions() {
	azureOpenAIKey := os.Getenv("AOAI_CHAT_COMPLETIONS_API_KEY")
	modelDeploymentID := os.Getenv("AOAI_CHAT_COMPLETIONS_MODEL")

	// Ex: "https://<your-azure-openai-host>.openai.azure.com"
	azureOpenAIEndpoint := os.Getenv("AOAI_CHAT_COMPLETIONS_ENDPOINT")

	if azureOpenAIKey == "" || modelDeploymentID == "" || azureOpenAIEndpoint == "" {
		fmt.Fprintf(os.Stderr, "Skipping example, environment variables missing\n")
		return
	}

	keyCredential := azcore.NewKeyCredential(azureOpenAIKey)

	// In Azure OpenAI you must deploy a model before you can use it in your client. For more information
	// see here: https://learn.microsoft.com/azure/cognitive-services/openai/how-to/create-resource
	client, err := azopenai.NewClientWithKeyCredential(azureOpenAIEndpoint, keyCredential, nil)

	if err != nil {
		// TODO: Update the following line with your application specific error handling logic
		log.Printf("ERROR: %s", err)
		return
	}

	// This is a conversation in progress.
	// NOTE: all messages, regardless of role, count against token usage for this API.
	messages := []azopenai.ChatRequestMessageClassification{
		// You set the tone and rules of the conversation with a prompt as the system role.
		&azopenai.ChatRequestSystemMessage{Content: azopenai.NewChatRequestSystemMessageContent("You are a helpful assistant. You will talk like a pirate.")},

		// The user asks a question
		&azopenai.ChatRequestUserMessage{Content: azopenai.NewChatRequestUserMessageContent("Can you help me?")},

		// The reply would come back from the ChatGPT. You'd add it to the conversation so we can maintain context.
		&azopenai.ChatRequestAssistantMessage{Content: azopenai.NewChatRequestAssistantMessageContent("Arrrr! Of course, me hearty! What can I do for ye?")},

		// The user answers the question based on the latest reply.
		&azopenai.ChatRequestUserMessage{Content: azopenai.NewChatRequestUserMessageContent("What's the best way to train a parrot?")},

		// from here you'd keep iterating, sending responses back from ChatGPT
	}

	gotReply := false

	resp, err := client.GetChatCompletions(context.TODO(), azopenai.ChatCompletionsOptions{
		// This is a conversation in progress.
		// NOTE: all messages count against token usage for this API.
		Messages:       messages,
		DeploymentName: &modelDeploymentID,
	}, nil)

	if err != nil {
		// TODO: Update the following line with your application specific error handling logic
		log.Printf("ERROR: %s", err)
		return
	}

	for _, choice := range resp.Choices {
		gotReply = true

		if choice.ContentFilterResults != nil {
			fmt.Fprintf(os.Stderr, "Content filter results\n")

			if choice.ContentFilterResults.Error != nil {
				fmt.Fprintf(os.Stderr, "  Error:%v\n", choice.ContentFilterResults.Error)
			}

			fmt.Fprintf(os.Stderr, "  Hate: sev: %v, filtered: %v\n", *choice.ContentFilterResults.Hate.Severity, *choice.ContentFilterResults.Hate.Filtered)
			fmt.Fprintf(os.Stderr, "  SelfHarm: sev: %v, filtered: %v\n", *choice.ContentFilterResults.SelfHarm.Severity, *choice.ContentFilterResults.SelfHarm.Filtered)
			fmt.Fprintf(os.Stderr, "  Sexual: sev: %v, filtered: %v\n", *choice.ContentFilterResults.Sexual.Severity, *choice.ContentFilterResults.Sexual.Filtered)
			fmt.Fprintf(os.Stderr, "  Violence: sev: %v, filtered: %v\n", *choice.ContentFilterResults.Violence.Severity, *choice.ContentFilterResults.Violence.Filtered)
		}

		if choice.Message != nil && choice.Message.Content != nil {
			fmt.Fprintf(os.Stderr, "Content[%d]: %s\n", *choice.Index, *choice.Message.Content)
		}

		if choice.FinishReason != nil {
			// this choice's conversation is complete.
			fmt.Fprintf(os.Stderr, "Finish reason[%d]: %s\n", *choice.Index, *choice.FinishReason)
		}
	}

	if gotReply {
		fmt.Fprintf(os.Stderr, "Got chat completions reply\n")
	}

	// Output:
}

func ExampleClient_GetChatCompletions_functions() {
	azureOpenAIKey := os.Getenv("AOAI_CHAT_COMPLETIONS_API_KEY")
	modelDeploymentID := os.Getenv("AOAI_CHAT_COMPLETIONS_MODEL")

	// Ex: "https://<your-azure-openai-host>.openai.azure.com"
	azureOpenAIEndpoint := os.Getenv("AOAI_CHAT_COMPLETIONS_ENDPOINT")

	if azureOpenAIKey == "" || modelDeploymentID == "" || azureOpenAIEndpoint == "" {
		fmt.Fprintf(os.Stderr, "Skipping example, environment variables missing\n")
		return
	}

	keyCredential := azcore.NewKeyCredential(azureOpenAIKey)

	// In Azure OpenAI you must deploy a model before you can use it in your client. For more information
	// see here: https://learn.microsoft.com/azure/cognitive-services/openai/how-to/create-resource
	client, err := azopenai.NewClientWithKeyCredential(azureOpenAIEndpoint, keyCredential, nil)

	if err != nil {
		// TODO: Update the following line with your application specific error handling logic
		log.Printf("ERROR: %s", err)
		return
	}

	jsonBytes, err := json.Marshal(map[string]any{
		"required": []string{"location"},
		"type":     "object",
		"properties": map[string]any{
			"location": map[string]any{
				"type":        "string",
				"description": "The city and state, e.g. San Francisco, CA",
			},
			"unit": map[string]any{
				"type": "string",
				"enum": []string{"celsius", "fahrenheit"},
			},
		},
	})

	if err != nil {
		panic(err)
	}

	funcDef := &azopenai.ChatCompletionsFunctionToolDefinitionFunction{
		Name:        to.Ptr("get_current_weather"),
		Description: to.Ptr("Get the current weather in a given location"),
		Parameters:  jsonBytes,
	}

	resp, err := client.GetChatCompletions(context.TODO(), azopenai.ChatCompletionsOptions{
		DeploymentName: &modelDeploymentID,
		Messages: []azopenai.ChatRequestMessageClassification{
			&azopenai.ChatRequestUserMessage{
				Content: azopenai.NewChatRequestUserMessageContent("What's the weather like in Boston, MA, in celsius?"),
			},
		},
		Tools: []azopenai.ChatCompletionsToolDefinitionClassification{
			&azopenai.ChatCompletionsFunctionToolDefinition{
				Function: funcDef,
			},
		},
		Temperature: to.Ptr[float32](0.0),
	}, nil)

	if err != nil {
		// TODO: Update the following line with your application specific error handling logic
		log.Printf("ERROR: %s", err)
		return
	}

	funcCall := resp.Choices[0].Message.ToolCalls[0].(*azopenai.ChatCompletionsFunctionToolCall).Function

	// This is the function name we gave in the call to GetCompletions
	// Prints: Function name: "get_current_weather"
	fmt.Fprintf(os.Stderr, "Function name: %q\n", *funcCall.Name)

	// The arguments for your function come back as a JSON string
	var funcParams *struct {
		Location string `json:"location"`
		Unit     string `json:"unit"`
	}
	err = json.Unmarshal([]byte(*funcCall.Arguments), &funcParams)

	if err != nil {
		// TODO: Update the following line with your application specific error handling logic
		log.Printf("ERROR: %s", err)
		return
	}

	// Prints:
	// Parameters: azopenai_test.location{Location:"Boston, MA", Unit:"celsius"}
	fmt.Fprintf(os.Stderr, "Parameters: %#v\n", *funcParams)

	// Output:
}

func ExampleClient_GetChatCompletions_legacyFunctions() {
	azureOpenAIKey := os.Getenv("AOAI_CHAT_COMPLETIONS_MODEL_LEGACY_FUNCTIONS_API_KEY")
	modelDeploymentID := os.Getenv("AOAI_CHAT_COMPLETIONS_MODEL_LEGACY_FUNCTIONS_MODEL")

	// Ex: "https://<your-azure-openai-host>.openai.azure.com"
	azureOpenAIEndpoint := os.Getenv("AOAI_CHAT_COMPLETIONS_MODEL_LEGACY_FUNCTIONS_ENDPOINT")

	if azureOpenAIKey == "" || modelDeploymentID == "" || azureOpenAIEndpoint == "" {
		fmt.Fprintf(os.Stderr, "Skipping example, environment variables missing\n")
		return
	}

	keyCredential := azcore.NewKeyCredential(azureOpenAIKey)

	// In Azure OpenAI you must deploy a model before you can use it in your client. For more information
	// see here: https://learn.microsoft.com/azure/cognitive-services/openai/how-to/create-resource
	client, err := azopenai.NewClientWithKeyCredential(azureOpenAIEndpoint, keyCredential, nil)

	if err != nil {
		// TODO: Update the following line with your application specific error handling logic
		log.Printf("ERROR: %s", err)
		return
	}

	parametersJSON, err := json.Marshal(map[string]any{
		"required": []string{"location"},
		"type":     "object",
		"properties": map[string]any{
			"location": map[string]any{
				"type":        "string",
				"description": "The city and state, e.g. San Francisco, CA",
			},
			"unit": map[string]any{
				"type": "string",
				"enum": []string{"celsius", "fahrenheit"},
			},
		},
	})

	if err != nil {
		// TODO: Update the following line with your application specific error handling logic
		log.Printf("ERROR: %s", err)
		return
	}

	resp, err := client.GetChatCompletions(context.TODO(), azopenai.ChatCompletionsOptions{
		DeploymentName: &modelDeploymentID,
		Messages: []azopenai.ChatRequestMessageClassification{
			&azopenai.ChatRequestUserMessage{
				Content: azopenai.NewChatRequestUserMessageContent("What's the weather like in Boston, MA, in celsius?"),
			},
		},
		FunctionCall: &azopenai.ChatCompletionsOptionsFunctionCall{
			Value: to.Ptr("auto"),
		},
		Functions: []azopenai.FunctionDefinition{
			{
				Name:        to.Ptr("get_current_weather"),
				Description: to.Ptr("Get the current weather in a given location"),

				Parameters: parametersJSON,
			},
		},
		Temperature: to.Ptr[float32](0.0),
	}, nil)

	if err != nil {
		// TODO: Update the following line with your application specific error handling logic
		log.Printf("ERROR: %s", err)
		return
	}

	funcCall := resp.Choices[0].Message.FunctionCall

	// This is the function name we gave in the call to GetCompletions
	// Prints: Function name: "get_current_weather"
	fmt.Fprintf(os.Stderr, "Function name: %q\n", *funcCall.Name)

	// The arguments for your function come back as a JSON string
	var funcParams *struct {
		Location string `json:"location"`
		Unit     string `json:"unit"`
	}
	err = json.Unmarshal([]byte(*funcCall.Arguments), &funcParams)

	if err != nil {
		// TODO: Update the following line with your application specific error handling logic
		log.Printf("ERROR: %s", err)
		return
	}

	// Prints:
	// Parameters: azopenai_test.location{Location:"Boston, MA", Unit:"celsius"}
	fmt.Fprintf(os.Stderr, "Parameters: %#v\n", *funcParams)

	// Output:
}

func ExampleClient_GetChatCompletionsStream() {
	azureOpenAIKey := os.Getenv("AOAI_CHAT_COMPLETIONS_API_KEY")
	modelDeploymentID := os.Getenv("AOAI_CHAT_COMPLETIONS_MODEL")

	// Ex: "https://<your-azure-openai-host>.openai.azure.com"
	azureOpenAIEndpoint := os.Getenv("AOAI_CHAT_COMPLETIONS_ENDPOINT")

	if azureOpenAIKey == "" || modelDeploymentID == "" || azureOpenAIEndpoint == "" {
		fmt.Fprintf(os.Stderr, "Skipping example, environment variables missing\n")
		return
	}

	keyCredential := azcore.NewKeyCredential(azureOpenAIKey)

	// In Azure OpenAI you must deploy a model before you can use it in your client. For more information
	// see here: https://learn.microsoft.com/azure/cognitive-services/openai/how-to/create-resource
	client, err := azopenai.NewClientWithKeyCredential(azureOpenAIEndpoint, keyCredential, nil)

	if err != nil {
		// TODO: Update the following line with your application specific error handling logic
		log.Printf("ERROR: %s", err)
		return
	}

	// This is a conversation in progress.
	// NOTE: all messages, regardless of role, count against token usage for this API.
	messages := []azopenai.ChatRequestMessageClassification{
		// You set the tone and rules of the conversation with a prompt as the system role.
		&azopenai.ChatRequestSystemMessage{Content: azopenai.NewChatRequestSystemMessageContent("You are a helpful assistant. You will talk like a pirate and limit your responses to 20 words or less.")},

		// The user asks a question
		&azopenai.ChatRequestUserMessage{Content: azopenai.NewChatRequestUserMessageContent("Can you help me?")},

		// The reply would come back from the ChatGPT. You'd add it to the conversation so we can maintain context.
		&azopenai.ChatRequestAssistantMessage{Content: azopenai.NewChatRequestAssistantMessageContent("Arrrr! Of course, me hearty! What can I do for ye?")},

		// The user answers the question based on the latest reply.
		&azopenai.ChatRequestUserMessage{Content: azopenai.NewChatRequestUserMessageContent("What's the best way to train a parrot?")},

		// from here you'd keep iterating, sending responses back from ChatGPT
	}

	resp, err := client.GetChatCompletionsStream(context.TODO(), azopenai.ChatCompletionsStreamOptions{
		// This is a conversation in progress.
		// NOTE: all messages count against token usage for this API.
		Messages:       messages,
		N:              to.Ptr[int32](1),
		DeploymentName: &modelDeploymentID,
	}, nil)

	if err != nil {
		// TODO: Update the following line with your application specific error handling logic
		log.Printf("ERROR: %s", err)
		return
	}

	defer resp.ChatCompletionsStream.Close()

	gotReply := false

	for {
		chatCompletions, err := resp.ChatCompletionsStream.Read()

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			//  TODO: Update the following line with your application specific error handling logic
			log.Printf("ERROR: %s", err)
			return
		}

		for _, choice := range chatCompletions.Choices {
			gotReply = true

			text := ""

			if choice.Delta.Content != nil {
				text = *choice.Delta.Content
			}

			role := ""

			if choice.Delta.Role != nil {
				role = string(*choice.Delta.Role)
			}

			fmt.Fprintf(os.Stderr, "Content[%d], role %q: %q\n", *choice.Index, role, text)
		}
	}

	if gotReply {
		fmt.Fprintf(os.Stderr, "Got chat completions streaming reply\n")
	}

	// Output:
}

func ExampleClient_GetChatCompletions_structuredOutputs() {
	// This example is a Go implementation of the example from this article:
	// https://openai.com/index/introducing-structured-outputs-in-the-api/
	modelDeploymentID := os.Getenv("AOAI_CHAT_COMPLETIONS_STRUCTURED_OUTPUTS_MODEL")

	// Ex: "https://<your-azure-openai-host>.openai.azure.com"
	azureOpenAIEndpoint := os.Getenv("AOAI_CHAT_COMPLETIONS_STRUCTURED_OUTPUTS_ENDPOINT")

	if modelDeploymentID == "" || azureOpenAIEndpoint == "" {
		fmt.Fprintf(os.Stderr, "Skipping example, environment variables missing\n")
		return
	}

	tokenCredential, err := azidentity.NewDefaultAzureCredential(nil)

	if err != nil {
		//  TODO: Update the following line with your application specific error handling logic
		log.Printf("ERROR: %s", err)
		return
	}

	client, err := azopenai.NewClient(azureOpenAIEndpoint, tokenCredential, nil)

	if err != nil {
		// TODO: Update the following line with your application specific error handling logic
		log.Printf("ERROR: %s", err)
		return
	}

	structuredJSONToolDefParams := []byte(`{
		"type": "object",
		"properties": {
			"table_name": {
				"type": "string",
				"enum": [
					"orders"
				]
			},
			"columns": {
				"type": "array",
				"items": {
					"type": "string",
					"enum": [
						"id",
						"status",
						"expected_delivery_date",
						"delivered_at",
						"shipped_at",
						"ordered_at",
						"canceled_at"
					]
				}
			},
			"conditions": {
				"type": "array",
				"items": {
					"type": "object",
					"properties": {
						"column": {
							"type": "string"
						},
						"operator": {
							"type": "string",
							"enum": [
								"=",
								">",
								"<",
								">=",
								"<=",
								"!="
							]
						},
						"value": {
							"anyOf": [
								{
									"type": "string"
								},
								{
									"type": "number"
								},
								{
									"type": "object",
									"properties": {
										"column_name": {
											"type": "string"
										}
									},
									"required": [
										"column_name"
									],
									"additionalProperties": false
								}
							]
						}
					},
					"required": [
						"column",
						"operator",
						"value"
					],
					"additionalProperties": false
				}
			},
			"order_by": {
				"type": "string",
				"enum": [
					"asc",
					"desc"
				]
			}
		},
		"required": [
			"table_name",
			"columns",
			"conditions",
			"order_by"
		],
		"additionalProperties": false
	}`)

	resp, err := client.GetChatCompletions(context.TODO(), azopenai.ChatCompletionsOptions{
		DeploymentName: &modelDeploymentID,
		Messages: []azopenai.ChatRequestMessageClassification{
			&azopenai.ChatRequestAssistantMessage{
				Content: azopenai.NewChatRequestAssistantMessageContent("You are a helpful assistant. The current date is August 6, 2024. You help users query for the data they are looking for by calling the query function."),
			},
			&azopenai.ChatRequestUserMessage{
				Content: azopenai.NewChatRequestUserMessageContent("look up all my orders in may of last year that were fulfilled but not delivered on time"),
			},
		},
		Tools: []azopenai.ChatCompletionsToolDefinitionClassification{
			&azopenai.ChatCompletionsFunctionToolDefinition{
				Function: &azopenai.ChatCompletionsFunctionToolDefinitionFunction{
					Strict:     to.Ptr(true),
					Name:       to.Ptr("query"),
					Parameters: structuredJSONToolDefParams,
				},
			},
		},
	}, nil)

	if err != nil {
		// TODO: Update the following line with your application specific error handling logic
		log.Printf("ERROR: %s", err)
		return
	}

	// and now the results should contain proper JSON for the arguments, every time.
	fn := resp.Choices[0].Message.ToolCalls[0].(*azopenai.ChatCompletionsFunctionToolCall).Function

	argumentsObj := map[string]any{}

	err = json.Unmarshal([]byte(*fn.Arguments), &argumentsObj)

	if err != nil {
		//  TODO: Update the following line with your application specific error handling logic
		log.Printf("ERROR: %s", err)
		return
	}

	fmt.Printf("%#v\n", argumentsObj)
}

func ExampleClient_GetChatCompletions_structuredOutputs_responseFormat() {
	// This example is a Go implementation of the example from this article:
	// https://openai.com/index/introducing-structured-outputs-in-the-api/
	modelDeploymentID := os.Getenv("AOAI_CHAT_COMPLETIONS_STRUCTURED_OUTPUTS_MODEL")

	// Ex: "https://<your-azure-openai-host>.openai.azure.com"
	azureOpenAIEndpoint := os.Getenv("AOAI_CHAT_COMPLETIONS_STRUCTURED_OUTPUTS_ENDPOINT")

	if modelDeploymentID == "" || azureOpenAIEndpoint == "" {
		fmt.Fprintf(os.Stderr, "Skipping example, environment variables missing\n")
		return
	}

	// In Azure OpenAI you must deploy a model before you can use it in your client. For more information
	// see here: https://learn.microsoft.com/azure/cognitive-services/openai/how-to/create-resource
	tokenCredential, err := azidentity.NewDefaultAzureCredential(nil)

	if err != nil {
		//  TODO: Update the following line with your application specific error handling logic
		log.Printf("ERROR: %s", err)
		return
	}

	client, err := azopenai.NewClient(azureOpenAIEndpoint, tokenCredential, nil)

	if err != nil {
		// TODO: Update the following line with your application specific error handling logic
		log.Printf("ERROR: %s", err)
		return
	}

	structuredOutputJSONSchema := []byte(`{
	"type": "object",
	"properties": {
		"steps": {
			"type": "array",
			"items": {
				"type": "object",
				"properties": {
					"explanation": {
						"type": "string"
					},
					"output": {
						"type": "string"
					}
				},
				"required": [
					"explanation",
					"output"
				],
				"additionalProperties": false
			}
		},
		"final_answer": {
			"type": "string"
		}
	},
	"required": [
		"steps",
		"final_answer"
	],
	"additionalProperties": false
}`)

	resp, err := client.GetChatCompletions(context.TODO(), azopenai.ChatCompletionsOptions{
		DeploymentName: &modelDeploymentID,
		Messages: []azopenai.ChatRequestMessageClassification{
			&azopenai.ChatRequestAssistantMessage{
				Content: azopenai.NewChatRequestAssistantMessageContent("You are a helpful math tutor."),
			},
			&azopenai.ChatRequestUserMessage{
				Content: azopenai.NewChatRequestUserMessageContent("solve 8x + 31 = 2"),
			},
		},
		ResponseFormat: &azopenai.ChatCompletionsJSONSchemaResponseFormat{
			JSONSchema: &azopenai.ChatCompletionsJSONSchemaResponseFormatJSONSchema{
				Name:   to.Ptr("math_response"),
				Strict: to.Ptr(true),
				Schema: structuredOutputJSONSchema,
			},
		},
	}, nil)

	if err != nil {
		// TODO: Update the following line with your application specific error handling logic
		log.Printf("ERROR: %s", err)
		return
	}

	responseObj := map[string]any{}
	err = json.Unmarshal([]byte(*resp.Choices[0].Message.Content), &responseObj)

	if err != nil {
		//  TODO: Update the following line with your application specific error handling logic
		log.Printf("ERROR: %s", err)
		return
	}

	fmt.Printf("%#v", responseObj)
}
