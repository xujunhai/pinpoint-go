package trace

/**
 *  sandbox is from 900 to 999. These values will not be assigned to anything.
 *
 * <table>
 * <tr><td>-1</td><td>args[0]</td></tr>
 * <tr><td>-2</td><td>args[1]</td></tr>
 * <tr><td>-3</td><td>args[2]</td></tr>
 * <tr><td>-4</td><td>args[3]</td></tr>
 * <tr><td>-5</td><td>args[4]</td></tr>
 * <tr><td>-6</td><td>args[5]</td></tr>
 * <tr><td>-7</td><td>args[6]</td></tr>
 * <tr><td>-8</td><td>args[7]</td></tr>
 * <tr><td>-9</td><td>args[8]</td></tr>
 * <tr><td>-10</td><td>args[9]</td></tr>
 * <tr><td>-11</td><td>args[N]</td></tr>
 * <tr><td>-30</td><td>cached_args[0]</td></tr>
 * <tr><td>-31</td><td>cached_args[1]</td></tr>
 * <tr><td>-32</td><td>cached_args[2]</td></tr>
 * <tr><td>-33</td><td>cached_args[3]</td></tr>
 * <tr><td>-34</td><td>cached_args[4]</td></tr>
 * <tr><td>-35</td><td>cached_args[5]</td></tr>
 * <tr><td>-36</td><td>cached_args[6]</td></tr>
 * <tr><td>-37</td><td>cached_args[7]</td></tr>
 * <tr><td>-38</td><td>cached_args[8]</td></tr>
 * <tr><td>-39</td><td>cached_args[9]</td></tr>
 * <tr><td>-40</td><td>cached_args[N]</td></tr>
 * <tr><td>-50</td><td>Exception</td></tr>
 * <tr><td>-51</td><td>ExceptionClass</td></tr>
 * <tr><td>-100</td><td>Asynchronous Invocation</td></tr>
 * <tr><td>-9999</td><td>UNKNOWN</td></tr>
 *
 * <tr><td>12</td><td>API</td></tr>
 * <tr><td>13</td><td>API_METADATA</td></tr>
 * <tr><td>14</td><td>RETURN_DATA</td></tr>
 * <tr><td>15</td><td><i>RESERVED</i></td></tr>
 * <tr><td>16</td><td><i>RESERVED</i></td></tr>
 * <tr><td>17</td><td><i>RESERVED</i></td></tr>
 * <tr><td>20</td><td>SQL-ID</td></tr>
 * <tr><td>21</td><td>SQL</td></tr>
 * <tr><td>22</td><td>SQL-METADATA</td></tr>
 * <tr><td>23</td><td>SQL-PARAM</td></tr>
 * <tr><td>24</td><td>SQL-BindValue</td></tr>
 * <tr><td>30</td><td>STRING_ID</td></tr>
 * <tr><td>40</td><td>http.url</td></tr>
 * <tr><td>41</td><td>http.param</td></tr>
 * <tr><td>42</td><td>http.entity</td></tr>
 * <tr><td>45</td><td>http.cookie</td></tr>
 * <tr><td>46</td><td>http.status.code</td></tr>
 * <tr><td>48</td><td>http.internal.display</td></tr>
 * <tr><td>49</td><td>http.io</td></tr>
 * <tr><td>50</td><td>arcus.command</td></tr>
 * <tr><td>60</td><td><i>RESERVED</i></td></tr>
 * <tr><td>61</td><td><i>RESERVED</i></td></tr>
 * <tr><td>62</td><td><i>RESERVED</i></td></tr>
 * <tr><td>70</td><td><i>RESERVED</i></td></tr>
 * <tr><td>71</td><td><i>RESERVED</i></td></tr>
 * <tr><td>72</td><td><i>RESERVED</i></td></tr>
 * <tr><td>73</td><td><i>RESERVED</i></td></tr>
 * <tr><td>80</td><td>thrift.url</td></tr>
 * <tr><td>81</td><td>thrift.args</td></tr>
 * <tr><td>82</td><td>thrift.result</td></tr>
 * <tr><td>90</td><td>dubbo.args</td></tr>
 * <tr><td>91</td><td>dubbo.result</td></tr>
 * <tr><td>110</td><td></td>hystrix.command</tr>
 * <tr><td>111</td><td></td>hystrix.command.execution</tr>
 * <tr><td>112</td><td></td>hystrix.command.fallback.cause</tr>
 * <tr><td>113</td><td></td>hystrix.command.exception</tr>
 * <tr><td>115</td><td></td>hystrix.command.key</tr>
 * <tr><td>116</td><td></td>hystrix.command.group.key</tr>
 * <tr><td>117</td><td></td>hystrix.thread.pool.key</tr>
 * <tr><td>118</td><td></td>hystrix.collapser.key</tr>
 * <tr><td>120</td><td>netty.address</td></tr>
 * <tr><td>130</td><td>rabbitmq.properties</td></tr>
 * <tr><td>131</td><td>rabbitmq.body</td></tr>
 * <tr><td>132</td><td>rabbitmq.properties</td></tr>
 * <tr><td>133</td><td>rabbitmq.body</td></tr>
 * <tr><td>150</td><td>mongo.json.data</td></tr>
 * <tr><td>151</td><td>mongo.collection.info</td></tr>
 * <tr><td>152</td><td>mongo.collection.option</td></tr>
 * <tr><td>153</td><td>mongo.json</td></tr>
 * <tr><td>154</td><td>mongo.json.bindvalue</td></tr>
 * <tr><td>160</td><td>grpc.status</td></tr>
 * <tr><td>171</td><td>es.args</td></tr>
 * <tr><td>172</td><td>es.url</td></tr>
 * <tr><td>173</td><td>es.dsl</td></tr>
 * <tr><td>174</td><td>es.action</td></tr>
 * <tr><td>175</td><td>es.responseHandle</td></tr>
 * <tr><td>176</td><td>es.version</td></tr>
 * <tr><td>180</td><td>process.command</td></tr>
 * <tr><td>181</td><td>process.pid</td></tr>
 *
 * <tr><td><s>200</s></td><td><s>cxf.operation</s></td></tr>
 * <tr><td><s>201</s></td><td><s>cxf.args</s></td></tr>
 * <tr><td>203</td><td>cxf.address</td></tr>
 * <tr><td>204</td><td>cxf.response.code</td></tr>
 * <tr><td>205</td><td>cxf.encoding</td></tr>
 * <tr><td>206</td><td>cxf.http.method</td></tr>
 * <tr><td>207</td><td>cxf.content.type</td></tr>
 * <tr><td>208</td><td>cxf.headers</td></tr>
 * <tr><td>209</td><td>cxf.messages</td></tr>
 * <tr><td>210</td><td>cxf.payload</td></tr>
 * <tr><td>300</td><td>PROXY_HTTP_HEADER</td></tr>
 * <tr><td>310</td><td>REDIS.IO</td></tr>
 * <tr><td>320</td><td>hbase.client.params</td></tr>
 * <tr><td>923</td><td>marker.message</td></tr>
 * <tr><td>9000</td><td>gson.json.length</td></tr>
 * <tr><td>9001</td><td>jackson.json.length</td></tr>
 * <tr><td>9002</td><td>json-lib.json.length</td></tr>
 * <tr><td>9003</td><td>fastjson.json.length</td></tr>
 * <tr><td>10015</td><td>API-TAG</td></tr>
 * <tr><td>10000010</td><td>API-METADATA-ERROR</td></tr>
 * <tr><td>10000011</td><td>API-METADATA-AGENT-INFO-NOT-FOUND</td></tr>
 * <tr><td>10000012</td><td>API-METADATA-IDENTIFIER-CHECK_ERROR</td></tr>
 * <tr><td>10000013</td><td>API-METADATA-NOT-FOUND</td></tr>
 * <tr><td>10000014</td><td>API-METADATA-DID-COLLSION</td></tr>
 * </table>
 *
 * @author netspider
 * @author emeroad
 * @author Jongho Moon
 */

