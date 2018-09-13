// Package firestore provides access to the Cloud Firestore API.
//
// This package is DEPRECATED. Use package cloud.google.com/go/firestore instead.
//
// See https://cloud.google.com/firestore
//
// Usage example:
//
//   import "google.golang.org/api/firestore/v1beta1"
//   ...
//   firestoreService, err := firestore.New(oauthHttpClient)
package firestore // import "google.golang.org/api/firestore/v1beta1"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "firestore:v1beta1"
const apiName = "firestore"
const apiVersion = "v1beta1"
const basePath = "https://firestore.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View and manage your Google Cloud Datastore data
	DatastoreScope = "https://www.googleapis.com/auth/datastore"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Databases = NewProjectsDatabasesService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Databases *ProjectsDatabasesService
}

func NewProjectsDatabasesService(s *Service) *ProjectsDatabasesService {
	rs := &ProjectsDatabasesService{s: s}
	rs.Documents = NewProjectsDatabasesDocumentsService(s)
	rs.Indexes = NewProjectsDatabasesIndexesService(s)
	return rs
}

type ProjectsDatabasesService struct {
	s *Service

	Documents *ProjectsDatabasesDocumentsService

	Indexes *ProjectsDatabasesIndexesService
}

func NewProjectsDatabasesDocumentsService(s *Service) *ProjectsDatabasesDocumentsService {
	rs := &ProjectsDatabasesDocumentsService{s: s}
	return rs
}

type ProjectsDatabasesDocumentsService struct {
	s *Service
}

func NewProjectsDatabasesIndexesService(s *Service) *ProjectsDatabasesIndexesService {
	rs := &ProjectsDatabasesIndexesService{s: s}
	return rs
}

type ProjectsDatabasesIndexesService struct {
	s *Service
}

