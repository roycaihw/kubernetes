/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package openapi

const NamespacedTemplateWithStatusScale = []byte(```{
  "swagger": "2.0",
  "info": {
   "title": "Kubernetes CRD",
   "version": "v0.1.0"
  },
  "paths": {
   "/apis/tmpGroup.io/tmpVersion/namespaces/{namespaces}/tmpPlural": {
    "get": {
     "description": "list objects of kind tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "listTmpGroupIoTmpVersionNamespacedtmpKind",
     "parameters": [
      {
       "uniqueItems": true,
       "type": "string",
       "description": "The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \"next key\".\n\nThis field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.",
       "name": "continue",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
       "name": "fieldSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
       "name": "labelSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.\n\nThe server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.",
       "name": "limit",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.",
       "name": "resourceVersion",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.",
       "name": "timeoutSeconds",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "boolean",
       "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
       "name": "watch",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpListKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "post": {
     "description": "create a tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "createTmpGroupIoTmpVersionNamespacedtmpKind",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "201": {
       "description": "Created",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "202": {
       "description": "Accepted",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "post",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "delete": {
     "description": "delete collection of tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "deleteTmpGroupIoTmpVersionCollectionNamespacedtmpKind",
     "parameters": [
      {
       "uniqueItems": true,
       "type": "string",
       "description": "The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \"next key\".\n\nThis field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.",
       "name": "continue",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
       "name": "fieldSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
       "name": "labelSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.\n\nThe server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.",
       "name": "limit",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.",
       "name": "resourceVersion",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.",
       "name": "timeoutSeconds",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "boolean",
       "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
       "name": "watch",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "delete",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "boolean",
      "description": "If true, partially initialized resources are included in the response.",
      "name": "includeUninitialized",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     }
    ]
   },
   "/apis/tmpGroup.io/tmpVersion/namespaces/{namespaces}/tmpPlural/{name}": {
    "get": {
     "description": "read the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "readTmpGroupIoTmpVersionNamespacedtmpKind",
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "put": {
     "description": "replace the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "replaceTmpGroupIoTmpVersionNamespacedtmpKind",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "201": {
       "description": "Created",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "put",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "delete": {
     "description": "delete a tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "deleteTmpGroupIoTmpVersionNamespacedtmpKind",
     "parameters": [
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately.",
       "name": "gracePeriodSeconds",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "boolean",
       "description": "Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \"orphan\" finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both.",
       "name": "orphanDependents",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground.",
       "name": "propagationPolicy",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
       }
      },
      "202": {
       "description": "Accepted",
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "delete",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "patch": {
     "description": "partially update the specified tmpKind",
     "consumes": [
      "application/json-patch+json",
      "application/merge-patch+json",
      "application/strategic-merge-patch+json"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "patchTmpGroupIoTmpVersionNamespacedtmpKind",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Patch"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "patch",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     }
    ]
   },
   "/apis/tmpGroup.io/tmpVersion/namespaces/{namespaces}/tmpPlural/{name}/scale": {
    "get": {
     "description": "read scale of the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "readTmpGroupIoTmpVersionNamespacedtmpKindScale",
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.api.autoscaling.v1.Scale"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "put": {
     "description": "replace scale of the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "replaceTmpGroupIoTmpVersionNamespacedtmpKindScale",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.k8s.api.autoscaling.v1.Scale"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.api.autoscaling.v1.Scale"
       }
      },
      "201": {
       "description": "Created",
       "schema": {
        "$ref": "#/definitions/io.k8s.api.autoscaling.v1.Scale"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "put",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "patch": {
     "description": "partially update scale of the specified tmpKind",
     "consumes": [
      "application/json-patch+json",
      "application/merge-patch+json",
      "application/strategic-merge-patch+json"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "patchTmpGroupIoTmpVersionNamespacedtmpKindScale",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Patch"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.api.autoscaling.v1.Scale"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "patch",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     }
    ]
   },
   "/apis/tmpGroup.io/tmpVersion/namespaces/{namespaces}/tmpPlural/{name}/status": {
    "get": {
     "description": "read status of the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "readTmpGroupIoTmpVersionNamespacedtmpKindStatus",
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "put": {
     "description": "replace status of the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "replaceTmpGroupIoTmpVersionNamespacedtmpKindStatus",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "201": {
       "description": "Created",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "put",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "patch": {
     "description": "partially update status of the specified tmpKind",
     "consumes": [
      "application/json-patch+json",
      "application/merge-patch+json",
      "application/strategic-merge-patch+json"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "patchTmpGroupIoTmpVersionNamespacedtmpKindStatus",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Patch"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "patch",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     }
    ]
   },
   "/apis/tmpGroup.io/tmpVersion/tmpPlural": {
    "get": {
     "description": "list objects of kind tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "listTmpGroupIoTmpVersionTmpKindForAllNamespaces",
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpListKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "string",
      "description": "The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \"next key\".\n\nThis field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.",
      "name": "continue",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "string",
      "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
      "name": "fieldSelector",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "boolean",
      "description": "If true, partially initialized resources are included in the response.",
      "name": "includeUninitialized",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "string",
      "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
      "name": "labelSelector",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "integer",
      "description": "limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.\n\nThe server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.",
      "name": "limit",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "string",
      "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.",
      "name": "resourceVersion",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "integer",
      "description": "Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.",
      "name": "timeoutSeconds",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "boolean",
      "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
      "name": "watch",
      "in": "query"
     }
    ]
   }
  },
  "definitions": {
   "io.k8s.api.autoscaling.v1.Scale": {
    "description": "Scale represents a scaling request for a resource.",
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "metadata": {
      "description": "Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata.",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta"
     },
     "spec": {
      "description": "defines the behavior of the scale. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status.",
      "$ref": "#/definitions/io.k8s.api.autoscaling.v1.ScaleSpec"
     },
     "status": {
      "description": "current status of the scale. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status. Read-only.",
      "$ref": "#/definitions/io.k8s.api.autoscaling.v1.ScaleStatus"
     }
    }
   },
   "io.k8s.api.autoscaling.v1.ScaleSpec": {
    "description": "ScaleSpec describes the attributes of a scale subresource.",
    "properties": {
     "replicas": {
      "description": "desired number of instances for the scaled object.",
      "type": "integer",
      "format": "int32"
     }
    }
   },
   "io.k8s.api.autoscaling.v1.ScaleStatus": {
    "description": "ScaleStatus represents the current status of a scale subresource.",
    "required": [
     "replicas"
    ],
    "properties": {
     "replicas": {
      "description": "actual number of observed instances of the scaled object.",
      "type": "integer",
      "format": "int32"
     },
     "selector": {
      "description": "label query over pods that should match the replicas count. This is same as the label selector but in the string format to avoid introspection by clients. The string will be in the same format as the query-param syntax. More info about label selectors: http://kubernetes.io/docs/user-guide/labels#label-selectors",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Initializer": {
    "description": "Initializer is information about an initializer that has not yet completed.",
    "required": [
     "name"
    ],
    "properties": {
     "name": {
      "description": "name of the process that is responsible for initializing this object.",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Initializers": {
    "description": "Initializers tracks the progress of initialization.",
    "required": [
     "pending"
    ],
    "properties": {
     "pending": {
      "description": "Pending is a list of initializers that must execute in order before this object is visible. When the last pending initializer is removed, and no failing result is set, the initializers struct will be set to nil and the object is considered as initialized and visible to all clients.",
      "type": "array",
      "items": {
       "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Initializer"
      },
      "x-kubernetes-patch-merge-key": "name",
      "x-kubernetes-patch-strategy": "merge"
     },
     "result": {
      "description": "If result is set with the Failure field, the object will be persisted to storage and then deleted, ensuring that other clients can observe the deletion.",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.ListMeta": {
    "description": "ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.",
    "properties": {
     "continue": {
      "description": "continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response, unless you have received this token from an error message.",
      "type": "string"
     },
     "resourceVersion": {
      "description": "String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
      "type": "string"
     },
     "selfLink": {
      "description": "selfLink is a URL representing this object. Populated by the system. Read-only.",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta": {
    "description": "ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.",
    "properties": {
     "annotations": {
      "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
      "type": "object",
      "additionalProperties": {
       "type": "string"
      }
     },
     "clusterName": {
      "description": "The name of the cluster which the object belongs to. This is used to distinguish resources with same name and namespace in different clusters. This field is not set anywhere right now and apiserver is going to ignore it if set in create or update request.",
      "type": "string"
     },
     "creationTimestamp": {
      "description": "CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.\n\nPopulated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Time"
     },
     "deletionGracePeriodSeconds": {
      "description": "Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only.",
      "type": "integer",
      "format": "int64"
     },
     "deletionTimestamp": {
      "description": "DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This field is set by the server when a graceful deletion is requested by the user, and is not directly settable by a client. The resource is expected to be deleted (no longer visible from resource lists, and not reachable by name) after the time in this field, once the finalizers list is empty. As long as the finalizers list contains items, deletion is blocked. Once the deletionTimestamp is set, this value may not be unset or be set further into the future, although it may be shortened or the resource may be deleted prior to this time. For example, a user may request that a pod is deleted in 30 seconds. The Kubelet will react by sending a graceful termination signal to the containers in the pod. After that 30 seconds, the Kubelet will send a hard termination signal (SIGKILL) to the container and after cleanup, remove the pod from the API. In the presence of network partitions, this object may still exist after this timestamp, until an administrator or automated process can determine the resource is fully terminated. If not set, graceful deletion of the object has not been requested.\n\nPopulated by the system when a graceful deletion is requested. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Time"
     },
     "finalizers": {
      "description": "Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed.",
      "type": "array",
      "items": {
       "type": "string"
      },
      "x-kubernetes-patch-strategy": "merge"
     },
     "generateName": {
      "description": "GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server.\n\nIf this field is specified and the generated name exists, the server will NOT return a 409 - instead, it will either return 201 Created or 500 with Reason ServerTimeout indicating a unique name could not be found in the time allotted, and the client should retry (optionally after the time indicated in the Retry-After header).\n\nApplied only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency",
      "type": "string"
     },
     "generation": {
      "description": "A sequence number representing a specific generation of the desired state. Populated by the system. Read-only.",
      "type": "integer",
      "format": "int64"
     },
     "initializers": {
      "description": "An initializer is a controller which enforces some system invariant at object creation time. This field is a list of initializers that have not yet acted on this object. If nil or empty, this object has been completely initialized. Otherwise, the object is considered uninitialized and is hidden (in list/watch and get calls) from clients that haven't explicitly asked to observe uninitialized objects.\n\nWhen an object is created, the system will populate this list with the current set of initializers. Only privileged users may set or modify this list. Once it is empty, it may not be modified further by any user.",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Initializers"
     },
     "labels": {
      "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
      "type": "object",
      "additionalProperties": {
       "type": "string"
      }
     },
     "name": {
      "description": "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
      "type": "string"
     },
     "namespace": {
      "description": "Namespace defines the space within each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.\n\nMust be a DNS_LABEL. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces",
      "type": "string"
     },
     "ownerReferences": {
      "description": "List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller.",
      "type": "array",
      "items": {
       "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference"
      },
      "x-kubernetes-patch-merge-key": "uid",
      "x-kubernetes-patch-strategy": "merge"
     },
     "resourceVersion": {
      "description": "An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources.\n\nPopulated by the system. Read-only. Value must be treated as opaque by clients and . More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
      "type": "string"
     },
     "selfLink": {
      "description": "SelfLink is a URL representing this object. Populated by the system. Read-only.",
      "type": "string"
     },
     "uid": {
      "description": "UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.\n\nPopulated by the system. Read-only. More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference": {
    "description": "OwnerReference contains enough information to let you identify an owning object. An owning object must be in the same namespace as the dependent, or be cluster-scoped, so there is no namespace field.",
    "required": [
     "apiVersion",
     "kind",
     "name",
     "uid"
    ],
    "properties": {
     "apiVersion": {
      "description": "API version of the referent.",
      "type": "string"
     },
     "blockOwnerDeletion": {
      "description": "If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. Defaults to false. To set this field, a user needs \"delete\" permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.",
      "type": "boolean"
     },
     "controller": {
      "description": "If true, this reference points to the managing controller.",
      "type": "boolean"
     },
     "kind": {
      "description": "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "name": {
      "description": "Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
      "type": "string"
     },
     "uid": {
      "description": "UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Patch": {
    "description": "Patch is provided to give a concrete name and type to the Kubernetes PATCH request body."
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Status": {
    "description": "Status is a return value for calls that don't return other objects.",
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "code": {
      "description": "Suggested HTTP return code for this status, 0 if not set.",
      "type": "integer",
      "format": "int32"
     },
     "details": {
      "description": "Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.StatusDetails"
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "message": {
      "description": "A human-readable description of the status of this operation.",
      "type": "string"
     },
     "metadata": {
      "description": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ListMeta"
     },
     "reason": {
      "description": "A machine-readable description of why this operation is in the \"Failure\" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.",
      "type": "string"
     },
     "status": {
      "description": "Status of the operation. One of: \"Success\" or \"Failure\". More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.StatusCause": {
    "description": "StatusCause provides more information about an api.Status failure, including cases when multiple errors are encountered.",
    "properties": {
     "field": {
      "description": "The field of the resource that has caused this error, as named by its JSON serialization. May include dot and postfix notation for nested attributes. Arrays are zero-indexed.  Fields may appear more than once in an array of causes due to fields having multiple errors. Optional.\n\nExamples:\n  \"name\" - the field \"name\" on the current resource\n  \"items[0].name\" - the field \"name\" on the first array entry in \"items\"",
      "type": "string"
     },
     "message": {
      "description": "A human-readable description of the cause of the error.  This field may be presented as-is to a reader.",
      "type": "string"
     },
     "reason": {
      "description": "A machine-readable description of the cause of the error. If this value is empty there is no information available.",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.StatusDetails": {
    "description": "StatusDetails is a set of additional properties that MAY be set by the server to provide additional information about a response. The Reason field of a Status object defines what attributes will be set. Clients must ignore fields that do not match the defined type of each attribute, and should assume that any attribute may be empty, invalid, or under defined.",
    "properties": {
     "causes": {
      "description": "The Causes array includes more details associated with the StatusReason failure. Not all StatusReasons may provide detailed causes.",
      "type": "array",
      "items": {
       "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.StatusCause"
      }
     },
     "group": {
      "description": "The group attribute of the resource associated with the status StatusReason.",
      "type": "string"
     },
     "kind": {
      "description": "The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "name": {
      "description": "The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described).",
      "type": "string"
     },
     "retryAfterSeconds": {
      "description": "If specified, the time in seconds before the operation should be retried. Some errors may indicate the client must take an alternate action - for those errors this field may indicate how long to wait before taking the alternate action.",
      "type": "integer",
      "format": "int32"
     },
     "uid": {
      "description": "UID of the resource. (when there is a single resource which can be described). More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Time": {
    "description": "Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.",
    "type": "string",
    "format": "date-time"
   },
   "io.tmpGroup.tmpVersion.tmpKind": {
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "metadata": {
      "description": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta"
     }
    },
    "x-kubernetes-group-version-kind": [
     {
      "group": "tmpGroup.io",
      "kind": "tmpKind",
      "version": "tmpVersion"
     }
    ]
   },
   "io.tmpGroup.tmpVersion.tmpListKind": {
    "description": "tmpListKind is a list of tmpKind",
    "required": [
     "items"
    ],
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "items": {
      "description": "List of tmpPlural. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md",
      "type": "array",
      "items": {
       "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
      }
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "metadata": {
      "description": "ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.",
      "properties": {
       "continue": {
        "description": "continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response, unless you have received this token from an error message.",
        "type": "string"
       },
       "resourceVersion": {
        "description": "String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
        "type": "string"
       },
       "selfLink": {
        "description": "selfLink is a URL representing this object. Populated by the system. Read-only.",
        "type": "string"
       }
      }
     }
    },
    "x-kubernetes-group-version-kind": [
     {
      "group": "tmpGroup.io",
      "kind": "tmpListKind",
      "version": "tmpVersion"
     }
    ]
   }
  }
 }```)