const (
	AkApi         int32 = 12
	AkApiMetadata int32 = 13
	AkReturnData  int32 = 14
	AkApiTag      int32 = 10015

	AkErrApiMetadataError                int32 = 10000010
	AkErrApiMetadataAgentInfoNotFound    int32 = 10000011
	AkErrApiMetadataIdentifierCheckError int32 = 10000012
	AkErrApiMetadataNotFound             int32 = 10000013
	AkErrApiMetadataDidCollsion          int32 = 10000014

	AkSqlId        int32 = 20
	AkSql          int32 = 21
	AkSqlMetadata  int32 = 22
	AkSqlParam     int32 = 23
	AkSqlBindvalue int32 = 24

	AkStringId int32 = 30

	AkHttpUrl             int32 = 40
	AkHttpParam           int32 = 41
	AkHttpParamEntity     int32 = 42
	AkHttpCookie          int32 = 45
	AkHttpStatusCode      int32 = 46
	AkHttpInternalDisplay int32 = 48
	AkHttpIo              int32 = 49

	AkMsgQueueUri int32 = 100

	AkArgs0 int32 = -1
	AkArgs1 int32 = -2
	AkArgs2 int32 = -3
	AkArgs3 int32 = -4
	AkArgs4 int32 = -5
	AkArgs5 int32 = -6
	AkArgs6 int32 = -7
	AkArgs7 int32 = -8
	AkArgs8 int32 = -9
	AkArgs9 int32 = -10
	AkArgsn int32 = -11

	AkCacheArgs0 int32 = -30
	AkCacheArgs1 int32 = -31
	AkCacheArgs2 int32 = -32
	AkCacheArgs3 int32 = -33
	AkCacheArgs4 int32 = -34
	AkCacheArgs5 int32 = -35
	AkCacheArgs6 int32 = -36
	AkCacheArgs7 int32 = -37
	AkCacheArgs8 int32 = -38
	AkCacheArgs9 int32 = -39
	AkCacheArgsn int32 = -40
)