// ArrayValue: An array value.
type ArrayValue struct {
	// Values: Values in the array.
	Values []*Value `json:"values,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Values") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Values") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ArrayValue) MarshalJSON() ([]byte, error) {
	type NoMethod ArrayValue
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BatchGetDocumentsRequest: The request for
// Firestore.BatchGetDocuments.
type BatchGetDocumentsRequest struct {
	// Documents: The names of the documents to retrieve. In the
	// format:
	// `projects/{project_id}/databases/{database_id}/documents/{docu
	// ment_path}`.
	// The request will fail if any of the document is not a child resource
	// of the
	// given `database`. Duplicate names will be elided.
	Documents []string `json:"documents,omitempty"`

	// Mask: The fields to return. If not set, returns all fields.
	//
	// If a document has a field that is not present in this mask, that
	// field will
	// not be returned in the response.
	Mask *DocumentMask `json:"mask,omitempty"`

	// NewTransaction: Starts a new transaction and reads the
	// documents.
	// Defaults to a read-only transaction.
	// The new transaction ID will be returned as the first response in
	// the
	// stream.
	NewTransaction *TransactionOptions `json:"newTransaction,omitempty"`

	// ReadTime: Reads documents as they were at the given time.
	// This may not be older than 60 seconds.
	ReadTime string `json:"readTime,omitempty"`

	// Transaction: Reads documents in a transaction.
	Transaction string `json:"transaction,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Documents") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Documents") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BatchGetDocumentsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod BatchGetDocumentsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BatchGetDocumentsResponse: The streamed response for
// Firestore.BatchGetDocuments.
type BatchGetDocumentsResponse struct {
	// Found: A document that was requested.
	Found *Document `json:"found,omitempty"`

	// Missing: A document name that was requested but does not exist. In
	// the
	// format:
	// `projects/{project_id}/databases/{database_id}/documents/{docu
	// ment_path}`.
	Missing string `json:"missing,omitempty"`

	// ReadTime: The time at which the document was read.
	// This may be monotically increasing, in this case the previous
	// documents in
	// the result stream are guaranteed not to have changed between
	// their
	// read_time and this one.
	ReadTime string `json:"readTime,omitempty"`

	// Transaction: The transaction that was started as part of this
	// request.
	// Will only be set in the first response, and only
	// if
	// BatchGetDocumentsRequest.new_transaction was set in the request.
	Transaction string `json:"transaction,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Found") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Found") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BatchGetDocumentsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod BatchGetDocumentsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BeginTransactionRequest: The request for Firestore.BeginTransaction.
type BeginTransactionRequest struct {
	// Options: The options for the transaction.
	// Defaults to a read-write transaction.
	Options *TransactionOptions `json:"options,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Options") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Options") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BeginTransactionRequest) MarshalJSON() ([]byte, error) {
	type NoMethod BeginTransactionRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BeginTransactionResponse: The response for
// Firestore.BeginTransaction.
type BeginTransactionResponse struct {
	// Transaction: The transaction that was started.
	Transaction string `json:"transaction,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Transaction") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Transaction") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BeginTransactionResponse) MarshalJSON() ([]byte, error) {
	type NoMethod BeginTransactionResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CollectionSelector: A selection of a collection, such as `messages as
// m1`.
type CollectionSelector struct {
	// AllDescendants: When false, selects only collections that are
	// immediate children of
	// the `parent` specified in the containing `RunQueryRequest`.
	// When true, selects all descendant collections.
	AllDescendants bool `json:"allDescendants,omitempty"`

	// CollectionId: The collection ID.
	// When set, selects only collections with this ID.
	CollectionId string `json:"collectionId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AllDescendants") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AllDescendants") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *CollectionSelector) MarshalJSON() ([]byte, error) {
	type NoMethod CollectionSelector
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CommitRequest: The request for Firestore.Commit.
type CommitRequest struct {
	// Transaction: If set, applies all writes in this transaction, and
	// commits it.
	Transaction string `json:"transaction,omitempty"`

	// Writes: The writes to apply.
	//
	// Always executed atomically and in order.
	Writes []*Write `json:"writes,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Transaction") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Transaction") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CommitRequest) MarshalJSON() ([]byte, error) {
	type NoMethod CommitRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CommitResponse: The response for Firestore.Commit.
type CommitResponse struct {
	// CommitTime: The time at which the commit occurred.
	CommitTime string `json:"commitTime,omitempty"`

	// WriteResults: The result of applying the writes.
	//
	// This i-th write result corresponds to the i-th write in the
	// request.
	WriteResults []*WriteResult `json:"writeResults,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CommitTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CommitTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CommitResponse) MarshalJSON() ([]byte, error) {
	type NoMethod CommitResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CompositeFilter: A filter that merges multiple other filters using
// the given operator.
type CompositeFilter struct {
	// Filters: The list of filters to combine.
	// Must contain at least one filter.
	Filters []*Filter `json:"filters,omitempty"`

	// Op: The operator for combining multiple filters.
	//
	// Possible values:
	//   "OPERATOR_UNSPECIFIED" - Unspecified. This value must not be used.
	//   "AND" - The results are required to satisfy each of the combined
	// filters.
	Op string `json:"op,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Filters") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Filters") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CompositeFilter) MarshalJSON() ([]byte, error) {
	type NoMethod CompositeFilter
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Cursor: A position in a query result set.
type Cursor struct {
	// Before: If the position is just before or just after the given
	// values, relative
	// to the sort order defined by the query.
	Before bool `json:"before,omitempty"`

	// Values: The values that represent a position, in the order they
	// appear in
	// the order by clause of a query.
	//
	// Can contain fewer values than specified in the order by clause.
	Values []*Value `json:"values,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Before") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Before") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Cursor) MarshalJSON() ([]byte, error) {
	type NoMethod Cursor
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Document: A Firestore document.
//
// Must not exceed 1 MiB - 4 bytes.
type Document struct {
	// CreateTime: Output only. The time at which the document was
	// created.
	//
	// This value increases monotonically when a document is deleted
	// then
	// recreated. It can also be compared to values from other documents
	// and
	// the `read_time` of a query.
	CreateTime string `json:"createTime,omitempty"`

	// Fields: The document's fields.
	//
	// The map keys represent field names.
	//
	// A simple field name contains only characters `a` to `z`, `A` to
	// `Z`,
	// `0` to `9`, or `_`, and must not start with `0` to `9`. For
	// example,
	// `foo_bar_17`.
	//
	// Field names matching the regular expression `__.*__` are reserved.
	// Reserved
	// field names are forbidden except in certain documented contexts. The
	// map
	// keys, represented as UTF-8, must not exceed 1,500 bytes and cannot
	// be
	// empty.
	//
	// Field paths may be used in other contexts to refer to structured
	// fields
	// defined here. For `map_value`, the field path is represented by the
	// simple
	// or quoted field names of the containing fields, delimited by `.`.
	// For
	// example, the structured field
	// "foo" : { map_value: { "x&y" : { string_value: "hello" }}}` would
	// be
	// represented by the field path `foo.x&y`.
	//
	// Within a field path, a quoted field name starts and ends with `` ` ``
	// and
	// may contain any character. Some characters, including `` ` ``, must
	// be
	// escaped using a `\`. For example, `` `x&y` `` represents `x&y` and
	// `` `bak\`tik` `` represents `` bak`tik ``.
	Fields map[string]Value `json:"fields,omitempty"`

	// Name: The resource name of the document, for
	// example
	// `projects/{project_id}/databases/{database_id}/documents/{docu
	// ment_path}`.
	Name string `json:"name,omitempty"`

	// UpdateTime: Output only. The time at which the document was last
	// changed.
	//
	// This value is initially set to the `create_time` then
	// increases
	// monotonically with each change to the document. It can also
	// be
	// compared to values from other documents and the `read_time` of a
	// query.
	UpdateTime string `json:"updateTime,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreateTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Document) MarshalJSON() ([]byte, error) {
	type NoMethod Document
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DocumentChange: A Document has changed.
//
// May be the result of multiple writes, including deletes,
// that
// ultimately resulted in a new value for the Document.
//
// Multiple DocumentChange messages may be returned for the same
// logical
// change, if multiple targets are affected.
type DocumentChange struct {
	// Document: The new state of the Document.
	//
	// If `mask` is set, contains only fields that were updated or added.
	Document *Document `json:"document,omitempty"`

	// RemovedTargetIds: A set of target IDs for targets that no longer
	// match this document.
	RemovedTargetIds []int64 `json:"removedTargetIds,omitempty"`

	// TargetIds: A set of target IDs of targets that match this document.
	TargetIds []int64 `json:"targetIds,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Document") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Document") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DocumentChange) MarshalJSON() ([]byte, error) {
	type NoMethod DocumentChange
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DocumentDelete: A Document has been deleted.
//
// May be the result of multiple writes, including updates, the
// last of which deleted the Document.
//
// Multiple DocumentDelete messages may be returned for the same
// logical
// delete, if multiple targets are affected.
type DocumentDelete struct {
	// Document: The resource name of the Document that was deleted.
	Document string `json:"document,omitempty"`

	// ReadTime: The read timestamp at which the delete was
	// observed.
	//
	// Greater or equal to the `commit_time` of the delete.
	ReadTime string `json:"readTime,omitempty"`

	// RemovedTargetIds: A set of target IDs for targets that previously
	// matched this entity.
	RemovedTargetIds []int64 `json:"removedTargetIds,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Document") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Document") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DocumentDelete) MarshalJSON() ([]byte, error) {
	type NoMethod DocumentDelete
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DocumentMask: A set of field paths on a document.
// Used to restrict a get or update operation on a document to a subset
// of its
// fields.
// This is different from standard field masks, as this is always scoped
// to a
// Document, and takes in account the dynamic nature of Value.
type DocumentMask struct {
	// FieldPaths: The list of field paths in the mask. See Document.fields
	// for a field
	// path syntax reference.
	FieldPaths []string `json:"fieldPaths,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FieldPaths") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FieldPaths") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DocumentMask) MarshalJSON() ([]byte, error) {
	type NoMethod DocumentMask
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DocumentRemove: A Document has been removed from the view of the
// targets.
//
// Sent if the document is no longer relevant to a target and is out of
// view.
// Can be sent instead of a DocumentDelete or a DocumentChange if the
// server
// can not send the new value of the document.
//
// Multiple DocumentRemove messages may be returned for the same
// logical
// write or delete, if multiple targets are affected.
type DocumentRemove struct {
	// Document: The resource name of the Document that has gone out of
	// view.
	Document string `json:"document,omitempty"`

	// ReadTime: The read timestamp at which the remove was
	// observed.
	//
	// Greater or equal to the `commit_time` of the change/delete/remove.
	ReadTime string `json:"readTime,omitempty"`

	// RemovedTargetIds: A set of target IDs for targets that previously
	// matched this document.
	RemovedTargetIds []int64 `json:"removedTargetIds,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Document") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Document") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DocumentRemove) MarshalJSON() ([]byte, error) {
	type NoMethod DocumentRemove
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DocumentTransform: A transformation of a document.
type DocumentTransform struct {
	// Document: The name of the document to transform.
	Document string `json:"document,omitempty"`

	// FieldTransforms: The list of transformations to apply to the fields
	// of the document, in
	// order.
	// This must not be empty.
	FieldTransforms []*FieldTransform `json:"fieldTransforms,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Document") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Document") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DocumentTransform) MarshalJSON() ([]byte, error) {
	type NoMethod DocumentTransform
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DocumentsTarget: A target specified by a set of documents names.
type DocumentsTarget struct {
	// Documents: The names of the documents to retrieve. In the
	// format:
	// `projects/{project_id}/databases/{database_id}/documents/{docu
	// ment_path}`.
	// The request will fail if any of the document is not a child resource
	// of
	// the given `database`. Duplicate names will be elided.
	Documents []string `json:"documents,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Documents") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Documents") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DocumentsTarget) MarshalJSON() ([]byte, error) {
	type NoMethod DocumentsTarget
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// ExistenceFilter: A digest of all the documents that match a given
// target.
type ExistenceFilter struct {
	// Count: The total count of documents that match target_id.
	//
	// If different from the count of documents in the client that match,
	// the
	// client must manually determine which documents no longer match the
	// target.
	Count int64 `json:"count,omitempty"`

	// TargetId: The target ID to which this filter applies.
	TargetId int64 `json:"targetId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Count") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Count") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ExistenceFilter) MarshalJSON() ([]byte, error) {
	type NoMethod ExistenceFilter
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// FieldFilter: A filter on a specific field.
type FieldFilter struct {
	// Field: The field to filter by.
	Field *FieldReference `json:"field,omitempty"`

	// Op: The operator to filter by.
	//
	// Possible values:
	//   "OPERATOR_UNSPECIFIED" - Unspecified. This value must not be used.
	//   "LESS_THAN" - Less than. Requires that the field come first in
	// `order_by`.
	//   "LESS_THAN_OR_EQUAL" - Less than or equal. Requires that the field
	// come first in `order_by`.
	//   "GREATER_THAN" - Greater than. Requires that the field come first
	// in `order_by`.
	//   "GREATER_THAN_OR_EQUAL" - Greater than or equal. Requires that the
	// field come first in
	// `order_by`.
	//   "EQUAL" - Equal.
	//   "ARRAY_CONTAINS" - Contains. Requires that the field is an array.
	Op string `json:"op,omitempty"`

	// Value: The value to compare to.
	Value *Value `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Field") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Field") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *FieldFilter) MarshalJSON() ([]byte, error) {
	type NoMethod FieldFilter
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// FieldReference: A reference to a field, such as `max(messages.time)
// as max_time`.
type FieldReference struct {
	FieldPath string `json:"fieldPath,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FieldPath") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FieldPath") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *FieldReference) MarshalJSON() ([]byte, error) {
	type NoMethod FieldReference
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// FieldTransform: A transformation of a field of the document.
type FieldTransform struct {
	// AppendMissingElements: Append the given elements in order if they are
	// not already present in
	// the current field value.
	// If the field is not an array, or if the field does not yet exist, it
	// is
	// first set to the empty array.
	//
	// Equivalent numbers of different types (e.g. 3L and 3.0)
	// are
	// considered equal when checking if a value is missing.
	// NaN is equal to NaN, and Null is equal to Null.
	// If the input contains multiple equivalent values, only the first
	// will
	// be considered.
	//
	// The corresponding transform_result will be the null value.
	AppendMissingElements *ArrayValue `json:"appendMissingElements,omitempty"`

	// FieldPath: The path of the field. See Document.fields for the field
	// path syntax
	// reference.
	FieldPath string `json:"fieldPath,omitempty"`

	// RemoveAllFromArray: Remove all of the given elements from the array
	// in the field.
	// If the field is not an array, or if the field does not yet exist, it
	// is
	// set to the empty array.
	//
	// Equivalent numbers of the different types (e.g. 3L and 3.0)
	// are
	// considered equal when deciding whether an element should be
	// removed.
	// NaN is equal to NaN, and Null is equal to Null.
	// This will remove all equivalent values if there are duplicates.
	//
	// The corresponding transform_result will be the null value.
	RemoveAllFromArray *ArrayValue `json:"removeAllFromArray,omitempty"`

	// SetToServerValue: Sets the field to the given server value.
	//
	// Possible values:
	//   "SERVER_VALUE_UNSPECIFIED" - Unspecified. This value must not be
	// used.
	//   "REQUEST_TIME" - The time at which the server processed the
	// request, with millisecond
	// precision.
	SetToServerValue string `json:"setToServerValue,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "AppendMissingElements") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AppendMissingElements") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *FieldTransform) MarshalJSON() ([]byte, error) {
	type NoMethod FieldTransform
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Filter: A filter.
type Filter struct {
	// CompositeFilter: A composite filter.
	CompositeFilter *CompositeFilter `json:"compositeFilter,omitempty"`

	// FieldFilter: A filter on a document field.
	FieldFilter *FieldFilter `json:"fieldFilter,omitempty"`

	// UnaryFilter: A filter that takes exactly one argument.
	UnaryFilter *UnaryFilter `json:"unaryFilter,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CompositeFilter") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CompositeFilter") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Filter) MarshalJSON() ([]byte, error) {
	type NoMethod Filter
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleFirestoreAdminV1beta1ExportDocumentsMetadata: Metadata for
// ExportDocuments operations.
type GoogleFirestoreAdminV1beta1ExportDocumentsMetadata struct {
	// CollectionIds: Which collection ids are being exported.
	CollectionIds []string `json:"collectionIds,omitempty"`

	// EndTime: The time the operation ended, either successfully or
	// otherwise. Unset if
	// the operation is still active.
	EndTime string `json:"endTime,omitempty"`

	// OperationState: The state of the export operation.
	//
	// Possible values:
	//   "STATE_UNSPECIFIED" - Unspecified.
	//   "INITIALIZING" - Request is being prepared for processing.
	//   "PROCESSING" - Request is actively being processed.
	//   "CANCELLING" - Request is in the process of being cancelled after
	// user called
	// google.longrunning.Operations.CancelOperation on the operation.
	//   "FINALIZING" - Request has been processed and is in its
	// finalization stage.
	//   "SUCCESSFUL" - Request has completed successfully.
	//   "FAILED" - Request has finished being processed, but encountered an
	// error.
	//   "CANCELLED" - Request has finished being cancelled after user
	// called
	// google.longrunning.Operations.CancelOperation.
	OperationState string `json:"operationState,omitempty"`

	// OutputUriPrefix: Where the entities are being exported to.
	OutputUriPrefix string `json:"outputUriPrefix,omitempty"`

	// ProgressBytes: An estimate of the number of bytes processed.
	ProgressBytes *GoogleFirestoreAdminV1beta1Progress `json:"progressBytes,omitempty"`

	// ProgressDocuments: An estimate of the number of documents processed.
	ProgressDocuments *GoogleFirestoreAdminV1beta1Progress `json:"progressDocuments,omitempty"`

	// StartTime: The time that work began on the operation.
	StartTime string `json:"startTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CollectionIds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CollectionIds") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleFirestoreAdminV1beta1ExportDocumentsMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirestoreAdminV1beta1ExportDocumentsMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleFirestoreAdminV1beta1ExportDocumentsRequest: The request for
// FirestoreAdmin.ExportDocuments.
type GoogleFirestoreAdminV1beta1ExportDocumentsRequest struct {
	// CollectionIds: Which collection ids to export. Unspecified means all
	// collections.
	CollectionIds []string `json:"collectionIds,omitempty"`

	// OutputUriPrefix: The output URI. Currently only supports Google Cloud
	// Storage URIs of the
	// form: `gs://BUCKET_NAME[/NAMESPACE_PATH]`, where `BUCKET_NAME` is the
	// name
	// of the Google Cloud Storage bucket and `NAMESPACE_PATH` is an
	// optional
	// Google Cloud Storage namespace path. When
	// choosing a name, be sure to consider Google Cloud Storage
	// naming
	// guidelines: https://cloud.google.com/storage/docs/naming.
	// If the URI is a bucket (without a namespace path), a prefix will
	// be
	// generated based on the start time.
	OutputUriPrefix string `json:"outputUriPrefix,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CollectionIds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CollectionIds") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleFirestoreAdminV1beta1ExportDocumentsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirestoreAdminV1beta1ExportDocumentsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleFirestoreAdminV1beta1ExportDocumentsResponse: Returned in the
// google.longrunning.Operation response field.
type GoogleFirestoreAdminV1beta1ExportDocumentsResponse struct {
	// OutputUriPrefix: Location of the output files. This can be used to
	// begin an import
	// into Cloud Firestore (this project or another project) after the
	// operation
	// completes successfully.
	OutputUriPrefix string `json:"outputUriPrefix,omitempty"`

	// ForceSendFields is a list of field names (e.g. "OutputUriPrefix") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "OutputUriPrefix") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleFirestoreAdminV1beta1ExportDocumentsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirestoreAdminV1beta1ExportDocumentsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleFirestoreAdminV1beta1ImportDocumentsMetadata: Metadata for
// ImportDocuments operations.
type GoogleFirestoreAdminV1beta1ImportDocumentsMetadata struct {
	// CollectionIds: Which collection ids are being imported.
	CollectionIds []string `json:"collectionIds,omitempty"`

	// EndTime: The time the operation ended, either successfully or
	// otherwise. Unset if
	// the operation is still active.
	EndTime string `json:"endTime,omitempty"`

	// InputUriPrefix: The location of the documents being imported.
	InputUriPrefix string `json:"inputUriPrefix,omitempty"`

	// OperationState: The state of the import operation.
	//
	// Possible values:
	//   "STATE_UNSPECIFIED" - Unspecified.
	//   "INITIALIZING" - Request is being prepared for processing.
	//   "PROCESSING" - Request is actively being processed.
	//   "CANCELLING" - Request is in the process of being cancelled after
	// user called
	// google.longrunning.Operations.CancelOperation on the operation.
	//   "FINALIZING" - Request has been processed and is in its
	// finalization stage.
	//   "SUCCESSFUL" - Request has completed successfully.
	//   "FAILED" - Request has finished being processed, but encountered an
	// error.
	//   "CANCELLED" - Request has finished being cancelled after user
	// called
	// google.longrunning.Operations.CancelOperation.
	OperationState string `json:"operationState,omitempty"`

	// ProgressBytes: An estimate of the number of bytes processed.
	ProgressBytes *GoogleFirestoreAdminV1beta1Progress `json:"progressBytes,omitempty"`

	// ProgressDocuments: An estimate of the number of documents processed.
	ProgressDocuments *GoogleFirestoreAdminV1beta1Progress `json:"progressDocuments,omitempty"`

	// StartTime: The time that work began on the operation.
	StartTime string `json:"startTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CollectionIds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CollectionIds") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleFirestoreAdminV1beta1ImportDocumentsMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirestoreAdminV1beta1ImportDocumentsMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleFirestoreAdminV1beta1ImportDocumentsRequest: The request for
// FirestoreAdmin.ImportDocuments.
type GoogleFirestoreAdminV1beta1ImportDocumentsRequest struct {
	// CollectionIds: Which collection ids to import. Unspecified means all
	// collections included
	// in the import.
	CollectionIds []string `json:"collectionIds,omitempty"`

	// InputUriPrefix: Location of the exported files.
	// This must match the output_uri_prefix of an ExportDocumentsResponse
	// from
	// an export that has completed
	// successfully.
	// See:
	// google.firestore.admin.v1beta1.ExportDocumentsRespo
	// nse.output_uri_prefix.
	InputUriPrefix string `json:"inputUriPrefix,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CollectionIds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CollectionIds") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleFirestoreAdminV1beta1ImportDocumentsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirestoreAdminV1beta1ImportDocumentsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleFirestoreAdminV1beta1Index: An index definition.
type GoogleFirestoreAdminV1beta1Index struct {
	// CollectionId: The collection ID to which this index applies.
	// Required.
	CollectionId string `json:"collectionId,omitempty"`

	// Fields: The fields to index.
	Fields []*GoogleFirestoreAdminV1beta1IndexField `json:"fields,omitempty"`

	// Name: The resource name of the index.
	// Output only.
	Name string `json:"name,omitempty"`

	// State: The state of the index.
	// Output only.
	//
	// Possible values:
	//   "STATE_UNSPECIFIED" - The state is unspecified.
	//   "CREATING" - The index is being created.
	// There is an active long-running operation for the index.
	// The index is updated when writing a document.
	// Some index data may exist.
	//   "READY" - The index is ready to be used.
	// The index is updated when writing a document.
	// The index is fully populated from all stored documents it applies to.
	//   "ERROR" - The index was being created, but something went
	// wrong.
	// There is no active long-running operation for the index,
	// and the most recently finished long-running operation failed.
	// The index is not updated when writing a document.
	// Some index data may exist.
	State string `json:"state,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CollectionId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CollectionId") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleFirestoreAdminV1beta1Index) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirestoreAdminV1beta1Index
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleFirestoreAdminV1beta1IndexField: A field of an index.
type GoogleFirestoreAdminV1beta1IndexField struct {
	// FieldPath: The path of the field. Must match the field path
	// specification described
	// by google.firestore.v1beta1.Document.fields.
	// Special field path `__name__` may be used by itself or at the end of
	// a
	// path. `__type__` may be used only at the end of path.
	FieldPath string `json:"fieldPath,omitempty"`

	// Mode: The field's mode.
	//
	// Possible values:
	//   "MODE_UNSPECIFIED" - The mode is unspecified.
	//   "ASCENDING" - The field's values are indexed so as to support
	// sequencing in
	// ascending order and also query by <, >, <=, >=, and =.
	//   "DESCENDING" - The field's values are indexed so as to support
	// sequencing in
	// descending order and also query by <, >, <=, >=, and =.
	//   "ARRAY_CONTAINS" - The field's array values are indexed so as to
	// support membership using
	// ARRAY_CONTAINS queries.
	Mode string `json:"mode,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FieldPath") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FieldPath") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleFirestoreAdminV1beta1IndexField) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirestoreAdminV1beta1IndexField
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleFirestoreAdminV1beta1IndexOperationMetadata: Metadata for index
// operations. This metadata populates
// the metadata field of google.longrunning.Operation.
type GoogleFirestoreAdminV1beta1IndexOperationMetadata struct {
	// Cancelled: True if the [google.longrunning.Operation] was cancelled.
	// If the
	// cancellation is in progress, cancelled will be true
	// but
	// google.longrunning.Operation.done will be false.
	Cancelled bool `json:"cancelled,omitempty"`

	// DocumentProgress: Progress of the existing operation, measured in
	// number of documents.
	DocumentProgress *GoogleFirestoreAdminV1beta1Progress `json:"documentProgress,omitempty"`

	// EndTime: The time the operation ended, either successfully or
	// otherwise. Unset if
	// the operation is still active.
	EndTime string `json:"endTime,omitempty"`

	// Index: The index resource that this operation is acting on. For
	// example:
	// `projects/{project_id}/databases/{database_id}/indexes/{index
	// _id}`
	Index string `json:"index,omitempty"`

	// OperationType: The type of index operation.
	//
	// Possible values:
	//   "OPERATION_TYPE_UNSPECIFIED" - Unspecified. Never set by server.
	//   "CREATING_INDEX" - The operation is creating the index. Initiated
	// by a `CreateIndex` call.
	OperationType string `json:"operationType,omitempty"`

	// StartTime: The time that work began on the operation.
	StartTime string `json:"startTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Cancelled") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Cancelled") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleFirestoreAdminV1beta1IndexOperationMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirestoreAdminV1beta1IndexOperationMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleFirestoreAdminV1beta1ListIndexesResponse: The response for
// FirestoreAdmin.ListIndexes.
type GoogleFirestoreAdminV1beta1ListIndexesResponse struct {
	// Indexes: The indexes.
	Indexes []*GoogleFirestoreAdminV1beta1Index `json:"indexes,omitempty"`

	// NextPageToken: The standard List next-page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Indexes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Indexes") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleFirestoreAdminV1beta1ListIndexesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirestoreAdminV1beta1ListIndexesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleFirestoreAdminV1beta1LocationMetadata: The metadata message for
// google.cloud.location.Location.metadata.
type GoogleFirestoreAdminV1beta1LocationMetadata struct {
}

// GoogleFirestoreAdminV1beta1Progress: Measures the progress of a
// particular metric.
type GoogleFirestoreAdminV1beta1Progress struct {
	// WorkCompleted: An estimate of how much work has been completed. Note
	// that this may be
	// greater than `work_estimated`.
	WorkCompleted int64 `json:"workCompleted,omitempty,string"`

	// WorkEstimated: An estimate of how much work needs to be performed.
	// Zero if the
	// work estimate is unavailable. May change as work progresses.
	WorkEstimated int64 `json:"workEstimated,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "WorkCompleted") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "WorkCompleted") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleFirestoreAdminV1beta1Progress) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleFirestoreAdminV1beta1Progress
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleLongrunningOperation: This resource represents a long-running
// operation that is the result of a
// network API call.
type GoogleLongrunningOperation struct {
	// Done: If the value is `false`, it means the operation is still in
	// progress.
	// If `true`, the operation is completed, and either `error` or
	// `response` is
	// available.
	Done bool `json:"done,omitempty"`

	// Error: The error result of the operation in case of failure or
	// cancellation.
	Error *Status `json:"error,omitempty"`

	// Metadata: Service-specific metadata associated with the operation.
	// It typically
	// contains progress information and common metadata such as create
	// time.
	// Some services might not provide such metadata.  Any method that
	// returns a
	// long-running operation should document the metadata type, if any.
	Metadata googleapi.RawMessage `json:"metadata,omitempty"`

	// Name: The server-assigned name, which is only unique within the same
	// service that
	// originally returns it. If you use the default HTTP mapping,
	// the
	// `name` should have the format of `operations/some/unique/name`.
	Name string `json:"name,omitempty"`

	// Response: The normal response of the operation in case of success.
	// If the original
	// method returns no data on success, such as `Delete`, the response
	// is
	// `google.protobuf.Empty`.  If the original method is
	// standard
	// `Get`/`Create`/`Update`, the response should be the resource.  For
	// other
	// methods, the response should have the type `XxxResponse`, where
	// `Xxx`
	// is the original method name.  For example, if the original method
	// name
	// is `TakeSnapshot()`, the inferred response type
	// is
	// `TakeSnapshotResponse`.
	Response googleapi.RawMessage `json:"response,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Done") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Done") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleLongrunningOperation) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleLongrunningOperation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LatLng: An object representing a latitude/longitude pair. This is
// expressed as a pair
// of doubles representing degrees latitude and degrees longitude.
// Unless
// specified otherwise, this must conform to the
// <a
// href="http://www.unoosa.org/pdf/icg/2012/template/WGS_84.pdf">WGS84
// st
// andard</a>. Values must be within normalized ranges.
type LatLng struct {
	// Latitude: The latitude in degrees. It must be in the range [-90.0,
	// +90.0].
	Latitude float64 `json:"latitude,omitempty"`

	// Longitude: The longitude in degrees. It must be in the range [-180.0,
	// +180.0].
	Longitude float64 `json:"longitude,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Latitude") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Latitude") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *LatLng) MarshalJSON() ([]byte, error) {
	type NoMethod LatLng
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *LatLng) UnmarshalJSON(data []byte) error {
	type NoMethod LatLng
	var s1 struct {
		Latitude  gensupport.JSONFloat64 `json:"latitude"`
		Longitude gensupport.JSONFloat64 `json:"longitude"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Latitude = float64(s1.Latitude)
	s.Longitude = float64(s1.Longitude)
	return nil
}

// ListCollectionIdsRequest: The request for
// Firestore.ListCollectionIds.
type ListCollectionIdsRequest struct {
	// PageSize: The maximum number of results to return.
	PageSize int64 `json:"pageSize,omitempty"`

	// PageToken: A page token. Must be a value
	// from
	// ListCollectionIdsResponse.
	PageToken string `json:"pageToken,omitempty"`

	// ForceSendFields is a list of field names (e.g. "PageSize") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "PageSize") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListCollectionIdsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod ListCollectionIdsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListCollectionIdsResponse: The response from
// Firestore.ListCollectionIds.
type ListCollectionIdsResponse struct {
	// CollectionIds: The collection ids.
	CollectionIds []string `json:"collectionIds,omitempty"`

	// NextPageToken: A page token that may be used to continue the list.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CollectionIds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CollectionIds") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListCollectionIdsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListCollectionIdsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListDocumentsResponse: The response for Firestore.ListDocuments.
type ListDocumentsResponse struct {
	// Documents: The Documents found.
	Documents []*Document `json:"documents,omitempty"`

	// NextPageToken: The next page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Documents") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Documents") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListDocumentsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListDocumentsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListenRequest: A request for Firestore.Listen
type ListenRequest struct {
	// AddTarget: A target to add to this stream.
	AddTarget *Target `json:"addTarget,omitempty"`

	// Labels: Labels associated with this target change.
	Labels map[string]string `json:"labels,omitempty"`

	// RemoveTarget: The ID of a target to remove from this stream.
	RemoveTarget int64 `json:"removeTarget,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AddTarget") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AddTarget") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListenRequest) MarshalJSON() ([]byte, error) {
	type NoMethod ListenRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListenResponse: The response for Firestore.Listen.
type ListenResponse struct {
	// DocumentChange: A Document has changed.
	DocumentChange *DocumentChange `json:"documentChange,omitempty"`

	// DocumentDelete: A Document has been deleted.
	DocumentDelete *DocumentDelete `json:"documentDelete,omitempty"`

	// DocumentRemove: A Document has been removed from a target (because it
	// is no longer
	// relevant to that target).
	DocumentRemove *DocumentRemove `json:"documentRemove,omitempty"`

	// Filter: A filter to apply to the set of documents previously returned
	// for the
	// given target.
	//
	// Returned when documents may have been removed from the given target,
	// but
	// the exact documents are unknown.
	Filter *ExistenceFilter `json:"filter,omitempty"`

	// TargetChange: Targets have changed.
	TargetChange *TargetChange `json:"targetChange,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "DocumentChange") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DocumentChange") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ListenResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListenResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// MapValue: A map value.
type MapValue struct {
	// Fields: The map's fields.
	//
	// The map keys represent field names. Field names matching the
	// regular
	// expression `__.*__` are reserved. Reserved field names are forbidden
	// except
	// in certain documented contexts. The map keys, represented as UTF-8,
	// must
	// not exceed 1,500 bytes and cannot be empty.
	Fields map[string]Value `json:"fields,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Fields") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Fields") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *MapValue) MarshalJSON() ([]byte, error) {
	type NoMethod MapValue
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Order: An order on a field.
type Order struct {
	// Direction: The direction to order by. Defaults to `ASCENDING`.
	//
	// Possible values:
	//   "DIRECTION_UNSPECIFIED" - Unspecified.
	//   "ASCENDING" - Ascending.
	//   "DESCENDING" - Descending.
	Direction string `json:"direction,omitempty"`

	// Field: The field to order by.
	Field *FieldReference `json:"field,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Direction") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Direction") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Order) MarshalJSON() ([]byte, error) {
	type NoMethod Order
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Precondition: A precondition on a document, used for conditional
// operations.
type Precondition struct {
	// Exists: When set to `true`, the target document must exist.
	// When set to `false`, the target document must not exist.
	Exists bool `json:"exists,omitempty"`

	// UpdateTime: When set, the target document must exist and have been
	// last updated at
	// that time.
	UpdateTime string `json:"updateTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Exists") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Exists") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Precondition) MarshalJSON() ([]byte, error) {
	type NoMethod Precondition
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Projection: The projection of document's fields to return.
type Projection struct {
	// Fields: The fields to return.
	//
	// If empty, all fields are returned. To only return the name
	// of the document, use `['__name__']`.
	Fields []*FieldReference `json:"fields,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Fields") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Fields") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Projection) MarshalJSON() ([]byte, error) {
	type NoMethod Projection
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// QueryTarget: A target specified by a query.
type QueryTarget struct {
	// Parent: The parent resource name. In the
	// format:
	// `projects/{project_id}/databases/{database_id}/documents`
	// or
	// `projects/{project_id}/databases/{database_id}/documents/{document_
	// path}`.
	// For example:
	// `projects/my-project/databases/my-database/documents`
	// or
	// `projects/my-project/databases/my-database/documents/chatrooms/my-c
	// hatroom`
	Parent string `json:"parent,omitempty"`

	// StructuredQuery: A structured query.
	StructuredQuery *StructuredQuery `json:"structuredQuery,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Parent") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Parent") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *QueryTarget) MarshalJSON() ([]byte, error) {
	type NoMethod QueryTarget
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ReadOnly: Options for a transaction that can only be used to read
// documents.
type ReadOnly struct {
	// ReadTime: Reads documents at the given time.
	// This may not be older than 60 seconds.
	ReadTime string `json:"readTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ReadTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ReadTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ReadOnly) MarshalJSON() ([]byte, error) {
	type NoMethod ReadOnly
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ReadWrite: Options for a transaction that can be used to read and
// write documents.
type ReadWrite struct {
	// RetryTransaction: An optional transaction to retry.
	RetryTransaction string `json:"retryTransaction,omitempty"`

	// ForceSendFields is a list of field names (e.g. "RetryTransaction") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "RetryTransaction") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ReadWrite) MarshalJSON() ([]byte, error) {
	type NoMethod ReadWrite
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RollbackRequest: The request for Firestore.Rollback.
type RollbackRequest struct {
	// Transaction: The transaction to roll back.
	Transaction string `json:"transaction,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Transaction") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Transaction") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RollbackRequest) MarshalJSON() ([]byte, error) {
	type NoMethod RollbackRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RunQueryRequest: The request for Firestore.RunQuery.
type RunQueryRequest struct {
	// NewTransaction: Starts a new transaction and reads the
	// documents.
	// Defaults to a read-only transaction.
	// The new transaction ID will be returned as the first response in
	// the
	// stream.
	NewTransaction *TransactionOptions `json:"newTransaction,omitempty"`

	// ReadTime: Reads documents as they were at the given time.
	// This may not be older than 60 seconds.
	ReadTime string `json:"readTime,omitempty"`

	// StructuredQuery: A structured query.
	StructuredQuery *StructuredQuery `json:"structuredQuery,omitempty"`

	// Transaction: Reads documents in a transaction.
	Transaction string `json:"transaction,omitempty"`

	// ForceSendFields is a list of field names (e.g. "NewTransaction") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NewTransaction") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *RunQueryRequest) MarshalJSON() ([]byte, error) {
	type NoMethod RunQueryRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RunQueryResponse: The response for Firestore.RunQuery.
type RunQueryResponse struct {
	// Document: A query result.
	// Not set when reporting partial progress.
	Document *Document `json:"document,omitempty"`

	// ReadTime: The time at which the document was read. This may be
	// monotonically
	// increasing; in this case, the previous documents in the result stream
	// are
	// guaranteed not to have changed between their `read_time` and this
	// one.
	//
	// If the query returns no results, a response with `read_time` and
	// no
	// `document` will be sent, and this represents the time at which the
	// query
	// was run.
	ReadTime string `json:"readTime,omitempty"`

	// SkippedResults: The number of results that have been skipped due to
	// an offset between
	// the last response and the current response.
	SkippedResults int64 `json:"skippedResults,omitempty"`

	// Transaction: The transaction that was started as part of this
	// request.
	// Can only be set in the first response, and only
	// if
	// RunQueryRequest.new_transaction was set in the request.
	// If set, no other fields will be set in this response.
	Transaction string `json:"transaction,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Document") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Document") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RunQueryResponse) MarshalJSON() ([]byte, error) {
	type NoMethod RunQueryResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Status: The `Status` type defines a logical error model that is
// suitable for different
// programming environments, including REST APIs and RPC APIs. It is
// used by
// [gRPC](https://github.com/grpc). The error model is designed to
// be:
//
// - Simple to use and understand for most users
// - Flexible enough to meet unexpected needs
//
// # Overview
//
// The `Status` message contains three pieces of data: error code, error
// message,
// and error details. The error code should be an enum value
// of
// google.rpc.Code, but it may accept additional error codes if needed.
// The
// error message should be a developer-facing English message that
// helps
// developers *understand* and *resolve* the error. If a localized
// user-facing
// error message is needed, put the localized message in the error
// details or
// localize it in the client. The optional error details may contain
// arbitrary
// information about the error. There is a predefined set of error
// detail types
// in the package `google.rpc` that can be used for common error
// conditions.
//
// # Language mapping
//
// The `Status` message is the logical representation of the error
// model, but it
// is not necessarily the actual wire format. When the `Status` message
// is
// exposed in different client libraries and different wire protocols,
// it can be
// mapped differently. For example, it will likely be mapped to some
// exceptions
// in Java, but more likely mapped to some error codes in C.
//
// # Other uses
//
// The error model and the `Status` message can be used in a variety
// of
// environments, either with or without APIs, to provide a
// consistent developer experience across different
// environments.
//
// Example uses of this error model include:
//
// - Partial errors. If a service needs to return partial errors to the
// client,
//     it may embed the `Status` in the normal response to indicate the
// partial
//     errors.
//
// - Workflow errors. A typical workflow has multiple steps. Each step
// may
//     have a `Status` message for error reporting.
//
// - Batch operations. If a client uses batch request and batch
// response, the
//     `Status` message should be used directly inside batch response,
// one for
//     each error sub-response.
//
// - Asynchronous operations. If an API call embeds asynchronous
// operation
//     results in its response, the status of those operations should
// be
//     represented directly using the `Status` message.
//
// - Logging. If some API errors are stored in logs, the message
// `Status` could
//     be used directly after any stripping needed for security/privacy
// reasons.
type Status struct {
	// Code: The status code, which should be an enum value of
	// google.rpc.Code.
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details.  There is a
	// common set of
	// message types for APIs to use.
	Details []googleapi.RawMessage `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any
	// user-facing error message should be localized and sent in
	// the
	// google.rpc.Status.details field, or localized by the client.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Status) MarshalJSON() ([]byte, error) {
	type NoMethod Status
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// StructuredQuery: A Firestore query.
type StructuredQuery struct {
	// EndAt: A end point for the query results.
	EndAt *Cursor `json:"endAt,omitempty"`

	// From: The collections to query.
	From []*CollectionSelector `json:"from,omitempty"`

	// Limit: The maximum number of results to return.
	//
	// Applies after all other constraints.
	// Must be >= 0 if specified.
	Limit int64 `json:"limit,omitempty"`

	// Offset: The number of results to skip.
	//
	// Applies before limit, but after all other constraints. Must be >= 0
	// if
	// specified.
	Offset int64 `json:"offset,omitempty"`

	// OrderBy: The order to apply to the query results.
	//
	// Firestore guarantees a stable ordering through the following rules:
	//
	//  * Any field required to appear in `order_by`, that is not already
	//    specified in `order_by`, is appended to the order in field name
	// order
	//    by default.
	//  * If an order on `__name__` is not specified, it is appended by
	// default.
	//
	// Fields are appended with the same sort direction as the last
	// order
	// specified, or 'ASCENDING' if no order was specified. For example:
	//
	//  * `SELECT * FROM Foo ORDER BY A` becomes
	//    `SELECT * FROM Foo ORDER BY A, __name__`
	//  * `SELECT * FROM Foo ORDER BY A DESC` becomes
	//    `SELECT * FROM Foo ORDER BY A DESC, __name__ DESC`
	//  * `SELECT * FROM Foo WHERE A > 1` becomes
	//    `SELECT * FROM Foo WHERE A > 1 ORDER BY A, __name__`
	OrderBy []*Order `json:"orderBy,omitempty"`

	// Select: The projection to return.
	Select *Projection `json:"select,omitempty"`

	// StartAt: A starting point for the query results.
	StartAt *Cursor `json:"startAt,omitempty"`

	// Where: The filter to apply.
	Where *Filter `json:"where,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EndAt") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EndAt") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *StructuredQuery) MarshalJSON() ([]byte, error) {
	type NoMethod StructuredQuery
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Target: A specification of a set of documents to listen to.
type Target struct {
	// Documents: A target specified by a set of document names.
	Documents *DocumentsTarget `json:"documents,omitempty"`

	// Once: If the target should be removed once it is current and
	// consistent.
	Once bool `json:"once,omitempty"`

	// Query: A target specified by a query.
	Query *QueryTarget `json:"query,omitempty"`

	// ReadTime: Start listening after a specific `read_time`.
	//
	// The client must know the state of matching documents at this time.
	ReadTime string `json:"readTime,omitempty"`

	// ResumeToken: A resume token from a prior TargetChange for an
	// identical target.
	//
	// Using a resume token with a different target is unsupported and may
	// fail.
	ResumeToken string `json:"resumeToken,omitempty"`

	// TargetId: A client provided target ID.
	//
	// If not set, the server will assign an ID for the target.
	//
	// Used for resuming a target without changing IDs. The IDs can either
	// be
	// client-assigned or be server-assigned in a previous stream. All
	// targets
	// with client provided IDs must be added before adding a target that
	// needs
	// a server-assigned id.
	TargetId int64 `json:"targetId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Documents") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Documents") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Target) MarshalJSON() ([]byte, error) {
	type NoMethod Target
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TargetChange: Targets being watched have changed.
type TargetChange struct {
	// Cause: The error that resulted in this change, if applicable.
	Cause *Status `json:"cause,omitempty"`

	// ReadTime: The consistent `read_time` for the given `target_ids`
	// (omitted when the
	// target_ids are not at a consistent snapshot).
	//
	// The stream is guaranteed to send a `read_time` with `target_ids`
	// empty
	// whenever the entire stream reaches a new consistent snapshot.
	// ADD,
	// CURRENT, and RESET messages are guaranteed to (eventually) result in
	// a
	// new consistent snapshot (while NO_CHANGE and REMOVE messages are
	// not).
	//
	// For a given stream, `read_time` is guaranteed to be
	// monotonically
	// increasing.
	ReadTime string `json:"readTime,omitempty"`

	// ResumeToken: A token that can be used to resume the stream for the
	// given `target_ids`,
	// or all targets if `target_ids` is empty.
	//
	// Not set on every target change.
	ResumeToken string `json:"resumeToken,omitempty"`

	// TargetChangeType: The type of change that occurred.
	//
	// Possible values:
	//   "NO_CHANGE" - No change has occurred. Used only to send an updated
	// `resume_token`.
	//   "ADD" - The targets have been added.
	//   "REMOVE" - The targets have been removed.
	//   "CURRENT" - The targets reflect all changes committed before the
	// targets were added
	// to the stream.
	//
	// This will be sent after or with a `read_time` that is greater than
	// or
	// equal to the time at which the targets were added.
	//
	// Listeners can wait for this change if read-after-write semantics
	// are desired.
	//   "RESET" - The targets have been reset, and a new initial state for
	// the targets
	// will be returned in subsequent changes.
	//
	// After the initial state is complete, `CURRENT` will be returned
	// even
	// if the target was previously indicated to be `CURRENT`.
	TargetChangeType string `json:"targetChangeType,omitempty"`

	// TargetIds: The target IDs of targets that have changed.
	//
	// If empty, the change applies to all targets.
	//
	// For `target_change_type=ADD`, the order of the target IDs matches the
	// order
	// of the requests to add the targets. This allows clients to
	// unambiguously
	// associate server-assigned target IDs with added targets.
	//
	// For other states, the order of the target IDs is not defined.
	TargetIds []int64 `json:"targetIds,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Cause") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Cause") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TargetChange) MarshalJSON() ([]byte, error) {
	type NoMethod TargetChange
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TransactionOptions: Options for creating a new transaction.
type TransactionOptions struct {
	// ReadOnly: The transaction can only be used for read operations.
	ReadOnly *ReadOnly `json:"readOnly,omitempty"`

	// ReadWrite: The transaction can be used for both read and write
	// operations.
	ReadWrite *ReadWrite `json:"readWrite,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ReadOnly") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ReadOnly") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TransactionOptions) MarshalJSON() ([]byte, error) {
	type NoMethod TransactionOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UnaryFilter: A filter with a single operand.
type UnaryFilter struct {
	// Field: The field to which to apply the operator.
	Field *FieldReference `json:"field,omitempty"`

	// Op: The unary operator to apply.
	//
	// Possible values:
	//   "OPERATOR_UNSPECIFIED" - Unspecified. This value must not be used.
	//   "IS_NAN" - Test if a field is equal to NaN.
	//   "IS_NULL" - Test if an exprestion evaluates to Null.
	Op string `json:"op,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Field") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Field") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UnaryFilter) MarshalJSON() ([]byte, error) {
	type NoMethod UnaryFilter
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Value: A message that can hold any of the supported value types.
type Value struct {
	// ArrayValue: An array value.
	//
	// Cannot directly contain another array value, though can contain
	// an
	// map which contains another array.
	ArrayValue *ArrayValue `json:"arrayValue,omitempty"`

	// BooleanValue: A boolean value.
	BooleanValue bool `json:"booleanValue,omitempty"`

	// BytesValue: A bytes value.
	//
	// Must not exceed 1 MiB - 89 bytes.
	// Only the first 1,500 bytes are considered by queries.
	BytesValue string `json:"bytesValue,omitempty"`

	// DoubleValue: A double value.
	DoubleValue float64 `json:"doubleValue,omitempty"`

	// GeoPointValue: A geo point value representing a point on the surface
	// of Earth.
	GeoPointValue *LatLng `json:"geoPointValue,omitempty"`

	// IntegerValue: An integer value.
	IntegerValue int64 `json:"integerValue,omitempty,string"`

	// MapValue: A map value.
	MapValue *MapValue `json:"mapValue,omitempty"`

	// NullValue: A null value.
	//
	// Possible values:
	//   "NULL_VALUE" - Null value.
	NullValue string `json:"nullValue,omitempty"`

	// ReferenceValue: A reference to a document. For
	// example:
	// `projects/{project_id}/databases/{database_id}/documents/{doc
	// ument_path}`.
	ReferenceValue string `json:"referenceValue,omitempty"`

	// StringValue: A string value.
	//
	// The string, represented as UTF-8, must not exceed 1 MiB - 89
	// bytes.
	// Only the first 1,500 bytes of the UTF-8 representation are considered
	// by
	// queries.
	StringValue string `json:"stringValue,omitempty"`

	// TimestampValue: A timestamp value.
	//
	// Precise only to microseconds. When stored, any additional precision
	// is
	// rounded down.
	TimestampValue string `json:"timestampValue,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ArrayValue") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ArrayValue") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Value) MarshalJSON() ([]byte, error) {
	type NoMethod Value
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *Value) UnmarshalJSON(data []byte) error {
	type NoMethod Value
	var s1 struct {
		DoubleValue gensupport.JSONFloat64 `json:"doubleValue"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.DoubleValue = float64(s1.DoubleValue)
	return nil
}

// Write: A write on a document.
type Write struct {
	// CurrentDocument: An optional precondition on the document.
	//
	// The write will fail if this is set and not met by the target
	// document.
	CurrentDocument *Precondition `json:"currentDocument,omitempty"`

	// Delete: A document name to delete. In the
	// format:
	// `projects/{project_id}/databases/{database_id}/documents/{docu
	// ment_path}`.
	Delete string `json:"delete,omitempty"`

	// Transform: Applies a tranformation to a document.
	// At most one `transform` per document is allowed in a given
	// request.
	// An `update` cannot follow a `transform` on the same document in a
	// given
	// request.
	Transform *DocumentTransform `json:"transform,omitempty"`

	// Update: A document to write.
	Update *Document `json:"update,omitempty"`

	// UpdateMask: The fields to update in this write.
	//
	// This field can be set only when the operation is `update`.
	// If the mask is not set for an `update` and the document exists,
	// any
	// existing data will be overwritten.
	// If the mask is set and the document on the server has fields not
	// covered by
	// the mask, they are left unchanged.
	// Fields referenced in the mask, but not present in the input document,
	// are
	// deleted from the document on the server.
	// The field paths in this mask must not contain a reserved field name.
	UpdateMask *DocumentMask `json:"updateMask,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CurrentDocument") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CurrentDocument") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Write) MarshalJSON() ([]byte, error) {
	type NoMethod Write
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// WriteRequest: The request for Firestore.Write.
//
// The first request creates a stream, or resumes an existing one from a
// token.
//
// When creating a new stream, the server replies with a response
// containing
// only an ID and a token, to use in the next request.
//
// When resuming a stream, the server first streams any responses later
// than the
// given token, then a response containing only an up-to-date token, to
// use in
// the next request.
type WriteRequest struct {
	// Labels: Labels associated with this write request.
	Labels map[string]string `json:"labels,omitempty"`

	// StreamId: The ID of the write stream to resume.
	// This may only be set in the first message. When left empty, a new
	// write
	// stream will be created.
	StreamId string `json:"streamId,omitempty"`

	// StreamToken: A stream token that was previously sent by the
	// server.
	//
	// The client should set this field to the token from the most
	// recent
	// WriteResponse it has received. This acknowledges that the client
	// has
	// received responses up to this token. After sending this token,
	// earlier
	// tokens may not be used anymore.
	//
	// The server may close the stream if there are too many
	// unacknowledged
	// responses.
	//
	// Leave this field unset when creating a new stream. To resume a stream
	// at
	// a specific point, set this field and the `stream_id` field.
	//
	// Leave this field unset when creating a new stream.
	StreamToken string `json:"streamToken,omitempty"`

	// Writes: The writes to apply.
	//
	// Always executed atomically and in order.
	// This must be empty on the first request.
	// This may be empty on the last request.
	// This must not be empty on all other requests.
	Writes []*Write `json:"writes,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Labels") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Labels") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *WriteRequest) MarshalJSON() ([]byte, error) {
	type NoMethod WriteRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// WriteResponse: The response for Firestore.Write.
type WriteResponse struct {
	// CommitTime: The time at which the commit occurred.
	CommitTime string `json:"commitTime,omitempty"`

	// StreamId: The ID of the stream.
	// Only set on the first message, when a new stream was created.
	StreamId string `json:"streamId,omitempty"`

	// StreamToken: A token that represents the position of this response in
	// the stream.
	// This can be used by a client to resume the stream at this
	// point.
	//
	// This field is always set.
	StreamToken string `json:"streamToken,omitempty"`

	// WriteResults: The result of applying the writes.
	//
	// This i-th write result corresponds to the i-th write in the
	// request.
	WriteResults []*WriteResult `json:"writeResults,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CommitTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CommitTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *WriteResponse) MarshalJSON() ([]byte, error) {
	type NoMethod WriteResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// WriteResult: The result of applying a write.
type WriteResult struct {
	// TransformResults: The results of applying each
	// DocumentTransform.FieldTransform, in the
	// same order.
	TransformResults []*Value `json:"transformResults,omitempty"`

	// UpdateTime: The last update time of the document after applying the
	// write. Not set
	// after a `delete`.
	//
	// If the write did not actually change the document, this will be
	// the
	// previous update_time.
	UpdateTime string `json:"updateTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "TransformResults") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "TransformResults") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *WriteResult) MarshalJSON() ([]byte, error) {
	type NoMethod WriteResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "firestore.projects.databases.exportDocuments":

type ProjectsDatabasesExportDocumentsCall struct {
	s                                                 *Service
	name                                              string
	googlefirestoreadminv1beta1exportdocumentsrequest *GoogleFirestoreAdminV1beta1ExportDocumentsRequest
	urlParams_                                        gensupport.URLParams
	ctx_                                              context.Context
	header_                                           http.Header
}

// ExportDocuments: Exports a copy of all or a subset of documents from
// Google Cloud Firestore
// to another storage system, such as Google Cloud Storage. Recent
// updates to
// documents may not be reflected in the export. The export occurs in
// the
// background and its progress can be monitored and managed via
// the
// Operation resource that is created. The output of an export may only
// be
// used once the associated operation is done. If an export operation
// is
// cancelled before completion it may leave partial data behind in
// Google
// Cloud Storage.
func (r *ProjectsDatabasesService) ExportDocuments(name string, googlefirestoreadminv1beta1exportdocumentsrequest *GoogleFirestoreAdminV1beta1ExportDocumentsRequest) *ProjectsDatabasesExportDocumentsCall {
	c := &ProjectsDatabasesExportDocumentsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.googlefirestoreadminv1beta1exportdocumentsrequest = googlefirestoreadminv1beta1exportdocumentsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDatabasesExportDocumentsCall) Fields(s ...googleapi.Field) *ProjectsDatabasesExportDocumentsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDatabasesExportDocumentsCall) Context(ctx context.Context) *ProjectsDatabasesExportDocumentsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDatabasesExportDocumentsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDatabasesExportDocumentsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googlefirestoreadminv1beta1exportdocumentsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}:exportDocuments")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firestore.projects.databases.exportDocuments" call.
// Exactly one of *GoogleLongrunningOperation or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GoogleLongrunningOperation.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsDatabasesExportDocumentsCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunningOperation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleLongrunningOperation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Exports a copy of all or a subset of documents from Google Cloud Firestore\nto another storage system, such as Google Cloud Storage. Recent updates to\ndocuments may not be reflected in the export. The export occurs in the\nbackground and its progress can be monitored and managed via the\nOperation resource that is created. The output of an export may only be\nused once the associated operation is done. If an export operation is\ncancelled before completion it may leave partial data behind in Google\nCloud Storage.",
	//   "flatPath": "v1beta1/projects/{projectsId}/databases/{databasesId}:exportDocuments",
	//   "httpMethod": "POST",
	//   "id": "firestore.projects.databases.exportDocuments",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Database to export. Should be of the form:\n`projects/{project_id}/databases/{database_id}`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/databases/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}:exportDocuments",
	//   "request": {
	//     "$ref": "GoogleFirestoreAdminV1beta1ExportDocumentsRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleLongrunningOperation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "firestore.projects.databases.importDocuments":

type ProjectsDatabasesImportDocumentsCall struct {
	s                                                 *Service
	name                                              string
	googlefirestoreadminv1beta1importdocumentsrequest *GoogleFirestoreAdminV1beta1ImportDocumentsRequest
	urlParams_                                        gensupport.URLParams
	ctx_                                              context.Context
	header_                                           http.Header
}

// ImportDocuments: Imports documents into Google Cloud Firestore.
// Existing documents with the
// same name are overwritten. The import occurs in the background and
// its
// progress can be monitored and managed via the Operation resource that
// is
// created. If an ImportDocuments operation is cancelled, it is
// possible
// that a subset of the data has already been imported to Cloud
// Firestore.
func (r *ProjectsDatabasesService) ImportDocuments(name string, googlefirestoreadminv1beta1importdocumentsrequest *GoogleFirestoreAdminV1beta1ImportDocumentsRequest) *ProjectsDatabasesImportDocumentsCall {
	c := &ProjectsDatabasesImportDocumentsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.googlefirestoreadminv1beta1importdocumentsrequest = googlefirestoreadminv1beta1importdocumentsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDatabasesImportDocumentsCall) Fields(s ...googleapi.Field) *ProjectsDatabasesImportDocumentsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDatabasesImportDocumentsCall) Context(ctx context.Context) *ProjectsDatabasesImportDocumentsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDatabasesImportDocumentsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDatabasesImportDocumentsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googlefirestoreadminv1beta1importdocumentsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}:importDocuments")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firestore.projects.databases.importDocuments" call.
// Exactly one of *GoogleLongrunningOperation or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GoogleLongrunningOperation.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsDatabasesImportDocumentsCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunningOperation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleLongrunningOperation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Imports documents into Google Cloud Firestore. Existing documents with the\nsame name are overwritten. The import occurs in the background and its\nprogress can be monitored and managed via the Operation resource that is\ncreated. If an ImportDocuments operation is cancelled, it is possible\nthat a subset of the data has already been imported to Cloud Firestore.",
	//   "flatPath": "v1beta1/projects/{projectsId}/databases/{databasesId}:importDocuments",
	//   "httpMethod": "POST",
	//   "id": "firestore.projects.databases.importDocuments",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Database to import into. Should be of the form:\n`projects/{project_id}/databases/{database_id}`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/databases/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}:importDocuments",
	//   "request": {
	//     "$ref": "GoogleFirestoreAdminV1beta1ImportDocumentsRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleLongrunningOperation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "firestore.projects.databases.documents.batchGet":

type ProjectsDatabasesDocumentsBatchGetCall struct {
	s                        *Service
	database                 string
	batchgetdocumentsrequest *BatchGetDocumentsRequest
	urlParams_               gensupport.URLParams
	ctx_                     context.Context
	header_                  http.Header
}

// BatchGet: Gets multiple documents.
//
// Documents returned by this method are not guaranteed to be returned
// in the
// same order that they were requested.
func (r *ProjectsDatabasesDocumentsService) BatchGet(database string, batchgetdocumentsrequest *BatchGetDocumentsRequest) *ProjectsDatabasesDocumentsBatchGetCall {
	c := &ProjectsDatabasesDocumentsBatchGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.database = database
	c.batchgetdocumentsrequest = batchgetdocumentsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDatabasesDocumentsBatchGetCall) Fields(s ...googleapi.Field) *ProjectsDatabasesDocumentsBatchGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDatabasesDocumentsBatchGetCall) Context(ctx context.Context) *ProjectsDatabasesDocumentsBatchGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDatabasesDocumentsBatchGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDatabasesDocumentsBatchGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.batchgetdocumentsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+database}/documents:batchGet")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"database": c.database,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firestore.projects.databases.documents.batchGet" call.
// Exactly one of *BatchGetDocumentsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *BatchGetDocumentsResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsDatabasesDocumentsBatchGetCall) Do(opts ...googleapi.CallOption) (*BatchGetDocumentsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &BatchGetDocumentsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets multiple documents.\n\nDocuments returned by this method are not guaranteed to be returned in the\nsame order that they were requested.",
	//   "flatPath": "v1beta1/projects/{projectsId}/databases/{databasesId}/documents:batchGet",
	//   "httpMethod": "POST",
	//   "id": "firestore.projects.databases.documents.batchGet",
	//   "parameterOrder": [
	//     "database"
	//   ],
	//   "parameters": {
	//     "database": {
	//       "description": "The database name. In the format:\n`projects/{project_id}/databases/{database_id}`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/databases/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+database}/documents:batchGet",
	//   "request": {
	//     "$ref": "BatchGetDocumentsRequest"
	//   },
	//   "response": {
	//     "$ref": "BatchGetDocumentsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "firestore.projects.databases.documents.beginTransaction":

type ProjectsDatabasesDocumentsBeginTransactionCall struct {
	s                       *Service
	database                string
	begintransactionrequest *BeginTransactionRequest
	urlParams_              gensupport.URLParams
	ctx_                    context.Context
	header_                 http.Header
}

// BeginTransaction: Starts a new transaction.
func (r *ProjectsDatabasesDocumentsService) BeginTransaction(database string, begintransactionrequest *BeginTransactionRequest) *ProjectsDatabasesDocumentsBeginTransactionCall {
	c := &ProjectsDatabasesDocumentsBeginTransactionCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.database = database
	c.begintransactionrequest = begintransactionrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDatabasesDocumentsBeginTransactionCall) Fields(s ...googleapi.Field) *ProjectsDatabasesDocumentsBeginTransactionCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDatabasesDocumentsBeginTransactionCall) Context(ctx context.Context) *ProjectsDatabasesDocumentsBeginTransactionCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDatabasesDocumentsBeginTransactionCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDatabasesDocumentsBeginTransactionCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.begintransactionrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+database}/documents:beginTransaction")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"database": c.database,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firestore.projects.databases.documents.beginTransaction" call.
// Exactly one of *BeginTransactionResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *BeginTransactionResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsDatabasesDocumentsBeginTransactionCall) Do(opts ...googleapi.CallOption) (*BeginTransactionResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &BeginTransactionResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Starts a new transaction.",
	//   "flatPath": "v1beta1/projects/{projectsId}/databases/{databasesId}/documents:beginTransaction",
	//   "httpMethod": "POST",
	//   "id": "firestore.projects.databases.documents.beginTransaction",
	//   "parameterOrder": [
	//     "database"
	//   ],
	//   "parameters": {
	//     "database": {
	//       "description": "The database name. In the format:\n`projects/{project_id}/databases/{database_id}`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/databases/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+database}/documents:beginTransaction",
	//   "request": {
	//     "$ref": "BeginTransactionRequest"
	//   },
	//   "response": {
	//     "$ref": "BeginTransactionResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "firestore.projects.databases.documents.commit":

type ProjectsDatabasesDocumentsCommitCall struct {
	s             *Service
	database      string
	commitrequest *CommitRequest
	urlParams_    gensupport.URLParams
	ctx_          context.Context
	header_       http.Header
}

// Commit: Commits a transaction, while optionally updating documents.
func (r *ProjectsDatabasesDocumentsService) Commit(database string, commitrequest *CommitRequest) *ProjectsDatabasesDocumentsCommitCall {
	c := &ProjectsDatabasesDocumentsCommitCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.database = database
	c.commitrequest = commitrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDatabasesDocumentsCommitCall) Fields(s ...googleapi.Field) *ProjectsDatabasesDocumentsCommitCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDatabasesDocumentsCommitCall) Context(ctx context.Context) *ProjectsDatabasesDocumentsCommitCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDatabasesDocumentsCommitCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDatabasesDocumentsCommitCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.commitrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+database}/documents:commit")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"database": c.database,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firestore.projects.databases.documents.commit" call.
