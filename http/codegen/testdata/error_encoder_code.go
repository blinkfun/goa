package testdata

var PrimitiveErrorResponseEncoderCode = `// EncodeMethodPrimitiveErrorResponseError returns an encoder for errors
// returned by the MethodPrimitiveErrorResponse ServicePrimitiveErrorResponse
// endpoint.
func EncodeMethodPrimitiveErrorResponseError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "bad_request":
			res := v.(serviceprimitiveerrorresponse.BadRequest)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewMethodPrimitiveErrorResponseBadRequestResponseBody(res)
			}
			w.Header().Set("x-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "internal_error":
			res := v.(serviceprimitiveerrorresponse.InternalError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewMethodPrimitiveErrorResponseInternalErrorResponseBody(res)
			}
			w.Header().Set("x-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}
`

var PrimitiveErrorInResponseHeaderEncoderCode = `// EncodeMethodPrimitiveErrorInResponseHeaderError returns an encoder for
// errors returned by the MethodPrimitiveErrorInResponseHeader
// ServicePrimitiveErrorInResponseHeader endpoint.
func EncodeMethodPrimitiveErrorInResponseHeaderError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "bad_request":
			res := v.(serviceprimitiveerrorinresponseheader.BadRequest)
			{
				val := string(res)
				string_s := val
				w.Header().Set("String", string_s)
			}
			w.Header().Set("x-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return nil
		case "internal_error":
			res := v.(serviceprimitiveerrorinresponseheader.InternalError)
			{
				val := int(res)
				int_s := strconv.Itoa(val)
				w.Header().Set("Int", int_s)
			}
			w.Header().Set("x-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return nil
		default:
			return encodeError(ctx, w, v)
		}
	}
}
`

var DefaultErrorResponseEncoderCode = `// EncodeMethodDefaultErrorResponseError returns an encoder for errors returned
// by the MethodDefaultErrorResponse ServiceDefaultErrorResponse endpoint.
func EncodeMethodDefaultErrorResponseError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "bad_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewMethodDefaultErrorResponseBadRequestResponseBody(res)
			}
			w.Header().Set("x-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}
`

var DefaultErrorResponseWithContentTypeEncoderCode = `// EncodeMethodDefaultErrorResponseError returns an encoder for errors returned
// by the MethodDefaultErrorResponse ServiceDefaultErrorResponse endpoint.
func EncodeMethodDefaultErrorResponseError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "bad_request":
			res := v.(*goa.ServiceError)
			ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/xml")
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewMethodDefaultErrorResponseBadRequestResponseBody(res)
			}
			w.Header().Set("x-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}
`

var ServiceErrorResponseEncoderCode = `// EncodeMethodServiceErrorResponseError returns an encoder for errors returned
// by the MethodServiceErrorResponse ServiceServiceErrorResponse endpoint.
func EncodeMethodServiceErrorResponseError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "internal_error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewMethodServiceErrorResponseInternalErrorResponseBody(res)
			}
			w.Header().Set("x-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "bad_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewMethodServiceErrorResponseBadRequestResponseBody(res)
			}
			w.Header().Set("x-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}
`

var ServiceErrorResponseWithContentTypeEncoderCode = `// EncodeMethodServiceErrorResponseError returns an encoder for errors returned
// by the MethodServiceErrorResponse ServiceServiceErrorResponse endpoint.
func EncodeMethodServiceErrorResponseError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "internal_error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewMethodServiceErrorResponseInternalErrorResponseBody(res)
			}
			w.Header().Set("x-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "bad_request":
			res := v.(*goa.ServiceError)
			ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/xml")
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewMethodServiceErrorResponseBadRequestResponseBody(res)
			}
			w.Header().Set("x-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}
`

var NoBodyErrorResponseEncoderCode = `// EncodeMethodServiceErrorResponseError returns an encoder for errors returned
// by the MethodServiceErrorResponse ServiceNoBodyErrorResponse endpoint.
func EncodeMethodServiceErrorResponseError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "bad_request":
			res := v.(*servicenobodyerrorresponse.StringError)
			if res.Header != nil {
				w.Header().Set("Header", *res.Header)
			}
			w.Header().Set("x-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return nil
		default:
			return encodeError(ctx, w, v)
		}
	}
}
`

var NoBodyErrorResponseWithContentTypeEncoderCode = `// EncodeMethodServiceErrorResponseError returns an encoder for errors returned
// by the MethodServiceErrorResponse ServiceNoBodyErrorResponse endpoint.
func EncodeMethodServiceErrorResponseError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "bad_request":
			res := v.(*servicenobodyerrorresponse.StringError)
			ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/xml")
			if res.Header != nil {
				w.Header().Set("Header", *res.Header)
			}
			w.Header().Set("x-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return nil
		default:
			return encodeError(ctx, w, v)
		}
	}
}
`

var EmptyErrorResponseBodyEncoderCode = `// EncodeMethodEmptyErrorResponseBodyError returns an encoder for errors
// returned by the MethodEmptyErrorResponseBody ServiceEmptyErrorResponseBody
// endpoint.
func EncodeMethodEmptyErrorResponseBodyError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "internal_error":
			res := v.(*goa.ServiceError)
			w.Header().Set("Error-Name", res.Name)
			w.Header().Set("X-Attribute-Id", res.ID)
			w.Header().Set("X-Attribute-Message", res.Message)
			w.Header().Set("x-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return nil
		case "not_found":
			res := v.(serviceemptyerrorresponsebody.NotFound)
			{
				val := string(res)
				inHeaders := val
				w.Header().Set("In-Header", inHeaders)
			}
			w.Header().Set("x-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return nil
		default:
			return encodeError(ctx, w, v)
		}
	}
}
`

var EmptyCustomErrorResponseBodyEncoderCode = `// EncodeMethodEmptyCustomErrorResponseBodyError returns an encoder for errors
// returned by the MethodEmptyCustomErrorResponseBody
// ServiceEmptyCustomErrorResponseBody endpoint.
func EncodeMethodEmptyCustomErrorResponseBodyError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "internal_error":
			res := v.(*serviceemptycustomerrorresponsebody.Error)
			w.Header().Set("x-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return nil
		default:
			return encodeError(ctx, w, v)
		}
	}
}
`