const NamespacedTemplateWithStatus = []byte(```{
  "swagger": "2.0",
  "info": {
   "title": "Kubernetes CRD",
   "version": "v0.1.0"
  },
  "paths": {
   "/apis/tmpGroup.io/tmpVersion/namespaces/{namespaces}/tmpPlural": {
    "get": {
     "description": "list objects of kind tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "listTmpGroupIoTmpVersionNamespacedtmpKind",
     "parameters": [
      {
       "uniqueItems": true,
       "type": "string",
       "description": "The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \"next key\".\n\nThis field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.",
       "name": "continue",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
       "name": "fieldSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
       "name": "labelSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.\n\nThe server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.",
       "name": "limit",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.",
       "name": "resourceVersion",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.",
       "name": "timeoutSeconds",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "boolean",
       "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
       "name": "watch",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpListKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "post": {
     "description": "create a tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "createTmpGroupIoTmpVersionNamespacedtmpKind",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "201": {
       "description": "Created",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "202": {
       "description": "Accepted",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "post",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "delete": {
     "description": "delete collection of tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "deleteTmpGroupIoTmpVersionCollectionNamespacedtmpKind",
     "parameters": [
      {
       "uniqueItems": true,
       "type": "string",
       "description": "The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \"next key\".\n\nThis field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.",
       "name": "continue",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
       "name": "fieldSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
       "name": "labelSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.\n\nThe server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.",
       "name": "limit",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.",
       "name": "resourceVersion",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.",
       "name": "timeoutSeconds",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "boolean",
       "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
       "name": "watch",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "delete",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "boolean",
      "description": "If true, partially initialized resources are included in the response.",
      "name": "includeUninitialized",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     }
    ]
   },
   "/apis/tmpGroup.io/tmpVersion/namespaces/{namespaces}/tmpPlural/{name}": {
    "get": {
     "description": "read the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "readTmpGroupIoTmpVersionNamespacedtmpKind",
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "put": {
     "description": "replace the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "replaceTmpGroupIoTmpVersionNamespacedtmpKind",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "201": {
       "description": "Created",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "put",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "delete": {
     "description": "delete a tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "deleteTmpGroupIoTmpVersionNamespacedtmpKind",
     "parameters": [
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately.",
       "name": "gracePeriodSeconds",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "boolean",
       "description": "Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \"orphan\" finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both.",
       "name": "orphanDependents",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground.",
       "name": "propagationPolicy",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
       }
      },
      "202": {
       "description": "Accepted",
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "delete",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "patch": {
     "description": "partially update the specified tmpKind",
     "consumes": [
      "application/json-patch+json",
      "application/merge-patch+json",
      "application/strategic-merge-patch+json"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "patchTmpGroupIoTmpVersionNamespacedtmpKind",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Patch"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "patch",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     }
    ]
   },
   "/apis/tmpGroup.io/tmpVersion/namespaces/{namespaces}/tmpPlural/{name}/status": {
    "get": {
     "description": "read status of the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "readTmpGroupIoTmpVersionNamespacedtmpKindStatus",
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "put": {
     "description": "replace status of the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "replaceTmpGroupIoTmpVersionNamespacedtmpKindStatus",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "201": {
       "description": "Created",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "put",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "patch": {
     "description": "partially update status of the specified tmpKind",
     "consumes": [
      "application/json-patch+json",
      "application/merge-patch+json",
      "application/strategic-merge-patch+json"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "patchTmpGroupIoTmpVersionNamespacedtmpKindStatus",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Patch"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "patch",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     }
    ]
   },
   "/apis/tmpGroup.io/tmpVersion/tmpPlural": {
    "get": {
     "description": "list objects of kind tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "listTmpGroupIoTmpVersionTmpKindForAllNamespaces",
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpListKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "string",
      "description": "The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \"next key\".\n\nThis field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.",
      "name": "continue",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "string",
      "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
      "name": "fieldSelector",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "boolean",
      "description": "If true, partially initialized resources are included in the response.",
      "name": "includeUninitialized",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "string",
      "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
      "name": "labelSelector",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "integer",
      "description": "limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.\n\nThe server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.",
      "name": "limit",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "string",
      "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.",
      "name": "resourceVersion",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "integer",
      "description": "Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.",
      "name": "timeoutSeconds",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "boolean",
      "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
      "name": "watch",
      "in": "query"
     }
    ]
   }
  },
  "definitions": {
   "io.k8s.apimachinery.pkg.apis.meta.v1.ListMeta": {
    "description": "ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.",
    "properties": {
     "continue": {
      "description": "continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response, unless you have received this token from an error message.",
      "type": "string"
     },
     "resourceVersion": {
      "description": "String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
      "type": "string"
     },
     "selfLink": {
      "description": "selfLink is a URL representing this object. Populated by the system. Read-only.",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Patch": {
    "description": "Patch is provided to give a concrete name and type to the Kubernetes PATCH request body."
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Status": {
    "description": "Status is a return value for calls that don't return other objects.",
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "code": {
      "description": "Suggested HTTP return code for this status, 0 if not set.",
      "type": "integer",
      "format": "int32"
     },
     "details": {
      "description": "Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.StatusDetails"
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "message": {
      "description": "A human-readable description of the status of this operation.",
      "type": "string"
     },
     "metadata": {
      "description": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ListMeta"
     },
     "reason": {
      "description": "A machine-readable description of why this operation is in the \"Failure\" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.",
      "type": "string"
     },
     "status": {
      "description": "Status of the operation. One of: \"Success\" or \"Failure\". More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.StatusCause": {
    "description": "StatusCause provides more information about an api.Status failure, including cases when multiple errors are encountered.",
    "properties": {
     "field": {
      "description": "The field of the resource that has caused this error, as named by its JSON serialization. May include dot and postfix notation for nested attributes. Arrays are zero-indexed.  Fields may appear more than once in an array of causes due to fields having multiple errors. Optional.\n\nExamples:\n  \"name\" - the field \"name\" on the current resource\n  \"items[0].name\" - the field \"name\" on the first array entry in \"items\"",
      "type": "string"
     },
     "message": {
      "description": "A human-readable description of the cause of the error.  This field may be presented as-is to a reader.",
      "type": "string"
     },
     "reason": {
      "description": "A machine-readable description of the cause of the error. If this value is empty there is no information available.",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.StatusDetails": {
    "description": "StatusDetails is a set of additional properties that MAY be set by the server to provide additional information about a response. The Reason field of a Status object defines what attributes will be set. Clients must ignore fields that do not match the defined type of each attribute, and should assume that any attribute may be empty, invalid, or under defined.",
    "properties": {
     "causes": {
      "description": "The Causes array includes more details associated with the StatusReason failure. Not all StatusReasons may provide detailed causes.",
      "type": "array",
      "items": {
       "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.StatusCause"
      }
     },
     "group": {
      "description": "The group attribute of the resource associated with the status StatusReason.",
      "type": "string"
     },
     "kind": {
      "description": "The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "name": {
      "description": "The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described).",
      "type": "string"
     },
     "retryAfterSeconds": {
      "description": "If specified, the time in seconds before the operation should be retried. Some errors may indicate the client must take an alternate action - for those errors this field may indicate how long to wait before taking the alternate action.",
      "type": "integer",
      "format": "int32"
     },
     "uid": {
      "description": "UID of the resource. (when there is a single resource which can be described). More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
      "type": "string"
     }
    }
   },
   "io.tmpGroup.tmpVersion.tmpKind": {
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "metadata": {
      "description": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta"
     }
    },
    "x-kubernetes-group-version-kind": [
     {
      "group": "tmpGroup.io",
      "kind": "tmpKind",
      "version": "tmpVersion"
     }
    ]
   },
   "io.tmpGroup.tmpVersion.tmpListKind": {
    "description": "tmpListKind is a list of tmpKind",
    "required": [
     "items"
    ],
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "items": {
      "description": "List of tmpPlural. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md",
      "type": "array",
      "items": {
       "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
      }
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "metadata": {
      "description": "ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.",
      "properties": {
       "continue": {
        "description": "continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response, unless you have received this token from an error message.",
        "type": "string"
       },
       "resourceVersion": {
        "description": "String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
        "type": "string"
       },
       "selfLink": {
        "description": "selfLink is a URL representing this object. Populated by the system. Read-only.",
        "type": "string"
       }
      }
     }
    },
    "x-kubernetes-group-version-kind": [
     {
      "group": "tmpGroup.io",
      "kind": "tmpListKind",
      "version": "tmpVersion"
     }
    ]
   }
  }
 }```)