// Exactly one of *CommitResponse or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *CommitResponse.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsDatabasesDocumentsCommitCall) Do(opts ...googleapi.CallOption) (*CommitResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &CommitResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Commits a transaction, while optionally updating documents.",
	//   "flatPath": "v1beta1/projects/{projectsId}/databases/{databasesId}/documents:commit",
	//   "httpMethod": "POST",
	//   "id": "firestore.projects.databases.documents.commit",
	//   "parameterOrder": [
	//     "database"
	//   ],
	//   "parameters": {
	//     "database": {
	//       "description": "The database name. In the format:\n`projects/{project_id}/databases/{database_id}`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/databases/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+database}/documents:commit",
	//   "request": {
	//     "$ref": "CommitRequest"
	//   },
	//   "response": {
	//     "$ref": "CommitResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "firestore.projects.databases.documents.createDocument":

type ProjectsDatabasesDocumentsCreateDocumentCall struct {
	s            *Service
	parent       string
	collectionId string
	document     *Document
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// CreateDocument: Creates a new document.
func (r *ProjectsDatabasesDocumentsService) CreateDocument(parent string, collectionId string, document *Document) *ProjectsDatabasesDocumentsCreateDocumentCall {
	c := &ProjectsDatabasesDocumentsCreateDocumentCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.collectionId = collectionId
	c.document = document
	return c
}

// DocumentId sets the optional parameter "documentId": The
// client-assigned document ID to use for this document.
//
//  If not specified, an ID will be assigned by the service.
func (c *ProjectsDatabasesDocumentsCreateDocumentCall) DocumentId(documentId string) *ProjectsDatabasesDocumentsCreateDocumentCall {
	c.urlParams_.Set("documentId", documentId)
	return c
}

// MaskFieldPaths sets the optional parameter "mask.fieldPaths": The
// list of field paths in the mask. See Document.fields for a field
// path syntax reference.
func (c *ProjectsDatabasesDocumentsCreateDocumentCall) MaskFieldPaths(maskFieldPaths ...string) *ProjectsDatabasesDocumentsCreateDocumentCall {
	c.urlParams_.SetMulti("mask.fieldPaths", append([]string{}, maskFieldPaths...))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDatabasesDocumentsCreateDocumentCall) Fields(s ...googleapi.Field) *ProjectsDatabasesDocumentsCreateDocumentCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDatabasesDocumentsCreateDocumentCall) Context(ctx context.Context) *ProjectsDatabasesDocumentsCreateDocumentCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDatabasesDocumentsCreateDocumentCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDatabasesDocumentsCreateDocumentCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.document)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/{collectionId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent":       c.parent,
		"collectionId": c.collectionId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firestore.projects.databases.documents.createDocument" call.
// Exactly one of *Document or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Document.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsDatabasesDocumentsCreateDocumentCall) Do(opts ...googleapi.CallOption) (*Document, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Document{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new document.",
	//   "flatPath": "v1beta1/projects/{projectsId}/databases/{databasesId}/documents/{documentsId}/{collectionId}",
	//   "httpMethod": "POST",
	//   "id": "firestore.projects.databases.documents.createDocument",
	//   "parameterOrder": [
	//     "parent",
	//     "collectionId"
	//   ],
	//   "parameters": {
	//     "collectionId": {
	//       "description": "The collection ID, relative to `parent`, to list. For example: `chatrooms`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "documentId": {
	//       "description": "The client-assigned document ID to use for this document.\n\nOptional. If not specified, an ID will be assigned by the service.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "mask.fieldPaths": {
	//       "description": "The list of field paths in the mask. See Document.fields for a field\npath syntax reference.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "The parent resource. For example:\n`projects/{project_id}/databases/{database_id}/documents` or\n`projects/{project_id}/databases/{database_id}/documents/chatrooms/{chatroom_id}`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/databases/[^/]+/documents/.+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/{collectionId}",
	//   "request": {
	//     "$ref": "Document"
	//   },
	//   "response": {
	//     "$ref": "Document"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "firestore.projects.databases.documents.delete":

type ProjectsDatabasesDocumentsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes a document.
func (r *ProjectsDatabasesDocumentsService) Delete(name string) *ProjectsDatabasesDocumentsDeleteCall {
	c := &ProjectsDatabasesDocumentsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// CurrentDocumentExists sets the optional parameter
// "currentDocument.exists": When set to `true`, the target document
// must exist.
// When set to `false`, the target document must not exist.
func (c *ProjectsDatabasesDocumentsDeleteCall) CurrentDocumentExists(currentDocumentExists bool) *ProjectsDatabasesDocumentsDeleteCall {
	c.urlParams_.Set("currentDocument.exists", fmt.Sprint(currentDocumentExists))
	return c
}

// CurrentDocumentUpdateTime sets the optional parameter
// "currentDocument.updateTime": When set, the target document must
// exist and have been last updated at
// that time.
func (c *ProjectsDatabasesDocumentsDeleteCall) CurrentDocumentUpdateTime(currentDocumentUpdateTime string) *ProjectsDatabasesDocumentsDeleteCall {
	c.urlParams_.Set("currentDocument.updateTime", currentDocumentUpdateTime)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDatabasesDocumentsDeleteCall) Fields(s ...googleapi.Field) *ProjectsDatabasesDocumentsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDatabasesDocumentsDeleteCall) Context(ctx context.Context) *ProjectsDatabasesDocumentsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDatabasesDocumentsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDatabasesDocumentsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firestore.projects.databases.documents.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsDatabasesDocumentsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a document.",
	//   "flatPath": "v1beta1/projects/{projectsId}/databases/{databasesId}/documents/{documentsId}/{documentsId1}",
	//   "httpMethod": "DELETE",
	//   "id": "firestore.projects.databases.documents.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "currentDocument.exists": {
	//       "description": "When set to `true`, the target document must exist.\nWhen set to `false`, the target document must not exist.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "currentDocument.updateTime": {
	//       "description": "When set, the target document must exist and have been last updated at\nthat time.",
	//       "format": "google-datetime",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The resource name of the Document to delete. In the format:\n`projects/{project_id}/databases/{database_id}/documents/{document_path}`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/databases/[^/]+/documents/[^/]+/.+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "firestore.projects.databases.documents.get":

type ProjectsDatabasesDocumentsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a single document.
func (r *ProjectsDatabasesDocumentsService) Get(name string) *ProjectsDatabasesDocumentsGetCall {
	c := &ProjectsDatabasesDocumentsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// MaskFieldPaths sets the optional parameter "mask.fieldPaths": The
// list of field paths in the mask. See Document.fields for a field
// path syntax reference.
func (c *ProjectsDatabasesDocumentsGetCall) MaskFieldPaths(maskFieldPaths ...string) *ProjectsDatabasesDocumentsGetCall {
	c.urlParams_.SetMulti("mask.fieldPaths", append([]string{}, maskFieldPaths...))
	return c
}

// ReadTime sets the optional parameter "readTime": Reads the version of
// the document at the given time.
// This may not be older than 60 seconds.
func (c *ProjectsDatabasesDocumentsGetCall) ReadTime(readTime string) *ProjectsDatabasesDocumentsGetCall {
	c.urlParams_.Set("readTime", readTime)
	return c
}

// Transaction sets the optional parameter "transaction": Reads the
// document in a transaction.
func (c *ProjectsDatabasesDocumentsGetCall) Transaction(transaction string) *ProjectsDatabasesDocumentsGetCall {
	c.urlParams_.Set("transaction", transaction)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDatabasesDocumentsGetCall) Fields(s ...googleapi.Field) *ProjectsDatabasesDocumentsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsDatabasesDocumentsGetCall) IfNoneMatch(entityTag string) *ProjectsDatabasesDocumentsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDatabasesDocumentsGetCall) Context(ctx context.Context) *ProjectsDatabasesDocumentsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDatabasesDocumentsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDatabasesDocumentsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firestore.projects.databases.documents.get" call.
