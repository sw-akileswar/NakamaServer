FROM heroiclabs/nakama-pluginbuilder:3.23.0 AS builder

RUN apt-get update && apt-get install -y \
    curl \
    gnupg \
    lsb-release \
    && curl -fsSL https://deb.nodesource.com/setup_16.x | bash - \
    && apt-get install -y nodejs

ENV GO111MODULE on
ENV CGO_ENABLED 1
ENV GOPRIVATE "github.com/heroiclabs/nakama-project-template"

WORKDIR /backend
COPY go/ .
COPY go/vendor ./vendor

RUN go build --trimpath --mod=vendor --buildmode=plugin -o ./backend.so

COPY ts/ .
RUN npm install
RUN npx tsc

COPY lua/ .

FROM heroiclabs/nakama:3.23.0

COPY --from=builder /backend/backend.so /nakama/data/modules
COPY --from=builder /backend/*.lua /nakama/data/modules/
COPY --from=builder /backend/build/*.js /nakama/data/modules/build/
COPY --from=builder /backend/local.yml /nakama/data/
