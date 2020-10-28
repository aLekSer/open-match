// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/messages.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/rpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DoubleRangeFilter_Exclude int32

const (
	// No bounds should be excluded when evaluating the filter, i.e.: MIN <= x <= MAX
	DoubleRangeFilter_NONE DoubleRangeFilter_Exclude = 0
	// Only the minimum bound should be excluded when evaluating the filter, i.e.: MIN < x <= MAX
	DoubleRangeFilter_MIN DoubleRangeFilter_Exclude = 1
	// Only the maximum bound should be excluded when evaluating the filter, i.e.: MIN <= x < MAX
	DoubleRangeFilter_MAX DoubleRangeFilter_Exclude = 2
	// Both bounds should be excluded when evaluating the filter, i.e.: MIN < x < MAX
	DoubleRangeFilter_BOTH DoubleRangeFilter_Exclude = 3
)

var DoubleRangeFilter_Exclude_name = map[int32]string{
	0: "NONE",
	1: "MIN",
	2: "MAX",
	3: "BOTH",
}

var DoubleRangeFilter_Exclude_value = map[string]int32{
	"NONE": 0,
	"MIN":  1,
	"MAX":  2,
	"BOTH": 3,
}

func (x DoubleRangeFilter_Exclude) String() string {
	return proto.EnumName(DoubleRangeFilter_Exclude_name, int32(x))
}

func (DoubleRangeFilter_Exclude) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cb9fb1f207fd5b8c, []int{3, 0}
}

