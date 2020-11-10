#build using docker
FROM mcr.microsoft.com/dotnet/core/sdk:3.1-alpine as builder
WORKDIR /home/build
COPY . ./
RUN dotnet publish SoSafe.Web -o /home/build/pub -f netcoreapp3.1

FROM mcr.microsoft.com/dotnet/core/aspnet:3.1-alpine
WORKDIR /opt/dotnet-cqrs
EXPOSE 80
COPY --from=builder /home/build/pub/ ./
COPY start-server.sh ./
CMD [ "/opt/dotnet-cqrs/start-server.sh" ]