var (
	API         = AnnotationKeyFactory(12, "API")
	ApiMetadata = AnnotationKeyFactory(13, "API-METADATA")
	ReturnData  = AnnotationKeyFactory(14, "RETURN_DATA", ViewInRecordSet)
	ApiTag      = AnnotationKeyFactory(10015, "API-TAG")

	// when you don't know the correct cause of errors.
	ErrorapimetadataError = AnnotationKeyFactory(10000010, "API-METADATA-ERROR", ErrorApiMetadata)
	// when agentInfo not found
	ErrorapimetadataAgentInfoNotFound = AnnotationKeyFactory(10000011, "API-METADATA-AGENT-INFO-NOT-FOUND", ErrorApiMetadata)
	// when checksum is not correct even if agentInfo exists
	ErrorapimetadataIdentifierCheckError = AnnotationKeyFactory(10000012, "API-METADATA-IDENTIFIER-CHECK_ERROR", ErrorApiMetadata)
	// when  meta data itself not found
	ErrorapimetadataNotFound = AnnotationKeyFactory(10000013, "API-METADATA-NOT-FOUND", ErrorApiMetadata)
	// when the same hashId of meta data exists
	ErrorapimetadataDidCollsion = AnnotationKeyFactory(10000014, "API-METADATA-DID-COLLSION", ErrorApiMetadata)

	// it's not clear to handle a error code.  so ApiMetaDataError with searching ERROR_API_META_DATA has been used.
	// automatically generated id

	SqlId        = AnnotationKeyFactory(20, "SQL-ID")
	SQL          = AnnotationKeyFactory(21, "SQL", ViewInRecordSet)
	SqlMetadata  = AnnotationKeyFactory(22, "SQL-METADATA")
	SqlParam     = AnnotationKeyFactory(23, "SQL-PARAM")
	SqlBindvalue = AnnotationKeyFactory(24, "SQL-BindValue", ViewInRecordSet)

	StringId = AnnotationKeyFactory(30, "STRING_ID")

	// HTTP_URL is replaced by argument. So viewInRecordSet parameter name is not true.
	HttpUrl             = AnnotationKeyFactory(40, "http.url")
	HttpParam           = AnnotationKeyFactory(41, "http.param", ViewInRecordSet)
	HttpParamEntity     = AnnotationKeyFactory(42, "http.entity", ViewInRecordSet)
	HttpCookie          = AnnotationKeyFactory(45, "http.cookie", ViewInRecordSet)
	HttpStatusCode      = AnnotationKeyFactory(46, "http.status.code", ViewInRecordSet)
	HttpRequestHeader   = AnnotationKeyFactory(47, "http.req.header", ViewInRecordSet)
	HttpInternalDisplay = AnnotationKeyFactory(48, "http.internal.display")
	HttpIo              = AnnotationKeyFactory(49, "http.io", ViewInRecordSet)
	// post method parameter of httpclient

	MessageQueueUri = AnnotationKeyFactory(100, "message.queue.url")

	ARGS0 = AnnotationKeyFactory(-1, "args[0]")
	ARGS1 = AnnotationKeyFactory(-2, "args[1]")
	ARGS2 = AnnotationKeyFactory(-3, "args[2]")
	ARGS3 = AnnotationKeyFactory(-4, "args[3]")
	ARGS4 = AnnotationKeyFactory(-5, "args[4]")
	ARGS5 = AnnotationKeyFactory(-6, "args[5]")
	ARGS6 = AnnotationKeyFactory(-7, "args[6]")
	ARGS7 = AnnotationKeyFactory(-8, "args[7]")
	ARGS8 = AnnotationKeyFactory(-9, "args[8]")
	ARGS9 = AnnotationKeyFactory(-10, "args[9]")
	ARGSN = AnnotationKeyFactory(-11, "args[N]")

	CacheArgs0 = AnnotationKeyFactory(-30, "cached_args[0]")
	CacheArgs1 = AnnotationKeyFactory(-31, "cached_args[1]")
	CacheArgs2 = AnnotationKeyFactory(-32, "cached_args[2]")
	CacheArgs3 = AnnotationKeyFactory(-33, "cached_args[3]")
	CacheArgs4 = AnnotationKeyFactory(-34, "cached_args[4]")
	CacheArgs5 = AnnotationKeyFactory(-35, "cached_args[5]")
	CacheArgs6 = AnnotationKeyFactory(-36, "cached_args[6]")
	CacheArgs7 = AnnotationKeyFactory(-37, "cached_args[7]")
	CacheArgs8 = AnnotationKeyFactory(-38, "cached_args[8]")
	CacheArgs9 = AnnotationKeyFactory(-39, "cached_args[9]")
	CacheArgsn = AnnotationKeyFactory(-40, "cached_args[N]")
	//@Deprecated
	EXCEPTION = AnnotationKeyFactory(-50, "Exception", ViewInRecordSet)
	//@Deprecated
	ExceptionClass = AnnotationKeyFactory(-51, "ExceptionClass")
	UNKNOWN        = AnnotationKeyFactory(-9999, "UNKNOWN")

	ASYNC = AnnotationKeyFactory(-100, "Asynchronous Invocation", ViewInRecordSet)

	ProxyHttpHeader = AnnotationKeyFactory(300, "PROXY_HTTP_HEADER", ViewInRecordSet)
	RedisIo         = AnnotationKeyFactory(310, "redis.io")
)

type AnnotationKey struct {
	code             int
	name             string
	viewInRecordSet  bool
	errorApiMetadata bool
}
