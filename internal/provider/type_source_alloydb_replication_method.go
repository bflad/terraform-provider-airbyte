// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

type SourceAlloydbReplicationMethod struct {
	SourceAlloydbReplicationMethodLogicalReplicationCDC       *SourceAlloydbReplicationMethodLogicalReplicationCDC       `tfsdk:"source_alloydb_replication_method_logical_replication_cdc"`
	SourceAlloydbReplicationMethodStandard                    *SourceAlloydbReplicationMethodStandard                    `tfsdk:"source_alloydb_replication_method_standard"`
	SourceAlloydbReplicationMethodStandardXmin                *SourceAlloydbReplicationMethodStandardXmin                `tfsdk:"source_alloydb_replication_method_standard_xmin"`
	SourceAlloydbUpdateReplicationMethodLogicalReplicationCDC *SourceAlloydbUpdateReplicationMethodLogicalReplicationCDC `tfsdk:"source_alloydb_update_replication_method_logical_replication_cdc"`
	SourceAlloydbUpdateReplicationMethodStandard              *SourceAlloydbReplicationMethodStandard                    `tfsdk:"source_alloydb_update_replication_method_standard"`
	SourceAlloydbUpdateReplicationMethodStandardXmin          *SourceAlloydbReplicationMethodStandardXmin                `tfsdk:"source_alloydb_update_replication_method_standard_xmin"`
}
