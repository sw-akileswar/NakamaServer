let rpcPing: nkruntime.RpcFunction = function (ctx: nkruntime.Context, logger: nkruntime.Logger, nk: nkruntime.Nakama, payload: string): string {
    logger.debug("Ping received with payload: %s", payload);

    const body = JSON.parse(payload);
    const requestID = body.RequestID;

    return JSON.stringify({
        requestID: requestID,
        payload: {
          OpCode: "PingResponse",
          Message: "pong",
      }
    });

    //return JSON.stringify("pong");
};