// A Ticket is a basic matchmaking entity in Open Match. A Ticket may represent
// an individual 'Player', a 'Group' of players, or any other concepts unique to
// your use case. Open Match will not interpret what the Ticket represents but
// just treat it as a matchmaking unit with a set of SearchFields. Open Match
// stores the Ticket in state storage and enables an Assignment to be set on the
// Ticket.
type Ticket struct {
	// Id represents an auto-generated Id issued by Open Match.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// An Assignment represents a game server assignment associated with a Ticket,
	// or whatever finalized matched state means for your use case.
	// Open Match does not require or inspect any fields on Assignment.
	Assignment *Assignment `protobuf:"bytes,3,opt,name=assignment,proto3" json:"assignment,omitempty"`
	// Search fields are the fields which Open Match is aware of, and can be used
	// when specifying filters.
	SearchFields *SearchFields `protobuf:"bytes,4,opt,name=search_fields,json=searchFields,proto3" json:"search_fields,omitempty"`
	// Customized information not inspected by Open Match, to be used by the match
	// making function, evaluator, and components making calls to Open Match.
	// Optional, depending on the requirements of the connected systems.
	Extensions map[string]*any.Any `protobuf:"bytes,5,rep,name=extensions,proto3" json:"extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Create time is the time the Ticket was created. It is populated by Open
	// Match at the time of Ticket creation.
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Ticket) Reset()         { *m = Ticket{} }
func (m *Ticket) String() string { return proto.CompactTextString(m) }
func (*Ticket) ProtoMessage()    {}
func (*Ticket) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb9fb1f207fd5b8c, []int{0}
}

func (m *Ticket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ticket.Unmarshal(m, b)
}
func (m *Ticket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ticket.Marshal(b, m, deterministic)
}
func (m *Ticket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ticket.Merge(m, src)
}
func (m *Ticket) XXX_Size() int {
	return xxx_messageInfo_Ticket.Size(m)
}
func (m *Ticket) XXX_DiscardUnknown() {
	xxx_messageInfo_Ticket.DiscardUnknown(m)
}

var xxx_messageInfo_Ticket proto.InternalMessageInfo

func (m *Ticket) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Ticket) GetAssignment() *Assignment {
	if m != nil {
		return m.Assignment
	}
	return nil
}

func (m *Ticket) GetSearchFields() *SearchFields {
	if m != nil {
		return m.SearchFields
	}
	return nil
}

func (m *Ticket) GetExtensions() map[string]*any.Any {
	if m != nil {
		return m.Extensions
	}
	return nil
}

func (m *Ticket) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

// Search fields are the fields which Open Match is aware of, and can be used
// when specifying filters.
type SearchFields struct {
	// Float arguments.  Filterable on ranges.
	DoubleArgs map[string]float64 `protobuf:"bytes,1,rep,name=double_args,json=doubleArgs,proto3" json:"double_args,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	// String arguments.  Filterable on equality.
	StringArgs map[string]string `protobuf:"bytes,2,rep,name=string_args,json=stringArgs,proto3" json:"string_args,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Filterable on presence or absence of given value.
	Tags                 []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchFields) Reset()         { *m = SearchFields{} }
func (m *SearchFields) String() string { return proto.CompactTextString(m) }
func (*SearchFields) ProtoMessage()    {}
func (*SearchFields) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb9fb1f207fd5b8c, []int{1}
}

func (m *SearchFields) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchFields.Unmarshal(m, b)
}
func (m *SearchFields) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchFields.Marshal(b, m, deterministic)
}
func (m *SearchFields) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchFields.Merge(m, src)
}
func (m *SearchFields) XXX_Size() int {
	return xxx_messageInfo_SearchFields.Size(m)
}
func (m *SearchFields) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchFields.DiscardUnknown(m)
}

var xxx_messageInfo_SearchFields proto.InternalMessageInfo

func (m *SearchFields) GetDoubleArgs() map[string]float64 {
	if m != nil {
		return m.DoubleArgs
	}
	return nil
}

func (m *SearchFields) GetStringArgs() map[string]string {
	if m != nil {
		return m.StringArgs
	}
	return nil
}

func (m *SearchFields) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// An Assignment represents a game server assignment associated with a Ticket.
// Open Match does not require or inspect any fields on assignment.
type Assignment struct {
	// Connection information for this Assignment.
	Connection string `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	// Customized information not inspected by Open Match, to be used by the match
	// making function, evaluator, and components making calls to Open Match.
	// Optional, depending on the requirements of the connected systems.
	Extensions           map[string]*any.Any `protobuf:"bytes,4,rep,name=extensions,proto3" json:"extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Assignment) Reset()         { *m = Assignment{} }
func (m *Assignment) String() string { return proto.CompactTextString(m) }
func (*Assignment) ProtoMessage()    {}
func (*Assignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb9fb1f207fd5b8c, []int{2}
}

func (m *Assignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Assignment.Unmarshal(m, b)
}
func (m *Assignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Assignment.Marshal(b, m, deterministic)
}
func (m *Assignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Assignment.Merge(m, src)
}
func (m *Assignment) XXX_Size() int {
	return xxx_messageInfo_Assignment.Size(m)
}
func (m *Assignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Assignment.DiscardUnknown(m)
}

var xxx_messageInfo_Assignment proto.InternalMessageInfo

func (m *Assignment) GetConnection() string {
	if m != nil {
		return m.Connection
	}
	return ""
}

func (m *Assignment) GetExtensions() map[string]*any.Any {
	if m != nil {
		return m.Extensions
	}
	return nil
}

// Filters numerical values to only those within a range.
//   double_arg: "foo"
//   max: 10
//   min: 5
// matches:
//   {"foo": 5}
//   {"foo": 7.5}
//   {"foo": 10}
// does not match:
//   {"foo": 4}
//   {"foo": 10.01}
//   {"foo": "7.5"}
//   {}
type DoubleRangeFilter struct {
	// Name of the ticket's search_fields.double_args this Filter operates on.
	DoubleArg string `protobuf:"bytes,1,opt,name=double_arg,json=doubleArg,proto3" json:"double_arg,omitempty"`
	// Maximum value.
	Max float64 `protobuf:"fixed64,2,opt,name=max,proto3" json:"max,omitempty"`
	// Minimum value.
	Min float64 `protobuf:"fixed64,3,opt,name=min,proto3" json:"min,omitempty"`
	// Which bounds would be excluded when comparing with a ticket's search_fields.double_args value.
	//
	// BETA FEATURE WARNING: This field and the associated values are
	// not finalized and still subject to possible change or removal.
	Exclude              DoubleRangeFilter_Exclude `protobuf:"varint,4,opt,name=exclude,proto3,enum=openmatch.DoubleRangeFilter_Exclude" json:"exclude,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *DoubleRangeFilter) Reset()         { *m = DoubleRangeFilter{} }
func (m *DoubleRangeFilter) String() string { return proto.CompactTextString(m) }
func (*DoubleRangeFilter) ProtoMessage()    {}
func (*DoubleRangeFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb9fb1f207fd5b8c, []int{3}
}

func (m *DoubleRangeFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoubleRangeFilter.Unmarshal(m, b)
}
func (m *DoubleRangeFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoubleRangeFilter.Marshal(b, m, deterministic)
}
func (m *DoubleRangeFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoubleRangeFilter.Merge(m, src)
}
func (m *DoubleRangeFilter) XXX_Size() int {
	return xxx_messageInfo_DoubleRangeFilter.Size(m)
}
func (m *DoubleRangeFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_DoubleRangeFilter.DiscardUnknown(m)
}

var xxx_messageInfo_DoubleRangeFilter proto.InternalMessageInfo

func (m *DoubleRangeFilter) GetDoubleArg() string {
	if m != nil {
		return m.DoubleArg
	}
	return ""
}

func (m *DoubleRangeFilter) GetMax() float64 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *DoubleRangeFilter) GetMin() float64 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *DoubleRangeFilter) GetExclude() DoubleRangeFilter_Exclude {
	if m != nil {
		return m.Exclude
	}
	return DoubleRangeFilter_NONE
}

