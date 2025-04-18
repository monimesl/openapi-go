// Package openapi3 contains JSON mapping structures.
package openapi3

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	orderedmap "github.com/wk8/go-ordered-map/v2"
	"reflect"
	"regexp"
)

// Spec structure is generated from "#".
//
// Validation schema for OpenAPI Specification 3.0.X.
type Spec struct {
	// Value must match pattern: `^3\.0\.\d(-.+)?$`.
	// Required.
	Openapi       string                 `json:"openapi"`
	Info          Info                   `json:"info"` // Required.
	ExternalDocs  *ExternalDocumentation `json:"externalDocs,omitempty"`
	Servers       []Server               `json:"servers,omitempty"`
	Security      []map[string][]string  `json:"security,omitempty"`
	Tags          []Tag                  `json:"tags,omitempty"`
	Paths         Paths                  `json:"paths"` // Required.
	Components    *Components            `json:"components,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"` // Key must match pattern: `^x-`.
}

// WithOpenapi sets Openapi value.
func (s *Spec) WithOpenapi(val string) *Spec {
	s.Openapi = val
	return s
}

// WithInfo sets Info value.
func (s *Spec) WithInfo(val Info) *Spec {
	s.Info = val
	return s
}

// WithExternalDocs sets ExternalDocs value.
func (s *Spec) WithExternalDocs(val ExternalDocumentation) *Spec {
	s.ExternalDocs = &val
	return s
}

// ExternalDocsEns ensures returned ExternalDocs is not nil.
func (s *Spec) ExternalDocsEns() *ExternalDocumentation {
	if s.ExternalDocs == nil {
		s.ExternalDocs = new(ExternalDocumentation)
	}

	return s.ExternalDocs
}

// WithServers sets Servers value.
func (s *Spec) WithServers(val ...Server) *Spec {
	s.Servers = val
	return s
}

// WithSecurity sets Security value.
func (s *Spec) WithSecurity(val ...map[string][]string) *Spec {
	s.Security = val
	return s
}

// WithTags sets Tags value.
func (s *Spec) WithTags(val ...Tag) *Spec {
	s.Tags = val
	return s
}

// WithPaths sets Paths value.
func (s *Spec) WithPaths(val Paths) *Spec {
	s.Paths = val
	return s
}

// WithComponents sets Components value.
func (s *Spec) WithComponents(val Components) *Spec {
	s.Components = &val
	return s
}

// ComponentsEns ensures returned Components is not nil.
func (s *Spec) ComponentsEns() *Components {
	if s.Components == nil {
		s.Components = new(Components)
	}

	return s.Components
}

