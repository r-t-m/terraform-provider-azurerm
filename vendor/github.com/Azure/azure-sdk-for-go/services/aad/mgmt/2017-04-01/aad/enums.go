package aad

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Category enumerates the values for category.
type Category string

const (
	// AuditLogs ...
	AuditLogs Category = "AuditLogs"
	// SignInLogs ...
	SignInLogs Category = "SignInLogs"
)

// PossibleCategoryValues returns an array of possible values for the Category const type.
func PossibleCategoryValues() []Category {
	return []Category{AuditLogs, SignInLogs}
}

// CategoryType enumerates the values for category type.
type CategoryType string

const (
	// Logs ...
	Logs CategoryType = "Logs"
)

// PossibleCategoryTypeValues returns an array of possible values for the CategoryType const type.
func PossibleCategoryTypeValues() []CategoryType {
	return []CategoryType{Logs}
}
