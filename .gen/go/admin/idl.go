// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by thriftrw v1.11.0. DO NOT EDIT.
// @generated

package admin

import (
	"github.com/uber/cadence/.gen/go/shared"
	"go.uber.org/thriftrw/thriftreflect"
)

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "admin",
	Package:  "github.com/uber/cadence/.gen/go/admin",
	FilePath: "admin.thrift",
	SHA1:     "bb61fbdf93f8f77bd9957ecfcacf5508cedfe81d",
	Includes: []*thriftreflect.ThriftModule{
		shared.ThriftModule,
	},
	Raw: rawIDL,
}

const rawIDL = "// Copyright (c) 2017 Uber Technologies, Inc.\n//\n// Permission is hereby granted, free of charge, to any person obtaining a copy\n// of this software and associated documentation files (the \"Software\"), to deal\n// in the Software without restriction, including without limitation the rights\n// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell\n// copies of the Software, and to permit persons to whom the Software is\n// furnished to do so, subject to the following conditions:\n//\n// The above copyright notice and this permission notice shall be included in\n// all copies or substantial portions of the Software.\n//\n// THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR\n// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,\n// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE\n// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER\n// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,\n// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN\n// THE SOFTWARE.\n\nnamespace java com.uber.cadence.admin\n\ninclude \"shared.thrift\"\n\n/**\n* AdminService provides advanced APIs for debugging and analysis with admin privillege\n**/\nservice AdminService {\n  /**\n  * DescribeWorkflowExecution returns information about the internal states of workflow execution.\n  **/\n  DescribeWorkflowExecutionResponse DescribeWorkflowExecution(1: DescribeWorkflowExecutionRequest request)\n    throws (\n      1: shared.BadRequestError         badRequestError,\n      2: shared.InternalServiceError    internalServiceError,\n      3: shared.EntityNotExistsError    entityNotExistError,\n      4: shared.AccessDeniedError       accessDeniedError,\n    )\n\n  /**\n    * DescribeHistoryHost returns information about the internal states of a history host\n    **/\n    shared.DescribeHistoryHostResponse DescribeHistoryHost(1: shared.DescribeHistoryHostRequest request)\n      throws (\n        1: shared.BadRequestError       badRequestError,\n        2: shared.InternalServiceError  internalServiceError,\n        3: shared.AccessDeniedError     accessDeniedError,\n      )\n}\n\nstruct DescribeWorkflowExecutionRequest {\n  10: optional string                       domain\n  20: optional shared.WorkflowExecution     execution\n}\n\nstruct DescribeWorkflowExecutionResponse{\n  10: optional string shardId\n  20: optional string historyAddr\n  40: optional string mutableStateInCache\n  50: optional string mutableStateInDatabase\n}"