const NamespacedTemplateWithScale = []byte(```{
  "swagger": "2.0",
  "info": {
   "title": "Kubernetes CRD",
   "version": "v0.1.0"
  },
  "paths": {
   "/apis/tmpGroup.io/tmpVersion/namespaces/{namespaces}/tmpPlural": {
    "get": {
     "description": "list objects of kind tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "listTmpGroupIoTmpVersionNamespacedtmpKind",
     "parameters": [
      {
       "uniqueItems": true,
       "type": "string",
       "description": "The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \"next key\".\n\nThis field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.",
       "name": "continue",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
       "name": "fieldSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
       "name": "labelSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.\n\nThe server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.",
       "name": "limit",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.",
       "name": "resourceVersion",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.",
       "name": "timeoutSeconds",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "boolean",
       "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
       "name": "watch",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpListKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "post": {
     "description": "create a tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "createTmpGroupIoTmpVersionNamespacedtmpKind",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "201": {
       "description": "Created",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "202": {
       "description": "Accepted",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "post",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "delete": {
     "description": "delete collection of tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "deleteTmpGroupIoTmpVersionCollectionNamespacedtmpKind",
     "parameters": [
      {
       "uniqueItems": true,
       "type": "string",
       "description": "The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \"next key\".\n\nThis field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.",
       "name": "continue",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
       "name": "fieldSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
       "name": "labelSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.\n\nThe server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.",
       "name": "limit",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.",
       "name": "resourceVersion",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.",
       "name": "timeoutSeconds",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "boolean",
       "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
       "name": "watch",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "delete",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "boolean",
      "description": "If true, partially initialized resources are included in the response.",
      "name": "includeUninitialized",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     }
    ]
   },
   "/apis/tmpGroup.io/tmpVersion/namespaces/{namespaces}/tmpPlural/{name}": {
    "get": {
     "description": "read the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "readTmpGroupIoTmpVersionNamespacedtmpKind",
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "put": {
     "description": "replace the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "replaceTmpGroupIoTmpVersionNamespacedtmpKind",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "201": {
       "description": "Created",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "put",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "delete": {
     "description": "delete a tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "deleteTmpGroupIoTmpVersionNamespacedtmpKind",
     "parameters": [
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately.",
       "name": "gracePeriodSeconds",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "boolean",
       "description": "Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \"orphan\" finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both.",
       "name": "orphanDependents",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground.",
       "name": "propagationPolicy",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
       }
      },
      "202": {
       "description": "Accepted",
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "delete",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "patch": {
     "description": "partially update the specified tmpKind",
     "consumes": [
      "application/json-patch+json",
      "application/merge-patch+json",
      "application/strategic-merge-patch+json"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "patchTmpGroupIoTmpVersionNamespacedtmpKind",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Patch"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "patch",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     }
    ]
   },
   "/apis/tmpGroup.io/tmpVersion/namespaces/{namespaces}/tmpPlural/{name}/scale": {
    "get": {
     "description": "read scale of the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "readTmpGroupIoTmpVersionNamespacedtmpKindScale",
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.api.autoscaling.v1.Scale"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "put": {
     "description": "replace scale of the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "replaceTmpGroupIoTmpVersionNamespacedtmpKindScale",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.k8s.api.autoscaling.v1.Scale"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.api.autoscaling.v1.Scale"
       }
      },
      "201": {
       "description": "Created",
       "schema": {
        "$ref": "#/definitions/io.k8s.api.autoscaling.v1.Scale"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "put",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "patch": {
     "description": "partially update scale of the specified tmpKind",
     "consumes": [
      "application/json-patch+json",
      "application/merge-patch+json",
      "application/strategic-merge-patch+json"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "patchTmpGroupIoTmpVersionNamespacedtmpKindScale",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Patch"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.api.autoscaling.v1.Scale"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "patch",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     }
    ]
   },
   "/apis/tmpGroup.io/tmpVersion/tmpPlural": {
    "get": {
     "description": "list objects of kind tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "listTmpGroupIoTmpVersionTmpKindForAllNamespaces",
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpListKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "string",
      "description": "The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \"next key\".\n\nThis field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.",
      "name": "continue",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "string",
      "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
      "name": "fieldSelector",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "boolean",
      "description": "If true, partially initialized resources are included in the response.",
      "name": "includeUninitialized",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "string",
      "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
      "name": "labelSelector",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "integer",
      "description": "limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.\n\nThe server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.",
      "name": "limit",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "string",
      "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.",
      "name": "resourceVersion",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "integer",
      "description": "Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.",
      "name": "timeoutSeconds",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "boolean",
      "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
      "name": "watch",
      "in": "query"
     }
    ]
   }
  },
  "definitions": {
   "io.k8s.api.autoscaling.v1.Scale": {
    "description": "Scale represents a scaling request for a resource.",
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "metadata": {
      "description": "Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata.",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta"
     },
     "spec": {
      "description": "defines the behavior of the scale. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status.",
      "$ref": "#/definitions/io.k8s.api.autoscaling.v1.ScaleSpec"
     },
     "status": {
      "description": "current status of the scale. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status. Read-only.",
      "$ref": "#/definitions/io.k8s.api.autoscaling.v1.ScaleStatus"
     }
    }
   },
   "io.k8s.api.autoscaling.v1.ScaleSpec": {
    "description": "ScaleSpec describes the attributes of a scale subresource.",
    "properties": {
     "replicas": {
      "description": "desired number of instances for the scaled object.",
      "type": "integer",
      "format": "int32"
     }
    }
   },
   "io.k8s.api.autoscaling.v1.ScaleStatus": {
    "description": "ScaleStatus represents the current status of a scale subresource.",
    "required": [
     "replicas"
    ],
    "properties": {
     "replicas": {
      "description": "actual number of observed instances of the scaled object.",
      "type": "integer",
      "format": "int32"
     },
     "selector": {
      "description": "label query over pods that should match the replicas count. This is same as the label selector but in the string format to avoid introspection by clients. The string will be in the same format as the query-param syntax. More info about label selectors: http://kubernetes.io/docs/user-guide/labels#label-selectors",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Initializer": {
    "description": "Initializer is information about an initializer that has not yet completed.",
    "required": [
     "name"
    ],
    "properties": {
     "name": {
      "description": "name of the process that is responsible for initializing this object.",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Initializers": {
    "description": "Initializers tracks the progress of initialization.",
    "required": [
     "pending"
    ],
    "properties": {
     "pending": {
      "description": "Pending is a list of initializers that must execute in order before this object is visible. When the last pending initializer is removed, and no failing result is set, the initializers struct will be set to nil and the object is considered as initialized and visible to all clients.",
      "type": "array",
      "items": {
       "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Initializer"
      },
      "x-kubernetes-patch-merge-key": "name",
      "x-kubernetes-patch-strategy": "merge"
     },
     "result": {
      "description": "If result is set with the Failure field, the object will be persisted to storage and then deleted, ensuring that other clients can observe the deletion.",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.ListMeta": {
    "description": "ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.",
    "properties": {
     "continue": {
      "description": "continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response, unless you have received this token from an error message.",
      "type": "string"
     },
     "resourceVersion": {
      "description": "String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
      "type": "string"
     },
     "selfLink": {
      "description": "selfLink is a URL representing this object. Populated by the system. Read-only.",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta": {
    "description": "ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.",
    "properties": {
     "annotations": {
      "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
      "type": "object",
      "additionalProperties": {
       "type": "string"
      }
     },
     "clusterName": {
      "description": "The name of the cluster which the object belongs to. This is used to distinguish resources with same name and namespace in different clusters. This field is not set anywhere right now and apiserver is going to ignore it if set in create or update request.",
      "type": "string"
     },
     "creationTimestamp": {
      "description": "CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.\n\nPopulated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Time"
     },
     "deletionGracePeriodSeconds": {
      "description": "Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only.",
      "type": "integer",
      "format": "int64"
     },
     "deletionTimestamp": {
      "description": "DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This field is set by the server when a graceful deletion is requested by the user, and is not directly settable by a client. The resource is expected to be deleted (no longer visible from resource lists, and not reachable by name) after the time in this field, once the finalizers list is empty. As long as the finalizers list contains items, deletion is blocked. Once the deletionTimestamp is set, this value may not be unset or be set further into the future, although it may be shortened or the resource may be deleted prior to this time. For example, a user may request that a pod is deleted in 30 seconds. The Kubelet will react by sending a graceful termination signal to the containers in the pod. After that 30 seconds, the Kubelet will send a hard termination signal (SIGKILL) to the container and after cleanup, remove the pod from the API. In the presence of network partitions, this object may still exist after this timestamp, until an administrator or automated process can determine the resource is fully terminated. If not set, graceful deletion of the object has not been requested.\n\nPopulated by the system when a graceful deletion is requested. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Time"
     },
     "finalizers": {
      "description": "Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed.",
      "type": "array",
      "items": {
       "type": "string"
      },
      "x-kubernetes-patch-strategy": "merge"
     },
     "generateName": {
      "description": "GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server.\n\nIf this field is specified and the generated name exists, the server will NOT return a 409 - instead, it will either return 201 Created or 500 with Reason ServerTimeout indicating a unique name could not be found in the time allotted, and the client should retry (optionally after the time indicated in the Retry-After header).\n\nApplied only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency",
      "type": "string"
     },
     "generation": {
      "description": "A sequence number representing a specific generation of the desired state. Populated by the system. Read-only.",
      "type": "integer",
      "format": "int64"
     },
     "initializers": {
      "description": "An initializer is a controller which enforces some system invariant at object creation time. This field is a list of initializers that have not yet acted on this object. If nil or empty, this object has been completely initialized. Otherwise, the object is considered uninitialized and is hidden (in list/watch and get calls) from clients that haven't explicitly asked to observe uninitialized objects.\n\nWhen an object is created, the system will populate this list with the current set of initializers. Only privileged users may set or modify this list. Once it is empty, it may not be modified further by any user.",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Initializers"
     },
     "labels": {
      "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
      "type": "object",
      "additionalProperties": {
       "type": "string"
      }
     },
     "name": {
      "description": "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
      "type": "string"
     },
     "namespace": {
      "description": "Namespace defines the space within each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.\n\nMust be a DNS_LABEL. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces",
      "type": "string"
     },
     "ownerReferences": {
      "description": "List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller.",
      "type": "array",
      "items": {
       "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference"
      },
      "x-kubernetes-patch-merge-key": "uid",
      "x-kubernetes-patch-strategy": "merge"
     },
     "resourceVersion": {
      "description": "An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources.\n\nPopulated by the system. Read-only. Value must be treated as opaque by clients and . More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
      "type": "string"
     },
     "selfLink": {
      "description": "SelfLink is a URL representing this object. Populated by the system. Read-only.",
      "type": "string"
     },
     "uid": {
      "description": "UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.\n\nPopulated by the system. Read-only. More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference": {
    "description": "OwnerReference contains enough information to let you identify an owning object. An owning object must be in the same namespace as the dependent, or be cluster-scoped, so there is no namespace field.",
    "required": [
     "apiVersion",
     "kind",
     "name",
     "uid"
    ],
    "properties": {
     "apiVersion": {
      "description": "API version of the referent.",
      "type": "string"
     },
     "blockOwnerDeletion": {
      "description": "If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. Defaults to false. To set this field, a user needs \"delete\" permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.",
      "type": "boolean"
     },
     "controller": {
      "description": "If true, this reference points to the managing controller.",
      "type": "boolean"
     },
     "kind": {
      "description": "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "name": {
      "description": "Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
      "type": "string"
     },
     "uid": {
      "description": "UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Patch": {
    "description": "Patch is provided to give a concrete name and type to the Kubernetes PATCH request body."
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Status": {
    "description": "Status is a return value for calls that don't return other objects.",
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "code": {
      "description": "Suggested HTTP return code for this status, 0 if not set.",
      "type": "integer",
      "format": "int32"
     },
     "details": {
      "description": "Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.StatusDetails"
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "message": {
      "description": "A human-readable description of the status of this operation.",
      "type": "string"
     },
     "metadata": {
      "description": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ListMeta"
     },
     "reason": {
      "description": "A machine-readable description of why this operation is in the \"Failure\" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.",
      "type": "string"
     },
     "status": {
      "description": "Status of the operation. One of: \"Success\" or \"Failure\". More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.StatusCause": {
    "description": "StatusCause provides more information about an api.Status failure, including cases when multiple errors are encountered.",
    "properties": {
     "field": {
      "description": "The field of the resource that has caused this error, as named by its JSON serialization. May include dot and postfix notation for nested attributes. Arrays are zero-indexed.  Fields may appear more than once in an array of causes due to fields having multiple errors. Optional.\n\nExamples:\n  \"name\" - the field \"name\" on the current resource\n  \"items[0].name\" - the field \"name\" on the first array entry in \"items\"",
      "type": "string"
     },
     "message": {
      "description": "A human-readable description of the cause of the error.  This field may be presented as-is to a reader.",
      "type": "string"
     },
     "reason": {
      "description": "A machine-readable description of the cause of the error. If this value is empty there is no information available.",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.StatusDetails": {
    "description": "StatusDetails is a set of additional properties that MAY be set by the server to provide additional information about a response. The Reason field of a Status object defines what attributes will be set. Clients must ignore fields that do not match the defined type of each attribute, and should assume that any attribute may be empty, invalid, or under defined.",
    "properties": {
     "causes": {
      "description": "The Causes array includes more details associated with the StatusReason failure. Not all StatusReasons may provide detailed causes.",
      "type": "array",
      "items": {
       "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.StatusCause"
      }
     },
     "group": {
      "description": "The group attribute of the resource associated with the status StatusReason.",
      "type": "string"
     },
     "kind": {
      "description": "The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "name": {
      "description": "The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described).",
      "type": "string"
     },
     "retryAfterSeconds": {
      "description": "If specified, the time in seconds before the operation should be retried. Some errors may indicate the client must take an alternate action - for those errors this field may indicate how long to wait before taking the alternate action.",
      "type": "integer",
      "format": "int32"
     },
     "uid": {
      "description": "UID of the resource. (when there is a single resource which can be described). More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Time": {
    "description": "Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.",
    "type": "string",
    "format": "date-time"
   },
   "io.tmpGroup.tmpVersion.tmpKind": {
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "metadata": {
      "description": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta"
     }
    },
    "x-kubernetes-group-version-kind": [
     {
      "group": "tmpGroup.io",
      "kind": "tmpKind",
      "version": "tmpVersion"
     }
    ]
   },
   "io.tmpGroup.tmpVersion.tmpListKind": {
    "description": "tmpListKind is a list of tmpKind",
    "required": [
     "items"
    ],
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "items": {
      "description": "List of tmpPlural. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md",
      "type": "array",
      "items": {
       "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
      }
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "metadata": {
      "description": "ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.",
      "properties": {
       "continue": {
        "description": "continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response, unless you have received this token from an error message.",
        "type": "string"
       },
       "resourceVersion": {
        "description": "String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
        "type": "string"
       },
       "selfLink": {
        "description": "selfLink is a URL representing this object. Populated by the system. Read-only.",
        "type": "string"
       }
      }
     }
    },
    "x-kubernetes-group-version-kind": [
     {
      "group": "tmpGroup.io",
      "kind": "tmpListKind",
      "version": "tmpVersion"
     }
    ]
   }
  }
 }```)

