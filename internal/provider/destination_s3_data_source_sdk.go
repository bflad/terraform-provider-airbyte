// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationS3DataSourceModel) RefreshFromGetResponse(resp *shared.DestinationS3GetResponse) {
	if resp.Configuration.AccessKeyID != nil {
		r.Configuration.AccessKeyID = types.StringValue(*resp.Configuration.AccessKeyID)
	} else {
		r.Configuration.AccessKeyID = types.StringNull()
	}
	r.Configuration.DestinationType = types.StringValue(string(resp.Configuration.DestinationType))
	if resp.Configuration.FileNamePattern != nil {
		r.Configuration.FileNamePattern = types.StringValue(*resp.Configuration.FileNamePattern)
	} else {
		r.Configuration.FileNamePattern = types.StringNull()
	}
	if resp.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro != nil {
		r.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro = &DestinationS3OutputFormatAvroApacheAvro{}
		if resp.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecBzip2 != nil {
			r.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecBzip2 = &DestinationGcsOutputFormatAvroApacheAvroCompressionCodecBzip2{}
			r.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecBzip2.Codec = types.StringValue(string(resp.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecBzip2.Codec))
		}
		if resp.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecDeflate != nil {
			r.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecDeflate = &DestinationGcsOutputFormatAvroApacheAvroCompressionCodecDeflate{}
			r.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecDeflate.Codec = types.StringValue(string(resp.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecDeflate.Codec))
			r.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecDeflate.CompressionLevel = types.Int64Value(resp.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecDeflate.CompressionLevel)
		}
		if resp.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecNoCompression != nil {
			r.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecNoCompression = &DestinationGcsOutputFormatAvroApacheAvroCompressionCodecNoCompression{}
			r.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecNoCompression.Codec = types.StringValue(string(resp.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecNoCompression.Codec))
		}
		if resp.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecSnappy != nil {
			r.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecSnappy = &DestinationGcsOutputFormatAvroApacheAvroCompressionCodecSnappy{}
			r.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecSnappy.Codec = types.StringValue(string(resp.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecSnappy.Codec))
		}
		if resp.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecXz != nil {
			r.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecXz = &DestinationGcsOutputFormatAvroApacheAvroCompressionCodecXz{}
			r.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecXz.Codec = types.StringValue(string(resp.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecXz.Codec))
			r.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecXz.CompressionLevel = types.Int64Value(resp.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecXz.CompressionLevel)
		}
		if resp.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecZstandard != nil {
			r.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecZstandard = &DestinationGcsOutputFormatAvroApacheAvroCompressionCodecZstandard{}
			r.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecZstandard.Codec = types.StringValue(string(resp.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecZstandard.Codec))
			r.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecZstandard.CompressionLevel = types.Int64Value(resp.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecZstandard.CompressionLevel)
			if resp.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecZstandard.IncludeChecksum != nil {
				r.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecZstandard.IncludeChecksum = types.BoolValue(*resp.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecZstandard.IncludeChecksum)
			} else {
				r.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.CompressionCodec.DestinationS3OutputFormatAvroApacheAvroCompressionCodecZstandard.IncludeChecksum = types.BoolNull()
			}
		}
		r.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.FormatType = types.StringValue(string(resp.Configuration.Format.DestinationS3OutputFormatAvroApacheAvro.FormatType))
	}
	if resp.Configuration.Format.DestinationS3OutputFormatCSVCommaSeparatedValues != nil {
		r.Configuration.Format.DestinationS3OutputFormatCSVCommaSeparatedValues = &DestinationS3OutputFormatCSVCommaSeparatedValues{}
		if resp.Configuration.Format.DestinationS3OutputFormatCSVCommaSeparatedValues.Compression == nil {
			r.Configuration.Format.DestinationS3OutputFormatCSVCommaSeparatedValues.Compression = nil
		} else {
			r.Configuration.Format.DestinationS3OutputFormatCSVCommaSeparatedValues.Compression = &DestinationS3OutputFormatCSVCommaSeparatedValuesCompression{}
			if resp.Configuration.Format.DestinationS3OutputFormatCSVCommaSeparatedValues.Compression.DestinationS3OutputFormatCSVCommaSeparatedValuesCompressionGZIP != nil {
				r.Configuration.Format.DestinationS3OutputFormatCSVCommaSeparatedValues.Compression.DestinationS3OutputFormatCSVCommaSeparatedValuesCompressionGZIP = &DestinationGcsOutputFormatCSVCommaSeparatedValuesCompressionGZIP{}
				if resp.Configuration.Format.DestinationS3OutputFormatCSVCommaSeparatedValues.Compression.DestinationS3OutputFormatCSVCommaSeparatedValuesCompressionGZIP.CompressionType != nil {
					r.Configuration.Format.DestinationS3OutputFormatCSVCommaSeparatedValues.Compression.DestinationS3OutputFormatCSVCommaSeparatedValuesCompressionGZIP.CompressionType = types.StringValue(string(*resp.Configuration.Format.DestinationS3OutputFormatCSVCommaSeparatedValues.Compression.DestinationS3OutputFormatCSVCommaSeparatedValuesCompressionGZIP.CompressionType))
				} else {
					r.Configuration.Format.DestinationS3OutputFormatCSVCommaSeparatedValues.Compression.DestinationS3OutputFormatCSVCommaSeparatedValuesCompressionGZIP.CompressionType = types.StringNull()
				}
			}
			if resp.Configuration.Format.DestinationS3OutputFormatCSVCommaSeparatedValues.Compression.DestinationS3OutputFormatCSVCommaSeparatedValuesCompressionNoCompression != nil {
				r.Configuration.Format.DestinationS3OutputFormatCSVCommaSeparatedValues.Compression.DestinationS3OutputFormatCSVCommaSeparatedValuesCompressionNoCompression = &DestinationGcsOutputFormatCSVCommaSeparatedValuesCompressionNoCompression{}
				if resp.Configuration.Format.DestinationS3OutputFormatCSVCommaSeparatedValues.Compression.DestinationS3OutputFormatCSVCommaSeparatedValuesCompressionNoCompression.CompressionType != nil {
					r.Configuration.Format.DestinationS3OutputFormatCSVCommaSeparatedValues.Compression.DestinationS3OutputFormatCSVCommaSeparatedValuesCompressionNoCompression.CompressionType = types.StringValue(string(*resp.Configuration.Format.DestinationS3OutputFormatCSVCommaSeparatedValues.Compression.DestinationS3OutputFormatCSVCommaSeparatedValuesCompressionNoCompression.CompressionType))
				} else {
					r.Configuration.Format.DestinationS3OutputFormatCSVCommaSeparatedValues.Compression.DestinationS3OutputFormatCSVCommaSeparatedValuesCompressionNoCompression.CompressionType = types.StringNull()
				}
			}
		}
		r.Configuration.Format.DestinationS3OutputFormatCSVCommaSeparatedValues.Flattening = types.StringValue(string(resp.Configuration.Format.DestinationS3OutputFormatCSVCommaSeparatedValues.Flattening))
		r.Configuration.Format.DestinationS3OutputFormatCSVCommaSeparatedValues.FormatType = types.StringValue(string(resp.Configuration.Format.DestinationS3OutputFormatCSVCommaSeparatedValues.FormatType))
	}
	if resp.Configuration.Format.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSON != nil {
		r.Configuration.Format.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSON = &DestinationS3OutputFormatJSONLinesNewlineDelimitedJSON{}
		if resp.Configuration.Format.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSON.Compression == nil {
			r.Configuration.Format.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSON.Compression = nil
		} else {
			r.Configuration.Format.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSON.Compression = &DestinationS3OutputFormatJSONLinesNewlineDelimitedJSONCompression{}
			if resp.Configuration.Format.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP != nil {
				r.Configuration.Format.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP = &DestinationGcsOutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP{}
				if resp.Configuration.Format.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP.CompressionType != nil {
					r.Configuration.Format.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP.CompressionType = types.StringValue(string(*resp.Configuration.Format.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP.CompressionType))
				} else {
					r.Configuration.Format.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSONCompressionGZIP.CompressionType = types.StringNull()
				}
			}
			if resp.Configuration.Format.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression != nil {
				r.Configuration.Format.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression = &DestinationGcsOutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression{}
				if resp.Configuration.Format.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression.CompressionType != nil {
					r.Configuration.Format.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression.CompressionType = types.StringValue(string(*resp.Configuration.Format.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression.CompressionType))
				} else {
					r.Configuration.Format.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSON.Compression.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSONCompressionNoCompression.CompressionType = types.StringNull()
				}
			}
		}
		if resp.Configuration.Format.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSON.Flattening != nil {
			r.Configuration.Format.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSON.Flattening = types.StringValue(string(*resp.Configuration.Format.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSON.Flattening))
		} else {
			r.Configuration.Format.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSON.Flattening = types.StringNull()
		}
		r.Configuration.Format.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSON.FormatType = types.StringValue(string(resp.Configuration.Format.DestinationS3OutputFormatJSONLinesNewlineDelimitedJSON.FormatType))
	}
	if resp.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage != nil {
		r.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage = &DestinationS3OutputFormatParquetColumnarStorage{}
		if resp.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage.BlockSizeMb != nil {
			r.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage.BlockSizeMb = types.Int64Value(*resp.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage.BlockSizeMb)
		} else {
			r.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage.BlockSizeMb = types.Int64Null()
		}
		if resp.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage.CompressionCodec != nil {
			r.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage.CompressionCodec = types.StringValue(string(*resp.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage.CompressionCodec))
		} else {
			r.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage.CompressionCodec = types.StringNull()
		}
		if resp.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage.DictionaryEncoding != nil {
			r.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage.DictionaryEncoding = types.BoolValue(*resp.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage.DictionaryEncoding)
		} else {
			r.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage.DictionaryEncoding = types.BoolNull()
		}
		if resp.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage.DictionaryPageSizeKb != nil {
			r.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage.DictionaryPageSizeKb = types.Int64Value(*resp.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage.DictionaryPageSizeKb)
		} else {
			r.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage.DictionaryPageSizeKb = types.Int64Null()
		}
		r.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage.FormatType = types.StringValue(string(resp.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage.FormatType))
		if resp.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage.MaxPaddingSizeMb != nil {
			r.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage.MaxPaddingSizeMb = types.Int64Value(*resp.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage.MaxPaddingSizeMb)
		} else {
			r.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage.MaxPaddingSizeMb = types.Int64Null()
		}
		if resp.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage.PageSizeKb != nil {
			r.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage.PageSizeKb = types.Int64Value(*resp.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage.PageSizeKb)
		} else {
			r.Configuration.Format.DestinationS3OutputFormatParquetColumnarStorage.PageSizeKb = types.Int64Null()
		}
	}
	if resp.Configuration.Format.DestinationS3UpdateOutputFormatAvroApacheAvro != nil {
		r.Configuration.Format.DestinationS3UpdateOutputFormatAvroApacheAvro = &DestinationS3UpdateOutputFormatAvroApacheAvro{}
	}
	if resp.Configuration.Format.DestinationS3UpdateOutputFormatCSVCommaSeparatedValues != nil {
		r.Configuration.Format.DestinationS3UpdateOutputFormatCSVCommaSeparatedValues = &DestinationS3UpdateOutputFormatCSVCommaSeparatedValues{}
	}
	if resp.Configuration.Format.DestinationS3UpdateOutputFormatJSONLinesNewlineDelimitedJSON != nil {
		r.Configuration.Format.DestinationS3UpdateOutputFormatJSONLinesNewlineDelimitedJSON = &DestinationS3UpdateOutputFormatJSONLinesNewlineDelimitedJSON{}
	}
	if resp.Configuration.Format.DestinationS3UpdateOutputFormatParquetColumnarStorage != nil {
		r.Configuration.Format.DestinationS3UpdateOutputFormatParquetColumnarStorage = &DestinationS3OutputFormatParquetColumnarStorage{}
	}
	r.Configuration.S3BucketName = types.StringValue(resp.Configuration.S3BucketName)
	r.Configuration.S3BucketPath = types.StringValue(resp.Configuration.S3BucketPath)
	r.Configuration.S3BucketRegion = types.StringValue(string(resp.Configuration.S3BucketRegion))
	if resp.Configuration.S3Endpoint != nil {
		r.Configuration.S3Endpoint = types.StringValue(*resp.Configuration.S3Endpoint)
	} else {
		r.Configuration.S3Endpoint = types.StringNull()
	}
	if resp.Configuration.S3PathFormat != nil {
		r.Configuration.S3PathFormat = types.StringValue(*resp.Configuration.S3PathFormat)
	} else {
		r.Configuration.S3PathFormat = types.StringNull()
	}
	if resp.Configuration.SecretAccessKey != nil {
		r.Configuration.SecretAccessKey = types.StringValue(*resp.Configuration.SecretAccessKey)
	} else {
		r.Configuration.SecretAccessKey = types.StringNull()
	}
	if resp.DestinationID != nil {
		r.DestinationID = types.StringValue(*resp.DestinationID)
	} else {
		r.DestinationID = types.StringNull()
	}
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