// Filters strings exactly equaling a value.
//   string_arg: "foo"
//   value: "bar"
// matches:
//   {"foo": "bar"}
// does not match:
//   {"foo": "baz"}
//   {"bar": "foo"}
//   {}
type StringEqualsFilter struct {
	// Name of the ticket's search_fields.string_args this Filter operates on.
	StringArg            string   `protobuf:"bytes,1,opt,name=string_arg,json=stringArg,proto3" json:"string_arg,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringEqualsFilter) Reset()         { *m = StringEqualsFilter{} }
func (m *StringEqualsFilter) String() string { return proto.CompactTextString(m) }
func (*StringEqualsFilter) ProtoMessage()    {}
func (*StringEqualsFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb9fb1f207fd5b8c, []int{4}
}

func (m *StringEqualsFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringEqualsFilter.Unmarshal(m, b)
}
func (m *StringEqualsFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringEqualsFilter.Marshal(b, m, deterministic)
}
func (m *StringEqualsFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringEqualsFilter.Merge(m, src)
}
func (m *StringEqualsFilter) XXX_Size() int {
	return xxx_messageInfo_StringEqualsFilter.Size(m)
}
func (m *StringEqualsFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_StringEqualsFilter.DiscardUnknown(m)
}

var xxx_messageInfo_StringEqualsFilter proto.InternalMessageInfo

func (m *StringEqualsFilter) GetStringArg() string {
	if m != nil {
		return m.StringArg
	}
	return ""
}

func (m *StringEqualsFilter) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Filters to the tag being present on the search_fields.
//   tag: "foo"
// matches:
//   ["foo"]
//   ["bar","foo"]
// does not match:
//   ["bar"]
//   []
type TagPresentFilter struct {
	Tag                  string   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TagPresentFilter) Reset()         { *m = TagPresentFilter{} }
func (m *TagPresentFilter) String() string { return proto.CompactTextString(m) }
func (*TagPresentFilter) ProtoMessage()    {}
func (*TagPresentFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb9fb1f207fd5b8c, []int{5}
}

func (m *TagPresentFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagPresentFilter.Unmarshal(m, b)
}
func (m *TagPresentFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagPresentFilter.Marshal(b, m, deterministic)
}
func (m *TagPresentFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagPresentFilter.Merge(m, src)
}
func (m *TagPresentFilter) XXX_Size() int {
	return xxx_messageInfo_TagPresentFilter.Size(m)
}
func (m *TagPresentFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_TagPresentFilter.DiscardUnknown(m)
}

var xxx_messageInfo_TagPresentFilter proto.InternalMessageInfo

func (m *TagPresentFilter) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

// Pool specfies a set of criteria that are used to select a subset of Tickets
// that meet all the criteria.
type Pool struct {
	// A developer-chosen human-readable name for this Pool.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Set of Filters indicating the filtering criteria. Selected tickets must
	// match every Filter.
	DoubleRangeFilters  []*DoubleRangeFilter  `protobuf:"bytes,2,rep,name=double_range_filters,json=doubleRangeFilters,proto3" json:"double_range_filters,omitempty"`
	StringEqualsFilters []*StringEqualsFilter `protobuf:"bytes,4,rep,name=string_equals_filters,json=stringEqualsFilters,proto3" json:"string_equals_filters,omitempty"`
	TagPresentFilters   []*TagPresentFilter   `protobuf:"bytes,5,rep,name=tag_present_filters,json=tagPresentFilters,proto3" json:"tag_present_filters,omitempty"`
	// If specified, only Tickets created before the specified time are selected.
	CreatedBefore *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_before,json=createdBefore,proto3" json:"created_before,omitempty"`
	// If specified, only Tickets created after the specified time are selected.
	CreatedAfter         *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_after,json=createdAfter,proto3" json:"created_after,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb9fb1f207fd5b8c, []int{6}
}

func (m *Pool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pool.Unmarshal(m, b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return xxx_messageInfo_Pool.Size(m)
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

func (m *Pool) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Pool) GetDoubleRangeFilters() []*DoubleRangeFilter {
	if m != nil {
		return m.DoubleRangeFilters
	}
	return nil
}

func (m *Pool) GetStringEqualsFilters() []*StringEqualsFilter {
	if m != nil {
		return m.StringEqualsFilters
	}
	return nil
}

func (m *Pool) GetTagPresentFilters() []*TagPresentFilter {
	if m != nil {
		return m.TagPresentFilters
	}
	return nil
}

func (m *Pool) GetCreatedBefore() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedBefore
	}
	return nil
}