const NamespacedTemplate = []byte(```{
  "swagger": "2.0",
  "info": {
   "title": "Kubernetes CRD",
   "version": "v0.1.0"
  },
  "paths": {
   "/apis/tmpGroup.io/tmpVersion/namespaces/{namespaces}/tmpPlural": {
    "get": {
     "description": "list objects of kind tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "listTmpGroupIoTmpVersionNamespacedtmpKind",
     "parameters": [
      {
       "uniqueItems": true,
       "type": "string",
       "description": "The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \"next key\".\n\nThis field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.",
       "name": "continue",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
       "name": "fieldSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
       "name": "labelSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.\n\nThe server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.",
       "name": "limit",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.",
       "name": "resourceVersion",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.",
       "name": "timeoutSeconds",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "boolean",
       "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
       "name": "watch",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpListKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "post": {
     "description": "create a tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "createTmpGroupIoTmpVersionNamespacedtmpKind",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "201": {
       "description": "Created",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "202": {
       "description": "Accepted",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "post",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "delete": {
     "description": "delete collection of tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "deleteTmpGroupIoTmpVersionCollectionNamespacedtmpKind",
     "parameters": [
      {
       "uniqueItems": true,
       "type": "string",
       "description": "The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \"next key\".\n\nThis field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.",
       "name": "continue",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
       "name": "fieldSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
       "name": "labelSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.\n\nThe server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.",
       "name": "limit",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.",
       "name": "resourceVersion",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.",
       "name": "timeoutSeconds",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "boolean",
       "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
       "name": "watch",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "delete",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "boolean",
      "description": "If true, partially initialized resources are included in the response.",
      "name": "includeUninitialized",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     }
    ]
   },
   "/apis/tmpGroup.io/tmpVersion/namespaces/{namespaces}/tmpPlural/{name}": {
    "get": {
     "description": "read the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "readTmpGroupIoTmpVersionNamespacedtmpKind",
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "put": {
     "description": "replace the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "replaceTmpGroupIoTmpVersionNamespacedtmpKind",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "201": {
       "description": "Created",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "put",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "delete": {
     "description": "delete a tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "deleteTmpGroupIoTmpVersionNamespacedtmpKind",
     "parameters": [
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately.",
       "name": "gracePeriodSeconds",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "boolean",
       "description": "Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \"orphan\" finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both.",
       "name": "orphanDependents",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground.",
       "name": "propagationPolicy",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
       }
      },
      "202": {
       "description": "Accepted",
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "delete",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "patch": {
     "description": "partially update the specified tmpKind",
     "consumes": [
      "application/json-patch+json",
      "application/merge-patch+json",
      "application/strategic-merge-patch+json"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "patchTmpGroupIoTmpVersionNamespacedtmpKind",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Patch"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "patch",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     }
    ]
   },
   "/apis/tmpGroup.io/tmpVersion/tmpPlural": {
    "get": {
     "description": "list objects of kind tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "listTmpGroupIoTmpVersionTmpKindForAllNamespaces",
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpListKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "string",
      "description": "The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \"next key\".\n\nThis field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.",
      "name": "continue",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "string",
      "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
      "name": "fieldSelector",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "boolean",
      "description": "If true, partially initialized resources are included in the response.",
      "name": "includeUninitialized",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "string",
      "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
      "name": "labelSelector",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "integer",
      "description": "limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.\n\nThe server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.",
      "name": "limit",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "string",
      "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.",
      "name": "resourceVersion",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "integer",
      "description": "Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.",
      "name": "timeoutSeconds",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "boolean",
      "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
      "name": "watch",
      "in": "query"
     }
    ]
   }
  },
  "definitions": {
   "io.k8s.apimachinery.pkg.apis.meta.v1.ListMeta": {
    "description": "ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.",
    "properties": {
     "continue": {
      "description": "continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response, unless you have received this token from an error message.",
      "type": "string"
     },
     "resourceVersion": {
      "description": "String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
      "type": "string"
     },
     "selfLink": {
      "description": "selfLink is a URL representing this object. Populated by the system. Read-only.",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Patch": {
    "description": "Patch is provided to give a concrete name and type to the Kubernetes PATCH request body."
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Status": {
    "description": "Status is a return value for calls that don't return other objects.",
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "code": {
      "description": "Suggested HTTP return code for this status, 0 if not set.",
      "type": "integer",
      "format": "int32"
     },
     "details": {
      "description": "Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.StatusDetails"
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "message": {
      "description": "A human-readable description of the status of this operation.",
      "type": "string"
     },
     "metadata": {
      "description": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ListMeta"
     },
     "reason": {
      "description": "A machine-readable description of why this operation is in the \"Failure\" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.",
      "type": "string"
     },
     "status": {
      "description": "Status of the operation. One of: \"Success\" or \"Failure\". More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.StatusCause": {
    "description": "StatusCause provides more information about an api.Status failure, including cases when multiple errors are encountered.",
    "properties": {
     "field": {
      "description": "The field of the resource that has caused this error, as named by its JSON serialization. May include dot and postfix notation for nested attributes. Arrays are zero-indexed.  Fields may appear more than once in an array of causes due to fields having multiple errors. Optional.\n\nExamples:\n  \"name\" - the field \"name\" on the current resource\n  \"items[0].name\" - the field \"name\" on the first array entry in \"items\"",
      "type": "string"
     },
     "message": {
      "description": "A human-readable description of the cause of the error.  This field may be presented as-is to a reader.",
      "type": "string"
     },
     "reason": {
      "description": "A machine-readable description of the cause of the error. If this value is empty there is no information available.",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.StatusDetails": {
    "description": "StatusDetails is a set of additional properties that MAY be set by the server to provide additional information about a response. The Reason field of a Status object defines what attributes will be set. Clients must ignore fields that do not match the defined type of each attribute, and should assume that any attribute may be empty, invalid, or under defined.",
    "properties": {
     "causes": {
      "description": "The Causes array includes more details associated with the StatusReason failure. Not all StatusReasons may provide detailed causes.",
      "type": "array",
      "items": {
       "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.StatusCause"
      }
     },
     "group": {
      "description": "The group attribute of the resource associated with the status StatusReason.",
      "type": "string"
     },
     "kind": {
      "description": "The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "name": {
      "description": "The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described).",
      "type": "string"
     },
     "retryAfterSeconds": {
      "description": "If specified, the time in seconds before the operation should be retried. Some errors may indicate the client must take an alternate action - for those errors this field may indicate how long to wait before taking the alternate action.",
      "type": "integer",
      "format": "int32"
     },
     "uid": {
      "description": "UID of the resource. (when there is a single resource which can be described). More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
      "type": "string"
     }
    }
   },
   "io.tmpGroup.tmpVersion.tmpKind": {
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "metadata": {
      "description": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta"
     }
    },
    "x-kubernetes-group-version-kind": [
     {
      "group": "tmpGroup.io",
      "kind": "tmpKind",
      "version": "tmpVersion"
     }
    ]
   },
   "io.tmpGroup.tmpVersion.tmpListKind": {
    "description": "tmpListKind is a list of tmpKind",
    "required": [
     "items"
    ],
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "items": {
      "description": "List of tmpPlural. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md",
      "type": "array",
      "items": {
       "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
      }
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "metadata": {
      "description": "ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.",
      "properties": {
       "continue": {
        "description": "continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response, unless you have received this token from an error message.",
        "type": "string"
       },
       "resourceVersion": {
        "description": "String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
        "type": "string"
       },
       "selfLink": {
        "description": "selfLink is a URL representing this object. Populated by the system. Read-only.",
        "type": "string"
       }
      }
     }
    },
    "x-kubernetes-group-version-kind": [
     {
      "group": "tmpGroup.io",
      "kind": "tmpListKind",
      "version": "tmpVersion"
     }
    ]
   }
  }
 }```)

