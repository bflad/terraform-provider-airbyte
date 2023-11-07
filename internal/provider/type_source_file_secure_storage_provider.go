// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

type SourceFileSecureStorageProvider struct {
	AzBlobAzureBlobStorage         *AzBlobAzureBlobStorage              `tfsdk:"az_blob_azure_blob_storage"`
	GCSGoogleCloudStorage          *GCSGoogleCloudStorage               `tfsdk:"gcs_google_cloud_storage"`
	HTTPSPublicWeb                 *HTTPSPublicWeb                      `tfsdk:"https_public_web"`
	S3AmazonWebServices            *SourceFileSecureS3AmazonWebServices `tfsdk:"s3_amazon_web_services"`
	SCPSecureCopyProtocol          *SCPSecureCopyProtocol               `tfsdk:"scp_secure_copy_protocol"`
	SFTPSecureFileTransferProtocol *SCPSecureCopyProtocol               `tfsdk:"sftp_secure_file_transfer_protocol"`
	SSHSecureShell                 *SCPSecureCopyProtocol               `tfsdk:"ssh_secure_shell"`
}