// Exactly one of *Document or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Document.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsDatabasesDocumentsGetCall) Do(opts ...googleapi.CallOption) (*Document, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Document{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a single document.",
	//   "flatPath": "v1beta1/projects/{projectsId}/databases/{databasesId}/documents/{documentsId}/{documentsId1}",
	//   "httpMethod": "GET",
	//   "id": "firestore.projects.databases.documents.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "mask.fieldPaths": {
	//       "description": "The list of field paths in the mask. See Document.fields for a field\npath syntax reference.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The resource name of the Document to get. In the format:\n`projects/{project_id}/databases/{database_id}/documents/{document_path}`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/databases/[^/]+/documents/[^/]+/.+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "readTime": {
	//       "description": "Reads the version of the document at the given time.\nThis may not be older than 60 seconds.",
	//       "format": "google-datetime",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "transaction": {
	//       "description": "Reads the document in a transaction.",
	//       "format": "byte",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "Document"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "firestore.projects.databases.documents.list":

type ProjectsDatabasesDocumentsListCall struct {
	s            *Service
	parent       string
	collectionId string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists documents.
func (r *ProjectsDatabasesDocumentsService) List(parent string, collectionId string) *ProjectsDatabasesDocumentsListCall {
	c := &ProjectsDatabasesDocumentsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.collectionId = collectionId
	return c
}

// MaskFieldPaths sets the optional parameter "mask.fieldPaths": The
// list of field paths in the mask. See Document.fields for a field
// path syntax reference.
func (c *ProjectsDatabasesDocumentsListCall) MaskFieldPaths(maskFieldPaths ...string) *ProjectsDatabasesDocumentsListCall {
	c.urlParams_.SetMulti("mask.fieldPaths", append([]string{}, maskFieldPaths...))
	return c
}

// OrderBy sets the optional parameter "orderBy": The order to sort
// results by. For example: `priority desc, name`.
func (c *ProjectsDatabasesDocumentsListCall) OrderBy(orderBy string) *ProjectsDatabasesDocumentsListCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of documents to return.
func (c *ProjectsDatabasesDocumentsListCall) PageSize(pageSize int64) *ProjectsDatabasesDocumentsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The
// `next_page_token` value returned from a previous List request, if
// any.
func (c *ProjectsDatabasesDocumentsListCall) PageToken(pageToken string) *ProjectsDatabasesDocumentsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// ReadTime sets the optional parameter "readTime": Reads documents as
// they were at the given time.
// This may not be older than 60 seconds.
func (c *ProjectsDatabasesDocumentsListCall) ReadTime(readTime string) *ProjectsDatabasesDocumentsListCall {
	c.urlParams_.Set("readTime", readTime)
	return c
}

// ShowMissing sets the optional parameter "showMissing": If the list
// should show missing documents. A missing document is a
// document that does not exist but has sub-documents. These documents
// will
// be returned with a key but will not have fields,
// Document.create_time,
// or Document.update_time set.
//
// Requests with `show_missing` may not specify `where` or
// `order_by`.
func (c *ProjectsDatabasesDocumentsListCall) ShowMissing(showMissing bool) *ProjectsDatabasesDocumentsListCall {
	c.urlParams_.Set("showMissing", fmt.Sprint(showMissing))
	return c
}

// Transaction sets the optional parameter "transaction": Reads
// documents in a transaction.
func (c *ProjectsDatabasesDocumentsListCall) Transaction(transaction string) *ProjectsDatabasesDocumentsListCall {
	c.urlParams_.Set("transaction", transaction)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDatabasesDocumentsListCall) Fields(s ...googleapi.Field) *ProjectsDatabasesDocumentsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsDatabasesDocumentsListCall) IfNoneMatch(entityTag string) *ProjectsDatabasesDocumentsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDatabasesDocumentsListCall) Context(ctx context.Context) *ProjectsDatabasesDocumentsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDatabasesDocumentsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDatabasesDocumentsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/{collectionId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent":       c.parent,
		"collectionId": c.collectionId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firestore.projects.databases.documents.list" call.
// Exactly one of *ListDocumentsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListDocumentsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsDatabasesDocumentsListCall) Do(opts ...googleapi.CallOption) (*ListDocumentsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListDocumentsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists documents.",
	//   "flatPath": "v1beta1/projects/{projectsId}/databases/{databasesId}/documents/{documentsId}/{documentsId1}/{collectionId}",
	//   "httpMethod": "GET",
	//   "id": "firestore.projects.databases.documents.list",
	//   "parameterOrder": [
	//     "parent",
	//     "collectionId"
	//   ],
	//   "parameters": {
	//     "collectionId": {
	//       "description": "The collection ID, relative to `parent`, to list. For example: `chatrooms`\nor `messages`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "mask.fieldPaths": {
	//       "description": "The list of field paths in the mask. See Document.fields for a field\npath syntax reference.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "orderBy": {
	//       "description": "The order to sort results by. For example: `priority desc, name`.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of documents to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The `next_page_token` value returned from a previous List request, if any.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "The parent resource name. In the format:\n`projects/{project_id}/databases/{database_id}/documents` or\n`projects/{project_id}/databases/{database_id}/documents/{document_path}`.\nFor example:\n`projects/my-project/databases/my-database/documents` or\n`projects/my-project/databases/my-database/documents/chatrooms/my-chatroom`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/databases/[^/]+/documents/[^/]+/.+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "readTime": {
	//       "description": "Reads documents as they were at the given time.\nThis may not be older than 60 seconds.",
	//       "format": "google-datetime",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "showMissing": {
	//       "description": "If the list should show missing documents. A missing document is a\ndocument that does not exist but has sub-documents. These documents will\nbe returned with a key but will not have fields, Document.create_time,\nor Document.update_time set.\n\nRequests with `show_missing` may not specify `where` or\n`order_by`.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "transaction": {
	//       "description": "Reads documents in a transaction.",
	//       "format": "byte",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/{collectionId}",
	//   "response": {
	//     "$ref": "ListDocumentsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsDatabasesDocumentsListCall) Pages(ctx context.Context, f func(*ListDocumentsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "firestore.projects.databases.documents.listCollectionIds":

type ProjectsDatabasesDocumentsListCollectionIdsCall struct {
	s                        *Service
	parent                   string
	listcollectionidsrequest *ListCollectionIdsRequest
	urlParams_               gensupport.URLParams
	ctx_                     context.Context
	header_                  http.Header
}

// ListCollectionIds: Lists all the collection IDs underneath a
// document.
func (r *ProjectsDatabasesDocumentsService) ListCollectionIds(parent string, listcollectionidsrequest *ListCollectionIdsRequest) *ProjectsDatabasesDocumentsListCollectionIdsCall {
	c := &ProjectsDatabasesDocumentsListCollectionIdsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.listcollectionidsrequest = listcollectionidsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDatabasesDocumentsListCollectionIdsCall) Fields(s ...googleapi.Field) *ProjectsDatabasesDocumentsListCollectionIdsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDatabasesDocumentsListCollectionIdsCall) Context(ctx context.Context) *ProjectsDatabasesDocumentsListCollectionIdsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDatabasesDocumentsListCollectionIdsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDatabasesDocumentsListCollectionIdsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.listcollectionidsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}:listCollectionIds")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firestore.projects.databases.documents.listCollectionIds" call.