const TemplateWithStatusScale = []byte(```{
  "swagger": "2.0",
  "info": {
   "title": "Kubernetes CRD",
   "version": "v0.1.0"
  },
  "paths": {
   "/apis/tmpGroup.io/tmpVersion/tmpPlural": {
    "get": {
     "description": "list objects of kind tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "listTmpGroupIoTmpVersionTmpKind",
     "parameters": [
      {
       "uniqueItems": true,
       "type": "string",
       "description": "The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \"next key\".\n\nThis field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.",
       "name": "continue",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
       "name": "fieldSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
       "name": "labelSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.\n\nThe server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.",
       "name": "limit",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.",
       "name": "resourceVersion",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.",
       "name": "timeoutSeconds",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "boolean",
       "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
       "name": "watch",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpListKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "post": {
     "description": "create a tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "createTmpGroupIoTmpVersionTmpKind",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "201": {
       "description": "Created",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "202": {
       "description": "Accepted",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "post",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "delete": {
     "description": "delete collection of tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "deleteTmpGroupIoTmpVersionCollectiontmpKind",
     "parameters": [
      {
       "uniqueItems": true,
       "type": "string",
       "description": "The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \"next key\".\n\nThis field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.",
       "name": "continue",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
       "name": "fieldSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
       "name": "labelSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.\n\nThe server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.",
       "name": "limit",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.",
       "name": "resourceVersion",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.",
       "name": "timeoutSeconds",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "boolean",
       "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
       "name": "watch",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "delete",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "boolean",
      "description": "If true, partially initialized resources are included in the response.",
      "name": "includeUninitialized",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     }
    ]
   },
   "/apis/tmpGroup.io/tmpVersion/tmpPlural/{name}": {
    "get": {
     "description": "read the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "readTmpGroupIoTmpVersionTmpKind",
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "put": {
     "description": "replace the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "replaceTmpGroupIoTmpVersionTmpKind",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "201": {
       "description": "Created",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "put",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "delete": {
     "description": "delete a tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "deleteTmpGroupIoTmpVersionTmpKind",
     "parameters": [
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately.",
       "name": "gracePeriodSeconds",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "boolean",
       "description": "Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \"orphan\" finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both.",
       "name": "orphanDependents",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground.",
       "name": "propagationPolicy",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
       }
      },
      "202": {
       "description": "Accepted",
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "delete",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "patch": {
     "description": "partially update the specified tmpKind",
     "consumes": [
      "application/json-patch+json",
      "application/merge-patch+json",
      "application/strategic-merge-patch+json"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "patchTmpGroupIoTmpVersionTmpKind",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Patch"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "patch",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     }
    ]
   },
   "/apis/tmpGroup.io/tmpVersion/tmpPlural/{name}/scale": {
    "get": {
     "description": "read scale of the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "readTmpGroupIoTmpVersionTmpKindScale",
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.api.autoscaling.v1.Scale"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "put": {
     "description": "replace scale of the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "replaceTmpGroupIoTmpVersionTmpKindScale",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.k8s.api.autoscaling.v1.Scale"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.api.autoscaling.v1.Scale"
       }
      },
      "201": {
       "description": "Created",
       "schema": {
        "$ref": "#/definitions/io.k8s.api.autoscaling.v1.Scale"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "put",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "patch": {
     "description": "partially update scale of the specified tmpKind",
     "consumes": [
      "application/json-patch+json",
      "application/merge-patch+json",
      "application/strategic-merge-patch+json"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "patchTmpGroupIoTmpVersionTmpKindScale",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Patch"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.api.autoscaling.v1.Scale"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "patch",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     }
    ]
   },
   "/apis/tmpGroup.io/tmpVersion/tmpPlural/{name}/status": {
    "get": {
     "description": "read status of the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "readTmpGroupIoTmpVersionTmpKindStatus",
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "put": {
     "description": "replace status of the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "replaceTmpGroupIoTmpVersionTmpKindStatus",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "201": {
       "description": "Created",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "put",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "patch": {
     "description": "partially update status of the specified tmpKind",
     "consumes": [
      "application/json-patch+json",
      "application/merge-patch+json",
      "application/strategic-merge-patch+json"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "patchTmpGroupIoTmpVersionTmpKindStatus",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Patch"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "patch",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     }
    ]
   }
  },
  "definitions": {
   "io.k8s.api.autoscaling.v1.Scale": {
    "description": "Scale represents a scaling request for a resource.",
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "metadata": {
      "description": "Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata.",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta"
     },
     "spec": {
      "description": "defines the behavior of the scale. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status.",
      "$ref": "#/definitions/io.k8s.api.autoscaling.v1.ScaleSpec"
     },
     "status": {
      "description": "current status of the scale. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status. Read-only.",
      "$ref": "#/definitions/io.k8s.api.autoscaling.v1.ScaleStatus"
     }
    }
   },
   "io.k8s.api.autoscaling.v1.ScaleSpec": {
    "description": "ScaleSpec describes the attributes of a scale subresource.",
    "properties": {
     "replicas": {
      "description": "desired number of instances for the scaled object.",
      "type": "integer",
      "format": "int32"
     }
    }
   },
   "io.k8s.api.autoscaling.v1.ScaleStatus": {
    "description": "ScaleStatus represents the current status of a scale subresource.",
    "required": [
     "replicas"
    ],
    "properties": {
     "replicas": {
      "description": "actual number of observed instances of the scaled object.",
      "type": "integer",
      "format": "int32"
     },
     "selector": {
      "description": "label query over pods that should match the replicas count. This is same as the label selector but in the string format to avoid introspection by clients. The string will be in the same format as the query-param syntax. More info about label selectors: http://kubernetes.io/docs/user-guide/labels#label-selectors",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Initializer": {
    "description": "Initializer is information about an initializer that has not yet completed.",
    "required": [
     "name"
    ],
    "properties": {
     "name": {
      "description": "name of the process that is responsible for initializing this object.",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Initializers": {
    "description": "Initializers tracks the progress of initialization.",
    "required": [
     "pending"
    ],
    "properties": {
     "pending": {
      "description": "Pending is a list of initializers that must execute in order before this object is visible. When the last pending initializer is removed, and no failing result is set, the initializers struct will be set to nil and the object is considered as initialized and visible to all clients.",
      "type": "array",
      "items": {
       "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Initializer"
      },
      "x-kubernetes-patch-merge-key": "name",
      "x-kubernetes-patch-strategy": "merge"
     },
     "result": {
      "description": "If result is set with the Failure field, the object will be persisted to storage and then deleted, ensuring that other clients can observe the deletion.",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.ListMeta": {
    "description": "ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.",
    "properties": {
     "continue": {
      "description": "continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response, unless you have received this token from an error message.",
      "type": "string"
     },
     "resourceVersion": {
      "description": "String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
      "type": "string"
     },
     "selfLink": {
      "description": "selfLink is a URL representing this object. Populated by the system. Read-only.",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta": {
    "description": "ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.",
    "properties": {
     "annotations": {
      "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
      "type": "object",
      "additionalProperties": {
       "type": "string"
      }
     },
     "clusterName": {
      "description": "The name of the cluster which the object belongs to. This is used to distinguish resources with same name and namespace in different clusters. This field is not set anywhere right now and apiserver is going to ignore it if set in create or update request.",
      "type": "string"
     },
     "creationTimestamp": {
      "description": "CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.\n\nPopulated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Time"
     },
     "deletionGracePeriodSeconds": {
      "description": "Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only.",
      "type": "integer",
      "format": "int64"
     },
     "deletionTimestamp": {
      "description": "DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This field is set by the server when a graceful deletion is requested by the user, and is not directly settable by a client. The resource is expected to be deleted (no longer visible from resource lists, and not reachable by name) after the time in this field, once the finalizers list is empty. As long as the finalizers list contains items, deletion is blocked. Once the deletionTimestamp is set, this value may not be unset or be set further into the future, although it may be shortened or the resource may be deleted prior to this time. For example, a user may request that a pod is deleted in 30 seconds. The Kubelet will react by sending a graceful termination signal to the containers in the pod. After that 30 seconds, the Kubelet will send a hard termination signal (SIGKILL) to the container and after cleanup, remove the pod from the API. In the presence of network partitions, this object may still exist after this timestamp, until an administrator or automated process can determine the resource is fully terminated. If not set, graceful deletion of the object has not been requested.\n\nPopulated by the system when a graceful deletion is requested. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Time"
     },
     "finalizers": {
      "description": "Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed.",
      "type": "array",
      "items": {
       "type": "string"
      },
      "x-kubernetes-patch-strategy": "merge"
     },
     "generateName": {
      "description": "GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server.\n\nIf this field is specified and the generated name exists, the server will NOT return a 409 - instead, it will either return 201 Created or 500 with Reason ServerTimeout indicating a unique name could not be found in the time allotted, and the client should retry (optionally after the time indicated in the Retry-After header).\n\nApplied only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency",
      "type": "string"
     },
     "generation": {
      "description": "A sequence number representing a specific generation of the desired state. Populated by the system. Read-only.",
      "type": "integer",
      "format": "int64"
     },
     "initializers": {
      "description": "An initializer is a controller which enforces some system invariant at object creation time. This field is a list of initializers that have not yet acted on this object. If nil or empty, this object has been completely initialized. Otherwise, the object is considered uninitialized and is hidden (in list/watch and get calls) from clients that haven't explicitly asked to observe uninitialized objects.\n\nWhen an object is created, the system will populate this list with the current set of initializers. Only privileged users may set or modify this list. Once it is empty, it may not be modified further by any user.",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Initializers"
     },
     "labels": {
      "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
      "type": "object",
      "additionalProperties": {
       "type": "string"
      }
     },
     "name": {
      "description": "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
      "type": "string"
     },
     "namespace": {
      "description": "Namespace defines the space within each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.\n\nMust be a DNS_LABEL. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces",
      "type": "string"
     },
     "ownerReferences": {
      "description": "List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller.",
      "type": "array",
      "items": {
       "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference"
      },
      "x-kubernetes-patch-merge-key": "uid",
      "x-kubernetes-patch-strategy": "merge"
     },
     "resourceVersion": {
      "description": "An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources.\n\nPopulated by the system. Read-only. Value must be treated as opaque by clients and . More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
      "type": "string"
     },
     "selfLink": {
      "description": "SelfLink is a URL representing this object. Populated by the system. Read-only.",
      "type": "string"
     },
     "uid": {
      "description": "UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.\n\nPopulated by the system. Read-only. More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference": {
    "description": "OwnerReference contains enough information to let you identify an owning object. An owning object must be in the same namespace as the dependent, or be cluster-scoped, so there is no namespace field.",
    "required": [
     "apiVersion",
     "kind",
     "name",
     "uid"
    ],
    "properties": {
     "apiVersion": {
      "description": "API version of the referent.",
      "type": "string"
     },
     "blockOwnerDeletion": {
      "description": "If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. Defaults to false. To set this field, a user needs \"delete\" permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.",
      "type": "boolean"
     },
     "controller": {
      "description": "If true, this reference points to the managing controller.",
      "type": "boolean"
     },
     "kind": {
      "description": "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "name": {
      "description": "Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
      "type": "string"
     },
     "uid": {
      "description": "UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Patch": {
    "description": "Patch is provided to give a concrete name and type to the Kubernetes PATCH request body."
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Status": {
    "description": "Status is a return value for calls that don't return other objects.",
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "code": {
      "description": "Suggested HTTP return code for this status, 0 if not set.",
      "type": "integer",
      "format": "int32"
     },
     "details": {
      "description": "Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.StatusDetails"
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "message": {
      "description": "A human-readable description of the status of this operation.",
      "type": "string"
     },
     "metadata": {
      "description": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ListMeta"
     },
     "reason": {
      "description": "A machine-readable description of why this operation is in the \"Failure\" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.",
      "type": "string"
     },
     "status": {
      "description": "Status of the operation. One of: \"Success\" or \"Failure\". More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.StatusCause": {
    "description": "StatusCause provides more information about an api.Status failure, including cases when multiple errors are encountered.",
    "properties": {
     "field": {
      "description": "The field of the resource that has caused this error, as named by its JSON serialization. May include dot and postfix notation for nested attributes. Arrays are zero-indexed.  Fields may appear more than once in an array of causes due to fields having multiple errors. Optional.\n\nExamples:\n  \"name\" - the field \"name\" on the current resource\n  \"items[0].name\" - the field \"name\" on the first array entry in \"items\"",
      "type": "string"
     },
     "message": {
      "description": "A human-readable description of the cause of the error.  This field may be presented as-is to a reader.",
      "type": "string"
     },
     "reason": {
      "description": "A machine-readable description of the cause of the error. If this value is empty there is no information available.",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.StatusDetails": {
    "description": "StatusDetails is a set of additional properties that MAY be set by the server to provide additional information about a response. The Reason field of a Status object defines what attributes will be set. Clients must ignore fields that do not match the defined type of each attribute, and should assume that any attribute may be empty, invalid, or under defined.",
    "properties": {
     "causes": {
      "description": "The Causes array includes more details associated with the StatusReason failure. Not all StatusReasons may provide detailed causes.",
      "type": "array",
      "items": {
       "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.StatusCause"
      }
     },
     "group": {
      "description": "The group attribute of the resource associated with the status StatusReason.",
      "type": "string"
     },
     "kind": {
      "description": "The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "name": {
      "description": "The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described).",
      "type": "string"
     },
     "retryAfterSeconds": {
      "description": "If specified, the time in seconds before the operation should be retried. Some errors may indicate the client must take an alternate action - for those errors this field may indicate how long to wait before taking the alternate action.",
      "type": "integer",
      "format": "int32"
     },
     "uid": {
      "description": "UID of the resource. (when there is a single resource which can be described). More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Time": {
    "description": "Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.",
    "type": "string",
    "format": "date-time"
   },
   "io.tmpGroup.tmpVersion.tmpKind": {
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "metadata": {
      "description": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta"
     }
    },
    "x-kubernetes-group-version-kind": [
     {
      "group": "tmpGroup.io",
      "kind": "tmpKind",
      "version": "tmpVersion"
     }
    ]
   },
   "io.tmpGroup.tmpVersion.tmpListKind": {
    "description": "tmpListKind is a list of tmpKind",
    "required": [
     "items"
    ],
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "items": {
      "description": "List of tmpPlural. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md",
      "type": "array",
      "items": {
       "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
      }
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "metadata": {
      "description": "ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.",
      "properties": {
       "continue": {
        "description": "continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response, unless you have received this token from an error message.",
        "type": "string"
       },
       "resourceVersion": {
        "description": "String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
        "type": "string"
       },
       "selfLink": {
        "description": "selfLink is a URL representing this object. Populated by the system. Read-only.",
        "type": "string"
       }
      }
     }
    },
    "x-kubernetes-group-version-kind": [
     {
      "group": "tmpGroup.io",
      "kind": "tmpListKind",
      "version": "tmpVersion"
     }
    ]
   }
  }
 }```)

