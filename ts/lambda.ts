const lambdaFunctionUrl = "https://gon7npt4xmt4lomya5gqvsat4i0besqw.lambda-url.ap-south-1.on.aws/";

let executeLambda: nkruntime.RpcFunction = function (
    ctx: nkruntime.Context,
    logger: nkruntime.Logger,
    nk: nkruntime.Nakama,
    payload: string
): string | void {
    logger.debug("Ping received with payload: %s", payload);

    const body = JSON.parse(payload);
    const requestID = body.RequestID;
    const requestBody = JSON.parse(body.Body);


    //const requestBody = JSON.parse(payload);

    try {
        // Synchronously call the Lambda URL
        const response: nkruntime.HttpResponse = nk.httpRequest(
            lambdaFunctionUrl,
            "post",
            {
                "Content-Type": "application/json",
            },
            JSON.stringify(requestBody)
        );

        logger.info("Response from Lambda: %s", response.body);

        // const responseJson = {
        //     requestID: requestID,
        //     payload: {
        //       OpCode: "LambdaResponse",
        //       Message: response.body,
        //   }
        // }

        // Return the response body to the caller
        return JSON.stringify({
            requestID: requestID,
            payload: {
              OpCode: "LambdaResponse",
              Message: response.body,
          }
        });
    } catch (error:any) {
        logger.error("Error calling Lambda function: %s", error);

        // Return an error response
        return JSON.stringify({
            requestID: requestID,
            payload: {
              OpCode: "LambdaResponse",
              Message: error.message,
          }
        });
        // return JSON.stringify({
        //     success: false,
        //     error: error.message || "Unknown error"
        // });
    }
};