// WithMapOfAnything sets MapOfAnything value.
func (s *Spec) WithMapOfAnything(val map[string]interface{}) *Spec {
	s.MapOfAnything = val
	return s
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (s *Spec) WithMapOfAnythingItem(key string, val interface{}) *Spec {
	if s.MapOfAnything == nil {
		s.MapOfAnything = make(map[string]interface{}, 1)
	}

	s.MapOfAnything[key] = val

	return s
}

type marshalSpec Spec

var knownKeysSpec = []string{
	"openapi",
	"info",
	"externalDocs",
	"servers",
	"security",
	"tags",
	"paths",
	"components",
}

var requireKeysSpec = []string{
	"openapi",
	"info",
	"paths",
}

// UnmarshalJSON decodes JSON.
func (s *Spec) UnmarshalJSON(data []byte) error {
	var err error

	ms := marshalSpec(*s)

	err = json.Unmarshal(data, &ms)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysSpec {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysSpec {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ms.MapOfAnything == nil {
				ms.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ms.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Spec: %v", offendingKeys)
	}

	*s = Spec(ms)

	return nil
}

// MarshalJSON encodes JSON.
func (s Spec) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalSpec(s), s.MapOfAnything)
}

// Info structure is generated from "#/definitions/Info".
type Info struct {
	Title          string                 `json:"title"` // Required.
	Description    *string                `json:"description,omitempty"`
	TermsOfService *string                `json:"termsOfService,omitempty"` // Format: uri-reference.
	Contact        *Contact               `json:"contact,omitempty"`
	License        *License               `json:"license,omitempty"`
	Version        string                 `json:"version"` // Required.
	MapOfAnything  map[string]interface{} `json:"-"`       // Key must match pattern: `^x-`.
}

// WithTitle sets Title value.
func (i *Info) WithTitle(val string) *Info {
	i.Title = val
	return i
}

// WithDescription sets Description value.
func (i *Info) WithDescription(val string) *Info {
	i.Description = &val
	return i
}

// WithTermsOfService sets TermsOfService value.
func (i *Info) WithTermsOfService(val string) *Info {
	i.TermsOfService = &val
	return i
}

// WithContact sets Contact value.
func (i *Info) WithContact(val Contact) *Info {
	i.Contact = &val
	return i
}

// ContactEns ensures returned Contact is not nil.
func (i *Info) ContactEns() *Contact {
	if i.Contact == nil {
		i.Contact = new(Contact)
	}

	return i.Contact
}

// WithLicense sets License value.
func (i *Info) WithLicense(val License) *Info {
	i.License = &val
	return i
}

// LicenseEns ensures returned License is not nil.
func (i *Info) LicenseEns() *License {
	if i.License == nil {
		i.License = new(License)
	}

	return i.License
}

// WithVersion sets Version value.
func (i *Info) WithVersion(val string) *Info {
	i.Version = val
	return i
}

// WithMapOfAnything sets MapOfAnything value.
func (i *Info) WithMapOfAnything(val map[string]interface{}) *Info {
	i.MapOfAnything = val
	return i
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (i *Info) WithMapOfAnythingItem(key string, val interface{}) *Info {
	if i.MapOfAnything == nil {
		i.MapOfAnything = make(map[string]interface{}, 1)
	}

	i.MapOfAnything[key] = val

	return i
}

type marshalInfo Info

var knownKeysInfo = []string{
	"title",
	"description",
	"termsOfService",
	"contact",
	"license",
	"version",
}

var requireKeysInfo = []string{
	"title",
	"version",
}

// UnmarshalJSON decodes JSON.
func (i *Info) UnmarshalJSON(data []byte) error {
	var err error

	mi := marshalInfo(*i)

	err = json.Unmarshal(data, &mi)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysInfo {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysInfo {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if mi.MapOfAnything == nil {
				mi.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mi.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Info: %v", offendingKeys)
	}

	*i = Info(mi)

	return nil
}

// MarshalJSON encodes JSON.
func (i Info) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalInfo(i), i.MapOfAnything)
}

// Contact structure is generated from "#/definitions/Contact".
type Contact struct {
	Name          *string                `json:"name,omitempty"`
	URL           *string                `json:"url,omitempty"`   // Format: uri-reference.
	Email         *string                `json:"email,omitempty"` // Format: email.
	MapOfAnything map[string]interface{} `json:"-"`               // Key must match pattern: `^x-`.
}

// WithName sets Name value.
func (c *Contact) WithName(val string) *Contact {
	c.Name = &val
	return c
}

// WithURL sets URL value.
func (c *Contact) WithURL(val string) *Contact {
	c.URL = &val
	return c
}

// WithEmail sets Email value.
func (c *Contact) WithEmail(val string) *Contact {
	c.Email = &val
	return c
}

// WithMapOfAnything sets MapOfAnything value.
func (c *Contact) WithMapOfAnything(val map[string]interface{}) *Contact {
	c.MapOfAnything = val
	return c
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (c *Contact) WithMapOfAnythingItem(key string, val interface{}) *Contact {
	if c.MapOfAnything == nil {
		c.MapOfAnything = make(map[string]interface{}, 1)
	}

	c.MapOfAnything[key] = val

	return c
}

type marshalContact Contact

var knownKeysContact = []string{
	"name",
	"url",
	"email",
}

// UnmarshalJSON decodes JSON.
func (c *Contact) UnmarshalJSON(data []byte) error {
	var err error

	mc := marshalContact(*c)

	err = json.Unmarshal(data, &mc)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range knownKeysContact {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if mc.MapOfAnything == nil {
				mc.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mc.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Contact: %v", offendingKeys)
	}

	*c = Contact(mc)

	return nil
}

// MarshalJSON encodes JSON.
func (c Contact) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalContact(c), c.MapOfAnything)
}

// License structure is generated from "#/definitions/License".
type License struct {
	Name          string                 `json:"name"`          // Required.
	URL           *string                `json:"url,omitempty"` // Format: uri-reference.
	MapOfAnything map[string]interface{} `json:"-"`             // Key must match pattern: `^x-`.
}

// WithName sets Name value.
func (l *License) WithName(val string) *License {
	l.Name = val
	return l
}

// WithURL sets URL value.
func (l *License) WithURL(val string) *License {
	l.URL = &val
	return l
}

// WithMapOfAnything sets MapOfAnything value.
func (l *License) WithMapOfAnything(val map[string]interface{}) *License {
	l.MapOfAnything = val
	return l
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (l *License) WithMapOfAnythingItem(key string, val interface{}) *License {
	if l.MapOfAnything == nil {
		l.MapOfAnything = make(map[string]interface{}, 1)
	}

	l.MapOfAnything[key] = val

	return l
}

type marshalLicense License

var knownKeysLicense = []string{
	"name",
	"url",
}

var requireKeysLicense = []string{
	"name",
}

// UnmarshalJSON decodes JSON.
func (l *License) UnmarshalJSON(data []byte) error {
	var err error

	ml := marshalLicense(*l)

	err = json.Unmarshal(data, &ml)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysLicense {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysLicense {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ml.MapOfAnything == nil {
				ml.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ml.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in License: %v", offendingKeys)
	}

	*l = License(ml)

	return nil
}

// MarshalJSON encodes JSON.
func (l License) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalLicense(l), l.MapOfAnything)
}

// ExternalDocumentation structure is generated from "#/definitions/ExternalDocumentation".
type ExternalDocumentation struct {
	Description *string `json:"description,omitempty"`
	// Format: uri-reference.
	// Required.
	URL           string                 `json:"url"`
	MapOfAnything map[string]interface{} `json:"-"` // Key must match pattern: `^x-`.
}

// WithDescription sets Description value.
func (e *ExternalDocumentation) WithDescription(val string) *ExternalDocumentation {
	e.Description = &val
	return e
}

// WithURL sets URL value.
func (e *ExternalDocumentation) WithURL(val string) *ExternalDocumentation {
	e.URL = val
	return e
}

// WithMapOfAnything sets MapOfAnything value.
func (e *ExternalDocumentation) WithMapOfAnything(val map[string]interface{}) *ExternalDocumentation {
	e.MapOfAnything = val
	return e
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (e *ExternalDocumentation) WithMapOfAnythingItem(key string, val interface{}) *ExternalDocumentation {
	if e.MapOfAnything == nil {
		e.MapOfAnything = make(map[string]interface{}, 1)
	}

	e.MapOfAnything[key] = val

	return e
}

type marshalExternalDocumentation ExternalDocumentation

var knownKeysExternalDocumentation = []string{
	"description",
	"url",
}

var requireKeysExternalDocumentation = []string{
	"url",
}

// UnmarshalJSON decodes JSON.
func (e *ExternalDocumentation) UnmarshalJSON(data []byte) error {
	var err error

	me := marshalExternalDocumentation(*e)

	err = json.Unmarshal(data, &me)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysExternalDocumentation {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysExternalDocumentation {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if me.MapOfAnything == nil {
				me.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			me.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in ExternalDocumentation: %v", offendingKeys)
	}

	*e = ExternalDocumentation(me)

	return nil
}

// MarshalJSON encodes JSON.
func (e ExternalDocumentation) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalExternalDocumentation(e), e.MapOfAnything)
}

// Server structure is generated from "#/definitions/Server".
type Server struct {
	URL           string                    `json:"url"` // Required.
	Description   *string                   `json:"description,omitempty"`
	Variables     map[string]ServerVariable `json:"variables,omitempty"`
	MapOfAnything map[string]interface{}    `json:"-"` // Key must match pattern: `^x-`.
}

// WithURL sets URL value.
func (s *Server) WithURL(val string) *Server {
	s.URL = val
	return s
}

// WithDescription sets Description value.
func (s *Server) WithDescription(val string) *Server {
	s.Description = &val
	return s
}

// WithVariables sets Variables value.
func (s *Server) WithVariables(val map[string]ServerVariable) *Server {
	s.Variables = val
	return s
}

// WithVariablesItem sets Variables item value.
func (s *Server) WithVariablesItem(key string, val ServerVariable) *Server {
	if s.Variables == nil {
		s.Variables = make(map[string]ServerVariable, 1)
	}

	s.Variables[key] = val

	return s
}

// WithMapOfAnything sets MapOfAnything value.
func (s *Server) WithMapOfAnything(val map[string]interface{}) *Server {
	s.MapOfAnything = val
	return s
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (s *Server) WithMapOfAnythingItem(key string, val interface{}) *Server {
	if s.MapOfAnything == nil {
		s.MapOfAnything = make(map[string]interface{}, 1)
	}

	s.MapOfAnything[key] = val

	return s
}

type marshalServer Server

var knownKeysServer = []string{
	"url",
	"description",
	"variables",
}

var requireKeysServer = []string{
	"url",
}

// UnmarshalJSON decodes JSON.
func (s *Server) UnmarshalJSON(data []byte) error {
	var err error

	ms := marshalServer(*s)

	err = json.Unmarshal(data, &ms)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysServer {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysServer {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ms.MapOfAnything == nil {
				ms.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ms.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Server: %v", offendingKeys)
	}

	*s = Server(ms)

	return nil
}

// MarshalJSON encodes JSON.
func (s Server) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalServer(s), s.MapOfAnything)
}

// ServerVariable structure is generated from "#/definitions/ServerVariable".
type ServerVariable struct {
	Enum          []string               `json:"enum,omitempty"`
	Default       string                 `json:"default"` // Required.
	Description   *string                `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"` // Key must match pattern: `^x-`.
}

// WithEnum sets Enum value.
func (s *ServerVariable) WithEnum(val ...string) *ServerVariable {
	s.Enum = val
	return s
}

// WithDefault sets Default value.
func (s *ServerVariable) WithDefault(val string) *ServerVariable {
	s.Default = val
	return s
}

// WithDescription sets Description value.
func (s *ServerVariable) WithDescription(val string) *ServerVariable {
	s.Description = &val
	return s
}

// WithMapOfAnything sets MapOfAnything value.
func (s *ServerVariable) WithMapOfAnything(val map[string]interface{}) *ServerVariable {
	s.MapOfAnything = val
	return s
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (s *ServerVariable) WithMapOfAnythingItem(key string, val interface{}) *ServerVariable {
	if s.MapOfAnything == nil {
		s.MapOfAnything = make(map[string]interface{}, 1)
	}

	s.MapOfAnything[key] = val

	return s
}

type marshalServerVariable ServerVariable

var knownKeysServerVariable = []string{
	"enum",
	"default",
	"description",
}

var requireKeysServerVariable = []string{
	"default",
}

// UnmarshalJSON decodes JSON.
func (s *ServerVariable) UnmarshalJSON(data []byte) error {
	var err error

	ms := marshalServerVariable(*s)

	err = json.Unmarshal(data, &ms)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysServerVariable {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysServerVariable {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ms.MapOfAnything == nil {
				ms.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ms.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in ServerVariable: %v", offendingKeys)
	}

	*s = ServerVariable(ms)

	return nil
}

// MarshalJSON encodes JSON.
func (s ServerVariable) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalServerVariable(s), s.MapOfAnything)
}

// Tag structure is generated from "#/definitions/Tag".
type Tag struct {
	Name          string                 `json:"name"` // Required.
	Description   *string                `json:"description,omitempty"`
	ExternalDocs  *ExternalDocumentation `json:"externalDocs,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"` // Key must match pattern: `^x-`.
}

// WithName sets Name value.
func (t *Tag) WithName(val string) *Tag {
	t.Name = val
	return t
}

// WithDescription sets Description value.
func (t *Tag) WithDescription(val string) *Tag {
	t.Description = &val
	return t
}

// WithExternalDocs sets ExternalDocs value.
func (t *Tag) WithExternalDocs(val ExternalDocumentation) *Tag {
	t.ExternalDocs = &val
	return t
}

// ExternalDocsEns ensures returned ExternalDocs is not nil.
func (t *Tag) ExternalDocsEns() *ExternalDocumentation {
	if t.ExternalDocs == nil {
		t.ExternalDocs = new(ExternalDocumentation)
	}

	return t.ExternalDocs
}

// WithMapOfAnything sets MapOfAnything value.
func (t *Tag) WithMapOfAnything(val map[string]interface{}) *Tag {
	t.MapOfAnything = val
	return t
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (t *Tag) WithMapOfAnythingItem(key string, val interface{}) *Tag {
	if t.MapOfAnything == nil {
		t.MapOfAnything = make(map[string]interface{}, 1)
	}

	t.MapOfAnything[key] = val

	return t
}

type marshalTag Tag

var knownKeysTag = []string{
	"name",
	"description",
	"externalDocs",
}

var requireKeysTag = []string{
	"name",
}

// UnmarshalJSON decodes JSON.
func (t *Tag) UnmarshalJSON(data []byte) error {
	var err error

	mt := marshalTag(*t)

	err = json.Unmarshal(data, &mt)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysTag {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysTag {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if mt.MapOfAnything == nil {
				mt.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mt.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Tag: %v", offendingKeys)
	}

	*t = Tag(mt)

	return nil
}

// MarshalJSON encodes JSON.
func (t Tag) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalTag(t), t.MapOfAnything)
}

// PathItem structure is generated from "#/definitions/PathItem".
type PathItem struct {
	Ref                  *string                `json:"$ref,omitempty"`
	Summary              *string                `json:"summary,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Servers              []Server               `json:"servers,omitempty"`
	Parameters           []ParameterOrRef       `json:"parameters,omitempty"`
	MapOfOperationValues map[string]Operation   `json:"-"` // Key must match pattern: `^(get|put|post|delete|options|head|patch|trace)$`.
	MapOfAnything        map[string]interface{} `json:"-"` // Key must match pattern: `^x-`.
}

// WithRef sets Ref value.
func (p *PathItem) WithRef(val string) *PathItem {
	p.Ref = &val
	return p
}

// WithSummary sets Summary value.
func (p *PathItem) WithSummary(val string) *PathItem {
	p.Summary = &val
	return p
}

// WithDescription sets Description value.
func (p *PathItem) WithDescription(val string) *PathItem {
	p.Description = &val
	return p
}

// WithServers sets Servers value.
func (p *PathItem) WithServers(val ...Server) *PathItem {
	p.Servers = val
	return p
}

// WithParameters sets Parameters value.
func (p *PathItem) WithParameters(val ...ParameterOrRef) *PathItem {
	p.Parameters = val
	return p
}

// WithMapOfOperationValues sets MapOfOperationValues value.
func (p *PathItem) WithMapOfOperationValues(val map[string]Operation) *PathItem {
	p.MapOfOperationValues = val
	return p
}

// WithMapOfOperationValuesItem sets MapOfOperationValues item value.
func (p *PathItem) WithMapOfOperationValuesItem(key string, val Operation) *PathItem {
	if p.MapOfOperationValues == nil {
		p.MapOfOperationValues = make(map[string]Operation, 1)
	}

	p.MapOfOperationValues[key] = val

	return p
}

// WithMapOfAnything sets MapOfAnything value.
func (p *PathItem) WithMapOfAnything(val map[string]interface{}) *PathItem {
	p.MapOfAnything = val
	return p
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (p *PathItem) WithMapOfAnythingItem(key string, val interface{}) *PathItem {
	if p.MapOfAnything == nil {
		p.MapOfAnything = make(map[string]interface{}, 1)
	}

	p.MapOfAnything[key] = val

	return p
}

type marshalPathItem PathItem

var knownKeysPathItem = []string{
	"$ref",
	"summary",
	"description",
	"servers",
	"parameters",
}

// UnmarshalJSON decodes JSON.
func (p *PathItem) UnmarshalJSON(data []byte) error {
	var err error

	mp := marshalPathItem(*p)

	err = json.Unmarshal(data, &mp)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range knownKeysPathItem {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexGetPutPostDeleteOptionsHeadPatchTrace.MatchString(key) {
			matched = true

			if mp.MapOfOperationValues == nil {
				mp.MapOfOperationValues = make(map[string]Operation, 1)
			}

			var val Operation

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mp.MapOfOperationValues[key] = val
		}

		if regexX.MatchString(key) {
			matched = true

			if mp.MapOfAnything == nil {
				mp.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mp.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in PathItem: %v", offendingKeys)
	}

	*p = PathItem(mp)

	return nil
}

// MarshalJSON encodes JSON.
func (p PathItem) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalPathItem(p), p.MapOfOperationValues, p.MapOfAnything)
}

// ParameterReference structure is generated from "#/definitions/ParameterReference".
type ParameterReference struct {
	// Format: uri-reference.
	// Required.
	Ref string `json:"$ref"`
}

// WithRef sets Ref value.
func (p *ParameterReference) WithRef(val string) *ParameterReference {
	p.Ref = val
	return p
}

type marshalParameterReference ParameterReference

var knownKeysParameterReference = []string{
	"$ref",
}

var requireKeysParameterReference = []string{
	"$ref",
}

// UnmarshalJSON decodes JSON.
func (p *ParameterReference) UnmarshalJSON(data []byte) error {
	var err error

	mp := marshalParameterReference(*p)

	err = json.Unmarshal(data, &mp)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysParameterReference {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysParameterReference {
		delete(rawMap, key)
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in ParameterReference: %v", offendingKeys)
	}

	*p = ParameterReference(mp)

	return nil
}

// Parameter structure is generated from "#/definitions/Parameter".
type Parameter struct {
	Name             string                  `json:"name"` // Required.
	In               ParameterIn             `json:"in"`   // Required.
	Description      *string                 `json:"description,omitempty"`
	Required         *bool                   `json:"required,omitempty"`
	Deprecated       *bool                   `json:"deprecated,omitempty"`
	AllowEmptyValue  *bool                   `json:"allowEmptyValue,omitempty"`
	Style            *string                 `json:"style,omitempty"`
	Explode          *bool                   `json:"explode,omitempty"`
	AllowReserved    *bool                   `json:"allowReserved,omitempty"`
	Schema           *SchemaOrRef            `json:"schema,omitempty"`
	Content          map[string]MediaType    `json:"content,omitempty"`
	Example          *interface{}            `json:"example,omitempty"`
	Examples         map[string]ExampleOrRef `json:"examples,omitempty"`
	SchemaXORContent *SchemaXORContent       `json:"-"`
	Location         *ParameterLocation      `json:"-"`
	MapOfAnything    map[string]interface{}  `json:"-"` // Key must match pattern: `^x-`.
}

// WithName sets Name value.
func (p *Parameter) WithName(val string) *Parameter {
	p.Name = val
	return p
}

// WithIn sets In value.
func (p *Parameter) WithIn(val ParameterIn) *Parameter {
	p.In = val
	return p
}

// WithDescription sets Description value.
func (p *Parameter) WithDescription(val string) *Parameter {
	p.Description = &val
	return p
}

// WithRequired sets Required value.
func (p *Parameter) WithRequired(val bool) *Parameter {
	p.Required = &val
	return p
}

// WithDeprecated sets Deprecated value.
func (p *Parameter) WithDeprecated(val bool) *Parameter {
	p.Deprecated = &val
	return p
}

// WithAllowEmptyValue sets AllowEmptyValue value.
func (p *Parameter) WithAllowEmptyValue(val bool) *Parameter {
	p.AllowEmptyValue = &val
	return p
}

// WithStyle sets Style value.
func (p *Parameter) WithStyle(val string) *Parameter {
	p.Style = &val
	return p
}

// WithExplode sets Explode value.
func (p *Parameter) WithExplode(val bool) *Parameter {
	p.Explode = &val
	return p
}

// WithAllowReserved sets AllowReserved value.
func (p *Parameter) WithAllowReserved(val bool) *Parameter {
	p.AllowReserved = &val
	return p
}

// WithSchema sets Schema value.
func (p *Parameter) WithSchema(val SchemaOrRef) *Parameter {
	p.Schema = &val
	return p
}

// SchemaEns ensures returned Schema is not nil.
func (p *Parameter) SchemaEns() *SchemaOrRef {
	if p.Schema == nil {
		p.Schema = new(SchemaOrRef)
	}

	return p.Schema
}

// WithContent sets Content value.
func (p *Parameter) WithContent(val map[string]MediaType) *Parameter {
	p.Content = val
	return p
}

// WithContentItem sets Content item value.
func (p *Parameter) WithContentItem(key string, val MediaType) *Parameter {
	if p.Content == nil {
		p.Content = make(map[string]MediaType, 1)
	}

	p.Content[key] = val

	return p
}

// WithExample sets Example value.
func (p *Parameter) WithExample(val interface{}) *Parameter {
	p.Example = &val
	return p
}

// WithExamples sets Examples value.
func (p *Parameter) WithExamples(val map[string]ExampleOrRef) *Parameter {
	p.Examples = val
	return p
}

// WithExamplesItem sets Examples item value.
func (p *Parameter) WithExamplesItem(key string, val ExampleOrRef) *Parameter {
	if p.Examples == nil {
		p.Examples = make(map[string]ExampleOrRef, 1)
	}

	p.Examples[key] = val

	return p
}

// WithSchemaXORContent sets SchemaXORContent value.
func (p *Parameter) WithSchemaXORContent(val SchemaXORContent) *Parameter {
	p.SchemaXORContent = &val
	return p
}

// SchemaXORContentEns ensures returned SchemaXORContent is not nil.
func (p *Parameter) SchemaXORContentEns() *SchemaXORContent {
	if p.SchemaXORContent == nil {
		p.SchemaXORContent = new(SchemaXORContent)
	}

	return p.SchemaXORContent
}

// WithLocation sets Location value.
func (p *Parameter) WithLocation(val ParameterLocation) *Parameter {
	p.Location = &val
	return p
}

// LocationEns ensures returned Location is not nil.
func (p *Parameter) LocationEns() *ParameterLocation {
	if p.Location == nil {
		p.Location = new(ParameterLocation)
	}

	return p.Location
}

// WithMapOfAnything sets MapOfAnything value.
func (p *Parameter) WithMapOfAnything(val map[string]interface{}) *Parameter {
	p.MapOfAnything = val
	return p
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (p *Parameter) WithMapOfAnythingItem(key string, val interface{}) *Parameter {
	if p.MapOfAnything == nil {
		p.MapOfAnything = make(map[string]interface{}, 1)
	}

	p.MapOfAnything[key] = val

	return p
}

type marshalParameter Parameter

var knownKeysParameter = []string{
	"name",
	"in",
	"description",
	"required",
	"deprecated",
	"allowEmptyValue",
	"style",
	"explode",
	"allowReserved",
	"schema",
	"content",
	"example",
	"examples",
}

var requireKeysParameter = []string{
	"name",
	"in",
}

// UnmarshalJSON decodes JSON.
func (p *Parameter) UnmarshalJSON(data []byte) error {
	var err error

	mp := marshalParameter(*p)

	err = json.Unmarshal(data, &mp)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &mp.SchemaXORContent)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &mp.Location)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysParameter {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	if mp.Example == nil {
		if _, ok := rawMap["example"]; ok {
			var v interface{}
			mp.Example = &v
		}
	}

	for _, key := range knownKeysParameter {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if mp.MapOfAnything == nil {
				mp.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mp.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Parameter: %v", offendingKeys)
	}

	*p = Parameter(mp)

	return nil
}

// MarshalJSON encodes JSON.
func (p Parameter) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalParameter(p), p.MapOfAnything, p.SchemaXORContent, p.Location)
}

// Schema structure is generated from "#/definitions/Schema".
type Schema struct {
	Title            *string       `json:"title,omitempty"`
	MultipleOf       *float64      `json:"multipleOf,omitempty"`
	Maximum          *float64      `json:"maximum,omitempty"`
	ExclusiveMaximum *bool         `json:"exclusiveMaximum,omitempty"`
	Minimum          *float64      `json:"minimum,omitempty"`
	ExclusiveMinimum *bool         `json:"exclusiveMinimum,omitempty"`
	MaxLength        *int64        `json:"maxLength,omitempty"`
	MinLength        *int64        `json:"minLength,omitempty"`
	Pattern          *string       `json:"pattern,omitempty"` // Format: regex.
	MaxItems         *int64        `json:"maxItems,omitempty"`
	MinItems         *int64        `json:"minItems,omitempty"`
	UniqueItems      *bool         `json:"uniqueItems,omitempty"`
	MaxProperties    *int64        `json:"maxProperties,omitempty"`
	MinProperties    *int64        `json:"minProperties,omitempty"`
	Required         []string      `json:"required,omitempty"`
	Enum             []interface{} `json:"enum,omitempty"`
	Type             *SchemaType   `json:"type,omitempty"`
	Not              *SchemaOrRef  `json:"not,omitempty"`
	AllOf            []SchemaOrRef `json:"allOf,omitempty"`
	OneOf            []SchemaOrRef `json:"oneOf,omitempty"`
	AnyOf            []SchemaOrRef `json:"anyOf,omitempty"`
	Items            *SchemaOrRef  `json:"items,omitempty"`
	//Properties           map[string]SchemaOrRef      `json:"properties,omitempty"`
	Properties           *orderedmap.OrderedMap[string, SchemaOrRef] `json:"properties,omitempty"`
	AdditionalProperties *SchemaAdditionalProperties                 `json:"additionalProperties,omitempty"`
	Description          *string                                     `json:"description,omitempty"`
	Format               *string                                     `json:"format,omitempty"`
	Default              *interface{}                                `json:"default,omitempty"`
	Nullable             *bool                                       `json:"nullable,omitempty"`
	Discriminator        *Discriminator                              `json:"discriminator,omitempty"`
	ReadOnly             *bool                                       `json:"readOnly,omitempty"`
	WriteOnly            *bool                                       `json:"writeOnly,omitempty"`
	Example              *interface{}                                `json:"example,omitempty"`
	ExternalDocs         *ExternalDocumentation                      `json:"externalDocs,omitempty"`
	Deprecated           *bool                                       `json:"deprecated,omitempty"`
	XML                  *XML                                        `json:"xml,omitempty"`
	MapOfAnything        map[string]interface{}                      `json:"-"` // Key must match pattern: `^x-`.
	ReflectType          reflect.Type                                `json:"-"`
}

// WithTitle sets Title value.
func (s *Schema) WithTitle(val string) *Schema {
	s.Title = &val
	return s
}

// WithMultipleOf sets MultipleOf value.
func (s *Schema) WithMultipleOf(val float64) *Schema {
	s.MultipleOf = &val
	return s
}

// WithMaximum sets Maximum value.
func (s *Schema) WithMaximum(val float64) *Schema {
	s.Maximum = &val
	return s
}

// WithExclusiveMaximum sets ExclusiveMaximum value.
func (s *Schema) WithExclusiveMaximum(val bool) *Schema {
	s.ExclusiveMaximum = &val
	return s
}

// WithMinimum sets Minimum value.
func (s *Schema) WithMinimum(val float64) *Schema {
	s.Minimum = &val
	return s
}

// WithExclusiveMinimum sets ExclusiveMinimum value.
func (s *Schema) WithExclusiveMinimum(val bool) *Schema {
	s.ExclusiveMinimum = &val
	return s
}

// WithMaxLength sets MaxLength value.
func (s *Schema) WithMaxLength(val int64) *Schema {
	s.MaxLength = &val
	return s
}

// WithMinLength sets MinLength value.
func (s *Schema) WithMinLength(val int64) *Schema {
	s.MinLength = &val
	return s
}

// WithPattern sets Pattern value.
func (s *Schema) WithPattern(val string) *Schema {
	s.Pattern = &val
	return s
}

// WithMaxItems sets MaxItems value.
func (s *Schema) WithMaxItems(val int64) *Schema {
	s.MaxItems = &val
	return s
}

// WithMinItems sets MinItems value.
func (s *Schema) WithMinItems(val int64) *Schema {
	s.MinItems = &val
	return s
}

// WithUniqueItems sets UniqueItems value.
func (s *Schema) WithUniqueItems(val bool) *Schema {
	s.UniqueItems = &val
	return s
}

// WithMaxProperties sets MaxProperties value.
func (s *Schema) WithMaxProperties(val int64) *Schema {
	s.MaxProperties = &val
	return s
}

// WithMinProperties sets MinProperties value.
func (s *Schema) WithMinProperties(val int64) *Schema {
	s.MinProperties = &val
	return s
}

// WithRequired sets Required value.
func (s *Schema) WithRequired(val ...string) *Schema {
	s.Required = val
	return s
}

// WithEnum sets Enum value.
func (s *Schema) WithEnum(val ...interface{}) *Schema {
	s.Enum = val
	return s
}

// WithType sets Type value.
func (s *Schema) WithType(val SchemaType) *Schema {
	s.Type = &val
	return s
}

// WithNot sets Not value.
func (s *Schema) WithNot(val SchemaOrRef) *Schema {
	s.Not = &val
	return s
}

// NotEns ensures returned Not is not nil.
func (s *Schema) NotEns() *SchemaOrRef {
	if s.Not == nil {
		s.Not = new(SchemaOrRef)
	}

	return s.Not
}

// WithAllOf sets AllOf value.
func (s *Schema) WithAllOf(val ...SchemaOrRef) *Schema {
	s.AllOf = val
	return s
}

// WithOneOf sets OneOf value.
func (s *Schema) WithOneOf(val ...SchemaOrRef) *Schema {
	s.OneOf = val
	return s
}

// WithAnyOf sets AnyOf value.
func (s *Schema) WithAnyOf(val ...SchemaOrRef) *Schema {
	s.AnyOf = val
	return s
}

// WithItems sets Items value.
func (s *Schema) WithItems(val SchemaOrRef) *Schema {
	s.Items = &val
	return s
}

// ItemsEns ensures returned Items is not nil.
func (s *Schema) ItemsEns() *SchemaOrRef {
	if s.Items == nil {
		s.Items = new(SchemaOrRef)
	}

	return s.Items
}

// WithProperties sets Properties value.
func (s *Schema) WithProperties(val map[string]SchemaOrRef) *Schema {
	if s.Properties == nil {
		s.Properties = orderedmap.New[string, SchemaOrRef]()
	}
	for k, v := range val {
		s.Properties.Set(k, v)
	}
	return s
}

// WithPropertiesItem sets Properties item value.
func (s *Schema) WithPropertiesItem(key string, val SchemaOrRef) *Schema {
	if s.Properties == nil {
		s.Properties = orderedmap.New[string, SchemaOrRef]()
	}
	s.Properties.Set(key, val)
	return s
}

// WithAdditionalProperties sets AdditionalProperties value.
func (s *Schema) WithAdditionalProperties(val SchemaAdditionalProperties) *Schema {
	s.AdditionalProperties = &val
	return s
}

// AdditionalPropertiesEns ensures returned AdditionalProperties is not nil.
func (s *Schema) AdditionalPropertiesEns() *SchemaAdditionalProperties {
	if s.AdditionalProperties == nil {
		s.AdditionalProperties = new(SchemaAdditionalProperties)
	}

	return s.AdditionalProperties
}

// WithDescription sets Description value.
func (s *Schema) WithDescription(val string) *Schema {
	s.Description = &val
	return s
}

// WithFormat sets Format value.
func (s *Schema) WithFormat(val string) *Schema {
	s.Format = &val
	return s
}

// WithDefault sets Default value.
func (s *Schema) WithDefault(val interface{}) *Schema {
	s.Default = &val
	return s
}

// WithNullable sets Nullable value.
func (s *Schema) WithNullable(val bool) *Schema {
	s.Nullable = &val
	return s
}

// WithDiscriminator sets Discriminator value.
func (s *Schema) WithDiscriminator(val Discriminator) *Schema {
	s.Discriminator = &val
	return s
}

// DiscriminatorEns ensures returned Discriminator is not nil.
func (s *Schema) DiscriminatorEns() *Discriminator {
	if s.Discriminator == nil {
		s.Discriminator = new(Discriminator)
	}

	return s.Discriminator
}

// WithReadOnly sets ReadOnly value.
func (s *Schema) WithReadOnly(val bool) *Schema {
	s.ReadOnly = &val
	return s
}

// WithWriteOnly sets WriteOnly value.
func (s *Schema) WithWriteOnly(val bool) *Schema {
	s.WriteOnly = &val
	return s
}

// WithExample sets Example value.
func (s *Schema) WithExample(val interface{}) *Schema {
	s.Example = &val
	return s
}

// WithExternalDocs sets ExternalDocs value.
func (s *Schema) WithExternalDocs(val ExternalDocumentation) *Schema {
	s.ExternalDocs = &val
	return s
}

// ExternalDocsEns ensures returned ExternalDocs is not nil.
func (s *Schema) ExternalDocsEns() *ExternalDocumentation {
	if s.ExternalDocs == nil {
		s.ExternalDocs = new(ExternalDocumentation)
	}

	return s.ExternalDocs
}

// WithDeprecated sets Deprecated value.
func (s *Schema) WithDeprecated(val bool) *Schema {
	s.Deprecated = &val
	return s
}

// WithXML sets XML value.
func (s *Schema) WithXML(val XML) *Schema {
	s.XML = &val
	return s
}

// XMLEns ensures returned XML is not nil.
func (s *Schema) XMLEns() *XML {
	if s.XML == nil {
		s.XML = new(XML)
	}

	return s.XML
}

// WithMapOfAnything sets MapOfAnything value.
func (s *Schema) WithMapOfAnything(val map[string]interface{}) *Schema {
	s.MapOfAnything = val
	return s
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (s *Schema) WithMapOfAnythingItem(key string, val interface{}) *Schema {
	if s.MapOfAnything == nil {
		s.MapOfAnything = make(map[string]interface{}, 1)
	}

	s.MapOfAnything[key] = val

	return s
}

type marshalSchema Schema

var knownKeysSchema = []string{
	"title",
	"multipleOf",
	"maximum",
	"exclusiveMaximum",
	"minimum",
	"exclusiveMinimum",
	"maxLength",
	"minLength",
	"pattern",
	"maxItems",
	"minItems",
	"uniqueItems",
	"maxProperties",
	"minProperties",
	"required",
	"enum",
	"type",
	"not",
	"allOf",
	"oneOf",
	"anyOf",
	"items",
	"properties",
	"additionalProperties",
	"description",
	"format",
	"default",
	"nullable",
	"discriminator",
	"readOnly",
	"writeOnly",
	"example",
	"externalDocs",
	"deprecated",
	"xml",
}

// UnmarshalJSON decodes JSON.
func (s *Schema) UnmarshalJSON(data []byte) error {
	var err error

	ms := marshalSchema(*s)

	err = json.Unmarshal(data, &ms)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if ms.Default == nil {
		if _, ok := rawMap["default"]; ok {
			var v interface{}
			ms.Default = &v
		}
	}

	if ms.Example == nil {
		if _, ok := rawMap["example"]; ok {
			var v interface{}
			ms.Example = &v
		}
	}

	for _, key := range knownKeysSchema {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ms.MapOfAnything == nil {
				ms.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ms.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Schema: %v", offendingKeys)
	}

	*s = Schema(ms)

	return nil
}

// MarshalJSON encodes JSON.
func (s Schema) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalSchema(s), s.MapOfAnything)
}

// SchemaReference structure is generated from "#/definitions/SchemaReference".
type SchemaReference struct {
	// Format: uri-reference.
	// Required.
	Ref string `json:"$ref"`
}

// WithRef sets Ref value.
func (s *SchemaReference) WithRef(val string) *SchemaReference {
	s.Ref = val
	return s
}

type marshalSchemaReference SchemaReference

var knownKeysSchemaReference = []string{
	"$ref",
}

var requireKeysSchemaReference = []string{
	"$ref",
}

// UnmarshalJSON decodes JSON.
func (s *SchemaReference) UnmarshalJSON(data []byte) error {
	var err error

	ms := marshalSchemaReference(*s)

	err = json.Unmarshal(data, &ms)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysSchemaReference {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysSchemaReference {
		delete(rawMap, key)
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in SchemaReference: %v", offendingKeys)
	}

	*s = SchemaReference(ms)

	return nil
}

// SchemaOrRef structure is generated from "#/definitions/SchemaOrRef".
type SchemaOrRef struct {
	Schema          *Schema          `json:"-"`
	SchemaReference *SchemaReference `json:"-"`
}

// WithSchema sets Schema value.
func (s *SchemaOrRef) WithSchema(val Schema) *SchemaOrRef {
	s.Schema = &val
	return s
}

// SchemaEns ensures returned Schema is not nil.
func (s *SchemaOrRef) SchemaEns() *Schema {
	if s.Schema == nil {
		s.Schema = new(Schema)
	}

	return s.Schema
}

// WithSchemaReference sets SchemaReference value.
func (s *SchemaOrRef) WithSchemaReference(val SchemaReference) *SchemaOrRef {
	s.SchemaReference = &val
	return s
}

// SchemaReferenceEns ensures returned SchemaReference is not nil.
func (s *SchemaOrRef) SchemaReferenceEns() *SchemaReference {
	if s.SchemaReference == nil {
		s.SchemaReference = new(SchemaReference)
	}

	return s.SchemaReference
}

// UnmarshalJSON decodes JSON.
func (s *SchemaOrRef) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 2)
	oneOfValid := 0

	err = json.Unmarshal(data, &s.Schema)
	if err != nil {
		oneOfErrors["Schema"] = err
		s.Schema = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &s.SchemaReference)
	if err != nil {
		oneOfErrors["SchemaReference"] = err
		s.SchemaReference = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for SchemaOrRef with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (s SchemaOrRef) MarshalJSON() ([]byte, error) {
	return marshalUnion(s.Schema, s.SchemaReference)
}

// SchemaAdditionalProperties structure is generated from "#/definitions/Schema->additionalProperties".
type SchemaAdditionalProperties struct {
	SchemaOrRef *SchemaOrRef `json:"-"`
	Bool        *bool        `json:"-"`
}

// WithSchemaOrRef sets SchemaOrRef value.
func (s *SchemaAdditionalProperties) WithSchemaOrRef(val SchemaOrRef) *SchemaAdditionalProperties {
	s.SchemaOrRef = &val
	return s
}

// SchemaOrRefEns ensures returned SchemaOrRef is not nil.
func (s *SchemaAdditionalProperties) SchemaOrRefEns() *SchemaOrRef {
	if s.SchemaOrRef == nil {
		s.SchemaOrRef = new(SchemaOrRef)
	}

	return s.SchemaOrRef
}

// WithBool sets Bool value.
func (s *SchemaAdditionalProperties) WithBool(val bool) *SchemaAdditionalProperties {
	s.Bool = &val
	return s
}

// UnmarshalJSON decodes JSON.
func (s *SchemaAdditionalProperties) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 2)
	oneOfValid := 0

	err = json.Unmarshal(data, &s.SchemaOrRef)
	if err != nil {
		oneOfErrors["SchemaOrRef"] = err
		s.SchemaOrRef = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &s.Bool)
	if err != nil {
		oneOfErrors["Bool"] = err
		s.Bool = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for SchemaAdditionalProperties with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (s SchemaAdditionalProperties) MarshalJSON() ([]byte, error) {
	return marshalUnion(s.SchemaOrRef, s.Bool)
}

// Discriminator structure is generated from "#/definitions/Discriminator".
type Discriminator struct {
	PropertyName string            `json:"propertyName"` // Required.
	Mapping      map[string]string `json:"mapping,omitempty"`
}

// WithPropertyName sets PropertyName value.
func (d *Discriminator) WithPropertyName(val string) *Discriminator {
	d.PropertyName = val
	return d
}

// WithMapping sets Mapping value.
func (d *Discriminator) WithMapping(val map[string]string) *Discriminator {
	d.Mapping = val
	return d
}

// WithMappingItem sets Mapping item value.
func (d *Discriminator) WithMappingItem(key string, val string) *Discriminator {
	if d.Mapping == nil {
		d.Mapping = make(map[string]string, 1)
	}

	d.Mapping[key] = val

	return d
}

type marshalDiscriminator Discriminator

var requireKeysDiscriminator = []string{
	"propertyName",
}

// UnmarshalJSON decodes JSON.
func (d *Discriminator) UnmarshalJSON(data []byte) error {
	var err error

	md := marshalDiscriminator(*d)

	err = json.Unmarshal(data, &md)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysDiscriminator {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	*d = Discriminator(md)

	return nil
}

// XML structure is generated from "#/definitions/XML".
type XML struct {
	Name          *string                `json:"name,omitempty"`
	Namespace     *string                `json:"namespace,omitempty"` // Format: uri.
	Prefix        *string                `json:"prefix,omitempty"`
	Attribute     *bool                  `json:"attribute,omitempty"`
	Wrapped       *bool                  `json:"wrapped,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"` // Key must match pattern: `^x-`.
}

// WithName sets Name value.
func (x *XML) WithName(val string) *XML {
	x.Name = &val
	return x
}

// WithNamespace sets Namespace value.
func (x *XML) WithNamespace(val string) *XML {
	x.Namespace = &val
	return x
}

// WithPrefix sets Prefix value.
func (x *XML) WithPrefix(val string) *XML {
	x.Prefix = &val
	return x
}

// WithAttribute sets Attribute value.
func (x *XML) WithAttribute(val bool) *XML {
	x.Attribute = &val
	return x
}

// WithWrapped sets Wrapped value.
func (x *XML) WithWrapped(val bool) *XML {
	x.Wrapped = &val
	return x
}

// WithMapOfAnything sets MapOfAnything value.
func (x *XML) WithMapOfAnything(val map[string]interface{}) *XML {
	x.MapOfAnything = val
	return x
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (x *XML) WithMapOfAnythingItem(key string, val interface{}) *XML {
	if x.MapOfAnything == nil {
		x.MapOfAnything = make(map[string]interface{}, 1)
	}

	x.MapOfAnything[key] = val

	return x
}

type marshalXML XML

var knownKeysXML = []string{
	"name",
	"namespace",
	"prefix",
	"attribute",
	"wrapped",
}

// UnmarshalJSON decodes JSON.
func (x *XML) UnmarshalJSON(data []byte) error {
	var err error

	mx := marshalXML(*x)

	err = json.Unmarshal(data, &mx)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range knownKeysXML {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if mx.MapOfAnything == nil {
				mx.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mx.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in XML: %v", offendingKeys)
	}

	*x = XML(mx)

	return nil
}

// MarshalJSON encodes JSON.
func (x XML) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalXML(x), x.MapOfAnything)
}

// MediaType structure is generated from "#/definitions/MediaType".
type MediaType struct {
	Schema        *SchemaOrRef            `json:"schema,omitempty"`
	Example       *interface{}            `json:"example,omitempty"`
	Examples      map[string]ExampleOrRef `json:"examples,omitempty"`
	Encoding      map[string]Encoding     `json:"encoding,omitempty"`
	MapOfAnything map[string]interface{}  `json:"-"` // Key must match pattern: `^x-`.
}

// WithSchema sets Schema value.
func (m *MediaType) WithSchema(val SchemaOrRef) *MediaType {
	m.Schema = &val
	return m
}

// SchemaEns ensures returned Schema is not nil.
func (m *MediaType) SchemaEns() *SchemaOrRef {
	if m.Schema == nil {
		m.Schema = new(SchemaOrRef)
	}

	return m.Schema
}

// WithExample sets Example value.
func (m *MediaType) WithExample(val interface{}) *MediaType {
	m.Example = &val
	return m
}

// WithExamples sets Examples value.
func (m *MediaType) WithExamples(val map[string]ExampleOrRef) *MediaType {
	m.Examples = val
	return m
}

// WithExamplesItem sets Examples item value.
func (m *MediaType) WithExamplesItem(key string, val ExampleOrRef) *MediaType {
	if m.Examples == nil {
		m.Examples = make(map[string]ExampleOrRef, 1)
	}

	m.Examples[key] = val

	return m
}

// WithEncoding sets Encoding value.
func (m *MediaType) WithEncoding(val map[string]Encoding) *MediaType {
	m.Encoding = val
	return m
}

// WithEncodingItem sets Encoding item value.
func (m *MediaType) WithEncodingItem(key string, val Encoding) *MediaType {
	if m.Encoding == nil {
		m.Encoding = make(map[string]Encoding, 1)
	}

	m.Encoding[key] = val

	return m
}

// WithMapOfAnything sets MapOfAnything value.
func (m *MediaType) WithMapOfAnything(val map[string]interface{}) *MediaType {
	m.MapOfAnything = val
	return m
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (m *MediaType) WithMapOfAnythingItem(key string, val interface{}) *MediaType {
	if m.MapOfAnything == nil {
		m.MapOfAnything = make(map[string]interface{}, 1)
	}

	m.MapOfAnything[key] = val

	return m
}

type marshalMediaType MediaType

var knownKeysMediaType = []string{
	"schema",
	"example",
	"examples",
	"encoding",
}

// UnmarshalJSON decodes JSON.
func (m *MediaType) UnmarshalJSON(data []byte) error {
	var err error

	mm := marshalMediaType(*m)

	err = json.Unmarshal(data, &mm)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if mm.Example == nil {
		if _, ok := rawMap["example"]; ok {
			var v interface{}
			mm.Example = &v
		}
	}

	for _, key := range knownKeysMediaType {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if mm.MapOfAnything == nil {
				mm.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mm.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in MediaType: %v", offendingKeys)
	}

	*m = MediaType(mm)

	return nil
}

// MarshalJSON encodes JSON.
func (m MediaType) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalMediaType(m), m.MapOfAnything)
}

// ExampleReference structure is generated from "#/definitions/ExampleReference".
type ExampleReference struct {
	// Format: uri-reference.
	// Required.
	Ref string `json:"$ref"`
}

// WithRef sets Ref value.
func (e *ExampleReference) WithRef(val string) *ExampleReference {
	e.Ref = val
	return e
}

type marshalExampleReference ExampleReference

var knownKeysExampleReference = []string{
	"$ref",
}

var requireKeysExampleReference = []string{
	"$ref",
}

// UnmarshalJSON decodes JSON.
func (e *ExampleReference) UnmarshalJSON(data []byte) error {
	var err error

	me := marshalExampleReference(*e)

	err = json.Unmarshal(data, &me)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysExampleReference {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysExampleReference {
		delete(rawMap, key)
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in ExampleReference: %v", offendingKeys)
	}

	*e = ExampleReference(me)

	return nil
}

// Example structure is generated from "#/definitions/Example".
type Example struct {
	Summary       *string                `json:"summary,omitempty"`
	Description   *string                `json:"description,omitempty"`
	Value         *interface{}           `json:"value,omitempty"`
	ExternalValue *string                `json:"externalValue,omitempty"` // Format: uri-reference.
	MapOfAnything map[string]interface{} `json:"-"`                       // Key must match pattern: `^x-`.
}

// WithSummary sets Summary value.
func (e *Example) WithSummary(val string) *Example {
	e.Summary = &val
	return e
}

// WithDescription sets Description value.
func (e *Example) WithDescription(val string) *Example {
	e.Description = &val
	return e
}

// WithValue sets Value value.
func (e *Example) WithValue(val interface{}) *Example {
	e.Value = &val
	return e
}

// WithExternalValue sets ExternalValue value.
func (e *Example) WithExternalValue(val string) *Example {
	e.ExternalValue = &val
	return e
}

// WithMapOfAnything sets MapOfAnything value.
func (e *Example) WithMapOfAnything(val map[string]interface{}) *Example {
	e.MapOfAnything = val
	return e
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (e *Example) WithMapOfAnythingItem(key string, val interface{}) *Example {
	if e.MapOfAnything == nil {
		e.MapOfAnything = make(map[string]interface{}, 1)
	}

	e.MapOfAnything[key] = val

	return e
}

type marshalExample Example

var knownKeysExample = []string{
	"summary",
	"description",
	"value",
	"externalValue",
}

// UnmarshalJSON decodes JSON.
func (e *Example) UnmarshalJSON(data []byte) error {
	var err error

	me := marshalExample(*e)

	err = json.Unmarshal(data, &me)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if me.Value == nil {
		if _, ok := rawMap["value"]; ok {
			var v interface{}
			me.Value = &v
		}
	}

	for _, key := range knownKeysExample {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if me.MapOfAnything == nil {
				me.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			me.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Example: %v", offendingKeys)
	}

	*e = Example(me)

	return nil
}

// MarshalJSON encodes JSON.
func (e Example) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalExample(e), e.MapOfAnything)
}

// ExampleOrRef structure is generated from "#/definitions/ExampleOrRef".
type ExampleOrRef struct {
	ExampleReference *ExampleReference `json:"-"`
	Example          *Example          `json:"-"`
}

// WithExampleReference sets ExampleReference value.
func (e *ExampleOrRef) WithExampleReference(val ExampleReference) *ExampleOrRef {
	e.ExampleReference = &val
	return e
}

// ExampleReferenceEns ensures returned ExampleReference is not nil.
func (e *ExampleOrRef) ExampleReferenceEns() *ExampleReference {
	if e.ExampleReference == nil {
		e.ExampleReference = new(ExampleReference)
	}

	return e.ExampleReference
}

// WithExample sets Example value.
func (e *ExampleOrRef) WithExample(val Example) *ExampleOrRef {
	e.Example = &val
	return e
}

// ExampleEns ensures returned Example is not nil.
func (e *ExampleOrRef) ExampleEns() *Example {
	if e.Example == nil {
		e.Example = new(Example)
	}

	return e.Example
}

// UnmarshalJSON decodes JSON.
func (e *ExampleOrRef) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 2)
	oneOfValid := 0

	err = json.Unmarshal(data, &e.ExampleReference)
	if err != nil {
		oneOfErrors["ExampleReference"] = err
		e.ExampleReference = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &e.Example)
	if err != nil {
		oneOfErrors["Example"] = err
		e.Example = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for ExampleOrRef with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (e ExampleOrRef) MarshalJSON() ([]byte, error) {
	return marshalUnion(e.ExampleReference, e.Example)
}

// Encoding structure is generated from "#/definitions/Encoding".
type Encoding struct {
	ContentType   *string           `json:"contentType,omitempty"`
	Headers       map[string]Header `json:"headers,omitempty"`
	Style         *EncodingStyle    `json:"style,omitempty"`
	Explode       *bool             `json:"explode,omitempty"`
	AllowReserved *bool             `json:"allowReserved,omitempty"`
}

// WithContentType sets ContentType value.
func (e *Encoding) WithContentType(val string) *Encoding {
	e.ContentType = &val
	return e
}

// WithHeaders sets Headers value.
func (e *Encoding) WithHeaders(val map[string]Header) *Encoding {
	e.Headers = val
	return e
}

// WithHeadersItem sets Headers item value.
func (e *Encoding) WithHeadersItem(key string, val Header) *Encoding {
	if e.Headers == nil {
		e.Headers = make(map[string]Header, 1)
	}

	e.Headers[key] = val

	return e
}

// WithStyle sets Style value.
func (e *Encoding) WithStyle(val EncodingStyle) *Encoding {
	e.Style = &val
	return e
}

// WithExplode sets Explode value.
func (e *Encoding) WithExplode(val bool) *Encoding {
	e.Explode = &val
	return e
}

// WithAllowReserved sets AllowReserved value.
func (e *Encoding) WithAllowReserved(val bool) *Encoding {
	e.AllowReserved = &val
	return e
}

type marshalEncoding Encoding

var knownKeysEncoding = []string{
	"contentType",
	"headers",
	"style",
	"explode",
	"allowReserved",
}

// UnmarshalJSON decodes JSON.
func (e *Encoding) UnmarshalJSON(data []byte) error {
	var err error

	me := marshalEncoding(*e)

	err = json.Unmarshal(data, &me)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range knownKeysEncoding {
		delete(rawMap, key)
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Encoding: %v", offendingKeys)
	}

	*e = Encoding(me)

	return nil
}

// Header structure is generated from "#/definitions/Header".
type Header struct {
	Description     *string                 `json:"description,omitempty"`
	Required        *bool                   `json:"required,omitempty"`
	Deprecated      *bool                   `json:"deprecated,omitempty"`
	AllowEmptyValue *bool                   `json:"allowEmptyValue,omitempty"`
	Explode         *bool                   `json:"explode,omitempty"`
	AllowReserved   *bool                   `json:"allowReserved,omitempty"`
	Schema          *SchemaOrRef            `json:"schema,omitempty"`
	Content         map[string]MediaType    `json:"content,omitempty"`
	Example         *interface{}            `json:"example,omitempty"`
	Examples        map[string]ExampleOrRef `json:"examples,omitempty"`
	MapOfAnything   map[string]interface{}  `json:"-"` // Key must match pattern: `^x-`.
}

// WithDescription sets Description value.
func (h *Header) WithDescription(val string) *Header {
	h.Description = &val
	return h
}

// WithRequired sets Required value.
func (h *Header) WithRequired(val bool) *Header {
	h.Required = &val
	return h
}

// WithDeprecated sets Deprecated value.
func (h *Header) WithDeprecated(val bool) *Header {
	h.Deprecated = &val
	return h
}

// WithAllowEmptyValue sets AllowEmptyValue value.
func (h *Header) WithAllowEmptyValue(val bool) *Header {
	h.AllowEmptyValue = &val
	return h
}

// WithExplode sets Explode value.
func (h *Header) WithExplode(val bool) *Header {
	h.Explode = &val
	return h
}

// WithAllowReserved sets AllowReserved value.
func (h *Header) WithAllowReserved(val bool) *Header {
	h.AllowReserved = &val
	return h
}

// WithSchema sets Schema value.
func (h *Header) WithSchema(val SchemaOrRef) *Header {
	h.Schema = &val
	return h
}

// SchemaEns ensures returned Schema is not nil.
func (h *Header) SchemaEns() *SchemaOrRef {
	if h.Schema == nil {
		h.Schema = new(SchemaOrRef)
	}

	return h.Schema
}

// WithContent sets Content value.
func (h *Header) WithContent(val map[string]MediaType) *Header {
	h.Content = val
	return h
}

// WithContentItem sets Content item value.
func (h *Header) WithContentItem(key string, val MediaType) *Header {
	if h.Content == nil {
		h.Content = make(map[string]MediaType, 1)
	}

	h.Content[key] = val

	return h
}

// WithExample sets Example value.
func (h *Header) WithExample(val interface{}) *Header {
	h.Example = &val
	return h
}

// WithExamples sets Examples value.
func (h *Header) WithExamples(val map[string]ExampleOrRef) *Header {
	h.Examples = val
	return h
}

// WithExamplesItem sets Examples item value.
func (h *Header) WithExamplesItem(key string, val ExampleOrRef) *Header {
	if h.Examples == nil {
		h.Examples = make(map[string]ExampleOrRef, 1)
	}

	h.Examples[key] = val

	return h
}

// WithMapOfAnything sets MapOfAnything value.
func (h *Header) WithMapOfAnything(val map[string]interface{}) *Header {
	h.MapOfAnything = val
	return h
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (h *Header) WithMapOfAnythingItem(key string, val interface{}) *Header {
	if h.MapOfAnything == nil {
		h.MapOfAnything = make(map[string]interface{}, 1)
	}

	h.MapOfAnything[key] = val

	return h
}

type marshalHeader Header

var knownKeysHeader = []string{
	"description",
	"required",
	"deprecated",
	"allowEmptyValue",
	"explode",
	"allowReserved",
	"schema",
	"content",
	"example",
	"examples",
	"style",
}

// UnmarshalJSON decodes JSON.
func (h *Header) UnmarshalJSON(data []byte) error {
	var err error

	mh := marshalHeader(*h)

	err = json.Unmarshal(data, &mh)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, exists := rawMap["style"]; exists && string(v) != `"simple"` {
		return fmt.Errorf(`bad const value for "style" ("simple" expected, %s received)`, v)
	}

	delete(rawMap, "style")

	if mh.Example == nil {
		if _, ok := rawMap["example"]; ok {
			var v interface{}
			mh.Example = &v
		}
	}

	for _, key := range knownKeysHeader {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if mh.MapOfAnything == nil {
				mh.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mh.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Header: %v", offendingKeys)
	}

	*h = Header(mh)

	return nil
}

// constHeader is unconditionally added to JSON.
var constHeader = json.RawMessage(`{"style":"simple"}`)

// MarshalJSON encodes JSON.
func (h Header) MarshalJSON() ([]byte, error) {
	return marshalUnion(constHeader, marshalHeader(h), h.MapOfAnything)
}

// HasSchema structure is generated from "#/definitions/SchemaXORContent/oneOf/0".
//
// Has Schema.
type HasSchema struct {
	Schema interface{} `json:"schema"` // Required.
}

// WithSchema sets Schema value.
func (h *HasSchema) WithSchema(val interface{}) *HasSchema {
	h.Schema = val
	return h
}

type marshalHasSchema HasSchema

var requireKeysHasSchema = []string{
	"schema",
}

// UnmarshalJSON decodes JSON.
func (h *HasSchema) UnmarshalJSON(data []byte) error {
	var err error

	mh := marshalHasSchema(*h)

	err = json.Unmarshal(data, &mh)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysHasSchema {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	*h = HasSchema(mh)

	return nil
}

// HasContent structure is generated from "#/definitions/SchemaXORContent/oneOf/1".
//
// Has Content.
//
// Some properties are not allowed if content is present.
type HasContent struct {
	Content interface{} `json:"content"` // Required.
}

// WithContent sets Content value.
func (h *HasContent) WithContent(val interface{}) *HasContent {
	h.Content = val
	return h
}

type marshalHasContent HasContent

var requireKeysHasContent = []string{
	"content",
}

// UnmarshalJSON decodes JSON.
func (h *HasContent) UnmarshalJSON(data []byte) error {
	var err error

	mh := marshalHasContent(*h)

	err = json.Unmarshal(data, &mh)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysHasContent {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	*h = HasContent(mh)

	return nil
}

// SchemaXORContent structure is generated from "#/definitions/SchemaXORContent".
//
// Schema and content are mutually exclusive, at least one is required.
type SchemaXORContent struct {
	HasSchema  *HasSchema  `json:"-"`
	HasContent *HasContent `json:"-"`
}

// WithHasSchema sets HasSchema value.
func (s *SchemaXORContent) WithHasSchema(val HasSchema) *SchemaXORContent {
	s.HasSchema = &val
	return s
}

// HasSchemaEns ensures returned HasSchema is not nil.
func (s *SchemaXORContent) HasSchemaEns() *HasSchema {
	if s.HasSchema == nil {
		s.HasSchema = new(HasSchema)
	}

	return s.HasSchema
}

// WithHasContent sets HasContent value.
func (s *SchemaXORContent) WithHasContent(val HasContent) *SchemaXORContent {
	s.HasContent = &val
	return s
}

// HasContentEns ensures returned HasContent is not nil.
func (s *SchemaXORContent) HasContentEns() *HasContent {
	if s.HasContent == nil {
		s.HasContent = new(HasContent)
	}

	return s.HasContent
}

// UnmarshalJSON decodes JSON.
func (s *SchemaXORContent) UnmarshalJSON(data []byte) error {
	var err error

	var not SchemaXORContentNot

	if json.Unmarshal(data, &not) == nil {
		return errors.New("not constraint failed for SchemaXORContent")
	}

	oneOfErrors := make(map[string]error, 2)
	oneOfValid := 0

	err = json.Unmarshal(data, &s.HasSchema)
	if err != nil {
		oneOfErrors["HasSchema"] = err
		s.HasSchema = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &s.HasContent)
	if err != nil {
		oneOfErrors["HasContent"] = err
		s.HasContent = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for SchemaXORContent with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (s SchemaXORContent) MarshalJSON() ([]byte, error) {
	return marshalUnion(s.HasSchema, s.HasContent)
}

// SchemaXORContentNot structure is generated from "#/definitions/SchemaXORContent->not".
type SchemaXORContentNot struct {
	Schema  interface{} `json:"schema"`  // Required.
	Content interface{} `json:"content"` // Required.
}

// WithSchema sets Schema value.
func (s *SchemaXORContentNot) WithSchema(val interface{}) *SchemaXORContentNot {
	s.Schema = val
	return s
}

// WithContent sets Content value.
func (s *SchemaXORContentNot) WithContent(val interface{}) *SchemaXORContentNot {
	s.Content = val
	return s
}

type marshalSchemaXORContentNot SchemaXORContentNot

var requireKeysSchemaXORContentNot = []string{
	"schema",
	"content",
}

// UnmarshalJSON decodes JSON.
func (s *SchemaXORContentNot) UnmarshalJSON(data []byte) error {
	var err error

	ms := marshalSchemaXORContentNot(*s)

	err = json.Unmarshal(data, &ms)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysSchemaXORContentNot {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	*s = SchemaXORContentNot(ms)

	return nil
}

// PathParameter structure is generated from "#/definitions/ParameterLocation/oneOf/0".
//
// Path Parameter.
//
// Parameter in path.
type PathParameter struct {
	Style *PathParameterStyle `json:"style,omitempty"`
}

// WithStyle sets Style value.
func (p *PathParameter) WithStyle(val PathParameterStyle) *PathParameter {
	p.Style = &val
	return p
}

type marshalPathParameter PathParameter

var requireKeysPathParameter = []string{
	"required",
}

// UnmarshalJSON decodes JSON.
func (p *PathParameter) UnmarshalJSON(data []byte) error {
	var err error

	mp := marshalPathParameter(*p)

	err = json.Unmarshal(data, &mp)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysPathParameter {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	if v, exists := rawMap["in"]; exists && string(v) != `"path"` {
		return fmt.Errorf(`bad const value for "in" ("path" expected, %s received)`, v)
	}

	delete(rawMap, "in")

	if v, exists := rawMap["required"]; exists && string(v) != "true" {
		return fmt.Errorf(`bad const value for "required" (true expected, %s received)`, v)
	}

	delete(rawMap, "required")

	*p = PathParameter(mp)

	return nil
}

// constPathParameter is unconditionally added to JSON.
var constPathParameter = json.RawMessage(`{"in":"path","required":true}`)

// MarshalJSON encodes JSON.
func (p PathParameter) MarshalJSON() ([]byte, error) {
	return marshalUnion(constPathParameter, marshalPathParameter(p))
}

// QueryParameter structure is generated from "#/definitions/ParameterLocation/oneOf/1".
//
// Query Parameter.
//
// Parameter in query.
type QueryParameter struct {
	Style *QueryParameterStyle `json:"style,omitempty"`
}

// WithStyle sets Style value.
func (q *QueryParameter) WithStyle(val QueryParameterStyle) *QueryParameter {
	q.Style = &val
	return q
}

type marshalQueryParameter QueryParameter

// UnmarshalJSON decodes JSON.
func (q *QueryParameter) UnmarshalJSON(data []byte) error {
	var err error

	mq := marshalQueryParameter(*q)

	err = json.Unmarshal(data, &mq)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, exists := rawMap["in"]; exists && string(v) != `"query"` {
		return fmt.Errorf(`bad const value for "in" ("query" expected, %s received)`, v)
	}

	delete(rawMap, "in")

	*q = QueryParameter(mq)

	return nil
}

// constQueryParameter is unconditionally added to JSON.
var constQueryParameter = json.RawMessage(`{"in":"query"}`)

// MarshalJSON encodes JSON.
func (q QueryParameter) MarshalJSON() ([]byte, error) {
	return marshalUnion(constQueryParameter, marshalQueryParameter(q))
}

// HeaderParameter structure is generated from "#/definitions/ParameterLocation/oneOf/2".
//
// Header Parameter.
//
// Parameter in header.
type HeaderParameter struct{}

// UnmarshalJSON decodes JSON.
func (h *HeaderParameter) UnmarshalJSON(data []byte) error {
	var err error

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, exists := rawMap["in"]; exists && string(v) != `"header"` {
		return fmt.Errorf(`bad const value for "in" ("header" expected, %s received)`, v)
	}

	delete(rawMap, "in")

	if v, exists := rawMap["style"]; exists && string(v) != `"simple"` {
		return fmt.Errorf(`bad const value for "style" ("simple" expected, %s received)`, v)
	}

	delete(rawMap, "style")

	return nil
}

// constHeaderParameter is unconditionally added to JSON.
var constHeaderParameter = json.RawMessage(`{"in":"header","style":"simple"}`)

// MarshalJSON encodes JSON.
func (h HeaderParameter) MarshalJSON() ([]byte, error) {
	return marshalUnion(constHeaderParameter)
}

// CookieParameter structure is generated from "#/definitions/ParameterLocation/oneOf/3".
//
// Cookie Parameter.
//
// Parameter in cookie.
type CookieParameter struct{}

// UnmarshalJSON decodes JSON.
func (c *CookieParameter) UnmarshalJSON(data []byte) error {
	var err error

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, exists := rawMap["in"]; exists && string(v) != `"cookie"` {
		return fmt.Errorf(`bad const value for "in" ("cookie" expected, %s received)`, v)
	}

	delete(rawMap, "in")

	if v, exists := rawMap["style"]; exists && string(v) != `"form"` {
		return fmt.Errorf(`bad const value for "style" ("form" expected, %s received)`, v)
	}

	delete(rawMap, "style")

	return nil
}

// constCookieParameter is unconditionally added to JSON.
var constCookieParameter = json.RawMessage(`{"in":"cookie","style":"form"}`)

// MarshalJSON encodes JSON.
func (c CookieParameter) MarshalJSON() ([]byte, error) {
	return marshalUnion(constCookieParameter)
}

// ParameterLocation structure is generated from "#/definitions/ParameterLocation".
//
// Parameter location.
type ParameterLocation struct {
	PathParameter   *PathParameter   `json:"-"`
	QueryParameter  *QueryParameter  `json:"-"`
	HeaderParameter *HeaderParameter `json:"-"`
	CookieParameter *CookieParameter `json:"-"`
}

// WithPathParameter sets PathParameter value.
func (p *ParameterLocation) WithPathParameter(val PathParameter) *ParameterLocation {
	p.PathParameter = &val
	return p
}

// PathParameterEns ensures returned PathParameter is not nil.
func (p *ParameterLocation) PathParameterEns() *PathParameter {
	if p.PathParameter == nil {
		p.PathParameter = new(PathParameter)
	}

	return p.PathParameter
}

// WithQueryParameter sets QueryParameter value.
func (p *ParameterLocation) WithQueryParameter(val QueryParameter) *ParameterLocation {
	p.QueryParameter = &val
	return p
}

// QueryParameterEns ensures returned QueryParameter is not nil.
func (p *ParameterLocation) QueryParameterEns() *QueryParameter {
	if p.QueryParameter == nil {
		p.QueryParameter = new(QueryParameter)
	}

	return p.QueryParameter
}

// WithHeaderParameter sets HeaderParameter value.
func (p *ParameterLocation) WithHeaderParameter(val HeaderParameter) *ParameterLocation {
	p.HeaderParameter = &val
	return p
}

// HeaderParameterEns ensures returned HeaderParameter is not nil.
func (p *ParameterLocation) HeaderParameterEns() *HeaderParameter {
	if p.HeaderParameter == nil {
		p.HeaderParameter = new(HeaderParameter)
	}

	return p.HeaderParameter
}

// WithCookieParameter sets CookieParameter value.
func (p *ParameterLocation) WithCookieParameter(val CookieParameter) *ParameterLocation {
	p.CookieParameter = &val
	return p
}

// CookieParameterEns ensures returned CookieParameter is not nil.
func (p *ParameterLocation) CookieParameterEns() *CookieParameter {
	if p.CookieParameter == nil {
		p.CookieParameter = new(CookieParameter)
	}

	return p.CookieParameter
}

// UnmarshalJSON decodes JSON.
func (p *ParameterLocation) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 4)
	oneOfValid := 0

	err = json.Unmarshal(data, &p.PathParameter)
	if err != nil {
		oneOfErrors["PathParameter"] = err
		p.PathParameter = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &p.QueryParameter)
	if err != nil {
		oneOfErrors["QueryParameter"] = err
		p.QueryParameter = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &p.HeaderParameter)
	if err != nil {
		oneOfErrors["HeaderParameter"] = err
		p.HeaderParameter = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &p.CookieParameter)
	if err != nil {
		oneOfErrors["CookieParameter"] = err
		p.CookieParameter = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for ParameterLocation with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (p ParameterLocation) MarshalJSON() ([]byte, error) {
	return marshalUnion(p.PathParameter, p.QueryParameter, p.HeaderParameter, p.CookieParameter)
}

// ParameterOrRef structure is generated from "#/definitions/ParameterOrRef".
type ParameterOrRef struct {
	ParameterReference *ParameterReference `json:"-"`
	Parameter          *Parameter          `json:"-"`
}

// WithParameterReference sets ParameterReference value.
func (p *ParameterOrRef) WithParameterReference(val ParameterReference) *ParameterOrRef {
	p.ParameterReference = &val
	return p
}

// ParameterReferenceEns ensures returned ParameterReference is not nil.
func (p *ParameterOrRef) ParameterReferenceEns() *ParameterReference {
	if p.ParameterReference == nil {
		p.ParameterReference = new(ParameterReference)
	}

	return p.ParameterReference
}

// WithParameter sets Parameter value.
func (p *ParameterOrRef) WithParameter(val Parameter) *ParameterOrRef {
	p.Parameter = &val
	return p
}

// ParameterEns ensures returned Parameter is not nil.
func (p *ParameterOrRef) ParameterEns() *Parameter {
	if p.Parameter == nil {
		p.Parameter = new(Parameter)
	}

	return p.Parameter
}

// UnmarshalJSON decodes JSON.
func (p *ParameterOrRef) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 2)
	oneOfValid := 0

	err = json.Unmarshal(data, &p.ParameterReference)
	if err != nil {
		oneOfErrors["ParameterReference"] = err
		p.ParameterReference = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &p.Parameter)
	if err != nil {
		oneOfErrors["Parameter"] = err
		p.Parameter = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for ParameterOrRef with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (p ParameterOrRef) MarshalJSON() ([]byte, error) {
	return marshalUnion(p.ParameterReference, p.Parameter)
}

// Operation structure is generated from "#/definitions/Operation".
type Operation struct {
	Tags          []string                 `json:"tags,omitempty"`
	Summary       *string                  `json:"summary,omitempty"`
	Description   *string                  `json:"description,omitempty"`
	ExternalDocs  *ExternalDocumentation   `json:"externalDocs,omitempty"`
	ID            *string                  `json:"operationId,omitempty"`
	Parameters    []ParameterOrRef         `json:"parameters,omitempty"`
	RequestBody   *RequestBodyOrRef        `json:"requestBody,omitempty"`
	Responses     Responses                `json:"responses"` // Required.
	Callbacks     map[string]CallbackOrRef `json:"callbacks,omitempty"`
	Deprecated    *bool                    `json:"deprecated,omitempty"`
	Security      []map[string][]string    `json:"security,omitempty"`
	Servers       []Server                 `json:"servers,omitempty"`
	MapOfAnything map[string]interface{}   `json:"-"` // Key must match pattern: `^x-`.
}

// WithTags sets Tags value.
func (o *Operation) WithTags(val ...string) *Operation {
	o.Tags = val
	return o
}

// WithSummary sets Summary value.
func (o *Operation) WithSummary(val string) *Operation {
	o.Summary = &val
	return o
}

// WithDescription sets Description value.
func (o *Operation) WithDescription(val string) *Operation {
	o.Description = &val
	return o
}

// WithExternalDocs sets ExternalDocs value.
func (o *Operation) WithExternalDocs(val ExternalDocumentation) *Operation {
	o.ExternalDocs = &val
	return o
}

// ExternalDocsEns ensures returned ExternalDocs is not nil.
func (o *Operation) ExternalDocsEns() *ExternalDocumentation {
	if o.ExternalDocs == nil {
		o.ExternalDocs = new(ExternalDocumentation)
	}

	return o.ExternalDocs
}

// WithID sets ID value.
func (o *Operation) WithID(val string) *Operation {
	o.ID = &val
	return o
}

// WithParameters sets Parameters value.
func (o *Operation) WithParameters(val ...ParameterOrRef) *Operation {
	o.Parameters = val
	return o
}

// WithRequestBody sets RequestBody value.
func (o *Operation) WithRequestBody(val RequestBodyOrRef) *Operation {
	o.RequestBody = &val
	return o
}

// RequestBodyEns ensures returned RequestBody is not nil.
func (o *Operation) RequestBodyEns() *RequestBodyOrRef {
	if o.RequestBody == nil {
		o.RequestBody = new(RequestBodyOrRef)
	}

	return o.RequestBody
}

// WithResponses sets Responses value.
func (o *Operation) WithResponses(val Responses) *Operation {
	o.Responses = val
	return o
}

// WithCallbacks sets Callbacks value.
func (o *Operation) WithCallbacks(val map[string]CallbackOrRef) *Operation {
	o.Callbacks = val
	return o
}

// WithCallbacksItem sets Callbacks item value.
func (o *Operation) WithCallbacksItem(key string, val CallbackOrRef) *Operation {
	if o.Callbacks == nil {
		o.Callbacks = make(map[string]CallbackOrRef, 1)
	}

	o.Callbacks[key] = val

	return o
}

// WithDeprecated sets Deprecated value.
func (o *Operation) WithDeprecated(val bool) *Operation {
	o.Deprecated = &val
	return o
}

// WithSecurity sets Security value.
func (o *Operation) WithSecurity(val ...map[string][]string) *Operation {
	o.Security = val
	return o
}

// WithServers sets Servers value.
func (o *Operation) WithServers(val ...Server) *Operation {
	o.Servers = val
	return o
}

// WithMapOfAnything sets MapOfAnything value.
func (o *Operation) WithMapOfAnything(val map[string]interface{}) *Operation {
	o.MapOfAnything = val
	return o
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (o *Operation) WithMapOfAnythingItem(key string, val interface{}) *Operation {
	if o.MapOfAnything == nil {
		o.MapOfAnything = make(map[string]interface{}, 1)
	}

	o.MapOfAnything[key] = val

	return o
}

type marshalOperation Operation

var knownKeysOperation = []string{
	"tags",
	"summary",
	"description",
	"externalDocs",
	"operationId",
	"parameters",
	"requestBody",
	"responses",
	"callbacks",
	"deprecated",
	"security",
	"servers",
}

var requireKeysOperation = []string{
	"responses",
}

// UnmarshalJSON decodes JSON.
func (o *Operation) UnmarshalJSON(data []byte) error {
	var err error

	mo := marshalOperation(*o)

	err = json.Unmarshal(data, &mo)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysOperation {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysOperation {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if mo.MapOfAnything == nil {
				mo.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mo.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Operation: %v", offendingKeys)
	}

	*o = Operation(mo)

	return nil
}

// MarshalJSON encodes JSON.
func (o Operation) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOperation(o), o.MapOfAnything)
}

// RequestBodyReference structure is generated from "#/definitions/RequestBodyReference".
type RequestBodyReference struct {
	// Format: uri-reference.
	// Required.
	Ref string `json:"$ref"`
}

// WithRef sets Ref value.
func (r *RequestBodyReference) WithRef(val string) *RequestBodyReference {
	r.Ref = val
	return r
}

type marshalRequestBodyReference RequestBodyReference

var knownKeysRequestBodyReference = []string{
	"$ref",
}

var requireKeysRequestBodyReference = []string{
	"$ref",
}

// UnmarshalJSON decodes JSON.
func (r *RequestBodyReference) UnmarshalJSON(data []byte) error {
	var err error

	mr := marshalRequestBodyReference(*r)

	err = json.Unmarshal(data, &mr)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysRequestBodyReference {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysRequestBodyReference {
		delete(rawMap, key)
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in RequestBodyReference: %v", offendingKeys)
	}

	*r = RequestBodyReference(mr)

	return nil
}

// RequestBody structure is generated from "#/definitions/RequestBody".
type RequestBody struct {
	Description   *string                `json:"description,omitempty"`
	Content       map[string]MediaType   `json:"content"` // Required.
	Required      *bool                  `json:"required,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"` // Key must match pattern: `^x-`.
}

// WithDescription sets Description value.
func (r *RequestBody) WithDescription(val string) *RequestBody {
	r.Description = &val
	return r
}

// WithContent sets Content value.
func (r *RequestBody) WithContent(val map[string]MediaType) *RequestBody {
	r.Content = val
	return r
}

// WithContentItem sets Content item value.
func (r *RequestBody) WithContentItem(key string, val MediaType) *RequestBody {
	if r.Content == nil {
		r.Content = make(map[string]MediaType, 1)
	}

	r.Content[key] = val

	return r
}

// WithRequired sets Required value.
func (r *RequestBody) WithRequired(val bool) *RequestBody {
	r.Required = &val
	return r
}

// WithMapOfAnything sets MapOfAnything value.
func (r *RequestBody) WithMapOfAnything(val map[string]interface{}) *RequestBody {
	r.MapOfAnything = val
	return r
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (r *RequestBody) WithMapOfAnythingItem(key string, val interface{}) *RequestBody {
	if r.MapOfAnything == nil {
		r.MapOfAnything = make(map[string]interface{}, 1)
	}

	r.MapOfAnything[key] = val

	return r
}

type marshalRequestBody RequestBody

var knownKeysRequestBody = []string{
	"description",
	"content",
	"required",
}

var requireKeysRequestBody = []string{
	"content",
}

// UnmarshalJSON decodes JSON.
func (r *RequestBody) UnmarshalJSON(data []byte) error {
	var err error

	mr := marshalRequestBody(*r)

	err = json.Unmarshal(data, &mr)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysRequestBody {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysRequestBody {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if mr.MapOfAnything == nil {
				mr.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mr.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in RequestBody: %v", offendingKeys)
	}

	*r = RequestBody(mr)

	return nil
}

// MarshalJSON encodes JSON.
func (r RequestBody) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalRequestBody(r), r.MapOfAnything)
}

// RequestBodyOrRef structure is generated from "#/definitions/RequestBodyOrRef".
type RequestBodyOrRef struct {
	RequestBodyReference *RequestBodyReference `json:"-"`
	RequestBody          *RequestBody          `json:"-"`
}

// WithRequestBodyReference sets RequestBodyReference value.
func (r *RequestBodyOrRef) WithRequestBodyReference(val RequestBodyReference) *RequestBodyOrRef {
	r.RequestBodyReference = &val
	return r
}

// RequestBodyReferenceEns ensures returned RequestBodyReference is not nil.
func (r *RequestBodyOrRef) RequestBodyReferenceEns() *RequestBodyReference {
	if r.RequestBodyReference == nil {
		r.RequestBodyReference = new(RequestBodyReference)
	}

	return r.RequestBodyReference
}

// WithRequestBody sets RequestBody value.
func (r *RequestBodyOrRef) WithRequestBody(val RequestBody) *RequestBodyOrRef {
	r.RequestBody = &val
	return r
}

// RequestBodyEns ensures returned RequestBody is not nil.
func (r *RequestBodyOrRef) RequestBodyEns() *RequestBody {
	if r.RequestBody == nil {
		r.RequestBody = new(RequestBody)
	}

	return r.RequestBody
}

// UnmarshalJSON decodes JSON.
func (r *RequestBodyOrRef) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 2)
	oneOfValid := 0

	err = json.Unmarshal(data, &r.RequestBodyReference)
	if err != nil {
		oneOfErrors["RequestBodyReference"] = err
		r.RequestBodyReference = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &r.RequestBody)
	if err != nil {
		oneOfErrors["RequestBody"] = err
		r.RequestBody = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for RequestBodyOrRef with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (r RequestBodyOrRef) MarshalJSON() ([]byte, error) {
	return marshalUnion(r.RequestBodyReference, r.RequestBody)
}

// Responses structure is generated from "#/definitions/Responses".
type Responses struct {
	Default                  *ResponseOrRef           `json:"default,omitempty"`
	MapOfResponseOrRefValues map[string]ResponseOrRef `json:"-"` // Key must match pattern: `^[1-5](?:\d{2}|XX)$`.
	MapOfAnything            map[string]interface{}   `json:"-"` // Key must match pattern: `^x-`.
}

// WithDefault sets Default value.
func (r *Responses) WithDefault(val ResponseOrRef) *Responses {
	r.Default = &val
	return r
}

// DefaultEns ensures returned Default is not nil.
func (r *Responses) DefaultEns() *ResponseOrRef {
	if r.Default == nil {
		r.Default = new(ResponseOrRef)
	}

	return r.Default
}

// WithMapOfResponseOrRefValues sets MapOfResponseOrRefValues value.
func (r *Responses) WithMapOfResponseOrRefValues(val map[string]ResponseOrRef) *Responses {
	r.MapOfResponseOrRefValues = val
	return r
}

// WithMapOfResponseOrRefValuesItem sets MapOfResponseOrRefValues item value.
func (r *Responses) WithMapOfResponseOrRefValuesItem(key string, val ResponseOrRef) *Responses {
	if r.MapOfResponseOrRefValues == nil {
		r.MapOfResponseOrRefValues = make(map[string]ResponseOrRef, 1)
	}

	r.MapOfResponseOrRefValues[key] = val

	return r
}

// WithMapOfAnything sets MapOfAnything value.
func (r *Responses) WithMapOfAnything(val map[string]interface{}) *Responses {
	r.MapOfAnything = val
	return r
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (r *Responses) WithMapOfAnythingItem(key string, val interface{}) *Responses {
	if r.MapOfAnything == nil {
		r.MapOfAnything = make(map[string]interface{}, 1)
	}

	r.MapOfAnything[key] = val

	return r
}

type marshalResponses Responses

var knownKeysResponses = []string{
	"default",
}

// UnmarshalJSON decodes JSON.
func (r *Responses) UnmarshalJSON(data []byte) error {
	var err error

	mr := marshalResponses(*r)

	err = json.Unmarshal(data, &mr)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range knownKeysResponses {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regex15D2XX.MatchString(key) {
			matched = true

			if mr.MapOfResponseOrRefValues == nil {
				mr.MapOfResponseOrRefValues = make(map[string]ResponseOrRef, 1)
			}

			var val ResponseOrRef

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mr.MapOfResponseOrRefValues[key] = val
		}

		if regexX.MatchString(key) {
			matched = true

			if mr.MapOfAnything == nil {
				mr.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mr.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Responses: %v", offendingKeys)
	}

	*r = Responses(mr)

	return nil
}

// MarshalJSON encodes JSON.
func (r Responses) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalResponses(r), r.MapOfResponseOrRefValues, r.MapOfAnything)
}

// ResponseReference structure is generated from "#/definitions/ResponseReference".
type ResponseReference struct {
	// Format: uri-reference.
	// Required.
	Ref string `json:"$ref"`
}

// WithRef sets Ref value.
func (r *ResponseReference) WithRef(val string) *ResponseReference {
	r.Ref = val
	return r
}

type marshalResponseReference ResponseReference

var knownKeysResponseReference = []string{
	"$ref",
}

var requireKeysResponseReference = []string{
	"$ref",
}

// UnmarshalJSON decodes JSON.
func (r *ResponseReference) UnmarshalJSON(data []byte) error {
	var err error

	mr := marshalResponseReference(*r)

	err = json.Unmarshal(data, &mr)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysResponseReference {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysResponseReference {
		delete(rawMap, key)
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in ResponseReference: %v", offendingKeys)
	}

	*r = ResponseReference(mr)

	return nil
}

// Response structure is generated from "#/definitions/Response".
type Response struct {
	Description   string                 `json:"description"` // Required.
	Headers       map[string]HeaderOrRef `json:"headers,omitempty"`
	Content       map[string]MediaType   `json:"content,omitempty"`
	Links         map[string]LinkOrRef   `json:"links,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"` // Key must match pattern: `^x-`.
}

// WithDescription sets Description value.
func (r *Response) WithDescription(val string) *Response {
	r.Description = val
	return r
}

// WithHeaders sets Headers value.
func (r *Response) WithHeaders(val map[string]HeaderOrRef) *Response {
	r.Headers = val
	return r
}

// WithHeadersItem sets Headers item value.
func (r *Response) WithHeadersItem(key string, val HeaderOrRef) *Response {
	if r.Headers == nil {
		r.Headers = make(map[string]HeaderOrRef, 1)
	}

	r.Headers[key] = val

	return r
}

// WithContent sets Content value.
func (r *Response) WithContent(val map[string]MediaType) *Response {
	r.Content = val
	return r
}

// WithContentItem sets Content item value.
func (r *Response) WithContentItem(key string, val MediaType) *Response {
	if r.Content == nil {
		r.Content = make(map[string]MediaType, 1)
	}

	r.Content[key] = val

	return r
}

// WithLinks sets Links value.
func (r *Response) WithLinks(val map[string]LinkOrRef) *Response {
	r.Links = val
	return r
}

// WithLinksItem sets Links item value.
func (r *Response) WithLinksItem(key string, val LinkOrRef) *Response {
	if r.Links == nil {
		r.Links = make(map[string]LinkOrRef, 1)
	}

	r.Links[key] = val

	return r
}

// WithMapOfAnything sets MapOfAnything value.
func (r *Response) WithMapOfAnything(val map[string]interface{}) *Response {
	r.MapOfAnything = val
	return r
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (r *Response) WithMapOfAnythingItem(key string, val interface{}) *Response {
	if r.MapOfAnything == nil {
		r.MapOfAnything = make(map[string]interface{}, 1)
	}

	r.MapOfAnything[key] = val

	return r
}

type marshalResponse Response

var knownKeysResponse = []string{
	"description",
	"headers",
	"content",
	"links",
}

var requireKeysResponse = []string{
	"description",
}

// UnmarshalJSON decodes JSON.
func (r *Response) UnmarshalJSON(data []byte) error {
	var err error

	mr := marshalResponse(*r)

	err = json.Unmarshal(data, &mr)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysResponse {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysResponse {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if mr.MapOfAnything == nil {
				mr.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mr.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Response: %v", offendingKeys)
	}

	*r = Response(mr)

	return nil
}

// MarshalJSON encodes JSON.
func (r Response) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalResponse(r), r.MapOfAnything)
}

// HeaderReference structure is generated from "#/definitions/HeaderReference".
type HeaderReference struct {
	// Format: uri-reference.
	// Required.
	Ref string `json:"$ref"`
}

// WithRef sets Ref value.
func (h *HeaderReference) WithRef(val string) *HeaderReference {
	h.Ref = val
	return h
}

type marshalHeaderReference HeaderReference

var knownKeysHeaderReference = []string{
	"$ref",
}

var requireKeysHeaderReference = []string{
	"$ref",
}

// UnmarshalJSON decodes JSON.
func (h *HeaderReference) UnmarshalJSON(data []byte) error {
	var err error

	mh := marshalHeaderReference(*h)

	err = json.Unmarshal(data, &mh)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysHeaderReference {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysHeaderReference {
		delete(rawMap, key)
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in HeaderReference: %v", offendingKeys)
	}

	*h = HeaderReference(mh)

	return nil
}

// HeaderOrRef structure is generated from "#/definitions/HeaderOrRef".
type HeaderOrRef struct {
	HeaderReference *HeaderReference `json:"-"`
	Header          *Header          `json:"-"`
}

// WithHeaderReference sets HeaderReference value.
func (h *HeaderOrRef) WithHeaderReference(val HeaderReference) *HeaderOrRef {
	h.HeaderReference = &val
	return h
}

// HeaderReferenceEns ensures returned HeaderReference is not nil.
func (h *HeaderOrRef) HeaderReferenceEns() *HeaderReference {
	if h.HeaderReference == nil {
		h.HeaderReference = new(HeaderReference)
	}

	return h.HeaderReference
}

// WithHeader sets Header value.
func (h *HeaderOrRef) WithHeader(val Header) *HeaderOrRef {
	h.Header = &val
	return h
}

// HeaderEns ensures returned Header is not nil.
func (h *HeaderOrRef) HeaderEns() *Header {
	if h.Header == nil {
		h.Header = new(Header)
	}

	return h.Header
}

// UnmarshalJSON decodes JSON.
func (h *HeaderOrRef) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 2)
	oneOfValid := 0

	err = json.Unmarshal(data, &h.HeaderReference)
	if err != nil {
		oneOfErrors["HeaderReference"] = err
		h.HeaderReference = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &h.Header)
	if err != nil {
		oneOfErrors["Header"] = err
		h.Header = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for HeaderOrRef with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (h HeaderOrRef) MarshalJSON() ([]byte, error) {
	return marshalUnion(h.HeaderReference, h.Header)
}

// LinkReference structure is generated from "#/definitions/LinkReference".
type LinkReference struct {
	// Format: uri-reference.
	// Required.
	Ref string `json:"$ref"`
}

// WithRef sets Ref value.
func (l *LinkReference) WithRef(val string) *LinkReference {
	l.Ref = val
	return l
}

type marshalLinkReference LinkReference

var knownKeysLinkReference = []string{
	"$ref",
}

var requireKeysLinkReference = []string{
	"$ref",
}

// UnmarshalJSON decodes JSON.
func (l *LinkReference) UnmarshalJSON(data []byte) error {
	var err error

	ml := marshalLinkReference(*l)

	err = json.Unmarshal(data, &ml)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysLinkReference {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysLinkReference {
		delete(rawMap, key)
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in LinkReference: %v", offendingKeys)
	}

	*l = LinkReference(ml)

	return nil
}

// Link structure is generated from "#/definitions/Link".
type Link struct {
	OperationID   *string                `json:"operationId,omitempty"`
	OperationRef  *string                `json:"operationRef,omitempty"` // Format: uri-reference.
	Parameters    map[string]interface{} `json:"parameters,omitempty"`
	RequestBody   *interface{}           `json:"requestBody,omitempty"`
	Description   *string                `json:"description,omitempty"`
	Server        *Server                `json:"server,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"` // Key must match pattern: `^x-`.
}

// WithOperationID sets OperationID value.
func (l *Link) WithOperationID(val string) *Link {
	l.OperationID = &val
	return l
}

// WithOperationRef sets OperationRef value.
func (l *Link) WithOperationRef(val string) *Link {
	l.OperationRef = &val
	return l
}

// WithParameters sets Parameters value.
func (l *Link) WithParameters(val map[string]interface{}) *Link {
	l.Parameters = val
	return l
}

// WithParametersItem sets Parameters item value.
func (l *Link) WithParametersItem(key string, val interface{}) *Link {
	if l.Parameters == nil {
		l.Parameters = make(map[string]interface{}, 1)
	}

	l.Parameters[key] = val

	return l
}

// WithRequestBody sets RequestBody value.
func (l *Link) WithRequestBody(val interface{}) *Link {
	l.RequestBody = &val
	return l
}

// WithDescription sets Description value.
func (l *Link) WithDescription(val string) *Link {
	l.Description = &val
	return l
}

// WithServer sets Server value.
func (l *Link) WithServer(val Server) *Link {
	l.Server = &val
	return l
}

// ServerEns ensures returned Server is not nil.
func (l *Link) ServerEns() *Server {
	if l.Server == nil {
		l.Server = new(Server)
	}

	return l.Server
}

// WithMapOfAnything sets MapOfAnything value.
func (l *Link) WithMapOfAnything(val map[string]interface{}) *Link {
	l.MapOfAnything = val
	return l
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (l *Link) WithMapOfAnythingItem(key string, val interface{}) *Link {
	if l.MapOfAnything == nil {
		l.MapOfAnything = make(map[string]interface{}, 1)
	}

	l.MapOfAnything[key] = val

	return l
}

type marshalLink Link

var knownKeysLink = []string{
	"operationId",
	"operationRef",
	"parameters",
	"requestBody",
	"description",
	"server",
}

// UnmarshalJSON decodes JSON.
func (l *Link) UnmarshalJSON(data []byte) error {
	var err error

	var not LinkNot

	if json.Unmarshal(data, &not) == nil {
		return errors.New("not constraint failed for Link")
	}

	ml := marshalLink(*l)

	err = json.Unmarshal(data, &ml)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if ml.RequestBody == nil {
		if _, ok := rawMap["requestBody"]; ok {
			var v interface{}
			ml.RequestBody = &v
		}
	}

	for _, key := range knownKeysLink {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ml.MapOfAnything == nil {
				ml.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ml.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Link: %v", offendingKeys)
	}

	*l = Link(ml)

	return nil
}

// MarshalJSON encodes JSON.
func (l Link) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalLink(l), l.MapOfAnything)
}

// LinkNot structure is generated from "#/definitions/Link->not".
//
// Operation Id and Operation Ref are mutually exclusive.
type LinkNot struct {
	OperationID  interface{} `json:"operationId"`  // Required.
	OperationRef interface{} `json:"operationRef"` // Required.
}

// WithOperationID sets OperationID value.
func (l *LinkNot) WithOperationID(val interface{}) *LinkNot {
	l.OperationID = val
	return l
}

// WithOperationRef sets OperationRef value.
func (l *LinkNot) WithOperationRef(val interface{}) *LinkNot {
	l.OperationRef = val
	return l
}

type marshalLinkNot LinkNot

var requireKeysLinkNot = []string{
	"operationId",
	"operationRef",
}

// UnmarshalJSON decodes JSON.
func (l *LinkNot) UnmarshalJSON(data []byte) error {
	var err error

	ml := marshalLinkNot(*l)

	err = json.Unmarshal(data, &ml)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysLinkNot {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	*l = LinkNot(ml)

	return nil
}

// LinkOrRef structure is generated from "#/definitions/LinkOrRef".
type LinkOrRef struct {
	LinkReference *LinkReference `json:"-"`
	Link          *Link          `json:"-"`
}

// WithLinkReference sets LinkReference value.
func (l *LinkOrRef) WithLinkReference(val LinkReference) *LinkOrRef {
	l.LinkReference = &val
	return l
}

// LinkReferenceEns ensures returned LinkReference is not nil.
func (l *LinkOrRef) LinkReferenceEns() *LinkReference {
	if l.LinkReference == nil {
		l.LinkReference = new(LinkReference)
	}

	return l.LinkReference
}

// WithLink sets Link value.
func (l *LinkOrRef) WithLink(val Link) *LinkOrRef {
	l.Link = &val
	return l
}

// LinkEns ensures returned Link is not nil.
func (l *LinkOrRef) LinkEns() *Link {
	if l.Link == nil {
		l.Link = new(Link)
	}

	return l.Link
}

// UnmarshalJSON decodes JSON.
func (l *LinkOrRef) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 2)
	oneOfValid := 0

	err = json.Unmarshal(data, &l.LinkReference)
	if err != nil {
		oneOfErrors["LinkReference"] = err
		l.LinkReference = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &l.Link)
	if err != nil {
		oneOfErrors["Link"] = err
		l.Link = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for LinkOrRef with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (l LinkOrRef) MarshalJSON() ([]byte, error) {
	return marshalUnion(l.LinkReference, l.Link)
}

// ResponseOrRef structure is generated from "#/definitions/ResponseOrRef".
type ResponseOrRef struct {
	ResponseReference *ResponseReference `json:"-"`
	Response          *Response          `json:"-"`
}

// WithResponseReference sets ResponseReference value.
func (r *ResponseOrRef) WithResponseReference(val ResponseReference) *ResponseOrRef {
	r.ResponseReference = &val
	return r
}

// ResponseReferenceEns ensures returned ResponseReference is not nil.
func (r *ResponseOrRef) ResponseReferenceEns() *ResponseReference {
	if r.ResponseReference == nil {
		r.ResponseReference = new(ResponseReference)
	}

	return r.ResponseReference
}

// WithResponse sets Response value.
func (r *ResponseOrRef) WithResponse(val Response) *ResponseOrRef {
	r.Response = &val
	return r
}

// ResponseEns ensures returned Response is not nil.
func (r *ResponseOrRef) ResponseEns() *Response {
	if r.Response == nil {
		r.Response = new(Response)
	}

	return r.Response
}

// UnmarshalJSON decodes JSON.
func (r *ResponseOrRef) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 2)
	oneOfValid := 0

	err = json.Unmarshal(data, &r.ResponseReference)
	if err != nil {
		oneOfErrors["ResponseReference"] = err
		r.ResponseReference = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &r.Response)
	if err != nil {
		oneOfErrors["Response"] = err
		r.Response = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for ResponseOrRef with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (r ResponseOrRef) MarshalJSON() ([]byte, error) {
	return marshalUnion(r.ResponseReference, r.Response)
}

// CallbackReference structure is generated from "#/definitions/CallbackReference".
type CallbackReference struct {
	// Format: uri-reference.
	// Required.
	Ref string `json:"$ref"`
}

// WithRef sets Ref value.
func (c *CallbackReference) WithRef(val string) *CallbackReference {
	c.Ref = val
	return c
}

type marshalCallbackReference CallbackReference

var knownKeysCallbackReference = []string{
	"$ref",
}

var requireKeysCallbackReference = []string{
	"$ref",
}

// UnmarshalJSON decodes JSON.
func (c *CallbackReference) UnmarshalJSON(data []byte) error {
	var err error

	mc := marshalCallbackReference(*c)

	err = json.Unmarshal(data, &mc)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysCallbackReference {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysCallbackReference {
		delete(rawMap, key)
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in CallbackReference: %v", offendingKeys)
	}

	*c = CallbackReference(mc)

	return nil
}

// Callback structure is generated from "#/definitions/Callback".
type Callback struct {
	MapOfAnything        map[string]interface{} `json:"-"` // Key must match pattern: `^x-`.
	AdditionalProperties map[string]PathItem    `json:"-"` // All unmatched properties.
}

// WithMapOfAnything sets MapOfAnything value.
func (c *Callback) WithMapOfAnything(val map[string]interface{}) *Callback {
	c.MapOfAnything = val
	return c
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (c *Callback) WithMapOfAnythingItem(key string, val interface{}) *Callback {
	if c.MapOfAnything == nil {
		c.MapOfAnything = make(map[string]interface{}, 1)
	}

	c.MapOfAnything[key] = val

	return c
}

// WithAdditionalProperties sets AdditionalProperties value.
func (c *Callback) WithAdditionalProperties(val map[string]PathItem) *Callback {
	c.AdditionalProperties = val
	return c
}

// WithAdditionalPropertiesItem sets AdditionalProperties item value.
func (c *Callback) WithAdditionalPropertiesItem(key string, val PathItem) *Callback {
	if c.AdditionalProperties == nil {
		c.AdditionalProperties = make(map[string]PathItem, 1)
	}

	c.AdditionalProperties[key] = val

	return c
}

// UnmarshalJSON decodes JSON.
func (c *Callback) UnmarshalJSON(data []byte) error {
	var err error

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if c.MapOfAnything == nil {
				c.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			c.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	for key, rawValue := range rawMap {
		if c.AdditionalProperties == nil {
			c.AdditionalProperties = make(map[string]PathItem, 1)
		}

		var val PathItem

		err = json.Unmarshal(rawValue, &val)
		if err != nil {
			return err
		}

		c.AdditionalProperties[key] = val
	}

	return nil
}

// MarshalJSON encodes JSON.
func (c Callback) MarshalJSON() ([]byte, error) {
	return marshalUnion(c.MapOfAnything, c.AdditionalProperties)
}

// CallbackOrRef structure is generated from "#/definitions/CallbackOrRef".
type CallbackOrRef struct {
	CallbackReference *CallbackReference `json:"-"`
	Callback          *Callback          `json:"-"`
}

// WithCallbackReference sets CallbackReference value.
func (c *CallbackOrRef) WithCallbackReference(val CallbackReference) *CallbackOrRef {
	c.CallbackReference = &val
	return c
}

// CallbackReferenceEns ensures returned CallbackReference is not nil.
func (c *CallbackOrRef) CallbackReferenceEns() *CallbackReference {
	if c.CallbackReference == nil {
		c.CallbackReference = new(CallbackReference)
	}

	return c.CallbackReference
}

// WithCallback sets Callback value.
func (c *CallbackOrRef) WithCallback(val Callback) *CallbackOrRef {
	c.Callback = &val
	return c
}

// CallbackEns ensures returned Callback is not nil.
func (c *CallbackOrRef) CallbackEns() *Callback {
	if c.Callback == nil {
		c.Callback = new(Callback)
	}

	return c.Callback
}

// UnmarshalJSON decodes JSON.
func (c *CallbackOrRef) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 2)
	oneOfValid := 0

	err = json.Unmarshal(data, &c.CallbackReference)
	if err != nil {
		oneOfErrors["CallbackReference"] = err
		c.CallbackReference = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &c.Callback)
	if err != nil {
		oneOfErrors["Callback"] = err
		c.Callback = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for CallbackOrRef with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (c CallbackOrRef) MarshalJSON() ([]byte, error) {
	return marshalUnion(c.CallbackReference, c.Callback)
}

// Paths structure is generated from "#/definitions/Paths".
type Paths struct {
	MapOfPathItemValues map[string]PathItem    `json:"-"` // Key must match pattern: `^\/`.
	MapOfAnything       map[string]interface{} `json:"-"` // Key must match pattern: `^x-`.
}

// WithMapOfPathItemValues sets MapOfPathItemValues value.
func (p *Paths) WithMapOfPathItemValues(val map[string]PathItem) *Paths {
	p.MapOfPathItemValues = val
	return p
}

// WithMapOfPathItemValuesItem sets MapOfPathItemValues item value.
func (p *Paths) WithMapOfPathItemValuesItem(key string, val PathItem) *Paths {
	if p.MapOfPathItemValues == nil {
		p.MapOfPathItemValues = make(map[string]PathItem, 1)
	}

	p.MapOfPathItemValues[key] = val

	return p
}

// WithMapOfAnything sets MapOfAnything value.
func (p *Paths) WithMapOfAnything(val map[string]interface{}) *Paths {
	p.MapOfAnything = val
	return p
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (p *Paths) WithMapOfAnythingItem(key string, val interface{}) *Paths {
	if p.MapOfAnything == nil {
		p.MapOfAnything = make(map[string]interface{}, 1)
	}

	p.MapOfAnything[key] = val

	return p
}

// UnmarshalJSON decodes JSON.
func (p *Paths) UnmarshalJSON(data []byte) error {
	var err error

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for key, rawValue := range rawMap {
		matched := false

		if regex.MatchString(key) {
			matched = true

			if p.MapOfPathItemValues == nil {
				p.MapOfPathItemValues = make(map[string]PathItem, 1)
			}

			var val PathItem

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			p.MapOfPathItemValues[key] = val
		}

		if regexX.MatchString(key) {
			matched = true

			if p.MapOfAnything == nil {
				p.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			p.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Paths: %v", offendingKeys)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (p Paths) MarshalJSON() ([]byte, error) {
	return marshalUnion(p.MapOfPathItemValues, p.MapOfAnything)
}

// Components structure is generated from "#/definitions/Components".
type Components struct {
	Schemas         *ComponentsSchemas         `json:"schemas,omitempty"`
	Responses       *ComponentsResponses       `json:"responses,omitempty"`
	Parameters      *ComponentsParameters      `json:"parameters,omitempty"`
	Examples        *ComponentsExamples        `json:"examples,omitempty"`
	RequestBodies   *ComponentsRequestBodies   `json:"requestBodies,omitempty"`
	Headers         *ComponentsHeaders         `json:"headers,omitempty"`
	SecuritySchemes *ComponentsSecuritySchemes `json:"securitySchemes,omitempty"`
	Links           *ComponentsLinks           `json:"links,omitempty"`
	Callbacks       *ComponentsCallbacks       `json:"callbacks,omitempty"`
	MapOfAnything   map[string]interface{}     `json:"-"` // Key must match pattern: `^x-`.
}

// WithSchemas sets Schemas value.
func (c *Components) WithSchemas(val ComponentsSchemas) *Components {
	c.Schemas = &val
	return c
}

// SchemasEns ensures returned Schemas is not nil.
func (c *Components) SchemasEns() *ComponentsSchemas {
	if c.Schemas == nil {
		c.Schemas = new(ComponentsSchemas)
	}

	return c.Schemas
}

// WithResponses sets Responses value.
func (c *Components) WithResponses(val ComponentsResponses) *Components {
	c.Responses = &val
	return c
}

// ResponsesEns ensures returned Responses is not nil.
func (c *Components) ResponsesEns() *ComponentsResponses {
	if c.Responses == nil {
		c.Responses = new(ComponentsResponses)
	}

	return c.Responses
}

// WithParameters sets Parameters value.
func (c *Components) WithParameters(val ComponentsParameters) *Components {
	c.Parameters = &val
	return c
}

// ParametersEns ensures returned Parameters is not nil.
func (c *Components) ParametersEns() *ComponentsParameters {
	if c.Parameters == nil {
		c.Parameters = new(ComponentsParameters)
	}

	return c.Parameters
}

// WithExamples sets Examples value.
func (c *Components) WithExamples(val ComponentsExamples) *Components {
	c.Examples = &val
	return c
}

// ExamplesEns ensures returned Examples is not nil.
func (c *Components) ExamplesEns() *ComponentsExamples {
	if c.Examples == nil {
		c.Examples = new(ComponentsExamples)
	}

	return c.Examples
}

// WithRequestBodies sets RequestBodies value.
func (c *Components) WithRequestBodies(val ComponentsRequestBodies) *Components {
	c.RequestBodies = &val
	return c
}

// RequestBodiesEns ensures returned RequestBodies is not nil.
func (c *Components) RequestBodiesEns() *ComponentsRequestBodies {
	if c.RequestBodies == nil {
		c.RequestBodies = new(ComponentsRequestBodies)
	}

	return c.RequestBodies
}

// WithHeaders sets Headers value.
func (c *Components) WithHeaders(val ComponentsHeaders) *Components {
	c.Headers = &val
	return c
}

// HeadersEns ensures returned Headers is not nil.
func (c *Components) HeadersEns() *ComponentsHeaders {
	if c.Headers == nil {
		c.Headers = new(ComponentsHeaders)
	}

	return c.Headers
}

// WithSecuritySchemes sets SecuritySchemes value.
func (c *Components) WithSecuritySchemes(val ComponentsSecuritySchemes) *Components {
	c.SecuritySchemes = &val
	return c
}

// SecuritySchemesEns ensures returned SecuritySchemes is not nil.
func (c *Components) SecuritySchemesEns() *ComponentsSecuritySchemes {
	if c.SecuritySchemes == nil {
		c.SecuritySchemes = new(ComponentsSecuritySchemes)
	}

	return c.SecuritySchemes
}

// WithLinks sets Links value.
func (c *Components) WithLinks(val ComponentsLinks) *Components {
	c.Links = &val
	return c
}

// LinksEns ensures returned Links is not nil.
func (c *Components) LinksEns() *ComponentsLinks {
	if c.Links == nil {
		c.Links = new(ComponentsLinks)
	}

	return c.Links
}

// WithCallbacks sets Callbacks value.
func (c *Components) WithCallbacks(val ComponentsCallbacks) *Components {
	c.Callbacks = &val
	return c
}

// CallbacksEns ensures returned Callbacks is not nil.
func (c *Components) CallbacksEns() *ComponentsCallbacks {
	if c.Callbacks == nil {
		c.Callbacks = new(ComponentsCallbacks)
	}

	return c.Callbacks
}

// WithMapOfAnything sets MapOfAnything value.
func (c *Components) WithMapOfAnything(val map[string]interface{}) *Components {
	c.MapOfAnything = val
	return c
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (c *Components) WithMapOfAnythingItem(key string, val interface{}) *Components {
	if c.MapOfAnything == nil {
		c.MapOfAnything = make(map[string]interface{}, 1)
	}

	c.MapOfAnything[key] = val

	return c
}

type marshalComponents Components

var knownKeysComponents = []string{
	"schemas",
	"responses",
	"parameters",
	"examples",
	"requestBodies",
	"headers",
	"securitySchemes",
	"links",
	"callbacks",
}

// UnmarshalJSON decodes JSON.
func (c *Components) UnmarshalJSON(data []byte) error {
	var err error

	mc := marshalComponents(*c)

	err = json.Unmarshal(data, &mc)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range knownKeysComponents {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if mc.MapOfAnything == nil {
				mc.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mc.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in Components: %v", offendingKeys)
	}

	*c = Components(mc)

	return nil
}

// MarshalJSON encodes JSON.
func (c Components) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalComponents(c), c.MapOfAnything)
}

// ComponentsSchemas structure is generated from "#/definitions/Components->schemas".
type ComponentsSchemas struct {
	MapOfSchemaOrRefValues map[string]SchemaOrRef `json:"-"` // Key must match pattern: `^[a-zA-Z0-9\.\-_]+$`.
}

// WithMapOfSchemaOrRefValues sets MapOfSchemaOrRefValues value.
func (c *ComponentsSchemas) WithMapOfSchemaOrRefValues(val map[string]SchemaOrRef) *ComponentsSchemas {
	c.MapOfSchemaOrRefValues = val
	return c
}

// WithMapOfSchemaOrRefValuesItem sets MapOfSchemaOrRefValues item value.
func (c *ComponentsSchemas) WithMapOfSchemaOrRefValuesItem(key string, val SchemaOrRef) *ComponentsSchemas {
	if c.MapOfSchemaOrRefValues == nil {
		c.MapOfSchemaOrRefValues = make(map[string]SchemaOrRef, 1)
	}

	c.MapOfSchemaOrRefValues[key] = val

	return c
}

// UnmarshalJSON decodes JSON.
func (c *ComponentsSchemas) UnmarshalJSON(data []byte) error {
	var err error

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexAZAZ09.MatchString(key) {
			matched = true

			if c.MapOfSchemaOrRefValues == nil {
				c.MapOfSchemaOrRefValues = make(map[string]SchemaOrRef, 1)
			}

			var val SchemaOrRef

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			c.MapOfSchemaOrRefValues[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	return nil
}

// MarshalJSON encodes JSON.
func (c ComponentsSchemas) MarshalJSON() ([]byte, error) {
	return marshalUnion(c.MapOfSchemaOrRefValues)
}

// ComponentsResponses structure is generated from "#/definitions/Components->responses".
type ComponentsResponses struct {
	MapOfResponseOrRefValues map[string]ResponseOrRef `json:"-"` // Key must match pattern: `^[a-zA-Z0-9\.\-_]+$`.
}

// WithMapOfResponseOrRefValues sets MapOfResponseOrRefValues value.
func (c *ComponentsResponses) WithMapOfResponseOrRefValues(val map[string]ResponseOrRef) *ComponentsResponses {
	c.MapOfResponseOrRefValues = val
	return c
}

// WithMapOfResponseOrRefValuesItem sets MapOfResponseOrRefValues item value.
func (c *ComponentsResponses) WithMapOfResponseOrRefValuesItem(key string, val ResponseOrRef) *ComponentsResponses {
	if c.MapOfResponseOrRefValues == nil {
		c.MapOfResponseOrRefValues = make(map[string]ResponseOrRef, 1)
	}

	c.MapOfResponseOrRefValues[key] = val

	return c
}

// UnmarshalJSON decodes JSON.
func (c *ComponentsResponses) UnmarshalJSON(data []byte) error {
	var err error

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexAZAZ09.MatchString(key) {
			matched = true

			if c.MapOfResponseOrRefValues == nil {
				c.MapOfResponseOrRefValues = make(map[string]ResponseOrRef, 1)
			}

			var val ResponseOrRef

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			c.MapOfResponseOrRefValues[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	return nil
}

// MarshalJSON encodes JSON.
func (c ComponentsResponses) MarshalJSON() ([]byte, error) {
	return marshalUnion(c.MapOfResponseOrRefValues)
}

// ComponentsParameters structure is generated from "#/definitions/Components->parameters".
type ComponentsParameters struct {
	MapOfParameterOrRefValues map[string]ParameterOrRef `json:"-"` // Key must match pattern: `^[a-zA-Z0-9\.\-_]+$`.
}

// WithMapOfParameterOrRefValues sets MapOfParameterOrRefValues value.
func (c *ComponentsParameters) WithMapOfParameterOrRefValues(val map[string]ParameterOrRef) *ComponentsParameters {
	c.MapOfParameterOrRefValues = val
	return c
}

// WithMapOfParameterOrRefValuesItem sets MapOfParameterOrRefValues item value.
func (c *ComponentsParameters) WithMapOfParameterOrRefValuesItem(key string, val ParameterOrRef) *ComponentsParameters {
	if c.MapOfParameterOrRefValues == nil {
		c.MapOfParameterOrRefValues = make(map[string]ParameterOrRef, 1)
	}

	c.MapOfParameterOrRefValues[key] = val

	return c
}

// UnmarshalJSON decodes JSON.
func (c *ComponentsParameters) UnmarshalJSON(data []byte) error {
	var err error

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexAZAZ09.MatchString(key) {
			matched = true

			if c.MapOfParameterOrRefValues == nil {
				c.MapOfParameterOrRefValues = make(map[string]ParameterOrRef, 1)
			}

			var val ParameterOrRef

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			c.MapOfParameterOrRefValues[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	return nil
}

// MarshalJSON encodes JSON.
func (c ComponentsParameters) MarshalJSON() ([]byte, error) {
	return marshalUnion(c.MapOfParameterOrRefValues)
}

// ComponentsExamples structure is generated from "#/definitions/Components->examples".
type ComponentsExamples struct {
	MapOfExampleOrRefValues map[string]ExampleOrRef `json:"-"` // Key must match pattern: `^[a-zA-Z0-9\.\-_]+$`.
}

// WithMapOfExampleOrRefValues sets MapOfExampleOrRefValues value.
func (c *ComponentsExamples) WithMapOfExampleOrRefValues(val map[string]ExampleOrRef) *ComponentsExamples {
	c.MapOfExampleOrRefValues = val
	return c
}

// WithMapOfExampleOrRefValuesItem sets MapOfExampleOrRefValues item value.
func (c *ComponentsExamples) WithMapOfExampleOrRefValuesItem(key string, val ExampleOrRef) *ComponentsExamples {
	if c.MapOfExampleOrRefValues == nil {
		c.MapOfExampleOrRefValues = make(map[string]ExampleOrRef, 1)
	}

	c.MapOfExampleOrRefValues[key] = val

	return c
}

// UnmarshalJSON decodes JSON.
func (c *ComponentsExamples) UnmarshalJSON(data []byte) error {
	var err error

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexAZAZ09.MatchString(key) {
			matched = true

			if c.MapOfExampleOrRefValues == nil {
				c.MapOfExampleOrRefValues = make(map[string]ExampleOrRef, 1)
			}

			var val ExampleOrRef

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			c.MapOfExampleOrRefValues[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	return nil
}

// MarshalJSON encodes JSON.
func (c ComponentsExamples) MarshalJSON() ([]byte, error) {
	return marshalUnion(c.MapOfExampleOrRefValues)
}

// ComponentsRequestBodies structure is generated from "#/definitions/Components->requestBodies".
type ComponentsRequestBodies struct {
	MapOfRequestBodyOrRefValues map[string]RequestBodyOrRef `json:"-"` // Key must match pattern: `^[a-zA-Z0-9\.\-_]+$`.
}

// WithMapOfRequestBodyOrRefValues sets MapOfRequestBodyOrRefValues value.
func (c *ComponentsRequestBodies) WithMapOfRequestBodyOrRefValues(val map[string]RequestBodyOrRef) *ComponentsRequestBodies {
	c.MapOfRequestBodyOrRefValues = val
	return c
}

// WithMapOfRequestBodyOrRefValuesItem sets MapOfRequestBodyOrRefValues item value.
func (c *ComponentsRequestBodies) WithMapOfRequestBodyOrRefValuesItem(key string, val RequestBodyOrRef) *ComponentsRequestBodies {
	if c.MapOfRequestBodyOrRefValues == nil {
		c.MapOfRequestBodyOrRefValues = make(map[string]RequestBodyOrRef, 1)
	}

	c.MapOfRequestBodyOrRefValues[key] = val

	return c
}

// UnmarshalJSON decodes JSON.
func (c *ComponentsRequestBodies) UnmarshalJSON(data []byte) error {
	var err error

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexAZAZ09.MatchString(key) {
			matched = true

			if c.MapOfRequestBodyOrRefValues == nil {
				c.MapOfRequestBodyOrRefValues = make(map[string]RequestBodyOrRef, 1)
			}

			var val RequestBodyOrRef

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			c.MapOfRequestBodyOrRefValues[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	return nil
}

// MarshalJSON encodes JSON.
func (c ComponentsRequestBodies) MarshalJSON() ([]byte, error) {
	return marshalUnion(c.MapOfRequestBodyOrRefValues)
}

// ComponentsHeaders structure is generated from "#/definitions/Components->headers".
type ComponentsHeaders struct {
	MapOfHeaderOrRefValues map[string]HeaderOrRef `json:"-"` // Key must match pattern: `^[a-zA-Z0-9\.\-_]+$`.
}

// WithMapOfHeaderOrRefValues sets MapOfHeaderOrRefValues value.
func (c *ComponentsHeaders) WithMapOfHeaderOrRefValues(val map[string]HeaderOrRef) *ComponentsHeaders {
	c.MapOfHeaderOrRefValues = val
	return c
}

// WithMapOfHeaderOrRefValuesItem sets MapOfHeaderOrRefValues item value.
func (c *ComponentsHeaders) WithMapOfHeaderOrRefValuesItem(key string, val HeaderOrRef) *ComponentsHeaders {
	if c.MapOfHeaderOrRefValues == nil {
		c.MapOfHeaderOrRefValues = make(map[string]HeaderOrRef, 1)
	}

	c.MapOfHeaderOrRefValues[key] = val

	return c
}

// UnmarshalJSON decodes JSON.
func (c *ComponentsHeaders) UnmarshalJSON(data []byte) error {
	var err error

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexAZAZ09.MatchString(key) {
			matched = true

			if c.MapOfHeaderOrRefValues == nil {
				c.MapOfHeaderOrRefValues = make(map[string]HeaderOrRef, 1)
			}

			var val HeaderOrRef

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			c.MapOfHeaderOrRefValues[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	return nil
}

// MarshalJSON encodes JSON.
func (c ComponentsHeaders) MarshalJSON() ([]byte, error) {
	return marshalUnion(c.MapOfHeaderOrRefValues)
}

// SecuritySchemeReference structure is generated from "#/definitions/SecuritySchemeReference".
type SecuritySchemeReference struct {
	// Format: uri-reference.
	// Required.
	Ref string `json:"$ref"`
}

// WithRef sets Ref value.
func (s *SecuritySchemeReference) WithRef(val string) *SecuritySchemeReference {
	s.Ref = val
	return s
}

type marshalSecuritySchemeReference SecuritySchemeReference

var knownKeysSecuritySchemeReference = []string{
	"$ref",
}

var requireKeysSecuritySchemeReference = []string{
	"$ref",
}

// UnmarshalJSON decodes JSON.
func (s *SecuritySchemeReference) UnmarshalJSON(data []byte) error {
	var err error

	ms := marshalSecuritySchemeReference(*s)

	err = json.Unmarshal(data, &ms)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysSecuritySchemeReference {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysSecuritySchemeReference {
		delete(rawMap, key)
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in SecuritySchemeReference: %v", offendingKeys)
	}

	*s = SecuritySchemeReference(ms)

	return nil
}

// APIKeySecurityScheme structure is generated from "#/definitions/APIKeySecurityScheme".
type APIKeySecurityScheme struct {
	Name          string                 `json:"name"` // Required.
	In            APIKeySecuritySchemeIn `json:"in"`   // Required.
	Description   *string                `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"` // Key must match pattern: `^x-`.
}

// WithName sets Name value.
func (a *APIKeySecurityScheme) WithName(val string) *APIKeySecurityScheme {
	a.Name = val
	return a
}

// WithIn sets In value.
func (a *APIKeySecurityScheme) WithIn(val APIKeySecuritySchemeIn) *APIKeySecurityScheme {
	a.In = val
	return a
}

// WithDescription sets Description value.
func (a *APIKeySecurityScheme) WithDescription(val string) *APIKeySecurityScheme {
	a.Description = &val
	return a
}

// WithMapOfAnything sets MapOfAnything value.
func (a *APIKeySecurityScheme) WithMapOfAnything(val map[string]interface{}) *APIKeySecurityScheme {
	a.MapOfAnything = val
	return a
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (a *APIKeySecurityScheme) WithMapOfAnythingItem(key string, val interface{}) *APIKeySecurityScheme {
	if a.MapOfAnything == nil {
		a.MapOfAnything = make(map[string]interface{}, 1)
	}

	a.MapOfAnything[key] = val

	return a
}

type marshalAPIKeySecurityScheme APIKeySecurityScheme

var knownKeysAPIKeySecurityScheme = []string{
	"name",
	"in",
	"description",
	"type",
}

var requireKeysAPIKeySecurityScheme = []string{
	"type",
	"name",
	"in",
}

// UnmarshalJSON decodes JSON.
func (a *APIKeySecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	ma := marshalAPIKeySecurityScheme(*a)

	err = json.Unmarshal(data, &ma)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysAPIKeySecurityScheme {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	if v, exists := rawMap["type"]; exists && string(v) != `"apiKey"` {
		return fmt.Errorf(`bad const value for "type" ("apiKey" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range knownKeysAPIKeySecurityScheme {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ma.MapOfAnything == nil {
				ma.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ma.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in APIKeySecurityScheme: %v", offendingKeys)
	}

	*a = APIKeySecurityScheme(ma)

	return nil
}

// constAPIKeySecurityScheme is unconditionally added to JSON.
var constAPIKeySecurityScheme = json.RawMessage(`{"type":"apiKey"}`)

// MarshalJSON encodes JSON.
func (a APIKeySecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(constAPIKeySecurityScheme, marshalAPIKeySecurityScheme(a), a.MapOfAnything)
}

// HTTPSecurityScheme structure is generated from "#/definitions/HTTPSecurityScheme".
type HTTPSecurityScheme struct {
	Scheme        string                 `json:"scheme"` // Required.
	BearerFormat  *string                `json:"bearerFormat,omitempty"`
	Description   *string                `json:"description,omitempty"`
	Bearer        *Bearer                `json:"-"`
	NonBearer     *NonBearer             `json:"-"`
	MapOfAnything map[string]interface{} `json:"-"` // Key must match pattern: `^x-`.
}

// WithScheme sets Scheme value.
func (h *HTTPSecurityScheme) WithScheme(val string) *HTTPSecurityScheme {
	h.Scheme = val
	return h
}

// WithBearerFormat sets BearerFormat value.
func (h *HTTPSecurityScheme) WithBearerFormat(val string) *HTTPSecurityScheme {
	h.BearerFormat = &val
	return h
}

// WithDescription sets Description value.
func (h *HTTPSecurityScheme) WithDescription(val string) *HTTPSecurityScheme {
	h.Description = &val
	return h
}

// WithBearer sets Bearer value.
func (h *HTTPSecurityScheme) WithBearer(val Bearer) *HTTPSecurityScheme {
	h.Bearer = &val
	return h
}

// BearerEns ensures returned Bearer is not nil.
func (h *HTTPSecurityScheme) BearerEns() *Bearer {
	if h.Bearer == nil {
		h.Bearer = new(Bearer)
	}

	return h.Bearer
}

// WithNonBearer sets NonBearer value.
func (h *HTTPSecurityScheme) WithNonBearer(val NonBearer) *HTTPSecurityScheme {
	h.NonBearer = &val
	return h
}

// NonBearerEns ensures returned NonBearer is not nil.
func (h *HTTPSecurityScheme) NonBearerEns() *NonBearer {
	if h.NonBearer == nil {
		h.NonBearer = new(NonBearer)
	}

	return h.NonBearer
}

// WithMapOfAnything sets MapOfAnything value.
func (h *HTTPSecurityScheme) WithMapOfAnything(val map[string]interface{}) *HTTPSecurityScheme {
	h.MapOfAnything = val
	return h
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (h *HTTPSecurityScheme) WithMapOfAnythingItem(key string, val interface{}) *HTTPSecurityScheme {
	if h.MapOfAnything == nil {
		h.MapOfAnything = make(map[string]interface{}, 1)
	}

	h.MapOfAnything[key] = val

	return h
}

type marshalHTTPSecurityScheme HTTPSecurityScheme

var knownKeysHTTPSecurityScheme = []string{
	"scheme",
	"bearerFormat",
	"description",
	"type",
}

var requireKeysHTTPSecurityScheme = []string{
	"scheme",
	"type",
}

// UnmarshalJSON decodes JSON.
func (h *HTTPSecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	mh := marshalHTTPSecurityScheme(*h)

	err = json.Unmarshal(data, &mh)
	if err != nil {
		return err
	}

	oneOfErrors := make(map[string]error, 2)
	oneOfValid := 0

	err = json.Unmarshal(data, &mh.Bearer)
	if err != nil {
		oneOfErrors["Bearer"] = err
		mh.Bearer = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &mh.NonBearer)
	if err != nil {
		oneOfErrors["NonBearer"] = err
		mh.NonBearer = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for HTTPSecurityScheme with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysHTTPSecurityScheme {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	if v, exists := rawMap["type"]; exists && string(v) != `"http"` {
		return fmt.Errorf(`bad const value for "type" ("http" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range knownKeysHTTPSecurityScheme {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if mh.MapOfAnything == nil {
				mh.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mh.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in HTTPSecurityScheme: %v", offendingKeys)
	}

	*h = HTTPSecurityScheme(mh)

	return nil
}

// constHTTPSecurityScheme is unconditionally added to JSON.
var constHTTPSecurityScheme = json.RawMessage(`{"type":"http"}`)

// MarshalJSON encodes JSON.
func (h HTTPSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(constHTTPSecurityScheme, marshalHTTPSecurityScheme(h), h.MapOfAnything, h.Bearer, h.NonBearer)
}

// Bearer structure is generated from "#/definitions/HTTPSecurityScheme/oneOf/0".
//
// Bearer.
//
// Bearer.
type Bearer struct{}

// UnmarshalJSON decodes JSON.
func (b *Bearer) UnmarshalJSON(data []byte) error {
	var err error

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, exists := rawMap["scheme"]; exists && string(v) != `"bearer"` {
		return fmt.Errorf(`bad const value for "scheme" ("bearer" expected, %s received)`, v)
	}

	delete(rawMap, "scheme")

	return nil
}

// constBearer is unconditionally added to JSON.
var constBearer = json.RawMessage(`{"scheme":"bearer"}`)

// MarshalJSON encodes JSON.
func (b Bearer) MarshalJSON() ([]byte, error) {
	return marshalUnion(constBearer)
}

// NonBearer structure is generated from "#/definitions/HTTPSecurityScheme/oneOf/1".
//
// Non Bearer.
//
// Non Bearer.
type NonBearer struct {
	Scheme *interface{} `json:"scheme,omitempty"`
}

// UnmarshalJSON decodes JSON.
func (b *NonBearer) UnmarshalJSON(data []byte) error {
	var err error

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	if v, exists := rawMap["scheme"]; exists && string(v) == `"bearer"` {
		return fmt.Errorf(`bad const value for "scheme" (not "bearer" expected, %s received)`, v)
	}

	if _, exists := rawMap["bearerFormat"]; exists {
		return errors.New(`property "bearerFormat" should not exist`)
	}

	delete(rawMap, "scheme")

	return nil
}

// WithScheme sets Scheme value.
func (n *NonBearer) WithScheme(val interface{}) *NonBearer {
	n.Scheme = &val
	return n
}

// OAuth2SecurityScheme structure is generated from "#/definitions/OAuth2SecurityScheme".
type OAuth2SecurityScheme struct {
	Flows         OAuthFlows             `json:"flows"` // Required.
	Description   *string                `json:"description,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"` // Key must match pattern: `^x-`.
}

// WithFlows sets Flows value.
func (o *OAuth2SecurityScheme) WithFlows(val OAuthFlows) *OAuth2SecurityScheme {
	o.Flows = val
	return o
}

// WithDescription sets Description value.
func (o *OAuth2SecurityScheme) WithDescription(val string) *OAuth2SecurityScheme {
	o.Description = &val
	return o
}

// WithMapOfAnything sets MapOfAnything value.
func (o *OAuth2SecurityScheme) WithMapOfAnything(val map[string]interface{}) *OAuth2SecurityScheme {
	o.MapOfAnything = val
	return o
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (o *OAuth2SecurityScheme) WithMapOfAnythingItem(key string, val interface{}) *OAuth2SecurityScheme {
	if o.MapOfAnything == nil {
		o.MapOfAnything = make(map[string]interface{}, 1)
	}

	o.MapOfAnything[key] = val

	return o
}

type marshalOAuth2SecurityScheme OAuth2SecurityScheme

var knownKeysOAuth2SecurityScheme = []string{
	"flows",
	"description",
	"type",
}

var requireKeysOAuth2SecurityScheme = []string{
	"type",
	"flows",
}

// UnmarshalJSON decodes JSON.
func (o *OAuth2SecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	mo := marshalOAuth2SecurityScheme(*o)

	err = json.Unmarshal(data, &mo)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysOAuth2SecurityScheme {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	if v, exists := rawMap["type"]; exists && string(v) != `"oauth2"` {
		return fmt.Errorf(`bad const value for "type" ("oauth2" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range knownKeysOAuth2SecurityScheme {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if mo.MapOfAnything == nil {
				mo.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mo.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in OAuth2SecurityScheme: %v", offendingKeys)
	}

	*o = OAuth2SecurityScheme(mo)

	return nil
}

// constOAuth2SecurityScheme is unconditionally added to JSON.
var constOAuth2SecurityScheme = json.RawMessage(`{"type":"oauth2"}`)

// MarshalJSON encodes JSON.
func (o OAuth2SecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(constOAuth2SecurityScheme, marshalOAuth2SecurityScheme(o), o.MapOfAnything)
}

// OAuthFlows structure is generated from "#/definitions/OAuthFlows".
type OAuthFlows struct {
	Implicit          *ImplicitOAuthFlow          `json:"implicit,omitempty"`
	Password          *PasswordOAuthFlow          `json:"password,omitempty"`
	ClientCredentials *ClientCredentialsFlow      `json:"clientCredentials,omitempty"`
	AuthorizationCode *AuthorizationCodeOAuthFlow `json:"authorizationCode,omitempty"`
	MapOfAnything     map[string]interface{}      `json:"-"` // Key must match pattern: `^x-`.
}

// WithImplicit sets Implicit value.
func (o *OAuthFlows) WithImplicit(val ImplicitOAuthFlow) *OAuthFlows {
	o.Implicit = &val
	return o
}

// ImplicitEns ensures returned Implicit is not nil.
func (o *OAuthFlows) ImplicitEns() *ImplicitOAuthFlow {
	if o.Implicit == nil {
		o.Implicit = new(ImplicitOAuthFlow)
	}

	return o.Implicit
}

// WithPassword sets Password value.
func (o *OAuthFlows) WithPassword(val PasswordOAuthFlow) *OAuthFlows {
	o.Password = &val
	return o
}

// PasswordEns ensures returned Password is not nil.
func (o *OAuthFlows) PasswordEns() *PasswordOAuthFlow {
	if o.Password == nil {
		o.Password = new(PasswordOAuthFlow)
	}

	return o.Password
}

// WithClientCredentials sets ClientCredentials value.
func (o *OAuthFlows) WithClientCredentials(val ClientCredentialsFlow) *OAuthFlows {
	o.ClientCredentials = &val
	return o
}

// ClientCredentialsEns ensures returned ClientCredentials is not nil.
func (o *OAuthFlows) ClientCredentialsEns() *ClientCredentialsFlow {
	if o.ClientCredentials == nil {
		o.ClientCredentials = new(ClientCredentialsFlow)
	}

	return o.ClientCredentials
}

// WithAuthorizationCode sets AuthorizationCode value.
func (o *OAuthFlows) WithAuthorizationCode(val AuthorizationCodeOAuthFlow) *OAuthFlows {
	o.AuthorizationCode = &val
	return o
}

// AuthorizationCodeEns ensures returned AuthorizationCode is not nil.
func (o *OAuthFlows) AuthorizationCodeEns() *AuthorizationCodeOAuthFlow {
	if o.AuthorizationCode == nil {
		o.AuthorizationCode = new(AuthorizationCodeOAuthFlow)
	}

	return o.AuthorizationCode
}

// WithMapOfAnything sets MapOfAnything value.
func (o *OAuthFlows) WithMapOfAnything(val map[string]interface{}) *OAuthFlows {
	o.MapOfAnything = val
	return o
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (o *OAuthFlows) WithMapOfAnythingItem(key string, val interface{}) *OAuthFlows {
	if o.MapOfAnything == nil {
		o.MapOfAnything = make(map[string]interface{}, 1)
	}

	o.MapOfAnything[key] = val

	return o
}

type marshalOAuthFlows OAuthFlows

var knownKeysOAuthFlows = []string{
	"implicit",
	"password",
	"clientCredentials",
	"authorizationCode",
}

// UnmarshalJSON decodes JSON.
func (o *OAuthFlows) UnmarshalJSON(data []byte) error {
	var err error

	mo := marshalOAuthFlows(*o)

	err = json.Unmarshal(data, &mo)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range knownKeysOAuthFlows {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if mo.MapOfAnything == nil {
				mo.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mo.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in OAuthFlows: %v", offendingKeys)
	}

	*o = OAuthFlows(mo)

	return nil
}

// MarshalJSON encodes JSON.
func (o OAuthFlows) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalOAuthFlows(o), o.MapOfAnything)
}

// ImplicitOAuthFlow structure is generated from "#/definitions/ImplicitOAuthFlow".
type ImplicitOAuthFlow struct {
	// Format: uri-reference.
	// Required.
	AuthorizationURL string                 `json:"authorizationUrl"`
	RefreshURL       *string                `json:"refreshUrl,omitempty"` // Format: uri-reference.
	Scopes           map[string]string      `json:"scopes"`               // Required.
	MapOfAnything    map[string]interface{} `json:"-"`                    // Key must match pattern: `^x-`.
}

// WithAuthorizationURL sets AuthorizationURL value.
func (i *ImplicitOAuthFlow) WithAuthorizationURL(val string) *ImplicitOAuthFlow {
	i.AuthorizationURL = val
	return i
}

// WithRefreshURL sets RefreshURL value.
func (i *ImplicitOAuthFlow) WithRefreshURL(val string) *ImplicitOAuthFlow {
	i.RefreshURL = &val
	return i
}

// WithScopes sets Scopes value.
func (i *ImplicitOAuthFlow) WithScopes(val map[string]string) *ImplicitOAuthFlow {
	i.Scopes = val
	return i
}

// WithScopesItem sets Scopes item value.
func (i *ImplicitOAuthFlow) WithScopesItem(key string, val string) *ImplicitOAuthFlow {
	if i.Scopes == nil {
		i.Scopes = make(map[string]string, 1)
	}

	i.Scopes[key] = val

	return i
}

// WithMapOfAnything sets MapOfAnything value.
func (i *ImplicitOAuthFlow) WithMapOfAnything(val map[string]interface{}) *ImplicitOAuthFlow {
	i.MapOfAnything = val
	return i
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (i *ImplicitOAuthFlow) WithMapOfAnythingItem(key string, val interface{}) *ImplicitOAuthFlow {
	if i.MapOfAnything == nil {
		i.MapOfAnything = make(map[string]interface{}, 1)
	}

	i.MapOfAnything[key] = val

	return i
}

type marshalImplicitOAuthFlow ImplicitOAuthFlow

var knownKeysImplicitOAuthFlow = []string{
	"authorizationUrl",
	"refreshUrl",
	"scopes",
}

var requireKeysImplicitOAuthFlow = []string{
	"authorizationUrl",
	"scopes",
}

// UnmarshalJSON decodes JSON.
func (i *ImplicitOAuthFlow) UnmarshalJSON(data []byte) error {
	var err error

	mi := marshalImplicitOAuthFlow(*i)

	err = json.Unmarshal(data, &mi)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysImplicitOAuthFlow {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysImplicitOAuthFlow {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if mi.MapOfAnything == nil {
				mi.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mi.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in ImplicitOAuthFlow: %v", offendingKeys)
	}

	*i = ImplicitOAuthFlow(mi)

	return nil
}

// MarshalJSON encodes JSON.
func (i ImplicitOAuthFlow) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalImplicitOAuthFlow(i), i.MapOfAnything)
}

// PasswordOAuthFlow structure is generated from "#/definitions/PasswordOAuthFlow".
type PasswordOAuthFlow struct {
	// Format: uri-reference.
	// Required.
	TokenURL      string                 `json:"tokenUrl"`
	RefreshURL    *string                `json:"refreshUrl,omitempty"` // Format: uri-reference.
	Scopes        map[string]string      `json:"scopes,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"` // Key must match pattern: `^x-`.
}

// WithTokenURL sets TokenURL value.
func (p *PasswordOAuthFlow) WithTokenURL(val string) *PasswordOAuthFlow {
	p.TokenURL = val
	return p
}

// WithRefreshURL sets RefreshURL value.
func (p *PasswordOAuthFlow) WithRefreshURL(val string) *PasswordOAuthFlow {
	p.RefreshURL = &val
	return p
}

// WithScopes sets Scopes value.
func (p *PasswordOAuthFlow) WithScopes(val map[string]string) *PasswordOAuthFlow {
	p.Scopes = val
	return p
}

// WithScopesItem sets Scopes item value.
func (p *PasswordOAuthFlow) WithScopesItem(key string, val string) *PasswordOAuthFlow {
	if p.Scopes == nil {
		p.Scopes = make(map[string]string, 1)
	}

	p.Scopes[key] = val

	return p
}

// WithMapOfAnything sets MapOfAnything value.
func (p *PasswordOAuthFlow) WithMapOfAnything(val map[string]interface{}) *PasswordOAuthFlow {
	p.MapOfAnything = val
	return p
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (p *PasswordOAuthFlow) WithMapOfAnythingItem(key string, val interface{}) *PasswordOAuthFlow {
	if p.MapOfAnything == nil {
		p.MapOfAnything = make(map[string]interface{}, 1)
	}

	p.MapOfAnything[key] = val

	return p
}

type marshalPasswordOAuthFlow PasswordOAuthFlow

var knownKeysPasswordOAuthFlow = []string{
	"tokenUrl",
	"refreshUrl",
	"scopes",
}

var requireKeysPasswordOAuthFlow = []string{
	"tokenUrl",
}

// UnmarshalJSON decodes JSON.
func (p *PasswordOAuthFlow) UnmarshalJSON(data []byte) error {
	var err error

	mp := marshalPasswordOAuthFlow(*p)

	err = json.Unmarshal(data, &mp)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysPasswordOAuthFlow {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysPasswordOAuthFlow {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if mp.MapOfAnything == nil {
				mp.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mp.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in PasswordOAuthFlow: %v", offendingKeys)
	}

	*p = PasswordOAuthFlow(mp)

	return nil
}

// MarshalJSON encodes JSON.
func (p PasswordOAuthFlow) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalPasswordOAuthFlow(p), p.MapOfAnything)
}

// ClientCredentialsFlow structure is generated from "#/definitions/ClientCredentialsFlow".
type ClientCredentialsFlow struct {
	// Format: uri-reference.
	// Required.
	TokenURL      string                 `json:"tokenUrl"`
	RefreshURL    *string                `json:"refreshUrl,omitempty"` // Format: uri-reference.
	Scopes        map[string]string      `json:"scopes,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"` // Key must match pattern: `^x-`.
}

// WithTokenURL sets TokenURL value.
func (c *ClientCredentialsFlow) WithTokenURL(val string) *ClientCredentialsFlow {
	c.TokenURL = val
	return c
}

// WithRefreshURL sets RefreshURL value.
func (c *ClientCredentialsFlow) WithRefreshURL(val string) *ClientCredentialsFlow {
	c.RefreshURL = &val
	return c
}

// WithScopes sets Scopes value.
func (c *ClientCredentialsFlow) WithScopes(val map[string]string) *ClientCredentialsFlow {
	c.Scopes = val
	return c
}

// WithScopesItem sets Scopes item value.
func (c *ClientCredentialsFlow) WithScopesItem(key string, val string) *ClientCredentialsFlow {
	if c.Scopes == nil {
		c.Scopes = make(map[string]string, 1)
	}

	c.Scopes[key] = val

	return c
}

// WithMapOfAnything sets MapOfAnything value.
func (c *ClientCredentialsFlow) WithMapOfAnything(val map[string]interface{}) *ClientCredentialsFlow {
	c.MapOfAnything = val
	return c
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (c *ClientCredentialsFlow) WithMapOfAnythingItem(key string, val interface{}) *ClientCredentialsFlow {
	if c.MapOfAnything == nil {
		c.MapOfAnything = make(map[string]interface{}, 1)
	}

	c.MapOfAnything[key] = val

	return c
}

type marshalClientCredentialsFlow ClientCredentialsFlow

var knownKeysClientCredentialsFlow = []string{
	"tokenUrl",
	"refreshUrl",
	"scopes",
}

var requireKeysClientCredentialsFlow = []string{
	"tokenUrl",
}

// UnmarshalJSON decodes JSON.
func (c *ClientCredentialsFlow) UnmarshalJSON(data []byte) error {
	var err error

	mc := marshalClientCredentialsFlow(*c)

	err = json.Unmarshal(data, &mc)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysClientCredentialsFlow {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysClientCredentialsFlow {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if mc.MapOfAnything == nil {
				mc.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mc.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in ClientCredentialsFlow: %v", offendingKeys)
	}

	*c = ClientCredentialsFlow(mc)

	return nil
}

// MarshalJSON encodes JSON.
func (c ClientCredentialsFlow) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalClientCredentialsFlow(c), c.MapOfAnything)
}

// AuthorizationCodeOAuthFlow structure is generated from "#/definitions/AuthorizationCodeOAuthFlow".
type AuthorizationCodeOAuthFlow struct {
	// Format: uri-reference.
	// Required.
	AuthorizationURL string `json:"authorizationUrl"`
	// Format: uri-reference.
	// Required.
	TokenURL      string                 `json:"tokenUrl"`
	RefreshURL    *string                `json:"refreshUrl,omitempty"` // Format: uri-reference.
	Scopes        map[string]string      `json:"scopes,omitempty"`
	MapOfAnything map[string]interface{} `json:"-"` // Key must match pattern: `^x-`.
}

// WithAuthorizationURL sets AuthorizationURL value.
func (a *AuthorizationCodeOAuthFlow) WithAuthorizationURL(val string) *AuthorizationCodeOAuthFlow {
	a.AuthorizationURL = val
	return a
}

// WithTokenURL sets TokenURL value.
func (a *AuthorizationCodeOAuthFlow) WithTokenURL(val string) *AuthorizationCodeOAuthFlow {
	a.TokenURL = val
	return a
}

// WithRefreshURL sets RefreshURL value.
func (a *AuthorizationCodeOAuthFlow) WithRefreshURL(val string) *AuthorizationCodeOAuthFlow {
	a.RefreshURL = &val
	return a
}

// WithScopes sets Scopes value.
func (a *AuthorizationCodeOAuthFlow) WithScopes(val map[string]string) *AuthorizationCodeOAuthFlow {
	a.Scopes = val
	return a
}

// WithScopesItem sets Scopes item value.
func (a *AuthorizationCodeOAuthFlow) WithScopesItem(key string, val string) *AuthorizationCodeOAuthFlow {
	if a.Scopes == nil {
		a.Scopes = make(map[string]string, 1)
	}

	a.Scopes[key] = val

	return a
}

// WithMapOfAnything sets MapOfAnything value.
func (a *AuthorizationCodeOAuthFlow) WithMapOfAnything(val map[string]interface{}) *AuthorizationCodeOAuthFlow {
	a.MapOfAnything = val
	return a
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (a *AuthorizationCodeOAuthFlow) WithMapOfAnythingItem(key string, val interface{}) *AuthorizationCodeOAuthFlow {
	if a.MapOfAnything == nil {
		a.MapOfAnything = make(map[string]interface{}, 1)
	}

	a.MapOfAnything[key] = val

	return a
}

type marshalAuthorizationCodeOAuthFlow AuthorizationCodeOAuthFlow

var knownKeysAuthorizationCodeOAuthFlow = []string{
	"authorizationUrl",
	"tokenUrl",
	"refreshUrl",
	"scopes",
}

var requireKeysAuthorizationCodeOAuthFlow = []string{
	"authorizationUrl",
	"tokenUrl",
}

// UnmarshalJSON decodes JSON.
func (a *AuthorizationCodeOAuthFlow) UnmarshalJSON(data []byte) error {
	var err error

	ma := marshalAuthorizationCodeOAuthFlow(*a)

	err = json.Unmarshal(data, &ma)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysAuthorizationCodeOAuthFlow {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	for _, key := range knownKeysAuthorizationCodeOAuthFlow {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if ma.MapOfAnything == nil {
				ma.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			ma.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in AuthorizationCodeOAuthFlow: %v", offendingKeys)
	}

	*a = AuthorizationCodeOAuthFlow(ma)

	return nil
}

// MarshalJSON encodes JSON.
func (a AuthorizationCodeOAuthFlow) MarshalJSON() ([]byte, error) {
	return marshalUnion(marshalAuthorizationCodeOAuthFlow(a), a.MapOfAnything)
}

// OpenIDConnectSecurityScheme structure is generated from "#/definitions/OpenIdConnectSecurityScheme".
type OpenIDConnectSecurityScheme struct {
	// Format: uri-reference.
	// Required.
	OpenIDConnectURL string                 `json:"openIdConnectUrl"`
	Description      *string                `json:"description,omitempty"`
	MapOfAnything    map[string]interface{} `json:"-"` // Key must match pattern: `^x-`.
}

// WithOpenIDConnectURL sets OpenIDConnectURL value.
func (o *OpenIDConnectSecurityScheme) WithOpenIDConnectURL(val string) *OpenIDConnectSecurityScheme {
	o.OpenIDConnectURL = val
	return o
}

// WithDescription sets Description value.
func (o *OpenIDConnectSecurityScheme) WithDescription(val string) *OpenIDConnectSecurityScheme {
	o.Description = &val
	return o
}

// WithMapOfAnything sets MapOfAnything value.
func (o *OpenIDConnectSecurityScheme) WithMapOfAnything(val map[string]interface{}) *OpenIDConnectSecurityScheme {
	o.MapOfAnything = val
	return o
}

// WithMapOfAnythingItem sets MapOfAnything item value.
func (o *OpenIDConnectSecurityScheme) WithMapOfAnythingItem(key string, val interface{}) *OpenIDConnectSecurityScheme {
	if o.MapOfAnything == nil {
		o.MapOfAnything = make(map[string]interface{}, 1)
	}

	o.MapOfAnything[key] = val

	return o
}

type marshalOpenIDConnectSecurityScheme OpenIDConnectSecurityScheme

var knownKeysOpenIDConnectSecurityScheme = []string{
	"openIdConnectUrl",
	"description",
	"type",
}

var requireKeysOpenIDConnectSecurityScheme = []string{
	"type",
	"openIdConnectUrl",
}

// UnmarshalJSON decodes JSON.
func (o *OpenIDConnectSecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	mo := marshalOpenIDConnectSecurityScheme(*o)

	err = json.Unmarshal(data, &mo)
	if err != nil {
		return err
	}

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for _, key := range requireKeysOpenIDConnectSecurityScheme {
		if _, found := rawMap[key]; !found {
			return errors.New("required key missing: " + key)
		}
	}

	if v, exists := rawMap["type"]; exists && string(v) != `"openIdConnect"` {
		return fmt.Errorf(`bad const value for "type" ("openIdConnect" expected, %s received)`, v)
	}

	delete(rawMap, "type")

	for _, key := range knownKeysOpenIDConnectSecurityScheme {
		delete(rawMap, key)
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexX.MatchString(key) {
			matched = true

			if mo.MapOfAnything == nil {
				mo.MapOfAnything = make(map[string]interface{}, 1)
			}

			var val interface{}

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			mo.MapOfAnything[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	if len(rawMap) != 0 {
		offendingKeys := make([]string, 0, len(rawMap))

		for key := range rawMap {
			offendingKeys = append(offendingKeys, key)
		}

		return fmt.Errorf("additional properties not allowed in OpenIDConnectSecurityScheme: %v", offendingKeys)
	}

	*o = OpenIDConnectSecurityScheme(mo)

	return nil
}

// constOpenIDConnectSecurityScheme is unconditionally added to JSON.
var constOpenIDConnectSecurityScheme = json.RawMessage(`{"type":"openIdConnect"}`)

// MarshalJSON encodes JSON.
func (o OpenIDConnectSecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(constOpenIDConnectSecurityScheme, marshalOpenIDConnectSecurityScheme(o), o.MapOfAnything)
}

// SecurityScheme structure is generated from "#/definitions/SecurityScheme".
type SecurityScheme struct {
	APIKeySecurityScheme        *APIKeySecurityScheme        `json:"-"`
	HTTPSecurityScheme          *HTTPSecurityScheme          `json:"-"`
	OAuth2SecurityScheme        *OAuth2SecurityScheme        `json:"-"`
	OpenIDConnectSecurityScheme *OpenIDConnectSecurityScheme `json:"-"`
}

// WithAPIKeySecurityScheme sets APIKeySecurityScheme value.
func (s *SecurityScheme) WithAPIKeySecurityScheme(val APIKeySecurityScheme) *SecurityScheme {
	s.APIKeySecurityScheme = &val
	return s
}

// APIKeySecuritySchemeEns ensures returned APIKeySecurityScheme is not nil.
func (s *SecurityScheme) APIKeySecuritySchemeEns() *APIKeySecurityScheme {
	if s.APIKeySecurityScheme == nil {
		s.APIKeySecurityScheme = new(APIKeySecurityScheme)
	}

	return s.APIKeySecurityScheme
}

// WithHTTPSecurityScheme sets HTTPSecurityScheme value.
func (s *SecurityScheme) WithHTTPSecurityScheme(val HTTPSecurityScheme) *SecurityScheme {
	s.HTTPSecurityScheme = &val
	return s
}

// HTTPSecuritySchemeEns ensures returned HTTPSecurityScheme is not nil.
func (s *SecurityScheme) HTTPSecuritySchemeEns() *HTTPSecurityScheme {
	if s.HTTPSecurityScheme == nil {
		s.HTTPSecurityScheme = new(HTTPSecurityScheme)
	}

	return s.HTTPSecurityScheme
}

// WithOAuth2SecurityScheme sets OAuth2SecurityScheme value.
func (s *SecurityScheme) WithOAuth2SecurityScheme(val OAuth2SecurityScheme) *SecurityScheme {
	s.OAuth2SecurityScheme = &val
	return s
}

// OAuth2SecuritySchemeEns ensures returned OAuth2SecurityScheme is not nil.
func (s *SecurityScheme) OAuth2SecuritySchemeEns() *OAuth2SecurityScheme {
	if s.OAuth2SecurityScheme == nil {
		s.OAuth2SecurityScheme = new(OAuth2SecurityScheme)
	}

	return s.OAuth2SecurityScheme
}

// WithOpenIDConnectSecurityScheme sets OpenIDConnectSecurityScheme value.
func (s *SecurityScheme) WithOpenIDConnectSecurityScheme(val OpenIDConnectSecurityScheme) *SecurityScheme {
	s.OpenIDConnectSecurityScheme = &val
	return s
}

// OpenIDConnectSecuritySchemeEns ensures returned OpenIDConnectSecurityScheme is not nil.
func (s *SecurityScheme) OpenIDConnectSecuritySchemeEns() *OpenIDConnectSecurityScheme {
	if s.OpenIDConnectSecurityScheme == nil {
		s.OpenIDConnectSecurityScheme = new(OpenIDConnectSecurityScheme)
	}

	return s.OpenIDConnectSecurityScheme
}

// UnmarshalJSON decodes JSON.
func (s *SecurityScheme) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 4)
	oneOfValid := 0

	err = json.Unmarshal(data, &s.APIKeySecurityScheme)
	if err != nil {
		oneOfErrors["APIKeySecurityScheme"] = err
		s.APIKeySecurityScheme = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &s.HTTPSecurityScheme)
	if err != nil {
		oneOfErrors["HTTPSecurityScheme"] = err
		s.HTTPSecurityScheme = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &s.OAuth2SecurityScheme)
	if err != nil {
		oneOfErrors["OAuth2SecurityScheme"] = err
		s.OAuth2SecurityScheme = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &s.OpenIDConnectSecurityScheme)
	if err != nil {
		oneOfErrors["OpenIDConnectSecurityScheme"] = err
		s.OpenIDConnectSecurityScheme = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for SecurityScheme with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (s SecurityScheme) MarshalJSON() ([]byte, error) {
	return marshalUnion(s.APIKeySecurityScheme, s.HTTPSecurityScheme, s.OAuth2SecurityScheme, s.OpenIDConnectSecurityScheme)
}

// SecuritySchemeOrRef structure is generated from "#/definitions/SecuritySchemeOrRef".
type SecuritySchemeOrRef struct {
	SecuritySchemeReference *SecuritySchemeReference `json:"-"`
	SecurityScheme          *SecurityScheme          `json:"-"`
}

// WithSecuritySchemeReference sets SecuritySchemeReference value.
func (s *SecuritySchemeOrRef) WithSecuritySchemeReference(val SecuritySchemeReference) *SecuritySchemeOrRef {
	s.SecuritySchemeReference = &val
	return s
}

// SecuritySchemeReferenceEns ensures returned SecuritySchemeReference is not nil.
func (s *SecuritySchemeOrRef) SecuritySchemeReferenceEns() *SecuritySchemeReference {
	if s.SecuritySchemeReference == nil {
		s.SecuritySchemeReference = new(SecuritySchemeReference)
	}

	return s.SecuritySchemeReference
}

// WithSecurityScheme sets SecurityScheme value.
func (s *SecuritySchemeOrRef) WithSecurityScheme(val SecurityScheme) *SecuritySchemeOrRef {
	s.SecurityScheme = &val
	return s
}

// SecuritySchemeEns ensures returned SecurityScheme is not nil.
func (s *SecuritySchemeOrRef) SecuritySchemeEns() *SecurityScheme {
	if s.SecurityScheme == nil {
		s.SecurityScheme = new(SecurityScheme)
	}

	return s.SecurityScheme
}

// UnmarshalJSON decodes JSON.
func (s *SecuritySchemeOrRef) UnmarshalJSON(data []byte) error {
	var err error

	oneOfErrors := make(map[string]error, 2)
	oneOfValid := 0

	err = json.Unmarshal(data, &s.SecuritySchemeReference)
	if err != nil {
		oneOfErrors["SecuritySchemeReference"] = err
		s.SecuritySchemeReference = nil
	} else {
		oneOfValid++
	}

	err = json.Unmarshal(data, &s.SecurityScheme)
	if err != nil {
		oneOfErrors["SecurityScheme"] = err
		s.SecurityScheme = nil
	} else {
		oneOfValid++
	}

	if oneOfValid != 1 {
		return fmt.Errorf("oneOf constraint failed for SecuritySchemeOrRef with %d valid results: %v", oneOfValid, oneOfErrors)
	}

	return nil
}

// MarshalJSON encodes JSON.
func (s SecuritySchemeOrRef) MarshalJSON() ([]byte, error) {
	return marshalUnion(s.SecuritySchemeReference, s.SecurityScheme)
}

// ComponentsSecuritySchemes structure is generated from "#/definitions/Components->securitySchemes".
type ComponentsSecuritySchemes struct {
	MapOfSecuritySchemeOrRefValues map[string]SecuritySchemeOrRef `json:"-"` // Key must match pattern: `^[a-zA-Z0-9\.\-_]+$`.
}

// WithMapOfSecuritySchemeOrRefValues sets MapOfSecuritySchemeOrRefValues value.
func (c *ComponentsSecuritySchemes) WithMapOfSecuritySchemeOrRefValues(val map[string]SecuritySchemeOrRef) *ComponentsSecuritySchemes {
	c.MapOfSecuritySchemeOrRefValues = val
	return c
}

// WithMapOfSecuritySchemeOrRefValuesItem sets MapOfSecuritySchemeOrRefValues item value.
func (c *ComponentsSecuritySchemes) WithMapOfSecuritySchemeOrRefValuesItem(key string, val SecuritySchemeOrRef) *ComponentsSecuritySchemes {
	if c.MapOfSecuritySchemeOrRefValues == nil {
		c.MapOfSecuritySchemeOrRefValues = make(map[string]SecuritySchemeOrRef, 1)
	}

	c.MapOfSecuritySchemeOrRefValues[key] = val

	return c
}

// UnmarshalJSON decodes JSON.
func (c *ComponentsSecuritySchemes) UnmarshalJSON(data []byte) error {
	var err error

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexAZAZ09.MatchString(key) {
			matched = true

			if c.MapOfSecuritySchemeOrRefValues == nil {
				c.MapOfSecuritySchemeOrRefValues = make(map[string]SecuritySchemeOrRef, 1)
			}

			var val SecuritySchemeOrRef

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			c.MapOfSecuritySchemeOrRefValues[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	return nil
}

// MarshalJSON encodes JSON.
func (c ComponentsSecuritySchemes) MarshalJSON() ([]byte, error) {
	return marshalUnion(c.MapOfSecuritySchemeOrRefValues)
}

// ComponentsLinks structure is generated from "#/definitions/Components->links".
type ComponentsLinks struct {
	MapOfLinkOrRefValues map[string]LinkOrRef `json:"-"` // Key must match pattern: `^[a-zA-Z0-9\.\-_]+$`.
}

// WithMapOfLinkOrRefValues sets MapOfLinkOrRefValues value.
func (c *ComponentsLinks) WithMapOfLinkOrRefValues(val map[string]LinkOrRef) *ComponentsLinks {
	c.MapOfLinkOrRefValues = val
	return c
}

// WithMapOfLinkOrRefValuesItem sets MapOfLinkOrRefValues item value.
func (c *ComponentsLinks) WithMapOfLinkOrRefValuesItem(key string, val LinkOrRef) *ComponentsLinks {
	if c.MapOfLinkOrRefValues == nil {
		c.MapOfLinkOrRefValues = make(map[string]LinkOrRef, 1)
	}

	c.MapOfLinkOrRefValues[key] = val

	return c
}

// UnmarshalJSON decodes JSON.
func (c *ComponentsLinks) UnmarshalJSON(data []byte) error {
	var err error

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexAZAZ09.MatchString(key) {
			matched = true

			if c.MapOfLinkOrRefValues == nil {
				c.MapOfLinkOrRefValues = make(map[string]LinkOrRef, 1)
			}

			var val LinkOrRef

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			c.MapOfLinkOrRefValues[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	return nil
}

// MarshalJSON encodes JSON.
func (c ComponentsLinks) MarshalJSON() ([]byte, error) {
	return marshalUnion(c.MapOfLinkOrRefValues)
}

// ComponentsCallbacks structure is generated from "#/definitions/Components->callbacks".
type ComponentsCallbacks struct {
	MapOfCallbackOrRefValues map[string]CallbackOrRef `json:"-"` // Key must match pattern: `^[a-zA-Z0-9\.\-_]+$`.
}

// WithMapOfCallbackOrRefValues sets MapOfCallbackOrRefValues value.
func (c *ComponentsCallbacks) WithMapOfCallbackOrRefValues(val map[string]CallbackOrRef) *ComponentsCallbacks {
	c.MapOfCallbackOrRefValues = val
	return c
}

// WithMapOfCallbackOrRefValuesItem sets MapOfCallbackOrRefValues item value.
func (c *ComponentsCallbacks) WithMapOfCallbackOrRefValuesItem(key string, val CallbackOrRef) *ComponentsCallbacks {
	if c.MapOfCallbackOrRefValues == nil {
		c.MapOfCallbackOrRefValues = make(map[string]CallbackOrRef, 1)
	}

	c.MapOfCallbackOrRefValues[key] = val

	return c
}

// UnmarshalJSON decodes JSON.
func (c *ComponentsCallbacks) UnmarshalJSON(data []byte) error {
	var err error

	var rawMap map[string]json.RawMessage

	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		rawMap = nil
	}

	for key, rawValue := range rawMap {
		matched := false

		if regexAZAZ09.MatchString(key) {
			matched = true

			if c.MapOfCallbackOrRefValues == nil {
				c.MapOfCallbackOrRefValues = make(map[string]CallbackOrRef, 1)
			}

			var val CallbackOrRef

			err = json.Unmarshal(rawValue, &val)
			if err != nil {
				return err
			}

			c.MapOfCallbackOrRefValues[key] = val
		}

		if matched {
			delete(rawMap, key)
		}
	}

	return nil
}

// MarshalJSON encodes JSON.
func (c ComponentsCallbacks) MarshalJSON() ([]byte, error) {
	return marshalUnion(c.MapOfCallbackOrRefValues)
}

// ParameterIn is an enum type.
type ParameterIn string

// ParameterIn values enumeration.
const (
	ParameterInPath   = ParameterIn("path")
	ParameterInQuery  = ParameterIn("query")
	ParameterInHeader = ParameterIn("header")
	ParameterInCookie = ParameterIn("cookie")
)

// MarshalJSON encodes JSON.
func (i ParameterIn) MarshalJSON() ([]byte, error) {
	switch i {
	case ParameterInPath:
	case ParameterInQuery:
	case ParameterInHeader:
	case ParameterInCookie:

	default:
		return nil, fmt.Errorf("unexpected ParameterIn value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *ParameterIn) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := ParameterIn(ii)

	switch v {
	case ParameterInPath:
	case ParameterInQuery:
	case ParameterInHeader:
	case ParameterInCookie:

	default:
		return fmt.Errorf("unexpected ParameterIn value: %v", v)
	}

	*i = v

	return nil
}

// SchemaType is an enum type.
type SchemaType string

// SchemaType values enumeration.
const (
	SchemaTypeArray   = SchemaType("array")
	SchemaTypeBoolean = SchemaType("boolean")
	SchemaTypeInteger = SchemaType("integer")
	SchemaTypeNumber  = SchemaType("number")
	SchemaTypeObject  = SchemaType("object")
	SchemaTypeString  = SchemaType("string")
)

// MarshalJSON encodes JSON.
func (i SchemaType) MarshalJSON() ([]byte, error) {
	switch i {
	case SchemaTypeArray:
	case SchemaTypeBoolean:
	case SchemaTypeInteger:
	case SchemaTypeNumber:
	case SchemaTypeObject:
	case SchemaTypeString:

	default:
		return nil, fmt.Errorf("unexpected SchemaType value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *SchemaType) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := SchemaType(ii)

	switch v {
	case SchemaTypeArray:
	case SchemaTypeBoolean:
	case SchemaTypeInteger:
	case SchemaTypeNumber:
	case SchemaTypeObject:
	case SchemaTypeString:

	default:
		return fmt.Errorf("unexpected SchemaType value: %v", v)
	}

	*i = v

	return nil
}

// EncodingStyle is an enum type.
type EncodingStyle string

// EncodingStyle values enumeration.
const (
	EncodingStyleForm           = EncodingStyle("form")
	EncodingStyleSpaceDelimited = EncodingStyle("spaceDelimited")
	EncodingStylePipeDelimited  = EncodingStyle("pipeDelimited")
	EncodingStyleDeepObject     = EncodingStyle("deepObject")
)

// MarshalJSON encodes JSON.
func (i EncodingStyle) MarshalJSON() ([]byte, error) {
	switch i {
	case EncodingStyleForm:
	case EncodingStyleSpaceDelimited:
	case EncodingStylePipeDelimited:
	case EncodingStyleDeepObject:

	default:
		return nil, fmt.Errorf("unexpected EncodingStyle value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *EncodingStyle) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := EncodingStyle(ii)

	switch v {
	case EncodingStyleForm:
	case EncodingStyleSpaceDelimited:
	case EncodingStylePipeDelimited:
	case EncodingStyleDeepObject:

	default:
		return fmt.Errorf("unexpected EncodingStyle value: %v", v)
	}

	*i = v

	return nil
}

// PathParameterStyle is an enum type.
type PathParameterStyle string

// PathParameterStyle values enumeration.
const (
	PathParameterStyleMatrix = PathParameterStyle("matrix")
	PathParameterStyleLabel  = PathParameterStyle("label")
	PathParameterStyleSimple = PathParameterStyle("simple")
)

// MarshalJSON encodes JSON.
func (i PathParameterStyle) MarshalJSON() ([]byte, error) {
	switch i {
	case PathParameterStyleMatrix:
	case PathParameterStyleLabel:
	case PathParameterStyleSimple:

	default:
		return nil, fmt.Errorf("unexpected PathParameterStyle value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *PathParameterStyle) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := PathParameterStyle(ii)

	switch v {
	case PathParameterStyleMatrix:
	case PathParameterStyleLabel:
	case PathParameterStyleSimple:

	default:
		return fmt.Errorf("unexpected PathParameterStyle value: %v", v)
	}

	*i = v

	return nil
}

// QueryParameterStyle is an enum type.
type QueryParameterStyle string

// QueryParameterStyle values enumeration.
const (
	QueryParameterStyleForm           = QueryParameterStyle("form")
	QueryParameterStyleSpaceDelimited = QueryParameterStyle("spaceDelimited")
	QueryParameterStylePipeDelimited  = QueryParameterStyle("pipeDelimited")
	QueryParameterStyleDeepObject     = QueryParameterStyle("deepObject")
)

// MarshalJSON encodes JSON.
func (i QueryParameterStyle) MarshalJSON() ([]byte, error) {
	switch i {
	case QueryParameterStyleForm:
	case QueryParameterStyleSpaceDelimited:
	case QueryParameterStylePipeDelimited:
	case QueryParameterStyleDeepObject:

	default:
		return nil, fmt.Errorf("unexpected QueryParameterStyle value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *QueryParameterStyle) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := QueryParameterStyle(ii)

	switch v {
	case QueryParameterStyleForm:
	case QueryParameterStyleSpaceDelimited:
	case QueryParameterStylePipeDelimited:
	case QueryParameterStyleDeepObject:

	default:
		return fmt.Errorf("unexpected QueryParameterStyle value: %v", v)
	}

	*i = v

	return nil
}

// APIKeySecuritySchemeIn is an enum type.
type APIKeySecuritySchemeIn string

// APIKeySecuritySchemeIn values enumeration.
const (
	APIKeySecuritySchemeInHeader = APIKeySecuritySchemeIn("header")
	APIKeySecuritySchemeInQuery  = APIKeySecuritySchemeIn("query")
	APIKeySecuritySchemeInCookie = APIKeySecuritySchemeIn("cookie")
)

// MarshalJSON encodes JSON.
func (i APIKeySecuritySchemeIn) MarshalJSON() ([]byte, error) {
	switch i {
	case APIKeySecuritySchemeInHeader:
	case APIKeySecuritySchemeInQuery:
	case APIKeySecuritySchemeInCookie:

	default:
		return nil, fmt.Errorf("unexpected APIKeySecuritySchemeIn value: %v", i)
	}

	return json.Marshal(string(i))
}

// UnmarshalJSON decodes JSON.
func (i *APIKeySecuritySchemeIn) UnmarshalJSON(data []byte) error {
	var ii string

	err := json.Unmarshal(data, &ii)
	if err != nil {
		return err
	}

	v := APIKeySecuritySchemeIn(ii)

	switch v {
	case APIKeySecuritySchemeInHeader:
	case APIKeySecuritySchemeInQuery:
	case APIKeySecuritySchemeInCookie:

	default:
		return fmt.Errorf("unexpected APIKeySecuritySchemeIn value: %v", v)
	}

	*i = v

	return nil
}

func marshalUnion(maps ...interface{}) ([]byte, error) {
	result := []byte("{")
	isObject := true

	for _, m := range maps {
		j, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}

		if string(j) == "{}" {
			continue
		}

		if string(j) == "null" {
			continue
		}

		if j[0] != '{' {
			if len(result) == 1 && (isObject || bytes.Equal(result, j)) {
				result = j
				isObject = false

				continue
			}

			return nil, errors.New("failed to union map: object expected, " + string(j) + " received")
		}

		if !isObject {
			return nil, errors.New("failed to union " + string(result) + " and " + string(j))
		}

		if len(result) > 1 {
			result[len(result)-1] = ','
		}

		result = append(result, j[1:]...)
	}

	// Close empty result.
	if isObject && len(result) == 1 {
		result = append(result, '}')
	}

	return result, nil
}

// Regular expressions for pattern properties.
var (
	regexX                                     = regexp.MustCompile("^x-")
	regexGetPutPostDeleteOptionsHeadPatchTrace = regexp.MustCompile("^(get|put|post|delete|options|head|patch|trace)$")
	regex15D2XX                                = regexp.MustCompile(`^[1-5](?:\d{2}|XX)$`)
	regex                                      = regexp.MustCompile(`^\/`)
	regexAZAZ09                                = regexp.MustCompile(`^[a-zA-Z0-9\.\-_]+$`)
)