const TemplateWithStatus = []byte(```{
  "swagger": "2.0",
  "info": {
   "title": "Kubernetes CRD",
   "version": "v0.1.0"
  },
  "paths": {
   "/apis/tmpGroup.io/tmpVersion/tmpPlural": {
    "get": {
     "description": "list objects of kind tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "listTmpGroupIoTmpVersionTmpKind",
     "parameters": [
      {
       "uniqueItems": true,
       "type": "string",
       "description": "The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \"next key\".\n\nThis field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.",
       "name": "continue",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
       "name": "fieldSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
       "name": "labelSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.\n\nThe server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.",
       "name": "limit",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.",
       "name": "resourceVersion",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.",
       "name": "timeoutSeconds",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "boolean",
       "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
       "name": "watch",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpListKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "post": {
     "description": "create a tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "createTmpGroupIoTmpVersionTmpKind",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "201": {
       "description": "Created",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "202": {
       "description": "Accepted",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "post",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "delete": {
     "description": "delete collection of tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "deleteTmpGroupIoTmpVersionCollectiontmpKind",
     "parameters": [
      {
       "uniqueItems": true,
       "type": "string",
       "description": "The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \"next key\".\n\nThis field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.",
       "name": "continue",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
       "name": "fieldSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
       "name": "labelSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.\n\nThe server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.",
       "name": "limit",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.",
       "name": "resourceVersion",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.",
       "name": "timeoutSeconds",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "boolean",
       "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
       "name": "watch",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "delete",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "boolean",
      "description": "If true, partially initialized resources are included in the response.",
      "name": "includeUninitialized",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     }
    ]
   },
   "/apis/tmpGroup.io/tmpVersion/tmpPlural/{name}": {
    "get": {
     "description": "read the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "readTmpGroupIoTmpVersionTmpKind",
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "put": {
     "description": "replace the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "replaceTmpGroupIoTmpVersionTmpKind",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "201": {
       "description": "Created",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "put",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "delete": {
     "description": "delete a tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "deleteTmpGroupIoTmpVersionTmpKind",
     "parameters": [
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately.",
       "name": "gracePeriodSeconds",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "boolean",
       "description": "Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \"orphan\" finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both.",
       "name": "orphanDependents",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground.",
       "name": "propagationPolicy",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
       }
      },
      "202": {
       "description": "Accepted",
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "delete",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "patch": {
     "description": "partially update the specified tmpKind",
     "consumes": [
      "application/json-patch+json",
      "application/merge-patch+json",
      "application/strategic-merge-patch+json"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "patchTmpGroupIoTmpVersionTmpKind",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Patch"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "patch",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     }
    ]
   },
   "/apis/tmpGroup.io/tmpVersion/tmpPlural/{name}/status": {
    "get": {
     "description": "read status of the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "readTmpGroupIoTmpVersionTmpKindStatus",
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "put": {
     "description": "replace status of the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "replaceTmpGroupIoTmpVersionTmpKindStatus",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "201": {
       "description": "Created",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "put",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "patch": {
     "description": "partially update status of the specified tmpKind",
     "consumes": [
      "application/json-patch+json",
      "application/merge-patch+json",
      "application/strategic-merge-patch+json"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "patchTmpGroupIoTmpVersionTmpKindStatus",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Patch"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "patch",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     }
    ]
   }
  },
  "definitions": {
   "io.k8s.apimachinery.pkg.apis.meta.v1.ListMeta": {
    "description": "ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.",
    "properties": {
     "continue": {
      "description": "continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response, unless you have received this token from an error message.",
      "type": "string"
     },
     "resourceVersion": {
      "description": "String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
      "type": "string"
     },
     "selfLink": {
      "description": "selfLink is a URL representing this object. Populated by the system. Read-only.",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Patch": {
    "description": "Patch is provided to give a concrete name and type to the Kubernetes PATCH request body."
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Status": {
    "description": "Status is a return value for calls that don't return other objects.",
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "code": {
      "description": "Suggested HTTP return code for this status, 0 if not set.",
      "type": "integer",
      "format": "int32"
     },
     "details": {
      "description": "Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.StatusDetails"
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "message": {
      "description": "A human-readable description of the status of this operation.",
      "type": "string"
     },
     "metadata": {
      "description": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ListMeta"
     },
     "reason": {
      "description": "A machine-readable description of why this operation is in the \"Failure\" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.",
      "type": "string"
     },
     "status": {
      "description": "Status of the operation. One of: \"Success\" or \"Failure\". More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.StatusCause": {
    "description": "StatusCause provides more information about an api.Status failure, including cases when multiple errors are encountered.",
    "properties": {
     "field": {
      "description": "The field of the resource that has caused this error, as named by its JSON serialization. May include dot and postfix notation for nested attributes. Arrays are zero-indexed.  Fields may appear more than once in an array of causes due to fields having multiple errors. Optional.\n\nExamples:\n  \"name\" - the field \"name\" on the current resource\n  \"items[0].name\" - the field \"name\" on the first array entry in \"items\"",
      "type": "string"
     },
     "message": {
      "description": "A human-readable description of the cause of the error.  This field may be presented as-is to a reader.",
      "type": "string"
     },
     "reason": {
      "description": "A machine-readable description of the cause of the error. If this value is empty there is no information available.",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.StatusDetails": {
    "description": "StatusDetails is a set of additional properties that MAY be set by the server to provide additional information about a response. The Reason field of a Status object defines what attributes will be set. Clients must ignore fields that do not match the defined type of each attribute, and should assume that any attribute may be empty, invalid, or under defined.",
    "properties": {
     "causes": {
      "description": "The Causes array includes more details associated with the StatusReason failure. Not all StatusReasons may provide detailed causes.",
      "type": "array",
      "items": {
       "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.StatusCause"
      }
     },
     "group": {
      "description": "The group attribute of the resource associated with the status StatusReason.",
      "type": "string"
     },
     "kind": {
      "description": "The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "name": {
      "description": "The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described).",
      "type": "string"
     },
     "retryAfterSeconds": {
      "description": "If specified, the time in seconds before the operation should be retried. Some errors may indicate the client must take an alternate action - for those errors this field may indicate how long to wait before taking the alternate action.",
      "type": "integer",
      "format": "int32"
     },
     "uid": {
      "description": "UID of the resource. (when there is a single resource which can be described). More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
      "type": "string"
     }
    }
   },
   "io.tmpGroup.tmpVersion.tmpKind": {
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "metadata": {
      "description": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta"
     }
    },
    "x-kubernetes-group-version-kind": [
     {
      "group": "tmpGroup.io",
      "kind": "tmpKind",
      "version": "tmpVersion"
     }
    ]
   },
   "io.tmpGroup.tmpVersion.tmpListKind": {
    "description": "tmpListKind is a list of tmpKind",
    "required": [
     "items"
    ],
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "items": {
      "description": "List of tmpPlural. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md",
      "type": "array",
      "items": {
       "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
      }
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "metadata": {
      "description": "ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.",
      "properties": {
       "continue": {
        "description": "continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response, unless you have received this token from an error message.",
        "type": "string"
       },
       "resourceVersion": {
        "description": "String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
        "type": "string"
       },
       "selfLink": {
        "description": "selfLink is a URL representing this object. Populated by the system. Read-only.",
        "type": "string"
       }
      }
     }
    },
    "x-kubernetes-group-version-kind": [
     {
      "group": "tmpGroup.io",
      "kind": "tmpListKind",
      "version": "tmpVersion"
     }
    ]
   }
  }
 }```)