// Exactly one of *ListCollectionIdsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListCollectionIdsResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsDatabasesDocumentsListCollectionIdsCall) Do(opts ...googleapi.CallOption) (*ListCollectionIdsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListCollectionIdsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all the collection IDs underneath a document.",
	//   "flatPath": "v1beta1/projects/{projectsId}/databases/{databasesId}/documents/{documentsId}/{documentsId1}:listCollectionIds",
	//   "httpMethod": "POST",
	//   "id": "firestore.projects.databases.documents.listCollectionIds",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "The parent document. In the format:\n`projects/{project_id}/databases/{database_id}/documents/{document_path}`.\nFor example:\n`projects/my-project/databases/my-database/documents/chatrooms/my-chatroom`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/databases/[^/]+/documents/[^/]+/.+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}:listCollectionIds",
	//   "request": {
	//     "$ref": "ListCollectionIdsRequest"
	//   },
	//   "response": {
	//     "$ref": "ListCollectionIdsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsDatabasesDocumentsListCollectionIdsCall) Pages(ctx context.Context, f func(*ListCollectionIdsResponse) error) error {
	c.ctx_ = ctx
	defer func(pt string) { c.listcollectionidsrequest.PageToken = pt }(c.listcollectionidsrequest.PageToken) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.listcollectionidsrequest.PageToken = x.NextPageToken
	}
}

