package logger

type Field string

const (
	////////////////// General Fields //////////////////

	FieldInstanceID Field = "instance_id"
	FieldRequestID  Field = "request_id"

	////////////////// User Info //////////////////

	FieldUserID     Field = "user_id"
	FieldAppVersion Field = "app_version"
	FieldDeviceID   Field = "device_id"
	FieldDeviceInfo Field = "device_info"

	////////////////// Debug Fields //////////////////

	FieldCaller       Field = "caller"
	FieldSourceCaller Field = "caller_source"
	FieldCommitSha    Field = "commit_sha"
	FieldBuildDate    Field = "build_date"

	////////////////// Error Fields //////////////////

	FieldError     Field = "error"
	FieldErrorCode Field = "error_code"
	FieldUserError Field = "user_error"

	////////////////// Business Fields //////////////////

	FieldEntityType   Field = "entity_type"
	FieldEntityID     Field = "entity_id"
	FieldEntityStep   Field = "entity_step"
	FieldMethodName   Field = "method_name"
	FieldMethodInput  Field = "method_input"
	FieldMethodOutput Field = "method_output"

	////////////////// HTTP Fields //////////////////

	FieldHTTPRequest    Field = "http_request"
	FieldRequest        Field = "http_request_body"
	FieldResponse       Field = "http_response_body"
	FieldHTTPStatusCode Field = "http_status_code"
	FieldHTTPMethod     Field = "http_method"
	FieldHTTPURL        Field = "http_url"
	FieldHTTPLatency    Field = "http_latency"

	////////////////// DB Fields //////////////////

	FieldDBQuery       Field = "db_query"
	FieldDBRowAffected Field = "db_rows_affected"
	FieldDBDelay       Field = "db_delay"

	////////////////// Security Fields //////////////////

	FieldUserAgent Field = "user_agent"
	FieldClientIP  Field = "client_ip"

	////////////////// Performance Fields //////////////////

	FieldExecutionTime Field = "execution_time"
	FieldMemoryUsage   Field = "memory_usage"
)

func (f Field) String() string {
	return string(f)
}
