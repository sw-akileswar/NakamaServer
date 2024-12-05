let rpcPing: nkruntime.RpcFunction = function (ctx: nkruntime.Context, logger: nkruntime.Logger, nk: nkruntime.Nakama, payload: string): string {
    logger.debug("Ping received with payload: %s", payload);
    return JSON.stringify("pong");
};

