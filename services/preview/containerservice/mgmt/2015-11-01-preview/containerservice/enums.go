package containerservice

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// OchestratorTypes enumerates the values for ochestrator types.
type OchestratorTypes string

const (
	// DCOS ...
	DCOS OchestratorTypes = "DCOS"
	// Mesos ...
	Mesos OchestratorTypes = "Mesos"
	// SwarmPreview ...
	SwarmPreview OchestratorTypes = "SwarmPreview"
)

// PossibleOchestratorTypesValues returns an array of possible values for the OchestratorTypes const type.
func PossibleOchestratorTypesValues() []OchestratorTypes {
	return []OchestratorTypes{DCOS, Mesos, SwarmPreview}
}

// VMSizeTypes enumerates the values for vm size types.
type VMSizeTypes string

const (
	// StandardA0 ...
	StandardA0 VMSizeTypes = "Standard_A0"
	// StandardA1 ...
	StandardA1 VMSizeTypes = "Standard_A1"
	// StandardA10 ...
	StandardA10 VMSizeTypes = "Standard_A10"
	// StandardA11 ...
	StandardA11 VMSizeTypes = "Standard_A11"
	// StandardA2 ...
	StandardA2 VMSizeTypes = "Standard_A2"
	// StandardA3 ...
	StandardA3 VMSizeTypes = "Standard_A3"
	// StandardA4 ...
	StandardA4 VMSizeTypes = "Standard_A4"
	// StandardA5 ...
	StandardA5 VMSizeTypes = "Standard_A5"
	// StandardA6 ...
	StandardA6 VMSizeTypes = "Standard_A6"
	// StandardA7 ...
	StandardA7 VMSizeTypes = "Standard_A7"
	// StandardA8 ...
	StandardA8 VMSizeTypes = "Standard_A8"
	// StandardA9 ...
	StandardA9 VMSizeTypes = "Standard_A9"
	// StandardD1 ...
	StandardD1 VMSizeTypes = "Standard_D1"
	// StandardD11 ...
	StandardD11 VMSizeTypes = "Standard_D11"
	// StandardD11V2 ...
	StandardD11V2 VMSizeTypes = "Standard_D11_v2"
	// StandardD12 ...
	StandardD12 VMSizeTypes = "Standard_D12"
	// StandardD12V2 ...
	StandardD12V2 VMSizeTypes = "Standard_D12_v2"
	// StandardD13 ...
	StandardD13 VMSizeTypes = "Standard_D13"
	// StandardD13V2 ...
	StandardD13V2 VMSizeTypes = "Standard_D13_v2"
	// StandardD14 ...
	StandardD14 VMSizeTypes = "Standard_D14"
	// StandardD14V2 ...
	StandardD14V2 VMSizeTypes = "Standard_D14_v2"
	// StandardD1V2 ...
	StandardD1V2 VMSizeTypes = "Standard_D1_v2"
	// StandardD2 ...
	StandardD2 VMSizeTypes = "Standard_D2"
	// StandardD2V2 ...
	StandardD2V2 VMSizeTypes = "Standard_D2_v2"
	// StandardD3 ...
	StandardD3 VMSizeTypes = "Standard_D3"
	// StandardD3V2 ...
	StandardD3V2 VMSizeTypes = "Standard_D3_v2"
	// StandardD4 ...
	StandardD4 VMSizeTypes = "Standard_D4"
	// StandardD4V2 ...
	StandardD4V2 VMSizeTypes = "Standard_D4_v2"
	// StandardD5V2 ...
	StandardD5V2 VMSizeTypes = "Standard_D5_v2"
	// StandardDS1 ...
	StandardDS1 VMSizeTypes = "Standard_DS1"
	// StandardDS11 ...
	StandardDS11 VMSizeTypes = "Standard_DS11"
	// StandardDS12 ...
	StandardDS12 VMSizeTypes = "Standard_DS12"
	// StandardDS13 ...
	StandardDS13 VMSizeTypes = "Standard_DS13"
	// StandardDS14 ...
	StandardDS14 VMSizeTypes = "Standard_DS14"
	// StandardDS2 ...
	StandardDS2 VMSizeTypes = "Standard_DS2"
	// StandardDS3 ...
	StandardDS3 VMSizeTypes = "Standard_DS3"
	// StandardDS4 ...
	StandardDS4 VMSizeTypes = "Standard_DS4"
	// StandardG1 ...
	StandardG1 VMSizeTypes = "Standard_G1"
	// StandardG2 ...
	StandardG2 VMSizeTypes = "Standard_G2"
	// StandardG3 ...
	StandardG3 VMSizeTypes = "Standard_G3"
	// StandardG4 ...
	StandardG4 VMSizeTypes = "Standard_G4"
	// StandardG5 ...
	StandardG5 VMSizeTypes = "Standard_G5"
	// StandardGS1 ...
	StandardGS1 VMSizeTypes = "Standard_GS1"
	// StandardGS2 ...
	StandardGS2 VMSizeTypes = "Standard_GS2"
	// StandardGS3 ...
	StandardGS3 VMSizeTypes = "Standard_GS3"
	// StandardGS4 ...
	StandardGS4 VMSizeTypes = "Standard_GS4"
	// StandardGS5 ...
	StandardGS5 VMSizeTypes = "Standard_GS5"
)

// PossibleVMSizeTypesValues returns an array of possible values for the VMSizeTypes const type.
func PossibleVMSizeTypesValues() []VMSizeTypes {
	return []VMSizeTypes{StandardA0, StandardA1, StandardA10, StandardA11, StandardA2, StandardA3, StandardA4, StandardA5, StandardA6, StandardA7, StandardA8, StandardA9, StandardD1, StandardD11, StandardD11V2, StandardD12, StandardD12V2, StandardD13, StandardD13V2, StandardD14, StandardD14V2, StandardD1V2, StandardD2, StandardD2V2, StandardD3, StandardD3V2, StandardD4, StandardD4V2, StandardD5V2, StandardDS1, StandardDS11, StandardDS12, StandardDS13, StandardDS14, StandardDS2, StandardDS3, StandardDS4, StandardG1, StandardG2, StandardG3, StandardG4, StandardG5, StandardGS1, StandardGS2, StandardGS3, StandardGS4, StandardGS5}
}