// method id "firestore.projects.databases.documents.listen":

type ProjectsDatabasesDocumentsListenCall struct {
	s             *Service
	database      string
	listenrequest *ListenRequest
	urlParams_    gensupport.URLParams
	ctx_          context.Context
	header_       http.Header
}

// Listen: Listens to changes.
func (r *ProjectsDatabasesDocumentsService) Listen(database string, listenrequest *ListenRequest) *ProjectsDatabasesDocumentsListenCall {
	c := &ProjectsDatabasesDocumentsListenCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.database = database
	c.listenrequest = listenrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDatabasesDocumentsListenCall) Fields(s ...googleapi.Field) *ProjectsDatabasesDocumentsListenCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDatabasesDocumentsListenCall) Context(ctx context.Context) *ProjectsDatabasesDocumentsListenCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDatabasesDocumentsListenCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDatabasesDocumentsListenCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.listenrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+database}/documents:listen")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"database": c.database,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firestore.projects.databases.documents.listen" call.
// Exactly one of *ListenResponse or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *ListenResponse.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsDatabasesDocumentsListenCall) Do(opts ...googleapi.CallOption) (*ListenResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListenResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Listens to changes.",
	//   "flatPath": "v1beta1/projects/{projectsId}/databases/{databasesId}/documents:listen",
	//   "httpMethod": "POST",
	//   "id": "firestore.projects.databases.documents.listen",
	//   "parameterOrder": [
	//     "database"
	//   ],
	//   "parameters": {
	//     "database": {
	//       "description": "The database name. In the format:\n`projects/{project_id}/databases/{database_id}`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/databases/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+database}/documents:listen",
	//   "request": {
	//     "$ref": "ListenRequest"
	//   },
	//   "response": {
	//     "$ref": "ListenResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "firestore.projects.databases.documents.patch":

