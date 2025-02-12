package validator

import (
    "strings"
)

type ValidationError struct {
    Field   string `json:"field"`
    Message string `json:"message"`
}

func ValidateCreateDialogRequest(prompt string) []ValidationError {
    var errors []ValidationError
    
    if strings.TrimSpace(prompt) == "" {
        errors = append(errors, ValidationError{
            Field:   "prompt",
            Message: "Prompt cannot be empty",
        })
    }
    
    if len(prompt) > 1000 {
        errors = append(errors, ValidationError{
            Field:   "prompt",
            Message: "Prompt length cannot exceed 1000 characters",
        })
    }
    
    return errors
}