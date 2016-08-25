// +build windows

// Implements flags necessary for working with WinHTTP.

package nativeweb

const (
	INTERNET_DEFAULT_PORT       = 0
	INTERNET_DEFAULT_HTTP_PORT  = 80
	INTERNET_DEFAULT_HTTPS_PORT = 443

	WINHTTP_FLAG_ASYNC = 0x10000000

	WINHTTP_FLAG_SECURE               = 0x00800000
	WINHTTP_FLAG_ESCAPE_PERCENT       = 0x00000004
	WINHTTP_FLAG_NULL_CODEPAGE        = 0x00000008
	WINHTTP_FLAG_BYPASS_PROXY_CACHE   = 0x00000100
	WINHTTP_FLAG_REFRESH              = 0x00000100
	WINHTTP_FLAG_ESCAPE_DISABLE       = 0x00000040
	WINHTTP_FLAG_ESCAPE_DISABLE_QUERY = 0x00000080

	WINHTTP_AUTOPROXY_AUTO_DETECT         = 0x00000001
	WINHTTP_AUTOPROXY_CONFIG_URL          = 0x00000002
	WINHTTP_AUTOPROXY_HOST_KEEPCASE       = 0x00000004
	WINHTTP_AUTOPROXY_HOST_LOWERCASE      = 0x00000008
	WINHTTP_AUTOPROXY_RUN_INPROCESS       = 0x00010000
	WINHTTP_AUTOPROXY_RUN_OUTPROCESS_ONLY = 0x00020000
	WINHTTP_AUTOPROXY_NO_DIRECTACCESS     = 0x00040000
	WINHTTP_AUTOPROXY_NO_CACHE_CLIENT     = 0x00080000
	WINHTTP_AUTOPROXY_NO_CACHE_SVC        = 0x00100000

	WINHTTP_AUTOPROXY_SORT_RESULTS = 0x00400000

	WINHTTP_AUTO_DETECT_TYPE_DHCP  = 0x00000001
	WINHTTP_AUTO_DETECT_TYPE_DNS_A = 0x00000002

	WINHTTP_ACCESS_TYPE_DEFAULT_PROXY   = 0
	WINHTTP_ACCESS_TYPE_NO_PROXY        = 1
	WINHTTP_ACCESS_TYPE_NAMED_PROXY     = 3
	WINHTTP_ACCESS_TYPE_AUTOMATIC_PROXY = 4

	WINHTTP_NO_PROXY_NAME   = 0
	WINHTTP_NO_PROXY_BYPASS = 0

	WINHTTP_NO_ADDITIONAL_HEADERS = 0
	WINHTTP_NO_REQUEST_DATA       = 0

	WINHTTP_QUERY_MIME_VERSION              = 0
	WINHTTP_QUERY_CONTENT_TYPE              = 1
	WINHTTP_QUERY_CONTENT_TRANSFER_ENCODING = 2
	WINHTTP_QUERY_CONTENT_ID                = 3
	WINHTTP_QUERY_CONTENT_DESCRIPTION       = 4
	WINHTTP_QUERY_CONTENT_LENGTH            = 5
	WINHTTP_QUERY_CONTENT_LANGUAGE          = 6
	WINHTTP_QUERY_ALLOW                     = 7
	WINHTTP_QUERY_PUBLIC                    = 8
	WINHTTP_QUERY_DATE                      = 9
	WINHTTP_QUERY_EXPIRES                   = 10
	WINHTTP_QUERY_LAST_MODIFIED             = 11
	WINHTTP_QUERY_MESSAGE_ID                = 12
	WINHTTP_QUERY_URI                       = 13
	WINHTTP_QUERY_DERIVED_FROM              = 14
	WINHTTP_QUERY_COST                      = 15
	WINHTTP_QUERY_LINK                      = 16
	WINHTTP_QUERY_PRAGMA                    = 17
	WINHTTP_QUERY_VERSION                   = 18 // special: part of status line
	WINHTTP_QUERY_STATUS_CODE               = 19 // special: part of status line
	WINHTTP_QUERY_STATUS_TEXT               = 20 // special: part of status line
	WINHTTP_QUERY_RAW_HEADERS               = 21 // special: all headers as ASCIIZ
	WINHTTP_QUERY_RAW_HEADERS_CRLF          = 22 // special: all headers
	WINHTTP_QUERY_CONNECTION                = 23
	WINHTTP_QUERY_ACCEPT                    = 24
	WINHTTP_QUERY_ACCEPT_CHARSET            = 25
	WINHTTP_QUERY_ACCEPT_ENCODING           = 26
	WINHTTP_QUERY_ACCEPT_LANGUAGE           = 27
	WINHTTP_QUERY_AUTHORIZATION             = 28
	WINHTTP_QUERY_CONTENT_ENCODING          = 29
	WINHTTP_QUERY_FORWARDED                 = 30
	WINHTTP_QUERY_FROM                      = 31
	WINHTTP_QUERY_IF_MODIFIED_SINCE         = 32
	WINHTTP_QUERY_LOCATION                  = 33
	WINHTTP_QUERY_ORIG_URI                  = 34
	WINHTTP_QUERY_REFERER                   = 35
	WINHTTP_QUERY_RETRY_AFTER               = 36
	WINHTTP_QUERY_SERVER                    = 37
	WINHTTP_QUERY_TITLE                     = 38
	WINHTTP_QUERY_USER_AGENT                = 39
	WINHTTP_QUERY_WWW_AUTHENTICATE          = 40
	WINHTTP_QUERY_PROXY_AUTHENTICATE        = 41
	WINHTTP_QUERY_ACCEPT_RANGES             = 42
	WINHTTP_QUERY_SET_COOKIE                = 43
	WINHTTP_QUERY_COOKIE                    = 44
	WINHTTP_QUERY_REQUEST_METHOD            = 45 // special: GET/POST etc.
	WINHTTP_QUERY_REFRESH                   = 46
	WINHTTP_QUERY_CONTENT_DISPOSITION       = 47

	//
	// HTTP 1.1 defined headers
	//

	WINHTTP_QUERY_AGE                   = 48
	WINHTTP_QUERY_CACHE_CONTROL         = 49
	WINHTTP_QUERY_CONTENT_BASE          = 50
	WINHTTP_QUERY_CONTENT_LOCATION      = 51
	WINHTTP_QUERY_CONTENT_MD5           = 52
	WINHTTP_QUERY_CONTENT_RANGE         = 53
	WINHTTP_QUERY_ETAG                  = 54
	WINHTTP_QUERY_HOST                  = 55
	WINHTTP_QUERY_IF_MATCH              = 56
	WINHTTP_QUERY_IF_NONE_MATCH         = 57
	WINHTTP_QUERY_IF_RANGE              = 58
	WINHTTP_QUERY_IF_UNMODIFIED_SINCE   = 59
	WINHTTP_QUERY_MAX_FORWARDS          = 60
	WINHTTP_QUERY_PROXY_AUTHORIZATION   = 61
	WINHTTP_QUERY_RANGE                 = 62
	WINHTTP_QUERY_TRANSFER_ENCODING     = 63
	WINHTTP_QUERY_UPGRADE               = 64
	WINHTTP_QUERY_VARY                  = 65
	WINHTTP_QUERY_VIA                   = 66
	WINHTTP_QUERY_WARNING               = 67
	WINHTTP_QUERY_EXPECT                = 68
	WINHTTP_QUERY_PROXY_CONNECTION      = 69
	WINHTTP_QUERY_UNLESS_MODIFIED_SINCE = 70

	WINHTTP_HEADER_NAME_BY_INDEX = 0
	WINHTTP_NO_OUTPUT_BUFFER     = 0
	WINHTTP_NO_HEADER_INDEX      = 0
)