type ProjectsDatabasesDocumentsPatchCall struct {
	s          *Service
	name       string
	document   *Document
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Patch: Updates or inserts a document.
func (r *ProjectsDatabasesDocumentsService) Patch(name string, document *Document) *ProjectsDatabasesDocumentsPatchCall {
	c := &ProjectsDatabasesDocumentsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.document = document
	return c
}

// CurrentDocumentExists sets the optional parameter
// "currentDocument.exists": When set to `true`, the target document
// must exist.
// When set to `false`, the target document must not exist.
func (c *ProjectsDatabasesDocumentsPatchCall) CurrentDocumentExists(currentDocumentExists bool) *ProjectsDatabasesDocumentsPatchCall {
	c.urlParams_.Set("currentDocument.exists", fmt.Sprint(currentDocumentExists))
	return c
}

// CurrentDocumentUpdateTime sets the optional parameter
// "currentDocument.updateTime": When set, the target document must
// exist and have been last updated at
// that time.
func (c *ProjectsDatabasesDocumentsPatchCall) CurrentDocumentUpdateTime(currentDocumentUpdateTime string) *ProjectsDatabasesDocumentsPatchCall {
	c.urlParams_.Set("currentDocument.updateTime", currentDocumentUpdateTime)
	return c
}

// MaskFieldPaths sets the optional parameter "mask.fieldPaths": The
// list of field paths in the mask. See Document.fields for a field
// path syntax reference.
func (c *ProjectsDatabasesDocumentsPatchCall) MaskFieldPaths(maskFieldPaths ...string) *ProjectsDatabasesDocumentsPatchCall {
	c.urlParams_.SetMulti("mask.fieldPaths", append([]string{}, maskFieldPaths...))
	return c
}

// UpdateMaskFieldPaths sets the optional parameter
// "updateMask.fieldPaths": The list of field paths in the mask. See
// Document.fields for a field
// path syntax reference.
func (c *ProjectsDatabasesDocumentsPatchCall) UpdateMaskFieldPaths(updateMaskFieldPaths ...string) *ProjectsDatabasesDocumentsPatchCall {
	c.urlParams_.SetMulti("updateMask.fieldPaths", append([]string{}, updateMaskFieldPaths...))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDatabasesDocumentsPatchCall) Fields(s ...googleapi.Field) *ProjectsDatabasesDocumentsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDatabasesDocumentsPatchCall) Context(ctx context.Context) *ProjectsDatabasesDocumentsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDatabasesDocumentsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDatabasesDocumentsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.document)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firestore.projects.databases.documents.patch" call.
// Exactly one of *Document or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Document.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsDatabasesDocumentsPatchCall) Do(opts ...googleapi.CallOption) (*Document, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Document{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates or inserts a document.",
	//   "flatPath": "v1beta1/projects/{projectsId}/databases/{databasesId}/documents/{documentsId}/{documentsId1}",
	//   "httpMethod": "PATCH",
	//   "id": "firestore.projects.databases.documents.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "currentDocument.exists": {
	//       "description": "When set to `true`, the target document must exist.\nWhen set to `false`, the target document must not exist.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "currentDocument.updateTime": {
	//       "description": "When set, the target document must exist and have been last updated at\nthat time.",
	//       "format": "google-datetime",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "mask.fieldPaths": {
	//       "description": "The list of field paths in the mask. See Document.fields for a field\npath syntax reference.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The resource name of the document, for example\n`projects/{project_id}/databases/{database_id}/documents/{document_path}`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/databases/[^/]+/documents/[^/]+/.+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask.fieldPaths": {
	//       "description": "The list of field paths in the mask. See Document.fields for a field\npath syntax reference.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "request": {
	//     "$ref": "Document"
	//   },
	//   "response": {
	//     "$ref": "Document"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "firestore.projects.databases.documents.rollback":