const TemplateWithScale = []byte(```{
  "swagger": "2.0",
  "info": {
   "title": "Kubernetes CRD",
   "version": "v0.1.0"
  },
  "paths": {
   "/apis/tmpGroup.io/tmpVersion/tmpPlural": {
    "get": {
     "description": "list objects of kind tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "listTmpGroupIoTmpVersionTmpKind",
     "parameters": [
      {
       "uniqueItems": true,
       "type": "string",
       "description": "The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \"next key\".\n\nThis field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.",
       "name": "continue",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
       "name": "fieldSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
       "name": "labelSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.\n\nThe server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.",
       "name": "limit",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.",
       "name": "resourceVersion",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.",
       "name": "timeoutSeconds",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "boolean",
       "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
       "name": "watch",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpListKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "post": {
     "description": "create a tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "createTmpGroupIoTmpVersionTmpKind",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "201": {
       "description": "Created",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "202": {
       "description": "Accepted",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "post",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "delete": {
     "description": "delete collection of tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "deleteTmpGroupIoTmpVersionCollectiontmpKind",
     "parameters": [
      {
       "uniqueItems": true,
       "type": "string",
       "description": "The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \"next key\".\n\nThis field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.",
       "name": "continue",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
       "name": "fieldSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
       "name": "labelSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.\n\nThe server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.",
       "name": "limit",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.",
       "name": "resourceVersion",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.",
       "name": "timeoutSeconds",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "boolean",
       "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
       "name": "watch",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "delete",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "boolean",
      "description": "If true, partially initialized resources are included in the response.",
      "name": "includeUninitialized",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     }
    ]
   },
   "/apis/tmpGroup.io/tmpVersion/tmpPlural/{name}": {
    "get": {
     "description": "read the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "readTmpGroupIoTmpVersionTmpKind",
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "put": {
     "description": "replace the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "replaceTmpGroupIoTmpVersionTmpKind",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "201": {
       "description": "Created",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "put",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "delete": {
     "description": "delete a tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "deleteTmpGroupIoTmpVersionTmpKind",
     "parameters": [
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately.",
       "name": "gracePeriodSeconds",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "boolean",
       "description": "Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \"orphan\" finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both.",
       "name": "orphanDependents",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground.",
       "name": "propagationPolicy",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
       }
      },
      "202": {
       "description": "Accepted",
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "delete",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "patch": {
     "description": "partially update the specified tmpKind",
     "consumes": [
      "application/json-patch+json",
      "application/merge-patch+json",
      "application/strategic-merge-patch+json"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "patchTmpGroupIoTmpVersionTmpKind",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Patch"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "patch",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     }
    ]
   },
   "/apis/tmpGroup.io/tmpVersion/tmpPlural/{name}/scale": {
    "get": {
     "description": "read scale of the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "readTmpGroupIoTmpVersionTmpKindScale",
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.api.autoscaling.v1.Scale"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "put": {
     "description": "replace scale of the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "replaceTmpGroupIoTmpVersionTmpKindScale",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.k8s.api.autoscaling.v1.Scale"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.api.autoscaling.v1.Scale"
       }
      },
      "201": {
       "description": "Created",
       "schema": {
        "$ref": "#/definitions/io.k8s.api.autoscaling.v1.Scale"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "put",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "patch": {
     "description": "partially update scale of the specified tmpKind",
     "consumes": [
      "application/json-patch+json",
      "application/merge-patch+json",
      "application/strategic-merge-patch+json"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "patchTmpGroupIoTmpVersionTmpKindScale",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Patch"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.api.autoscaling.v1.Scale"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "patch",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     }
    ]
   }
  },
  "definitions": {
   "io.k8s.api.autoscaling.v1.Scale": {
    "description": "Scale represents a scaling request for a resource.",
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "metadata": {
      "description": "Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata.",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta"
     },
     "spec": {
      "description": "defines the behavior of the scale. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status.",
      "$ref": "#/definitions/io.k8s.api.autoscaling.v1.ScaleSpec"
     },
     "status": {
      "description": "current status of the scale. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status. Read-only.",
      "$ref": "#/definitions/io.k8s.api.autoscaling.v1.ScaleStatus"
     }
    }
   },
   "io.k8s.api.autoscaling.v1.ScaleSpec": {
    "description": "ScaleSpec describes the attributes of a scale subresource.",
    "properties": {
     "replicas": {
      "description": "desired number of instances for the scaled object.",
      "type": "integer",
      "format": "int32"
     }
    }
   },
   "io.k8s.api.autoscaling.v1.ScaleStatus": {
    "description": "ScaleStatus represents the current status of a scale subresource.",
    "required": [
     "replicas"
    ],
    "properties": {
     "replicas": {
      "description": "actual number of observed instances of the scaled object.",
      "type": "integer",
      "format": "int32"
     },
     "selector": {
      "description": "label query over pods that should match the replicas count. This is same as the label selector but in the string format to avoid introspection by clients. The string will be in the same format as the query-param syntax. More info about label selectors: http://kubernetes.io/docs/user-guide/labels#label-selectors",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Initializer": {
    "description": "Initializer is information about an initializer that has not yet completed.",
    "required": [
     "name"
    ],
    "properties": {
     "name": {
      "description": "name of the process that is responsible for initializing this object.",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Initializers": {
    "description": "Initializers tracks the progress of initialization.",
    "required": [
     "pending"
    ],
    "properties": {
     "pending": {
      "description": "Pending is a list of initializers that must execute in order before this object is visible. When the last pending initializer is removed, and no failing result is set, the initializers struct will be set to nil and the object is considered as initialized and visible to all clients.",
      "type": "array",
      "items": {
       "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Initializer"
      },
      "x-kubernetes-patch-merge-key": "name",
      "x-kubernetes-patch-strategy": "merge"
     },
     "result": {
      "description": "If result is set with the Failure field, the object will be persisted to storage and then deleted, ensuring that other clients can observe the deletion.",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.ListMeta": {
    "description": "ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.",
    "properties": {
     "continue": {
      "description": "continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response, unless you have received this token from an error message.",
      "type": "string"
     },
     "resourceVersion": {
      "description": "String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
      "type": "string"
     },
     "selfLink": {
      "description": "selfLink is a URL representing this object. Populated by the system. Read-only.",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta": {
    "description": "ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.",
    "properties": {
     "annotations": {
      "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
      "type": "object",
      "additionalProperties": {
       "type": "string"
      }
     },
     "clusterName": {
      "description": "The name of the cluster which the object belongs to. This is used to distinguish resources with same name and namespace in different clusters. This field is not set anywhere right now and apiserver is going to ignore it if set in create or update request.",
      "type": "string"
     },
     "creationTimestamp": {
      "description": "CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.\n\nPopulated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Time"
     },
     "deletionGracePeriodSeconds": {
      "description": "Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only.",
      "type": "integer",
      "format": "int64"
     },
     "deletionTimestamp": {
      "description": "DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This field is set by the server when a graceful deletion is requested by the user, and is not directly settable by a client. The resource is expected to be deleted (no longer visible from resource lists, and not reachable by name) after the time in this field, once the finalizers list is empty. As long as the finalizers list contains items, deletion is blocked. Once the deletionTimestamp is set, this value may not be unset or be set further into the future, although it may be shortened or the resource may be deleted prior to this time. For example, a user may request that a pod is deleted in 30 seconds. The Kubelet will react by sending a graceful termination signal to the containers in the pod. After that 30 seconds, the Kubelet will send a hard termination signal (SIGKILL) to the container and after cleanup, remove the pod from the API. In the presence of network partitions, this object may still exist after this timestamp, until an administrator or automated process can determine the resource is fully terminated. If not set, graceful deletion of the object has not been requested.\n\nPopulated by the system when a graceful deletion is requested. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Time"
     },
     "finalizers": {
      "description": "Must be empty before the object is deleted from the registry. Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed.",
      "type": "array",
      "items": {
       "type": "string"
      },
      "x-kubernetes-patch-strategy": "merge"
     },
     "generateName": {
      "description": "GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server.\n\nIf this field is specified and the generated name exists, the server will NOT return a 409 - instead, it will either return 201 Created or 500 with Reason ServerTimeout indicating a unique name could not be found in the time allotted, and the client should retry (optionally after the time indicated in the Retry-After header).\n\nApplied only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency",
      "type": "string"
     },
     "generation": {
      "description": "A sequence number representing a specific generation of the desired state. Populated by the system. Read-only.",
      "type": "integer",
      "format": "int64"
     },
     "initializers": {
      "description": "An initializer is a controller which enforces some system invariant at object creation time. This field is a list of initializers that have not yet acted on this object. If nil or empty, this object has been completely initialized. Otherwise, the object is considered uninitialized and is hidden (in list/watch and get calls) from clients that haven't explicitly asked to observe uninitialized objects.\n\nWhen an object is created, the system will populate this list with the current set of initializers. Only privileged users may set or modify this list. Once it is empty, it may not be modified further by any user.",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Initializers"
     },
     "labels": {
      "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
      "type": "object",
      "additionalProperties": {
       "type": "string"
      }
     },
     "name": {
      "description": "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
      "type": "string"
     },
     "namespace": {
      "description": "Namespace defines the space within each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.\n\nMust be a DNS_LABEL. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces",
      "type": "string"
     },
     "ownerReferences": {
      "description": "List of objects depended by this object. If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller.",
      "type": "array",
      "items": {
       "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference"
      },
      "x-kubernetes-patch-merge-key": "uid",
      "x-kubernetes-patch-strategy": "merge"
     },
     "resourceVersion": {
      "description": "An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources.\n\nPopulated by the system. Read-only. Value must be treated as opaque by clients and . More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
      "type": "string"
     },
     "selfLink": {
      "description": "SelfLink is a URL representing this object. Populated by the system. Read-only.",
      "type": "string"
     },
     "uid": {
      "description": "UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.\n\nPopulated by the system. Read-only. More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference": {
    "description": "OwnerReference contains enough information to let you identify an owning object. An owning object must be in the same namespace as the dependent, or be cluster-scoped, so there is no namespace field.",
    "required": [
     "apiVersion",
     "kind",
     "name",
     "uid"
    ],
    "properties": {
     "apiVersion": {
      "description": "API version of the referent.",
      "type": "string"
     },
     "blockOwnerDeletion": {
      "description": "If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. Defaults to false. To set this field, a user needs \"delete\" permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.",
      "type": "boolean"
     },
     "controller": {
      "description": "If true, this reference points to the managing controller.",
      "type": "boolean"
     },
     "kind": {
      "description": "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "name": {
      "description": "Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
      "type": "string"
     },
     "uid": {
      "description": "UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Patch": {
    "description": "Patch is provided to give a concrete name and type to the Kubernetes PATCH request body."
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Status": {
    "description": "Status is a return value for calls that don't return other objects.",
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "code": {
      "description": "Suggested HTTP return code for this status, 0 if not set.",
      "type": "integer",
      "format": "int32"
     },
     "details": {
      "description": "Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.StatusDetails"
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "message": {
      "description": "A human-readable description of the status of this operation.",
      "type": "string"
     },
     "metadata": {
      "description": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ListMeta"
     },
     "reason": {
      "description": "A machine-readable description of why this operation is in the \"Failure\" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.",
      "type": "string"
     },
     "status": {
      "description": "Status of the operation. One of: \"Success\" or \"Failure\". More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.StatusCause": {
    "description": "StatusCause provides more information about an api.Status failure, including cases when multiple errors are encountered.",
    "properties": {
     "field": {
      "description": "The field of the resource that has caused this error, as named by its JSON serialization. May include dot and postfix notation for nested attributes. Arrays are zero-indexed.  Fields may appear more than once in an array of causes due to fields having multiple errors. Optional.\n\nExamples:\n  \"name\" - the field \"name\" on the current resource\n  \"items[0].name\" - the field \"name\" on the first array entry in \"items\"",
      "type": "string"
     },
     "message": {
      "description": "A human-readable description of the cause of the error.  This field may be presented as-is to a reader.",
      "type": "string"
     },
     "reason": {
      "description": "A machine-readable description of the cause of the error. If this value is empty there is no information available.",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.StatusDetails": {
    "description": "StatusDetails is a set of additional properties that MAY be set by the server to provide additional information about a response. The Reason field of a Status object defines what attributes will be set. Clients must ignore fields that do not match the defined type of each attribute, and should assume that any attribute may be empty, invalid, or under defined.",
    "properties": {
     "causes": {
      "description": "The Causes array includes more details associated with the StatusReason failure. Not all StatusReasons may provide detailed causes.",
      "type": "array",
      "items": {
       "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.StatusCause"
      }
     },
     "group": {
      "description": "The group attribute of the resource associated with the status StatusReason.",
      "type": "string"
     },
     "kind": {
      "description": "The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "name": {
      "description": "The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described).",
      "type": "string"
     },
     "retryAfterSeconds": {
      "description": "If specified, the time in seconds before the operation should be retried. Some errors may indicate the client must take an alternate action - for those errors this field may indicate how long to wait before taking the alternate action.",
      "type": "integer",
      "format": "int32"
     },
     "uid": {
      "description": "UID of the resource. (when there is a single resource which can be described). More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Time": {
    "description": "Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.",
    "type": "string",
    "format": "date-time"
   },
   "io.tmpGroup.tmpVersion.tmpKind": {
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "metadata": {
      "description": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta"
     }
    },
    "x-kubernetes-group-version-kind": [
     {
      "group": "tmpGroup.io",
      "kind": "tmpKind",
      "version": "tmpVersion"
     }
    ]
   },
   "io.tmpGroup.tmpVersion.tmpListKind": {
    "description": "tmpListKind is a list of tmpKind",
    "required": [
     "items"
    ],
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "items": {
      "description": "List of tmpPlural. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md",
      "type": "array",
      "items": {
       "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
      }
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "metadata": {
      "description": "ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.",
      "properties": {
       "continue": {
        "description": "continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response, unless you have received this token from an error message.",
        "type": "string"
       },
       "resourceVersion": {
        "description": "String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
        "type": "string"
       },
       "selfLink": {
        "description": "selfLink is a URL representing this object. Populated by the system. Read-only.",
        "type": "string"
       }
      }
     }
    },
    "x-kubernetes-group-version-kind": [
     {
      "group": "tmpGroup.io",
      "kind": "tmpListKind",
      "version": "tmpVersion"
     }
    ]
   }
  }
 }```)

