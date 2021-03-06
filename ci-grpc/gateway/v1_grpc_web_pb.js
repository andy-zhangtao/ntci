/**
 * @fileoverview gRPC-Web generated client stub for 
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = require('./v1_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.GateWayRpcClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.GateWayRpcPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.BuildRequest,
 *   !proto.JobInfo>}
 */
const methodDescriptor_GateWayRpc_GetBuild = new grpc.web.MethodDescriptor(
  '/GateWayRpc/GetBuild',
  grpc.web.MethodType.UNARY,
  proto.BuildRequest,
  proto.JobInfo,
  /** @param {!proto.BuildRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.JobInfo.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.BuildRequest,
 *   !proto.JobInfo>}
 */
const methodInfo_GateWayRpc_GetBuild = new grpc.web.AbstractClientBase.MethodInfo(
  proto.JobInfo,
  /** @param {!proto.BuildRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.JobInfo.deserializeBinary
);


/**
 * @param {!proto.BuildRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.JobInfo)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.JobInfo>|undefined}
 *     The XHR Node Readable Stream
 */
proto.GateWayRpcClient.prototype.getBuild =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/GateWayRpc/GetBuild',
      request,
      metadata || {},
      methodDescriptor_GateWayRpc_GetBuild,
      callback);
};


/**
 * @param {!proto.BuildRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.JobInfo>}
 *     A native promise that resolves to the response
 */
proto.GateWayRpcPromiseClient.prototype.getBuild =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/GateWayRpc/GetBuild',
      request,
      metadata || {},
      methodDescriptor_GateWayRpc_GetBuild);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.Builder,
 *   !proto.Reply>}
 */
const methodDescriptor_GateWayRpc_JobStatus = new grpc.web.MethodDescriptor(
  '/GateWayRpc/JobStatus',
  grpc.web.MethodType.UNARY,
  proto.Builder,
  proto.Reply,
  /** @param {!proto.Builder} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.Reply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.Builder,
 *   !proto.Reply>}
 */
const methodInfo_GateWayRpc_JobStatus = new grpc.web.AbstractClientBase.MethodInfo(
  proto.Reply,
  /** @param {!proto.Builder} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.Reply.deserializeBinary
);


/**
 * @param {!proto.Builder} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.Reply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.Reply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.GateWayRpcClient.prototype.jobStatus =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/GateWayRpc/JobStatus',
      request,
      metadata || {},
      methodDescriptor_GateWayRpc_JobStatus,
      callback);
};


/**
 * @param {!proto.Builder} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.Reply>}
 *     A native promise that resolves to the response
 */
proto.GateWayRpcPromiseClient.prototype.jobStatus =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/GateWayRpc/JobStatus',
      request,
      metadata || {},
      methodDescriptor_GateWayRpc_JobStatus);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.Builder,
 *   !proto.Reply>}
 */
const methodDescriptor_GateWayRpc_RestartJob = new grpc.web.MethodDescriptor(
  '/GateWayRpc/RestartJob',
  grpc.web.MethodType.UNARY,
  proto.Builder,
  proto.Reply,
  /** @param {!proto.Builder} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.Reply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.Builder,
 *   !proto.Reply>}
 */
const methodInfo_GateWayRpc_RestartJob = new grpc.web.AbstractClientBase.MethodInfo(
  proto.Reply,
  /** @param {!proto.Builder} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.Reply.deserializeBinary
);


/**
 * @param {!proto.Builder} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.Reply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.Reply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.GateWayRpcClient.prototype.restartJob =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/GateWayRpc/RestartJob',
      request,
      metadata || {},
      methodDescriptor_GateWayRpc_RestartJob,
      callback);
};


/**
 * @param {!proto.Builder} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.Reply>}
 *     A native promise that resolves to the response
 */
proto.GateWayRpcPromiseClient.prototype.restartJob =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/GateWayRpc/RestartJob',
      request,
      metadata || {},
      methodDescriptor_GateWayRpc_RestartJob);
};


module.exports = proto;