type ProjectsDatabasesDocumentsRollbackCall struct {
	s               *Service
	database        string
	rollbackrequest *RollbackRequest
	urlParams_      gensupport.URLParams
	ctx_            context.Context
	header_         http.Header
}

// Rollback: Rolls back a transaction.
func (r *ProjectsDatabasesDocumentsService) Rollback(database string, rollbackrequest *RollbackRequest) *ProjectsDatabasesDocumentsRollbackCall {
	c := &ProjectsDatabasesDocumentsRollbackCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.database = database
	c.rollbackrequest = rollbackrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDatabasesDocumentsRollbackCall) Fields(s ...googleapi.Field) *ProjectsDatabasesDocumentsRollbackCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDatabasesDocumentsRollbackCall) Context(ctx context.Context) *ProjectsDatabasesDocumentsRollbackCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDatabasesDocumentsRollbackCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDatabasesDocumentsRollbackCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.rollbackrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+database}/documents:rollback")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"database": c.database,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firestore.projects.databases.documents.rollback" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsDatabasesDocumentsRollbackCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Rolls back a transaction.",
	//   "flatPath": "v1beta1/projects/{projectsId}/databases/{databasesId}/documents:rollback",
	//   "httpMethod": "POST",
	//   "id": "firestore.projects.databases.documents.rollback",
	//   "parameterOrder": [
	//     "database"
	//   ],
	//   "parameters": {
	//     "database": {
	//       "description": "The database name. In the format:\n`projects/{project_id}/databases/{database_id}`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/databases/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+database}/documents:rollback",
	//   "request": {
	//     "$ref": "RollbackRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "firestore.projects.databases.documents.runQuery":

type ProjectsDatabasesDocumentsRunQueryCall struct {
	s               *Service
	parent          string
	runqueryrequest *RunQueryRequest
	urlParams_      gensupport.URLParams
	ctx_            context.Context
	header_         http.Header
}

// RunQuery: Runs a query.
func (r *ProjectsDatabasesDocumentsService) RunQuery(parent string, runqueryrequest *RunQueryRequest) *ProjectsDatabasesDocumentsRunQueryCall {
	c := &ProjectsDatabasesDocumentsRunQueryCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.runqueryrequest = runqueryrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDatabasesDocumentsRunQueryCall) Fields(s ...googleapi.Field) *ProjectsDatabasesDocumentsRunQueryCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDatabasesDocumentsRunQueryCall) Context(ctx context.Context) *ProjectsDatabasesDocumentsRunQueryCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDatabasesDocumentsRunQueryCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDatabasesDocumentsRunQueryCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.runqueryrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}:runQuery")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firestore.projects.databases.documents.runQuery" call.
// Exactly one of *RunQueryResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *RunQueryResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsDatabasesDocumentsRunQueryCall) Do(opts ...googleapi.CallOption) (*RunQueryResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &RunQueryResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Runs a query.",
	//   "flatPath": "v1beta1/projects/{projectsId}/databases/{databasesId}/documents/{documentsId}/{documentsId1}:runQuery",
	//   "httpMethod": "POST",
	//   "id": "firestore.projects.databases.documents.runQuery",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "The parent resource name. In the format:\n`projects/{project_id}/databases/{database_id}/documents` or\n`projects/{project_id}/databases/{database_id}/documents/{document_path}`.\nFor example:\n`projects/my-project/databases/my-database/documents` or\n`projects/my-project/databases/my-database/documents/chatrooms/my-chatroom`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/databases/[^/]+/documents/[^/]+/.+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}:runQuery",
	//   "request": {
	//     "$ref": "RunQueryRequest"
	//   },
	//   "response": {
	//     "$ref": "RunQueryResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "firestore.projects.databases.documents.write":

type ProjectsDatabasesDocumentsWriteCall struct {
	s            *Service
	database     string
	writerequest *WriteRequest
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// Write: Streams batches of document updates and deletes, in order.
func (r *ProjectsDatabasesDocumentsService) Write(database string, writerequest *WriteRequest) *ProjectsDatabasesDocumentsWriteCall {
	c := &ProjectsDatabasesDocumentsWriteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.database = database
	c.writerequest = writerequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDatabasesDocumentsWriteCall) Fields(s ...googleapi.Field) *ProjectsDatabasesDocumentsWriteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDatabasesDocumentsWriteCall) Context(ctx context.Context) *ProjectsDatabasesDocumentsWriteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDatabasesDocumentsWriteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDatabasesDocumentsWriteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.writerequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+database}/documents:write")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"database": c.database,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firestore.projects.databases.documents.write" call.
// Exactly one of *WriteResponse or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *WriteResponse.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsDatabasesDocumentsWriteCall) Do(opts ...googleapi.CallOption) (*WriteResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &WriteResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Streams batches of document updates and deletes, in order.",
	//   "flatPath": "v1beta1/projects/{projectsId}/databases/{databasesId}/documents:write",
	//   "httpMethod": "POST",
	//   "id": "firestore.projects.databases.documents.write",
	//   "parameterOrder": [
	//     "database"
	//   ],
	//   "parameters": {
	//     "database": {
	//       "description": "The database name. In the format:\n`projects/{project_id}/databases/{database_id}`.\nThis is only required in the first message.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/databases/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+database}/documents:write",
	//   "request": {
	//     "$ref": "WriteRequest"
	//   },
	//   "response": {
	//     "$ref": "WriteResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "firestore.projects.databases.indexes.create":

type ProjectsDatabasesIndexesCreateCall struct {
	s                                *Service
	parent                           string
	googlefirestoreadminv1beta1index *GoogleFirestoreAdminV1beta1Index
	urlParams_                       gensupport.URLParams
	ctx_                             context.Context
	header_                          http.Header
}

// Create: Creates the specified index.
// A newly created index's initial state is `CREATING`. On completion of
// the
// returned google.longrunning.Operation, the state will be `READY`.
// If the index already exists, the call will return an
// `ALREADY_EXISTS`
// status.
//
// During creation, the process could result in an error, in which case
// the
// index will move to the `ERROR` state. The process can be recovered
// by
// fixing the data that caused the error, removing the index
// with
// delete, then re-creating the index with
// create.
//
// Indexes with a single field cannot be created.
func (r *ProjectsDatabasesIndexesService) Create(parent string, googlefirestoreadminv1beta1index *GoogleFirestoreAdminV1beta1Index) *ProjectsDatabasesIndexesCreateCall {
	c := &ProjectsDatabasesIndexesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googlefirestoreadminv1beta1index = googlefirestoreadminv1beta1index
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDatabasesIndexesCreateCall) Fields(s ...googleapi.Field) *ProjectsDatabasesIndexesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDatabasesIndexesCreateCall) Context(ctx context.Context) *ProjectsDatabasesIndexesCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDatabasesIndexesCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDatabasesIndexesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googlefirestoreadminv1beta1index)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/indexes")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firestore.projects.databases.indexes.create" call.
// Exactly one of *GoogleLongrunningOperation or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GoogleLongrunningOperation.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsDatabasesIndexesCreateCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunningOperation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleLongrunningOperation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates the specified index.\nA newly created index's initial state is `CREATING`. On completion of the\nreturned google.longrunning.Operation, the state will be `READY`.\nIf the index already exists, the call will return an `ALREADY_EXISTS`\nstatus.\n\nDuring creation, the process could result in an error, in which case the\nindex will move to the `ERROR` state. The process can be recovered by\nfixing the data that caused the error, removing the index with\ndelete, then re-creating the index with\ncreate.\n\nIndexes with a single field cannot be created.",
	//   "flatPath": "v1beta1/projects/{projectsId}/databases/{databasesId}/indexes",
	//   "httpMethod": "POST",
	//   "id": "firestore.projects.databases.indexes.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "The name of the database this index will apply to. For example:\n`projects/{project_id}/databases/{database_id}`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/databases/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/indexes",
	//   "request": {
	//     "$ref": "GoogleFirestoreAdminV1beta1Index"
	//   },
	//   "response": {
	//     "$ref": "GoogleLongrunningOperation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "firestore.projects.databases.indexes.delete":

type ProjectsDatabasesIndexesDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes an index.
func (r *ProjectsDatabasesIndexesService) Delete(name string) *ProjectsDatabasesIndexesDeleteCall {
	c := &ProjectsDatabasesIndexesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDatabasesIndexesDeleteCall) Fields(s ...googleapi.Field) *ProjectsDatabasesIndexesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDatabasesIndexesDeleteCall) Context(ctx context.Context) *ProjectsDatabasesIndexesDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDatabasesIndexesDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDatabasesIndexesDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firestore.projects.databases.indexes.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsDatabasesIndexesDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes an index.",
	//   "flatPath": "v1beta1/projects/{projectsId}/databases/{databasesId}/indexes/{indexesId}",
	//   "httpMethod": "DELETE",
	//   "id": "firestore.projects.databases.indexes.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The index name. For example:\n`projects/{project_id}/databases/{database_id}/indexes/{index_id}`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/databases/[^/]+/indexes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "firestore.projects.databases.indexes.get":

type ProjectsDatabasesIndexesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets an index.
func (r *ProjectsDatabasesIndexesService) Get(name string) *ProjectsDatabasesIndexesGetCall {
	c := &ProjectsDatabasesIndexesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDatabasesIndexesGetCall) Fields(s ...googleapi.Field) *ProjectsDatabasesIndexesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsDatabasesIndexesGetCall) IfNoneMatch(entityTag string) *ProjectsDatabasesIndexesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDatabasesIndexesGetCall) Context(ctx context.Context) *ProjectsDatabasesIndexesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDatabasesIndexesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDatabasesIndexesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firestore.projects.databases.indexes.get" call.
// Exactly one of *GoogleFirestoreAdminV1beta1Index or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleFirestoreAdminV1beta1Index.ServerResponse.Header or (if
// a response was returned at all) in error.(*googleapi.Error).Header.
// Use googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsDatabasesIndexesGetCall) Do(opts ...googleapi.CallOption) (*GoogleFirestoreAdminV1beta1Index, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleFirestoreAdminV1beta1Index{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets an index.",
	//   "flatPath": "v1beta1/projects/{projectsId}/databases/{databasesId}/indexes/{indexesId}",
	//   "httpMethod": "GET",
	//   "id": "firestore.projects.databases.indexes.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the index. For example:\n`projects/{project_id}/databases/{database_id}/indexes/{index_id}`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/databases/[^/]+/indexes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "GoogleFirestoreAdminV1beta1Index"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "firestore.projects.databases.indexes.list":

type ProjectsDatabasesIndexesListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the indexes that match the specified filters.
func (r *ProjectsDatabasesIndexesService) List(parent string) *ProjectsDatabasesIndexesListCall {
	c := &ProjectsDatabasesIndexesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// Filter sets the optional parameter "filter":
func (c *ProjectsDatabasesIndexesListCall) Filter(filter string) *ProjectsDatabasesIndexesListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": The standard List
// page size.
func (c *ProjectsDatabasesIndexesListCall) PageSize(pageSize int64) *ProjectsDatabasesIndexesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The standard List
// page token.
func (c *ProjectsDatabasesIndexesListCall) PageToken(pageToken string) *ProjectsDatabasesIndexesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDatabasesIndexesListCall) Fields(s ...googleapi.Field) *ProjectsDatabasesIndexesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsDatabasesIndexesListCall) IfNoneMatch(entityTag string) *ProjectsDatabasesIndexesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDatabasesIndexesListCall) Context(ctx context.Context) *ProjectsDatabasesIndexesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsDatabasesIndexesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsDatabasesIndexesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+parent}/indexes")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "firestore.projects.databases.indexes.list" call.
// Exactly one of *GoogleFirestoreAdminV1beta1ListIndexesResponse or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleFirestoreAdminV1beta1ListIndexesResponse.ServerResponse.Header
// or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsDatabasesIndexesListCall) Do(opts ...googleapi.CallOption) (*GoogleFirestoreAdminV1beta1ListIndexesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleFirestoreAdminV1beta1ListIndexesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the indexes that match the specified filters.",
	//   "flatPath": "v1beta1/projects/{projectsId}/databases/{databasesId}/indexes",
	//   "httpMethod": "GET",
	//   "id": "firestore.projects.databases.indexes.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The standard List page size.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The standard List page token.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "The database name. For example:\n`projects/{project_id}/databases/{database_id}`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/databases/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+parent}/indexes",
	//   "response": {
	//     "$ref": "GoogleFirestoreAdminV1beta1ListIndexesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsDatabasesIndexesListCall) Pages(ctx context.Context, f func(*GoogleFirestoreAdminV1beta1ListIndexesResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}
