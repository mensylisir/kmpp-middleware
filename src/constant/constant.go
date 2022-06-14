package constant

import "errors"

const (
	DefaultPassword = "Def@u1tpwd"
)

const (
	BatchOperationUpdate = "update"
	BatchOperationCreate = "create"
	BatchOperationDelete = "delete"
)

var (
	NotSupportedBatchOperation = errors.New("not supported operation")
)

const (
	CenterCluster = "center"
)

const (
	ClusterNormal   = "Normal"
	ClusterInnormal = "Innormal"
)

const (
	BASIC      = "basic"
	ADVANCE    = "advance"
	POSTGRESQL = "postgres"
)

const (
	PARAM_MISSING = "parameter %v is missing, please check your requests parameters"
	TOKEN_INVALID = "token is invalid"
	PARAM_EMPTY   = "param_empty"
	NOT_SUPPORT   = "not_support"
)

const (
	Azure = "AZURE"
	S3    = "S3"
	OSS   = "OSS"
	Sftp  = "SFTP"
	MINIO = "MINIO"
)

var (
	ErrOriginalNotMatch  = errors.New("ORIGINAL_NOT_MATCH")
	ErrUserNotFound      = errors.New("USER_NOT_FOUND")
	ErrUserIsNotActive   = errors.New("USER_IS_NOT_ACTIVE")
	ErrUserNameExist     = errors.New("NAME_EXISTS")
	ErrLdapDisable       = errors.New("LDAP_DISABLE")
	ErrEmailExist        = errors.New("EMAIL_EXIST")
	ErrNamePwdFailed     = errors.New("NAME_PASSWORD_SAME_FAILED")
	ErrEmailDisable      = errors.New("EMAIL_DISABLE")
	ErrEmailNotMatch     = errors.New("EMAIL_NOT_MATCH")
	ErrNameOrPasswordErr = errors.New("NAME_PASSWORD_ERROR")
	ErrResourceExist     = errors.New("RESOURCE_EXISTS")
)

const (
	SystemRoleSuperAdmin = 0
)

const (
	Local = "LOCAL"
	Ldap  = "LDAP"
)