func (m *Pool) GetCreatedAfter() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAfter
	}
	return nil
}

// A MatchProfile is Open Match's representation of a Match specification. It is
// used to indicate the criteria for selecting players for a match. A
// MatchProfile is the input to the API to get matches and is passed to the
// MatchFunction. It contains all the information required by the MatchFunction
// to generate match proposals.
type MatchProfile struct {
	// Name of this match profile.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Set of pools to be queried when generating a match for this MatchProfile.
	Pools []*Pool `protobuf:"bytes,3,rep,name=pools,proto3" json:"pools,omitempty"`
	// Customized information not inspected by Open Match, to be used by the match
	// making function, evaluator, and components making calls to Open Match.
	// Optional, depending on the requirements of the connected systems.
	Extensions           map[string]*any.Any `protobuf:"bytes,5,rep,name=extensions,proto3" json:"extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MatchProfile) Reset()         { *m = MatchProfile{} }
func (m *MatchProfile) String() string { return proto.CompactTextString(m) }
func (*MatchProfile) ProtoMessage()    {}
func (*MatchProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb9fb1f207fd5b8c, []int{7}
}

func (m *MatchProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchProfile.Unmarshal(m, b)
}
func (m *MatchProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchProfile.Marshal(b, m, deterministic)
}
func (m *MatchProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchProfile.Merge(m, src)
}
func (m *MatchProfile) XXX_Size() int {
	return xxx_messageInfo_MatchProfile.Size(m)
}
func (m *MatchProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchProfile.DiscardUnknown(m)
}

var xxx_messageInfo_MatchProfile proto.InternalMessageInfo

func (m *MatchProfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MatchProfile) GetPools() []*Pool {
	if m != nil {
		return m.Pools
	}
	return nil
}

func (m *MatchProfile) GetExtensions() map[string]*any.Any {
	if m != nil {
		return m.Extensions
	}
	return nil
}

// A Match is used to represent a completed match object. It can be generated by
// a MatchFunction as a proposal or can be returned by OpenMatch as a result in
// response to the FetchMatches call.
// When a match is returned by the FetchMatches call, it should contain at least
// one ticket to be considered as valid.
type Match struct {
	// A Match ID that should be passed through the stack for tracing.
	MatchId string `protobuf:"bytes,1,opt,name=match_id,json=matchId,proto3" json:"match_id,omitempty"`
	// Name of the match profile that generated this Match.
	MatchProfile string `protobuf:"bytes,2,opt,name=match_profile,json=matchProfile,proto3" json:"match_profile,omitempty"`
	// Name of the match function that generated this Match.
	MatchFunction string `protobuf:"bytes,3,opt,name=match_function,json=matchFunction,proto3" json:"match_function,omitempty"`
	// Tickets belonging to this match.
	Tickets []*Ticket `protobuf:"bytes,4,rep,name=tickets,proto3" json:"tickets,omitempty"`
	// Customized information not inspected by Open Match, to be used by the match
	// making function, evaluator, and components making calls to Open Match.
	// Optional, depending on the requirements of the connected systems.
	Extensions map[string]*any.Any `protobuf:"bytes,7,rep,name=extensions,proto3" json:"extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Backfill request which contains additional information to the match
	// and contains an association to a GameServer.
	//
	// BETA FEATURE WARNING: This field is not finalized and still subject
	// to possible change or removal.
	Backfill *Backfill `protobuf:"bytes,8,opt,name=backfill,proto3" json:"backfill,omitempty"`
	// AllocateGameServer signalise Director that Backfill is new and it should
	// allocate a GameServer, this Backfill would be assigned.
	//
	// BETA FEATURE WARNING: This field is not finalized and still subject
	// to possible change or removal.
	AllocateGameserver   bool     `protobuf:"varint,9,opt,name=allocate_gameserver,json=allocateGameserver,proto3" json:"allocate_gameserver,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Match) Reset()         { *m = Match{} }
func (m *Match) String() string { return proto.CompactTextString(m) }
func (*Match) ProtoMessage()    {}
func (*Match) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb9fb1f207fd5b8c, []int{8}
}

func (m *Match) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Match.Unmarshal(m, b)
}
func (m *Match) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Match.Marshal(b, m, deterministic)
}
func (m *Match) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Match.Merge(m, src)
}
func (m *Match) XXX_Size() int {
	return xxx_messageInfo_Match.Size(m)
}
func (m *Match) XXX_DiscardUnknown() {
	xxx_messageInfo_Match.DiscardUnknown(m)
}

var xxx_messageInfo_Match proto.InternalMessageInfo

func (m *Match) GetMatchId() string {
	if m != nil {
		return m.MatchId
	}
	return ""
}

func (m *Match) GetMatchProfile() string {
	if m != nil {
		return m.MatchProfile
	}
	return ""
}

func (m *Match) GetMatchFunction() string {
	if m != nil {
		return m.MatchFunction
	}
	return ""
}

func (m *Match) GetTickets() []*Ticket {
	if m != nil {
		return m.Tickets
	}
	return nil
}

func (m *Match) GetExtensions() map[string]*any.Any {
	if m != nil {
		return m.Extensions
	}
	return nil
}

func (m *Match) GetBackfill() *Backfill {
	if m != nil {
		return m.Backfill
	}
	return nil
}

func (m *Match) GetAllocateGameserver() bool {
	if m != nil {
		return m.AllocateGameserver
	}
	return false
}

// BETA FEATURE WARNING:  This call and the associated Request and Response
// messages are not finalized and still subject to possible change or removal.
// Represents a backfill entity which is used to fill partially full matches.
type Backfill struct {
	// Id represents an auto-generated ID issued by Open Match.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Search fields are the fields which Open Match is aware of, and can be used when specifying filters.
	SearchFields *SearchFields `protobuf:"bytes,2,opt,name=search_fields,json=searchFields,proto3" json:"search_fields,omitempty"`
	// Customized information not inspected by Open Match, to be used by the match making function, evaluator, and components making calls to Open Match.
	// Optional, depending on the requirements of the connected systems.
	Extensions map[string]*any.Any `protobuf:"bytes,3,rep,name=extensions,proto3" json:"extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Create time is the time the Ticket was created. It is populated by Open
	// Match at the time of Ticket creation.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Generation gets incremented on GameServers update operations.
	// Prevents the MMF from overriding a newer version from the game server.
	// It is not supposed to be updated by the MMF.
	Generation           int64    `protobuf:"varint,5,opt,name=generation,proto3" json:"generation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Backfill) Reset()         { *m = Backfill{} }
func (m *Backfill) String() string { return proto.CompactTextString(m) }
func (*Backfill) ProtoMessage()    {}
func (*Backfill) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb9fb1f207fd5b8c, []int{9}
}

func (m *Backfill) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Backfill.Unmarshal(m, b)
}
func (m *Backfill) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Backfill.Marshal(b, m, deterministic)
}
func (m *Backfill) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Backfill.Merge(m, src)
}
func (m *Backfill) XXX_Size() int {
	return xxx_messageInfo_Backfill.Size(m)
}
func (m *Backfill) XXX_DiscardUnknown() {
	xxx_messageInfo_Backfill.DiscardUnknown(m)
}

var xxx_messageInfo_Backfill proto.InternalMessageInfo

func (m *Backfill) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Backfill) GetSearchFields() *SearchFields {
	if m != nil {
		return m.SearchFields
	}
	return nil
}

func (m *Backfill) GetExtensions() map[string]*any.Any {
	if m != nil {
		return m.Extensions
	}
	return nil
}

func (m *Backfill) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Backfill) GetGeneration() int64 {
	if m != nil {
		return m.Generation
	}
	return 0
}

func init() {
	proto.RegisterEnum("openmatch.DoubleRangeFilter_Exclude", DoubleRangeFilter_Exclude_name, DoubleRangeFilter_Exclude_value)
	proto.RegisterType((*Ticket)(nil), "openmatch.Ticket")
	proto.RegisterMapType((map[string]*any.Any)(nil), "openmatch.Ticket.ExtensionsEntry")
	proto.RegisterType((*SearchFields)(nil), "openmatch.SearchFields")
	proto.RegisterMapType((map[string]float64)(nil), "openmatch.SearchFields.DoubleArgsEntry")
	proto.RegisterMapType((map[string]string)(nil), "openmatch.SearchFields.StringArgsEntry")
	proto.RegisterType((*Assignment)(nil), "openmatch.Assignment")
	proto.RegisterMapType((map[string]*any.Any)(nil), "openmatch.Assignment.ExtensionsEntry")
	proto.RegisterType((*DoubleRangeFilter)(nil), "openmatch.DoubleRangeFilter")
	proto.RegisterType((*StringEqualsFilter)(nil), "openmatch.StringEqualsFilter")
	proto.RegisterType((*TagPresentFilter)(nil), "openmatch.TagPresentFilter")
	proto.RegisterType((*Pool)(nil), "openmatch.Pool")
	proto.RegisterType((*MatchProfile)(nil), "openmatch.MatchProfile")
	proto.RegisterMapType((map[string]*any.Any)(nil), "openmatch.MatchProfile.ExtensionsEntry")
	proto.RegisterType((*Match)(nil), "openmatch.Match")
	proto.RegisterMapType((map[string]*any.Any)(nil), "openmatch.Match.ExtensionsEntry")
	proto.RegisterType((*Backfill)(nil), "openmatch.Backfill")
	proto.RegisterMapType((map[string]*any.Any)(nil), "openmatch.Backfill.ExtensionsEntry")
}

func init() { proto.RegisterFile("api/messages.proto", fileDescriptor_cb9fb1f207fd5b8c) }

var fileDescriptor_cb9fb1f207fd5b8c = []byte{
	// 992 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0x4d, 0x73, 0xe3, 0x34,
	0x18, 0xc6, 0x1f, 0x69, 0x92, 0xb7, 0xd9, 0xd6, 0x55, 0xbb, 0xb3, 0xde, 0xc0, 0x42, 0xf0, 0x6e,
	0x87, 0x0c, 0x0c, 0xf6, 0x4c, 0x19, 0x66, 0x18, 0xbe, 0x53, 0x48, 0x77, 0x5b, 0x66, 0xdb, 0xe2,
	0xf6, 0xc0, 0x70, 0xc9, 0x28, 0xb6, 0xe2, 0xf5, 0xd4, 0x91, 0x8d, 0xa5, 0x74, 0xda, 0xff, 0xc1,
	0x0f, 0xe0, 0xcc, 0x99, 0xff, 0xc0, 0x91, 0x0b, 0x3f, 0x82, 0x3b, 0x7f, 0x80, 0xb1, 0x64, 0x3b,
	0x6a, 0x12, 0xba, 0xf4, 0xd0, 0xe1, 0x26, 0xbd, 0x5f, 0xd2, 0xf3, 0xe8, 0xd1, 0x2b, 0x01, 0xc2,
	0x59, 0xec, 0x4d, 0x09, 0x63, 0x38, 0x22, 0xcc, 0xcd, 0xf2, 0x94, 0xa7, 0xa8, 0x9d, 0x66, 0x84,
	0x4e, 0x31, 0x0f, 0x5e, 0x75, 0x1f, 0x45, 0x69, 0x1a, 0x25, 0xc4, 0xcb, 0xb3, 0xc0, 0x63, 0x1c,
	0xf3, 0x59, 0x19, 0xd3, 0x7d, 0x5c, 0x3a, 0xc4, 0x6c, 0x3c, 0x9b, 0x78, 0x98, 0x5e, 0x97, 0xae,
	0x77, 0x16, 0x5d, 0x3c, 0x9e, 0x12, 0xc6, 0xf1, 0x34, 0x93, 0x01, 0xce, 0x5f, 0x3a, 0xac, 0x9d,
	0xc7, 0xc1, 0x05, 0xe1, 0x68, 0x03, 0xf4, 0x38, 0xb4, 0xb5, 0x9e, 0xd6, 0x6f, 0xfb, 0x7a, 0x1c,
	0xa2, 0x8f, 0x01, 0x30, 0x63, 0x71, 0x44, 0xa7, 0x84, 0x72, 0xdb, 0xe8, 0x69, 0xfd, 0xf5, 0xbd,
	0x87, 0x6e, 0xbd, 0x1f, 0x77, 0x50, 0x3b, 0x7d, 0x25, 0x10, 0x7d, 0x0e, 0x0f, 0x18, 0xc1, 0x79,
	0xf0, 0x6a, 0x34, 0x89, 0x49, 0x12, 0x32, 0xdb, 0x14, 0x99, 0x8f, 0x94, 0xcc, 0x33, 0xe1, 0x3f,
	0x10, 0x6e, 0xbf, 0xc3, 0x94, 0x19, 0x1a, 0x00, 0x90, 0x2b, 0x4e, 0x28, 0x8b, 0x53, 0xca, 0xec,
	0x46, 0xcf, 0xe8, 0xaf, 0xef, 0xbd, 0xab, 0xa4, 0xca, 0xbd, 0xba, 0xc3, 0x3a, 0x66, 0x48, 0x79,
	0x7e, 0xed, 0x2b, 0x49, 0xe8, 0x33, 0x58, 0x0f, 0x72, 0x82, 0x39, 0x19, 0x15, 0x60, 0xed, 0x35,
	0xb1, 0x7c, 0xd7, 0x95, 0x4c, 0xb8, 0x15, 0x13, 0xee, 0x79, 0xc5, 0x84, 0x0f, 0x32, 0xbc, 0x30,
	0x74, 0xcf, 0x60, 0x73, 0xa1, 0x36, 0xb2, 0xc0, 0xb8, 0x20, 0xd7, 0x25, 0x31, 0xc5, 0x10, 0xbd,
	0x0f, 0x8d, 0x4b, 0x9c, 0xcc, 0x88, 0xad, 0x8b, 0xda, 0x3b, 0x4b, 0xb5, 0x07, 0xf4, 0xda, 0x97,
	0x21, 0x9f, 0xea, 0x9f, 0x68, 0x47, 0x66, 0x4b, 0xb7, 0x0c, 0xe7, 0x37, 0x1d, 0x3a, 0x2a, 0x72,
	0xf4, 0x02, 0xd6, 0xc3, 0x74, 0x36, 0x4e, 0xc8, 0x08, 0xe7, 0x11, 0xb3, 0x35, 0x01, 0xf6, 0xbd,
	0x7f, 0xe1, 0xc9, 0xfd, 0x56, 0x84, 0x0e, 0xf2, 0xa8, 0x82, 0x1c, 0xd6, 0x86, 0xa2, 0x12, 0xe3,
	0x79, 0x4c, 0x23, 0x59, 0x49, 0xbf, 0xbd, 0xd2, 0x99, 0x08, 0x55, 0x2a, 0xb1, 0xda, 0x80, 0x10,
	0x98, 0x1c, 0x47, 0xcc, 0x36, 0x7a, 0x46, 0xbf, 0xed, 0x8b, 0x71, 0xf7, 0x0b, 0xd8, 0x5c, 0x58,
	0x7c, 0x05, 0x27, 0x3b, 0x2a, 0x27, 0x9a, 0x82, 0xbe, 0x48, 0x5f, 0x58, 0xf1, 0x75, 0xe9, 0x6d,
	0x25, 0xdd, 0xf9, 0x53, 0x03, 0x98, 0x4b, 0x0d, 0xbd, 0x0d, 0x10, 0xa4, 0x94, 0x92, 0x80, 0xc7,
	0x29, 0x2d, 0x2b, 0x28, 0x16, 0x34, 0xbc, 0x21, 0x20, 0x53, 0x30, 0xb1, 0xbb, 0x52, 0xb5, 0xb7,
	0x89, 0xe8, 0x1e, 0x75, 0x70, 0x64, 0xb6, 0x0c, 0xcb, 0x74, 0x7e, 0xd7, 0x60, 0x4b, 0xb2, 0xea,
	0x63, 0x1a, 0x91, 0x83, 0x38, 0xe1, 0x24, 0x47, 0x4f, 0x00, 0xe6, 0x92, 0x28, 0x97, 0x6a, 0xd7,
	0x07, 0x5d, 0x6c, 0x61, 0x8a, 0xaf, 0x4a, 0x8a, 0x8b, 0xa1, 0xb0, 0xc4, 0x54, 0xdc, 0xce, 0xc2,
	0x12, 0x53, 0xf4, 0x25, 0x34, 0xc9, 0x55, 0x90, 0xcc, 0x42, 0x22, 0x6e, 0xde, 0xc6, 0xde, 0x33,
	0x05, 0xfd, 0xd2, 0x8a, 0xee, 0x50, 0xc6, 0xfa, 0x55, 0x92, 0xe3, 0x41, 0xb3, 0xb4, 0xa1, 0x16,
	0x98, 0xc7, 0x27, 0xc7, 0x43, 0xeb, 0x0d, 0xd4, 0x04, 0xe3, 0xe5, 0xe1, 0xb1, 0xa5, 0x89, 0xc1,
	0xe0, 0x07, 0x4b, 0x2f, 0x7c, 0xfb, 0x27, 0xe7, 0x2f, 0x2c, 0xc3, 0x39, 0x04, 0x24, 0xcf, 0x77,
	0xf8, 0xd3, 0x0c, 0x27, 0x6c, 0x8e, 0x64, 0x2e, 0xc9, 0x0a, 0x49, 0x2d, 0xb4, 0xd5, 0xe7, 0xed,
	0x3c, 0x03, 0xeb, 0x1c, 0x47, 0xa7, 0x39, 0x61, 0x84, 0xf2, 0xb2, 0x90, 0x05, 0x06, 0xc7, 0x55,
	0x85, 0x62, 0xe8, 0xfc, 0x6c, 0x80, 0x79, 0x9a, 0xa6, 0x49, 0x21, 0x56, 0x8a, 0xa7, 0xa4, 0xf4,
	0x89, 0x31, 0x3a, 0x86, 0x9d, 0x92, 0xc1, 0xbc, 0x40, 0x39, 0x9a, 0x88, 0x2a, 0xd5, 0x9d, 0x78,
	0xeb, 0x36, 0x2e, 0x7c, 0x14, 0x2e, 0x9a, 0x18, 0xfa, 0x1e, 0x1e, 0x96, 0x38, 0x88, 0x80, 0x57,
	0x17, 0x94, 0xd2, 0x7a, 0xa2, 0x5e, 0xb2, 0x25, 0x16, 0xfc, 0x6d, 0xb6, 0x64, 0x63, 0xe8, 0x3b,
	0xd8, 0xe6, 0x38, 0x1a, 0x65, 0x12, 0x66, 0x5d, 0x50, 0x36, 0xbb, 0x37, 0xd5, 0x66, 0xb7, 0xc0,
	0x85, 0xbf, 0xc5, 0x17, 0x2c, 0x45, 0xc3, 0xdc, 0x90, 0xed, 0x2b, 0x1c, 0x8d, 0xc9, 0x24, 0xcd,
	0xff, 0x4b, 0xc3, 0x7b, 0x50, 0x66, 0xec, 0x8b, 0x04, 0xf4, 0x15, 0x54, 0x86, 0x11, 0x9e, 0x70,
	0x92, 0xdb, 0xcd, 0xd7, 0x56, 0xe8, 0x94, 0x09, 0x83, 0x22, 0xbe, 0x54, 0xf4, 0xdf, 0x1a, 0x74,
	0x5e, 0x16, 0xfb, 0x3e, 0xcd, 0xd3, 0x49, 0x9c, 0x90, 0x95, 0xc7, 0xb3, 0x0b, 0x8d, 0x2c, 0x4d,
	0x13, 0xd9, 0x60, 0xd6, 0xf7, 0x36, 0x15, 0xb4, 0xc5, 0x91, 0xfa, 0xd2, 0x8b, 0x9e, 0xaf, 0x78,
	0x06, 0xd4, 0x7e, 0xa6, 0xae, 0xf3, 0xff, 0xdd, 0x63, 0xd3, 0x6a, 0x38, 0xbf, 0x18, 0xd0, 0x10,
	0xbb, 0x41, 0x8f, 0xa1, 0x25, 0x36, 0x37, 0xaa, 0x5f, 0xd1, 0xa6, 0x98, 0x1f, 0x86, 0xe8, 0x29,
	0x3c, 0x90, 0xae, 0x4c, 0x6e, 0xb9, 0x54, 0x7d, 0x67, 0xaa, 0xd2, 0xb5, 0x0b, 0x1b, 0x32, 0x68,
	0x32, 0xa3, 0xb2, 0xbb, 0x19, 0x22, 0x4a, 0xa6, 0x1e, 0x94, 0x46, 0xf4, 0x01, 0x34, 0xb9, 0x78,
	0x04, 0x2b, 0x09, 0x6e, 0x2d, 0x3d, 0x8f, 0x7e, 0x15, 0x81, 0xbe, 0xbe, 0xc1, 0x63, 0x53, 0xc4,
	0xf7, 0x16, 0x79, 0xbc, 0xf5, 0x35, 0xf5, 0xa0, 0x35, 0xc6, 0xc1, 0xc5, 0x24, 0x4e, 0x12, 0xbb,
	0x25, 0xe8, 0xd9, 0x56, 0xf2, 0xf7, 0x4b, 0x97, 0x5f, 0x07, 0x21, 0x0f, 0xb6, 0x71, 0x92, 0xa4,
	0x41, 0xf1, 0x00, 0x47, 0x78, 0x4a, 0x18, 0xc9, 0x2f, 0x49, 0x6e, 0xb7, 0x7b, 0x5a, 0xbf, 0xe5,
	0xa3, 0xca, 0xf5, 0xbc, 0xf6, 0xdc, 0xd7, 0x11, 0x35, 0xac, 0xb5, 0x23, 0xb3, 0xb5, 0x66, 0x35,
	0x9d, 0x3f, 0x74, 0x68, 0x55, 0x1b, 0x5d, 0xfa, 0xe5, 0x2c, 0x7d, 0x57, 0xf4, 0xbb, 0x7c, 0x57,
	0xbe, 0xb9, 0xc1, 0xaf, 0xd4, 0xf4, 0xd3, 0x15, 0xfc, 0xdc, 0xe5, 0xc3, 0x62, 0xde, 0xe5, 0xc3,
	0x52, 0xbc, 0x87, 0x11, 0xa1, 0x24, 0xc7, 0x42, 0x31, 0x8d, 0x9e, 0xd6, 0x37, 0x7c, 0xc5, 0x72,
	0x2f, 0xec, 0xee, 0xbb, 0x3f, 0xf6, 0x0a, 0x8c, 0x1f, 0x4a, 0x90, 0x21, 0xb9, 0xf4, 0xe6, 0x53,
	0x2f, 0xbb, 0x88, 0xbc, 0x6c, 0xfc, 0xab, 0xde, 0x3e, 0xc9, 0x08, 0x15, 0x02, 0x1b, 0xaf, 0x89,
	0x42, 0x1f, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xaa, 0xfb, 0x6f, 0xe2, 0x0a, 0x00, 0x00,
}
