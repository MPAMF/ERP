/*
 * Medical Records API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type MedicalActResult struct {
	Id int `json:"id,omitempty"`

	MedicalActId int `json:"medical_act_id,omitempty"`

	FileName string `json:"file_name,omitempty"`

	FileBinaryData []byte `json:"-"`

	FileData string `json:"file_data,omitempty"`
}