const Template = []byte(```{
  "swagger": "2.0",
  "info": {
   "title": "Kubernetes CRD",
   "version": "v0.1.0"
  },
  "paths": {
   "/apis/tmpGroup.io/tmpVersion/tmpPlural": {
    "get": {
     "description": "list objects of kind tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "listTmpGroupIoTmpVersionTmpKind",
     "parameters": [
      {
       "uniqueItems": true,
       "type": "string",
       "description": "The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \"next key\".\n\nThis field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.",
       "name": "continue",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
       "name": "fieldSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
       "name": "labelSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.\n\nThe server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.",
       "name": "limit",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.",
       "name": "resourceVersion",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.",
       "name": "timeoutSeconds",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "boolean",
       "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
       "name": "watch",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpListKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "post": {
     "description": "create a tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "createTmpGroupIoTmpVersionTmpKind",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "201": {
       "description": "Created",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "202": {
       "description": "Accepted",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "post",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "delete": {
     "description": "delete collection of tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "deleteTmpGroupIoTmpVersionCollectiontmpKind",
     "parameters": [
      {
       "uniqueItems": true,
       "type": "string",
       "description": "The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \"next key\".\n\nThis field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.",
       "name": "continue",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their fields. Defaults to everything.",
       "name": "fieldSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "A selector to restrict the list of returned objects by their labels. Defaults to everything.",
       "name": "labelSelector",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.\n\nThe server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.",
       "name": "limit",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.",
       "name": "resourceVersion",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.",
       "name": "timeoutSeconds",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "boolean",
       "description": "Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.",
       "name": "watch",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "delete",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "boolean",
      "description": "If true, partially initialized resources are included in the response.",
      "name": "includeUninitialized",
      "in": "query"
     },
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     }
    ]
   },
   "/apis/tmpGroup.io/tmpVersion/tmpPlural/{name}": {
    "get": {
     "description": "read the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "readTmpGroupIoTmpVersionTmpKind",
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "get",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "put": {
     "description": "replace the specified tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "replaceTmpGroupIoTmpVersionTmpKind",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "201": {
       "description": "Created",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "put",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "delete": {
     "description": "delete a tmpKind",
     "consumes": [
      "*/*"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "deleteTmpGroupIoTmpVersionTmpKind",
     "parameters": [
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "integer",
       "description": "The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately.",
       "name": "gracePeriodSeconds",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "boolean",
       "description": "Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \"orphan\" finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both.",
       "name": "orphanDependents",
       "in": "query"
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground.",
       "name": "propagationPolicy",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
       }
      },
      "202": {
       "description": "Accepted",
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Status"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "delete",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "patch": {
     "description": "partially update the specified tmpKind",
     "consumes": [
      "application/json-patch+json",
      "application/merge-patch+json",
      "application/strategic-merge-patch+json"
     ],
     "produces": [
      "application/json",
      "application/yaml"
     ],
     "schemes": [
      "https"
     ],
     "tags": [
      "tmpGroupIo_tmpVersion"
     ],
     "operationId": "patchTmpGroupIoTmpVersionTmpKind",
     "parameters": [
      {
       "name": "body",
       "in": "body",
       "required": true,
       "schema": {
        "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Patch"
       }
      },
      {
       "uniqueItems": true,
       "type": "string",
       "description": "When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed",
       "name": "dryRun",
       "in": "query"
      }
     ],
     "responses": {
      "200": {
       "description": "OK",
       "schema": {
        "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
       }
      },
      "401": {
       "description": "Unauthorized"
      }
     },
     "x-kubernetes-action": "patch",
     "x-kubernetes-group-version-kind": {
      "group": "tmpGroup.io",
      "version": "tmpVersion",
      "kind": "tmpKind"
     }
    },
    "parameters": [
     {
      "uniqueItems": true,
      "type": "string",
      "description": "If 'true', then the output is pretty printed.",
      "name": "pretty",
      "in": "query"
     }
    ]
   }
  },
  "definitions": {
   "io.k8s.apimachinery.pkg.apis.meta.v1.ListMeta": {
    "description": "ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.",
    "properties": {
     "continue": {
      "description": "continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response, unless you have received this token from an error message.",
      "type": "string"
     },
     "resourceVersion": {
      "description": "String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
      "type": "string"
     },
     "selfLink": {
      "description": "selfLink is a URL representing this object. Populated by the system. Read-only.",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Patch": {
    "description": "Patch is provided to give a concrete name and type to the Kubernetes PATCH request body."
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.Status": {
    "description": "Status is a return value for calls that don't return other objects.",
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "code": {
      "description": "Suggested HTTP return code for this status, 0 if not set.",
      "type": "integer",
      "format": "int32"
     },
     "details": {
      "description": "Extended data associated with the reason.  Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.StatusDetails"
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "message": {
      "description": "A human-readable description of the status of this operation.",
      "type": "string"
     },
     "metadata": {
      "description": "Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ListMeta"
     },
     "reason": {
      "description": "A machine-readable description of why this operation is in the \"Failure\" status. If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.",
      "type": "string"
     },
     "status": {
      "description": "Status of the operation. One of: \"Success\" or \"Failure\". More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.StatusCause": {
    "description": "StatusCause provides more information about an api.Status failure, including cases when multiple errors are encountered.",
    "properties": {
     "field": {
      "description": "The field of the resource that has caused this error, as named by its JSON serialization. May include dot and postfix notation for nested attributes. Arrays are zero-indexed.  Fields may appear more than once in an array of causes due to fields having multiple errors. Optional.\n\nExamples:\n  \"name\" - the field \"name\" on the current resource\n  \"items[0].name\" - the field \"name\" on the first array entry in \"items\"",
      "type": "string"
     },
     "message": {
      "description": "A human-readable description of the cause of the error.  This field may be presented as-is to a reader.",
      "type": "string"
     },
     "reason": {
      "description": "A machine-readable description of the cause of the error. If this value is empty there is no information available.",
      "type": "string"
     }
    }
   },
   "io.k8s.apimachinery.pkg.apis.meta.v1.StatusDetails": {
    "description": "StatusDetails is a set of additional properties that MAY be set by the server to provide additional information about a response. The Reason field of a Status object defines what attributes will be set. Clients must ignore fields that do not match the defined type of each attribute, and should assume that any attribute may be empty, invalid, or under defined.",
    "properties": {
     "causes": {
      "description": "The Causes array includes more details associated with the StatusReason failure. Not all StatusReasons may provide detailed causes.",
      "type": "array",
      "items": {
       "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.StatusCause"
      }
     },
     "group": {
      "description": "The group attribute of the resource associated with the status StatusReason.",
      "type": "string"
     },
     "kind": {
      "description": "The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "name": {
      "description": "The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described).",
      "type": "string"
     },
     "retryAfterSeconds": {
      "description": "If specified, the time in seconds before the operation should be retried. Some errors may indicate the client must take an alternate action - for those errors this field may indicate how long to wait before taking the alternate action.",
      "type": "integer",
      "format": "int32"
     },
     "uid": {
      "description": "UID of the resource. (when there is a single resource which can be described). More info: http://kubernetes.io/docs/user-guide/identifiers#uids",
      "type": "string"
     }
    }
   },
   "io.tmpGroup.tmpVersion.tmpKind": {
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "metadata": {
      "description": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
      "$ref": "#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta"
     }
    },
    "x-kubernetes-group-version-kind": [
     {
      "group": "tmpGroup.io",
      "kind": "tmpKind",
      "version": "tmpVersion"
     }
    ]
   },
   "io.tmpGroup.tmpVersion.tmpListKind": {
    "description": "tmpListKind is a list of tmpKind",
    "required": [
     "items"
    ],
    "properties": {
     "apiVersion": {
      "description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
      "type": "string"
     },
     "items": {
      "description": "List of tmpPlural. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md",
      "type": "array",
      "items": {
       "$ref": "#/definitions/io.tmpGroup.tmpVersion.tmpKind"
      }
     },
     "kind": {
      "description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
      "type": "string"
     },
     "metadata": {
      "description": "ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.",
      "properties": {
       "continue": {
        "description": "continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response, unless you have received this token from an error message.",
        "type": "string"
       },
       "resourceVersion": {
        "description": "String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency",
        "type": "string"
       },
       "selfLink": {
        "description": "selfLink is a URL representing this object. Populated by the system. Read-only.",
        "type": "string"
       }
      }
     }
    },
    "x-kubernetes-group-version-kind": [
     {
      "group": "tmpGroup.io",
      "kind": "tmpListKind",
      "version": "tmpVersion"
     }
    ]
   }
  }
 }```)
