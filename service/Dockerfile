FROM oven/bun:1 AS base
WORKDIR /app

FROM base AS install
RUN mkdir -p /temp/dev
COPY package.json bun.lockb /temp/dev/
COPY packages/api/package.json /temp/dev/packages/api/
COPY packages/core/package.json /temp/dev/packages/core/
RUN cd /temp/dev && bun install --frozen-lockfile

RUN mkdir -p /temp/prod
COPY package.json bun.lockb /temp/prod/
COPY packages/api/package.json /temp/prod/packages/api/
COPY packages/core/package.json /temp/prod/packages/core/
RUN cd /temp/prod && bun install --frozen-lockfile --production

FROM base AS prerelease
COPY --from=install /temp/dev/node_modules node_modules
COPY . .

FROM base AS release
COPY --from=install /temp/prod/node_modules node_modules
COPY --from=prerelease /app/packages ./packages/
COPY --from=prerelease /app/package.json .


RUN bun run db:migrate
# RUN chmod +rw /app && chmod +rw /app/db.sqlite

EXPOSE 3000/tcp
ENTRYPOINT [ "bun", "run", "packages/api/src/index.ts" ]